# TransactionConfigurationTransactionDataCollectionTimeout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int32** | Length of time before data collection timeout expires; range is 0-30 minutes or 0-1800 seconds. | 
**TimeUnit** | [**EnumShortTimeUnit**](EnumShortTimeUnit.md) |  | 

## Methods

### NewTransactionConfigurationTransactionDataCollectionTimeout

`func NewTransactionConfigurationTransactionDataCollectionTimeout(duration int32, timeUnit EnumShortTimeUnit, ) *TransactionConfigurationTransactionDataCollectionTimeout`

NewTransactionConfigurationTransactionDataCollectionTimeout instantiates a new TransactionConfigurationTransactionDataCollectionTimeout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionConfigurationTransactionDataCollectionTimeoutWithDefaults

`func NewTransactionConfigurationTransactionDataCollectionTimeoutWithDefaults() *TransactionConfigurationTransactionDataCollectionTimeout`

NewTransactionConfigurationTransactionDataCollectionTimeoutWithDefaults instantiates a new TransactionConfigurationTransactionDataCollectionTimeout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *TransactionConfigurationTransactionDataCollectionTimeout) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TransactionConfigurationTransactionDataCollectionTimeout) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TransactionConfigurationTransactionDataCollectionTimeout) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetTimeUnit

`func (o *TransactionConfigurationTransactionDataCollectionTimeout) GetTimeUnit() EnumShortTimeUnit`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *TransactionConfigurationTransactionDataCollectionTimeout) GetTimeUnitOk() (*EnumShortTimeUnit, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *TransactionConfigurationTransactionDataCollectionTimeout) SetTimeUnit(v EnumShortTimeUnit)`

SetTimeUnit sets TimeUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


