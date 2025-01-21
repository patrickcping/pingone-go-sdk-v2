# GatewayTypeRADIUSAllOfRadiusClients

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlastRadiusMitigation** | Pointer to [**GatewayTypeRADIUSAllOfBlastRadiusMitigation**](GatewayTypeRADIUSAllOfBlastRadiusMitigation.md) |  | [optional] 
**Ip** | **string** | The IP of the RADIUS client. | 
**SharedSecret** | Pointer to **string** | The shared secret for the RADIUS client. If this value is not provided, the shared secret specified with &#x60;defaultSharedSecret&#x60; is used. If you are not providing a shared secret for the client, leave out &#x60;sharedSecret&#x60; or set it to null. | [optional] 

## Methods

### NewGatewayTypeRADIUSAllOfRadiusClients

`func NewGatewayTypeRADIUSAllOfRadiusClients(ip string, ) *GatewayTypeRADIUSAllOfRadiusClients`

NewGatewayTypeRADIUSAllOfRadiusClients instantiates a new GatewayTypeRADIUSAllOfRadiusClients object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayTypeRADIUSAllOfRadiusClientsWithDefaults

`func NewGatewayTypeRADIUSAllOfRadiusClientsWithDefaults() *GatewayTypeRADIUSAllOfRadiusClients`

NewGatewayTypeRADIUSAllOfRadiusClientsWithDefaults instantiates a new GatewayTypeRADIUSAllOfRadiusClients object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlastRadiusMitigation

`func (o *GatewayTypeRADIUSAllOfRadiusClients) GetBlastRadiusMitigation() GatewayTypeRADIUSAllOfBlastRadiusMitigation`

GetBlastRadiusMitigation returns the BlastRadiusMitigation field if non-nil, zero value otherwise.

### GetBlastRadiusMitigationOk

`func (o *GatewayTypeRADIUSAllOfRadiusClients) GetBlastRadiusMitigationOk() (*GatewayTypeRADIUSAllOfBlastRadiusMitigation, bool)`

GetBlastRadiusMitigationOk returns a tuple with the BlastRadiusMitigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlastRadiusMitigation

`func (o *GatewayTypeRADIUSAllOfRadiusClients) SetBlastRadiusMitigation(v GatewayTypeRADIUSAllOfBlastRadiusMitigation)`

SetBlastRadiusMitigation sets BlastRadiusMitigation field to given value.

### HasBlastRadiusMitigation

`func (o *GatewayTypeRADIUSAllOfRadiusClients) HasBlastRadiusMitigation() bool`

HasBlastRadiusMitigation returns a boolean if a field has been set.

### GetIp

`func (o *GatewayTypeRADIUSAllOfRadiusClients) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GatewayTypeRADIUSAllOfRadiusClients) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GatewayTypeRADIUSAllOfRadiusClients) SetIp(v string)`

SetIp sets Ip field to given value.


### GetSharedSecret

`func (o *GatewayTypeRADIUSAllOfRadiusClients) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *GatewayTypeRADIUSAllOfRadiusClients) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *GatewayTypeRADIUSAllOfRadiusClients) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *GatewayTypeRADIUSAllOfRadiusClients) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


