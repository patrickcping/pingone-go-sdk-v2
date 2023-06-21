# FIDO2PolicyBackupEligibility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allow** | **bool** | Set to &#x60;true&#x60; if you want to let users register and authenticate with a device that uses cloud-synced credentials. Set to &#x60;false&#x60; if you don&#39;t want to allow this. | 
**EnforceDuringAuthentication** | **bool** | Set to &#x60;true&#x60; if you want the backup eligibility of the device to be checked again at each authentication attempt and not just once during registration. Set to &#x60;false&#x60; to have it checked only at registration. | 

## Methods

### NewFIDO2PolicyBackupEligibility

`func NewFIDO2PolicyBackupEligibility(allow bool, enforceDuringAuthentication bool, ) *FIDO2PolicyBackupEligibility`

NewFIDO2PolicyBackupEligibility instantiates a new FIDO2PolicyBackupEligibility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFIDO2PolicyBackupEligibilityWithDefaults

`func NewFIDO2PolicyBackupEligibilityWithDefaults() *FIDO2PolicyBackupEligibility`

NewFIDO2PolicyBackupEligibilityWithDefaults instantiates a new FIDO2PolicyBackupEligibility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllow

`func (o *FIDO2PolicyBackupEligibility) GetAllow() bool`

GetAllow returns the Allow field if non-nil, zero value otherwise.

### GetAllowOk

`func (o *FIDO2PolicyBackupEligibility) GetAllowOk() (*bool, bool)`

GetAllowOk returns a tuple with the Allow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllow

`func (o *FIDO2PolicyBackupEligibility) SetAllow(v bool)`

SetAllow sets Allow field to given value.


### GetEnforceDuringAuthentication

`func (o *FIDO2PolicyBackupEligibility) GetEnforceDuringAuthentication() bool`

GetEnforceDuringAuthentication returns the EnforceDuringAuthentication field if non-nil, zero value otherwise.

### GetEnforceDuringAuthenticationOk

`func (o *FIDO2PolicyBackupEligibility) GetEnforceDuringAuthenticationOk() (*bool, bool)`

GetEnforceDuringAuthenticationOk returns a tuple with the EnforceDuringAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceDuringAuthentication

`func (o *FIDO2PolicyBackupEligibility) SetEnforceDuringAuthentication(v bool)`

SetEnforceDuringAuthentication sets EnforceDuringAuthentication field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


