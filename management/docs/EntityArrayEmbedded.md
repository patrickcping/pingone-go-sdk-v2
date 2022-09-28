# EntityArrayEmbedded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]SignOnPolicyAction**](SignOnPolicyAction.md) |  | [optional] 
**Agreements** | Pointer to [**[]Agreement**](Agreement.md) |  | [optional] 
**AlertChannels** | Pointer to [**[]AlertChannel**](AlertChannel.md) |  | [optional] 
**Attributes** | Pointer to [**[]EntityArrayEmbeddedAttributesInner**](EntityArrayEmbeddedAttributesInner.md) |  | [optional] 
**Applications** | Pointer to [**[]ReadOneApplication200Response**](ReadOneApplication200Response.md) |  | [optional] 
**Certificates** | Pointer to [**[]Certificate**](Certificate.md) |  | [optional] 
**Credentials** | Pointer to [**[]GatewayCredential**](GatewayCredential.md) |  | [optional] 
**CustomDomains** | Pointer to [**[]CustomDomain**](CustomDomain.md) |  | [optional] 
**EmailDomains** | Pointer to [**[]EmailDomain**](EmailDomain.md) |  | [optional] 
**Environments** | Pointer to [**[]Environment**](Environment.md) |  | [optional] 
**GatewayInstances** | Pointer to [**[]GatewayInstance**](GatewayInstance.md) |  | [optional] 
**Gateways** | Pointer to [**[]EntityArrayEmbeddedGatewaysInner**](EntityArrayEmbeddedGatewaysInner.md) |  | [optional] 
**Grants** | Pointer to [**[]ApplicationResourceGrant**](ApplicationResourceGrant.md) |  | [optional] 
**Groups** | Pointer to [**[]Group**](Group.md) |  | [optional] 
**GroupMemberships** | Pointer to [**[]GroupMembership**](GroupMembership.md) |  | [optional] 
**IdentityProviders** | Pointer to [**[]IdentityProvider**](IdentityProvider.md) |  | [optional] 
**Keys** | Pointer to [**[]Certificate**](Certificate.md) |  | [optional] 
**Languages** | Pointer to [**[]AgreementLanguage**](AgreementLanguage.md) |  | [optional] 
**Licenses** | Pointer to [**[]License**](License.md) |  | [optional] 
**Organizations** | Pointer to [**[]Organization**](Organization.md) |  | [optional] 
**PasswordPolicies** | Pointer to [**[]PasswordPolicy**](PasswordPolicy.md) |  | [optional] 
**Populations** | Pointer to [**[]Population**](Population.md) |  | [optional] 
**Resources** | Pointer to [**[]Resource**](Resource.md) |  | [optional] 
**Revisions** | Pointer to [**[]AgreementLanguageRevision**](AgreementLanguageRevision.md) |  | [optional] 
**Scopes** | Pointer to [**[]ResourceScope**](ResourceScope.md) |  | [optional] 
**SignOnPolicies** | Pointer to [**[]SignOnPolicy**](SignOnPolicy.md) |  | [optional] 
**SignOnPolicyAssignments** | Pointer to [**[]SignOnPolicyAssignment**](SignOnPolicyAssignment.md) |  | [optional] 
**Subscriptions** | Pointer to [**[]Subscription**](Subscription.md) |  | [optional] 
**TrustedEmails** | Pointer to [**[]EmailDomainTrustedEmail**](EmailDomainTrustedEmail.md) |  | [optional] 
**RoleAssignments** | Pointer to [**[]RoleAssignment**](RoleAssignment.md) |  | [optional] 
**Roles** | Pointer to [**[]Role**](Role.md) |  | [optional] 
**Schemas** | Pointer to [**[]Schema**](Schema.md) |  | [optional] 
**Users** | Pointer to [**[]User**](User.md) |  | [optional] 

## Methods

### NewEntityArrayEmbedded

`func NewEntityArrayEmbedded() *EntityArrayEmbedded`

NewEntityArrayEmbedded instantiates a new EntityArrayEmbedded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityArrayEmbeddedWithDefaults

`func NewEntityArrayEmbeddedWithDefaults() *EntityArrayEmbedded`

NewEntityArrayEmbeddedWithDefaults instantiates a new EntityArrayEmbedded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *EntityArrayEmbedded) GetActions() []SignOnPolicyAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *EntityArrayEmbedded) GetActionsOk() (*[]SignOnPolicyAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *EntityArrayEmbedded) SetActions(v []SignOnPolicyAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *EntityArrayEmbedded) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetAgreements

`func (o *EntityArrayEmbedded) GetAgreements() []Agreement`

GetAgreements returns the Agreements field if non-nil, zero value otherwise.

### GetAgreementsOk

`func (o *EntityArrayEmbedded) GetAgreementsOk() (*[]Agreement, bool)`

GetAgreementsOk returns a tuple with the Agreements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreements

`func (o *EntityArrayEmbedded) SetAgreements(v []Agreement)`

SetAgreements sets Agreements field to given value.

### HasAgreements

`func (o *EntityArrayEmbedded) HasAgreements() bool`

HasAgreements returns a boolean if a field has been set.

### GetAlertChannels

`func (o *EntityArrayEmbedded) GetAlertChannels() []AlertChannel`

GetAlertChannels returns the AlertChannels field if non-nil, zero value otherwise.

### GetAlertChannelsOk

`func (o *EntityArrayEmbedded) GetAlertChannelsOk() (*[]AlertChannel, bool)`

GetAlertChannelsOk returns a tuple with the AlertChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertChannels

`func (o *EntityArrayEmbedded) SetAlertChannels(v []AlertChannel)`

SetAlertChannels sets AlertChannels field to given value.

### HasAlertChannels

`func (o *EntityArrayEmbedded) HasAlertChannels() bool`

HasAlertChannels returns a boolean if a field has been set.

### GetAttributes

`func (o *EntityArrayEmbedded) GetAttributes() []EntityArrayEmbeddedAttributesInner`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EntityArrayEmbedded) GetAttributesOk() (*[]EntityArrayEmbeddedAttributesInner, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EntityArrayEmbedded) SetAttributes(v []EntityArrayEmbeddedAttributesInner)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *EntityArrayEmbedded) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetApplications

`func (o *EntityArrayEmbedded) GetApplications() []ReadOneApplication200Response`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *EntityArrayEmbedded) GetApplicationsOk() (*[]ReadOneApplication200Response, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *EntityArrayEmbedded) SetApplications(v []ReadOneApplication200Response)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *EntityArrayEmbedded) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetCertificates

`func (o *EntityArrayEmbedded) GetCertificates() []Certificate`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *EntityArrayEmbedded) GetCertificatesOk() (*[]Certificate, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *EntityArrayEmbedded) SetCertificates(v []Certificate)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *EntityArrayEmbedded) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### GetCredentials

`func (o *EntityArrayEmbedded) GetCredentials() []GatewayCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *EntityArrayEmbedded) GetCredentialsOk() (*[]GatewayCredential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *EntityArrayEmbedded) SetCredentials(v []GatewayCredential)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *EntityArrayEmbedded) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetCustomDomains

`func (o *EntityArrayEmbedded) GetCustomDomains() []CustomDomain`

GetCustomDomains returns the CustomDomains field if non-nil, zero value otherwise.

### GetCustomDomainsOk

`func (o *EntityArrayEmbedded) GetCustomDomainsOk() (*[]CustomDomain, bool)`

GetCustomDomainsOk returns a tuple with the CustomDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDomains

`func (o *EntityArrayEmbedded) SetCustomDomains(v []CustomDomain)`

SetCustomDomains sets CustomDomains field to given value.

### HasCustomDomains

`func (o *EntityArrayEmbedded) HasCustomDomains() bool`

HasCustomDomains returns a boolean if a field has been set.

### GetEmailDomains

`func (o *EntityArrayEmbedded) GetEmailDomains() []EmailDomain`

GetEmailDomains returns the EmailDomains field if non-nil, zero value otherwise.

### GetEmailDomainsOk

`func (o *EntityArrayEmbedded) GetEmailDomainsOk() (*[]EmailDomain, bool)`

GetEmailDomainsOk returns a tuple with the EmailDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailDomains

`func (o *EntityArrayEmbedded) SetEmailDomains(v []EmailDomain)`

SetEmailDomains sets EmailDomains field to given value.

### HasEmailDomains

`func (o *EntityArrayEmbedded) HasEmailDomains() bool`

HasEmailDomains returns a boolean if a field has been set.

### GetEnvironments

`func (o *EntityArrayEmbedded) GetEnvironments() []Environment`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *EntityArrayEmbedded) GetEnvironmentsOk() (*[]Environment, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *EntityArrayEmbedded) SetEnvironments(v []Environment)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *EntityArrayEmbedded) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetGatewayInstances

`func (o *EntityArrayEmbedded) GetGatewayInstances() []GatewayInstance`

GetGatewayInstances returns the GatewayInstances field if non-nil, zero value otherwise.

### GetGatewayInstancesOk

`func (o *EntityArrayEmbedded) GetGatewayInstancesOk() (*[]GatewayInstance, bool)`

GetGatewayInstancesOk returns a tuple with the GatewayInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayInstances

`func (o *EntityArrayEmbedded) SetGatewayInstances(v []GatewayInstance)`

SetGatewayInstances sets GatewayInstances field to given value.

### HasGatewayInstances

`func (o *EntityArrayEmbedded) HasGatewayInstances() bool`

HasGatewayInstances returns a boolean if a field has been set.

### GetGateways

`func (o *EntityArrayEmbedded) GetGateways() []EntityArrayEmbeddedGatewaysInner`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *EntityArrayEmbedded) GetGatewaysOk() (*[]EntityArrayEmbeddedGatewaysInner, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *EntityArrayEmbedded) SetGateways(v []EntityArrayEmbeddedGatewaysInner)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *EntityArrayEmbedded) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetGrants

`func (o *EntityArrayEmbedded) GetGrants() []ApplicationResourceGrant`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### GetGrantsOk

`func (o *EntityArrayEmbedded) GetGrantsOk() (*[]ApplicationResourceGrant, bool)`

GetGrantsOk returns a tuple with the Grants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrants

`func (o *EntityArrayEmbedded) SetGrants(v []ApplicationResourceGrant)`

SetGrants sets Grants field to given value.

### HasGrants

`func (o *EntityArrayEmbedded) HasGrants() bool`

HasGrants returns a boolean if a field has been set.

### GetGroups

`func (o *EntityArrayEmbedded) GetGroups() []Group`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *EntityArrayEmbedded) GetGroupsOk() (*[]Group, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *EntityArrayEmbedded) SetGroups(v []Group)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *EntityArrayEmbedded) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetGroupMemberships

`func (o *EntityArrayEmbedded) GetGroupMemberships() []GroupMembership`

GetGroupMemberships returns the GroupMemberships field if non-nil, zero value otherwise.

### GetGroupMembershipsOk

`func (o *EntityArrayEmbedded) GetGroupMembershipsOk() (*[]GroupMembership, bool)`

GetGroupMembershipsOk returns a tuple with the GroupMemberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberships

`func (o *EntityArrayEmbedded) SetGroupMemberships(v []GroupMembership)`

SetGroupMemberships sets GroupMemberships field to given value.

### HasGroupMemberships

`func (o *EntityArrayEmbedded) HasGroupMemberships() bool`

HasGroupMemberships returns a boolean if a field has been set.

### GetIdentityProviders

`func (o *EntityArrayEmbedded) GetIdentityProviders() []IdentityProvider`

GetIdentityProviders returns the IdentityProviders field if non-nil, zero value otherwise.

### GetIdentityProvidersOk

`func (o *EntityArrayEmbedded) GetIdentityProvidersOk() (*[]IdentityProvider, bool)`

GetIdentityProvidersOk returns a tuple with the IdentityProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviders

`func (o *EntityArrayEmbedded) SetIdentityProviders(v []IdentityProvider)`

SetIdentityProviders sets IdentityProviders field to given value.

### HasIdentityProviders

`func (o *EntityArrayEmbedded) HasIdentityProviders() bool`

HasIdentityProviders returns a boolean if a field has been set.

### GetKeys

`func (o *EntityArrayEmbedded) GetKeys() []Certificate`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *EntityArrayEmbedded) GetKeysOk() (*[]Certificate, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *EntityArrayEmbedded) SetKeys(v []Certificate)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *EntityArrayEmbedded) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetLanguages

`func (o *EntityArrayEmbedded) GetLanguages() []AgreementLanguage`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *EntityArrayEmbedded) GetLanguagesOk() (*[]AgreementLanguage, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *EntityArrayEmbedded) SetLanguages(v []AgreementLanguage)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *EntityArrayEmbedded) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetLicenses

`func (o *EntityArrayEmbedded) GetLicenses() []License`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *EntityArrayEmbedded) GetLicensesOk() (*[]License, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *EntityArrayEmbedded) SetLicenses(v []License)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *EntityArrayEmbedded) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetOrganizations

`func (o *EntityArrayEmbedded) GetOrganizations() []Organization`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *EntityArrayEmbedded) GetOrganizationsOk() (*[]Organization, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *EntityArrayEmbedded) SetOrganizations(v []Organization)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *EntityArrayEmbedded) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetPasswordPolicies

`func (o *EntityArrayEmbedded) GetPasswordPolicies() []PasswordPolicy`

GetPasswordPolicies returns the PasswordPolicies field if non-nil, zero value otherwise.

### GetPasswordPoliciesOk

`func (o *EntityArrayEmbedded) GetPasswordPoliciesOk() (*[]PasswordPolicy, bool)`

GetPasswordPoliciesOk returns a tuple with the PasswordPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicies

`func (o *EntityArrayEmbedded) SetPasswordPolicies(v []PasswordPolicy)`

SetPasswordPolicies sets PasswordPolicies field to given value.

### HasPasswordPolicies

`func (o *EntityArrayEmbedded) HasPasswordPolicies() bool`

HasPasswordPolicies returns a boolean if a field has been set.

### GetPopulations

`func (o *EntityArrayEmbedded) GetPopulations() []Population`

GetPopulations returns the Populations field if non-nil, zero value otherwise.

### GetPopulationsOk

`func (o *EntityArrayEmbedded) GetPopulationsOk() (*[]Population, bool)`

GetPopulationsOk returns a tuple with the Populations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulations

`func (o *EntityArrayEmbedded) SetPopulations(v []Population)`

SetPopulations sets Populations field to given value.

### HasPopulations

`func (o *EntityArrayEmbedded) HasPopulations() bool`

HasPopulations returns a boolean if a field has been set.

### GetResources

`func (o *EntityArrayEmbedded) GetResources() []Resource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *EntityArrayEmbedded) GetResourcesOk() (*[]Resource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *EntityArrayEmbedded) SetResources(v []Resource)`

SetResources sets Resources field to given value.

### HasResources

`func (o *EntityArrayEmbedded) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRevisions

`func (o *EntityArrayEmbedded) GetRevisions() []AgreementLanguageRevision`

GetRevisions returns the Revisions field if non-nil, zero value otherwise.

### GetRevisionsOk

`func (o *EntityArrayEmbedded) GetRevisionsOk() (*[]AgreementLanguageRevision, bool)`

GetRevisionsOk returns a tuple with the Revisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisions

`func (o *EntityArrayEmbedded) SetRevisions(v []AgreementLanguageRevision)`

SetRevisions sets Revisions field to given value.

### HasRevisions

`func (o *EntityArrayEmbedded) HasRevisions() bool`

HasRevisions returns a boolean if a field has been set.

### GetScopes

`func (o *EntityArrayEmbedded) GetScopes() []ResourceScope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *EntityArrayEmbedded) GetScopesOk() (*[]ResourceScope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *EntityArrayEmbedded) SetScopes(v []ResourceScope)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *EntityArrayEmbedded) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSignOnPolicies

`func (o *EntityArrayEmbedded) GetSignOnPolicies() []SignOnPolicy`

GetSignOnPolicies returns the SignOnPolicies field if non-nil, zero value otherwise.

### GetSignOnPoliciesOk

`func (o *EntityArrayEmbedded) GetSignOnPoliciesOk() (*[]SignOnPolicy, bool)`

GetSignOnPoliciesOk returns a tuple with the SignOnPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnPolicies

`func (o *EntityArrayEmbedded) SetSignOnPolicies(v []SignOnPolicy)`

SetSignOnPolicies sets SignOnPolicies field to given value.

### HasSignOnPolicies

`func (o *EntityArrayEmbedded) HasSignOnPolicies() bool`

HasSignOnPolicies returns a boolean if a field has been set.

### GetSignOnPolicyAssignments

`func (o *EntityArrayEmbedded) GetSignOnPolicyAssignments() []SignOnPolicyAssignment`

GetSignOnPolicyAssignments returns the SignOnPolicyAssignments field if non-nil, zero value otherwise.

### GetSignOnPolicyAssignmentsOk

`func (o *EntityArrayEmbedded) GetSignOnPolicyAssignmentsOk() (*[]SignOnPolicyAssignment, bool)`

GetSignOnPolicyAssignmentsOk returns a tuple with the SignOnPolicyAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnPolicyAssignments

`func (o *EntityArrayEmbedded) SetSignOnPolicyAssignments(v []SignOnPolicyAssignment)`

SetSignOnPolicyAssignments sets SignOnPolicyAssignments field to given value.

### HasSignOnPolicyAssignments

`func (o *EntityArrayEmbedded) HasSignOnPolicyAssignments() bool`

HasSignOnPolicyAssignments returns a boolean if a field has been set.

### GetSubscriptions

`func (o *EntityArrayEmbedded) GetSubscriptions() []Subscription`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *EntityArrayEmbedded) GetSubscriptionsOk() (*[]Subscription, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *EntityArrayEmbedded) SetSubscriptions(v []Subscription)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *EntityArrayEmbedded) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetTrustedEmails

`func (o *EntityArrayEmbedded) GetTrustedEmails() []EmailDomainTrustedEmail`

GetTrustedEmails returns the TrustedEmails field if non-nil, zero value otherwise.

### GetTrustedEmailsOk

`func (o *EntityArrayEmbedded) GetTrustedEmailsOk() (*[]EmailDomainTrustedEmail, bool)`

GetTrustedEmailsOk returns a tuple with the TrustedEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedEmails

`func (o *EntityArrayEmbedded) SetTrustedEmails(v []EmailDomainTrustedEmail)`

SetTrustedEmails sets TrustedEmails field to given value.

### HasTrustedEmails

`func (o *EntityArrayEmbedded) HasTrustedEmails() bool`

HasTrustedEmails returns a boolean if a field has been set.

### GetRoleAssignments

`func (o *EntityArrayEmbedded) GetRoleAssignments() []RoleAssignment`

GetRoleAssignments returns the RoleAssignments field if non-nil, zero value otherwise.

### GetRoleAssignmentsOk

`func (o *EntityArrayEmbedded) GetRoleAssignmentsOk() (*[]RoleAssignment, bool)`

GetRoleAssignmentsOk returns a tuple with the RoleAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAssignments

`func (o *EntityArrayEmbedded) SetRoleAssignments(v []RoleAssignment)`

SetRoleAssignments sets RoleAssignments field to given value.

### HasRoleAssignments

`func (o *EntityArrayEmbedded) HasRoleAssignments() bool`

HasRoleAssignments returns a boolean if a field has been set.

### GetRoles

`func (o *EntityArrayEmbedded) GetRoles() []Role`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *EntityArrayEmbedded) GetRolesOk() (*[]Role, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *EntityArrayEmbedded) SetRoles(v []Role)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *EntityArrayEmbedded) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSchemas

`func (o *EntityArrayEmbedded) GetSchemas() []Schema`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *EntityArrayEmbedded) GetSchemasOk() (*[]Schema, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *EntityArrayEmbedded) SetSchemas(v []Schema)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *EntityArrayEmbedded) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetUsers

`func (o *EntityArrayEmbedded) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *EntityArrayEmbedded) GetUsersOk() (*[]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *EntityArrayEmbedded) SetUsers(v []User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *EntityArrayEmbedded) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


