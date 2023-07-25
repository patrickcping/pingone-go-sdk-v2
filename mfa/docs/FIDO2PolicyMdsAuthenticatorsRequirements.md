# FIDO2PolicyMdsAuthenticatorsRequirements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedAuthenticators** | Pointer to [**[]FIDO2PolicyMdsAuthenticatorsRequirementsAllowedAuthenticatorsInner**](FIDO2PolicyMdsAuthenticatorsRequirementsAllowedAuthenticatorsInner.md) | If you set &#x60;mdsAuthenticatorsRequirements.option&#x60; to &#x60;SPECIFIC&#x60;, use this array to specify the authenticators that you want to allow. | [optional] 
**EnforceDuringAuthentication** | **bool** | Set to true if you want the device characteristics related to attestation to be checked again at each authentication attempt and not just once during registration. Set to false to have them checked only at registration. | 
**Option** | [**EnumFIDO2PolicyMDSAuthenticatorOption**](EnumFIDO2PolicyMDSAuthenticatorOption.md) |  | 

## Methods

### NewFIDO2PolicyMdsAuthenticatorsRequirements

`func NewFIDO2PolicyMdsAuthenticatorsRequirements(enforceDuringAuthentication bool, option EnumFIDO2PolicyMDSAuthenticatorOption, ) *FIDO2PolicyMdsAuthenticatorsRequirements`

NewFIDO2PolicyMdsAuthenticatorsRequirements instantiates a new FIDO2PolicyMdsAuthenticatorsRequirements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFIDO2PolicyMdsAuthenticatorsRequirementsWithDefaults

`func NewFIDO2PolicyMdsAuthenticatorsRequirementsWithDefaults() *FIDO2PolicyMdsAuthenticatorsRequirements`

NewFIDO2PolicyMdsAuthenticatorsRequirementsWithDefaults instantiates a new FIDO2PolicyMdsAuthenticatorsRequirements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedAuthenticators

`func (o *FIDO2PolicyMdsAuthenticatorsRequirements) GetAllowedAuthenticators() []FIDO2PolicyMdsAuthenticatorsRequirementsAllowedAuthenticatorsInner`

GetAllowedAuthenticators returns the AllowedAuthenticators field if non-nil, zero value otherwise.

### GetAllowedAuthenticatorsOk

`func (o *FIDO2PolicyMdsAuthenticatorsRequirements) GetAllowedAuthenticatorsOk() (*[]FIDO2PolicyMdsAuthenticatorsRequirementsAllowedAuthenticatorsInner, bool)`

GetAllowedAuthenticatorsOk returns a tuple with the AllowedAuthenticators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAuthenticators

`func (o *FIDO2PolicyMdsAuthenticatorsRequirements) SetAllowedAuthenticators(v []FIDO2PolicyMdsAuthenticatorsRequirementsAllowedAuthenticatorsInner)`

SetAllowedAuthenticators sets AllowedAuthenticators field to given value.

### HasAllowedAuthenticators

`func (o *FIDO2PolicyMdsAuthenticatorsRequirements) HasAllowedAuthenticators() bool`

HasAllowedAuthenticators returns a boolean if a field has been set.

### GetEnforceDuringAuthentication

`func (o *FIDO2PolicyMdsAuthenticatorsRequirements) GetEnforceDuringAuthentication() bool`

GetEnforceDuringAuthentication returns the EnforceDuringAuthentication field if non-nil, zero value otherwise.

### GetEnforceDuringAuthenticationOk

`func (o *FIDO2PolicyMdsAuthenticatorsRequirements) GetEnforceDuringAuthenticationOk() (*bool, bool)`

GetEnforceDuringAuthenticationOk returns a tuple with the EnforceDuringAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceDuringAuthentication

`func (o *FIDO2PolicyMdsAuthenticatorsRequirements) SetEnforceDuringAuthentication(v bool)`

SetEnforceDuringAuthentication sets EnforceDuringAuthentication field to given value.


### GetOption

`func (o *FIDO2PolicyMdsAuthenticatorsRequirements) GetOption() EnumFIDO2PolicyMDSAuthenticatorOption`

GetOption returns the Option field if non-nil, zero value otherwise.

### GetOptionOk

`func (o *FIDO2PolicyMdsAuthenticatorsRequirements) GetOptionOk() (*EnumFIDO2PolicyMDSAuthenticatorOption, bool)`

GetOptionOk returns a tuple with the Option field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption

`func (o *FIDO2PolicyMdsAuthenticatorsRequirements) SetOption(v EnumFIDO2PolicyMDSAuthenticatorOption)`

SetOption sets Option field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


