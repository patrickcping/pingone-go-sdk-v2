# TemplateContentPushAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | The push title (maximum 200 characters). If supported, this can include variables. | 
**Body** | **string** | The push text (maximum 400 characters for push text). If supported, this can include variables. | 
**PushCategory** | Pointer to [**EnumTemplateContentPushCategory**](EnumTemplateContentPushCategory.md) |  | [optional] 

## Methods

### NewTemplateContentPushAllOf

`func NewTemplateContentPushAllOf(title string, body string, ) *TemplateContentPushAllOf`

NewTemplateContentPushAllOf instantiates a new TemplateContentPushAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateContentPushAllOfWithDefaults

`func NewTemplateContentPushAllOfWithDefaults() *TemplateContentPushAllOf`

NewTemplateContentPushAllOfWithDefaults instantiates a new TemplateContentPushAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *TemplateContentPushAllOf) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TemplateContentPushAllOf) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TemplateContentPushAllOf) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *TemplateContentPushAllOf) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *TemplateContentPushAllOf) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *TemplateContentPushAllOf) SetBody(v string)`

SetBody sets Body field to given value.


### GetPushCategory

`func (o *TemplateContentPushAllOf) GetPushCategory() EnumTemplateContentPushCategory`

GetPushCategory returns the PushCategory field if non-nil, zero value otherwise.

### GetPushCategoryOk

`func (o *TemplateContentPushAllOf) GetPushCategoryOk() (*EnumTemplateContentPushCategory, bool)`

GetPushCategoryOk returns a tuple with the PushCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushCategory

`func (o *TemplateContentPushAllOf) SetPushCategory(v EnumTemplateContentPushCategory)`

SetPushCategory sets PushCategory field to given value.

### HasPushCategory

`func (o *TemplateContentPushAllOf) HasPushCategory() bool`

HasPushCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


