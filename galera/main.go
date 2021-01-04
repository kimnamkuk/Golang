package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	util "github.com/kimnamkuk/Golang/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	strPath := util.pwd()
	fmt.Println(strPath)

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

func init_pod() {
	// check files (token,SA .etc)
	if _, err := os.Stat("/var/run/secrets/kubernetes.io/serviceaccount/token"); os.IsNotExist(err) {
		fmt.Println("No Found /var/run/secrets/kubernetes.io/serviceaccount/token")
		fmt.Println(err.Error())
		panic(err.Error())
	}

	if _, err := os.Stat("/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"); os.IsNotExist(err) {
		fmt.Println("No Found /var/run/secrets/kubernetes.io/serviceaccount/ca.crt")
		fmt.Println(err.Error())
		panic(err.Error())
	}

	if strEnv := os.Getenv("SERVER_ID"); strEnv == "" {
		fmt.Println("No Set env of server_id")
		panic(errors.New("No set envirments"))
	}

	if strEnv := os.Getenv("SSH_PWD"); strEnv == "" {
		fmt.Println("No Set env of ssh_pwd")
		panic(errors.New("No set envirments"))
	}

}

func check_mode() {
	clientset := util.K8sInCluster()
}
