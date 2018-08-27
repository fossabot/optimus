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

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type PipelineList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Items []Pipeline `json:"pipelines"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Pipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Project Project `json:"project"`
}

type Project struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	ArchiveArtifacts bool             `json:"archive"`
	ContainerBuilder ContainerBuilder `json:"container_builder"`
	Notifications    Notifications    `json:"notifications"`
	PollInterval     uint32           `json:"poll_interval"`
	Repository       string           `json:"repository"`
	Username         string           `json:"username"`

	Stages []Stage `json:"stages"`
}

type Stage struct {
	ExitOnError bool   `json:"exit_on_error"`
	Name        string `json:"stage_name"`
	Parallel    bool   `json:"parallel"`

	Commands []Command `json:"commands"`
}

type Command struct {
	RuntimeImage string   `json:"runtime_image"`
	ExitOnError  string   `json:"exit_on_error"`
	Cmd          []string `json:"commands"`

	PodTemplate *v1.ResourceRequirements `json:"pod_template,omitempty"`
}

type ContainerBuilder struct {
	URI      string            `json:"uri"`
	Registry ContainerRegistry `json:"registry"`
}

type ContainerRegistry struct {
	URI string `json:"uri"`
}

type Storage struct {
	URI string `json:"uri"`
}

type Notifications struct {
	URI string `json:"uri"`
}
