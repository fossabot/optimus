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
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	optimusClient "github.com/cloudflavor/optimus/pkg/client/clientset/versioned"
	optimusInformers "github.com/cloudflavor/optimus/pkg/client/informers/externalversions"
	optimusController "github.com/cloudflavor/optimus/pkg/controller"
	utils "github.com/cloudflavor/optimus/pkg/utils"
)

var (
	k8sConfig = flag.String("kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	master    = flag.String("master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
)

func main() {
	flag.Parse()

	logLevel := os.Getenv("OPTIMUS_LOG_LEVEL")
	if logLevel != "" {
		level, err := strconv.Atoi(logLevel)
		if err != nil {
			logrus.Fatalf("Failed to convert assigned log level! %#v", err)
		}
		logrus.SetLevel(logrus.AllLevels[level])
	}
	logrus.SetOutput(os.Stdout)

	cfg, err := clientcmd.BuildConfigFromFlags(*master, *k8sConfig)
	if err != nil {
		logrus.Fatalf("Error building configuration: %v", err)
	}

	localClient, err := optimusClient.NewForConfig(cfg)
	if err != nil {
		logrus.Fatalf("Error building clientset: %v", err)
	}

	client, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		logrus.Fatalf("Error creating new kubernetes client, %#v", err)
	}

	informersFactory := informers.NewSharedInformerFactory(client, time.Second*10)
	optimusInformersFactory := optimusInformers.NewSharedInformerFactory(localClient, time.Second*10)

	controller := optimusController.NewController(
		client,
		localClient,
		informersFactory,
		optimusInformersFactory,
	)
	stopCh := utils.SetupSignalHandler()
	logrus.Info("Starting informers factory")
	informersFactory.Start(stopCh)

	logrus.Info("Starting optimus informers factory")
	optimusInformersFactory.Start(stopCh)

	if err := controller.Start(2, stopCh); err != nil {
		logrus.Fatalf("Failed to start optimus controller: %s", err.Error())
	}
}
