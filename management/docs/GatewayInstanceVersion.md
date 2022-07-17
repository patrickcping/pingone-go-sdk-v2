# GatewayInstanceVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionNumber** | Pointer to **string** | A string that specifies the version number of the gateway running for the instance. This is a required property. | [optional] 
**UpdateStatus** | Pointer to [**EnumUpdateStatus**](EnumUpdateStatus.md) |  | [optional] 

## Methods

### NewGatewayInstanceVersion

`func NewGatewayInstanceVersion() *GatewayInstanceVersion`

NewGatewayInstanceVersion instantiates a new GatewayInstanceVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayInstanceVersionWithDefaults

`func NewGatewayInstanceVersionWithDefaults() *GatewayInstanceVersion`

NewGatewayInstanceVersionWithDefaults instantiates a new GatewayInstanceVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionNumber

`func (o *GatewayInstanceVersion) GetVersionNumber() string`

GetVersionNumber returns the VersionNumber field if non-nil, zero value otherwise.

### GetVersionNumberOk

`func (o *GatewayInstanceVersion) GetVersionNumberOk() (*string, bool)`

GetVersionNumberOk returns a tuple with the VersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionNumber

`func (o *GatewayInstanceVersion) SetVersionNumber(v string)`

SetVersionNumber sets VersionNumber field to given value.

### HasVersionNumber

`func (o *GatewayInstanceVersion) HasVersionNumber() bool`

HasVersionNumber returns a boolean if a field has been set.

### GetUpdateStatus

`func (o *GatewayInstanceVersion) GetUpdateStatus() EnumUpdateStatus`

GetUpdateStatus returns the UpdateStatus field if non-nil, zero value otherwise.

### GetUpdateStatusOk

`func (o *GatewayInstanceVersion) GetUpdateStatusOk() (*EnumUpdateStatus, bool)`

GetUpdateStatusOk returns a tuple with the UpdateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateStatus

`func (o *GatewayInstanceVersion) SetUpdateStatus(v EnumUpdateStatus)`

SetUpdateStatus sets UpdateStatus field to given value.

### HasUpdateStatus

`func (o *GatewayInstanceVersion) HasUpdateStatus() bool`

HasUpdateStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


