# TransactionConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeout** | Pointer to [**TransactionConfigurationTimeout**](TransactionConfigurationTimeout.md) |  | [optional] 
**DataCollection** | Pointer to [**TransactionConfigurationDataCollection**](TransactionConfigurationDataCollection.md) |  | [optional] 
**DataCollectionOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewTransactionConfiguration

`func NewTransactionConfiguration() *TransactionConfiguration`

NewTransactionConfiguration instantiates a new TransactionConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionConfigurationWithDefaults

`func NewTransactionConfigurationWithDefaults() *TransactionConfiguration`

NewTransactionConfigurationWithDefaults instantiates a new TransactionConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeout

`func (o *TransactionConfiguration) GetTimeout() TransactionConfigurationTimeout`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *TransactionConfiguration) GetTimeoutOk() (*TransactionConfigurationTimeout, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *TransactionConfiguration) SetTimeout(v TransactionConfigurationTimeout)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *TransactionConfiguration) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetDataCollection

`func (o *TransactionConfiguration) GetDataCollection() TransactionConfigurationDataCollection`

GetDataCollection returns the DataCollection field if non-nil, zero value otherwise.

### GetDataCollectionOk

`func (o *TransactionConfiguration) GetDataCollectionOk() (*TransactionConfigurationDataCollection, bool)`

GetDataCollectionOk returns a tuple with the DataCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCollection

`func (o *TransactionConfiguration) SetDataCollection(v TransactionConfigurationDataCollection)`

SetDataCollection sets DataCollection field to given value.

### HasDataCollection

`func (o *TransactionConfiguration) HasDataCollection() bool`

HasDataCollection returns a boolean if a field has been set.

### GetDataCollectionOnly

`func (o *TransactionConfiguration) GetDataCollectionOnly() bool`

GetDataCollectionOnly returns the DataCollectionOnly field if non-nil, zero value otherwise.

### GetDataCollectionOnlyOk

`func (o *TransactionConfiguration) GetDataCollectionOnlyOk() (*bool, bool)`

GetDataCollectionOnlyOk returns a tuple with the DataCollectionOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCollectionOnly

`func (o *TransactionConfiguration) SetDataCollectionOnly(v bool)`

SetDataCollectionOnly sets DataCollectionOnly field to given value.

### HasDataCollectionOnly

`func (o *TransactionConfiguration) HasDataCollectionOnly() bool`

HasDataCollectionOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


