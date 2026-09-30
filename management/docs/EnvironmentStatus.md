# EnvironmentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**EnumEnvironmentStatus**](EnumEnvironmentStatus.md) |  | 
**License** | Pointer to [**EnvironmentStatusLicense**](EnvironmentStatusLicense.md) |  | [optional] 

## Methods

### NewEnvironmentStatus

`func NewEnvironmentStatus(status EnumEnvironmentStatus, ) *EnvironmentStatus`

NewEnvironmentStatus instantiates a new EnvironmentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentStatusWithDefaults

`func NewEnvironmentStatusWithDefaults() *EnvironmentStatus`

NewEnvironmentStatusWithDefaults instantiates a new EnvironmentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *EnvironmentStatus) GetStatus() EnumEnvironmentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnvironmentStatus) GetStatusOk() (*EnumEnvironmentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnvironmentStatus) SetStatus(v EnumEnvironmentStatus)`

SetStatus sets Status field to given value.


### GetLicense

`func (o *EnvironmentStatus) GetLicense() EnvironmentStatusLicense`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *EnvironmentStatus) GetLicenseOk() (*EnvironmentStatusLicense, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *EnvironmentStatus) SetLicense(v EnvironmentStatusLicense)`

SetLicense sets License field to given value.

### HasLicense

`func (o *EnvironmentStatus) HasLicense() bool`

HasLicense returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


