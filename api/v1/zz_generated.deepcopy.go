//go:build !ignore_autogenerated

/*
Copyright 2024.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSidecarConfiguration) DeepCopyInto(out *InstanceSidecarConfiguration) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSidecarConfiguration.
func (in *InstanceSidecarConfiguration) DeepCopy() *InstanceSidecarConfiguration {
	if in == nil {
		return nil
	}
	out := new(InstanceSidecarConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStore) DeepCopyInto(out *ObjectStore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStore.
func (in *ObjectStore) DeepCopy() *ObjectStore {
	if in == nil {
		return nil
	}
	out := new(ObjectStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObjectStore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStoreList) DeepCopyInto(out *ObjectStoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ObjectStore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStoreList.
func (in *ObjectStoreList) DeepCopy() *ObjectStoreList {
	if in == nil {
		return nil
	}
	out := new(ObjectStoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObjectStoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStoreSpec) DeepCopyInto(out *ObjectStoreSpec) {
	*out = *in
	in.Configuration.DeepCopyInto(&out.Configuration)
	in.InstanceSidecarConfiguration.DeepCopyInto(&out.InstanceSidecarConfiguration)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStoreSpec.
func (in *ObjectStoreSpec) DeepCopy() *ObjectStoreSpec {
	if in == nil {
		return nil
	}
	out := new(ObjectStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStoreStatus) DeepCopyInto(out *ObjectStoreStatus) {
	*out = *in
	if in.ServerRecoveryWindow != nil {
		in, out := &in.ServerRecoveryWindow, &out.ServerRecoveryWindow
		*out = make(map[string]RecoveryWindow, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStoreStatus.
func (in *ObjectStoreStatus) DeepCopy() *ObjectStoreStatus {
	if in == nil {
		return nil
	}
	out := new(ObjectStoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecoveryWindow) DeepCopyInto(out *RecoveryWindow) {
	*out = *in
	if in.FirstRecoverabilityPoint != nil {
		in, out := &in.FirstRecoverabilityPoint, &out.FirstRecoverabilityPoint
		*out = (*in).DeepCopy()
	}
	if in.LastSuccessfulBackupTime != nil {
		in, out := &in.LastSuccessfulBackupTime, &out.LastSuccessfulBackupTime
		*out = (*in).DeepCopy()
	}
	if in.FirstWALSubmissionTime != nil {
		in, out := &in.FirstWALSubmissionTime, &out.FirstWALSubmissionTime
		*out = (*in).DeepCopy()
	}
	if in.LastWALSubmissionTime != nil {
		in, out := &in.LastWALSubmissionTime, &out.LastWALSubmissionTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecoveryWindow.
func (in *RecoveryWindow) DeepCopy() *RecoveryWindow {
	if in == nil {
		return nil
	}
	out := new(RecoveryWindow)
	in.DeepCopyInto(out)
	return out
}
