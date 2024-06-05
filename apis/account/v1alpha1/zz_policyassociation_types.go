/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PolicyAssociationInitParameters struct {

	// The account to operate on.
	// The account Id.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Comma separated list of the account’s available policies. These policies can be applied to the websites in the account.
	// e.g. available_policy_ids = format("%s,%s", incapsula_policy.acl1-policy.id, incapsula_policy.waf3-policy.id)
	// Specify this argument only for a parent account trying to update policy availability for its subaccounts. To remove availability for all policies, specify "NO_AVAILABLE_POLICIES".
	// Comma separated list of The account’s available policies. These policies can be applied to the websites in the account. e.g. available_policy_ids = format("%s,%s", incapsula_policy.acl1-policy.id, incapsula_policy.waf3-policy.id) Specify this argument only if you are a parent account trying to update your child account policies availability in order to remove availability for all policies please specify "NO_AVAILABLE_POLICIES".
	AvailablePolicyIds *string `json:"availablePolicyIds,omitempty" tf:"available_policy_ids,omitempty"`

	// This list is currently relevant to Allowlist and ACL policies. More than one policy can be set as default.
	// The default policies can be set for the current account, or if used by users with credentials of the parent account can also be set for sub-accounts.
	// Default setting – empty list. No default policy. Providing an empty list or omitting this argument will clear all the non-mandatory default policies.
	// This list is currently relevant to whitelist and acl policies. More than one policy can be set as default. providing an empty list or omitting this argument will clear all the non mandatory default policies.
	// +crossplane:generate:reference:type=github.com/kernel-kun/provider-incapsula/apis/policy/v1alpha1.Policy
	// +listType=set
	DefaultNonMandatoryPolicyIds []*string `json:"defaultNonMandatoryPolicyIds,omitempty" tf:"default_non_mandatory_policy_ids,omitempty"`

	// References to Policy in policy to populate defaultNonMandatoryPolicyIds.
	// +kubebuilder:validation:Optional
	DefaultNonMandatoryPolicyIdsRefs []v1.Reference `json:"defaultNonMandatoryPolicyIdsRefs,omitempty" tf:"-"`

	// Selector for a list of Policy in policy to populate defaultNonMandatoryPolicyIds.
	// +kubebuilder:validation:Optional
	DefaultNonMandatoryPolicyIdsSelector *v1.Selector `json:"defaultNonMandatoryPolicyIdsSelector,omitempty" tf:"-"`

	// The WAF policy which is set as default for the account. The account can only have 1 such ID.
	// The Default policy will be applied automatically to sites that are created after setting it to default.
	// This default setting can be set for the current account, or if used by users with credentials of the parent account can also be set for sub-accounts.
	// This parameter is MANDATORY for customers that have account level WAF RULES policies enabled. This means that a default WAF RULES policy resource must be created.
	// For customers who were not migrated yet, this parameter should not be set. However, after migration, a WAF RULES policy must be added and set as default.
	// Default setting - None.
	// The WAF policy which is set as default to the account. The account can only have 1 such id.
	// The Default policy will be applied automatically to sites that were create after setting it to default.
	DefaultWafPolicyID *string `json:"defaultWafPolicyId,omitempty" tf:"default_waf_policy_id,omitempty"`
}

type PolicyAssociationObservation struct {

	// The account to operate on.
	// The account Id.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Comma separated list of the account’s available policies. These policies can be applied to the websites in the account.
	// e.g. available_policy_ids = format("%s,%s", incapsula_policy.acl1-policy.id, incapsula_policy.waf3-policy.id)
	// Specify this argument only for a parent account trying to update policy availability for its subaccounts. To remove availability for all policies, specify "NO_AVAILABLE_POLICIES".
	// Comma separated list of The account’s available policies. These policies can be applied to the websites in the account. e.g. available_policy_ids = format("%s,%s", incapsula_policy.acl1-policy.id, incapsula_policy.waf3-policy.id) Specify this argument only if you are a parent account trying to update your child account policies availability in order to remove availability for all policies please specify "NO_AVAILABLE_POLICIES".
	AvailablePolicyIds *string `json:"availablePolicyIds,omitempty" tf:"available_policy_ids,omitempty"`

	// This list is currently relevant to Allowlist and ACL policies. More than one policy can be set as default.
	// The default policies can be set for the current account, or if used by users with credentials of the parent account can also be set for sub-accounts.
	// Default setting – empty list. No default policy. Providing an empty list or omitting this argument will clear all the non-mandatory default policies.
	// This list is currently relevant to whitelist and acl policies. More than one policy can be set as default. providing an empty list or omitting this argument will clear all the non mandatory default policies.
	// +listType=set
	DefaultNonMandatoryPolicyIds []*string `json:"defaultNonMandatoryPolicyIds,omitempty" tf:"default_non_mandatory_policy_ids,omitempty"`

	// The WAF policy which is set as default for the account. The account can only have 1 such ID.
	// The Default policy will be applied automatically to sites that are created after setting it to default.
	// This default setting can be set for the current account, or if used by users with credentials of the parent account can also be set for sub-accounts.
	// This parameter is MANDATORY for customers that have account level WAF RULES policies enabled. This means that a default WAF RULES policy resource must be created.
	// For customers who were not migrated yet, this parameter should not be set. However, after migration, a WAF RULES policy must be added and set as default.
	// Default setting - None.
	// The WAF policy which is set as default to the account. The account can only have 1 such id.
	// The Default policy will be applied automatically to sites that were create after setting it to default.
	DefaultWafPolicyID *string `json:"defaultWafPolicyId,omitempty" tf:"default_waf_policy_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PolicyAssociationParameters struct {

	// The account to operate on.
	// The account Id.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Comma separated list of the account’s available policies. These policies can be applied to the websites in the account.
	// e.g. available_policy_ids = format("%s,%s", incapsula_policy.acl1-policy.id, incapsula_policy.waf3-policy.id)
	// Specify this argument only for a parent account trying to update policy availability for its subaccounts. To remove availability for all policies, specify "NO_AVAILABLE_POLICIES".
	// Comma separated list of The account’s available policies. These policies can be applied to the websites in the account. e.g. available_policy_ids = format("%s,%s", incapsula_policy.acl1-policy.id, incapsula_policy.waf3-policy.id) Specify this argument only if you are a parent account trying to update your child account policies availability in order to remove availability for all policies please specify "NO_AVAILABLE_POLICIES".
	// +kubebuilder:validation:Optional
	AvailablePolicyIds *string `json:"availablePolicyIds,omitempty" tf:"available_policy_ids,omitempty"`

	// This list is currently relevant to Allowlist and ACL policies. More than one policy can be set as default.
	// The default policies can be set for the current account, or if used by users with credentials of the parent account can also be set for sub-accounts.
	// Default setting – empty list. No default policy. Providing an empty list or omitting this argument will clear all the non-mandatory default policies.
	// This list is currently relevant to whitelist and acl policies. More than one policy can be set as default. providing an empty list or omitting this argument will clear all the non mandatory default policies.
	// +crossplane:generate:reference:type=github.com/kernel-kun/provider-incapsula/apis/policy/v1alpha1.Policy
	// +kubebuilder:validation:Optional
	// +listType=set
	DefaultNonMandatoryPolicyIds []*string `json:"defaultNonMandatoryPolicyIds,omitempty" tf:"default_non_mandatory_policy_ids,omitempty"`

	// References to Policy in policy to populate defaultNonMandatoryPolicyIds.
	// +kubebuilder:validation:Optional
	DefaultNonMandatoryPolicyIdsRefs []v1.Reference `json:"defaultNonMandatoryPolicyIdsRefs,omitempty" tf:"-"`

	// Selector for a list of Policy in policy to populate defaultNonMandatoryPolicyIds.
	// +kubebuilder:validation:Optional
	DefaultNonMandatoryPolicyIdsSelector *v1.Selector `json:"defaultNonMandatoryPolicyIdsSelector,omitempty" tf:"-"`

	// The WAF policy which is set as default for the account. The account can only have 1 such ID.
	// The Default policy will be applied automatically to sites that are created after setting it to default.
	// This default setting can be set for the current account, or if used by users with credentials of the parent account can also be set for sub-accounts.
	// This parameter is MANDATORY for customers that have account level WAF RULES policies enabled. This means that a default WAF RULES policy resource must be created.
	// For customers who were not migrated yet, this parameter should not be set. However, after migration, a WAF RULES policy must be added and set as default.
	// Default setting - None.
	// The WAF policy which is set as default to the account. The account can only have 1 such id.
	// The Default policy will be applied automatically to sites that were create after setting it to default.
	// +kubebuilder:validation:Optional
	DefaultWafPolicyID *string `json:"defaultWafPolicyId,omitempty" tf:"default_waf_policy_id,omitempty"`
}

// PolicyAssociationSpec defines the desired state of PolicyAssociation
type PolicyAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyAssociationParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider PolicyAssociationInitParameters `json:"initProvider,omitempty"`
}

// PolicyAssociationStatus defines the observed state of PolicyAssociation.
type PolicyAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PolicyAssociation is the Schema for the PolicyAssociations API. Provides an Incapsula Account Policy Association resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,incapsula}
type PolicyAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.accountId) || (has(self.initProvider) && has(self.initProvider.accountId))",message="spec.forProvider.accountId is a required parameter"
	Spec   PolicyAssociationSpec   `json:"spec"`
	Status PolicyAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyAssociationList contains a list of PolicyAssociations
type PolicyAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyAssociation `json:"items"`
}

// Repository type metadata.
var (
	PolicyAssociation_Kind             = "PolicyAssociation"
	PolicyAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PolicyAssociation_Kind}.String()
	PolicyAssociation_KindAPIVersion   = PolicyAssociation_Kind + "." + CRDGroupVersion.String()
	PolicyAssociation_GroupVersionKind = CRDGroupVersion.WithKind(PolicyAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&PolicyAssociation{}, &PolicyAssociationList{})
}
