# AdministratorSecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**AllowedMethods** | Pointer to [**EnumAdministratorSecurityAllowedMethods**](EnumAdministratorSecurityAllowedMethods.md) |  | [optional] 
**AuthenticationMethod** | [**EnumAdministratorSecurityAuthenticationMethod**](EnumAdministratorSecurityAuthenticationMethod.md) |  | [default to ENUMADMINISTRATORSECURITYAUTHENTICATIONMETHOD_PINGONE]
**CreatedAt** | Pointer to **time.Time** | The time the resource was created. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**MfaStatus** | Pointer to [**EnumAdministratorSecurityMfaStatus**](EnumAdministratorSecurityMfaStatus.md) |  | [optional] [default to ENUMADMINISTRATORSECURITYMFASTATUS_OUT]
**Policy** | Pointer to [**AdministratorSecurityPolicy**](AdministratorSecurityPolicy.md) |  | [optional] 
**Provider** | Pointer to [**AdministratorSecurityProvider**](AdministratorSecurityProvider.md) |  | [optional] 
**Recovery** | **bool** | Indicates whether to allow account recovery within the admin policy. | 
**UpdatedAt** | Pointer to **time.Time** | The time the resource was last updated. | [optional] [readonly] 

## Methods

### NewAdministratorSecurity

`func NewAdministratorSecurity(authenticationMethod EnumAdministratorSecurityAuthenticationMethod, recovery bool, ) *AdministratorSecurity`

NewAdministratorSecurity instantiates a new AdministratorSecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdministratorSecurityWithDefaults

`func NewAdministratorSecurityWithDefaults() *AdministratorSecurity`

NewAdministratorSecurityWithDefaults instantiates a new AdministratorSecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AdministratorSecurity) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AdministratorSecurity) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AdministratorSecurity) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AdministratorSecurity) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAllowedMethods

`func (o *AdministratorSecurity) GetAllowedMethods() EnumAdministratorSecurityAllowedMethods`

GetAllowedMethods returns the AllowedMethods field if non-nil, zero value otherwise.

### GetAllowedMethodsOk

`func (o *AdministratorSecurity) GetAllowedMethodsOk() (*EnumAdministratorSecurityAllowedMethods, bool)`

GetAllowedMethodsOk returns a tuple with the AllowedMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMethods

`func (o *AdministratorSecurity) SetAllowedMethods(v EnumAdministratorSecurityAllowedMethods)`

SetAllowedMethods sets AllowedMethods field to given value.

### HasAllowedMethods

`func (o *AdministratorSecurity) HasAllowedMethods() bool`

HasAllowedMethods returns a boolean if a field has been set.

### GetAuthenticationMethod

`func (o *AdministratorSecurity) GetAuthenticationMethod() EnumAdministratorSecurityAuthenticationMethod`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *AdministratorSecurity) GetAuthenticationMethodOk() (*EnumAdministratorSecurityAuthenticationMethod, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *AdministratorSecurity) SetAuthenticationMethod(v EnumAdministratorSecurityAuthenticationMethod)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.


### GetCreatedAt

`func (o *AdministratorSecurity) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdministratorSecurity) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdministratorSecurity) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AdministratorSecurity) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEnvironment

`func (o *AdministratorSecurity) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AdministratorSecurity) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AdministratorSecurity) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *AdministratorSecurity) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetMfaStatus

`func (o *AdministratorSecurity) GetMfaStatus() EnumAdministratorSecurityMfaStatus`

GetMfaStatus returns the MfaStatus field if non-nil, zero value otherwise.

### GetMfaStatusOk

`func (o *AdministratorSecurity) GetMfaStatusOk() (*EnumAdministratorSecurityMfaStatus, bool)`

GetMfaStatusOk returns a tuple with the MfaStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaStatus

`func (o *AdministratorSecurity) SetMfaStatus(v EnumAdministratorSecurityMfaStatus)`

SetMfaStatus sets MfaStatus field to given value.

### HasMfaStatus

`func (o *AdministratorSecurity) HasMfaStatus() bool`

HasMfaStatus returns a boolean if a field has been set.

### GetPolicy

`func (o *AdministratorSecurity) GetPolicy() AdministratorSecurityPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *AdministratorSecurity) GetPolicyOk() (*AdministratorSecurityPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *AdministratorSecurity) SetPolicy(v AdministratorSecurityPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *AdministratorSecurity) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetProvider

`func (o *AdministratorSecurity) GetProvider() AdministratorSecurityProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AdministratorSecurity) GetProviderOk() (*AdministratorSecurityProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AdministratorSecurity) SetProvider(v AdministratorSecurityProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *AdministratorSecurity) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRecovery

`func (o *AdministratorSecurity) GetRecovery() bool`

GetRecovery returns the Recovery field if non-nil, zero value otherwise.

### GetRecoveryOk

`func (o *AdministratorSecurity) GetRecoveryOk() (*bool, bool)`

GetRecoveryOk returns a tuple with the Recovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecovery

`func (o *AdministratorSecurity) SetRecovery(v bool)`

SetRecovery sets Recovery field to given value.


### GetUpdatedAt

`func (o *AdministratorSecurity) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdministratorSecurity) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdministratorSecurity) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdministratorSecurity) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


