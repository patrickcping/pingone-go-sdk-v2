# VoiceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verify** | [**EnumVerify**](EnumVerify.md) |  | 
**Enrollment** | **bool** |  | 
**TextDependent** | Pointer to [**VoiceConfigurationTextDependent**](VoiceConfigurationTextDependent.md) |  | [optional] 
**Comparison** | [**VoiceConfigurationThreshold**](VoiceConfigurationThreshold.md) |  | 
**Liveness** | [**VoiceConfigurationThreshold**](VoiceConfigurationThreshold.md) |  | 
**ReferenceData** | Pointer to [**VoiceConfigurationReferenceData**](VoiceConfigurationReferenceData.md) |  | [optional] 

## Methods

### NewVoiceConfiguration

`func NewVoiceConfiguration(verify EnumVerify, enrollment bool, comparison VoiceConfigurationThreshold, liveness VoiceConfigurationThreshold, ) *VoiceConfiguration`

NewVoiceConfiguration instantiates a new VoiceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoiceConfigurationWithDefaults

`func NewVoiceConfigurationWithDefaults() *VoiceConfiguration`

NewVoiceConfigurationWithDefaults instantiates a new VoiceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerify

`func (o *VoiceConfiguration) GetVerify() EnumVerify`

GetVerify returns the Verify field if non-nil, zero value otherwise.

### GetVerifyOk

`func (o *VoiceConfiguration) GetVerifyOk() (*EnumVerify, bool)`

GetVerifyOk returns a tuple with the Verify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerify

`func (o *VoiceConfiguration) SetVerify(v EnumVerify)`

SetVerify sets Verify field to given value.


### GetEnrollment

`func (o *VoiceConfiguration) GetEnrollment() bool`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *VoiceConfiguration) GetEnrollmentOk() (*bool, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *VoiceConfiguration) SetEnrollment(v bool)`

SetEnrollment sets Enrollment field to given value.


### GetTextDependent

`func (o *VoiceConfiguration) GetTextDependent() VoiceConfigurationTextDependent`

GetTextDependent returns the TextDependent field if non-nil, zero value otherwise.

### GetTextDependentOk

`func (o *VoiceConfiguration) GetTextDependentOk() (*VoiceConfigurationTextDependent, bool)`

GetTextDependentOk returns a tuple with the TextDependent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextDependent

`func (o *VoiceConfiguration) SetTextDependent(v VoiceConfigurationTextDependent)`

SetTextDependent sets TextDependent field to given value.

### HasTextDependent

`func (o *VoiceConfiguration) HasTextDependent() bool`

HasTextDependent returns a boolean if a field has been set.

### GetComparison

`func (o *VoiceConfiguration) GetComparison() VoiceConfigurationThreshold`

GetComparison returns the Comparison field if non-nil, zero value otherwise.

### GetComparisonOk

`func (o *VoiceConfiguration) GetComparisonOk() (*VoiceConfigurationThreshold, bool)`

GetComparisonOk returns a tuple with the Comparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparison

`func (o *VoiceConfiguration) SetComparison(v VoiceConfigurationThreshold)`

SetComparison sets Comparison field to given value.


### GetLiveness

`func (o *VoiceConfiguration) GetLiveness() VoiceConfigurationThreshold`

GetLiveness returns the Liveness field if non-nil, zero value otherwise.

### GetLivenessOk

`func (o *VoiceConfiguration) GetLivenessOk() (*VoiceConfigurationThreshold, bool)`

GetLivenessOk returns a tuple with the Liveness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveness

`func (o *VoiceConfiguration) SetLiveness(v VoiceConfigurationThreshold)`

SetLiveness sets Liveness field to given value.


### GetReferenceData

`func (o *VoiceConfiguration) GetReferenceData() VoiceConfigurationReferenceData`

GetReferenceData returns the ReferenceData field if non-nil, zero value otherwise.

### GetReferenceDataOk

`func (o *VoiceConfiguration) GetReferenceDataOk() (*VoiceConfigurationReferenceData, bool)`

GetReferenceDataOk returns a tuple with the ReferenceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceData

`func (o *VoiceConfiguration) SetReferenceData(v VoiceConfigurationReferenceData)`

SetReferenceData sets ReferenceData field to given value.

### HasReferenceData

`func (o *VoiceConfiguration) HasReferenceData() bool`

HasReferenceData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


