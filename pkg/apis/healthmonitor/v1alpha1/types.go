package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DeploymentHealthMonitor monitors the health of all deployments in the given namespace and reports them as statuses.
// It is a specification of the deploymenthealthmonitor resource.
type DeploymentHealthMonitor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HealthMonitorSpec   `json:"spec"`
	Status HealthMonitorStatus `json:"status"`
}

// Generic type as opposed to deployment specific because in future if we define sts resource etc,. Spec and status types can be reused.
type HealthMonitorSpec struct {
	Namespace string `json:"namespace"`
}

type HealthMonitorStatus struct {
	Results []HealthMonitorStatus `json:"results,omitempty"`
}

// Reports status of a workload resource. Status is roughly 2 parts:
// - When was this last upgraded?
//   - Is the upgrade currently failing?
//
// - Is the resource currently unhealthy?
//   - when did we last observe unhealthiness?
type HealthMonitorResourceResult struct {
	Name string `json:"name"`

	ObservedGeneration   int64       `json:"observedGeneration,omitempty"`
	LastUpgradeStartTime metav1.Time `json:"lastUpgradeStartTime,omitempty"`
	LastUpgradeEndTime   metav1.Time `json:"lastUpgradeEndTime,omitempty"`
	UpgradeFailureReason string      `json:"upgradeFailureReason,omitempty"`

	IsUnhealthy                   bool            `json:"isUnhealthy,omitempty"`
	LastObservedUnhealthyTime     metav1.Time     `json:"lastObservedUnhealthy,omitempty"`
	LastObservedUnhealthyDuration metav1.Duration `json:"lastObservedUnhealthyDuration,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DeploymentHealthMonitorList is a list of DeploymentHealthMonitor resources
type DeploymentHealthMonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []DeploymentHealthMonitor `json:"items"`
}
