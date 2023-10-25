# BillOfMaterialsProductsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A string that specifies the BOM ID | [optional] [readonly] 
**Type** | [**EnumProductType**](EnumProductType.md) |  | 
**Description** | Pointer to **string** | A string that specifies the description of the product or standalone service | [optional] [readonly] 
**Console** | Pointer to [**BillOfMaterialsProductsInnerConsole**](BillOfMaterialsProductsInnerConsole.md) |  | [optional] 
**Tags** | Pointer to [**[]EnumBillOfMaterialsProductTags**](EnumBillOfMaterialsProductTags.md) |  | [optional] 
**Deployment** | Pointer to [**BillOfMaterialsProductsInnerDeployment**](BillOfMaterialsProductsInnerDeployment.md) |  | [optional] 
**Bookmarks** | Pointer to [**[]BillOfMaterialsProductsInnerBookmarksInner**](BillOfMaterialsProductsInnerBookmarksInner.md) | Optional array of custom bookmarks. Maximum of five bookmarks per product. | [optional] 

## Methods

### NewBillOfMaterialsProductsInner

`func NewBillOfMaterialsProductsInner(type_ EnumProductType, ) *BillOfMaterialsProductsInner`

NewBillOfMaterialsProductsInner instantiates a new BillOfMaterialsProductsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillOfMaterialsProductsInnerWithDefaults

`func NewBillOfMaterialsProductsInnerWithDefaults() *BillOfMaterialsProductsInner`

NewBillOfMaterialsProductsInnerWithDefaults instantiates a new BillOfMaterialsProductsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillOfMaterialsProductsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillOfMaterialsProductsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillOfMaterialsProductsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillOfMaterialsProductsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *BillOfMaterialsProductsInner) GetType() EnumProductType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BillOfMaterialsProductsInner) GetTypeOk() (*EnumProductType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BillOfMaterialsProductsInner) SetType(v EnumProductType)`

SetType sets Type field to given value.


### GetDescription

`func (o *BillOfMaterialsProductsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillOfMaterialsProductsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillOfMaterialsProductsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillOfMaterialsProductsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConsole

`func (o *BillOfMaterialsProductsInner) GetConsole() BillOfMaterialsProductsInnerConsole`

GetConsole returns the Console field if non-nil, zero value otherwise.

### GetConsoleOk

`func (o *BillOfMaterialsProductsInner) GetConsoleOk() (*BillOfMaterialsProductsInnerConsole, bool)`

GetConsoleOk returns a tuple with the Console field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsole

`func (o *BillOfMaterialsProductsInner) SetConsole(v BillOfMaterialsProductsInnerConsole)`

SetConsole sets Console field to given value.

### HasConsole

`func (o *BillOfMaterialsProductsInner) HasConsole() bool`

HasConsole returns a boolean if a field has been set.

### GetTags

`func (o *BillOfMaterialsProductsInner) GetTags() []EnumBillOfMaterialsProductTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BillOfMaterialsProductsInner) GetTagsOk() (*[]EnumBillOfMaterialsProductTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BillOfMaterialsProductsInner) SetTags(v []EnumBillOfMaterialsProductTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BillOfMaterialsProductsInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetDeployment

`func (o *BillOfMaterialsProductsInner) GetDeployment() BillOfMaterialsProductsInnerDeployment`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *BillOfMaterialsProductsInner) GetDeploymentOk() (*BillOfMaterialsProductsInnerDeployment, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *BillOfMaterialsProductsInner) SetDeployment(v BillOfMaterialsProductsInnerDeployment)`

SetDeployment sets Deployment field to given value.

### HasDeployment

`func (o *BillOfMaterialsProductsInner) HasDeployment() bool`

HasDeployment returns a boolean if a field has been set.

### GetBookmarks

`func (o *BillOfMaterialsProductsInner) GetBookmarks() []BillOfMaterialsProductsInnerBookmarksInner`

GetBookmarks returns the Bookmarks field if non-nil, zero value otherwise.

### GetBookmarksOk

`func (o *BillOfMaterialsProductsInner) GetBookmarksOk() (*[]BillOfMaterialsProductsInnerBookmarksInner, bool)`

GetBookmarksOk returns a tuple with the Bookmarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookmarks

`func (o *BillOfMaterialsProductsInner) SetBookmarks(v []BillOfMaterialsProductsInnerBookmarksInner)`

SetBookmarks sets Bookmarks field to given value.

### HasBookmarks

`func (o *BillOfMaterialsProductsInner) HasBookmarks() bool`

HasBookmarks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


