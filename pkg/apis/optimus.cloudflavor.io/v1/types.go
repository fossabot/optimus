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

package v1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PipelineList is a list of pipeline CRDs.

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type PipelineList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Items []Pipeline `json:"items"`
}

// Pipeline is the utmost CRD type that defines a new pipeline.
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Pipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Project Project `json:"project"`
}

// Project holds the specifics of a pipeline project.
type Project struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	ArchiveArtifacts bool              `json:"archive"`
	Registry         ContainerRegistry `json:"registry"`
	Notifications    Notifications     `json:"notifications,omitempty"`
	PollInterval     uint32            `json:"poll_interval"`
	Repository       string            `json:"repository"`
	Username         string            `json:"username"`
	RunInterval      *RunInterval      `json:"run_interval,omitempty"`

	Stages []Stage `json:"stages"`
}

// Stage represents a stage in the pipeline.
type Stage struct {
	Name        string `json:"stage_name"`
	ExitOnError bool   `json:"exit_on_error"`
	Parallel    bool   `json:"parallel"`

	Commands []Command `json:"commands"`
}

// Command represents a chain of commands that is related to a stage in the
// pipeline.
type Command struct {
	RuntimeImage string   `json:"runtime_image"`
	ExitOnError  string   `json:"exit_on_error"`
	Cmd          []string `json:"commands"`

	PodTemplate *v1.ResourceRequirements `json:"pod_template,omitempty"`
}

// ContainerRegistry holds information about a registry where an image will be
// pushed once an image has been built.
// NOTE: add username and password/token???
type ContainerRegistry struct {
	URI string `json:"uri"`
}

// Storage is the interface that is used for abstracting away the object storage
// client that will be used to store and archive artifacts. For now this will
// be represented by minio.
type Storage struct {
	URI string `json:"uri"`
}

// Notifications represents a webhook notification that is triggered by an
// action in the pipeline.
type Notifications struct {
	URI   string `json:"uri"`
	Token string `json:"token,omitempty"`
}

// RunInterval holds the time date interval for an automatic pipeline to run
// in cron format.
// TODO: add kubernetes cronjob types to this, avoid reinventing the wheel.
type RunInterval struct{}
