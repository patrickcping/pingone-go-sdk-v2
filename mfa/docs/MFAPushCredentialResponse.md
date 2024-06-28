# MFAPushCredentialResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Id** | Pointer to **string** | A string that specifies the push credential ID. | [optional] [readonly] 
**Type** | Pointer to [**EnumMFAPushCredentialAttrType**](EnumMFAPushCredentialAttrType.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The time the resource was last updated. | [optional] [readonly] 

## Methods

### NewMFAPushCredentialResponse

`func NewMFAPushCredentialResponse() *MFAPushCredentialResponse`

NewMFAPushCredentialResponse instantiates a new MFAPushCredentialResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMFAPushCredentialResponseWithDefaults

`func NewMFAPushCredentialResponseWithDefaults() *MFAPushCredentialResponse`

NewMFAPushCredentialResponseWithDefaults instantiates a new MFAPushCredentialResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *MFAPushCredentialResponse) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MFAPushCredentialResponse) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MFAPushCredentialResponse) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MFAPushCredentialResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *MFAPushCredentialResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MFAPushCredentialResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MFAPushCredentialResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MFAPushCredentialResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *MFAPushCredentialResponse) GetType() EnumMFAPushCredentialAttrType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MFAPushCredentialResponse) GetTypeOk() (*EnumMFAPushCredentialAttrType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MFAPushCredentialResponse) SetType(v EnumMFAPushCredentialAttrType)`

SetType sets Type field to given value.

### HasType

`func (o *MFAPushCredentialResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MFAPushCredentialResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MFAPushCredentialResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MFAPushCredentialResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MFAPushCredentialResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MFAPushCredentialResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MFAPushCredentialResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MFAPushCredentialResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MFAPushCredentialResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


