# ImageTargetsOriginal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | A string that specifies the URL or fully qualified path to the image source file. | [optional] [readonly] 
**Id** | Pointer to **string** | A string that specifies the UUID of the target image. | [optional] [readonly] 
**Type** | Pointer to [**EnumImageFormat**](EnumImageFormat.md) |  | [optional] 
**Width** | Pointer to **int32** | The width of the image (in pixels). | [optional] [readonly] 
**Height** | Pointer to **int32** | The height of the image (in pixels). | [optional] [readonly] 

## Methods

### NewImageTargetsOriginal

`func NewImageTargetsOriginal() *ImageTargetsOriginal`

NewImageTargetsOriginal instantiates a new ImageTargetsOriginal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageTargetsOriginalWithDefaults

`func NewImageTargetsOriginalWithDefaults() *ImageTargetsOriginal`

NewImageTargetsOriginalWithDefaults instantiates a new ImageTargetsOriginal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *ImageTargetsOriginal) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ImageTargetsOriginal) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ImageTargetsOriginal) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ImageTargetsOriginal) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *ImageTargetsOriginal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageTargetsOriginal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageTargetsOriginal) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ImageTargetsOriginal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ImageTargetsOriginal) GetType() EnumImageFormat`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImageTargetsOriginal) GetTypeOk() (*EnumImageFormat, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImageTargetsOriginal) SetType(v EnumImageFormat)`

SetType sets Type field to given value.

### HasType

`func (o *ImageTargetsOriginal) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWidth

`func (o *ImageTargetsOriginal) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ImageTargetsOriginal) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ImageTargetsOriginal) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ImageTargetsOriginal) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *ImageTargetsOriginal) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ImageTargetsOriginal) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ImageTargetsOriginal) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ImageTargetsOriginal) HasHeight() bool`

HasHeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


