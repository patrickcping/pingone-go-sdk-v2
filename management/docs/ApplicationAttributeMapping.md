# ApplicationAttributeMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**LinksHATEOAS**](LinksHATEOAS.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the application ID. | [optional] [readonly] 
**Application** | Pointer to [**ApplicationAttributeMappingApplication**](ApplicationAttributeMappingApplication.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time the resource was created. | [optional] [readonly] 
**MappingType** | Pointer to [**EnumAttributeMappingType**](EnumAttributeMappingType.md) |  | [optional] 
**Name** | **string** | A string that specifies the name of attribute and must be unique within an application. For SAML applications, the samlAssertion.subject name is a reserved case-insensitive name which indicates the mapping to be used for the subject in an assertion. For OpenID Connect applications, the following names are reserved and cannot be used acr, amr, at_hash, aud, auth_time, azp, client_id, exp, iat, iss, jti, nbf, nonce, org, scope, sid, sub  This is a required property. | 
**Required** | **bool** | A boolean to specify whether a mapping value is required for this attribute. If true, a value must be set and a non-empty value must be available in the SAML assertion or ID token. | 
**UpdatedAt** | Pointer to **time.Time** | The time the resource was updated. | [optional] [readonly] 
**Value** | **string** | A string that specifies the string constants or expression for mapping the attribute path against a specific source. The expression format is ${&lt;source&gt;.&lt;attribute_path&gt;}. The only supported source is user (for example, ${user.id}). This is a required property. | 
**NameFormat** | Pointer to **string** | A URI reference representing the classification of the attribute. Helps the service provider interpret the attribute format. | [optional] 
**IdToken** | Pointer to **bool** | Whether the attribute mapping should be available in the ID Token. This property is applicable only when the application&#39;s &#x60;protocol&#x60; property is &#x60;OPENID_CONNECT&#x60;. If omitted, the default is &#x60;true&#x60;. Note that the &#x60;idToken&#x60; and &#x60;userInfo&#x60; properties cannot both be set to &#x60;false&#x60;. At least one of these properties must have a value of true. | [optional] [default to true]
**UserInfo** | Pointer to **bool** | Whether the attribute mapping should be available through the &#x60;/as/userinfo&#x60; endpoint. This property is applicable only when the application&#39;s protocol property is &#x60;OPENID_CONNECT&#x60;. If omitted, the default is &#x60;true&#x60;. Note that the &#x60;idToken&#x60; and &#x60;userInfo&#x60; properties cannot both be set to &#x60;false&#x60;. At least one of these properties must have a value of &#x60;true&#x60;. | [optional] [default to true]
**OidcScopes** | Pointer to **[]string** | OIDC resource scope IDs that this attribute mapping is available for exclusively. This setting overrides any global OIDC resource scopes that contain an attribute mapping with the same name. The list can contain only scope IDs that have been granted for the application through the &#x60;/grants&#x60; endpoint. A null value is accepted for backwards compatibility. However, an empty set is invalid, and one scope ID is expected. If null, the response includes this mapping in the &#x60;openid&#x60; scope. | [optional] 

## Methods

### NewApplicationAttributeMapping

`func NewApplicationAttributeMapping(name string, required bool, value string, ) *ApplicationAttributeMapping`

NewApplicationAttributeMapping instantiates a new ApplicationAttributeMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationAttributeMappingWithDefaults

`func NewApplicationAttributeMappingWithDefaults() *ApplicationAttributeMapping`

NewApplicationAttributeMappingWithDefaults instantiates a new ApplicationAttributeMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ApplicationAttributeMapping) GetLinks() LinksHATEOAS`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApplicationAttributeMapping) GetLinksOk() (*LinksHATEOAS, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApplicationAttributeMapping) SetLinks(v LinksHATEOAS)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApplicationAttributeMapping) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *ApplicationAttributeMapping) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationAttributeMapping) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationAttributeMapping) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationAttributeMapping) HasId() bool`

HasId returns a boolean if a field has been set.

### GetApplication

`func (o *ApplicationAttributeMapping) GetApplication() ApplicationAttributeMappingApplication`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *ApplicationAttributeMapping) GetApplicationOk() (*ApplicationAttributeMappingApplication, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *ApplicationAttributeMapping) SetApplication(v ApplicationAttributeMappingApplication)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *ApplicationAttributeMapping) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ApplicationAttributeMapping) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApplicationAttributeMapping) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApplicationAttributeMapping) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApplicationAttributeMapping) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMappingType

`func (o *ApplicationAttributeMapping) GetMappingType() EnumAttributeMappingType`

GetMappingType returns the MappingType field if non-nil, zero value otherwise.

### GetMappingTypeOk

`func (o *ApplicationAttributeMapping) GetMappingTypeOk() (*EnumAttributeMappingType, bool)`

GetMappingTypeOk returns a tuple with the MappingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingType

`func (o *ApplicationAttributeMapping) SetMappingType(v EnumAttributeMappingType)`

SetMappingType sets MappingType field to given value.

### HasMappingType

`func (o *ApplicationAttributeMapping) HasMappingType() bool`

HasMappingType returns a boolean if a field has been set.

### GetName

`func (o *ApplicationAttributeMapping) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationAttributeMapping) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationAttributeMapping) SetName(v string)`

SetName sets Name field to given value.


### GetRequired

`func (o *ApplicationAttributeMapping) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ApplicationAttributeMapping) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ApplicationAttributeMapping) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetUpdatedAt

`func (o *ApplicationAttributeMapping) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ApplicationAttributeMapping) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ApplicationAttributeMapping) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ApplicationAttributeMapping) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetValue

`func (o *ApplicationAttributeMapping) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ApplicationAttributeMapping) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ApplicationAttributeMapping) SetValue(v string)`

SetValue sets Value field to given value.


### GetNameFormat

`func (o *ApplicationAttributeMapping) GetNameFormat() string`

GetNameFormat returns the NameFormat field if non-nil, zero value otherwise.

### GetNameFormatOk

`func (o *ApplicationAttributeMapping) GetNameFormatOk() (*string, bool)`

GetNameFormatOk returns a tuple with the NameFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameFormat

`func (o *ApplicationAttributeMapping) SetNameFormat(v string)`

SetNameFormat sets NameFormat field to given value.

### HasNameFormat

`func (o *ApplicationAttributeMapping) HasNameFormat() bool`

HasNameFormat returns a boolean if a field has been set.

### GetIdToken

`func (o *ApplicationAttributeMapping) GetIdToken() bool`

GetIdToken returns the IdToken field if non-nil, zero value otherwise.

### GetIdTokenOk

`func (o *ApplicationAttributeMapping) GetIdTokenOk() (*bool, bool)`

GetIdTokenOk returns a tuple with the IdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdToken

`func (o *ApplicationAttributeMapping) SetIdToken(v bool)`

SetIdToken sets IdToken field to given value.

### HasIdToken

`func (o *ApplicationAttributeMapping) HasIdToken() bool`

HasIdToken returns a boolean if a field has been set.

### GetUserInfo

`func (o *ApplicationAttributeMapping) GetUserInfo() bool`

GetUserInfo returns the UserInfo field if non-nil, zero value otherwise.

### GetUserInfoOk

`func (o *ApplicationAttributeMapping) GetUserInfoOk() (*bool, bool)`

GetUserInfoOk returns a tuple with the UserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfo

`func (o *ApplicationAttributeMapping) SetUserInfo(v bool)`

SetUserInfo sets UserInfo field to given value.

### HasUserInfo

`func (o *ApplicationAttributeMapping) HasUserInfo() bool`

HasUserInfo returns a boolean if a field has been set.

### GetOidcScopes

`func (o *ApplicationAttributeMapping) GetOidcScopes() []string`

GetOidcScopes returns the OidcScopes field if non-nil, zero value otherwise.

### GetOidcScopesOk

`func (o *ApplicationAttributeMapping) GetOidcScopesOk() (*[]string, bool)`

GetOidcScopesOk returns a tuple with the OidcScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcScopes

`func (o *ApplicationAttributeMapping) SetOidcScopes(v []string)`

SetOidcScopes sets OidcScopes field to given value.

### HasOidcScopes

`func (o *ApplicationAttributeMapping) HasOidcScopes() bool`

HasOidcScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


