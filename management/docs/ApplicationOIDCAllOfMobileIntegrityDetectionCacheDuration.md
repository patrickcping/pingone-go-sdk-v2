# ApplicationOIDCAllOfMobileIntegrityDetectionCacheDuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | An integer that specifies the number of minutes or hours that specify the duration between successful integrity detection calls. Every attestation request entails a certain time tradeoff. You can choose to cache successful integrity detection calls for a predefined duration, between a minimum of 1 minute and a maximum of 48 hours. If mobile.integrityDetection.mode is ENABLED, the cache duration must be set. | [optional] 
**Units** | Pointer to [**EnumDurationUnitMinsHours**](EnumDurationUnitMinsHours.md) |  | [optional] 

## Methods

### NewApplicationOIDCAllOfMobileIntegrityDetectionCacheDuration

`func NewApplicationOIDCAllOfMobileIntegrityDetectionCacheDuration() *ApplicationOIDCAllOfMobileIntegrityDetectionCacheDuration`

NewApplicationOIDCAllOfMobileIntegrityDetectionCacheDuration instantiates a new ApplicationOIDCAllOfMobileIntegrityDetectionCacheDuration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationOIDCAllOfMobileIntegrityDetectionCacheDurationWithDefaults

`func NewApplicationOIDCAllOfMobileIntegrityDetectionCacheDurationWithDefaults() *ApplicationOIDCAllOfMobileIntegrityDetectionCacheDuration`

NewApplicationOIDCAllOfMobileIntegrityDetectionCacheDurationWithDefaults instantiates a new ApplicationOIDCAllOfMobileIntegrityDetectionCacheDuration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ApplicationOIDCAllOfMobileIntegrityDetectionCacheDuration) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ApplicationOIDCAllOfMobileIntegrityDetectionCacheDuration) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ApplicationOIDCAllOfMobileIntegrityDetectionCacheDuration) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ApplicationOIDCAllOfMobileIntegrityDetectionCacheDuration) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetUnits

`func (o *ApplicationOIDCAllOfMobileIntegrityDetectionCacheDuration) GetUnits() EnumDurationUnitMinsHours`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *ApplicationOIDCAllOfMobileIntegrityDetectionCacheDuration) GetUnitsOk() (*EnumDurationUnitMinsHours, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *ApplicationOIDCAllOfMobileIntegrityDetectionCacheDuration) SetUnits(v EnumDurationUnitMinsHours)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *ApplicationOIDCAllOfMobileIntegrityDetectionCacheDuration) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


