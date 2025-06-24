package keel

import (
    "context"
    "fmt"

    corev1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NamespaceManager struct {
    client *Client
}

func NewNamespaceManager(client *Client) *NamespaceManager {
    return &NamespaceManager{client: client}
}

func (n *NamespaceManager) Create(name string) (*corev1.Namespace, error) {
    ns := &corev1.Namespace{
        ObjectMeta: metav1.ObjectMeta{
            Name: name,
        },
    }

    return n.client.KubeClient.CoreV1().Namespaces().Create(context.TODO(), ns, metav1.CreateOptions{})
}

func (n *NamespaceManager) Get(name string) (*corev1.Namespace, error) {
    return n.client.KubeClient.CoreV1().Namespaces().Get(context.TODO(), name, metav1.GetOptions{})
}

func (n *NamespaceManager) List() ([]corev1.Namespace, error) {
    nsList, err := n.client.KubeClient.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
    if err != nil {
        return nil, err
    }
    return nsList.Items, nil
}

func (n *NamespaceManager) Delete(name string) error {
    return n.client.KubeClient.CoreV1().Namespaces().Delete(context.TODO(), name, metav1.DeleteOptions{})
}
