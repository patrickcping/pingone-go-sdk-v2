# UserPhoto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** | A string that specifies the URI that is a uniform resource locator (as defined in Section 1.1.3 of RFC 3986) that points to a resource location representing the user’s image. This can be removed from a user by setting the photo attribute to null. If provided, the resource must be a file (for example, a GIF, JPEG, or PNG image file) rather than a web page containing an image. It must be a valid URL that starts with the HTTP or HTTPS scheme. | 

## Methods

### NewUserPhoto

`func NewUserPhoto(href string, ) *UserPhoto`

NewUserPhoto instantiates a new UserPhoto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPhotoWithDefaults

`func NewUserPhotoWithDefaults() *UserPhoto`

NewUserPhotoWithDefaults instantiates a new UserPhoto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *UserPhoto) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *UserPhoto) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *UserPhoto) SetHref(v string)`

SetHref sets Href field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


