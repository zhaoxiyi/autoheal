// +build !ignore_autogenerated

/*
Copyright (c) 2018 Red Hat, Inc.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha2

import (
	unsafe "unsafe"

	autoheal "github.com/openshift/autoheal/pkg/apis/autoheal"
	v1 "k8s.io/api/batch/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha2_AWXJobAction_To_autoheal_AWXJobAction,
		Convert_autoheal_AWXJobAction_To_v1alpha2_AWXJobAction,
		Convert_v1alpha2_HealingRule_To_autoheal_HealingRule,
		Convert_autoheal_HealingRule_To_v1alpha2_HealingRule,
		Convert_v1alpha2_HealingRuleList_To_autoheal_HealingRuleList,
		Convert_autoheal_HealingRuleList_To_v1alpha2_HealingRuleList,
	)
}

func autoConvert_v1alpha2_AWXJobAction_To_autoheal_AWXJobAction(in *AWXJobAction, out *autoheal.AWXJobAction, s conversion.Scope) error {
	out.Template = in.Template
	out.ExtraVars = *(*autoheal.JsonDoc)(unsafe.Pointer(&in.ExtraVars))
	return nil
}

// Convert_v1alpha2_AWXJobAction_To_autoheal_AWXJobAction is an autogenerated conversion function.
func Convert_v1alpha2_AWXJobAction_To_autoheal_AWXJobAction(in *AWXJobAction, out *autoheal.AWXJobAction, s conversion.Scope) error {
	return autoConvert_v1alpha2_AWXJobAction_To_autoheal_AWXJobAction(in, out, s)
}

func autoConvert_autoheal_AWXJobAction_To_v1alpha2_AWXJobAction(in *autoheal.AWXJobAction, out *AWXJobAction, s conversion.Scope) error {
	out.Template = in.Template
	out.ExtraVars = *(*JsonDoc)(unsafe.Pointer(&in.ExtraVars))
	return nil
}

// Convert_autoheal_AWXJobAction_To_v1alpha2_AWXJobAction is an autogenerated conversion function.
func Convert_autoheal_AWXJobAction_To_v1alpha2_AWXJobAction(in *autoheal.AWXJobAction, out *AWXJobAction, s conversion.Scope) error {
	return autoConvert_autoheal_AWXJobAction_To_v1alpha2_AWXJobAction(in, out, s)
}

func autoConvert_v1alpha2_HealingRule_To_autoheal_HealingRule(in *HealingRule, out *autoheal.HealingRule, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Labels = *(*map[string]string)(unsafe.Pointer(&in.Labels))
	out.Annotations = *(*map[string]string)(unsafe.Pointer(&in.Annotations))
	out.AWXJob = (*autoheal.AWXJobAction)(unsafe.Pointer(in.AWXJob))
	out.BatchJob = (*v1.Job)(unsafe.Pointer(in.BatchJob))
	return nil
}

// Convert_v1alpha2_HealingRule_To_autoheal_HealingRule is an autogenerated conversion function.
func Convert_v1alpha2_HealingRule_To_autoheal_HealingRule(in *HealingRule, out *autoheal.HealingRule, s conversion.Scope) error {
	return autoConvert_v1alpha2_HealingRule_To_autoheal_HealingRule(in, out, s)
}

func autoConvert_autoheal_HealingRule_To_v1alpha2_HealingRule(in *autoheal.HealingRule, out *HealingRule, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Labels = *(*map[string]string)(unsafe.Pointer(&in.Labels))
	out.Annotations = *(*map[string]string)(unsafe.Pointer(&in.Annotations))
	out.AWXJob = (*AWXJobAction)(unsafe.Pointer(in.AWXJob))
	out.BatchJob = (*v1.Job)(unsafe.Pointer(in.BatchJob))
	return nil
}

// Convert_autoheal_HealingRule_To_v1alpha2_HealingRule is an autogenerated conversion function.
func Convert_autoheal_HealingRule_To_v1alpha2_HealingRule(in *autoheal.HealingRule, out *HealingRule, s conversion.Scope) error {
	return autoConvert_autoheal_HealingRule_To_v1alpha2_HealingRule(in, out, s)
}

func autoConvert_v1alpha2_HealingRuleList_To_autoheal_HealingRuleList(in *HealingRuleList, out *autoheal.HealingRuleList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]autoheal.HealingRule)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha2_HealingRuleList_To_autoheal_HealingRuleList is an autogenerated conversion function.
func Convert_v1alpha2_HealingRuleList_To_autoheal_HealingRuleList(in *HealingRuleList, out *autoheal.HealingRuleList, s conversion.Scope) error {
	return autoConvert_v1alpha2_HealingRuleList_To_autoheal_HealingRuleList(in, out, s)
}

func autoConvert_autoheal_HealingRuleList_To_v1alpha2_HealingRuleList(in *autoheal.HealingRuleList, out *HealingRuleList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]HealingRule)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_autoheal_HealingRuleList_To_v1alpha2_HealingRuleList is an autogenerated conversion function.
func Convert_autoheal_HealingRuleList_To_v1alpha2_HealingRuleList(in *autoheal.HealingRuleList, out *HealingRuleList, s conversion.Scope) error {
	return autoConvert_autoheal_HealingRuleList_To_v1alpha2_HealingRuleList(in, out, s)
}
