# ApplicationSecretPrevious

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | Pointer to **string** | A string that specifies the previous application secret. This property is returned in the response if the previous secret is not expired. | [optional] [readonly] 
**ExpiresAt** | **time.Time** | A timestamp that specifies how long this secret is saved (and can be used) before it expires. Supported time range is 1 minute to 30 days. | 
**LastUsed** | Pointer to **time.Time** | A timestamp that specifies when the previous secret was last used. | [optional] [readonly] 

## Methods

### NewApplicationSecretPrevious

`func NewApplicationSecretPrevious(expiresAt time.Time, ) *ApplicationSecretPrevious`

NewApplicationSecretPrevious instantiates a new ApplicationSecretPrevious object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationSecretPreviousWithDefaults

`func NewApplicationSecretPreviousWithDefaults() *ApplicationSecretPrevious`

NewApplicationSecretPreviousWithDefaults instantiates a new ApplicationSecretPrevious object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecret

`func (o *ApplicationSecretPrevious) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ApplicationSecretPrevious) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ApplicationSecretPrevious) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ApplicationSecretPrevious) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ApplicationSecretPrevious) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ApplicationSecretPrevious) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ApplicationSecretPrevious) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetLastUsed

`func (o *ApplicationSecretPrevious) GetLastUsed() time.Time`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *ApplicationSecretPrevious) GetLastUsedOk() (*time.Time, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *ApplicationSecretPrevious) SetLastUsed(v time.Time)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *ApplicationSecretPrevious) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


