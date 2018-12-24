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
	"errors"

	"github.com/sirupsen/logrus"

	//	"k8s.io/apimachinery/pkg/util/wait"
	informersFactory "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"

	optimusAPI "github.com/cloudflavor/optimus/pkg/apis/optimus.cloudflavor.io/v1"
	optimusAPIClient "github.com/cloudflavor/optimus/pkg/client/clientset/versioned"
	optimusIFactory "github.com/cloudflavor/optimus/pkg/client/informers/externalversions"
	optimusListers "github.com/cloudflavor/optimus/pkg/client/listers/optimus.cloudflavor.io/v1"
)

// NewController is a constructor for a controller struct.
func NewController(
	k8sClient *kubernetes.Clientset,
	optimusClient *optimusAPIClient.Clientset,
	informerFactory informersFactory.SharedInformerFactory,
	optimusInformersFactory optimusIFactory.SharedInformerFactory,
) *Controller {

	newController := &Controller{
		K8sClient:      k8sClient,
		OptimusClient:  optimusClient,
		Queue:          workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter()),
		PipelineLister: optimusInformersFactory.Optimus().V1().Pipelines().Lister(),
		PodLister:      informerFactory.Core().V1().Pods().Lister(),
		PVCLister:      informerFactory.Core().V1().PersistentVolumeClaims().Lister(),
		InformerSyncs: append(
			[]cache.InformerSynced{},
			optimusInformersFactory.Optimus().V1().Pipelines().Informer().HasSynced,
			informerFactory.Core().V1().Pods().Informer().HasSynced,
			informerFactory.Core().V1().PersistentVolumeClaims().Informer().HasSynced,
		),
	}

	optimusInformersFactory.Optimus().V1().Pipelines().Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc:    newController.HandleObject,
			DeleteFunc: newController.HandleObject,
			UpdateFunc: func(old, new interface{}) {
				oldPipeline := new.(*optimusAPI.Pipeline)
				newPipeline := old.(*optimusAPI.Pipeline)

				if oldPipeline.ResourceVersion == newPipeline.ResourceVersion {
					logrus.Info("Objects match, skipping update...")
					return
				}
				newController.HandleObject(new)
			},
		},
	)

	return newController
}

// Controller handles all the pipelines resource operations.
type Controller struct {
	K8sClient      *kubernetes.Clientset
	OptimusClient  *optimusAPIClient.Clientset
	Queue          workqueue.RateLimitingInterface
	PipelineLister optimusListers.PipelineLister
	PodLister      listers.PodLister
	PVCLister      listers.PersistentVolumeClaimLister
	InformerSyncs  []cache.InformerSynced
	Recorder       record.EventRecorder
}

// Start will start the controller
func (c *Controller) Start(threadinss int, stopCh <-chan struct{}) error {
	logrus.Info("Started controller")
	defer logrus.Info("Closing controller")

	logrus.Info("Waiting for informer caches to sync")
	if ok := cache.WaitForCacheSync(stopCh, c.InformerSyncs...); !ok {
		return errors.New("Failed to sync informer caches")
	}

	// for i:=0; i < threadinss; i++ {
	// 	go wait.
	// }
	<-stopCh
	return nil
}

// HandleObject will handle adding a new version of an object ot the queue.
func (c *Controller) HandleObject(obj interface{}) {
	logrus.Debugf("Handling object: %#v", obj)

	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		logrus.Errorf("Failed to handle object: %#v with error: %s", obj, err)
		return
	}
	c.Queue.AddRateLimited(key)
}

//func (c *Controller) processWorker
