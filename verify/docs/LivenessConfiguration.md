# LivenessConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retry** | Pointer to [**ObjectRetry**](ObjectRetry.md) |  | [optional] 
**Threshold** | [**EnumThreshold**](EnumThreshold.md) |  | 
**Verify** | [**EnumVerify**](EnumVerify.md) |  | 

## Methods

### NewLivenessConfiguration

`func NewLivenessConfiguration(threshold EnumThreshold, verify EnumVerify, ) *LivenessConfiguration`

NewLivenessConfiguration instantiates a new LivenessConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLivenessConfigurationWithDefaults

`func NewLivenessConfigurationWithDefaults() *LivenessConfiguration`

NewLivenessConfigurationWithDefaults instantiates a new LivenessConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetry

`func (o *LivenessConfiguration) GetRetry() ObjectRetry`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *LivenessConfiguration) GetRetryOk() (*ObjectRetry, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *LivenessConfiguration) SetRetry(v ObjectRetry)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *LivenessConfiguration) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


