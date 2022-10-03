# LicenseUsers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowPasswordManagementNotifications** | Pointer to **bool** | A read-only boolean that specifies whether the license supports sending password management notifications. | [optional] 
**AllowIdentityProviders** | Pointer to **bool** | A read-only boolean that specifies whether the license supports using external identity providers in the specified environment. | [optional] 
**AllowMyAccount** | Pointer to **bool** | A read-only boolean that specifies whether the license supports using My Account capabilities in the specified environment. | [optional] 
**AllowPasswordManagement** | Pointer to **bool** | A read-only boolean that specifies whether the license supports using password management capabilities in the specified environment. | [optional] 
**AllowPasswordOnlyAuthentication** | Pointer to **bool** | A read-only boolean that specifies whether the license supports using password only login capabilities in the specified environment. | [optional] 
**AllowPasswordPolicy** | Pointer to **bool** | A read-only boolean that specifies whether the license supports using password policies in the specified environment. | [optional] 
**AllowProvisioning** | Pointer to **bool** | A read-only boolean that specifies whether the license supports using provisioning capabilities in the specified environment. | [optional] 
**AllowInboundProvisioning** | Pointer to **bool** |  | [optional] 
**AllowRoleAssignment** | Pointer to **bool** | A read-only boolean that specifies whether the license supports role assignments in the specified environment. | [optional] 
**AllowVerificationFlow** | Pointer to **bool** | A read-only boolean that specifies whether the license supports using verification flows in the specified environment. | [optional] 
**AllowUpdateSelf** | Pointer to **bool** | A read-only boolean that specifies whether the license supports allowing users to update their own profile. | [optional] 
**EntitledToSupport** | Pointer to **bool** | A read-only boolean that specifies whether the license allows PingOne support. | [optional] 
**Max** | Pointer to **int32** | An read-only integer that specifies the maximum number of users allowed per environment. | [optional] 
**HardLimitMax** | Pointer to **int32** |  | [optional] 
**AnnualActiveIncluded** | Pointer to **int32** | A read-only integer that specifies a soft limit on the number of active identities across all environments on the license per year. This property is not visible if a value is not provided at the time the license is created. | [optional] 
**MonthlyActiveIncluded** | Pointer to **int32** | A read-only integer that specifies a soft limit on the number of active identities across all environments on the license per month. This property is not visible if a value is not provided at the time the license is created. | [optional] 

## Methods

### NewLicenseUsers

`func NewLicenseUsers() *LicenseUsers`

NewLicenseUsers instantiates a new LicenseUsers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseUsersWithDefaults

`func NewLicenseUsersWithDefaults() *LicenseUsers`

NewLicenseUsersWithDefaults instantiates a new LicenseUsers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowPasswordManagementNotifications

`func (o *LicenseUsers) GetAllowPasswordManagementNotifications() bool`

GetAllowPasswordManagementNotifications returns the AllowPasswordManagementNotifications field if non-nil, zero value otherwise.

### GetAllowPasswordManagementNotificationsOk

`func (o *LicenseUsers) GetAllowPasswordManagementNotificationsOk() (*bool, bool)`

GetAllowPasswordManagementNotificationsOk returns a tuple with the AllowPasswordManagementNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPasswordManagementNotifications

`func (o *LicenseUsers) SetAllowPasswordManagementNotifications(v bool)`

SetAllowPasswordManagementNotifications sets AllowPasswordManagementNotifications field to given value.

### HasAllowPasswordManagementNotifications

`func (o *LicenseUsers) HasAllowPasswordManagementNotifications() bool`

HasAllowPasswordManagementNotifications returns a boolean if a field has been set.

### GetAllowIdentityProviders

`func (o *LicenseUsers) GetAllowIdentityProviders() bool`

GetAllowIdentityProviders returns the AllowIdentityProviders field if non-nil, zero value otherwise.

### GetAllowIdentityProvidersOk

`func (o *LicenseUsers) GetAllowIdentityProvidersOk() (*bool, bool)`

GetAllowIdentityProvidersOk returns a tuple with the AllowIdentityProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowIdentityProviders

`func (o *LicenseUsers) SetAllowIdentityProviders(v bool)`

SetAllowIdentityProviders sets AllowIdentityProviders field to given value.

### HasAllowIdentityProviders

`func (o *LicenseUsers) HasAllowIdentityProviders() bool`

HasAllowIdentityProviders returns a boolean if a field has been set.

### GetAllowMyAccount

`func (o *LicenseUsers) GetAllowMyAccount() bool`

GetAllowMyAccount returns the AllowMyAccount field if non-nil, zero value otherwise.

### GetAllowMyAccountOk

`func (o *LicenseUsers) GetAllowMyAccountOk() (*bool, bool)`

GetAllowMyAccountOk returns a tuple with the AllowMyAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMyAccount

`func (o *LicenseUsers) SetAllowMyAccount(v bool)`

SetAllowMyAccount sets AllowMyAccount field to given value.

### HasAllowMyAccount

`func (o *LicenseUsers) HasAllowMyAccount() bool`

HasAllowMyAccount returns a boolean if a field has been set.

### GetAllowPasswordManagement

`func (o *LicenseUsers) GetAllowPasswordManagement() bool`

GetAllowPasswordManagement returns the AllowPasswordManagement field if non-nil, zero value otherwise.

### GetAllowPasswordManagementOk

`func (o *LicenseUsers) GetAllowPasswordManagementOk() (*bool, bool)`

GetAllowPasswordManagementOk returns a tuple with the AllowPasswordManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPasswordManagement

`func (o *LicenseUsers) SetAllowPasswordManagement(v bool)`

SetAllowPasswordManagement sets AllowPasswordManagement field to given value.

### HasAllowPasswordManagement

`func (o *LicenseUsers) HasAllowPasswordManagement() bool`

HasAllowPasswordManagement returns a boolean if a field has been set.

### GetAllowPasswordOnlyAuthentication

`func (o *LicenseUsers) GetAllowPasswordOnlyAuthentication() bool`

GetAllowPasswordOnlyAuthentication returns the AllowPasswordOnlyAuthentication field if non-nil, zero value otherwise.

### GetAllowPasswordOnlyAuthenticationOk

`func (o *LicenseUsers) GetAllowPasswordOnlyAuthenticationOk() (*bool, bool)`

GetAllowPasswordOnlyAuthenticationOk returns a tuple with the AllowPasswordOnlyAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPasswordOnlyAuthentication

`func (o *LicenseUsers) SetAllowPasswordOnlyAuthentication(v bool)`

SetAllowPasswordOnlyAuthentication sets AllowPasswordOnlyAuthentication field to given value.

### HasAllowPasswordOnlyAuthentication

`func (o *LicenseUsers) HasAllowPasswordOnlyAuthentication() bool`

HasAllowPasswordOnlyAuthentication returns a boolean if a field has been set.

### GetAllowPasswordPolicy

`func (o *LicenseUsers) GetAllowPasswordPolicy() bool`

GetAllowPasswordPolicy returns the AllowPasswordPolicy field if non-nil, zero value otherwise.

### GetAllowPasswordPolicyOk

`func (o *LicenseUsers) GetAllowPasswordPolicyOk() (*bool, bool)`

GetAllowPasswordPolicyOk returns a tuple with the AllowPasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPasswordPolicy

`func (o *LicenseUsers) SetAllowPasswordPolicy(v bool)`

SetAllowPasswordPolicy sets AllowPasswordPolicy field to given value.

### HasAllowPasswordPolicy

`func (o *LicenseUsers) HasAllowPasswordPolicy() bool`

HasAllowPasswordPolicy returns a boolean if a field has been set.

### GetAllowProvisioning

`func (o *LicenseUsers) GetAllowProvisioning() bool`

GetAllowProvisioning returns the AllowProvisioning field if non-nil, zero value otherwise.

### GetAllowProvisioningOk

`func (o *LicenseUsers) GetAllowProvisioningOk() (*bool, bool)`

GetAllowProvisioningOk returns a tuple with the AllowProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowProvisioning

`func (o *LicenseUsers) SetAllowProvisioning(v bool)`

SetAllowProvisioning sets AllowProvisioning field to given value.

### HasAllowProvisioning

`func (o *LicenseUsers) HasAllowProvisioning() bool`

HasAllowProvisioning returns a boolean if a field has been set.

### GetAllowInboundProvisioning

`func (o *LicenseUsers) GetAllowInboundProvisioning() bool`

GetAllowInboundProvisioning returns the AllowInboundProvisioning field if non-nil, zero value otherwise.

### GetAllowInboundProvisioningOk

`func (o *LicenseUsers) GetAllowInboundProvisioningOk() (*bool, bool)`

GetAllowInboundProvisioningOk returns a tuple with the AllowInboundProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInboundProvisioning

`func (o *LicenseUsers) SetAllowInboundProvisioning(v bool)`

SetAllowInboundProvisioning sets AllowInboundProvisioning field to given value.

### HasAllowInboundProvisioning

`func (o *LicenseUsers) HasAllowInboundProvisioning() bool`

HasAllowInboundProvisioning returns a boolean if a field has been set.

### GetAllowRoleAssignment

`func (o *LicenseUsers) GetAllowRoleAssignment() bool`

GetAllowRoleAssignment returns the AllowRoleAssignment field if non-nil, zero value otherwise.

### GetAllowRoleAssignmentOk

`func (o *LicenseUsers) GetAllowRoleAssignmentOk() (*bool, bool)`

GetAllowRoleAssignmentOk returns a tuple with the AllowRoleAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRoleAssignment

`func (o *LicenseUsers) SetAllowRoleAssignment(v bool)`

SetAllowRoleAssignment sets AllowRoleAssignment field to given value.

### HasAllowRoleAssignment

`func (o *LicenseUsers) HasAllowRoleAssignment() bool`

HasAllowRoleAssignment returns a boolean if a field has been set.

### GetAllowVerificationFlow

`func (o *LicenseUsers) GetAllowVerificationFlow() bool`

GetAllowVerificationFlow returns the AllowVerificationFlow field if non-nil, zero value otherwise.

### GetAllowVerificationFlowOk

`func (o *LicenseUsers) GetAllowVerificationFlowOk() (*bool, bool)`

GetAllowVerificationFlowOk returns a tuple with the AllowVerificationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowVerificationFlow

`func (o *LicenseUsers) SetAllowVerificationFlow(v bool)`

SetAllowVerificationFlow sets AllowVerificationFlow field to given value.

### HasAllowVerificationFlow

`func (o *LicenseUsers) HasAllowVerificationFlow() bool`

HasAllowVerificationFlow returns a boolean if a field has been set.

### GetAllowUpdateSelf

`func (o *LicenseUsers) GetAllowUpdateSelf() bool`

GetAllowUpdateSelf returns the AllowUpdateSelf field if non-nil, zero value otherwise.

### GetAllowUpdateSelfOk

`func (o *LicenseUsers) GetAllowUpdateSelfOk() (*bool, bool)`

GetAllowUpdateSelfOk returns a tuple with the AllowUpdateSelf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUpdateSelf

`func (o *LicenseUsers) SetAllowUpdateSelf(v bool)`

SetAllowUpdateSelf sets AllowUpdateSelf field to given value.

### HasAllowUpdateSelf

`func (o *LicenseUsers) HasAllowUpdateSelf() bool`

HasAllowUpdateSelf returns a boolean if a field has been set.

### GetEntitledToSupport

`func (o *LicenseUsers) GetEntitledToSupport() bool`

GetEntitledToSupport returns the EntitledToSupport field if non-nil, zero value otherwise.

### GetEntitledToSupportOk

`func (o *LicenseUsers) GetEntitledToSupportOk() (*bool, bool)`

GetEntitledToSupportOk returns a tuple with the EntitledToSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitledToSupport

`func (o *LicenseUsers) SetEntitledToSupport(v bool)`

SetEntitledToSupport sets EntitledToSupport field to given value.

### HasEntitledToSupport

`func (o *LicenseUsers) HasEntitledToSupport() bool`

HasEntitledToSupport returns a boolean if a field has been set.

### GetMax

`func (o *LicenseUsers) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *LicenseUsers) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *LicenseUsers) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *LicenseUsers) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetHardLimitMax

`func (o *LicenseUsers) GetHardLimitMax() int32`

GetHardLimitMax returns the HardLimitMax field if non-nil, zero value otherwise.

### GetHardLimitMaxOk

`func (o *LicenseUsers) GetHardLimitMaxOk() (*int32, bool)`

GetHardLimitMaxOk returns a tuple with the HardLimitMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardLimitMax

`func (o *LicenseUsers) SetHardLimitMax(v int32)`

SetHardLimitMax sets HardLimitMax field to given value.

### HasHardLimitMax

`func (o *LicenseUsers) HasHardLimitMax() bool`

HasHardLimitMax returns a boolean if a field has been set.

### GetAnnualActiveIncluded

`func (o *LicenseUsers) GetAnnualActiveIncluded() int32`

GetAnnualActiveIncluded returns the AnnualActiveIncluded field if non-nil, zero value otherwise.

### GetAnnualActiveIncludedOk

`func (o *LicenseUsers) GetAnnualActiveIncludedOk() (*int32, bool)`

GetAnnualActiveIncludedOk returns a tuple with the AnnualActiveIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnualActiveIncluded

`func (o *LicenseUsers) SetAnnualActiveIncluded(v int32)`

SetAnnualActiveIncluded sets AnnualActiveIncluded field to given value.

### HasAnnualActiveIncluded

`func (o *LicenseUsers) HasAnnualActiveIncluded() bool`

HasAnnualActiveIncluded returns a boolean if a field has been set.

### GetMonthlyActiveIncluded

`func (o *LicenseUsers) GetMonthlyActiveIncluded() int32`

GetMonthlyActiveIncluded returns the MonthlyActiveIncluded field if non-nil, zero value otherwise.

### GetMonthlyActiveIncludedOk

`func (o *LicenseUsers) GetMonthlyActiveIncludedOk() (*int32, bool)`

GetMonthlyActiveIncludedOk returns a tuple with the MonthlyActiveIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyActiveIncluded

`func (o *LicenseUsers) SetMonthlyActiveIncluded(v int32)`

SetMonthlyActiveIncluded sets MonthlyActiveIncluded field to given value.

### HasMonthlyActiveIncluded

`func (o *LicenseUsers) HasMonthlyActiveIncluded() bool`

HasMonthlyActiveIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


