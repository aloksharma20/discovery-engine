//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveredPolicy) DeepCopyInto(out *DiscoveredPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveredPolicy.
func (in *DiscoveredPolicy) DeepCopy() *DiscoveredPolicy {
	if in == nil {
		return nil
	}
	out := new(DiscoveredPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DiscoveredPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveredPolicyList) DeepCopyInto(out *DiscoveredPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DiscoveredPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveredPolicyList.
func (in *DiscoveredPolicyList) DeepCopy() *DiscoveredPolicyList {
	if in == nil {
		return nil
	}
	out := new(DiscoveredPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DiscoveredPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveredPolicySpec) DeepCopyInto(out *DiscoveredPolicySpec) {
	*out = *in
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(apiextensionsv1.JSON)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveredPolicySpec.
func (in *DiscoveredPolicySpec) DeepCopy() *DiscoveredPolicySpec {
	if in == nil {
		return nil
	}
	out := new(DiscoveredPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveredPolicyStatus) DeepCopyInto(out *DiscoveredPolicyStatus) {
	*out = *in
	in.LastUpdatedTime.DeepCopyInto(&out.LastUpdatedTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveredPolicyStatus.
func (in *DiscoveredPolicyStatus) DeepCopy() *DiscoveredPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(DiscoveredPolicyStatus)
	in.DeepCopyInto(out)
	return out
}
