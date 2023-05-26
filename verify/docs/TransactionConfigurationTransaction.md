# TransactionConfigurationTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeout** | Pointer to [**TransactionConfigurationTransactionTimeout**](TransactionConfigurationTransactionTimeout.md) |  | [optional] 
**DataCollection** | Pointer to [**TransactionConfigurationTransactionDataCollection**](TransactionConfigurationTransactionDataCollection.md) |  | [optional] 
**DataCollectionOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewTransactionConfigurationTransaction

`func NewTransactionConfigurationTransaction() *TransactionConfigurationTransaction`

NewTransactionConfigurationTransaction instantiates a new TransactionConfigurationTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionConfigurationTransactionWithDefaults

`func NewTransactionConfigurationTransactionWithDefaults() *TransactionConfigurationTransaction`

NewTransactionConfigurationTransactionWithDefaults instantiates a new TransactionConfigurationTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeout

`func (o *TransactionConfigurationTransaction) GetTimeout() TransactionConfigurationTransactionTimeout`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *TransactionConfigurationTransaction) GetTimeoutOk() (*TransactionConfigurationTransactionTimeout, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *TransactionConfigurationTransaction) SetTimeout(v TransactionConfigurationTransactionTimeout)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *TransactionConfigurationTransaction) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetDataCollection

`func (o *TransactionConfigurationTransaction) GetDataCollection() TransactionConfigurationTransactionDataCollection`

GetDataCollection returns the DataCollection field if non-nil, zero value otherwise.

### GetDataCollectionOk

`func (o *TransactionConfigurationTransaction) GetDataCollectionOk() (*TransactionConfigurationTransactionDataCollection, bool)`

GetDataCollectionOk returns a tuple with the DataCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCollection

`func (o *TransactionConfigurationTransaction) SetDataCollection(v TransactionConfigurationTransactionDataCollection)`

SetDataCollection sets DataCollection field to given value.

### HasDataCollection

`func (o *TransactionConfigurationTransaction) HasDataCollection() bool`

HasDataCollection returns a boolean if a field has been set.

### GetDataCollectionOnly

`func (o *TransactionConfigurationTransaction) GetDataCollectionOnly() bool`

GetDataCollectionOnly returns the DataCollectionOnly field if non-nil, zero value otherwise.

### GetDataCollectionOnlyOk

`func (o *TransactionConfigurationTransaction) GetDataCollectionOnlyOk() (*bool, bool)`

GetDataCollectionOnlyOk returns a tuple with the DataCollectionOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCollectionOnly

`func (o *TransactionConfigurationTransaction) SetDataCollectionOnly(v bool)`

SetDataCollectionOnly sets DataCollectionOnly field to given value.

### HasDataCollectionOnly

`func (o *TransactionConfigurationTransaction) HasDataCollectionOnly() bool`

HasDataCollectionOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


