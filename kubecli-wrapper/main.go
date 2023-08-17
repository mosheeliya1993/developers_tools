package main

import (
	"context"
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func getKubeClient() (*kubernetes.Clientset, error) {
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		kubeconfig = os.Getenv("HOME") + "/.kube/config"
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

func listPods(namespace string) {
	clientset, err := getKubeClient()
	if err != nil {
		fmt.Println("Error creating Kubernetes client:", err)
		return
	}

	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error listing pods:", err)
		return
	}

	for _, pod := range pods.Items {
		fmt.Println("Pod:", pod.Name)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: kubecli-wrapper <command> [args]")
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "list-pods":
		namespace := "default" // Default namespace
		if len(args) > 0 {
			namespace = args[0]
		}
		listPods(namespace)
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}
}
