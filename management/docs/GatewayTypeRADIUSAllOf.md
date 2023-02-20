# GatewayTypeRADIUSAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Davinci** | [**GatewayTypeRADIUSAllOfDavinci**](GatewayTypeRADIUSAllOfDavinci.md) |  | 
**DefaultSharedSecret** | Pointer to **string** | Value to use for the shared secret if the shared secret is not provided for one or more of the RADIUS clients specified. | [optional] 
**RadiusClients** | [**[]GatewayTypeRADIUSAllOfRadiusClients**](GatewayTypeRADIUSAllOfRadiusClients.md) | Collection of RADIUS clients. | 

## Methods

### NewGatewayTypeRADIUSAllOf

`func NewGatewayTypeRADIUSAllOf(davinci GatewayTypeRADIUSAllOfDavinci, radiusClients []GatewayTypeRADIUSAllOfRadiusClients, ) *GatewayTypeRADIUSAllOf`

NewGatewayTypeRADIUSAllOf instantiates a new GatewayTypeRADIUSAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayTypeRADIUSAllOfWithDefaults

`func NewGatewayTypeRADIUSAllOfWithDefaults() *GatewayTypeRADIUSAllOf`

NewGatewayTypeRADIUSAllOfWithDefaults instantiates a new GatewayTypeRADIUSAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDavinci

`func (o *GatewayTypeRADIUSAllOf) GetDavinci() GatewayTypeRADIUSAllOfDavinci`

GetDavinci returns the Davinci field if non-nil, zero value otherwise.

### GetDavinciOk

`func (o *GatewayTypeRADIUSAllOf) GetDavinciOk() (*GatewayTypeRADIUSAllOfDavinci, bool)`

GetDavinciOk returns a tuple with the Davinci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDavinci

`func (o *GatewayTypeRADIUSAllOf) SetDavinci(v GatewayTypeRADIUSAllOfDavinci)`

SetDavinci sets Davinci field to given value.


### GetDefaultSharedSecret

`func (o *GatewayTypeRADIUSAllOf) GetDefaultSharedSecret() string`

GetDefaultSharedSecret returns the DefaultSharedSecret field if non-nil, zero value otherwise.

### GetDefaultSharedSecretOk

`func (o *GatewayTypeRADIUSAllOf) GetDefaultSharedSecretOk() (*string, bool)`

GetDefaultSharedSecretOk returns a tuple with the DefaultSharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSharedSecret

`func (o *GatewayTypeRADIUSAllOf) SetDefaultSharedSecret(v string)`

SetDefaultSharedSecret sets DefaultSharedSecret field to given value.

### HasDefaultSharedSecret

`func (o *GatewayTypeRADIUSAllOf) HasDefaultSharedSecret() bool`

HasDefaultSharedSecret returns a boolean if a field has been set.

### GetRadiusClients

`func (o *GatewayTypeRADIUSAllOf) GetRadiusClients() []GatewayTypeRADIUSAllOfRadiusClients`

GetRadiusClients returns the RadiusClients field if non-nil, zero value otherwise.

### GetRadiusClientsOk

`func (o *GatewayTypeRADIUSAllOf) GetRadiusClientsOk() (*[]GatewayTypeRADIUSAllOfRadiusClients, bool)`

GetRadiusClientsOk returns a tuple with the RadiusClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusClients

`func (o *GatewayTypeRADIUSAllOf) SetRadiusClients(v []GatewayTypeRADIUSAllOfRadiusClients)`

SetRadiusClients sets RadiusClients field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


