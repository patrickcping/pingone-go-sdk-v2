# PropagationStoreSyncStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to **string** | Details of any synchronization errors. | [optional] [readonly] 
**FailedCount** | Pointer to **int32** | A count of failed synchronization events since the last revision. | [optional] [readonly] 
**FailedDeprovisionCount** | Pointer to **int32** | A count of failed deprovisioning synchronization events since the last revision. | [optional] [readonly] 
**LastSyncAt** | Pointer to **time.Time** | The last synchronization in yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSS&#39;Z&#39; format. | [optional] [readonly] 
**SuccessCount** | Pointer to **int32** | A count of successful synchronization events since the last revision. | [optional] [readonly] 
**SyncState** | Pointer to [**EnumPropagationStoreSyncState**](EnumPropagationStoreSyncState.md) |  | [optional] 
**UserTotal** | Pointer to **int32** | A count of users that will synchronize to the target store based on the Rule’s filter. | [optional] [readonly] 

## Methods

### NewPropagationStoreSyncStatus

`func NewPropagationStoreSyncStatus() *PropagationStoreSyncStatus`

NewPropagationStoreSyncStatus instantiates a new PropagationStoreSyncStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropagationStoreSyncStatusWithDefaults

`func NewPropagationStoreSyncStatusWithDefaults() *PropagationStoreSyncStatus`

NewPropagationStoreSyncStatusWithDefaults instantiates a new PropagationStoreSyncStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *PropagationStoreSyncStatus) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *PropagationStoreSyncStatus) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *PropagationStoreSyncStatus) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *PropagationStoreSyncStatus) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetFailedCount

`func (o *PropagationStoreSyncStatus) GetFailedCount() int32`

GetFailedCount returns the FailedCount field if non-nil, zero value otherwise.

### GetFailedCountOk

`func (o *PropagationStoreSyncStatus) GetFailedCountOk() (*int32, bool)`

GetFailedCountOk returns a tuple with the FailedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedCount

`func (o *PropagationStoreSyncStatus) SetFailedCount(v int32)`

SetFailedCount sets FailedCount field to given value.

### HasFailedCount

`func (o *PropagationStoreSyncStatus) HasFailedCount() bool`

HasFailedCount returns a boolean if a field has been set.

### GetFailedDeprovisionCount

`func (o *PropagationStoreSyncStatus) GetFailedDeprovisionCount() int32`

GetFailedDeprovisionCount returns the FailedDeprovisionCount field if non-nil, zero value otherwise.

### GetFailedDeprovisionCountOk

`func (o *PropagationStoreSyncStatus) GetFailedDeprovisionCountOk() (*int32, bool)`

GetFailedDeprovisionCountOk returns a tuple with the FailedDeprovisionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedDeprovisionCount

`func (o *PropagationStoreSyncStatus) SetFailedDeprovisionCount(v int32)`

SetFailedDeprovisionCount sets FailedDeprovisionCount field to given value.

### HasFailedDeprovisionCount

`func (o *PropagationStoreSyncStatus) HasFailedDeprovisionCount() bool`

HasFailedDeprovisionCount returns a boolean if a field has been set.

### GetLastSyncAt

`func (o *PropagationStoreSyncStatus) GetLastSyncAt() time.Time`

GetLastSyncAt returns the LastSyncAt field if non-nil, zero value otherwise.

### GetLastSyncAtOk

`func (o *PropagationStoreSyncStatus) GetLastSyncAtOk() (*time.Time, bool)`

GetLastSyncAtOk returns a tuple with the LastSyncAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncAt

`func (o *PropagationStoreSyncStatus) SetLastSyncAt(v time.Time)`

SetLastSyncAt sets LastSyncAt field to given value.

### HasLastSyncAt

`func (o *PropagationStoreSyncStatus) HasLastSyncAt() bool`

HasLastSyncAt returns a boolean if a field has been set.

### GetSuccessCount

`func (o *PropagationStoreSyncStatus) GetSuccessCount() int32`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *PropagationStoreSyncStatus) GetSuccessCountOk() (*int32, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCount

`func (o *PropagationStoreSyncStatus) SetSuccessCount(v int32)`

SetSuccessCount sets SuccessCount field to given value.

### HasSuccessCount

`func (o *PropagationStoreSyncStatus) HasSuccessCount() bool`

HasSuccessCount returns a boolean if a field has been set.

### GetSyncState

`func (o *PropagationStoreSyncStatus) GetSyncState() EnumPropagationStoreSyncState`

GetSyncState returns the SyncState field if non-nil, zero value otherwise.

### GetSyncStateOk

`func (o *PropagationStoreSyncStatus) GetSyncStateOk() (*EnumPropagationStoreSyncState, bool)`

GetSyncStateOk returns a tuple with the SyncState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncState

`func (o *PropagationStoreSyncStatus) SetSyncState(v EnumPropagationStoreSyncState)`

SetSyncState sets SyncState field to given value.

### HasSyncState

`func (o *PropagationStoreSyncStatus) HasSyncState() bool`

HasSyncState returns a boolean if a field has been set.

### GetUserTotal

`func (o *PropagationStoreSyncStatus) GetUserTotal() int32`

GetUserTotal returns the UserTotal field if non-nil, zero value otherwise.

### GetUserTotalOk

`func (o *PropagationStoreSyncStatus) GetUserTotalOk() (*int32, bool)`

GetUserTotalOk returns a tuple with the UserTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTotal

`func (o *PropagationStoreSyncStatus) SetUserTotal(v int32)`

SetUserTotal sets UserTotal field to given value.

### HasUserTotal

`func (o *PropagationStoreSyncStatus) HasUserTotal() bool`

HasUserTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


