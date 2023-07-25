# CredentialTypeMetaDataFieldsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attribute** | Pointer to **string** | A string that specifies the name of the PingOne Directory attribute if field.type is Directory Attribute. | [optional] 
**Id** | **string** | A string that specifies the identifier of the field. | 
**Title** | **string** | A string that specifies the descriptive text when showing the field. | 
**IsVisible** | **bool** | A boolean value that specifies whether the field should be visible to viewers of the credential. | 
**Type** | [**EnumCredentialTypeMetaDataFieldsType**](EnumCredentialTypeMetaDataFieldsType.md) |  | 
**Value** | Pointer to **string** | A string that specifies the text to appear on the credential for a type of Alphanumeric Text. | [optional] 

## Methods

### NewCredentialTypeMetaDataFieldsInner

`func NewCredentialTypeMetaDataFieldsInner(id string, title string, isVisible bool, type_ EnumCredentialTypeMetaDataFieldsType, ) *CredentialTypeMetaDataFieldsInner`

NewCredentialTypeMetaDataFieldsInner instantiates a new CredentialTypeMetaDataFieldsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialTypeMetaDataFieldsInnerWithDefaults

`func NewCredentialTypeMetaDataFieldsInnerWithDefaults() *CredentialTypeMetaDataFieldsInner`

NewCredentialTypeMetaDataFieldsInnerWithDefaults instantiates a new CredentialTypeMetaDataFieldsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttribute

`func (o *CredentialTypeMetaDataFieldsInner) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *CredentialTypeMetaDataFieldsInner) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *CredentialTypeMetaDataFieldsInner) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *CredentialTypeMetaDataFieldsInner) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetId

`func (o *CredentialTypeMetaDataFieldsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CredentialTypeMetaDataFieldsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CredentialTypeMetaDataFieldsInner) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *CredentialTypeMetaDataFieldsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CredentialTypeMetaDataFieldsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CredentialTypeMetaDataFieldsInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetIsVisible

`func (o *CredentialTypeMetaDataFieldsInner) GetIsVisible() bool`

GetIsVisible returns the IsVisible field if non-nil, zero value otherwise.

### GetIsVisibleOk

`func (o *CredentialTypeMetaDataFieldsInner) GetIsVisibleOk() (*bool, bool)`

GetIsVisibleOk returns a tuple with the IsVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVisible

`func (o *CredentialTypeMetaDataFieldsInner) SetIsVisible(v bool)`

SetIsVisible sets IsVisible field to given value.


### GetType

`func (o *CredentialTypeMetaDataFieldsInner) GetType() EnumCredentialTypeMetaDataFieldsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CredentialTypeMetaDataFieldsInner) GetTypeOk() (*EnumCredentialTypeMetaDataFieldsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CredentialTypeMetaDataFieldsInner) SetType(v EnumCredentialTypeMetaDataFieldsType)`

SetType sets Type field to given value.


### GetValue

`func (o *CredentialTypeMetaDataFieldsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CredentialTypeMetaDataFieldsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CredentialTypeMetaDataFieldsInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CredentialTypeMetaDataFieldsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


