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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	v1alpha1 "k8s.io/api/admissionregistration/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	admissionregistration "k8s.io/kubernetes/pkg/apis/admissionregistration"
	admissionregistrationv1 "k8s.io/kubernetes/pkg/apis/admissionregistration/v1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v1alpha1.AuditAnnotation)(nil), (*admissionregistration.AuditAnnotation)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_AuditAnnotation_To_admissionregistration_AuditAnnotation(a.(*v1alpha1.AuditAnnotation), b.(*admissionregistration.AuditAnnotation), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*admissionregistration.AuditAnnotation)(nil), (*v1alpha1.AuditAnnotation)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_admissionregistration_AuditAnnotation_To_v1alpha1_AuditAnnotation(a.(*admissionregistration.AuditAnnotation), b.(*v1alpha1.AuditAnnotation), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.MatchResources)(nil), (*admissionregistration.MatchResources)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_MatchResources_To_admissionregistration_MatchResources(a.(*v1alpha1.MatchResources), b.(*admissionregistration.MatchResources), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*admissionregistration.MatchResources)(nil), (*v1alpha1.MatchResources)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_admissionregistration_MatchResources_To_v1alpha1_MatchResources(a.(*admissionregistration.MatchResources), b.(*v1alpha1.MatchResources), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.NamedRuleWithOperations)(nil), (*admissionregistration.NamedRuleWithOperations)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_NamedRuleWithOperations_To_admissionregistration_NamedRuleWithOperations(a.(*v1alpha1.NamedRuleWithOperations), b.(*admissionregistration.NamedRuleWithOperations), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*admissionregistration.NamedRuleWithOperations)(nil), (*v1alpha1.NamedRuleWithOperations)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_admissionregistration_NamedRuleWithOperations_To_v1alpha1_NamedRuleWithOperations(a.(*admissionregistration.NamedRuleWithOperations), b.(*v1alpha1.NamedRuleWithOperations), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.ParamKind)(nil), (*admissionregistration.ParamKind)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ParamKind_To_admissionregistration_ParamKind(a.(*v1alpha1.ParamKind), b.(*admissionregistration.ParamKind), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*admissionregistration.ParamKind)(nil), (*v1alpha1.ParamKind)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_admissionregistration_ParamKind_To_v1alpha1_ParamKind(a.(*admissionregistration.ParamKind), b.(*v1alpha1.ParamKind), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.ParamRef)(nil), (*admissionregistration.ParamRef)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ParamRef_To_admissionregistration_ParamRef(a.(*v1alpha1.ParamRef), b.(*admissionregistration.ParamRef), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*admissionregistration.ParamRef)(nil), (*v1alpha1.ParamRef)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_admissionregistration_ParamRef_To_v1alpha1_ParamRef(a.(*admissionregistration.ParamRef), b.(*v1alpha1.ParamRef), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.ValidatingAdmissionPolicy)(nil), (*admissionregistration.ValidatingAdmissionPolicy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ValidatingAdmissionPolicy_To_admissionregistration_ValidatingAdmissionPolicy(a.(*v1alpha1.ValidatingAdmissionPolicy), b.(*admissionregistration.ValidatingAdmissionPolicy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*admissionregistration.ValidatingAdmissionPolicy)(nil), (*v1alpha1.ValidatingAdmissionPolicy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_admissionregistration_ValidatingAdmissionPolicy_To_v1alpha1_ValidatingAdmissionPolicy(a.(*admissionregistration.ValidatingAdmissionPolicy), b.(*v1alpha1.ValidatingAdmissionPolicy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.ValidatingAdmissionPolicyBinding)(nil), (*admissionregistration.ValidatingAdmissionPolicyBinding)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ValidatingAdmissionPolicyBinding_To_admissionregistration_ValidatingAdmissionPolicyBinding(a.(*v1alpha1.ValidatingAdmissionPolicyBinding), b.(*admissionregistration.ValidatingAdmissionPolicyBinding), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*admissionregistration.ValidatingAdmissionPolicyBinding)(nil), (*v1alpha1.ValidatingAdmissionPolicyBinding)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_admissionregistration_ValidatingAdmissionPolicyBinding_To_v1alpha1_ValidatingAdmissionPolicyBinding(a.(*admissionregistration.ValidatingAdmissionPolicyBinding), b.(*v1alpha1.ValidatingAdmissionPolicyBinding), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.ValidatingAdmissionPolicyBindingList)(nil), (*admissionregistration.ValidatingAdmissionPolicyBindingList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ValidatingAdmissionPolicyBindingList_To_admissionregistration_ValidatingAdmissionPolicyBindingList(a.(*v1alpha1.ValidatingAdmissionPolicyBindingList), b.(*admissionregistration.ValidatingAdmissionPolicyBindingList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*admissionregistration.ValidatingAdmissionPolicyBindingList)(nil), (*v1alpha1.ValidatingAdmissionPolicyBindingList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_admissionregistration_ValidatingAdmissionPolicyBindingList_To_v1alpha1_ValidatingAdmissionPolicyBindingList(a.(*admissionregistration.ValidatingAdmissionPolicyBindingList), b.(*v1alpha1.ValidatingAdmissionPolicyBindingList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.ValidatingAdmissionPolicyBindingSpec)(nil), (*admissionregistration.ValidatingAdmissionPolicyBindingSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ValidatingAdmissionPolicyBindingSpec_To_admissionregistration_ValidatingAdmissionPolicyBindingSpec(a.(*v1alpha1.ValidatingAdmissionPolicyBindingSpec), b.(*admissionregistration.ValidatingAdmissionPolicyBindingSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*admissionregistration.ValidatingAdmissionPolicyBindingSpec)(nil), (*v1alpha1.ValidatingAdmissionPolicyBindingSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_admissionregistration_ValidatingAdmissionPolicyBindingSpec_To_v1alpha1_ValidatingAdmissionPolicyBindingSpec(a.(*admissionregistration.ValidatingAdmissionPolicyBindingSpec), b.(*v1alpha1.ValidatingAdmissionPolicyBindingSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.ValidatingAdmissionPolicyList)(nil), (*admissionregistration.ValidatingAdmissionPolicyList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ValidatingAdmissionPolicyList_To_admissionregistration_ValidatingAdmissionPolicyList(a.(*v1alpha1.ValidatingAdmissionPolicyList), b.(*admissionregistration.ValidatingAdmissionPolicyList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*admissionregistration.ValidatingAdmissionPolicyList)(nil), (*v1alpha1.ValidatingAdmissionPolicyList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_admissionregistration_ValidatingAdmissionPolicyList_To_v1alpha1_ValidatingAdmissionPolicyList(a.(*admissionregistration.ValidatingAdmissionPolicyList), b.(*v1alpha1.ValidatingAdmissionPolicyList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.ValidatingAdmissionPolicySpec)(nil), (*admissionregistration.ValidatingAdmissionPolicySpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ValidatingAdmissionPolicySpec_To_admissionregistration_ValidatingAdmissionPolicySpec(a.(*v1alpha1.ValidatingAdmissionPolicySpec), b.(*admissionregistration.ValidatingAdmissionPolicySpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*admissionregistration.ValidatingAdmissionPolicySpec)(nil), (*v1alpha1.ValidatingAdmissionPolicySpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_admissionregistration_ValidatingAdmissionPolicySpec_To_v1alpha1_ValidatingAdmissionPolicySpec(a.(*admissionregistration.ValidatingAdmissionPolicySpec), b.(*v1alpha1.ValidatingAdmissionPolicySpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha1.Validation)(nil), (*admissionregistration.Validation)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Validation_To_admissionregistration_Validation(a.(*v1alpha1.Validation), b.(*admissionregistration.Validation), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*admissionregistration.Validation)(nil), (*v1alpha1.Validation)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_admissionregistration_Validation_To_v1alpha1_Validation(a.(*admissionregistration.Validation), b.(*v1alpha1.Validation), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_AuditAnnotation_To_admissionregistration_AuditAnnotation(in *v1alpha1.AuditAnnotation, out *admissionregistration.AuditAnnotation, s conversion.Scope) error {
	out.Key = in.Key
	out.ValueExpression = in.ValueExpression
	return nil
}

// Convert_v1alpha1_AuditAnnotation_To_admissionregistration_AuditAnnotation is an autogenerated conversion function.
func Convert_v1alpha1_AuditAnnotation_To_admissionregistration_AuditAnnotation(in *v1alpha1.AuditAnnotation, out *admissionregistration.AuditAnnotation, s conversion.Scope) error {
	return autoConvert_v1alpha1_AuditAnnotation_To_admissionregistration_AuditAnnotation(in, out, s)
}

func autoConvert_admissionregistration_AuditAnnotation_To_v1alpha1_AuditAnnotation(in *admissionregistration.AuditAnnotation, out *v1alpha1.AuditAnnotation, s conversion.Scope) error {
	out.Key = in.Key
	out.ValueExpression = in.ValueExpression
	return nil
}

// Convert_admissionregistration_AuditAnnotation_To_v1alpha1_AuditAnnotation is an autogenerated conversion function.
func Convert_admissionregistration_AuditAnnotation_To_v1alpha1_AuditAnnotation(in *admissionregistration.AuditAnnotation, out *v1alpha1.AuditAnnotation, s conversion.Scope) error {
	return autoConvert_admissionregistration_AuditAnnotation_To_v1alpha1_AuditAnnotation(in, out, s)
}

func autoConvert_v1alpha1_MatchResources_To_admissionregistration_MatchResources(in *v1alpha1.MatchResources, out *admissionregistration.MatchResources, s conversion.Scope) error {
	out.NamespaceSelector = (*v1.LabelSelector)(unsafe.Pointer(in.NamespaceSelector))
	out.ObjectSelector = (*v1.LabelSelector)(unsafe.Pointer(in.ObjectSelector))
	if in.ResourceRules != nil {
		in, out := &in.ResourceRules, &out.ResourceRules
		*out = make([]admissionregistration.NamedRuleWithOperations, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_NamedRuleWithOperations_To_admissionregistration_NamedRuleWithOperations(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.ResourceRules = nil
	}
	if in.ExcludeResourceRules != nil {
		in, out := &in.ExcludeResourceRules, &out.ExcludeResourceRules
		*out = make([]admissionregistration.NamedRuleWithOperations, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_NamedRuleWithOperations_To_admissionregistration_NamedRuleWithOperations(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.ExcludeResourceRules = nil
	}
	out.MatchPolicy = (*admissionregistration.MatchPolicyType)(unsafe.Pointer(in.MatchPolicy))
	return nil
}

// Convert_v1alpha1_MatchResources_To_admissionregistration_MatchResources is an autogenerated conversion function.
func Convert_v1alpha1_MatchResources_To_admissionregistration_MatchResources(in *v1alpha1.MatchResources, out *admissionregistration.MatchResources, s conversion.Scope) error {
	return autoConvert_v1alpha1_MatchResources_To_admissionregistration_MatchResources(in, out, s)
}

func autoConvert_admissionregistration_MatchResources_To_v1alpha1_MatchResources(in *admissionregistration.MatchResources, out *v1alpha1.MatchResources, s conversion.Scope) error {
	out.NamespaceSelector = (*v1.LabelSelector)(unsafe.Pointer(in.NamespaceSelector))
	out.ObjectSelector = (*v1.LabelSelector)(unsafe.Pointer(in.ObjectSelector))
	if in.ResourceRules != nil {
		in, out := &in.ResourceRules, &out.ResourceRules
		*out = make([]v1alpha1.NamedRuleWithOperations, len(*in))
		for i := range *in {
			if err := Convert_admissionregistration_NamedRuleWithOperations_To_v1alpha1_NamedRuleWithOperations(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.ResourceRules = nil
	}
	if in.ExcludeResourceRules != nil {
		in, out := &in.ExcludeResourceRules, &out.ExcludeResourceRules
		*out = make([]v1alpha1.NamedRuleWithOperations, len(*in))
		for i := range *in {
			if err := Convert_admissionregistration_NamedRuleWithOperations_To_v1alpha1_NamedRuleWithOperations(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.ExcludeResourceRules = nil
	}
	out.MatchPolicy = (*v1alpha1.MatchPolicyType)(unsafe.Pointer(in.MatchPolicy))
	return nil
}

// Convert_admissionregistration_MatchResources_To_v1alpha1_MatchResources is an autogenerated conversion function.
func Convert_admissionregistration_MatchResources_To_v1alpha1_MatchResources(in *admissionregistration.MatchResources, out *v1alpha1.MatchResources, s conversion.Scope) error {
	return autoConvert_admissionregistration_MatchResources_To_v1alpha1_MatchResources(in, out, s)
}

func autoConvert_v1alpha1_NamedRuleWithOperations_To_admissionregistration_NamedRuleWithOperations(in *v1alpha1.NamedRuleWithOperations, out *admissionregistration.NamedRuleWithOperations, s conversion.Scope) error {
	out.ResourceNames = *(*[]string)(unsafe.Pointer(&in.ResourceNames))
	if err := admissionregistrationv1.Convert_v1_RuleWithOperations_To_admissionregistration_RuleWithOperations(&in.RuleWithOperations, &out.RuleWithOperations, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_NamedRuleWithOperations_To_admissionregistration_NamedRuleWithOperations is an autogenerated conversion function.
func Convert_v1alpha1_NamedRuleWithOperations_To_admissionregistration_NamedRuleWithOperations(in *v1alpha1.NamedRuleWithOperations, out *admissionregistration.NamedRuleWithOperations, s conversion.Scope) error {
	return autoConvert_v1alpha1_NamedRuleWithOperations_To_admissionregistration_NamedRuleWithOperations(in, out, s)
}

func autoConvert_admissionregistration_NamedRuleWithOperations_To_v1alpha1_NamedRuleWithOperations(in *admissionregistration.NamedRuleWithOperations, out *v1alpha1.NamedRuleWithOperations, s conversion.Scope) error {
	out.ResourceNames = *(*[]string)(unsafe.Pointer(&in.ResourceNames))
	if err := admissionregistrationv1.Convert_admissionregistration_RuleWithOperations_To_v1_RuleWithOperations(&in.RuleWithOperations, &out.RuleWithOperations, s); err != nil {
		return err
	}
	return nil
}

// Convert_admissionregistration_NamedRuleWithOperations_To_v1alpha1_NamedRuleWithOperations is an autogenerated conversion function.
func Convert_admissionregistration_NamedRuleWithOperations_To_v1alpha1_NamedRuleWithOperations(in *admissionregistration.NamedRuleWithOperations, out *v1alpha1.NamedRuleWithOperations, s conversion.Scope) error {
	return autoConvert_admissionregistration_NamedRuleWithOperations_To_v1alpha1_NamedRuleWithOperations(in, out, s)
}

func autoConvert_v1alpha1_ParamKind_To_admissionregistration_ParamKind(in *v1alpha1.ParamKind, out *admissionregistration.ParamKind, s conversion.Scope) error {
	out.APIVersion = in.APIVersion
	out.Kind = in.Kind
	return nil
}

// Convert_v1alpha1_ParamKind_To_admissionregistration_ParamKind is an autogenerated conversion function.
func Convert_v1alpha1_ParamKind_To_admissionregistration_ParamKind(in *v1alpha1.ParamKind, out *admissionregistration.ParamKind, s conversion.Scope) error {
	return autoConvert_v1alpha1_ParamKind_To_admissionregistration_ParamKind(in, out, s)
}

func autoConvert_admissionregistration_ParamKind_To_v1alpha1_ParamKind(in *admissionregistration.ParamKind, out *v1alpha1.ParamKind, s conversion.Scope) error {
	out.APIVersion = in.APIVersion
	out.Kind = in.Kind
	return nil
}

// Convert_admissionregistration_ParamKind_To_v1alpha1_ParamKind is an autogenerated conversion function.
func Convert_admissionregistration_ParamKind_To_v1alpha1_ParamKind(in *admissionregistration.ParamKind, out *v1alpha1.ParamKind, s conversion.Scope) error {
	return autoConvert_admissionregistration_ParamKind_To_v1alpha1_ParamKind(in, out, s)
}

func autoConvert_v1alpha1_ParamRef_To_admissionregistration_ParamRef(in *v1alpha1.ParamRef, out *admissionregistration.ParamRef, s conversion.Scope) error {
	out.Name = in.Name
	out.Namespace = in.Namespace
	return nil
}

// Convert_v1alpha1_ParamRef_To_admissionregistration_ParamRef is an autogenerated conversion function.
func Convert_v1alpha1_ParamRef_To_admissionregistration_ParamRef(in *v1alpha1.ParamRef, out *admissionregistration.ParamRef, s conversion.Scope) error {
	return autoConvert_v1alpha1_ParamRef_To_admissionregistration_ParamRef(in, out, s)
}

func autoConvert_admissionregistration_ParamRef_To_v1alpha1_ParamRef(in *admissionregistration.ParamRef, out *v1alpha1.ParamRef, s conversion.Scope) error {
	out.Name = in.Name
	out.Namespace = in.Namespace
	return nil
}

// Convert_admissionregistration_ParamRef_To_v1alpha1_ParamRef is an autogenerated conversion function.
func Convert_admissionregistration_ParamRef_To_v1alpha1_ParamRef(in *admissionregistration.ParamRef, out *v1alpha1.ParamRef, s conversion.Scope) error {
	return autoConvert_admissionregistration_ParamRef_To_v1alpha1_ParamRef(in, out, s)
}

func autoConvert_v1alpha1_ValidatingAdmissionPolicy_To_admissionregistration_ValidatingAdmissionPolicy(in *v1alpha1.ValidatingAdmissionPolicy, out *admissionregistration.ValidatingAdmissionPolicy, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_ValidatingAdmissionPolicySpec_To_admissionregistration_ValidatingAdmissionPolicySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ValidatingAdmissionPolicy_To_admissionregistration_ValidatingAdmissionPolicy is an autogenerated conversion function.
func Convert_v1alpha1_ValidatingAdmissionPolicy_To_admissionregistration_ValidatingAdmissionPolicy(in *v1alpha1.ValidatingAdmissionPolicy, out *admissionregistration.ValidatingAdmissionPolicy, s conversion.Scope) error {
	return autoConvert_v1alpha1_ValidatingAdmissionPolicy_To_admissionregistration_ValidatingAdmissionPolicy(in, out, s)
}

func autoConvert_admissionregistration_ValidatingAdmissionPolicy_To_v1alpha1_ValidatingAdmissionPolicy(in *admissionregistration.ValidatingAdmissionPolicy, out *v1alpha1.ValidatingAdmissionPolicy, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_admissionregistration_ValidatingAdmissionPolicySpec_To_v1alpha1_ValidatingAdmissionPolicySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_admissionregistration_ValidatingAdmissionPolicy_To_v1alpha1_ValidatingAdmissionPolicy is an autogenerated conversion function.
func Convert_admissionregistration_ValidatingAdmissionPolicy_To_v1alpha1_ValidatingAdmissionPolicy(in *admissionregistration.ValidatingAdmissionPolicy, out *v1alpha1.ValidatingAdmissionPolicy, s conversion.Scope) error {
	return autoConvert_admissionregistration_ValidatingAdmissionPolicy_To_v1alpha1_ValidatingAdmissionPolicy(in, out, s)
}

func autoConvert_v1alpha1_ValidatingAdmissionPolicyBinding_To_admissionregistration_ValidatingAdmissionPolicyBinding(in *v1alpha1.ValidatingAdmissionPolicyBinding, out *admissionregistration.ValidatingAdmissionPolicyBinding, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_ValidatingAdmissionPolicyBindingSpec_To_admissionregistration_ValidatingAdmissionPolicyBindingSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ValidatingAdmissionPolicyBinding_To_admissionregistration_ValidatingAdmissionPolicyBinding is an autogenerated conversion function.
func Convert_v1alpha1_ValidatingAdmissionPolicyBinding_To_admissionregistration_ValidatingAdmissionPolicyBinding(in *v1alpha1.ValidatingAdmissionPolicyBinding, out *admissionregistration.ValidatingAdmissionPolicyBinding, s conversion.Scope) error {
	return autoConvert_v1alpha1_ValidatingAdmissionPolicyBinding_To_admissionregistration_ValidatingAdmissionPolicyBinding(in, out, s)
}

func autoConvert_admissionregistration_ValidatingAdmissionPolicyBinding_To_v1alpha1_ValidatingAdmissionPolicyBinding(in *admissionregistration.ValidatingAdmissionPolicyBinding, out *v1alpha1.ValidatingAdmissionPolicyBinding, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_admissionregistration_ValidatingAdmissionPolicyBindingSpec_To_v1alpha1_ValidatingAdmissionPolicyBindingSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_admissionregistration_ValidatingAdmissionPolicyBinding_To_v1alpha1_ValidatingAdmissionPolicyBinding is an autogenerated conversion function.
func Convert_admissionregistration_ValidatingAdmissionPolicyBinding_To_v1alpha1_ValidatingAdmissionPolicyBinding(in *admissionregistration.ValidatingAdmissionPolicyBinding, out *v1alpha1.ValidatingAdmissionPolicyBinding, s conversion.Scope) error {
	return autoConvert_admissionregistration_ValidatingAdmissionPolicyBinding_To_v1alpha1_ValidatingAdmissionPolicyBinding(in, out, s)
}

func autoConvert_v1alpha1_ValidatingAdmissionPolicyBindingList_To_admissionregistration_ValidatingAdmissionPolicyBindingList(in *v1alpha1.ValidatingAdmissionPolicyBindingList, out *admissionregistration.ValidatingAdmissionPolicyBindingList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]admissionregistration.ValidatingAdmissionPolicyBinding, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_ValidatingAdmissionPolicyBinding_To_admissionregistration_ValidatingAdmissionPolicyBinding(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha1_ValidatingAdmissionPolicyBindingList_To_admissionregistration_ValidatingAdmissionPolicyBindingList is an autogenerated conversion function.
func Convert_v1alpha1_ValidatingAdmissionPolicyBindingList_To_admissionregistration_ValidatingAdmissionPolicyBindingList(in *v1alpha1.ValidatingAdmissionPolicyBindingList, out *admissionregistration.ValidatingAdmissionPolicyBindingList, s conversion.Scope) error {
	return autoConvert_v1alpha1_ValidatingAdmissionPolicyBindingList_To_admissionregistration_ValidatingAdmissionPolicyBindingList(in, out, s)
}

func autoConvert_admissionregistration_ValidatingAdmissionPolicyBindingList_To_v1alpha1_ValidatingAdmissionPolicyBindingList(in *admissionregistration.ValidatingAdmissionPolicyBindingList, out *v1alpha1.ValidatingAdmissionPolicyBindingList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1alpha1.ValidatingAdmissionPolicyBinding, len(*in))
		for i := range *in {
			if err := Convert_admissionregistration_ValidatingAdmissionPolicyBinding_To_v1alpha1_ValidatingAdmissionPolicyBinding(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_admissionregistration_ValidatingAdmissionPolicyBindingList_To_v1alpha1_ValidatingAdmissionPolicyBindingList is an autogenerated conversion function.
func Convert_admissionregistration_ValidatingAdmissionPolicyBindingList_To_v1alpha1_ValidatingAdmissionPolicyBindingList(in *admissionregistration.ValidatingAdmissionPolicyBindingList, out *v1alpha1.ValidatingAdmissionPolicyBindingList, s conversion.Scope) error {
	return autoConvert_admissionregistration_ValidatingAdmissionPolicyBindingList_To_v1alpha1_ValidatingAdmissionPolicyBindingList(in, out, s)
}

func autoConvert_v1alpha1_ValidatingAdmissionPolicyBindingSpec_To_admissionregistration_ValidatingAdmissionPolicyBindingSpec(in *v1alpha1.ValidatingAdmissionPolicyBindingSpec, out *admissionregistration.ValidatingAdmissionPolicyBindingSpec, s conversion.Scope) error {
	out.PolicyName = in.PolicyName
	out.ParamRef = (*admissionregistration.ParamRef)(unsafe.Pointer(in.ParamRef))
	if in.MatchResources != nil {
		in, out := &in.MatchResources, &out.MatchResources
		*out = new(admissionregistration.MatchResources)
		if err := Convert_v1alpha1_MatchResources_To_admissionregistration_MatchResources(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.MatchResources = nil
	}
	out.ValidationActions = *(*[]admissionregistration.ValidationAction)(unsafe.Pointer(&in.ValidationActions))
	return nil
}

// Convert_v1alpha1_ValidatingAdmissionPolicyBindingSpec_To_admissionregistration_ValidatingAdmissionPolicyBindingSpec is an autogenerated conversion function.
func Convert_v1alpha1_ValidatingAdmissionPolicyBindingSpec_To_admissionregistration_ValidatingAdmissionPolicyBindingSpec(in *v1alpha1.ValidatingAdmissionPolicyBindingSpec, out *admissionregistration.ValidatingAdmissionPolicyBindingSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_ValidatingAdmissionPolicyBindingSpec_To_admissionregistration_ValidatingAdmissionPolicyBindingSpec(in, out, s)
}

func autoConvert_admissionregistration_ValidatingAdmissionPolicyBindingSpec_To_v1alpha1_ValidatingAdmissionPolicyBindingSpec(in *admissionregistration.ValidatingAdmissionPolicyBindingSpec, out *v1alpha1.ValidatingAdmissionPolicyBindingSpec, s conversion.Scope) error {
	out.PolicyName = in.PolicyName
	out.ParamRef = (*v1alpha1.ParamRef)(unsafe.Pointer(in.ParamRef))
	if in.MatchResources != nil {
		in, out := &in.MatchResources, &out.MatchResources
		*out = new(v1alpha1.MatchResources)
		if err := Convert_admissionregistration_MatchResources_To_v1alpha1_MatchResources(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.MatchResources = nil
	}
	out.ValidationActions = *(*[]v1alpha1.ValidationAction)(unsafe.Pointer(&in.ValidationActions))
	return nil
}

// Convert_admissionregistration_ValidatingAdmissionPolicyBindingSpec_To_v1alpha1_ValidatingAdmissionPolicyBindingSpec is an autogenerated conversion function.
func Convert_admissionregistration_ValidatingAdmissionPolicyBindingSpec_To_v1alpha1_ValidatingAdmissionPolicyBindingSpec(in *admissionregistration.ValidatingAdmissionPolicyBindingSpec, out *v1alpha1.ValidatingAdmissionPolicyBindingSpec, s conversion.Scope) error {
	return autoConvert_admissionregistration_ValidatingAdmissionPolicyBindingSpec_To_v1alpha1_ValidatingAdmissionPolicyBindingSpec(in, out, s)
}

func autoConvert_v1alpha1_ValidatingAdmissionPolicyList_To_admissionregistration_ValidatingAdmissionPolicyList(in *v1alpha1.ValidatingAdmissionPolicyList, out *admissionregistration.ValidatingAdmissionPolicyList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]admissionregistration.ValidatingAdmissionPolicy, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_ValidatingAdmissionPolicy_To_admissionregistration_ValidatingAdmissionPolicy(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha1_ValidatingAdmissionPolicyList_To_admissionregistration_ValidatingAdmissionPolicyList is an autogenerated conversion function.
func Convert_v1alpha1_ValidatingAdmissionPolicyList_To_admissionregistration_ValidatingAdmissionPolicyList(in *v1alpha1.ValidatingAdmissionPolicyList, out *admissionregistration.ValidatingAdmissionPolicyList, s conversion.Scope) error {
	return autoConvert_v1alpha1_ValidatingAdmissionPolicyList_To_admissionregistration_ValidatingAdmissionPolicyList(in, out, s)
}

func autoConvert_admissionregistration_ValidatingAdmissionPolicyList_To_v1alpha1_ValidatingAdmissionPolicyList(in *admissionregistration.ValidatingAdmissionPolicyList, out *v1alpha1.ValidatingAdmissionPolicyList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1alpha1.ValidatingAdmissionPolicy, len(*in))
		for i := range *in {
			if err := Convert_admissionregistration_ValidatingAdmissionPolicy_To_v1alpha1_ValidatingAdmissionPolicy(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_admissionregistration_ValidatingAdmissionPolicyList_To_v1alpha1_ValidatingAdmissionPolicyList is an autogenerated conversion function.
func Convert_admissionregistration_ValidatingAdmissionPolicyList_To_v1alpha1_ValidatingAdmissionPolicyList(in *admissionregistration.ValidatingAdmissionPolicyList, out *v1alpha1.ValidatingAdmissionPolicyList, s conversion.Scope) error {
	return autoConvert_admissionregistration_ValidatingAdmissionPolicyList_To_v1alpha1_ValidatingAdmissionPolicyList(in, out, s)
}

func autoConvert_v1alpha1_ValidatingAdmissionPolicySpec_To_admissionregistration_ValidatingAdmissionPolicySpec(in *v1alpha1.ValidatingAdmissionPolicySpec, out *admissionregistration.ValidatingAdmissionPolicySpec, s conversion.Scope) error {
	out.ParamKind = (*admissionregistration.ParamKind)(unsafe.Pointer(in.ParamKind))
	if in.MatchConstraints != nil {
		in, out := &in.MatchConstraints, &out.MatchConstraints
		*out = new(admissionregistration.MatchResources)
		if err := Convert_v1alpha1_MatchResources_To_admissionregistration_MatchResources(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.MatchConstraints = nil
	}
	out.Validations = *(*[]admissionregistration.Validation)(unsafe.Pointer(&in.Validations))
	out.FailurePolicy = (*admissionregistration.FailurePolicyType)(unsafe.Pointer(in.FailurePolicy))
	out.AuditAnnotations = *(*[]admissionregistration.AuditAnnotation)(unsafe.Pointer(&in.AuditAnnotations))
	return nil
}

// Convert_v1alpha1_ValidatingAdmissionPolicySpec_To_admissionregistration_ValidatingAdmissionPolicySpec is an autogenerated conversion function.
func Convert_v1alpha1_ValidatingAdmissionPolicySpec_To_admissionregistration_ValidatingAdmissionPolicySpec(in *v1alpha1.ValidatingAdmissionPolicySpec, out *admissionregistration.ValidatingAdmissionPolicySpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_ValidatingAdmissionPolicySpec_To_admissionregistration_ValidatingAdmissionPolicySpec(in, out, s)
}

func autoConvert_admissionregistration_ValidatingAdmissionPolicySpec_To_v1alpha1_ValidatingAdmissionPolicySpec(in *admissionregistration.ValidatingAdmissionPolicySpec, out *v1alpha1.ValidatingAdmissionPolicySpec, s conversion.Scope) error {
	out.ParamKind = (*v1alpha1.ParamKind)(unsafe.Pointer(in.ParamKind))
	if in.MatchConstraints != nil {
		in, out := &in.MatchConstraints, &out.MatchConstraints
		*out = new(v1alpha1.MatchResources)
		if err := Convert_admissionregistration_MatchResources_To_v1alpha1_MatchResources(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.MatchConstraints = nil
	}
	out.Validations = *(*[]v1alpha1.Validation)(unsafe.Pointer(&in.Validations))
	out.FailurePolicy = (*v1alpha1.FailurePolicyType)(unsafe.Pointer(in.FailurePolicy))
	out.AuditAnnotations = *(*[]v1alpha1.AuditAnnotation)(unsafe.Pointer(&in.AuditAnnotations))
	return nil
}

// Convert_admissionregistration_ValidatingAdmissionPolicySpec_To_v1alpha1_ValidatingAdmissionPolicySpec is an autogenerated conversion function.
func Convert_admissionregistration_ValidatingAdmissionPolicySpec_To_v1alpha1_ValidatingAdmissionPolicySpec(in *admissionregistration.ValidatingAdmissionPolicySpec, out *v1alpha1.ValidatingAdmissionPolicySpec, s conversion.Scope) error {
	return autoConvert_admissionregistration_ValidatingAdmissionPolicySpec_To_v1alpha1_ValidatingAdmissionPolicySpec(in, out, s)
}

func autoConvert_v1alpha1_Validation_To_admissionregistration_Validation(in *v1alpha1.Validation, out *admissionregistration.Validation, s conversion.Scope) error {
	out.Expression = in.Expression
	out.Message = in.Message
	out.Reason = (*v1.StatusReason)(unsafe.Pointer(in.Reason))
	out.MessageExpression = in.MessageExpression
	return nil
}

// Convert_v1alpha1_Validation_To_admissionregistration_Validation is an autogenerated conversion function.
func Convert_v1alpha1_Validation_To_admissionregistration_Validation(in *v1alpha1.Validation, out *admissionregistration.Validation, s conversion.Scope) error {
	return autoConvert_v1alpha1_Validation_To_admissionregistration_Validation(in, out, s)
}

func autoConvert_admissionregistration_Validation_To_v1alpha1_Validation(in *admissionregistration.Validation, out *v1alpha1.Validation, s conversion.Scope) error {
	out.Expression = in.Expression
	out.Message = in.Message
	out.Reason = (*v1.StatusReason)(unsafe.Pointer(in.Reason))
	out.MessageExpression = in.MessageExpression
	return nil
}

// Convert_admissionregistration_Validation_To_v1alpha1_Validation is an autogenerated conversion function.
func Convert_admissionregistration_Validation_To_v1alpha1_Validation(in *admissionregistration.Validation, out *v1alpha1.Validation, s conversion.Scope) error {
	return autoConvert_admissionregistration_Validation_To_v1alpha1_Validation(in, out, s)
}
