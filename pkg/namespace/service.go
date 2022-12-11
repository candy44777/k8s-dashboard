package namespace

import (
	"context"
	"github.com/candy44777/k8s-dashboard/conf"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Service struct {
	UnimplementedServiceServer
	Config *Config
}

func NewService() *Service {
	return &Service{
		Config: NewConfig(),
	}
}

func (s *Service) QueryAll(ctx context.Context, req *RequestByNamespace) (*NamespaceSet, error) {
	ns := &NamespaceSet{}
	nl, err := conf.K8sCLientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, item := range nl.Items {
		ns.Add(&item)
	}
	return ns, nil
}

func (s *Service) GetByName(ctx context.Context, req *RequestByNamespace) (*Namespace, error) {
	ns := &corev1.Namespace{}
	ns.Name = req.Name
	ret := MapStorage.Get(ns)
	return dataConversion(ret), nil
}

func (s *Service) Delete(ctx context.Context, req *RequestByNamespace) (*Namespace, error) {
	ns := &corev1.Namespace{}
	ns.Name = req.Name
	if err := MapStorage.Delete(ns); err != nil {
		return nil, err
	}

	conf.K8sCLientset.CoreV1().Namespaces().Delete(ctx, req.Name, metav1.DeleteOptions{})
	return dataConversion(ns), nil
}

func (s *Service) Create(ctx context.Context, req *Namespace) (*Namespace, error) {
	ns := &corev1.Namespace{}
	ns.Name = req.Name
	if req.Label != nil {
		ns.Labels = map[string]string{
			req.Label[0].Key: req.Label[0].Value,
		}
	}

	newNs, err := conf.K8sCLientset.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return dataConversion(newNs), nil
}

func (s *Service) OverwriteLabel(ctx context.Context, req *RequestByLabel) (*Namespace, error) {

	if req.Label != nil {
		for _, v := range req.Label {
			label = map[string]interface{}{
				"metadata": map[string]map[string]string{
					"labels": {
						v.Key: v.Value,
					},
				},
			}
		}
	}

	ns, err := PatchNamespaceLabel(ctx, conf.K8sCLientset.CoreV1(), req.Namespace.Name, label)
	if err != nil {
		return nil, err
	}

	return dataConversion(ns), nil
}
