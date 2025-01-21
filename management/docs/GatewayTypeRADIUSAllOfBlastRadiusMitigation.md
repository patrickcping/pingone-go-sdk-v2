# GatewayTypeRADIUSAllOfBlastRadiusMitigation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequireMsgAuth** | Pointer to **bool** | Set to &#x60;true&#x60; to require that all requests from the client include the Message-Authenticator attribute. Any requests without the attribute will be ignored. | [optional] 
**LimitProxyState** | Pointer to **bool** | For older clients that don&#39;t support sending the Message-Authenticator attribute, you can set this field to &#x60;true&#x60;. This instructs the gateway to ignore requests that don&#39;t contain the Message-Authenticator attribute but contain the Proxy-State attribute. | [optional] 

## Methods

### NewGatewayTypeRADIUSAllOfBlastRadiusMitigation

`func NewGatewayTypeRADIUSAllOfBlastRadiusMitigation() *GatewayTypeRADIUSAllOfBlastRadiusMitigation`

NewGatewayTypeRADIUSAllOfBlastRadiusMitigation instantiates a new GatewayTypeRADIUSAllOfBlastRadiusMitigation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayTypeRADIUSAllOfBlastRadiusMitigationWithDefaults

`func NewGatewayTypeRADIUSAllOfBlastRadiusMitigationWithDefaults() *GatewayTypeRADIUSAllOfBlastRadiusMitigation`

NewGatewayTypeRADIUSAllOfBlastRadiusMitigationWithDefaults instantiates a new GatewayTypeRADIUSAllOfBlastRadiusMitigation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequireMsgAuth

`func (o *GatewayTypeRADIUSAllOfBlastRadiusMitigation) GetRequireMsgAuth() bool`

GetRequireMsgAuth returns the RequireMsgAuth field if non-nil, zero value otherwise.

### GetRequireMsgAuthOk

`func (o *GatewayTypeRADIUSAllOfBlastRadiusMitigation) GetRequireMsgAuthOk() (*bool, bool)`

GetRequireMsgAuthOk returns a tuple with the RequireMsgAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireMsgAuth

`func (o *GatewayTypeRADIUSAllOfBlastRadiusMitigation) SetRequireMsgAuth(v bool)`

SetRequireMsgAuth sets RequireMsgAuth field to given value.

### HasRequireMsgAuth

`func (o *GatewayTypeRADIUSAllOfBlastRadiusMitigation) HasRequireMsgAuth() bool`

HasRequireMsgAuth returns a boolean if a field has been set.

### GetLimitProxyState

`func (o *GatewayTypeRADIUSAllOfBlastRadiusMitigation) GetLimitProxyState() bool`

GetLimitProxyState returns the LimitProxyState field if non-nil, zero value otherwise.

### GetLimitProxyStateOk

`func (o *GatewayTypeRADIUSAllOfBlastRadiusMitigation) GetLimitProxyStateOk() (*bool, bool)`

GetLimitProxyStateOk returns a tuple with the LimitProxyState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitProxyState

`func (o *GatewayTypeRADIUSAllOfBlastRadiusMitigation) SetLimitProxyState(v bool)`

SetLimitProxyState sets LimitProxyState field to given value.

### HasLimitProxyState

`func (o *GatewayTypeRADIUSAllOfBlastRadiusMitigation) HasLimitProxyState() bool`

HasLimitProxyState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


