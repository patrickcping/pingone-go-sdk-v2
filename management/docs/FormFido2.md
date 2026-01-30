# FormFIDO2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Trigger** | [**EnumFormFIDO2Trigger**](EnumFormFIDO2Trigger.md) |  | 
**Action** | [**EnumFormFIDO2Action**](EnumFormFIDO2Action.md) |  | 
**Label** | **string** | A string that specifies the text label for the FIDO2 button. | 

## Methods

### NewFormFIDO2

`func NewFormFIDO2(trigger EnumFormFIDO2Trigger, action EnumFormFIDO2Action, label string, ) *FormFIDO2`

NewFormFIDO2 instantiates a new FormFIDO2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFIDO2WithDefaults

`func NewFormFIDO2WithDefaults() *FormFIDO2`

NewFormFIDO2WithDefaults instantiates a new FormFIDO2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrigger

`func (o *FormFIDO2) GetTrigger() EnumFormFIDO2Trigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *FormFIDO2) GetTriggerOk() (*EnumFormFIDO2Trigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *FormFIDO2) SetTrigger(v EnumFormFIDO2Trigger)`

SetTrigger sets Trigger field to given value.


### GetAction

`func (o *FormFIDO2) GetAction() EnumFormFIDO2Action`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FormFIDO2) GetActionOk() (*EnumFormFIDO2Action, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FormFIDO2) SetAction(v EnumFormFIDO2Action)`

SetAction sets Action field to given value.


### GetLabel

`func (o *FormFIDO2) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FormFIDO2) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FormFIDO2) SetLabel(v string)`

SetLabel sets Label field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


