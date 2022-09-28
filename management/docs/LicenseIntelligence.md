# LicenseIntelligence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowGeoVelocity** | Pointer to **bool** | A read-only boolean that specifies whether to use the intelligence geo-velocity feature. For &#x60;TRIAL&#x60; (unpaid) licenses, the default value is true. For &#x60;ADMIN&#x60;, &#x60;GLOBAL&#x60;, &#x60;RISK&#x60;, and &#x60;MFARISK&#x60;, the default value is true. | [optional] 
**AllowAnonymousNetworkDetection** | Pointer to **bool** | A read-only boolean that specifies whether to use the intelligence anonymous network detection feature. For &#x60;TRIAL&#x60; (unpaid) licenses, the default value is true. For &#x60;ADMIN&#x60;, &#x60;GLOBAL&#x60;, &#x60;RISK&#x60;, and &#x60;MFARISK&#x60;, the default value is true. | [optional] 
**AllowReputation** | Pointer to **bool** | A read-only boolean that specifies whether to use the intelligence IP reputation feature. For &#x60;TRIAL&#x60; (unpaid) licenses, the default value is true. For &#x60;ADMIN&#x60;, &#x60;GLOBAL&#x60;, &#x60;RISK&#x60;, and &#x60;MFARISK&#x60;, the default value is true. | [optional] 
**AllowDataConsent** | Pointer to **bool** | A read-only boolean that specifies whether the customer has opted in to allow user and event behavior analytics (UEBA) data collection. | [optional] 
**AllowRisk** | Pointer to **bool** | A read-only boolean that specifies whether your license permits you to configure risk features such as sign-on policies that include rules to detect anomalous changes to your locations (such as impossible travel). This capability is supported for TRIAL, RISK, and MFARISK license packages. Note, The sharing of user data to enable our machine-learning engine, which is integral to PingOne Risk, is captured in the license property license.intelligence.allowDataConsent, but it is not set to true by default in any license package. This license capability always requires active consent by the customer before it can be enabled, and if consent is given, then it allows the full scope of intelligence features included in PingOne Risk (and PingOne Risk plus MFA). | [optional] 

## Methods

### NewLicenseIntelligence

`func NewLicenseIntelligence() *LicenseIntelligence`

NewLicenseIntelligence instantiates a new LicenseIntelligence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseIntelligenceWithDefaults

`func NewLicenseIntelligenceWithDefaults() *LicenseIntelligence`

NewLicenseIntelligenceWithDefaults instantiates a new LicenseIntelligence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowGeoVelocity

`func (o *LicenseIntelligence) GetAllowGeoVelocity() bool`

GetAllowGeoVelocity returns the AllowGeoVelocity field if non-nil, zero value otherwise.

### GetAllowGeoVelocityOk

`func (o *LicenseIntelligence) GetAllowGeoVelocityOk() (*bool, bool)`

GetAllowGeoVelocityOk returns a tuple with the AllowGeoVelocity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGeoVelocity

`func (o *LicenseIntelligence) SetAllowGeoVelocity(v bool)`

SetAllowGeoVelocity sets AllowGeoVelocity field to given value.

### HasAllowGeoVelocity

`func (o *LicenseIntelligence) HasAllowGeoVelocity() bool`

HasAllowGeoVelocity returns a boolean if a field has been set.

### GetAllowAnonymousNetworkDetection

`func (o *LicenseIntelligence) GetAllowAnonymousNetworkDetection() bool`

GetAllowAnonymousNetworkDetection returns the AllowAnonymousNetworkDetection field if non-nil, zero value otherwise.

### GetAllowAnonymousNetworkDetectionOk

`func (o *LicenseIntelligence) GetAllowAnonymousNetworkDetectionOk() (*bool, bool)`

GetAllowAnonymousNetworkDetectionOk returns a tuple with the AllowAnonymousNetworkDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAnonymousNetworkDetection

`func (o *LicenseIntelligence) SetAllowAnonymousNetworkDetection(v bool)`

SetAllowAnonymousNetworkDetection sets AllowAnonymousNetworkDetection field to given value.

### HasAllowAnonymousNetworkDetection

`func (o *LicenseIntelligence) HasAllowAnonymousNetworkDetection() bool`

HasAllowAnonymousNetworkDetection returns a boolean if a field has been set.

### GetAllowReputation

`func (o *LicenseIntelligence) GetAllowReputation() bool`

GetAllowReputation returns the AllowReputation field if non-nil, zero value otherwise.

### GetAllowReputationOk

`func (o *LicenseIntelligence) GetAllowReputationOk() (*bool, bool)`

GetAllowReputationOk returns a tuple with the AllowReputation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowReputation

`func (o *LicenseIntelligence) SetAllowReputation(v bool)`

SetAllowReputation sets AllowReputation field to given value.

### HasAllowReputation

`func (o *LicenseIntelligence) HasAllowReputation() bool`

HasAllowReputation returns a boolean if a field has been set.

### GetAllowDataConsent

`func (o *LicenseIntelligence) GetAllowDataConsent() bool`

GetAllowDataConsent returns the AllowDataConsent field if non-nil, zero value otherwise.

### GetAllowDataConsentOk

`func (o *LicenseIntelligence) GetAllowDataConsentOk() (*bool, bool)`

GetAllowDataConsentOk returns a tuple with the AllowDataConsent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDataConsent

`func (o *LicenseIntelligence) SetAllowDataConsent(v bool)`

SetAllowDataConsent sets AllowDataConsent field to given value.

### HasAllowDataConsent

`func (o *LicenseIntelligence) HasAllowDataConsent() bool`

HasAllowDataConsent returns a boolean if a field has been set.

### GetAllowRisk

`func (o *LicenseIntelligence) GetAllowRisk() bool`

GetAllowRisk returns the AllowRisk field if non-nil, zero value otherwise.

### GetAllowRiskOk

`func (o *LicenseIntelligence) GetAllowRiskOk() (*bool, bool)`

GetAllowRiskOk returns a tuple with the AllowRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRisk

`func (o *LicenseIntelligence) SetAllowRisk(v bool)`

SetAllowRisk sets AllowRisk field to given value.

### HasAllowRisk

`func (o *LicenseIntelligence) HasAllowRisk() bool`

HasAllowRisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


