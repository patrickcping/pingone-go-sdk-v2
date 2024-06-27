# FIDO2Policy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Id** | Pointer to **string** | FIDO policy&#39;s UUID. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The time the resource was last updated. | [optional] [readonly] 
**AttestationRequirements** | [**EnumFIDO2PolicyAttestationRequirements**](EnumFIDO2PolicyAttestationRequirements.md) |  | 
**AuthenticatorAttachment** | [**EnumFIDO2PolicyAuthenticatorAttachment**](EnumFIDO2PolicyAuthenticatorAttachment.md) |  | 
**BackupEligibility** | [**FIDO2PolicyBackupEligibility**](FIDO2PolicyBackupEligibility.md) |  | 
**Default** | Pointer to **bool** | Whether this policy should serve as the default FIDO policy. | [optional] 
**Description** | Pointer to **string** | Description of the FIDO policy. | [optional] 
**DeviceAuthenticationPolicies** | Pointer to **[]string** | The device authentication policies that use the relevant FIDO policy. If you include the parameter &#x60;expand&#x3D;deviceAuthenticationPolicies&#x60; in the URL of the request, this array is included in the response when reading FIDO policies. Each object in the array contains the ID and the name of the device authentication policy. | [optional] [readonly] 
**DeviceDisplayName** | **string** | The name to display for the device in registration and authentication windows. Can be up to 100 characters. If you want to use translatable text, you can use any of the keys listed on the *FIDO Policy* page of the *Self-Service* module and the *Sign On Policy* module. The value of the parameter should include only the part of the key name that comes after the module name, for example, &#x60;fidoPolicy.deviceDisplayName01&#x60; or &#x60;fidoPolicy.deviceDisplayName07&#x60;. See the pages in the UI for the full list of keys. For more information on translatable keys, see [Modifying translatable keys](https://docs.pingidentity.com/access/sources/dita/topic?category&#x3D;p1&amp;resourceid&#x3D;pingone_modifying_translatable_keys) in the PingOne documentation. | 
**DiscoverableCredentials** | [**EnumFIDO2PolicyDiscoverableCredentials**](EnumFIDO2PolicyDiscoverableCredentials.md) |  | 
**MdsAuthenticatorsRequirements** | [**FIDO2PolicyMdsAuthenticatorsRequirements**](FIDO2PolicyMdsAuthenticatorsRequirements.md) |  | 
**Name** | **string** | The name to use for the FIDO policy. Can be up to 256 characters. | 
**RelyingPartyId** | **string** | The ID of the relying party. The value should be a domain name, such as &#x60;example.com&#x60; (in lower-case characters). | 
**UserDisplayNameAttributes** | [**FIDO2PolicyUserDisplayNameAttributes**](FIDO2PolicyUserDisplayNameAttributes.md) |  | 
**UserVerification** | [**FIDO2PolicyUserVerification**](FIDO2PolicyUserVerification.md) |  | 

## Methods

### NewFIDO2Policy

`func NewFIDO2Policy(attestationRequirements EnumFIDO2PolicyAttestationRequirements, authenticatorAttachment EnumFIDO2PolicyAuthenticatorAttachment, backupEligibility FIDO2PolicyBackupEligibility, deviceDisplayName string, discoverableCredentials EnumFIDO2PolicyDiscoverableCredentials, mdsAuthenticatorsRequirements FIDO2PolicyMdsAuthenticatorsRequirements, name string, relyingPartyId string, userDisplayNameAttributes FIDO2PolicyUserDisplayNameAttributes, userVerification FIDO2PolicyUserVerification, ) *FIDO2Policy`

NewFIDO2Policy instantiates a new FIDO2Policy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFIDO2PolicyWithDefaults

`func NewFIDO2PolicyWithDefaults() *FIDO2Policy`

NewFIDO2PolicyWithDefaults instantiates a new FIDO2Policy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *FIDO2Policy) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FIDO2Policy) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FIDO2Policy) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FIDO2Policy) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *FIDO2Policy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FIDO2Policy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FIDO2Policy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FIDO2Policy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *FIDO2Policy) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *FIDO2Policy) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *FIDO2Policy) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *FIDO2Policy) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FIDO2Policy) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FIDO2Policy) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FIDO2Policy) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FIDO2Policy) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FIDO2Policy) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FIDO2Policy) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FIDO2Policy) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FIDO2Policy) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAttestationRequirements

`func (o *FIDO2Policy) GetAttestationRequirements() EnumFIDO2PolicyAttestationRequirements`

GetAttestationRequirements returns the AttestationRequirements field if non-nil, zero value otherwise.

### GetAttestationRequirementsOk

`func (o *FIDO2Policy) GetAttestationRequirementsOk() (*EnumFIDO2PolicyAttestationRequirements, bool)`

GetAttestationRequirementsOk returns a tuple with the AttestationRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestationRequirements

`func (o *FIDO2Policy) SetAttestationRequirements(v EnumFIDO2PolicyAttestationRequirements)`

SetAttestationRequirements sets AttestationRequirements field to given value.


### GetAuthenticatorAttachment

`func (o *FIDO2Policy) GetAuthenticatorAttachment() EnumFIDO2PolicyAuthenticatorAttachment`

GetAuthenticatorAttachment returns the AuthenticatorAttachment field if non-nil, zero value otherwise.

### GetAuthenticatorAttachmentOk

`func (o *FIDO2Policy) GetAuthenticatorAttachmentOk() (*EnumFIDO2PolicyAuthenticatorAttachment, bool)`

GetAuthenticatorAttachmentOk returns a tuple with the AuthenticatorAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorAttachment

`func (o *FIDO2Policy) SetAuthenticatorAttachment(v EnumFIDO2PolicyAuthenticatorAttachment)`

SetAuthenticatorAttachment sets AuthenticatorAttachment field to given value.


### GetBackupEligibility

`func (o *FIDO2Policy) GetBackupEligibility() FIDO2PolicyBackupEligibility`

GetBackupEligibility returns the BackupEligibility field if non-nil, zero value otherwise.

### GetBackupEligibilityOk

`func (o *FIDO2Policy) GetBackupEligibilityOk() (*FIDO2PolicyBackupEligibility, bool)`

GetBackupEligibilityOk returns a tuple with the BackupEligibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupEligibility

`func (o *FIDO2Policy) SetBackupEligibility(v FIDO2PolicyBackupEligibility)`

SetBackupEligibility sets BackupEligibility field to given value.


### GetDefault

`func (o *FIDO2Policy) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *FIDO2Policy) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *FIDO2Policy) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *FIDO2Policy) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDescription

`func (o *FIDO2Policy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FIDO2Policy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FIDO2Policy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FIDO2Policy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceAuthenticationPolicies

`func (o *FIDO2Policy) GetDeviceAuthenticationPolicies() []string`

GetDeviceAuthenticationPolicies returns the DeviceAuthenticationPolicies field if non-nil, zero value otherwise.

### GetDeviceAuthenticationPoliciesOk

`func (o *FIDO2Policy) GetDeviceAuthenticationPoliciesOk() (*[]string, bool)`

GetDeviceAuthenticationPoliciesOk returns a tuple with the DeviceAuthenticationPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAuthenticationPolicies

`func (o *FIDO2Policy) SetDeviceAuthenticationPolicies(v []string)`

SetDeviceAuthenticationPolicies sets DeviceAuthenticationPolicies field to given value.

### HasDeviceAuthenticationPolicies

`func (o *FIDO2Policy) HasDeviceAuthenticationPolicies() bool`

HasDeviceAuthenticationPolicies returns a boolean if a field has been set.

### GetDeviceDisplayName

`func (o *FIDO2Policy) GetDeviceDisplayName() string`

GetDeviceDisplayName returns the DeviceDisplayName field if non-nil, zero value otherwise.

### GetDeviceDisplayNameOk

`func (o *FIDO2Policy) GetDeviceDisplayNameOk() (*string, bool)`

GetDeviceDisplayNameOk returns a tuple with the DeviceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDisplayName

`func (o *FIDO2Policy) SetDeviceDisplayName(v string)`

SetDeviceDisplayName sets DeviceDisplayName field to given value.


### GetDiscoverableCredentials

`func (o *FIDO2Policy) GetDiscoverableCredentials() EnumFIDO2PolicyDiscoverableCredentials`

GetDiscoverableCredentials returns the DiscoverableCredentials field if non-nil, zero value otherwise.

### GetDiscoverableCredentialsOk

`func (o *FIDO2Policy) GetDiscoverableCredentialsOk() (*EnumFIDO2PolicyDiscoverableCredentials, bool)`

GetDiscoverableCredentialsOk returns a tuple with the DiscoverableCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverableCredentials

`func (o *FIDO2Policy) SetDiscoverableCredentials(v EnumFIDO2PolicyDiscoverableCredentials)`

SetDiscoverableCredentials sets DiscoverableCredentials field to given value.


### GetMdsAuthenticatorsRequirements

`func (o *FIDO2Policy) GetMdsAuthenticatorsRequirements() FIDO2PolicyMdsAuthenticatorsRequirements`

GetMdsAuthenticatorsRequirements returns the MdsAuthenticatorsRequirements field if non-nil, zero value otherwise.

### GetMdsAuthenticatorsRequirementsOk

`func (o *FIDO2Policy) GetMdsAuthenticatorsRequirementsOk() (*FIDO2PolicyMdsAuthenticatorsRequirements, bool)`

GetMdsAuthenticatorsRequirementsOk returns a tuple with the MdsAuthenticatorsRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdsAuthenticatorsRequirements

`func (o *FIDO2Policy) SetMdsAuthenticatorsRequirements(v FIDO2PolicyMdsAuthenticatorsRequirements)`

SetMdsAuthenticatorsRequirements sets MdsAuthenticatorsRequirements field to given value.


### GetName

`func (o *FIDO2Policy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FIDO2Policy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FIDO2Policy) SetName(v string)`

SetName sets Name field to given value.


### GetRelyingPartyId

`func (o *FIDO2Policy) GetRelyingPartyId() string`

GetRelyingPartyId returns the RelyingPartyId field if non-nil, zero value otherwise.

### GetRelyingPartyIdOk

`func (o *FIDO2Policy) GetRelyingPartyIdOk() (*string, bool)`

GetRelyingPartyIdOk returns a tuple with the RelyingPartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelyingPartyId

`func (o *FIDO2Policy) SetRelyingPartyId(v string)`

SetRelyingPartyId sets RelyingPartyId field to given value.


### GetUserDisplayNameAttributes

`func (o *FIDO2Policy) GetUserDisplayNameAttributes() FIDO2PolicyUserDisplayNameAttributes`

GetUserDisplayNameAttributes returns the UserDisplayNameAttributes field if non-nil, zero value otherwise.

### GetUserDisplayNameAttributesOk

`func (o *FIDO2Policy) GetUserDisplayNameAttributesOk() (*FIDO2PolicyUserDisplayNameAttributes, bool)`

GetUserDisplayNameAttributesOk returns a tuple with the UserDisplayNameAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDisplayNameAttributes

`func (o *FIDO2Policy) SetUserDisplayNameAttributes(v FIDO2PolicyUserDisplayNameAttributes)`

SetUserDisplayNameAttributes sets UserDisplayNameAttributes field to given value.


### GetUserVerification

`func (o *FIDO2Policy) GetUserVerification() FIDO2PolicyUserVerification`

GetUserVerification returns the UserVerification field if non-nil, zero value otherwise.

### GetUserVerificationOk

`func (o *FIDO2Policy) GetUserVerificationOk() (*FIDO2PolicyUserVerification, bool)`

GetUserVerificationOk returns a tuple with the UserVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVerification

`func (o *FIDO2Policy) SetUserVerification(v FIDO2PolicyUserVerification)`

SetUserVerification sets UserVerification field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


