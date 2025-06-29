package keel

import (
    "flag"
    "path/filepath"
	"os"
	"fmt"
	"sigs.k8s.io/yaml"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
)

type Client struct {
    KubeClient *kubernetes.Clientset
}

func NewClient() (*Client, error) {
    var kubeconfig *string
    if home := homeDir(); home != "" {
        kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
    } else {
        kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
    }
    flag.Parse()

    config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
    if err != nil {
        return nil, err
    }

    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        return nil, err
    }

    return &Client{KubeClient: clientset}, nil
}

func homeDir() string {
    if h := os.Getenv("HOME"); h != "" {
        return h
    }
    return os.Getenv("USERPROFILE") // for Windows
}

// MarshalYAML marshals a Kubernetes object into YAML bytes.
func MarshalYAML(obj interface{}) ([]byte, error) {
    return yaml.Marshal(obj)
}

// MarshalYAMLToStdout marshals the object and prints it to stdout.
func MarshalYAMLToStdout(obj interface{}) error {
    out, err := MarshalYAML(obj)
    if err != nil {
        return err
    }
    fmt.Println("---")
    fmt.Println(string(out))
    return nil
}
