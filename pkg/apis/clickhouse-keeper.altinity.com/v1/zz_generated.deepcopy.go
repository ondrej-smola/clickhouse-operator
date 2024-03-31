//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	clickhousealtinitycomv1 "github.com/altinity/clickhouse-operator/pkg/apis/clickhouse.altinity.com/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChkCluster) DeepCopyInto(out *ChkCluster) {
	*out = *in
	if in.Layout != nil {
		in, out := &in.Layout, &out.Layout
		*out = new(ChkClusterLayout)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChkCluster.
func (in *ChkCluster) DeepCopy() *ChkCluster {
	if in == nil {
		return nil
	}
	out := new(ChkCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChkClusterLayout) DeepCopyInto(out *ChkClusterLayout) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChkClusterLayout.
func (in *ChkClusterLayout) DeepCopy() *ChkClusterLayout {
	if in == nil {
		return nil
	}
	out := new(ChkClusterLayout)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChkConfiguration) DeepCopyInto(out *ChkConfiguration) {
	*out = *in
	if in.Settings != nil {
		in, out := &in.Settings, &out.Settings
		*out = new(clickhousealtinitycomv1.Settings)
		(*in).DeepCopyInto(*out)
	}
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]*ChkCluster, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ChkCluster)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChkConfiguration.
func (in *ChkConfiguration) DeepCopy() *ChkConfiguration {
	if in == nil {
		return nil
	}
	out := new(ChkConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChkSpec) DeepCopyInto(out *ChkSpec) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(ChkConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = new(clickhousealtinitycomv1.ChiTemplates)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChkSpec.
func (in *ChkSpec) DeepCopy() *ChkSpec {
	if in == nil {
		return nil
	}
	out := new(ChkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChkStatus) DeepCopyInto(out *ChkStatus) {
	*out = *in
	if in.ReadyReplicas != nil {
		in, out := &in.ReadyReplicas, &out.ReadyReplicas
		*out = make([]clickhousealtinitycomv1.ChiZookeeperNode, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PodIPs != nil {
		in, out := &in.PodIPs, &out.PodIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FQDNs != nil {
		in, out := &in.FQDNs, &out.FQDNs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NormalizedCHK != nil {
		in, out := &in.NormalizedCHK, &out.NormalizedCHK
		*out = new(ClickHouseKeeperInstallation)
		(*in).DeepCopyInto(*out)
	}
	if in.NormalizedCHKCompleted != nil {
		in, out := &in.NormalizedCHKCompleted, &out.NormalizedCHKCompleted
		*out = new(ClickHouseKeeperInstallation)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChkStatus.
func (in *ChkStatus) DeepCopy() *ChkStatus {
	if in == nil {
		return nil
	}
	out := new(ChkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClickHouseKeeperInstallation) DeepCopyInto(out *ClickHouseKeeperInstallation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(ChkStatus)
		(*in).DeepCopyInto(*out)
	}
	out.Runtime = in.Runtime
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClickHouseKeeperInstallation.
func (in *ClickHouseKeeperInstallation) DeepCopy() *ClickHouseKeeperInstallation {
	if in == nil {
		return nil
	}
	out := new(ClickHouseKeeperInstallation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClickHouseKeeperInstallation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClickHouseKeeperInstallationList) DeepCopyInto(out *ClickHouseKeeperInstallationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClickHouseKeeperInstallation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClickHouseKeeperInstallationList.
func (in *ClickHouseKeeperInstallationList) DeepCopy() *ClickHouseKeeperInstallationList {
	if in == nil {
		return nil
	}
	out := new(ClickHouseKeeperInstallationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClickHouseKeeperInstallationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClickHouseKeeperInstallationRuntime) DeepCopyInto(out *ClickHouseKeeperInstallationRuntime) {
	*out = *in
	out.statusCreatorMutex = in.statusCreatorMutex
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClickHouseKeeperInstallationRuntime.
func (in *ClickHouseKeeperInstallationRuntime) DeepCopy() *ClickHouseKeeperInstallationRuntime {
	if in == nil {
		return nil
	}
	out := new(ClickHouseKeeperInstallationRuntime)
	in.DeepCopyInto(out)
	return out
}
