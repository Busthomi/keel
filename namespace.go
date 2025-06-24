package keel

import (
    corev1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NamespaceManager struct{}

func NewNamespaceManager() *NamespaceManager {
    return &NamespaceManager{}
}

// Build creates a namespace object (but doesn't apply it to Kubernetes).
func (n *NamespaceManager) Build(name string) *corev1.Namespace {
    return &corev1.Namespace{
        TypeMeta: metav1.TypeMeta{
            Kind:       "Namespace",
            APIVersion: "v1",
        },
        ObjectMeta: metav1.ObjectMeta{
            Name: name,
        },
    }
}
