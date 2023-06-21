# FIDO2PolicyUserVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnforceDuringAuthentication** | **bool** | Set to &#x60;true&#x60; if you want the device characteristics related to user verification to be checked again at each authentication attempt and not just once during registration. Set to &#x60;false&#x60; to have them checked only at registration. | 
**Option** | [**EnumFIDO2PolicyUserVerificationOption**](EnumFIDO2PolicyUserVerificationOption.md) |  | 

## Methods

### NewFIDO2PolicyUserVerification

`func NewFIDO2PolicyUserVerification(enforceDuringAuthentication bool, option EnumFIDO2PolicyUserVerificationOption, ) *FIDO2PolicyUserVerification`

NewFIDO2PolicyUserVerification instantiates a new FIDO2PolicyUserVerification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFIDO2PolicyUserVerificationWithDefaults

`func NewFIDO2PolicyUserVerificationWithDefaults() *FIDO2PolicyUserVerification`

NewFIDO2PolicyUserVerificationWithDefaults instantiates a new FIDO2PolicyUserVerification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnforceDuringAuthentication

`func (o *FIDO2PolicyUserVerification) GetEnforceDuringAuthentication() bool`

GetEnforceDuringAuthentication returns the EnforceDuringAuthentication field if non-nil, zero value otherwise.

### GetEnforceDuringAuthenticationOk

`func (o *FIDO2PolicyUserVerification) GetEnforceDuringAuthenticationOk() (*bool, bool)`

GetEnforceDuringAuthenticationOk returns a tuple with the EnforceDuringAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceDuringAuthentication

`func (o *FIDO2PolicyUserVerification) SetEnforceDuringAuthentication(v bool)`

SetEnforceDuringAuthentication sets EnforceDuringAuthentication field to given value.


### GetOption

`func (o *FIDO2PolicyUserVerification) GetOption() EnumFIDO2PolicyUserVerificationOption`

GetOption returns the Option field if non-nil, zero value otherwise.

### GetOptionOk

`func (o *FIDO2PolicyUserVerification) GetOptionOk() (*EnumFIDO2PolicyUserVerificationOption, bool)`

GetOptionOk returns a tuple with the Option field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption

`func (o *FIDO2PolicyUserVerification) SetOption(v EnumFIDO2PolicyUserVerificationOption)`

SetOption sets Option field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


