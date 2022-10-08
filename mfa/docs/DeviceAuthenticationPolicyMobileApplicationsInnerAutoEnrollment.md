# DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Set to true if you want the application to allow Auto Enrollment. Auto Enrollment means that the user can authenticate for the first time from an unpaired device, and the successful authentication will result in the pairing of the device for MFA. | 

## Methods

### NewDeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment

`func NewDeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment(enabled bool, ) *DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment`

NewDeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment instantiates a new DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollmentWithDefaults

`func NewDeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollmentWithDefaults() *DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment`

NewDeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollmentWithDefaults instantiates a new DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


