package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	pipelines "github.com/PI-Victor/pipelines/pkg/apis/cloudflavor.io/v1"
)

var SchemeGroupVersion = schema.GroupVersion{
	Group:   pipelines.GroupName,
	Version: "v1",
}
