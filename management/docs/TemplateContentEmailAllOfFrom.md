# TemplateContentEmailAllOfFrom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The email&#39;s sender name. If the environment uses the Ping Identity email sender, the name \&quot;PingOne\&quot; is used. You can configure other email sender names per environment. See [Note](https://apidocs.pingidentity.com/pingone/platform/v1/api/#from-replyTo-note) for details.  | [optional] [default to "PingOne"]
**Address** | Pointer to **string** | The sender email address. If the environment uses the Ping Identity email sender, or if the address field is empty, the address \&quot;noreply@pingidentity.com\&quot; is used. You can configure other email sender addresses per environment. See [Note](https://apidocs.pingidentity.com/pingone/platform/v1/api/#from-replyTo-note) for details.  | [optional] [default to "noreply@pingidentity.com"]

## Methods

### NewTemplateContentEmailAllOfFrom

`func NewTemplateContentEmailAllOfFrom() *TemplateContentEmailAllOfFrom`

NewTemplateContentEmailAllOfFrom instantiates a new TemplateContentEmailAllOfFrom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateContentEmailAllOfFromWithDefaults

`func NewTemplateContentEmailAllOfFromWithDefaults() *TemplateContentEmailAllOfFrom`

NewTemplateContentEmailAllOfFromWithDefaults instantiates a new TemplateContentEmailAllOfFrom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TemplateContentEmailAllOfFrom) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateContentEmailAllOfFrom) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateContentEmailAllOfFrom) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplateContentEmailAllOfFrom) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *TemplateContentEmailAllOfFrom) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TemplateContentEmailAllOfFrom) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TemplateContentEmailAllOfFrom) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TemplateContentEmailAllOfFrom) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


