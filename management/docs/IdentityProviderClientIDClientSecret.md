# IdentityProviderClientIDClientSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | Pointer to **string** | The description of the IdP. | [optional] 
**Enabled** | **bool** | The current enabled state of the IdP. | 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Icon** | Pointer to [**IdentityProviderCommonIcon**](IdentityProviderCommonIcon.md) |  | [optional] 
**Id** | Pointer to **string** | The resource ID. | [optional] [readonly] 
**LoginButtonIcon** | Pointer to [**IdentityProviderCommonLoginButtonIcon**](IdentityProviderCommonLoginButtonIcon.md) |  | [optional] 
**Name** | **string** | The name of the IdP. | 
**Registration** | Pointer to [**IdentityProviderCommonRegistration**](IdentityProviderCommonRegistration.md) |  | [optional] 
**Type** | [**EnumIdentityProviderExt**](EnumIdentityProviderExt.md) |  | 
**CreatedAt** | Pointer to **string** | The time the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | The time the resource was last updated. | [optional] [readonly] 
**ClientId** | **string** | A string that specifies the application ID from the provider. This is a required property. | 
**ClientSecret** | **string** | A string that specifies the application secret from the provider. This is a required property. | 

## Methods

### NewIdentityProviderClientIDClientSecret

`func NewIdentityProviderClientIDClientSecret(enabled bool, name string, type_ EnumIdentityProviderExt, clientId string, clientSecret string, ) *IdentityProviderClientIDClientSecret`

NewIdentityProviderClientIDClientSecret instantiates a new IdentityProviderClientIDClientSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderClientIDClientSecretWithDefaults

`func NewIdentityProviderClientIDClientSecretWithDefaults() *IdentityProviderClientIDClientSecret`

NewIdentityProviderClientIDClientSecretWithDefaults instantiates a new IdentityProviderClientIDClientSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *IdentityProviderClientIDClientSecret) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IdentityProviderClientIDClientSecret) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IdentityProviderClientIDClientSecret) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *IdentityProviderClientIDClientSecret) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetDescription

`func (o *IdentityProviderClientIDClientSecret) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IdentityProviderClientIDClientSecret) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IdentityProviderClientIDClientSecret) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IdentityProviderClientIDClientSecret) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *IdentityProviderClientIDClientSecret) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IdentityProviderClientIDClientSecret) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IdentityProviderClientIDClientSecret) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEnvironment

`func (o *IdentityProviderClientIDClientSecret) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *IdentityProviderClientIDClientSecret) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *IdentityProviderClientIDClientSecret) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *IdentityProviderClientIDClientSecret) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetIcon

`func (o *IdentityProviderClientIDClientSecret) GetIcon() IdentityProviderCommonIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *IdentityProviderClientIDClientSecret) GetIconOk() (*IdentityProviderCommonIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *IdentityProviderClientIDClientSecret) SetIcon(v IdentityProviderCommonIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *IdentityProviderClientIDClientSecret) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetId

`func (o *IdentityProviderClientIDClientSecret) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityProviderClientIDClientSecret) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityProviderClientIDClientSecret) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityProviderClientIDClientSecret) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoginButtonIcon

`func (o *IdentityProviderClientIDClientSecret) GetLoginButtonIcon() IdentityProviderCommonLoginButtonIcon`

GetLoginButtonIcon returns the LoginButtonIcon field if non-nil, zero value otherwise.

### GetLoginButtonIconOk

`func (o *IdentityProviderClientIDClientSecret) GetLoginButtonIconOk() (*IdentityProviderCommonLoginButtonIcon, bool)`

GetLoginButtonIconOk returns a tuple with the LoginButtonIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginButtonIcon

`func (o *IdentityProviderClientIDClientSecret) SetLoginButtonIcon(v IdentityProviderCommonLoginButtonIcon)`

SetLoginButtonIcon sets LoginButtonIcon field to given value.

### HasLoginButtonIcon

`func (o *IdentityProviderClientIDClientSecret) HasLoginButtonIcon() bool`

HasLoginButtonIcon returns a boolean if a field has been set.

### GetName

`func (o *IdentityProviderClientIDClientSecret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityProviderClientIDClientSecret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityProviderClientIDClientSecret) SetName(v string)`

SetName sets Name field to given value.


### GetRegistration

`func (o *IdentityProviderClientIDClientSecret) GetRegistration() IdentityProviderCommonRegistration`

GetRegistration returns the Registration field if non-nil, zero value otherwise.

### GetRegistrationOk

`func (o *IdentityProviderClientIDClientSecret) GetRegistrationOk() (*IdentityProviderCommonRegistration, bool)`

GetRegistrationOk returns a tuple with the Registration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistration

`func (o *IdentityProviderClientIDClientSecret) SetRegistration(v IdentityProviderCommonRegistration)`

SetRegistration sets Registration field to given value.

### HasRegistration

`func (o *IdentityProviderClientIDClientSecret) HasRegistration() bool`

HasRegistration returns a boolean if a field has been set.

### GetType

`func (o *IdentityProviderClientIDClientSecret) GetType() EnumIdentityProviderExt`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdentityProviderClientIDClientSecret) GetTypeOk() (*EnumIdentityProviderExt, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdentityProviderClientIDClientSecret) SetType(v EnumIdentityProviderExt)`

SetType sets Type field to given value.


### GetCreatedAt

`func (o *IdentityProviderClientIDClientSecret) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IdentityProviderClientIDClientSecret) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IdentityProviderClientIDClientSecret) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IdentityProviderClientIDClientSecret) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IdentityProviderClientIDClientSecret) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IdentityProviderClientIDClientSecret) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IdentityProviderClientIDClientSecret) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IdentityProviderClientIDClientSecret) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetClientId

`func (o *IdentityProviderClientIDClientSecret) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IdentityProviderClientIDClientSecret) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IdentityProviderClientIDClientSecret) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *IdentityProviderClientIDClientSecret) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *IdentityProviderClientIDClientSecret) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *IdentityProviderClientIDClientSecret) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


