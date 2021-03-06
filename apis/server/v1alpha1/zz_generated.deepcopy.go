//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright AppsCode Inc. and Contributors

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

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Server) DeepCopyInto(out *Server) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Server.
func (in *Server) DeepCopy() *Server {
	if in == nil {
		return nil
	}
	out := new(Server)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Server) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerList) DeepCopyInto(out *ServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Server, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerList.
func (in *ServerList) DeepCopy() *ServerList {
	if in == nil {
		return nil
	}
	out := new(ServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerSpec) DeepCopyInto(out *ServerSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ServerSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerSpec.
func (in *ServerSpec) DeepCopy() *ServerSpec {
	if in == nil {
		return nil
	}
	out := new(ServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerSpecLogin) DeepCopyInto(out *ServerSpecLogin) {
	*out = *in
	if in.CreatePassword != nil {
		in, out := &in.CreatePassword, &out.CreatePassword
		*out = new(bool)
		**out = **in
	}
	if in.Keys != nil {
		in, out := &in.Keys, &out.Keys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PasswordDelivery != nil {
		in, out := &in.PasswordDelivery, &out.PasswordDelivery
		*out = new(string)
		**out = **in
	}
	if in.User != nil {
		in, out := &in.User, &out.User
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerSpecLogin.
func (in *ServerSpecLogin) DeepCopy() *ServerSpecLogin {
	if in == nil {
		return nil
	}
	out := new(ServerSpecLogin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerSpecNetworkInterface) DeepCopyInto(out *ServerSpecNetworkInterface) {
	*out = *in
	if in.Bootable != nil {
		in, out := &in.Bootable, &out.Bootable
		*out = new(bool)
		**out = **in
	}
	if in.IpAddress != nil {
		in, out := &in.IpAddress, &out.IpAddress
		*out = new(string)
		**out = **in
	}
	if in.IpAddressFamily != nil {
		in, out := &in.IpAddressFamily, &out.IpAddressFamily
		*out = new(string)
		**out = **in
	}
	if in.IpAddressFloating != nil {
		in, out := &in.IpAddressFloating, &out.IpAddressFloating
		*out = new(bool)
		**out = **in
	}
	if in.MacAddress != nil {
		in, out := &in.MacAddress, &out.MacAddress
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(string)
		**out = **in
	}
	if in.SourceIPFiltering != nil {
		in, out := &in.SourceIPFiltering, &out.SourceIPFiltering
		*out = new(bool)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerSpecNetworkInterface.
func (in *ServerSpecNetworkInterface) DeepCopy() *ServerSpecNetworkInterface {
	if in == nil {
		return nil
	}
	out := new(ServerSpecNetworkInterface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerSpecResource) DeepCopyInto(out *ServerSpecResource) {
	*out = *in
	if in.Cpu != nil {
		in, out := &in.Cpu, &out.Cpu
		*out = new(int64)
		**out = **in
	}
	if in.Firewall != nil {
		in, out := &in.Firewall, &out.Firewall
		*out = new(bool)
		**out = **in
	}
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(int64)
		**out = **in
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	if in.Login != nil {
		in, out := &in.Login, &out.Login
		*out = new(ServerSpecLogin)
		(*in).DeepCopyInto(*out)
	}
	if in.Mem != nil {
		in, out := &in.Mem, &out.Mem
		*out = new(int64)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(bool)
		**out = **in
	}
	if in.NetworkInterface != nil {
		in, out := &in.NetworkInterface, &out.NetworkInterface
		*out = make([]ServerSpecNetworkInterface, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Plan != nil {
		in, out := &in.Plan, &out.Plan
		*out = new(string)
		**out = **in
	}
	if in.SimpleBackup != nil {
		in, out := &in.SimpleBackup, &out.SimpleBackup
		*out = new(ServerSpecSimpleBackup)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageDevices != nil {
		in, out := &in.StorageDevices, &out.StorageDevices
		*out = make([]ServerSpecStorageDevices, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(ServerSpecTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
	if in.UserData != nil {
		in, out := &in.UserData, &out.UserData
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerSpecResource.
func (in *ServerSpecResource) DeepCopy() *ServerSpecResource {
	if in == nil {
		return nil
	}
	out := new(ServerSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerSpecSimpleBackup) DeepCopyInto(out *ServerSpecSimpleBackup) {
	*out = *in
	if in.Plan != nil {
		in, out := &in.Plan, &out.Plan
		*out = new(string)
		**out = **in
	}
	if in.Time != nil {
		in, out := &in.Time, &out.Time
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerSpecSimpleBackup.
func (in *ServerSpecSimpleBackup) DeepCopy() *ServerSpecSimpleBackup {
	if in == nil {
		return nil
	}
	out := new(ServerSpecSimpleBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerSpecStorageDevices) DeepCopyInto(out *ServerSpecStorageDevices) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerSpecStorageDevices.
func (in *ServerSpecStorageDevices) DeepCopy() *ServerSpecStorageDevices {
	if in == nil {
		return nil
	}
	out := new(ServerSpecStorageDevices)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerSpecTemplate) DeepCopyInto(out *ServerSpecTemplate) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.BackupRule != nil {
		in, out := &in.BackupRule, &out.BackupRule
		*out = new(ServerSpecTemplateBackupRule)
		(*in).DeepCopyInto(*out)
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(int64)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(string)
		**out = **in
	}
	if in.Tier != nil {
		in, out := &in.Tier, &out.Tier
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerSpecTemplate.
func (in *ServerSpecTemplate) DeepCopy() *ServerSpecTemplate {
	if in == nil {
		return nil
	}
	out := new(ServerSpecTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerSpecTemplateBackupRule) DeepCopyInto(out *ServerSpecTemplateBackupRule) {
	*out = *in
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(string)
		**out = **in
	}
	if in.Retention != nil {
		in, out := &in.Retention, &out.Retention
		*out = new(int64)
		**out = **in
	}
	if in.Time != nil {
		in, out := &in.Time, &out.Time
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerSpecTemplateBackupRule.
func (in *ServerSpecTemplateBackupRule) DeepCopy() *ServerSpecTemplateBackupRule {
	if in == nil {
		return nil
	}
	out := new(ServerSpecTemplateBackupRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerStatus) DeepCopyInto(out *ServerStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerStatus.
func (in *ServerStatus) DeepCopy() *ServerStatus {
	if in == nil {
		return nil
	}
	out := new(ServerStatus)
	in.DeepCopyInto(out)
	return out
}
