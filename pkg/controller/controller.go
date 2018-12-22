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

package controller

import (
	"github.com/golang/glog"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	//	optimusV1 "github.com/cloudflavor/optimus/pkg/apis/optimus.cloudflavor.io/v1"
	localClient "github.com/cloudflavor/optimus/pkg/client/clientset/versioned"
)

func NewController(
	k8sClient *kubernetes.Clientset,
	optimusClient *localClient.Clientset,
) *Controller {
	return &Controller{
		OptimusClient: optimusClient,
		K8sClient:     k8sClient,
	}
}

type Controller struct {
	K8sClient     *kubernetes.Clientset
	OptimusClient *localClient.Clientset
}

// Run will start the pipeline.
func (c *Controller) Run(stopCh chan struct{}) error {
	pipes, err := c.OptimusClient.OptimusV1().Pipelines("optimus").List(metav1.ListOptions{})

	if err != nil {
		glog.Fatalf("Failed to list pipelines: %s", err)
	}

	for _, pipe := range pipes.Items {
		//		glog.V(0).Infof("Found pipeline: %#v", pipe.Jobs)
		for _, job := range pipe.Jobs {
			glog.V(0).Infof("job: %#v", job)
			for _, stage := range job.Stages {
				glog.V(0).Infof("stage: %#v", stage)
			}
		}
	}
	return nil
}
