# RateLimitConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The UUID of the rate limit configuration. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | The date and time the rate limit configuration was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The date and time the rate limit configuration was updated. | [optional] [readonly] 
**Type** | [**EnumRateLimitConfigurationType**](EnumRateLimitConfigurationType.md) |  | 
**Value** | **string** | The IP address (v4 or v6), or a CIDR range, for the IP address or addresses to be excluded. | 

## Methods

### NewRateLimitConfiguration

`func NewRateLimitConfiguration(type_ EnumRateLimitConfigurationType, value string, ) *RateLimitConfiguration`

NewRateLimitConfiguration instantiates a new RateLimitConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitConfigurationWithDefaults

`func NewRateLimitConfigurationWithDefaults() *RateLimitConfiguration`

NewRateLimitConfigurationWithDefaults instantiates a new RateLimitConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RateLimitConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RateLimitConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RateLimitConfiguration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RateLimitConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RateLimitConfiguration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RateLimitConfiguration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RateLimitConfiguration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RateLimitConfiguration) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RateLimitConfiguration) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RateLimitConfiguration) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RateLimitConfiguration) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RateLimitConfiguration) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetType

`func (o *RateLimitConfiguration) GetType() EnumRateLimitConfigurationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RateLimitConfiguration) GetTypeOk() (*EnumRateLimitConfigurationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RateLimitConfiguration) SetType(v EnumRateLimitConfigurationType)`

SetType sets Type field to given value.


### GetValue

`func (o *RateLimitConfiguration) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RateLimitConfiguration) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RateLimitConfiguration) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


