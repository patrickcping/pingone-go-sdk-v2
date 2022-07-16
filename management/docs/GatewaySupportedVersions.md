# GatewaySupportedVersions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | A string that specifies the gateway version number. | [optional] 
**Image** | Pointer to **string** | A string that identifies the gateway image path. | [optional] 
**Recommended** | Pointer to **bool** | A boolean that specifies whether this is the recommended LDAP gateway version. | [optional] 
**Latest** | Pointer to **bool** | A boolean that specifies whether this is the latest LDAP gateway version. | [optional] 

## Methods

### NewGatewaySupportedVersions

`func NewGatewaySupportedVersions() *GatewaySupportedVersions`

NewGatewaySupportedVersions instantiates a new GatewaySupportedVersions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewaySupportedVersionsWithDefaults

`func NewGatewaySupportedVersionsWithDefaults() *GatewaySupportedVersions`

NewGatewaySupportedVersionsWithDefaults instantiates a new GatewaySupportedVersions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *GatewaySupportedVersions) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GatewaySupportedVersions) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GatewaySupportedVersions) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GatewaySupportedVersions) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetImage

`func (o *GatewaySupportedVersions) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *GatewaySupportedVersions) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *GatewaySupportedVersions) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *GatewaySupportedVersions) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetRecommended

`func (o *GatewaySupportedVersions) GetRecommended() bool`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *GatewaySupportedVersions) GetRecommendedOk() (*bool, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *GatewaySupportedVersions) SetRecommended(v bool)`

SetRecommended sets Recommended field to given value.

### HasRecommended

`func (o *GatewaySupportedVersions) HasRecommended() bool`

HasRecommended returns a boolean if a field has been set.

### GetLatest

`func (o *GatewaySupportedVersions) GetLatest() bool`

GetLatest returns the Latest field if non-nil, zero value otherwise.

### GetLatestOk

`func (o *GatewaySupportedVersions) GetLatestOk() (*bool, bool)`

GetLatestOk returns a tuple with the Latest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatest

`func (o *GatewaySupportedVersions) SetLatest(v bool)`

SetLatest sets Latest field to given value.

### HasLatest

`func (o *GatewaySupportedVersions) HasLatest() bool`

HasLatest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


