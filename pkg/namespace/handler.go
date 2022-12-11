package namespace

import (
	corev1 "k8s.io/api/core/v1"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) OnAdd(obj interface{}) {
	if namespace, ok := obj.(*corev1.Namespace); ok {
		MapStorage.Add(namespace)
	}

}

func (h *Handler) OnUpdate(oldObj interface{}, newObj interface{}) {
	if namespace, ok := newObj.(*corev1.Namespace); ok {
		MapStorage.Update(namespace)
	}

}

func (h *Handler) OnDelete(obj interface{}) {
	if namespace, ok := obj.(*corev1.Namespace); ok {
		MapStorage.Delete(namespace)
	}
}
