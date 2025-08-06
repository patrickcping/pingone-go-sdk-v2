# CredentialTypeVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifier (UUID) of this version of the credential type. The service assigns identifiers. | [optional] [readonly] 
**Number** | Pointer to **int32** | Version of this credential type. The service assigns versions. | [optional] [readonly] 
**Uri** | Pointer to **string** | A URI to of this version of the credential type. The service assigns URIs. | [optional] [readonly] 

## Methods

### NewCredentialTypeVersion

`func NewCredentialTypeVersion() *CredentialTypeVersion`

NewCredentialTypeVersion instantiates a new CredentialTypeVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialTypeVersionWithDefaults

`func NewCredentialTypeVersionWithDefaults() *CredentialTypeVersion`

NewCredentialTypeVersionWithDefaults instantiates a new CredentialTypeVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CredentialTypeVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CredentialTypeVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CredentialTypeVersion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CredentialTypeVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumber

`func (o *CredentialTypeVersion) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CredentialTypeVersion) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CredentialTypeVersion) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *CredentialTypeVersion) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetUri

`func (o *CredentialTypeVersion) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *CredentialTypeVersion) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *CredentialTypeVersion) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *CredentialTypeVersion) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


