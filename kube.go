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

func checkErr(e error) {
	if e != nil {
		log.Println(e)
	}
}

func hello() {
	cfg, err := clientcmd.BuildConfigFromFlags(
		"",
		filepath.Join(homedir.HomeDir(), ".kube", "config"),
	)
	checkErr(err)

	k8s, err := kubernetes.NewForConfig(cfg)
	checkErr(err)

	// nsList, err := k8s.CoreV1().
	// 	Namespaces().
	// 	List(context.Background(), metav1.ListOptions{})
	// checkErr(err)

	// for _, n := range nsList.Items {
	// 	fmt.Println(n.Name)
	// }

	pod := &core.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-pod",
			Namespace: "default",
			Labels: map[string]string{
				"app": "test-pod",
			},
		},
		Spec: core.PodSpec{
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

	checkErr(err)

}
