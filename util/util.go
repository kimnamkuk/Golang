package util

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strconv"

	"k8s.io/client-go/kubernetes"
	k8s "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
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
func Cat(strFile string) []byte {
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

	return output
}

func K8sInCluster() *k8s.Clientset {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}

func K8sOutOfCluster() *k8s.Clientset {
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

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
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

	set_chown(strDirpath, uid, gid)
}
