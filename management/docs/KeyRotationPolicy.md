# KeyRotationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**LinksHATEOAS**](LinksHATEOAS.md) |  | [optional] 
**Algorithm** | [**EnumKeyRotationPolicyAlgorithm**](EnumKeyRotationPolicyAlgorithm.md) |  | 
**CurrentKeyId** | Pointer to **string** | The &#x60;kid&#x60; (key identifier) of the &#x60;KrpKey&#x60; designated as &#x60;CURRENT&#x60;. | [optional] [readonly] 
**Dn** | **string** | The DN this KRP will apply to generated &#x60;KrpKeys&#x60;. Is applied as both &#x60;issuerDN&#x60; and &#x60;subjectDN&#x60; because generated &#x60;KrpKeys&#x60; are self-signed. | 
**Id** | Pointer to **string** | The resource’s unique identifier. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**KeyLength** | **int32** | The number of bytes of a cryptographic key this KRP will apply to generated &#x60;KrpKeys&#x60;. | 
**Name** | **string** | Human-readable name displayed in the admin console. | 
**NextKeyId** | Pointer to **string** | The &#x60;kid&#x60; (key identifier) of the &#x60;KrpKey&#x60; designated as &#x60;NEXT&#x60;. | [optional] [readonly] 
**RotatedAt** | Pointer to **time.Time** | The last time this KRP was rotated. | [optional] [readonly] 
**RotationPeriod** | Pointer to **int32** | The number of days to elapse before this KRP rotates &#x60;KrpKeys&#x60;. The default value is &#x60;90&#x60; days. The minimum value is &#x60;30&#x60; days. The maximum value is 1 day less than the &#x60;validityPeriod&#x60; value. | [optional] [default to 90]
**SignatureAlgorithm** | [**EnumKeyRotationPolicySigAlgorithm**](EnumKeyRotationPolicySigAlgorithm.md) |  | 
**UsageType** | [**EnumKeyRotationPolicyUsageType**](EnumKeyRotationPolicyUsageType.md) |  | 
**ValidityPeriod** | Pointer to **int32** | Controls the &#x60;startsAt&#x60; and &#x60;expiresAt&#x60; fields this KRP will apply to generated &#x60;KrpKeys&#x60;. The default value is &#x60;365&#x60; days. | [optional] [default to 365]

## Methods

### NewKeyRotationPolicy

`func NewKeyRotationPolicy(algorithm EnumKeyRotationPolicyAlgorithm, dn string, keyLength int32, name string, signatureAlgorithm EnumKeyRotationPolicySigAlgorithm, usageType EnumKeyRotationPolicyUsageType, ) *KeyRotationPolicy`

NewKeyRotationPolicy instantiates a new KeyRotationPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyRotationPolicyWithDefaults

`func NewKeyRotationPolicyWithDefaults() *KeyRotationPolicy`

NewKeyRotationPolicyWithDefaults instantiates a new KeyRotationPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *KeyRotationPolicy) GetLinks() LinksHATEOAS`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *KeyRotationPolicy) GetLinksOk() (*LinksHATEOAS, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *KeyRotationPolicy) SetLinks(v LinksHATEOAS)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *KeyRotationPolicy) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAlgorithm

`func (o *KeyRotationPolicy) GetAlgorithm() EnumKeyRotationPolicyAlgorithm`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *KeyRotationPolicy) GetAlgorithmOk() (*EnumKeyRotationPolicyAlgorithm, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *KeyRotationPolicy) SetAlgorithm(v EnumKeyRotationPolicyAlgorithm)`

SetAlgorithm sets Algorithm field to given value.


### GetCurrentKeyId

`func (o *KeyRotationPolicy) GetCurrentKeyId() string`

GetCurrentKeyId returns the CurrentKeyId field if non-nil, zero value otherwise.

### GetCurrentKeyIdOk

`func (o *KeyRotationPolicy) GetCurrentKeyIdOk() (*string, bool)`

GetCurrentKeyIdOk returns a tuple with the CurrentKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentKeyId

`func (o *KeyRotationPolicy) SetCurrentKeyId(v string)`

SetCurrentKeyId sets CurrentKeyId field to given value.

### HasCurrentKeyId

`func (o *KeyRotationPolicy) HasCurrentKeyId() bool`

HasCurrentKeyId returns a boolean if a field has been set.

### GetDn

`func (o *KeyRotationPolicy) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *KeyRotationPolicy) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *KeyRotationPolicy) SetDn(v string)`

SetDn sets Dn field to given value.


### GetId

`func (o *KeyRotationPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyRotationPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyRotationPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KeyRotationPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *KeyRotationPolicy) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *KeyRotationPolicy) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *KeyRotationPolicy) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *KeyRotationPolicy) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetKeyLength

`func (o *KeyRotationPolicy) GetKeyLength() int32`

GetKeyLength returns the KeyLength field if non-nil, zero value otherwise.

### GetKeyLengthOk

`func (o *KeyRotationPolicy) GetKeyLengthOk() (*int32, bool)`

GetKeyLengthOk returns a tuple with the KeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyLength

`func (o *KeyRotationPolicy) SetKeyLength(v int32)`

SetKeyLength sets KeyLength field to given value.


### GetName

`func (o *KeyRotationPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyRotationPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyRotationPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetNextKeyId

`func (o *KeyRotationPolicy) GetNextKeyId() string`

GetNextKeyId returns the NextKeyId field if non-nil, zero value otherwise.

### GetNextKeyIdOk

`func (o *KeyRotationPolicy) GetNextKeyIdOk() (*string, bool)`

GetNextKeyIdOk returns a tuple with the NextKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextKeyId

`func (o *KeyRotationPolicy) SetNextKeyId(v string)`

SetNextKeyId sets NextKeyId field to given value.

### HasNextKeyId

`func (o *KeyRotationPolicy) HasNextKeyId() bool`

HasNextKeyId returns a boolean if a field has been set.

### GetRotatedAt

`func (o *KeyRotationPolicy) GetRotatedAt() time.Time`

GetRotatedAt returns the RotatedAt field if non-nil, zero value otherwise.

### GetRotatedAtOk

`func (o *KeyRotationPolicy) GetRotatedAtOk() (*time.Time, bool)`

GetRotatedAtOk returns a tuple with the RotatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatedAt

`func (o *KeyRotationPolicy) SetRotatedAt(v time.Time)`

SetRotatedAt sets RotatedAt field to given value.

### HasRotatedAt

`func (o *KeyRotationPolicy) HasRotatedAt() bool`

HasRotatedAt returns a boolean if a field has been set.

### GetRotationPeriod

`func (o *KeyRotationPolicy) GetRotationPeriod() int32`

GetRotationPeriod returns the RotationPeriod field if non-nil, zero value otherwise.

### GetRotationPeriodOk

`func (o *KeyRotationPolicy) GetRotationPeriodOk() (*int32, bool)`

GetRotationPeriodOk returns a tuple with the RotationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationPeriod

`func (o *KeyRotationPolicy) SetRotationPeriod(v int32)`

SetRotationPeriod sets RotationPeriod field to given value.

### HasRotationPeriod

`func (o *KeyRotationPolicy) HasRotationPeriod() bool`

HasRotationPeriod returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *KeyRotationPolicy) GetSignatureAlgorithm() EnumKeyRotationPolicySigAlgorithm`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *KeyRotationPolicy) GetSignatureAlgorithmOk() (*EnumKeyRotationPolicySigAlgorithm, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *KeyRotationPolicy) SetSignatureAlgorithm(v EnumKeyRotationPolicySigAlgorithm)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.


### GetUsageType

`func (o *KeyRotationPolicy) GetUsageType() EnumKeyRotationPolicyUsageType`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *KeyRotationPolicy) GetUsageTypeOk() (*EnumKeyRotationPolicyUsageType, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *KeyRotationPolicy) SetUsageType(v EnumKeyRotationPolicyUsageType)`

SetUsageType sets UsageType field to given value.


### GetValidityPeriod

`func (o *KeyRotationPolicy) GetValidityPeriod() int32`

GetValidityPeriod returns the ValidityPeriod field if non-nil, zero value otherwise.

### GetValidityPeriodOk

`func (o *KeyRotationPolicy) GetValidityPeriodOk() (*int32, bool)`

GetValidityPeriodOk returns a tuple with the ValidityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityPeriod

`func (o *KeyRotationPolicy) SetValidityPeriod(v int32)`

SetValidityPeriod sets ValidityPeriod field to given value.

### HasValidityPeriod

`func (o *KeyRotationPolicy) HasValidityPeriod() bool`

HasValidityPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


