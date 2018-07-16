/*
Copyright 2018 Victor Palade.
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
	"fmt"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sClientCmd "k8s.io/client-go/tools/clientcmd"
	k8sClient "k8s.io/go-client"

	versionedClient "github.com/PI-Victor/pipelines/pkg/client/clientset/versioned"
)

var (
	kuberconfig = flag.String("kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	master      = flag.String("master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
)

func main() {
	// set logging capabilities. Flush logs and set stderr as default..
	defer glog.Flush()

	flag.Set("logtostderr", "true")
	if os.Getenv("PIPELINE_LOG_LEVEL") != "" {
		flag.Set("v", os.Getenv("PIPELINE_LOG_LEVEL"))
	}
	flag.Parse()

	cfg, err := client.BuildConfigFromFlags(*master, *kuberconfig)
	if err != nil {
		glog.Fatalf("Error building kubeconfig: %v", err)
	}

	client, err := versionedClient.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building example clientset: %v", err)
	}

	// list, err := exampleClient.ExampleV1().Databases("default").List(metav1.ListOptions{})
	// if err != nil {
	// 	glog.Fatalf("Error listing all databases: %v", err)
	// }
	//
	// for _, db := range list.Items {
	// 	fmt.Printf("database %s with user %q\n", db.Name, db.Spec.User)
	// }
}
