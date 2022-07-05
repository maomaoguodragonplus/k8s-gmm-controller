package v1

import (
	gmm "github.com/maomaoguodragonplus/k8s-gmm-controller/pkg/apis/gmmcrd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var SchemaGroupVersion = schema.GroupVersion{
	Group:   github.com/maomaoguodragonplus/k8s-gmm-controller.GroupName,
	Version: github.com/maomaoguodragonplus/k8s-gmm-controller.Version,
}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

func Resource(resource string) schema.GroupResource {
	return SchemaGroupVersion.WithResource(resource).GroupResource()
}
func Kind(kind string) schema.GroupKind {
	return SchemaGroupVersion.WithKind(kind).GroupKind()
}
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(
		SchemaGroupVersion,
		&Database{},
		&DatabaseList{},
	)
	metav1.AddToGroupVersion(scheme, SchemaGroupVersion)
	return nil
}
