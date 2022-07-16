# GatewayInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connected** | Pointer to **bool** | A boolean that specifies whether or not the gateway instance has one or more connections. This is a required property. | [optional] 
**Id** | Pointer to **string** | A string that specifies the instance ID of the gateway. The gateway instance ID is created by the gateway when it starts up. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Gateway** | Pointer to [**GatewayInstanceGateway**](GatewayInstanceGateway.md) |  | [optional] 
**Credential** | Pointer to [**GatewayInstanceCredential**](GatewayInstanceCredential.md) |  | [optional] 
**CurrentErrors** | Pointer to **[]string** | An array of strings that lists the messages that are maintained by the gateway instance. | [optional] 
**HealthStatus** | Pointer to **string** | An enumeration that specifies whether or not the gateway is in a healthy state. Options are HEALTHY, DEGRADED, UNHEALTHY. This is a required property. | [optional] 
**Hostname** | Pointer to **string** | A string that specifies the hostname of the container running in the customer’s infrastructure. This is a required property. | [optional] 
**InitializedAt** | Pointer to **string** | A timestamp that specifies when gateway was initialized (when the first connect to PingOne was made). | [optional] 
**LastReportedAt** | Pointer to **string** | A timestamp that specifies the last reported timestamp, heartbeat, or other message. | [optional] 
**Version** | Pointer to [**GatewayInstanceVersion**](GatewayInstanceVersion.md) |  | [optional] 
**BusyPercentage** | Pointer to **int32** | An integer that specifies the gateway instance&#39;s busy percentage. When this percentage is high, then more instances should be added. | [optional] 
**OperationsPerSecond** | Pointer to **float32** | A number that specifies the recent throughput of the gateway instance. | [optional] 
**ResponseTimeMillis** | Pointer to **float32** | A number that specifies the processing time of the gateway instance in milliseconds. | [optional] 

## Methods

### NewGatewayInstance

`func NewGatewayInstance() *GatewayInstance`

NewGatewayInstance instantiates a new GatewayInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayInstanceWithDefaults

`func NewGatewayInstanceWithDefaults() *GatewayInstance`

NewGatewayInstanceWithDefaults instantiates a new GatewayInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnected

`func (o *GatewayInstance) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *GatewayInstance) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *GatewayInstance) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *GatewayInstance) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetId

`func (o *GatewayInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayInstance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GatewayInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *GatewayInstance) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *GatewayInstance) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *GatewayInstance) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *GatewayInstance) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetGateway

`func (o *GatewayInstance) GetGateway() GatewayInstanceGateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *GatewayInstance) GetGatewayOk() (*GatewayInstanceGateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *GatewayInstance) SetGateway(v GatewayInstanceGateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *GatewayInstance) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetCredential

`func (o *GatewayInstance) GetCredential() GatewayInstanceCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *GatewayInstance) GetCredentialOk() (*GatewayInstanceCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *GatewayInstance) SetCredential(v GatewayInstanceCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *GatewayInstance) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetCurrentErrors

`func (o *GatewayInstance) GetCurrentErrors() []string`

GetCurrentErrors returns the CurrentErrors field if non-nil, zero value otherwise.

### GetCurrentErrorsOk

`func (o *GatewayInstance) GetCurrentErrorsOk() (*[]string, bool)`

GetCurrentErrorsOk returns a tuple with the CurrentErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentErrors

`func (o *GatewayInstance) SetCurrentErrors(v []string)`

SetCurrentErrors sets CurrentErrors field to given value.

### HasCurrentErrors

`func (o *GatewayInstance) HasCurrentErrors() bool`

HasCurrentErrors returns a boolean if a field has been set.

### GetHealthStatus

`func (o *GatewayInstance) GetHealthStatus() string`

GetHealthStatus returns the HealthStatus field if non-nil, zero value otherwise.

### GetHealthStatusOk

`func (o *GatewayInstance) GetHealthStatusOk() (*string, bool)`

GetHealthStatusOk returns a tuple with the HealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatus

`func (o *GatewayInstance) SetHealthStatus(v string)`

SetHealthStatus sets HealthStatus field to given value.

### HasHealthStatus

`func (o *GatewayInstance) HasHealthStatus() bool`

HasHealthStatus returns a boolean if a field has been set.

### GetHostname

`func (o *GatewayInstance) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *GatewayInstance) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *GatewayInstance) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *GatewayInstance) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetInitializedAt

`func (o *GatewayInstance) GetInitializedAt() string`

GetInitializedAt returns the InitializedAt field if non-nil, zero value otherwise.

### GetInitializedAtOk

`func (o *GatewayInstance) GetInitializedAtOk() (*string, bool)`

GetInitializedAtOk returns a tuple with the InitializedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializedAt

`func (o *GatewayInstance) SetInitializedAt(v string)`

SetInitializedAt sets InitializedAt field to given value.

### HasInitializedAt

`func (o *GatewayInstance) HasInitializedAt() bool`

HasInitializedAt returns a boolean if a field has been set.

### GetLastReportedAt

`func (o *GatewayInstance) GetLastReportedAt() string`

GetLastReportedAt returns the LastReportedAt field if non-nil, zero value otherwise.

### GetLastReportedAtOk

`func (o *GatewayInstance) GetLastReportedAtOk() (*string, bool)`

GetLastReportedAtOk returns a tuple with the LastReportedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedAt

`func (o *GatewayInstance) SetLastReportedAt(v string)`

SetLastReportedAt sets LastReportedAt field to given value.

### HasLastReportedAt

`func (o *GatewayInstance) HasLastReportedAt() bool`

HasLastReportedAt returns a boolean if a field has been set.

### GetVersion

`func (o *GatewayInstance) GetVersion() GatewayInstanceVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GatewayInstance) GetVersionOk() (*GatewayInstanceVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GatewayInstance) SetVersion(v GatewayInstanceVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GatewayInstance) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetBusyPercentage

`func (o *GatewayInstance) GetBusyPercentage() int32`

GetBusyPercentage returns the BusyPercentage field if non-nil, zero value otherwise.

### GetBusyPercentageOk

`func (o *GatewayInstance) GetBusyPercentageOk() (*int32, bool)`

GetBusyPercentageOk returns a tuple with the BusyPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusyPercentage

`func (o *GatewayInstance) SetBusyPercentage(v int32)`

SetBusyPercentage sets BusyPercentage field to given value.

### HasBusyPercentage

`func (o *GatewayInstance) HasBusyPercentage() bool`

HasBusyPercentage returns a boolean if a field has been set.

### GetOperationsPerSecond

`func (o *GatewayInstance) GetOperationsPerSecond() float32`

GetOperationsPerSecond returns the OperationsPerSecond field if non-nil, zero value otherwise.

### GetOperationsPerSecondOk

`func (o *GatewayInstance) GetOperationsPerSecondOk() (*float32, bool)`

GetOperationsPerSecondOk returns a tuple with the OperationsPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationsPerSecond

`func (o *GatewayInstance) SetOperationsPerSecond(v float32)`

SetOperationsPerSecond sets OperationsPerSecond field to given value.

### HasOperationsPerSecond

`func (o *GatewayInstance) HasOperationsPerSecond() bool`

HasOperationsPerSecond returns a boolean if a field has been set.

### GetResponseTimeMillis

`func (o *GatewayInstance) GetResponseTimeMillis() float32`

GetResponseTimeMillis returns the ResponseTimeMillis field if non-nil, zero value otherwise.

### GetResponseTimeMillisOk

`func (o *GatewayInstance) GetResponseTimeMillisOk() (*float32, bool)`

GetResponseTimeMillisOk returns a tuple with the ResponseTimeMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeMillis

`func (o *GatewayInstance) SetResponseTimeMillis(v float32)`

SetResponseTimeMillis sets ResponseTimeMillis field to given value.

### HasResponseTimeMillis

`func (o *GatewayInstance) HasResponseTimeMillis() bool`

HasResponseTimeMillis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


