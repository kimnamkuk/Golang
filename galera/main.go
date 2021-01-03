package galera

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	os.Setenv("KUBERNETES_MASTER", "172.17.16.160")
	var kubeconfig *string
	kubeconfig = flag.String("kubeconifg", filepath.Join("C:\\Users\\knk10\\.kube", "spk_config"), "(optional) absolute path to the kubeconfig file!")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	//clientset, err := kubernetes.NewForConfig(config)
	clientset, err := kubernetes.NewForConfig(config)

	appsv1Cient, err := appsv1.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	clients := appsv1Cient.StatefulSets("kafka")

	stss, err := clients.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, sts := range stss.Items {
		fmt.Println("Sts's name is ", sts.Name)
		fmt.Println("Sts's status is", sts.Status.ReadyReplicas)
		fmt.Println("Sts's status is", sts.Status.Replicas)
		fmt.Println("Sts's status is", sts.Spec.ServiceName)

		pods, err := clientset.CoreV1().Pods(sts.Name).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}

		for _, pod := range pods.Items {
			fmt.Println("Pod's name is", pod.Name)

		}

	}
}
