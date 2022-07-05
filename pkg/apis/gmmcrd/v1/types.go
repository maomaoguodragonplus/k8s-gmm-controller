package v1


import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Database describes a database.
// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Database struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec DatabaseSpec `json:"spec"`
}

// DatabaseSpec is the spec for a Foo resource
type DatabaseSpec struct {
	Image    string `json:"image"`
	Replicas int32  `json:"replicas"`
}

// DatabaseList is a list of Database resources
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type DatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Database `json:"items"`
}
