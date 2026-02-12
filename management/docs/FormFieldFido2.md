# FormFieldFIDO2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumFormFieldType**](EnumFormFieldType.md) |  | 
**Position** | [**FormFieldCommonPosition**](FormFieldCommonPosition.md) |  | 
**Visibility** | Pointer to [**FormFieldCommonVisibility**](FormFieldCommonVisibility.md) |  | [optional] 
**Trigger** | [**EnumFormFIDO2Trigger**](EnumFormFIDO2Trigger.md) |  | 
**Action** | [**EnumFormFIDO2Action**](EnumFormFIDO2Action.md) |  | 
**Label** | **string** | A string that specifies the text label for the FIDO2 button. | 

## Methods

### NewFormFieldFIDO2

`func NewFormFieldFIDO2(type_ EnumFormFieldType, position FormFieldCommonPosition, trigger EnumFormFIDO2Trigger, action EnumFormFIDO2Action, label string, ) *FormFieldFIDO2`

NewFormFieldFIDO2 instantiates a new FormFieldFIDO2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldFIDO2WithDefaults

`func NewFormFieldFIDO2WithDefaults() *FormFieldFIDO2`

NewFormFieldFIDO2WithDefaults instantiates a new FormFieldFIDO2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FormFieldFIDO2) GetType() EnumFormFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormFieldFIDO2) GetTypeOk() (*EnumFormFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormFieldFIDO2) SetType(v EnumFormFieldType)`

SetType sets Type field to given value.


### GetPosition

`func (o *FormFieldFIDO2) GetPosition() FormFieldCommonPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FormFieldFIDO2) GetPositionOk() (*FormFieldCommonPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FormFieldFIDO2) SetPosition(v FormFieldCommonPosition)`

SetPosition sets Position field to given value.


### GetVisibility

`func (o *FormFieldFIDO2) GetVisibility() FormFieldCommonVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *FormFieldFIDO2) GetVisibilityOk() (*FormFieldCommonVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *FormFieldFIDO2) SetVisibility(v FormFieldCommonVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *FormFieldFIDO2) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTrigger

`func (o *FormFieldFIDO2) GetTrigger() EnumFormFIDO2Trigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *FormFieldFIDO2) GetTriggerOk() (*EnumFormFIDO2Trigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *FormFieldFIDO2) SetTrigger(v EnumFormFIDO2Trigger)`

SetTrigger sets Trigger field to given value.


### GetAction

`func (o *FormFieldFIDO2) GetAction() EnumFormFIDO2Action`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FormFieldFIDO2) GetActionOk() (*EnumFormFIDO2Action, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FormFieldFIDO2) SetAction(v EnumFormFIDO2Action)`

SetAction sets Action field to given value.


### GetLabel

`func (o *FormFieldFIDO2) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FormFieldFIDO2) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FormFieldFIDO2) SetLabel(v string)`

SetLabel sets Label field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


