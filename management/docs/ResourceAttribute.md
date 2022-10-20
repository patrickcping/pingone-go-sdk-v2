# ResourceAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
**Name** | **string** | A string that specifies the name of the custom resource attribute to be included in the access token. The following are reserved names and cannot be used. Thesese reserved names are applicable only when the resource&#39;s type property is &#x60;OPENID_CONNECT&#x60;: - &#x60;acr&#x60; - &#x60;amr&#x60; - &#x60;aud&#x60; - &#x60;auth_time&#x60; - &#x60;client_id&#x60; - &#x60;env&#x60; - &#x60;exp&#x60; - &#x60;iat&#x60; - &#x60;iss&#x60; - &#x60;jti&#x60; - &#x60;org&#x60; - &#x60;p1.*&#x60; (any name starting with the p1. prefix) - &#x60;scope&#x60; - &#x60;sid&#x60; - &#x60;sub&#x60;  | 
**Type** | Pointer to [**EnumResourceAttributeType**](EnumResourceAttributeType.md) |  | [optional] 
**Value** | **string** | A string that specifies the value of the custom resource attribute. This value can be a placeholder that references an attribute in the user schema, expressed as &#x60;${user.path.to.value}&#x60;, or it can be a static string. Placeholders must be valid, enabled attributes in the environment’s user schema. Examples fo valid values are &#x60;${user.email}&#x60;, &#x60;${user.name.family}&#x60;, and &#x60;myClaimValueString&#x60; | 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Resource** | Pointer to [**IdentityProviderAttributeIdentityProvider**](IdentityProviderAttributeIdentityProvider.md) |  | [optional] 
**IdToken** | Pointer to **bool** | A boolean that specifies whether the attribute mapping should be available in the ID Token. This property is applicable only when the application&#39;s protocol property is &#x60;OPENID_CONNECT&#x60;. If omitted, the default is &#x60;true&#x60;. Note that the &#x60;idToken&#x60; and &#x60;userInfo&#x60; properties cannot both be set to &#x60;false&#x60;. At least one of these properties must have a value of &#x60;true&#x60;. | [optional] 
**UserInfo** | Pointer to **bool** | A boolean that specifies whether the attribute mapping should be available through the &#x60;/as/userinfo&#x60; endpoint. This property is applicable only when the application&#39;s protocol property is &#x60;OPENID_CONNECT&#x60;. If omitted, the default is &#x60;true&#x60;. Note that the &#x60;idToken&#x60; and &#x60;userInfo&#x60; properties cannot both be set to &#x60;false&#x60;. At least one of these properties must have a value of &#x60;true&#x60;. | [optional] 

## Methods

### NewResourceAttribute

`func NewResourceAttribute(name string, value string, ) *ResourceAttribute`

NewResourceAttribute instantiates a new ResourceAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceAttributeWithDefaults

`func NewResourceAttributeWithDefaults() *ResourceAttribute`

NewResourceAttributeWithDefaults instantiates a new ResourceAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceAttribute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceAttribute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceAttribute) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceAttribute) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ResourceAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceAttribute) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ResourceAttribute) GetType() EnumResourceAttributeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceAttribute) GetTypeOk() (*EnumResourceAttributeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceAttribute) SetType(v EnumResourceAttributeType)`

SetType sets Type field to given value.

### HasType

`func (o *ResourceAttribute) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *ResourceAttribute) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ResourceAttribute) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ResourceAttribute) SetValue(v string)`

SetValue sets Value field to given value.


### GetEnvironment

`func (o *ResourceAttribute) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ResourceAttribute) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ResourceAttribute) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ResourceAttribute) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetResource

`func (o *ResourceAttribute) GetResource() IdentityProviderAttributeIdentityProvider`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ResourceAttribute) GetResourceOk() (*IdentityProviderAttributeIdentityProvider, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ResourceAttribute) SetResource(v IdentityProviderAttributeIdentityProvider)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ResourceAttribute) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetIdToken

`func (o *ResourceAttribute) GetIdToken() bool`

GetIdToken returns the IdToken field if non-nil, zero value otherwise.

### GetIdTokenOk

`func (o *ResourceAttribute) GetIdTokenOk() (*bool, bool)`

GetIdTokenOk returns a tuple with the IdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdToken

`func (o *ResourceAttribute) SetIdToken(v bool)`

SetIdToken sets IdToken field to given value.

### HasIdToken

`func (o *ResourceAttribute) HasIdToken() bool`

HasIdToken returns a boolean if a field has been set.

### GetUserInfo

`func (o *ResourceAttribute) GetUserInfo() bool`

GetUserInfo returns the UserInfo field if non-nil, zero value otherwise.

### GetUserInfoOk

`func (o *ResourceAttribute) GetUserInfoOk() (*bool, bool)`

GetUserInfoOk returns a tuple with the UserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfo

`func (o *ResourceAttribute) SetUserInfo(v bool)`

SetUserInfo sets UserInfo field to given value.

### HasUserInfo

`func (o *ResourceAttribute) HasUserInfo() bool`

HasUserInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


