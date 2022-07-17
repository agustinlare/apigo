package main

import (
	"context"
	"log"
	"path/filepath"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func newPod() {
	cfg, err := clientcmd.BuildConfigFromFlags(
		"",
		filepath.Join(homedir.HomeDir(), ".kube", "config"),
	)

	if err != nil {
		log.Println(err)
	}

	k8s, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Println(err)
	}

	pod := &core.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-pod",
			Namespace: "default",
			Labels: map[string]string{
				"app": "test-pod",
			},
		},
		Spec: core.PodSpec{
			InitContainers: []core.Container{
				{
					Name:            "awscli ",
					Image:           "amazon/aws-cli",
					ImagePullPolicy: core.PullIfNotPresent,
					Command: []string{
						"sleep",
						"3600",
					},
				},
			},
			Containers: []core.Container{
				{
					Name:            "busybox",
					Image:           "busybox",
					ImagePullPolicy: core.PullIfNotPresent,
					Command: []string{
						"sleep",
						"3600",
					},
				},
			},
		},
	}

	_, err = k8s.CoreV1().Pods("default").Create(
		context.Background(),
		pod,
		metav1.CreateOptions{},
	)

	if err != nil {
		log.Println(err)
	}

}
