# FormFieldCommonVisibility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumFormFieldVisibility**](EnumFormFieldVisibility.md) |  | 
**Key** | Pointer to **string** | A non-unique string associated with the field when visibility is evaluated by DaVinci at runtime. If the &#x60;type&#x60; property is set to &#x60;SHOW_BY_DEFAULT&#x60; or &#x60;HIDE_BY_DEFAULT&#x60;, then this property is required. | [optional] 

## Methods

### NewFormFieldCommonVisibility

`func NewFormFieldCommonVisibility(type_ EnumFormFieldVisibility, ) *FormFieldCommonVisibility`

NewFormFieldCommonVisibility instantiates a new FormFieldCommonVisibility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldCommonVisibilityWithDefaults

`func NewFormFieldCommonVisibilityWithDefaults() *FormFieldCommonVisibility`

NewFormFieldCommonVisibilityWithDefaults instantiates a new FormFieldCommonVisibility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FormFieldCommonVisibility) GetType() EnumFormFieldVisibility`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormFieldCommonVisibility) GetTypeOk() (*EnumFormFieldVisibility, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormFieldCommonVisibility) SetType(v EnumFormFieldVisibility)`

SetType sets Type field to given value.


### GetKey

`func (o *FormFieldCommonVisibility) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FormFieldCommonVisibility) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FormFieldCommonVisibility) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FormFieldCommonVisibility) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


