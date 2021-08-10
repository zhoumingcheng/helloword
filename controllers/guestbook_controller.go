/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	"helloword/utils"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	goruntime "runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"time"

	wepappv1 "helloword/api/v1"
)

// GuestbookReconciler reconciles a Guestbook object
type GuestbookReconciler struct {
	client.Client
	Log      logr.Logger
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

// +kubebuilder:rbac:groups=wepapp.com.bolingcavalry,resources=guestbooks,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=wepapp.com.bolingcavalry,resources=guestbooks/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=batch,resources=pod,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=batch,resources=pod/status,verbs=get
// +kubebuilder:rbac:groups="",resources=events,verbs=create;patch

func (r *GuestbookReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("guestbook", req.NamespacedName)

	// your logic here
	r.Log.Info(fmt.Sprintf("1.%v", req))
	r.Log.Info(fmt.Sprintf("2.%v", goruntime.NumGoroutine()))
	//这里的guest实现一个简单的类似deployment的副本控制
	//按照namespace加载guestbook
	guestbook := &wepappv1.Guestbook{}

	err := r.Get(ctx, req.NamespacedName, guestbook)
	if err != nil {
		//如果没有查询到实例则直接返回
		if errors.IsNotFound(err) {
			log.Info("未查询到实例！")
			return ctrl.Result{}, nil
		}
		log.Error(err, "")
		return ctrl.Result{}, err
	}
	//列出所有的属于该guestbook下的pod
	podList := &corev1.PodList{}
	if err := r.List(ctx, podList, client.InNamespace(req.Namespace),
		client.MatchingLabels{"app": guestbook.Spec.Selector.MatchLabels["app"]}); err != nil {
		log.Error(err, "unable to list child Pods")
		return ctrl.Result{}, err
	}

	//判断pod数量是否和replicas相同
	realityReplicas := len(podList.Items)
	if realityReplicas != guestbook.Spec.Replicas {
		//添加event信息表明实际replicas与期望不同，正在达成期望
		r.Recorder.Eventf(guestbook, corev1.EventTypeNormal, "Scheduled", "正在尝试调整pod数达到副本预期数")
		if realityReplicas < guestbook.Spec.Replicas {
			//当实际pod的副本书小于guestbook期望的副本书时创建新的pod
			for i := 0; i < guestbook.Spec.Replicas-realityReplicas; i++ {
				pod := utils.GetPod(guestbook)
				if err := r.Create(ctx, pod); err != nil {
					log.Error(err, "创建pod失败！")
					return ctrl.Result{}, err
				}
			}
			r.Recorder.Eventf(guestbook, corev1.EventTypeNormal, "Created", "Created pod success")
		} else {
			//当实际pod副本书大于guestbook期望副本书时删除多余的pod，当前根据pod创建先后进行删除 先>后
			deletePodNumber := realityReplicas - guestbook.Spec.Replicas
			for _, item := range podList.Items {
				pod := &corev1.Pod{
					TypeMeta: metav1.TypeMeta{},
					ObjectMeta: metav1.ObjectMeta{
						Name:      item.Name,
						Namespace: item.Namespace,
					},
				}
				if err := r.Delete(ctx, pod); err != nil {
					return ctrl.Result{}, err
				}
				deletePodNumber--
				if deletePodNumber <= 0 {
					break
				}
			}
			r.Recorder.Eventf(guestbook, corev1.EventTypeNormal, "Delete", "Delete pod success")
		}
	}
	//更新guestbook当前的状态
	guestbook.Status.AvailableReplicas = guestbook.Spec.Replicas
	if err := r.Update(ctx, guestbook); err != nil {
		log.Error(err, "更新状态失败")
		return ctrl.Result{}, err
	}
	// 声明 finalizer 字段，类型为字符串
	myFinalizerName := "storage.finalizers.tutorial.kubebuilder.io"

	// 通过检查 DeletionTimestamp 字段是否为0 判断资源是否被删除
	if guestbook.ObjectMeta.DeletionTimestamp.IsZero() {
		// 如果为0 ，则资源未被删除，我们需要检测是否存在 finalizer，如果不存在，则添加，并更新到资源对象中
		if !containsString(guestbook.ObjectMeta.Finalizers, myFinalizerName) {
			guestbook.ObjectMeta.Finalizers = append(guestbook.ObjectMeta.Finalizers, myFinalizerName)
			if err := r.Update(context.Background(), guestbook); err != nil {
				return ctrl.Result{}, err
			}
		}
	} else {
		// 如果不为 0 ，则对象处于删除中
		if containsString(guestbook.ObjectMeta.Finalizers, myFinalizerName) {
			// 如果存在 finalizer 且与上述声明的 finalizer 匹配，那么执行对应 hook 逻辑
			if err := r.deleteExternalResources(ctx, guestbook); err != nil {
				// 如果删除失败，则直接返回对应 err，controller 会自动执行重试逻辑
				return ctrl.Result{}, err
			}

			// 如果对应 hook 执行成功，那么清空 finalizers， k8s 删除对应资源
			guestbook.ObjectMeta.Finalizers = removeString(guestbook.ObjectMeta.Finalizers, myFinalizerName)
			if err := r.Update(context.Background(), guestbook); err != nil {
				return ctrl.Result{}, err
			}
		}

		return ctrl.Result{}, err
	}
	return ctrl.Result{RequeueAfter: time.Second}, nil
}

func (r *GuestbookReconciler) deleteExternalResources(ctx context.Context, guestbook *wepappv1.Guestbook) error {
	// 删除 guestbook关联的pods
	pod := &corev1.Pod{}
	if err := r.DeleteAllOf(ctx, pod, client.InNamespace(guestbook.Namespace),
		client.MatchingLabels{"app": guestbook.Spec.Selector.MatchLabels["app"]}); err != nil {
		r.Log.Error(err, "删除guestbook关联pod失败")
		return err
	}
	return nil
}

func containsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func removeString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}

func (r *GuestbookReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&wepappv1.Guestbook{}).
		Complete(r)
}
