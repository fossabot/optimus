/*
Copyright 2018 Cloudflavor Org contributors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	//	"time"
	"os"

	"github.com/golang/glog"

	//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	optimusClient "github.com/cloudflavor/optimus/pkg/client/clientset/versioned"
	//	optimusInformers "github.com/cloudflavor/optimus/pkg/client/informers/externalversions"
	optimusController "github.com/cloudflavor/optimus/pkg/controller"
)

var (
	k8sConfig = flag.String("kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	master    = flag.String("master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
)

func init() {
	// set logging capabilities. Flush logs and set stderr as default..
	defer glog.Flush()
	flag.Set("logtostderr", "true")
	if os.Getenv("OPTIMUS_LOG_LEVEL") != "" {
		flag.Set("v", os.Getenv("OPTIMUS_LOG_LEVEL"))
	}
	flag.Parse()
}

func main() {
	flag.Parse()
	cfg, err := clientcmd.BuildConfigFromFlags(*master, *k8sConfig)
	if err != nil {
		glog.Fatalf("Error building kubeconfig: %v", err)
	}
	localClient, err := optimusClient.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building example clientset: %v", err)
	}

	k8sClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error creating new kubernetes client, %#v", err)
	}

	//	optimusInformersFactory := optimusInformers.NewSharedInformerFactory(localClient, time.Second*10)
	//	k8sInformerFactory := informers.NewSharedInformerFactory(k8sClient, time.Second*10)

	controller := optimusController.NewController(k8sClient, localClient)

	stop := make(chan struct{})
	defer close(stop)
	go controller.Run(stop)

	select {}
}
