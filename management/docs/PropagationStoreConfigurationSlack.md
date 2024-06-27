# PropagationStoreConfigurationSlack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CREATE_USERS** | Pointer to **bool** | Whether or not users are allowed to be created. | [optional] 
**DEPROVISION_USERS** | Pointer to **bool** | Whether or not users are allowed to be deprovisioned (removed) following action specified in &#x60;REMOVE_ACTION&#x60;. | [optional] 
**DISABLE_USERS** | Pointer to **bool** | Whether or not users are allowed to be disabled. Must be set to &#x60;true&#x60;. | [optional] 
**OAUTH_ACCESS_TOKEN** | **string** | OAuth 2 access token. | 
**REMOVE_ACTION** | Pointer to [**EnumPropagationStoreTypeRemoveActionDisable**](EnumPropagationStoreTypeRemoveActionDisable.md) |  | [optional] 
**UNIQUE_USER_IDENTIFIER** | [**EnumPropagationStoreTypeSlackUniqueUserIdentifier**](EnumPropagationStoreTypeSlackUniqueUserIdentifier.md) |  | 
**UPDATE_USERS** | Pointer to **bool** | Whether or not users are allowed to be updated. | [optional] 

## Methods

### NewPropagationStoreConfigurationSlack

`func NewPropagationStoreConfigurationSlack(oAUTHACCESSTOKEN string, uNIQUEUSERIDENTIFIER EnumPropagationStoreTypeSlackUniqueUserIdentifier, ) *PropagationStoreConfigurationSlack`

NewPropagationStoreConfigurationSlack instantiates a new PropagationStoreConfigurationSlack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropagationStoreConfigurationSlackWithDefaults

`func NewPropagationStoreConfigurationSlackWithDefaults() *PropagationStoreConfigurationSlack`

NewPropagationStoreConfigurationSlackWithDefaults instantiates a new PropagationStoreConfigurationSlack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCREATE_USERS

`func (o *PropagationStoreConfigurationSlack) GetCREATE_USERS() bool`

GetCREATE_USERS returns the CREATE_USERS field if non-nil, zero value otherwise.

### GetCREATE_USERSOk

`func (o *PropagationStoreConfigurationSlack) GetCREATE_USERSOk() (*bool, bool)`

GetCREATE_USERSOk returns a tuple with the CREATE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATE_USERS

`func (o *PropagationStoreConfigurationSlack) SetCREATE_USERS(v bool)`

SetCREATE_USERS sets CREATE_USERS field to given value.

### HasCREATE_USERS

`func (o *PropagationStoreConfigurationSlack) HasCREATE_USERS() bool`

HasCREATE_USERS returns a boolean if a field has been set.

### GetDEPROVISION_USERS

`func (o *PropagationStoreConfigurationSlack) GetDEPROVISION_USERS() bool`

GetDEPROVISION_USERS returns the DEPROVISION_USERS field if non-nil, zero value otherwise.

### GetDEPROVISION_USERSOk

`func (o *PropagationStoreConfigurationSlack) GetDEPROVISION_USERSOk() (*bool, bool)`

GetDEPROVISION_USERSOk returns a tuple with the DEPROVISION_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEPROVISION_USERS

`func (o *PropagationStoreConfigurationSlack) SetDEPROVISION_USERS(v bool)`

SetDEPROVISION_USERS sets DEPROVISION_USERS field to given value.

### HasDEPROVISION_USERS

`func (o *PropagationStoreConfigurationSlack) HasDEPROVISION_USERS() bool`

HasDEPROVISION_USERS returns a boolean if a field has been set.

### GetDISABLE_USERS

`func (o *PropagationStoreConfigurationSlack) GetDISABLE_USERS() bool`

GetDISABLE_USERS returns the DISABLE_USERS field if non-nil, zero value otherwise.

### GetDISABLE_USERSOk

`func (o *PropagationStoreConfigurationSlack) GetDISABLE_USERSOk() (*bool, bool)`

GetDISABLE_USERSOk returns a tuple with the DISABLE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDISABLE_USERS

`func (o *PropagationStoreConfigurationSlack) SetDISABLE_USERS(v bool)`

SetDISABLE_USERS sets DISABLE_USERS field to given value.

### HasDISABLE_USERS

`func (o *PropagationStoreConfigurationSlack) HasDISABLE_USERS() bool`

HasDISABLE_USERS returns a boolean if a field has been set.

### GetOAUTH_ACCESS_TOKEN

`func (o *PropagationStoreConfigurationSlack) GetOAUTH_ACCESS_TOKEN() string`

GetOAUTH_ACCESS_TOKEN returns the OAUTH_ACCESS_TOKEN field if non-nil, zero value otherwise.

### GetOAUTH_ACCESS_TOKENOk

`func (o *PropagationStoreConfigurationSlack) GetOAUTH_ACCESS_TOKENOk() (*string, bool)`

GetOAUTH_ACCESS_TOKENOk returns a tuple with the OAUTH_ACCESS_TOKEN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAUTH_ACCESS_TOKEN

`func (o *PropagationStoreConfigurationSlack) SetOAUTH_ACCESS_TOKEN(v string)`

SetOAUTH_ACCESS_TOKEN sets OAUTH_ACCESS_TOKEN field to given value.


### GetREMOVE_ACTION

`func (o *PropagationStoreConfigurationSlack) GetREMOVE_ACTION() EnumPropagationStoreTypeRemoveActionDisable`

GetREMOVE_ACTION returns the REMOVE_ACTION field if non-nil, zero value otherwise.

### GetREMOVE_ACTIONOk

`func (o *PropagationStoreConfigurationSlack) GetREMOVE_ACTIONOk() (*EnumPropagationStoreTypeRemoveActionDisable, bool)`

GetREMOVE_ACTIONOk returns a tuple with the REMOVE_ACTION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVE_ACTION

`func (o *PropagationStoreConfigurationSlack) SetREMOVE_ACTION(v EnumPropagationStoreTypeRemoveActionDisable)`

SetREMOVE_ACTION sets REMOVE_ACTION field to given value.

### HasREMOVE_ACTION

`func (o *PropagationStoreConfigurationSlack) HasREMOVE_ACTION() bool`

HasREMOVE_ACTION returns a boolean if a field has been set.

### GetUNIQUE_USER_IDENTIFIER

`func (o *PropagationStoreConfigurationSlack) GetUNIQUE_USER_IDENTIFIER() EnumPropagationStoreTypeSlackUniqueUserIdentifier`

GetUNIQUE_USER_IDENTIFIER returns the UNIQUE_USER_IDENTIFIER field if non-nil, zero value otherwise.

### GetUNIQUE_USER_IDENTIFIEROk

`func (o *PropagationStoreConfigurationSlack) GetUNIQUE_USER_IDENTIFIEROk() (*EnumPropagationStoreTypeSlackUniqueUserIdentifier, bool)`

GetUNIQUE_USER_IDENTIFIEROk returns a tuple with the UNIQUE_USER_IDENTIFIER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUNIQUE_USER_IDENTIFIER

`func (o *PropagationStoreConfigurationSlack) SetUNIQUE_USER_IDENTIFIER(v EnumPropagationStoreTypeSlackUniqueUserIdentifier)`

SetUNIQUE_USER_IDENTIFIER sets UNIQUE_USER_IDENTIFIER field to given value.


### GetUPDATE_USERS

`func (o *PropagationStoreConfigurationSlack) GetUPDATE_USERS() bool`

GetUPDATE_USERS returns the UPDATE_USERS field if non-nil, zero value otherwise.

### GetUPDATE_USERSOk

`func (o *PropagationStoreConfigurationSlack) GetUPDATE_USERSOk() (*bool, bool)`

GetUPDATE_USERSOk returns a tuple with the UPDATE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATE_USERS

`func (o *PropagationStoreConfigurationSlack) SetUPDATE_USERS(v bool)`

SetUPDATE_USERS sets UPDATE_USERS field to given value.

### HasUPDATE_USERS

`func (o *PropagationStoreConfigurationSlack) HasUPDATE_USERS() bool`

HasUPDATE_USERS returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


