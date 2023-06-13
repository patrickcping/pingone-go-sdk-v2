# OTPDeviceConfigurationOtpDeliveries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | Count of OTP deliveries. | 
**Cooldown** | [**OTPDeviceConfigurationOtpDeliveriesCooldown**](OTPDeviceConfigurationOtpDeliveriesCooldown.md) |  | 

## Methods

### NewOTPDeviceConfigurationOtpDeliveries

`func NewOTPDeviceConfigurationOtpDeliveries(count int32, cooldown OTPDeviceConfigurationOtpDeliveriesCooldown, ) *OTPDeviceConfigurationOtpDeliveries`

NewOTPDeviceConfigurationOtpDeliveries instantiates a new OTPDeviceConfigurationOtpDeliveries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOTPDeviceConfigurationOtpDeliveriesWithDefaults

`func NewOTPDeviceConfigurationOtpDeliveriesWithDefaults() *OTPDeviceConfigurationOtpDeliveries`

NewOTPDeviceConfigurationOtpDeliveriesWithDefaults instantiates a new OTPDeviceConfigurationOtpDeliveries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *OTPDeviceConfigurationOtpDeliveries) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *OTPDeviceConfigurationOtpDeliveries) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *OTPDeviceConfigurationOtpDeliveries) SetCount(v int32)`

SetCount sets Count field to given value.


### GetCooldown

`func (o *OTPDeviceConfigurationOtpDeliveries) GetCooldown() OTPDeviceConfigurationOtpDeliveriesCooldown`

GetCooldown returns the Cooldown field if non-nil, zero value otherwise.

### GetCooldownOk

`func (o *OTPDeviceConfigurationOtpDeliveries) GetCooldownOk() (*OTPDeviceConfigurationOtpDeliveriesCooldown, bool)`

GetCooldownOk returns a tuple with the Cooldown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldown

`func (o *OTPDeviceConfigurationOtpDeliveries) SetCooldown(v OTPDeviceConfigurationOtpDeliveriesCooldown)`

SetCooldown sets Cooldown field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


