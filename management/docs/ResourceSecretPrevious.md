# ResourceSecretPrevious

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | Pointer to **string** | A string that specifies the resource&#39;s previous secret. This property is returned in the response if the previous secret is not expired. | [optional] [readonly] 
**ExpiresAt** | **time.Time** | A timestamp that specifies how long this secret is saved (and can be used) before it expires. Supported time range is 1 minute to 30 days. | 
**LastUsed** | Pointer to **time.Time** | A timestamp that specifies when the previous secret was last used. | [optional] [readonly] 

## Methods

### NewResourceSecretPrevious

`func NewResourceSecretPrevious(expiresAt time.Time, ) *ResourceSecretPrevious`

NewResourceSecretPrevious instantiates a new ResourceSecretPrevious object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSecretPreviousWithDefaults

`func NewResourceSecretPreviousWithDefaults() *ResourceSecretPrevious`

NewResourceSecretPreviousWithDefaults instantiates a new ResourceSecretPrevious object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecret

`func (o *ResourceSecretPrevious) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ResourceSecretPrevious) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ResourceSecretPrevious) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ResourceSecretPrevious) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ResourceSecretPrevious) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ResourceSecretPrevious) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ResourceSecretPrevious) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetLastUsed

`func (o *ResourceSecretPrevious) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *ResourceSecretPrevious) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *ResourceSecretPrevious) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *ResourceSecretPrevious) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


