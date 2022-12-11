package namespace

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/json"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"

	"context"
)

func dataConversion(item *corev1.Namespace) *Namespace {
	ns := &Namespace{
		Name: item.Name,
	}

	for k, v := range item.GetLabels() {
		ns.Label = append(ns.Label, &Label{
			Key:   k,
			Value: v,
		})
	}
	return ns
}

func (x *NamespaceSet) Add(item *corev1.Namespace) {
	x.Namespace = append(x.Namespace, dataConversion(item))
}

var label = make(map[string]interface{}, 0)

func PatchNamespaceLabel(ctx context.Context, client v1.CoreV1Interface, name string, patchData map[string]interface{}) (
	*corev1.Namespace, error) {

	playLoadBytes, _ := json.Marshal(patchData)
	newNs, err := client.Namespaces().Patch(ctx, name, types.StrategicMergePatchType, playLoadBytes, metav1.PatchOptions{})

	if err != nil {
		return nil, err
	}

	return newNs, nil
}
