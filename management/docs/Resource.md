# Resource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**AccessTokenValiditySeconds** | Pointer to **int32** | An integer that specifies the number of seconds that the access token is valid. If a value is not specified, the default is 3600. The minimum value is 300 seconds (5 minutes); the maximum value is 2592000 seconds (30 days). | [optional] 
**ApplicationPermissionsSettings** | Pointer to [**ResourceApplicationPermissionsSettings**](ResourceApplicationPermissionsSettings.md) |  | [optional] 
**Audience** | Pointer to **string** | A string that specifies a URL without a fragment or &#x60;@ObjectName&#x60; and must not contain &#x60;pingone&#x60; or &#x60;pingidentity&#x60; (for example, &#x60;https://api.bxretail.org&#x60;). If a URL is not specified, the resource name is used. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time the resource was created. | [optional] [readonly] 
**Description** | Pointer to **string** | A string that specifies the description of the resource. | [optional] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Resource** | Pointer to [**IdentityProviderAttributeIdentityProvider**](IdentityProviderAttributeIdentityProvider.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
**Name** | **string** | A string that specifies the resource name, which must be provided and must be unique within an environment. | 
**IntrospectEndpointAuthMethod** | Pointer to [**EnumResourceIntrospectEndpointAuthMethod**](EnumResourceIntrospectEndpointAuthMethod.md) |  | [optional] 
**Type** | Pointer to [**EnumResourceType**](EnumResourceType.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The time the resource was last updated. | [optional] [readonly] 

## Methods

### NewResource

`func NewResource(name string, ) *Resource`

NewResource instantiates a new Resource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceWithDefaults

`func NewResourceWithDefaults() *Resource`

NewResourceWithDefaults instantiates a new Resource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Resource) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Resource) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Resource) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Resource) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAccessTokenValiditySeconds

`func (o *Resource) GetAccessTokenValiditySeconds() int32`

GetAccessTokenValiditySeconds returns the AccessTokenValiditySeconds field if non-nil, zero value otherwise.

### GetAccessTokenValiditySecondsOk

`func (o *Resource) GetAccessTokenValiditySecondsOk() (*int32, bool)`

GetAccessTokenValiditySecondsOk returns a tuple with the AccessTokenValiditySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenValiditySeconds

`func (o *Resource) SetAccessTokenValiditySeconds(v int32)`

SetAccessTokenValiditySeconds sets AccessTokenValiditySeconds field to given value.

### HasAccessTokenValiditySeconds

`func (o *Resource) HasAccessTokenValiditySeconds() bool`

HasAccessTokenValiditySeconds returns a boolean if a field has been set.

### GetApplicationPermissionsSettings

`func (o *Resource) GetApplicationPermissionsSettings() ResourceApplicationPermissionsSettings`

GetApplicationPermissionsSettings returns the ApplicationPermissionsSettings field if non-nil, zero value otherwise.

### GetApplicationPermissionsSettingsOk

`func (o *Resource) GetApplicationPermissionsSettingsOk() (*ResourceApplicationPermissionsSettings, bool)`

GetApplicationPermissionsSettingsOk returns a tuple with the ApplicationPermissionsSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationPermissionsSettings

`func (o *Resource) SetApplicationPermissionsSettings(v ResourceApplicationPermissionsSettings)`

SetApplicationPermissionsSettings sets ApplicationPermissionsSettings field to given value.

### HasApplicationPermissionsSettings

`func (o *Resource) HasApplicationPermissionsSettings() bool`

HasApplicationPermissionsSettings returns a boolean if a field has been set.

### GetAudience

`func (o *Resource) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *Resource) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *Resource) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *Resource) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Resource) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Resource) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Resource) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Resource) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *Resource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Resource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Resource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Resource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *Resource) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Resource) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Resource) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *Resource) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetResource

`func (o *Resource) GetResource() IdentityProviderAttributeIdentityProvider`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *Resource) GetResourceOk() (*IdentityProviderAttributeIdentityProvider, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *Resource) SetResource(v IdentityProviderAttributeIdentityProvider)`

SetResource sets Resource field to given value.

### HasResource

`func (o *Resource) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetId

`func (o *Resource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Resource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Resource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Resource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Resource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Resource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Resource) SetName(v string)`

SetName sets Name field to given value.


### GetIntrospectEndpointAuthMethod

`func (o *Resource) GetIntrospectEndpointAuthMethod() EnumResourceIntrospectEndpointAuthMethod`

GetIntrospectEndpointAuthMethod returns the IntrospectEndpointAuthMethod field if non-nil, zero value otherwise.

### GetIntrospectEndpointAuthMethodOk

`func (o *Resource) GetIntrospectEndpointAuthMethodOk() (*EnumResourceIntrospectEndpointAuthMethod, bool)`

GetIntrospectEndpointAuthMethodOk returns a tuple with the IntrospectEndpointAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntrospectEndpointAuthMethod

`func (o *Resource) SetIntrospectEndpointAuthMethod(v EnumResourceIntrospectEndpointAuthMethod)`

SetIntrospectEndpointAuthMethod sets IntrospectEndpointAuthMethod field to given value.

### HasIntrospectEndpointAuthMethod

`func (o *Resource) HasIntrospectEndpointAuthMethod() bool`

HasIntrospectEndpointAuthMethod returns a boolean if a field has been set.

### GetType

`func (o *Resource) GetType() EnumResourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Resource) GetTypeOk() (*EnumResourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Resource) SetType(v EnumResourceType)`

SetType sets Type field to given value.

### HasType

`func (o *Resource) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Resource) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Resource) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Resource) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Resource) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


