# CredentialTypeVersionNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifier (UUID) of this version of the credential type. The service assigns identifiers. | [optional] [readonly] 
**Number** | Pointer to **int32** | Version of this credential type. The service assigns versions. | [optional] [readonly] 
**Uri** | Pointer to **string** | A URI to of this version of the credential type. The service assigns URIs. | [optional] [readonly] 

## Methods

### NewCredentialTypeVersionNumber

`func NewCredentialTypeVersionNumber() *CredentialTypeVersionNumber`

NewCredentialTypeVersionNumber instantiates a new CredentialTypeVersionNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialTypeVersionNumberWithDefaults

`func NewCredentialTypeVersionNumberWithDefaults() *CredentialTypeVersionNumber`

NewCredentialTypeVersionNumberWithDefaults instantiates a new CredentialTypeVersionNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CredentialTypeVersionNumber) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CredentialTypeVersionNumber) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CredentialTypeVersionNumber) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CredentialTypeVersionNumber) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumber

`func (o *CredentialTypeVersionNumber) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CredentialTypeVersionNumber) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CredentialTypeVersionNumber) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *CredentialTypeVersionNumber) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetUri

`func (o *CredentialTypeVersionNumber) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *CredentialTypeVersionNumber) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *CredentialTypeVersionNumber) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *CredentialTypeVersionNumber) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


