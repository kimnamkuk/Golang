package util

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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

func FindFile(strTargetDir string, lstPattern []string) {

	for _, v := range lstPattern {
		matches, err := filepath.Glob(strTargetDir + v)

		if err != nil {
			log.Println(err)
		}

		if len(matches) != 0 {
			log.Println("Found : ", matches)
		}
	}
}

func IsFindFile(strTargetDir string, strFileName string) bool {
	matches, err := filepath.Glob(strTargetDir + strFileName)
	if err != nil {
		log.Println(err)
		return false
	}

	if len(matches) != 0 {
		log.Println("Found : ", matches)
		return true
	}

	return false
}

//You must set null of strAbspath, if you want to use $home/.kube/config
func GetConfigOutofCluster(strAbspath string, strMasterEnv string) *rest.Config {
	var kubeconfig *string
	if strMasterEnv != "" {
		os.Setenv("KUBERNETES_MASTER", strMasterEnv)
	}

	if strAbspath == "" {
		home := homedir.HomeDir()
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", strAbspath, "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	return config
}

func GetConfigInCluster() *rest.Config {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	return config
}

func GetAppsv1(config *rest.Config) *appsv1.AppsV1Client {

	appsv1Client, err := appsv1.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return appsv1Client
}

func GetClientset(config *rest.Config) *k8s.Clientset {

	clientset, err := k8s.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}
