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

package v1

import (
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GuestbookSpec defines the desired state of Guestbook
type GuestbookSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example fielkd of Guestbook. Edit Guestbook_types.go to remove/update
	Foo      string               `json:"foo,omitempty"`
	Replicas int                  `json:"replicas"`
	Selector metav1.LabelSelector `json:"selector"`
	Template Template             `json:"template"`
}

type Template struct {
	Spec Spec `json:"spec"`
}

type Spec struct {
	Volumes                       []corev1.Volume                   `json:"volumes,omitempty" patchStrategy:"merge,retainKeys" patchMergeKey:"name" protobuf:"bytes,1,rep,name=volumes"`
	InitContainers                []Container                       `json:"initContainers,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,20,rep,name=initContainers"`
	Containers                    []Container                       `json:"containers" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,2,rep,name=containers"`
	EphemeralContainers           []corev1.EphemeralContainer       `json:"ephemeralContainers,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,34,rep,name=ephemeralContainers"`
	RestartPolicy                 corev1.RestartPolicy              `json:"restartPolicy,omitempty" protobuf:"bytes,3,opt,name=restartPolicy,casttype=RestartPolicy"`
	TerminationGracePeriodSeconds *int64                            `json:"terminationGracePeriodSeconds,omitempty" protobuf:"varint,4,opt,name=terminationGracePeriodSeconds"`
	ActiveDeadlineSeconds         *int64                            `json:"activeDeadlineSeconds,omitempty" protobuf:"varint,5,opt,name=activeDeadlineSeconds"`
	DNSPolicy                     corev1.DNSPolicy                  `json:"dnsPolicy,omitempty" protobuf:"bytes,6,opt,name=dnsPolicy,casttype=DNSPolicy"`
	NodeSelector                  map[string]string                 `json:"nodeSelector,omitempty" protobuf:"bytes,7,rep,name=nodeSelector"`
	ServiceAccountName            string                            `json:"serviceAccountName,omitempty" protobuf:"bytes,8,opt,name=serviceAccountName"`
	DeprecatedServiceAccount      string                            `json:"serviceAccount,omitempty" protobuf:"bytes,9,opt,name=serviceAccount"`
	AutomountServiceAccountToken  *bool                             `json:"automountServiceAccountToken,omitempty" protobuf:"varint,21,opt,name=automountServiceAccountToken"`
	NodeName                      string                            `json:"nodeName,omitempty" protobuf:"bytes,10,opt,name=nodeName"`
	HostNetwork                   bool                              `json:"hostNetwork,omitempty" protobuf:"varint,11,opt,name=hostNetwork"`
	HostPID                       bool                              `json:"hostPID,omitempty" protobuf:"varint,12,opt,name=hostPID"`
	HostIPC                       bool                              `json:"hostIPC,omitempty" protobuf:"varint,13,opt,name=hostIPC"`
	ShareProcessNamespace         *bool                             `json:"shareProcessNamespace,omitempty" protobuf:"varint,27,opt,name=shareProcessNamespace"`
	SecurityContext               *corev1.PodSecurityContext        `json:"securityContext,omitempty" protobuf:"bytes,14,opt,name=securityContext"`
	ImagePullSecrets              []corev1.LocalObjectReference     `json:"imagePullSecrets,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,15,rep,name=imagePullSecrets"`
	Hostname                      string                            `json:"hostname,omitempty" protobuf:"bytes,16,opt,name=hostname"`
	Subdomain                     string                            `json:"subdomain,omitempty" protobuf:"bytes,17,opt,name=subdomain"`
	Affinity                      *corev1.Affinity                  `json:"affinity,omitempty" protobuf:"bytes,18,opt,name=affinity"`
	SchedulerName                 string                            `json:"schedulerName,omitempty" protobuf:"bytes,19,opt,name=schedulerName"`
	Tolerations                   []corev1.Toleration               `json:"tolerations,omitempty" protobuf:"bytes,22,opt,name=tolerations"`
	HostAliases                   []corev1.HostAlias                `json:"hostAliases,omitempty" patchStrategy:"merge" patchMergeKey:"ip" protobuf:"bytes,23,rep,name=hostAliases"`
	PriorityClassName             string                            `json:"priorityClassName,omitempty" protobuf:"bytes,24,opt,name=priorityClassName"`
	Priority                      *int32                            `json:"priority,omitempty" protobuf:"bytes,25,opt,name=priority"`
	DNSConfig                     *corev1.PodDNSConfig              `json:"dnsConfig,omitempty" protobuf:"bytes,26,opt,name=dnsConfig"`
	ReadinessGates                []corev1.PodReadinessGate         `json:"readinessGates,omitempty" protobuf:"bytes,28,opt,name=readinessGates"`
	RuntimeClassName              *string                           `json:"runtimeClassName,omitempty" protobuf:"bytes,29,opt,name=runtimeClassName"`
	EnableServiceLinks            *bool                             `json:"enableServiceLinks,omitempty" protobuf:"varint,30,opt,name=enableServiceLinks"`
	PreemptionPolicy              *corev1.PreemptionPolicy          `json:"preemptionPolicy,omitempty" protobuf:"bytes,31,opt,name=preemptionPolicy"`
	Overhead                      corev1.ResourceList               `json:"overhead,omitempty" protobuf:"bytes,32,opt,name=overhead"`
	TopologySpreadConstraints     []corev1.TopologySpreadConstraint `json:"topologySpreadConstraints,omitempty" patchStrategy:"merge" patchMergeKey:"topologyKey" protobuf:"bytes,33,opt,name=topologySpreadConstraints"`
}

type Container struct {
	Name                     string                          `json:"name" protobuf:"bytes,1,opt,name=name"`
	Image                    string                          `json:"image,omitempty" protobuf:"bytes,2,opt,name=image"`
	Command                  []string                        `json:"command,omitempty" protobuf:"bytes,3,rep,name=command"`
	Args                     []string                        `json:"args,omitempty" protobuf:"bytes,4,rep,name=args"`
	WorkingDir               string                          `json:"workingDir,omitempty" protobuf:"bytes,5,opt,name=workingDir"`
	Ports                    []ContainerPort                 `json:"ports,omitempty" patchStrategy:"merge" patchMergeKey:"containerPort" protobuf:"bytes,6,rep,name=ports"`
	EnvFrom                  []corev1.EnvFromSource          `json:"envFrom,omitempty" protobuf:"bytes,19,rep,name=envFrom"`
	Env                      []corev1.EnvVar                 `json:"env,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,7,rep,name=env"`
	Resources                corev1.ResourceRequirements     `json:"resources,omitempty" protobuf:"bytes,8,opt,name=resources"`
	VolumeMounts             []corev1.VolumeMount            `json:"volumeMounts,omitempty" patchStrategy:"merge" patchMergeKey:"mountPath" protobuf:"bytes,9,rep,name=volumeMounts"`
	VolumeDevices            []corev1.VolumeDevice           `json:"volumeDevices,omitempty" patchStrategy:"merge" patchMergeKey:"devicePath" protobuf:"bytes,21,rep,name=volumeDevices"`
	LivenessProbe            *corev1.Probe                   `json:"livenessProbe,omitempty" protobuf:"bytes,10,opt,name=livenessProbe"`
	ReadinessProbe           *corev1.Probe                   `json:"readinessProbe,omitempty" protobuf:"bytes,11,opt,name=readinessProbe"`
	StartupProbe             *corev1.Probe                   `json:"startupProbe,omitempty" protobuf:"bytes,22,opt,name=startupProbe"`
	Lifecycle                *corev1.Lifecycle               `json:"lifecycle,omitempty" protobuf:"bytes,12,opt,name=lifecycle"`
	TerminationMessagePath   string                          `json:"terminationMessagePath,omitempty" protobuf:"bytes,13,opt,name=terminationMessagePath"`
	TerminationMessagePolicy corev1.TerminationMessagePolicy `json:"terminationMessagePolicy,omitempty" protobuf:"bytes,20,opt,name=terminationMessagePolicy,casttype=TerminationMessagePolicy"`
	ImagePullPolicy          corev1.PullPolicy               `json:"imagePullPolicy,omitempty" protobuf:"bytes,14,opt,name=imagePullPolicy,casttype=PullPolicy"`
	SecurityContext          *corev1.SecurityContext         `json:"securityContext,omitempty" protobuf:"bytes,15,opt,name=securityContext"`
	Stdin                    bool                            `json:"stdin,omitempty" protobuf:"varint,16,opt,name=stdin"`
	StdinOnce                bool                            `json:"stdinOnce,omitempty" protobuf:"varint,17,opt,name=stdinOnce"`
	TTY                      bool                            `json:"tty,omitempty" protobuf:"varint,18,opt,name=tty"`
}

type ContainerPort struct {
	// If specified, this must be an IANA_SVC_NAME and unique within the pod. Each
	// named port in a pod must have a unique name. Name for the port that can be
	// referred to by services.
	// +optional
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	// Number of port to expose on the host.
	// If specified, this must be a valid port number, 0 < x < 65536.
	// If HostNetwork is specified, this must match ContainerPort.
	// Most containers do not need this.
	// +optional
	HostPort int32 `json:"hostPort,omitempty" protobuf:"varint,2,opt,name=hostPort"`
	// Number of port to expose on the pod's IP address.
	// This must be a valid port number, 0 < x < 65536.
	ContainerPort int32 `json:"containerPort" protobuf:"varint,3,opt,name=containerPort"`
	// Protocol for port. Must be UDP, TCP, or SCTP.
	// Defaults to "TCP".
	// +optional
	Protocol corev1.Protocol `json:"protocol,omitempty" protobuf:"bytes,4,opt,name=protocol,casttype=Protocol"`
	// What host IP to bind the external port to.
	// +optional
	HostIP string `json:"hostIP,omitempty" protobuf:"bytes,5,opt,name=hostIP"`
}

// GuestbookStatus defines the observed state of Guestbook
type GuestbookStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Status            v1.DeploymentStatus `json:"status"`
	AvailableReplicas int                 `json:"availableReplicas"` //当前实际的副本数
}

// +kubebuilder:object:root=true

// Guestbook is the Schema for the guestbooks API
type Guestbook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	//dev 该了啊 laile
	Spec   GuestbookSpec   `json:"spec,omitempty"`
	Status GuestbookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GuestbookList contains a list of Guestbook
type GuestbookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Guestbook `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Guestbook{}, &GuestbookList{})
}
