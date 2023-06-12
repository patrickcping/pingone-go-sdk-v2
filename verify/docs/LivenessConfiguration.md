# LivenessConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verify** | [**EnumVerify**](EnumVerify.md) |  | 
**Threshold** | [**EnumThreshold**](EnumThreshold.md) |  | 

## Methods

### NewLivenessConfiguration

`func NewLivenessConfiguration(verify EnumVerify, threshold EnumThreshold, ) *LivenessConfiguration`

NewLivenessConfiguration instantiates a new LivenessConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLivenessConfigurationWithDefaults

`func NewLivenessConfigurationWithDefaults() *LivenessConfiguration`

NewLivenessConfigurationWithDefaults instantiates a new LivenessConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerify

`func (o *LivenessConfiguration) GetVerify() EnumVerify`

GetVerify returns the Verify field if non-nil, zero value otherwise.

### GetVerifyOk

`func (o *LivenessConfiguration) GetVerifyOk() (*EnumVerify, bool)`

GetVerifyOk returns a tuple with the Verify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerify

`func (o *LivenessConfiguration) SetVerify(v EnumVerify)`

SetVerify sets Verify field to given value.


### GetThreshold

`func (o *LivenessConfiguration) GetThreshold() EnumThreshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *LivenessConfiguration) GetThresholdOk() (*EnumThreshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *LivenessConfiguration) SetThreshold(v EnumThreshold)`

SetThreshold sets Threshold field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


