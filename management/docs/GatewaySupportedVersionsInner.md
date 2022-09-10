# GatewaySupportedVersionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | A string that specifies the gateway version number. | [optional] 
**Image** | Pointer to **string** | A string that identifies the gateway image path. | [optional] 
**Recommended** | Pointer to **bool** | A boolean that specifies whether this is the recommended LDAP gateway version. | [optional] 
**Latest** | Pointer to **bool** | A boolean that specifies whether this is the latest LDAP gateway version. | [optional] 
**SupportEndsOn** | Pointer to **time.Time** |  | [optional] 
**DaysUntilSupportEnds** | Pointer to **int32** |  | [optional] 

## Methods

### NewGatewaySupportedVersionsInner

`func NewGatewaySupportedVersionsInner() *GatewaySupportedVersionsInner`

NewGatewaySupportedVersionsInner instantiates a new GatewaySupportedVersionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewaySupportedVersionsInnerWithDefaults

`func NewGatewaySupportedVersionsInnerWithDefaults() *GatewaySupportedVersionsInner`

NewGatewaySupportedVersionsInnerWithDefaults instantiates a new GatewaySupportedVersionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *GatewaySupportedVersionsInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GatewaySupportedVersionsInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GatewaySupportedVersionsInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GatewaySupportedVersionsInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetImage

`func (o *GatewaySupportedVersionsInner) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *GatewaySupportedVersionsInner) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *GatewaySupportedVersionsInner) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *GatewaySupportedVersionsInner) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetRecommended

`func (o *GatewaySupportedVersionsInner) GetRecommended() bool`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *GatewaySupportedVersionsInner) GetRecommendedOk() (*bool, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *GatewaySupportedVersionsInner) SetRecommended(v bool)`

SetRecommended sets Recommended field to given value.

### HasRecommended

`func (o *GatewaySupportedVersionsInner) HasRecommended() bool`

HasRecommended returns a boolean if a field has been set.

### GetLatest

`func (o *GatewaySupportedVersionsInner) GetLatest() bool`

GetLatest returns the Latest field if non-nil, zero value otherwise.

### GetLatestOk

`func (o *GatewaySupportedVersionsInner) GetLatestOk() (*bool, bool)`

GetLatestOk returns a tuple with the Latest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatest

`func (o *GatewaySupportedVersionsInner) SetLatest(v bool)`

SetLatest sets Latest field to given value.

### HasLatest

`func (o *GatewaySupportedVersionsInner) HasLatest() bool`

HasLatest returns a boolean if a field has been set.

### GetSupportEndsOn

`func (o *GatewaySupportedVersionsInner) GetSupportEndsOn() time.Time`

GetSupportEndsOn returns the SupportEndsOn field if non-nil, zero value otherwise.

### GetSupportEndsOnOk

`func (o *GatewaySupportedVersionsInner) GetSupportEndsOnOk() (*time.Time, bool)`

GetSupportEndsOnOk returns a tuple with the SupportEndsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEndsOn

`func (o *GatewaySupportedVersionsInner) SetSupportEndsOn(v time.Time)`

SetSupportEndsOn sets SupportEndsOn field to given value.

### HasSupportEndsOn

`func (o *GatewaySupportedVersionsInner) HasSupportEndsOn() bool`

HasSupportEndsOn returns a boolean if a field has been set.

### GetDaysUntilSupportEnds

`func (o *GatewaySupportedVersionsInner) GetDaysUntilSupportEnds() int32`

GetDaysUntilSupportEnds returns the DaysUntilSupportEnds field if non-nil, zero value otherwise.

### GetDaysUntilSupportEndsOk

`func (o *GatewaySupportedVersionsInner) GetDaysUntilSupportEndsOk() (*int32, bool)`

GetDaysUntilSupportEndsOk returns a tuple with the DaysUntilSupportEnds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysUntilSupportEnds

`func (o *GatewaySupportedVersionsInner) SetDaysUntilSupportEnds(v int32)`

SetDaysUntilSupportEnds sets DaysUntilSupportEnds field to given value.

### HasDaysUntilSupportEnds

`func (o *GatewaySupportedVersionsInner) HasDaysUntilSupportEnds() bool`

HasDaysUntilSupportEnds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


