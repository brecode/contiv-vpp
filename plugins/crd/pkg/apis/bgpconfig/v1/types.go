// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	"github.com/contiv/vpp/plugins/crd/pkg/apis/bgpconfig"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CRD Constants
const (
	CRDGroup                   string = bgpconfig.GroupName
	CRDGroupVersion            string = "v1"
	CRDContivBgpConfigPlural   string = "bgpconfigs"
	CRDFullContivBgpConfigName string = CRDContivBgpConfigPlural + "." + CRDGroup
)

// BgpConfig describes contiv bgp configuration custom resource
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type BgpConfig struct {
	// TypeMeta is the metadata for the resource, like kind and apiversion
	metav1.TypeMeta `json:",inline"`
	// ObjectMeta contains the metadata for the particular object
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Spec is the custom resource spec
	Spec   BgpConfigSpec   `json:"spec,omitempty"`
	Status BgpConfigStatus `json:"status,omitempty"`
}

// BgpConfigSpec is the spec for the contiv bgp configuration resource.
type BgpConfigSpec struct {
	// FILL STRUCT HERE
}

// BgpConfigList is a list of bgp configuration resource
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type BgpConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []BgpConfig `json:"items"`
}

// BgpConfigStatus is the state for the contiv ode configuration
type BgpConfigStatus struct {
	State   string `json:"state,omitempty"`
	Message string `json:"message,omitempty"`
}
