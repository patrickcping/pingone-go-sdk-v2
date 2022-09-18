/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// EntityArrayEmbedded struct for EntityArrayEmbedded
type EntityArrayEmbedded struct {
	Actions []SignOnPolicyAction `json:"actions,omitempty"`
	Agreements []Agreement `json:"agreements,omitempty"`
	AlertChannels []AlertChannel `json:"alertChannels,omitempty"`
	Attributes []EntityArrayEmbeddedAttributesInner `json:"attributes,omitempty"`
	Applications []ReadOneApplication200Response `json:"applications,omitempty"`
	Certificates []Certificate `json:"certificates,omitempty"`
	Credentials []GatewayCredential `json:"credentials,omitempty"`
	CustomDomains []CustomDomain `json:"customDomains,omitempty"`
	EmailDomains []EmailDomain `json:"emailDomains,omitempty"`
	Environments []Environment `json:"environments,omitempty"`
	GatewayInstances []GatewayInstance `json:"gatewayInstances,omitempty"`
	Gateways []EntityArrayEmbeddedGatewaysInner `json:"gateways,omitempty"`
	Grants []ApplicationResourceGrant `json:"grants,omitempty"`
	Groups []Group `json:"groups,omitempty"`
	GroupMemberships []GroupMembership `json:"groupMemberships,omitempty"`
	IdentityProviders []IdentityProvider `json:"identityProviders,omitempty"`
	Keys []Certificate `json:"keys,omitempty"`
	Languages []AgreementLanguage `json:"languages,omitempty"`
	Organizations []Organization `json:"organizations,omitempty"`
	PasswordPolicies []PasswordPolicy `json:"passwordPolicies,omitempty"`
	Populations []Population `json:"populations,omitempty"`
	Resources []Resource `json:"resources,omitempty"`
	Revisions []AgreementLanguageRevision `json:"revisions,omitempty"`
	Scopes []ResourceScope `json:"scopes,omitempty"`
	SignOnPolicies []SignOnPolicy `json:"signOnPolicies,omitempty"`
	SignOnPolicyAssignments []SignOnPolicyAssignment `json:"signOnPolicyAssignments,omitempty"`
	Subscriptions []Subscription `json:"subscriptions,omitempty"`
	TrustedEmails []EmailDomainTrustedEmail `json:"trustedEmails,omitempty"`
	RoleAssignments []RoleAssignment `json:"roleAssignments,omitempty"`
	Roles []Role `json:"roles,omitempty"`
	Schemas []Schema `json:"schemas,omitempty"`
	Users []User `json:"users,omitempty"`
}

// NewEntityArrayEmbedded instantiates a new EntityArrayEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityArrayEmbedded() *EntityArrayEmbedded {
	this := EntityArrayEmbedded{}
	return &this
}

// NewEntityArrayEmbeddedWithDefaults instantiates a new EntityArrayEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityArrayEmbeddedWithDefaults() *EntityArrayEmbedded {
	this := EntityArrayEmbedded{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetActions() []SignOnPolicyAction {
	if o == nil || o.Actions == nil {
		var ret []SignOnPolicyAction
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetActionsOk() ([]SignOnPolicyAction, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given []SignOnPolicyAction and assigns it to the Actions field.
func (o *EntityArrayEmbedded) SetActions(v []SignOnPolicyAction) {
	o.Actions = v
}

// GetAgreements returns the Agreements field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetAgreements() []Agreement {
	if o == nil || o.Agreements == nil {
		var ret []Agreement
		return ret
	}
	return o.Agreements
}

// GetAgreementsOk returns a tuple with the Agreements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetAgreementsOk() ([]Agreement, bool) {
	if o == nil || o.Agreements == nil {
		return nil, false
	}
	return o.Agreements, true
}

// HasAgreements returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasAgreements() bool {
	if o != nil && o.Agreements != nil {
		return true
	}

	return false
}

// SetAgreements gets a reference to the given []Agreement and assigns it to the Agreements field.
func (o *EntityArrayEmbedded) SetAgreements(v []Agreement) {
	o.Agreements = v
}

// GetAlertChannels returns the AlertChannels field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetAlertChannels() []AlertChannel {
	if o == nil || o.AlertChannels == nil {
		var ret []AlertChannel
		return ret
	}
	return o.AlertChannels
}

// GetAlertChannelsOk returns a tuple with the AlertChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetAlertChannelsOk() ([]AlertChannel, bool) {
	if o == nil || o.AlertChannels == nil {
		return nil, false
	}
	return o.AlertChannels, true
}

// HasAlertChannels returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasAlertChannels() bool {
	if o != nil && o.AlertChannels != nil {
		return true
	}

	return false
}

// SetAlertChannels gets a reference to the given []AlertChannel and assigns it to the AlertChannels field.
func (o *EntityArrayEmbedded) SetAlertChannels(v []AlertChannel) {
	o.AlertChannels = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetAttributes() []EntityArrayEmbeddedAttributesInner {
	if o == nil || o.Attributes == nil {
		var ret []EntityArrayEmbeddedAttributesInner
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetAttributesOk() ([]EntityArrayEmbeddedAttributesInner, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []EntityArrayEmbeddedAttributesInner and assigns it to the Attributes field.
func (o *EntityArrayEmbedded) SetAttributes(v []EntityArrayEmbeddedAttributesInner) {
	o.Attributes = v
}

// GetApplications returns the Applications field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetApplications() []ReadOneApplication200Response {
	if o == nil || o.Applications == nil {
		var ret []ReadOneApplication200Response
		return ret
	}
	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetApplicationsOk() ([]ReadOneApplication200Response, bool) {
	if o == nil || o.Applications == nil {
		return nil, false
	}
	return o.Applications, true
}

// HasApplications returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasApplications() bool {
	if o != nil && o.Applications != nil {
		return true
	}

	return false
}

// SetApplications gets a reference to the given []ReadOneApplication200Response and assigns it to the Applications field.
func (o *EntityArrayEmbedded) SetApplications(v []ReadOneApplication200Response) {
	o.Applications = v
}

// GetCertificates returns the Certificates field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetCertificates() []Certificate {
	if o == nil || o.Certificates == nil {
		var ret []Certificate
		return ret
	}
	return o.Certificates
}

// GetCertificatesOk returns a tuple with the Certificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetCertificatesOk() ([]Certificate, bool) {
	if o == nil || o.Certificates == nil {
		return nil, false
	}
	return o.Certificates, true
}

// HasCertificates returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasCertificates() bool {
	if o != nil && o.Certificates != nil {
		return true
	}

	return false
}

// SetCertificates gets a reference to the given []Certificate and assigns it to the Certificates field.
func (o *EntityArrayEmbedded) SetCertificates(v []Certificate) {
	o.Certificates = v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetCredentials() []GatewayCredential {
	if o == nil || o.Credentials == nil {
		var ret []GatewayCredential
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetCredentialsOk() ([]GatewayCredential, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given []GatewayCredential and assigns it to the Credentials field.
func (o *EntityArrayEmbedded) SetCredentials(v []GatewayCredential) {
	o.Credentials = v
}

// GetCustomDomains returns the CustomDomains field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetCustomDomains() []CustomDomain {
	if o == nil || o.CustomDomains == nil {
		var ret []CustomDomain
		return ret
	}
	return o.CustomDomains
}

// GetCustomDomainsOk returns a tuple with the CustomDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetCustomDomainsOk() ([]CustomDomain, bool) {
	if o == nil || o.CustomDomains == nil {
		return nil, false
	}
	return o.CustomDomains, true
}

// HasCustomDomains returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasCustomDomains() bool {
	if o != nil && o.CustomDomains != nil {
		return true
	}

	return false
}

// SetCustomDomains gets a reference to the given []CustomDomain and assigns it to the CustomDomains field.
func (o *EntityArrayEmbedded) SetCustomDomains(v []CustomDomain) {
	o.CustomDomains = v
}

// GetEmailDomains returns the EmailDomains field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetEmailDomains() []EmailDomain {
	if o == nil || o.EmailDomains == nil {
		var ret []EmailDomain
		return ret
	}
	return o.EmailDomains
}

// GetEmailDomainsOk returns a tuple with the EmailDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetEmailDomainsOk() ([]EmailDomain, bool) {
	if o == nil || o.EmailDomains == nil {
		return nil, false
	}
	return o.EmailDomains, true
}

// HasEmailDomains returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasEmailDomains() bool {
	if o != nil && o.EmailDomains != nil {
		return true
	}

	return false
}

// SetEmailDomains gets a reference to the given []EmailDomain and assigns it to the EmailDomains field.
func (o *EntityArrayEmbedded) SetEmailDomains(v []EmailDomain) {
	o.EmailDomains = v
}

// GetEnvironments returns the Environments field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetEnvironments() []Environment {
	if o == nil || o.Environments == nil {
		var ret []Environment
		return ret
	}
	return o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetEnvironmentsOk() ([]Environment, bool) {
	if o == nil || o.Environments == nil {
		return nil, false
	}
	return o.Environments, true
}

// HasEnvironments returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasEnvironments() bool {
	if o != nil && o.Environments != nil {
		return true
	}

	return false
}

// SetEnvironments gets a reference to the given []Environment and assigns it to the Environments field.
func (o *EntityArrayEmbedded) SetEnvironments(v []Environment) {
	o.Environments = v
}

// GetGatewayInstances returns the GatewayInstances field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetGatewayInstances() []GatewayInstance {
	if o == nil || o.GatewayInstances == nil {
		var ret []GatewayInstance
		return ret
	}
	return o.GatewayInstances
}

// GetGatewayInstancesOk returns a tuple with the GatewayInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetGatewayInstancesOk() ([]GatewayInstance, bool) {
	if o == nil || o.GatewayInstances == nil {
		return nil, false
	}
	return o.GatewayInstances, true
}

// HasGatewayInstances returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasGatewayInstances() bool {
	if o != nil && o.GatewayInstances != nil {
		return true
	}

	return false
}

// SetGatewayInstances gets a reference to the given []GatewayInstance and assigns it to the GatewayInstances field.
func (o *EntityArrayEmbedded) SetGatewayInstances(v []GatewayInstance) {
	o.GatewayInstances = v
}

// GetGateways returns the Gateways field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetGateways() []EntityArrayEmbeddedGatewaysInner {
	if o == nil || o.Gateways == nil {
		var ret []EntityArrayEmbeddedGatewaysInner
		return ret
	}
	return o.Gateways
}

// GetGatewaysOk returns a tuple with the Gateways field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetGatewaysOk() ([]EntityArrayEmbeddedGatewaysInner, bool) {
	if o == nil || o.Gateways == nil {
		return nil, false
	}
	return o.Gateways, true
}

// HasGateways returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasGateways() bool {
	if o != nil && o.Gateways != nil {
		return true
	}

	return false
}

// SetGateways gets a reference to the given []EntityArrayEmbeddedGatewaysInner and assigns it to the Gateways field.
func (o *EntityArrayEmbedded) SetGateways(v []EntityArrayEmbeddedGatewaysInner) {
	o.Gateways = v
}

// GetGrants returns the Grants field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetGrants() []ApplicationResourceGrant {
	if o == nil || o.Grants == nil {
		var ret []ApplicationResourceGrant
		return ret
	}
	return o.Grants
}

// GetGrantsOk returns a tuple with the Grants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetGrantsOk() ([]ApplicationResourceGrant, bool) {
	if o == nil || o.Grants == nil {
		return nil, false
	}
	return o.Grants, true
}

// HasGrants returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasGrants() bool {
	if o != nil && o.Grants != nil {
		return true
	}

	return false
}

// SetGrants gets a reference to the given []ApplicationResourceGrant and assigns it to the Grants field.
func (o *EntityArrayEmbedded) SetGrants(v []ApplicationResourceGrant) {
	o.Grants = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetGroups() []Group {
	if o == nil || o.Groups == nil {
		var ret []Group
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetGroupsOk() ([]Group, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []Group and assigns it to the Groups field.
func (o *EntityArrayEmbedded) SetGroups(v []Group) {
	o.Groups = v
}

// GetGroupMemberships returns the GroupMemberships field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetGroupMemberships() []GroupMembership {
	if o == nil || o.GroupMemberships == nil {
		var ret []GroupMembership
		return ret
	}
	return o.GroupMemberships
}

// GetGroupMembershipsOk returns a tuple with the GroupMemberships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetGroupMembershipsOk() ([]GroupMembership, bool) {
	if o == nil || o.GroupMemberships == nil {
		return nil, false
	}
	return o.GroupMemberships, true
}

// HasGroupMemberships returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasGroupMemberships() bool {
	if o != nil && o.GroupMemberships != nil {
		return true
	}

	return false
}

// SetGroupMemberships gets a reference to the given []GroupMembership and assigns it to the GroupMemberships field.
func (o *EntityArrayEmbedded) SetGroupMemberships(v []GroupMembership) {
	o.GroupMemberships = v
}

// GetIdentityProviders returns the IdentityProviders field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetIdentityProviders() []IdentityProvider {
	if o == nil || o.IdentityProviders == nil {
		var ret []IdentityProvider
		return ret
	}
	return o.IdentityProviders
}

// GetIdentityProvidersOk returns a tuple with the IdentityProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetIdentityProvidersOk() ([]IdentityProvider, bool) {
	if o == nil || o.IdentityProviders == nil {
		return nil, false
	}
	return o.IdentityProviders, true
}

// HasIdentityProviders returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasIdentityProviders() bool {
	if o != nil && o.IdentityProviders != nil {
		return true
	}

	return false
}

// SetIdentityProviders gets a reference to the given []IdentityProvider and assigns it to the IdentityProviders field.
func (o *EntityArrayEmbedded) SetIdentityProviders(v []IdentityProvider) {
	o.IdentityProviders = v
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetKeys() []Certificate {
	if o == nil || o.Keys == nil {
		var ret []Certificate
		return ret
	}
	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetKeysOk() ([]Certificate, bool) {
	if o == nil || o.Keys == nil {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasKeys() bool {
	if o != nil && o.Keys != nil {
		return true
	}

	return false
}

// SetKeys gets a reference to the given []Certificate and assigns it to the Keys field.
func (o *EntityArrayEmbedded) SetKeys(v []Certificate) {
	o.Keys = v
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetLanguages() []AgreementLanguage {
	if o == nil || o.Languages == nil {
		var ret []AgreementLanguage
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetLanguagesOk() ([]AgreementLanguage, bool) {
	if o == nil || o.Languages == nil {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasLanguages() bool {
	if o != nil && o.Languages != nil {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []AgreementLanguage and assigns it to the Languages field.
func (o *EntityArrayEmbedded) SetLanguages(v []AgreementLanguage) {
	o.Languages = v
}

// GetOrganizations returns the Organizations field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetOrganizations() []Organization {
	if o == nil || o.Organizations == nil {
		var ret []Organization
		return ret
	}
	return o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetOrganizationsOk() ([]Organization, bool) {
	if o == nil || o.Organizations == nil {
		return nil, false
	}
	return o.Organizations, true
}

// HasOrganizations returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasOrganizations() bool {
	if o != nil && o.Organizations != nil {
		return true
	}

	return false
}

// SetOrganizations gets a reference to the given []Organization and assigns it to the Organizations field.
func (o *EntityArrayEmbedded) SetOrganizations(v []Organization) {
	o.Organizations = v
}

// GetPasswordPolicies returns the PasswordPolicies field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetPasswordPolicies() []PasswordPolicy {
	if o == nil || o.PasswordPolicies == nil {
		var ret []PasswordPolicy
		return ret
	}
	return o.PasswordPolicies
}

// GetPasswordPoliciesOk returns a tuple with the PasswordPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetPasswordPoliciesOk() ([]PasswordPolicy, bool) {
	if o == nil || o.PasswordPolicies == nil {
		return nil, false
	}
	return o.PasswordPolicies, true
}

// HasPasswordPolicies returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasPasswordPolicies() bool {
	if o != nil && o.PasswordPolicies != nil {
		return true
	}

	return false
}

// SetPasswordPolicies gets a reference to the given []PasswordPolicy and assigns it to the PasswordPolicies field.
func (o *EntityArrayEmbedded) SetPasswordPolicies(v []PasswordPolicy) {
	o.PasswordPolicies = v
}

// GetPopulations returns the Populations field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetPopulations() []Population {
	if o == nil || o.Populations == nil {
		var ret []Population
		return ret
	}
	return o.Populations
}

// GetPopulationsOk returns a tuple with the Populations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetPopulationsOk() ([]Population, bool) {
	if o == nil || o.Populations == nil {
		return nil, false
	}
	return o.Populations, true
}

// HasPopulations returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasPopulations() bool {
	if o != nil && o.Populations != nil {
		return true
	}

	return false
}

// SetPopulations gets a reference to the given []Population and assigns it to the Populations field.
func (o *EntityArrayEmbedded) SetPopulations(v []Population) {
	o.Populations = v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetResources() []Resource {
	if o == nil || o.Resources == nil {
		var ret []Resource
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetResourcesOk() ([]Resource, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasResources() bool {
	if o != nil && o.Resources != nil {
		return true
	}

	return false
}

// SetResources gets a reference to the given []Resource and assigns it to the Resources field.
func (o *EntityArrayEmbedded) SetResources(v []Resource) {
	o.Resources = v
}

// GetRevisions returns the Revisions field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetRevisions() []AgreementLanguageRevision {
	if o == nil || o.Revisions == nil {
		var ret []AgreementLanguageRevision
		return ret
	}
	return o.Revisions
}

// GetRevisionsOk returns a tuple with the Revisions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetRevisionsOk() ([]AgreementLanguageRevision, bool) {
	if o == nil || o.Revisions == nil {
		return nil, false
	}
	return o.Revisions, true
}

// HasRevisions returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasRevisions() bool {
	if o != nil && o.Revisions != nil {
		return true
	}

	return false
}

// SetRevisions gets a reference to the given []AgreementLanguageRevision and assigns it to the Revisions field.
func (o *EntityArrayEmbedded) SetRevisions(v []AgreementLanguageRevision) {
	o.Revisions = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetScopes() []ResourceScope {
	if o == nil || o.Scopes == nil {
		var ret []ResourceScope
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetScopesOk() ([]ResourceScope, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []ResourceScope and assigns it to the Scopes field.
func (o *EntityArrayEmbedded) SetScopes(v []ResourceScope) {
	o.Scopes = v
}

// GetSignOnPolicies returns the SignOnPolicies field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetSignOnPolicies() []SignOnPolicy {
	if o == nil || o.SignOnPolicies == nil {
		var ret []SignOnPolicy
		return ret
	}
	return o.SignOnPolicies
}

// GetSignOnPoliciesOk returns a tuple with the SignOnPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetSignOnPoliciesOk() ([]SignOnPolicy, bool) {
	if o == nil || o.SignOnPolicies == nil {
		return nil, false
	}
	return o.SignOnPolicies, true
}

// HasSignOnPolicies returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasSignOnPolicies() bool {
	if o != nil && o.SignOnPolicies != nil {
		return true
	}

	return false
}

// SetSignOnPolicies gets a reference to the given []SignOnPolicy and assigns it to the SignOnPolicies field.
func (o *EntityArrayEmbedded) SetSignOnPolicies(v []SignOnPolicy) {
	o.SignOnPolicies = v
}

// GetSignOnPolicyAssignments returns the SignOnPolicyAssignments field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetSignOnPolicyAssignments() []SignOnPolicyAssignment {
	if o == nil || o.SignOnPolicyAssignments == nil {
		var ret []SignOnPolicyAssignment
		return ret
	}
	return o.SignOnPolicyAssignments
}

// GetSignOnPolicyAssignmentsOk returns a tuple with the SignOnPolicyAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetSignOnPolicyAssignmentsOk() ([]SignOnPolicyAssignment, bool) {
	if o == nil || o.SignOnPolicyAssignments == nil {
		return nil, false
	}
	return o.SignOnPolicyAssignments, true
}

// HasSignOnPolicyAssignments returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasSignOnPolicyAssignments() bool {
	if o != nil && o.SignOnPolicyAssignments != nil {
		return true
	}

	return false
}

// SetSignOnPolicyAssignments gets a reference to the given []SignOnPolicyAssignment and assigns it to the SignOnPolicyAssignments field.
func (o *EntityArrayEmbedded) SetSignOnPolicyAssignments(v []SignOnPolicyAssignment) {
	o.SignOnPolicyAssignments = v
}

// GetSubscriptions returns the Subscriptions field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetSubscriptions() []Subscription {
	if o == nil || o.Subscriptions == nil {
		var ret []Subscription
		return ret
	}
	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetSubscriptionsOk() ([]Subscription, bool) {
	if o == nil || o.Subscriptions == nil {
		return nil, false
	}
	return o.Subscriptions, true
}

// HasSubscriptions returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasSubscriptions() bool {
	if o != nil && o.Subscriptions != nil {
		return true
	}

	return false
}

// SetSubscriptions gets a reference to the given []Subscription and assigns it to the Subscriptions field.
func (o *EntityArrayEmbedded) SetSubscriptions(v []Subscription) {
	o.Subscriptions = v
}

// GetTrustedEmails returns the TrustedEmails field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetTrustedEmails() []EmailDomainTrustedEmail {
	if o == nil || o.TrustedEmails == nil {
		var ret []EmailDomainTrustedEmail
		return ret
	}
	return o.TrustedEmails
}

// GetTrustedEmailsOk returns a tuple with the TrustedEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetTrustedEmailsOk() ([]EmailDomainTrustedEmail, bool) {
	if o == nil || o.TrustedEmails == nil {
		return nil, false
	}
	return o.TrustedEmails, true
}

// HasTrustedEmails returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasTrustedEmails() bool {
	if o != nil && o.TrustedEmails != nil {
		return true
	}

	return false
}

// SetTrustedEmails gets a reference to the given []EmailDomainTrustedEmail and assigns it to the TrustedEmails field.
func (o *EntityArrayEmbedded) SetTrustedEmails(v []EmailDomainTrustedEmail) {
	o.TrustedEmails = v
}

// GetRoleAssignments returns the RoleAssignments field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetRoleAssignments() []RoleAssignment {
	if o == nil || o.RoleAssignments == nil {
		var ret []RoleAssignment
		return ret
	}
	return o.RoleAssignments
}

// GetRoleAssignmentsOk returns a tuple with the RoleAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetRoleAssignmentsOk() ([]RoleAssignment, bool) {
	if o == nil || o.RoleAssignments == nil {
		return nil, false
	}
	return o.RoleAssignments, true
}

// HasRoleAssignments returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasRoleAssignments() bool {
	if o != nil && o.RoleAssignments != nil {
		return true
	}

	return false
}

// SetRoleAssignments gets a reference to the given []RoleAssignment and assigns it to the RoleAssignments field.
func (o *EntityArrayEmbedded) SetRoleAssignments(v []RoleAssignment) {
	o.RoleAssignments = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetRoles() []Role {
	if o == nil || o.Roles == nil {
		var ret []Role
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetRolesOk() ([]Role, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []Role and assigns it to the Roles field.
func (o *EntityArrayEmbedded) SetRoles(v []Role) {
	o.Roles = v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetSchemas() []Schema {
	if o == nil || o.Schemas == nil {
		var ret []Schema
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetSchemasOk() ([]Schema, bool) {
	if o == nil || o.Schemas == nil {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasSchemas() bool {
	if o != nil && o.Schemas != nil {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []Schema and assigns it to the Schemas field.
func (o *EntityArrayEmbedded) SetSchemas(v []Schema) {
	o.Schemas = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetUsers() []User {
	if o == nil || o.Users == nil {
		var ret []User
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetUsersOk() ([]User, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []User and assigns it to the Users field.
func (o *EntityArrayEmbedded) SetUsers(v []User) {
	o.Users = v
}

func (o EntityArrayEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.Agreements != nil {
		toSerialize["agreements"] = o.Agreements
	}
	if o.AlertChannels != nil {
		toSerialize["alertChannels"] = o.AlertChannels
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Applications != nil {
		toSerialize["applications"] = o.Applications
	}
	if o.Certificates != nil {
		toSerialize["certificates"] = o.Certificates
	}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	if o.CustomDomains != nil {
		toSerialize["customDomains"] = o.CustomDomains
	}
	if o.EmailDomains != nil {
		toSerialize["emailDomains"] = o.EmailDomains
	}
	if o.Environments != nil {
		toSerialize["environments"] = o.Environments
	}
	if o.GatewayInstances != nil {
		toSerialize["gatewayInstances"] = o.GatewayInstances
	}
	if o.Gateways != nil {
		toSerialize["gateways"] = o.Gateways
	}
	if o.Grants != nil {
		toSerialize["grants"] = o.Grants
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.GroupMemberships != nil {
		toSerialize["groupMemberships"] = o.GroupMemberships
	}
	if o.IdentityProviders != nil {
		toSerialize["identityProviders"] = o.IdentityProviders
	}
	if o.Keys != nil {
		toSerialize["keys"] = o.Keys
	}
	if o.Languages != nil {
		toSerialize["languages"] = o.Languages
	}
	if o.Organizations != nil {
		toSerialize["organizations"] = o.Organizations
	}
	if o.PasswordPolicies != nil {
		toSerialize["passwordPolicies"] = o.PasswordPolicies
	}
	if o.Populations != nil {
		toSerialize["populations"] = o.Populations
	}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	if o.Revisions != nil {
		toSerialize["revisions"] = o.Revisions
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if o.SignOnPolicies != nil {
		toSerialize["signOnPolicies"] = o.SignOnPolicies
	}
	if o.SignOnPolicyAssignments != nil {
		toSerialize["signOnPolicyAssignments"] = o.SignOnPolicyAssignments
	}
	if o.Subscriptions != nil {
		toSerialize["subscriptions"] = o.Subscriptions
	}
	if o.TrustedEmails != nil {
		toSerialize["trustedEmails"] = o.TrustedEmails
	}
	if o.RoleAssignments != nil {
		toSerialize["roleAssignments"] = o.RoleAssignments
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.Schemas != nil {
		toSerialize["schemas"] = o.Schemas
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	return json.Marshal(toSerialize)
}

type NullableEntityArrayEmbedded struct {
	value *EntityArrayEmbedded
	isSet bool
}

func (v NullableEntityArrayEmbedded) Get() *EntityArrayEmbedded {
	return v.value
}

func (v *NullableEntityArrayEmbedded) Set(val *EntityArrayEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityArrayEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityArrayEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityArrayEmbedded(val *EntityArrayEmbedded) *NullableEntityArrayEmbedded {
	return &NullableEntityArrayEmbedded{value: val, isSet: true}
}

func (v NullableEntityArrayEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityArrayEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


