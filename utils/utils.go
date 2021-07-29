package utils

import (
	wepappv1 "helloword/api/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
	"time"
)

func GetPod(guestbook *wepappv1.Guestbook) *corev1.Pod {
	initContainers := make([]corev1.Container, 0)
	for _, InitContainer := range guestbook.Spec.Template.Spec.InitContainers {
		ports := make([]corev1.ContainerPort, 0)
		for _, port := range InitContainer.Ports {
			ports = append(ports, corev1.ContainerPort{
				Name:          port.Name,
				HostPort:      port.HostPort,
				ContainerPort: port.ContainerPort,
				Protocol:      port.Protocol,
				HostIP:        port.HostIP,
			})
		}
		initContainers = append(initContainers, corev1.Container{
			Name:                     InitContainer.Name,
			Image:                    InitContainer.Image,
			Command:                  InitContainer.Command,
			Args:                     InitContainer.Args,
			WorkingDir:               InitContainer.WorkingDir,
			Ports:                    ports,
			EnvFrom:                  InitContainer.EnvFrom,
			Env:                      InitContainer.Env,
			Resources:                InitContainer.Resources,
			VolumeMounts:             InitContainer.VolumeMounts,
			VolumeDevices:            InitContainer.VolumeDevices,
			LivenessProbe:            InitContainer.LivenessProbe,
			ReadinessProbe:           InitContainer.ReadinessProbe,
			StartupProbe:             InitContainer.StartupProbe,
			Lifecycle:                InitContainer.Lifecycle,
			TerminationMessagePath:   InitContainer.TerminationMessagePath,
			TerminationMessagePolicy: InitContainer.TerminationMessagePolicy,
			ImagePullPolicy:          InitContainer.ImagePullPolicy,
			SecurityContext:          InitContainer.SecurityContext,
			Stdin:                    InitContainer.Stdin,
			StdinOnce:                InitContainer.StdinOnce,
			TTY:                      InitContainer.TTY,
		})
	}
	containers := make([]corev1.Container, 0)
	for _, InitContainer := range guestbook.Spec.Template.Spec.Containers {
		ports := make([]corev1.ContainerPort, 0)
		for _, port := range InitContainer.Ports {
			ports = append(ports, corev1.ContainerPort{
				Name:          port.Name,
				HostPort:      port.HostPort,
				ContainerPort: port.ContainerPort,
				Protocol:      port.Protocol,
				HostIP:        port.HostIP,
			})
		}
		containers = append(containers, corev1.Container{
			Name:                     InitContainer.Name,
			Image:                    InitContainer.Image,
			Command:                  InitContainer.Command,
			Args:                     InitContainer.Args,
			WorkingDir:               InitContainer.WorkingDir,
			Ports:                    ports,
			EnvFrom:                  InitContainer.EnvFrom,
			Env:                      InitContainer.Env,
			Resources:                InitContainer.Resources,
			VolumeMounts:             InitContainer.VolumeMounts,
			VolumeDevices:            InitContainer.VolumeDevices,
			LivenessProbe:            InitContainer.LivenessProbe,
			ReadinessProbe:           InitContainer.ReadinessProbe,
			StartupProbe:             InitContainer.StartupProbe,
			Lifecycle:                InitContainer.Lifecycle,
			TerminationMessagePath:   InitContainer.TerminationMessagePath,
			TerminationMessagePolicy: InitContainer.TerminationMessagePolicy,
			ImagePullPolicy:          InitContainer.ImagePullPolicy,
			SecurityContext:          InitContainer.SecurityContext,
			Stdin:                    InitContainer.Stdin,
			StdinOnce:                InitContainer.StdinOnce,
			TTY:                      InitContainer.TTY,
		})
	}
	pod := &corev1.Pod{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      guestbook.Name + "-" + strconv.FormatInt(time.Now().UnixNano(), 10),
			Namespace: guestbook.Namespace,
			Labels:    guestbook.Spec.Selector.MatchLabels,
		},
		Spec: corev1.PodSpec{
			Volumes:                       guestbook.Spec.Template.Spec.Volumes,
			InitContainers:                initContainers,
			Containers:                    containers,
			EphemeralContainers:           guestbook.Spec.Template.Spec.EphemeralContainers,
			RestartPolicy:                 guestbook.Spec.Template.Spec.RestartPolicy,
			TerminationGracePeriodSeconds: guestbook.Spec.Template.Spec.TerminationGracePeriodSeconds,
			ActiveDeadlineSeconds:         guestbook.Spec.Template.Spec.ActiveDeadlineSeconds,
			DNSPolicy:                     guestbook.Spec.Template.Spec.DNSPolicy,
			NodeSelector:                  guestbook.Spec.Template.Spec.NodeSelector,
			ServiceAccountName:            guestbook.Spec.Template.Spec.ServiceAccountName,
			DeprecatedServiceAccount:      guestbook.Spec.Template.Spec.DeprecatedServiceAccount,
			AutomountServiceAccountToken:  guestbook.Spec.Template.Spec.AutomountServiceAccountToken,
			NodeName:                      guestbook.Spec.Template.Spec.NodeName,
			HostNetwork:                   guestbook.Spec.Template.Spec.HostNetwork,
			HostPID:                       guestbook.Spec.Template.Spec.HostPID,
			HostIPC:                       guestbook.Spec.Template.Spec.HostIPC,
			ShareProcessNamespace:         guestbook.Spec.Template.Spec.ShareProcessNamespace,
			SecurityContext:               guestbook.Spec.Template.Spec.SecurityContext,
			ImagePullSecrets:              guestbook.Spec.Template.Spec.ImagePullSecrets,
			Hostname:                      guestbook.Spec.Template.Spec.Hostname,
			Subdomain:                     guestbook.Spec.Template.Spec.Subdomain,
			Affinity:                      guestbook.Spec.Template.Spec.Affinity,
			SchedulerName:                 guestbook.Spec.Template.Spec.SchedulerName,
			Tolerations:                   guestbook.Spec.Template.Spec.Tolerations,
			HostAliases:                   guestbook.Spec.Template.Spec.HostAliases,
			PriorityClassName:             guestbook.Spec.Template.Spec.PriorityClassName,
			Priority:                      guestbook.Spec.Template.Spec.Priority,
			DNSConfig:                     guestbook.Spec.Template.Spec.DNSConfig,
			ReadinessGates:                guestbook.Spec.Template.Spec.ReadinessGates,
			RuntimeClassName:              guestbook.Spec.Template.Spec.RuntimeClassName,
			EnableServiceLinks:            guestbook.Spec.Template.Spec.EnableServiceLinks,
			PreemptionPolicy:              guestbook.Spec.Template.Spec.PreemptionPolicy,
			Overhead:                      guestbook.Spec.Template.Spec.Overhead,
			TopologySpreadConstraints:     guestbook.Spec.Template.Spec.TopologySpreadConstraints,
		},
		Status: corev1.PodStatus{},
	}
	return pod
}
