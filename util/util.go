package util

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strconv"

	k8s "k8s.io/client-go/kubernetes"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	rest "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// Get Workdir of current
func Pwd() string {

	dir, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	return dir
}
func Cat(strFile string) string {
	var output []byte
	var err error
	file, err := os.Open(strFile)
	defer file.Close()

	if err != nil {
		panic(err.Error())
	}

	output, err = ioutil.ReadAll(file)

	if err != nil {
		panic(err.Error())
	}

	return string(output)
}

func Set_chown(strDirpath string, nUid, nGid int) {

	_, err := os.Stat(strDirpath)
	if os.IsNotExist(err) {
		fmt.Println("No found Directory of", strDirpath)
		panic(err.Error())
	}

	err = os.Chown(strDirpath, nUid, nGid)
	if err != nil {
		fmt.Println("Can't Change to  other permission mode")
		panic(err.Error())
	}
}

func Set_chown2(strDirpath, strUid, strGid string) {

	curuser, err := user.Lookup(strUid)
	if err != nil {
		fmt.Println("Can't Get uid from user's name")
		panic(err.Error())
	}

	curgroup, err := user.LookupGroup(strGid)
	if err != nil {
		fmt.Println("Can't Get gid from group's name")
		panic(err.Error())
	}

	uid, _ := strconv.Atoi(curuser.Uid)
	gid, _ := strconv.Atoi(curgroup.Gid)

	Set_chown(strDirpath, uid, gid)
}

func Get_Appsv1_In_Cluster() *appsv1.AppsV1Client {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	appsv1Client, err := appsv1.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return appsv1Client
}

func Get_Appsv1_Outof_Cluster() *appsv1.AppsV1Client {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	appsv1Client, err := appsv1.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return appsv1Client
}

func Get_Clientset_In_Cluster() *k8s.Clientset {

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err := k8s.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}

func Get_Clientset_OutOf_Cluster() *k8s.Clientset {

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := k8s.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}
