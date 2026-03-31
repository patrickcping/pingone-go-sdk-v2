# FormAgreement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | A string that specifies an identifier for the field component. | 
**InputType** | [**EnumFormAgreementInputType**](EnumFormAgreementInputType.md) |  | 
**TitleEnabled** | **bool** | Specifies whether the title is enabled. | 
**Agreement** | Pointer to [**FormAgreementAllOfAgreement**](FormAgreementAllOfAgreement.md) |  | [optional] 

## Methods

### NewFormAgreement

`func NewFormAgreement(key string, inputType EnumFormAgreementInputType, titleEnabled bool, ) *FormAgreement`

NewFormAgreement instantiates a new FormAgreement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormAgreementWithDefaults

`func NewFormAgreementWithDefaults() *FormAgreement`

NewFormAgreementWithDefaults instantiates a new FormAgreement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *FormAgreement) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FormAgreement) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FormAgreement) SetKey(v string)`

SetKey sets Key field to given value.


### GetInputType

`func (o *FormAgreement) GetInputType() EnumFormAgreementInputType`

GetInputType returns the InputType field if non-nil, zero value otherwise.

### GetInputTypeOk

`func (o *FormAgreement) GetInputTypeOk() (*EnumFormAgreementInputType, bool)`

GetInputTypeOk returns a tuple with the InputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputType

`func (o *FormAgreement) SetInputType(v EnumFormAgreementInputType)`

SetInputType sets InputType field to given value.


### GetTitleEnabled

`func (o *FormAgreement) GetTitleEnabled() bool`

GetTitleEnabled returns the TitleEnabled field if non-nil, zero value otherwise.

### GetTitleEnabledOk

`func (o *FormAgreement) GetTitleEnabledOk() (*bool, bool)`

GetTitleEnabledOk returns a tuple with the TitleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleEnabled

`func (o *FormAgreement) SetTitleEnabled(v bool)`

SetTitleEnabled sets TitleEnabled field to given value.


### GetAgreement

`func (o *FormAgreement) GetAgreement() FormAgreementAllOfAgreement`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *FormAgreement) GetAgreementOk() (*FormAgreementAllOfAgreement, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *FormAgreement) SetAgreement(v FormAgreementAllOfAgreement)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *FormAgreement) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


