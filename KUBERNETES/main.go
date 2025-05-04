package main

import (
	"context"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func 

func main() {
	masterURL := "https://127.0.0.1:39467"
	kubecConfigPath := "/home/orkhan/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags(masterURL, kubecConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	listOfPods, err := clientSet.CoreV1().Pods("devops-lab").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range listOfPods.Items {
		if item.Status.Phase != "Running" && item.Status.Phase != "Pending" {
			fmt.Println("Selected Pod : ", item.Name)
			clientSet.CoreV1().Pods("devops-lab").Delete(context.Background(), item.Name, metav1.DeleteOptions{})
		}
	}



}
