package namespace

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"sync"
)

type list []*corev1.Namespace

func (l list) Len() int {
	return len(l)
}

func (l list) Less(i, j int) bool {
	return l[i].CreationTimestamp.Time.After(l[j].CreationTimestamp.Time)
}

func (l list) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

var MapStorage = NewMapStorage()

type mapStorage struct {
	data sync.Map
}

func NewMapStorage() *mapStorage {
	return &mapStorage{}
}

func (m *mapStorage) Add(itme *corev1.Namespace) {
	m.data.Store(itme.Name, itme)
}

func (m *mapStorage) Update(itme *corev1.Namespace) {
	m.data.Store(itme.Name, itme)
}

func (m *mapStorage) Delete(itme *corev1.Namespace) error {
	if ns, ok := m.data.Load(itme.Name); ok {
		m.data.Delete(ns)
		return nil
	}
	return fmt.Errorf("resource %s is not", itme.Name)
}

func (m *mapStorage) Get(itme *corev1.Namespace) *corev1.Namespace {
	if ns, ok := m.data.Load(itme.Name); ok {
		return ns.(*corev1.Namespace)
	}
	return nil
}
