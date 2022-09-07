# SubscriptionHttpEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | A string that specifies a valid HTTPS URL to which event messages are sent. This is a required property. | 
**Headers** | Pointer to **[]map[string]interface{}** | An array that specifies the headers applied to the outbound request (for example, &#x60;Authorization&#x60; &#x60;Basic usernamepassword&#x60;. The purpose of these headers is for the HTTPS endpoint to authenticate the PingOne service, ensuring that the information from PingOne is from a trusted source. | [optional] 

## Methods

### NewSubscriptionHttpEndpoint

`func NewSubscriptionHttpEndpoint(url string, ) *SubscriptionHttpEndpoint`

NewSubscriptionHttpEndpoint instantiates a new SubscriptionHttpEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionHttpEndpointWithDefaults

`func NewSubscriptionHttpEndpointWithDefaults() *SubscriptionHttpEndpoint`

NewSubscriptionHttpEndpointWithDefaults instantiates a new SubscriptionHttpEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *SubscriptionHttpEndpoint) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SubscriptionHttpEndpoint) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SubscriptionHttpEndpoint) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeaders

`func (o *SubscriptionHttpEndpoint) GetHeaders() []map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *SubscriptionHttpEndpoint) GetHeadersOk() (*[]map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *SubscriptionHttpEndpoint) SetHeaders(v []map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *SubscriptionHttpEndpoint) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


