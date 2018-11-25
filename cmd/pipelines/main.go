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
	"os"

	"github.com/golang/glog"

	versionedClient "github.com/pi-victor/pipelines/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sCmd "k8s.io/client-go/tools/clientcmd"
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

	cfg, err := k8sCmd.BuildConfigFromFlags(*master, *kuberconfig)

	if err != nil {
		glog.Fatalf("Error building kubeconfig: %v", err)
	}

	client, err := versionedClient.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building example clientset: %v", err)
	}
	glog.V(0).Infof("%#v", cfg)

	list, err := client.CloudflavorV1().Pipelines("ci-namespace").List(metav1.ListOptions{})

	if err != nil {
		glog.Fatalf("Error listing all pipelines: %v", err)
	}
	glog.V(0).Infof("These are the pipelines: %#v", list)
	for {
		if list.Items != nil {
			for _, pipeline := range list.Items {
				glog.V(0).Infof("pipeline: %#v\n", pipeline)
				continue
			}
			glog.V(0).Info("No pipelines found, list is empty!!!")
		}
	}

}
