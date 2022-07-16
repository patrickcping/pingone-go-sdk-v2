# ImageTargets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | A string that specifies the URL or fully qualified path to the image source file. | [optional] [readonly] 
**Id** | Pointer to **string** | A string that specifies the UUID of the target image. | [optional] [readonly] 
**Type** | Pointer to **string** | A string that specifies the type of format used for the image. Options are jpg, png, and gif. | [optional] [readonly] 
**Width** | Pointer to **int32** | The width of the image (in pixels). | [optional] [readonly] 
**Height** | Pointer to **int32** | The height of the image (in pixels). | [optional] [readonly] 

## Methods

### NewImageTargets

`func NewImageTargets() *ImageTargets`

NewImageTargets instantiates a new ImageTargets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageTargetsWithDefaults

`func NewImageTargetsWithDefaults() *ImageTargets`

NewImageTargetsWithDefaults instantiates a new ImageTargets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *ImageTargets) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ImageTargets) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ImageTargets) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ImageTargets) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *ImageTargets) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageTargets) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageTargets) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ImageTargets) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ImageTargets) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImageTargets) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImageTargets) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ImageTargets) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWidth

`func (o *ImageTargets) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ImageTargets) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ImageTargets) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ImageTargets) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *ImageTargets) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ImageTargets) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ImageTargets) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ImageTargets) HasHeight() bool`

HasHeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


