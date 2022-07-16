# PopulationPasswordPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the password policy that is used for this population. If absent, the environment&#39;s default is used. Requried if &#x60;passwordPolicy&#x60; is used. | 

## Methods

### NewPopulationPasswordPolicy

`func NewPopulationPasswordPolicy(id string, ) *PopulationPasswordPolicy`

NewPopulationPasswordPolicy instantiates a new PopulationPasswordPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPopulationPasswordPolicyWithDefaults

`func NewPopulationPasswordPolicyWithDefaults() *PopulationPasswordPolicy`

NewPopulationPasswordPolicyWithDefaults instantiates a new PopulationPasswordPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PopulationPasswordPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PopulationPasswordPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PopulationPasswordPolicy) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


