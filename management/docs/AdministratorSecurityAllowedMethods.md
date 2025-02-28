# AdministratorSecurityAllowedMethods

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EMAIL** | **string** | Indicates whether to enable email for sign-on. | 
**FIDO2** | **string** | Indicates whether to enable FIDO2 for sign-on. | 
**TOTP** | **string** | Indicates whether to enable TOTP for sign-on. | 

## Methods

### NewAdministratorSecurityAllowedMethods

`func NewAdministratorSecurityAllowedMethods(eMAIL string, fIDO2 string, tOTP string, ) *AdministratorSecurityAllowedMethods`

NewAdministratorSecurityAllowedMethods instantiates a new AdministratorSecurityAllowedMethods object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdministratorSecurityAllowedMethodsWithDefaults

`func NewAdministratorSecurityAllowedMethodsWithDefaults() *AdministratorSecurityAllowedMethods`

NewAdministratorSecurityAllowedMethodsWithDefaults instantiates a new AdministratorSecurityAllowedMethods object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEMAIL

`func (o *AdministratorSecurityAllowedMethods) GetEMAIL() string`

GetEMAIL returns the EMAIL field if non-nil, zero value otherwise.

### GetEMAILOk

`func (o *AdministratorSecurityAllowedMethods) GetEMAILOk() (*string, bool)`

GetEMAILOk returns a tuple with the EMAIL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEMAIL

`func (o *AdministratorSecurityAllowedMethods) SetEMAIL(v string)`

SetEMAIL sets EMAIL field to given value.


### GetFIDO2

`func (o *AdministratorSecurityAllowedMethods) GetFIDO2() string`

GetFIDO2 returns the FIDO2 field if non-nil, zero value otherwise.

### GetFIDO2Ok

`func (o *AdministratorSecurityAllowedMethods) GetFIDO2Ok() (*string, bool)`

GetFIDO2Ok returns a tuple with the FIDO2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFIDO2

`func (o *AdministratorSecurityAllowedMethods) SetFIDO2(v string)`

SetFIDO2 sets FIDO2 field to given value.


### GetTOTP

`func (o *AdministratorSecurityAllowedMethods) GetTOTP() string`

GetTOTP returns the TOTP field if non-nil, zero value otherwise.

### GetTOTPOk

`func (o *AdministratorSecurityAllowedMethods) GetTOTPOk() (*string, bool)`

GetTOTPOk returns a tuple with the TOTP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTOTP

`func (o *AdministratorSecurityAllowedMethods) SetTOTP(v string)`

SetTOTP sets TOTP field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


