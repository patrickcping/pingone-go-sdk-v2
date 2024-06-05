# ApplicationResourceParent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumApplicationResourceParentType**](EnumApplicationResourceParentType.md) |  | 
**Id** | **string** | The application resource&#39;s parent ID. | 

## Methods

### NewApplicationResourceParent

`func NewApplicationResourceParent(type_ EnumApplicationResourceParentType, id string, ) *ApplicationResourceParent`

NewApplicationResourceParent instantiates a new ApplicationResourceParent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationResourceParentWithDefaults

`func NewApplicationResourceParentWithDefaults() *ApplicationResourceParent`

NewApplicationResourceParentWithDefaults instantiates a new ApplicationResourceParent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationResourceParent) GetType() EnumApplicationResourceParentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationResourceParent) GetTypeOk() (*EnumApplicationResourceParentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationResourceParent) SetType(v EnumApplicationResourceParentType)`

SetType sets Type field to given value.


### GetId

`func (o *ApplicationResourceParent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationResourceParent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationResourceParent) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


