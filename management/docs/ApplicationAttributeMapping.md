# ApplicationAttributeMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A string that specifies the application ID. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | The time the resource was created. | [optional] [readonly] 
**MappingType** | Pointer to **string** | A string that specifies the mapping type of the attribute. Options are CORE, SCOPE, and CUSTOM. The CORE and SCOPE mapping types are for reserved attributes managed by the API and cannot be removed. Attribute values for these mapping types can be updated. The CUSTOM mapping type is for user-defined attributes. Attributes of this type can be updated and deleted. | [optional] 
**Name** | **string** | A string that specifies the name of attribute and must be unique within an application. For SAML applications, the samlAssertion.subject name is a reserved case-insensitive name which indicates the mapping to be used for the subject in an assertion. For OpenID Connect applications, the following names are reserved and cannot be used acr, amr, at_hash, aud, auth_time, azp, client_id, exp, iat, iss, jti, nbf, nonce, org, scope, sid, sub  This is a required property. | 
**Required** | **bool** | A boolean to specify whether a mapping value is required for this attribute. If true, a value must be set and a non-empty value must be available in the SAML assertion or ID token. | 
**UpdatedAt** | Pointer to **string** | The time the resource was updated. | [optional] [readonly] 
**Value** | **string** | A string that specifies the string constants or expression for mapping the attribute path against a specific source. The expression format is ${&lt;source&gt;.&lt;attribute_path&gt;}. The only supported source is user (for example, ${user.id}). This is a required property. | 

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

### GetCreatedAt

`func (o *ApplicationAttributeMapping) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApplicationAttributeMapping) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApplicationAttributeMapping) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApplicationAttributeMapping) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMappingType

`func (o *ApplicationAttributeMapping) GetMappingType() string`

GetMappingType returns the MappingType field if non-nil, zero value otherwise.

### GetMappingTypeOk

`func (o *ApplicationAttributeMapping) GetMappingTypeOk() (*string, bool)`

GetMappingTypeOk returns a tuple with the MappingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingType

`func (o *ApplicationAttributeMapping) SetMappingType(v string)`

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

`func (o *ApplicationAttributeMapping) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ApplicationAttributeMapping) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ApplicationAttributeMapping) SetUpdatedAt(v string)`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


