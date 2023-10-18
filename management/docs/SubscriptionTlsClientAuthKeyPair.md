# SubscriptionTlsClientAuthKeyPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of a key to be used for outbound mutual TLS (mTLS) authentication. This key is used as a client credential to authenticate the webhook. The key must have a &#x60;usageType&#x60; of &#x60;OUTBOUND_MTLS&#x60;. See [Certificate management](https://apidocs.pingidentity.com/pingone/platform/v1/api/#certificate-management) for information on creating a key. If this property is set, &#x60;verifyTlsCertificates&#x60; must be set to &#x60;true&#x60;. | [optional] 

## Methods

### NewSubscriptionTlsClientAuthKeyPair

`func NewSubscriptionTlsClientAuthKeyPair() *SubscriptionTlsClientAuthKeyPair`

NewSubscriptionTlsClientAuthKeyPair instantiates a new SubscriptionTlsClientAuthKeyPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionTlsClientAuthKeyPairWithDefaults

`func NewSubscriptionTlsClientAuthKeyPairWithDefaults() *SubscriptionTlsClientAuthKeyPair`

NewSubscriptionTlsClientAuthKeyPairWithDefaults instantiates a new SubscriptionTlsClientAuthKeyPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionTlsClientAuthKeyPair) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionTlsClientAuthKeyPair) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionTlsClientAuthKeyPair) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionTlsClientAuthKeyPair) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


