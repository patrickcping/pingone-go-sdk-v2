# PropagationStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**LinksHATEOAS**](LinksHATEOAS.md) |  | [optional] 
**Configuration** | [**PropagationStoreConfiguration**](PropagationStoreConfiguration.md) |  | 
**Description** | Pointer to **string** | A description for the identity propagation store. | [optional] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
**Image** | Pointer to [**PropagationStoreImage**](PropagationStoreImage.md) |  | [optional] 
**Managed** | Pointer to **bool** | Indicates whether or not to enable deprovisioning of users for a store when it is deleted. The deprovisioning occurs when a new revision is created (&#x60;POST {{apiPath}}/environments/{{envID}}/propagation/revisions&#x60;). | [optional] 
**Name** | **string** | The name of the identity store. | 
**Status** | Pointer to [**EnumPropagationStoreStatus**](EnumPropagationStoreStatus.md) |  | [optional] 
**SyncStatus** | Pointer to [**PropagationStoreSyncStatus**](PropagationStoreSyncStatus.md) |  | [optional] 
**Type** | [**EnumPropagationStoreType**](EnumPropagationStoreType.md) |  | 

## Methods

### NewPropagationStore

`func NewPropagationStore(configuration PropagationStoreConfiguration, name string, type_ EnumPropagationStoreType, ) *PropagationStore`

NewPropagationStore instantiates a new PropagationStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropagationStoreWithDefaults

`func NewPropagationStoreWithDefaults() *PropagationStore`

NewPropagationStoreWithDefaults instantiates a new PropagationStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *PropagationStore) GetLinks() LinksHATEOAS`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PropagationStore) GetLinksOk() (*LinksHATEOAS, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PropagationStore) SetLinks(v LinksHATEOAS)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PropagationStore) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetConfiguration

`func (o *PropagationStore) GetConfiguration() PropagationStoreConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *PropagationStore) GetConfigurationOk() (*PropagationStoreConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *PropagationStore) SetConfiguration(v PropagationStoreConfiguration)`

SetConfiguration sets Configuration field to given value.


### GetDescription

`func (o *PropagationStore) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PropagationStore) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PropagationStore) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PropagationStore) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *PropagationStore) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PropagationStore) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PropagationStore) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *PropagationStore) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *PropagationStore) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PropagationStore) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PropagationStore) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PropagationStore) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *PropagationStore) GetImage() PropagationStoreImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PropagationStore) GetImageOk() (*PropagationStoreImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PropagationStore) SetImage(v PropagationStoreImage)`

SetImage sets Image field to given value.

### HasImage

`func (o *PropagationStore) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetManaged

`func (o *PropagationStore) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *PropagationStore) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *PropagationStore) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *PropagationStore) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetName

`func (o *PropagationStore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PropagationStore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PropagationStore) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *PropagationStore) GetStatus() EnumPropagationStoreStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PropagationStore) GetStatusOk() (*EnumPropagationStoreStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PropagationStore) SetStatus(v EnumPropagationStoreStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PropagationStore) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSyncStatus

`func (o *PropagationStore) GetSyncStatus() PropagationStoreSyncStatus`

GetSyncStatus returns the SyncStatus field if non-nil, zero value otherwise.

### GetSyncStatusOk

`func (o *PropagationStore) GetSyncStatusOk() (*PropagationStoreSyncStatus, bool)`

GetSyncStatusOk returns a tuple with the SyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStatus

`func (o *PropagationStore) SetSyncStatus(v PropagationStoreSyncStatus)`

SetSyncStatus sets SyncStatus field to given value.

### HasSyncStatus

`func (o *PropagationStore) HasSyncStatus() bool`

HasSyncStatus returns a boolean if a field has been set.

### GetType

`func (o *PropagationStore) GetType() EnumPropagationStoreType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PropagationStore) GetTypeOk() (*EnumPropagationStoreType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PropagationStore) SetType(v EnumPropagationStoreType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


