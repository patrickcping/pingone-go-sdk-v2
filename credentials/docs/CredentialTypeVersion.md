# CredentialTypeVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | A string that specifies the date and time the credential type version was created. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**CredentialType** | Pointer to [**ObjectCredentialType**](ObjectCredentialType.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the identifier (UUID) of the credential type version. | [optional] [readonly] 
**Snapshot** | Pointer to [**CredentialType**](CredentialType.md) |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] [readonly] 

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

### GetLinks

`func (o *CredentialTypeVersion) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CredentialTypeVersion) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CredentialTypeVersion) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CredentialTypeVersion) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CredentialTypeVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CredentialTypeVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CredentialTypeVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CredentialTypeVersion) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEnvironment

`func (o *CredentialTypeVersion) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CredentialTypeVersion) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CredentialTypeVersion) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CredentialTypeVersion) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetCredentialType

`func (o *CredentialTypeVersion) GetCredentialType() ObjectCredentialType`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *CredentialTypeVersion) GetCredentialTypeOk() (*ObjectCredentialType, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *CredentialTypeVersion) SetCredentialType(v ObjectCredentialType)`

SetCredentialType sets CredentialType field to given value.

### HasCredentialType

`func (o *CredentialTypeVersion) HasCredentialType() bool`

HasCredentialType returns a boolean if a field has been set.

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

### GetSnapshot

`func (o *CredentialTypeVersion) GetSnapshot() CredentialType`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *CredentialTypeVersion) GetSnapshotOk() (*CredentialType, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *CredentialTypeVersion) SetSnapshot(v CredentialType)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *CredentialTypeVersion) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetVersion

`func (o *CredentialTypeVersion) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CredentialTypeVersion) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CredentialTypeVersion) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CredentialTypeVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


