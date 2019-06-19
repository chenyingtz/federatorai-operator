// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaComponentSpec) DeepCopyInto(out *AlamedaComponentSpec) {
	*out = *in
	if in.Storages != nil {
		in, out := &in.Storages, &out.Storages
		*out = make([]StorageSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.BootStrapContainer = in.BootStrapContainer
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaComponentSpec.
func (in *AlamedaComponentSpec) DeepCopy() *AlamedaComponentSpec {
	if in == nil {
		return nil
	}
	out := new(AlamedaComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaService) DeepCopyInto(out *AlamedaService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaService.
func (in *AlamedaService) DeepCopy() *AlamedaService {
	if in == nil {
		return nil
	}
	out := new(AlamedaService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlamedaService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaServiceList) DeepCopyInto(out *AlamedaServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AlamedaService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaServiceList.
func (in *AlamedaServiceList) DeepCopy() *AlamedaServiceList {
	if in == nil {
		return nil
	}
	out := new(AlamedaServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlamedaServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaServiceSpec) DeepCopyInto(out *AlamedaServiceSpec) {
	*out = *in
	if in.Storages != nil {
		in, out := &in.Storages, &out.Storages
		*out = make([]StorageSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.InfluxdbSectionSet.DeepCopyInto(&out.InfluxdbSectionSet)
	in.GrafanaSectionSet.DeepCopyInto(&out.GrafanaSectionSet)
	in.AlamedaAISectionSet.DeepCopyInto(&out.AlamedaAISectionSet)
	in.AlamedaOperatorSectionSet.DeepCopyInto(&out.AlamedaOperatorSectionSet)
	in.AlamedaDatahubSectionSet.DeepCopyInto(&out.AlamedaDatahubSectionSet)
	in.AlamedaEvictionerSectionSet.DeepCopyInto(&out.AlamedaEvictionerSectionSet)
	in.AdmissionControllerSectionSet.DeepCopyInto(&out.AdmissionControllerSectionSet)
	in.AlamedaRecommenderSectionSet.DeepCopyInto(&out.AlamedaRecommenderSectionSet)
	in.AlamedaExecutorSectionSet.DeepCopyInto(&out.AlamedaExecutorSectionSet)
	in.AlamedaWeaveScopeSectionSet.DeepCopyInto(&out.AlamedaWeaveScopeSectionSet)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaServiceSpec.
func (in *AlamedaServiceSpec) DeepCopy() *AlamedaServiceSpec {
	if in == nil {
		return nil
	}
	out := new(AlamedaServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaServiceStatus) DeepCopyInto(out *AlamedaServiceStatus) {
	*out = *in
	out.CRDVersion = in.CRDVersion
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]AlamedaServiceStatusCondition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaServiceStatus.
func (in *AlamedaServiceStatus) DeepCopy() *AlamedaServiceStatus {
	if in == nil {
		return nil
	}
	out := new(AlamedaServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaServiceStatusCRDVersion) DeepCopyInto(out *AlamedaServiceStatusCRDVersion) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaServiceStatusCRDVersion.
func (in *AlamedaServiceStatusCRDVersion) DeepCopy() *AlamedaServiceStatusCRDVersion {
	if in == nil {
		return nil
	}
	out := new(AlamedaServiceStatusCRDVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaServiceStatusCondition) DeepCopyInto(out *AlamedaServiceStatusCondition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaServiceStatusCondition.
func (in *AlamedaServiceStatusCondition) DeepCopy() *AlamedaServiceStatusCondition {
	if in == nil {
		return nil
	}
	out := new(AlamedaServiceStatusCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Imagestruct) DeepCopyInto(out *Imagestruct) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Imagestruct.
func (in *Imagestruct) DeepCopy() *Imagestruct {
	if in == nil {
		return nil
	}
	out := new(Imagestruct)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageSpec) DeepCopyInto(out *StorageSpec) {
	*out = *in
	if in.Class != nil {
		in, out := &in.Class, &out.Class
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageSpec.
func (in *StorageSpec) DeepCopy() *StorageSpec {
	if in == nil {
		return nil
	}
	out := new(StorageSpec)
	in.DeepCopyInto(out)
	return out
}
