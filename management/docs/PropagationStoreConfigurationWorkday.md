# PropagationStoreConfigurationWorkday

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeContingentWorkers** | Pointer to **bool** | Whether or not contingent workers are excluded. | [optional] 
**ExcludeEmployees** | Pointer to **bool** | Whether or not employees are excluded. | [optional] 
**ExcludeInactiveWorkers** | Pointer to **bool** | Whether or not inactive workers are excluded. | [optional] 
**Host** | Pointer to **string** | The Workday API host. | [optional] 
**Password** | **string** | The password for account authentication. | 
**TenantId** | **string** | The Workday tenant ID. | 
**Username** | **string** | The user name for account authentication. | 

## Methods

### NewPropagationStoreConfigurationWorkday

`func NewPropagationStoreConfigurationWorkday(password string, tenantId string, username string, ) *PropagationStoreConfigurationWorkday`

NewPropagationStoreConfigurationWorkday instantiates a new PropagationStoreConfigurationWorkday object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropagationStoreConfigurationWorkdayWithDefaults

`func NewPropagationStoreConfigurationWorkdayWithDefaults() *PropagationStoreConfigurationWorkday`

NewPropagationStoreConfigurationWorkdayWithDefaults instantiates a new PropagationStoreConfigurationWorkday object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeContingentWorkers

`func (o *PropagationStoreConfigurationWorkday) GetExcludeContingentWorkers() bool`

GetExcludeContingentWorkers returns the ExcludeContingentWorkers field if non-nil, zero value otherwise.

### GetExcludeContingentWorkersOk

`func (o *PropagationStoreConfigurationWorkday) GetExcludeContingentWorkersOk() (*bool, bool)`

GetExcludeContingentWorkersOk returns a tuple with the ExcludeContingentWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeContingentWorkers

`func (o *PropagationStoreConfigurationWorkday) SetExcludeContingentWorkers(v bool)`

SetExcludeContingentWorkers sets ExcludeContingentWorkers field to given value.

### HasExcludeContingentWorkers

`func (o *PropagationStoreConfigurationWorkday) HasExcludeContingentWorkers() bool`

HasExcludeContingentWorkers returns a boolean if a field has been set.

### GetExcludeEmployees

`func (o *PropagationStoreConfigurationWorkday) GetExcludeEmployees() bool`

GetExcludeEmployees returns the ExcludeEmployees field if non-nil, zero value otherwise.

### GetExcludeEmployeesOk

`func (o *PropagationStoreConfigurationWorkday) GetExcludeEmployeesOk() (*bool, bool)`

GetExcludeEmployeesOk returns a tuple with the ExcludeEmployees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeEmployees

`func (o *PropagationStoreConfigurationWorkday) SetExcludeEmployees(v bool)`

SetExcludeEmployees sets ExcludeEmployees field to given value.

### HasExcludeEmployees

`func (o *PropagationStoreConfigurationWorkday) HasExcludeEmployees() bool`

HasExcludeEmployees returns a boolean if a field has been set.

### GetExcludeInactiveWorkers

`func (o *PropagationStoreConfigurationWorkday) GetExcludeInactiveWorkers() bool`

GetExcludeInactiveWorkers returns the ExcludeInactiveWorkers field if non-nil, zero value otherwise.

### GetExcludeInactiveWorkersOk

`func (o *PropagationStoreConfigurationWorkday) GetExcludeInactiveWorkersOk() (*bool, bool)`

GetExcludeInactiveWorkersOk returns a tuple with the ExcludeInactiveWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeInactiveWorkers

`func (o *PropagationStoreConfigurationWorkday) SetExcludeInactiveWorkers(v bool)`

SetExcludeInactiveWorkers sets ExcludeInactiveWorkers field to given value.

### HasExcludeInactiveWorkers

`func (o *PropagationStoreConfigurationWorkday) HasExcludeInactiveWorkers() bool`

HasExcludeInactiveWorkers returns a boolean if a field has been set.

### GetHost

`func (o *PropagationStoreConfigurationWorkday) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *PropagationStoreConfigurationWorkday) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *PropagationStoreConfigurationWorkday) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *PropagationStoreConfigurationWorkday) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPassword

`func (o *PropagationStoreConfigurationWorkday) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PropagationStoreConfigurationWorkday) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PropagationStoreConfigurationWorkday) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetTenantId

`func (o *PropagationStoreConfigurationWorkday) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PropagationStoreConfigurationWorkday) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PropagationStoreConfigurationWorkday) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetUsername

`func (o *PropagationStoreConfigurationWorkday) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PropagationStoreConfigurationWorkday) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PropagationStoreConfigurationWorkday) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


