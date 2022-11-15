# TemplateContentEmailAllOfReplyTo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The email&#39;s \&quot;reply to\&quot; name. If the environment uses the Ping Identity email sender, the name \&quot;PingOne\&quot; is used. You can configure other email \&quot;reply to\&quot; names per environment. See [Note](https://apidocs.pingidentity.com/pingone/platform/v1/api/#from-replyTo-note) for details.  | [optional] [default to "PingOne"]
**Address** | Pointer to **string** | The \&quot;reply to\&quot; email address. If the environment uses the Ping Identity email sender, or if the address field is empty, the address \&quot;noreply@pingidentity.com\&quot; is used. You can configure other email \&quot;reply to\&quot; addresses per environment. See [Note](https://apidocs.pingidentity.com/pingone/platform/v1/api/#from-replyTo-note) for details.  | [optional] [default to "noreply@pingidentity.com"]

## Methods

### NewTemplateContentEmailAllOfReplyTo

`func NewTemplateContentEmailAllOfReplyTo() *TemplateContentEmailAllOfReplyTo`

NewTemplateContentEmailAllOfReplyTo instantiates a new TemplateContentEmailAllOfReplyTo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateContentEmailAllOfReplyToWithDefaults

`func NewTemplateContentEmailAllOfReplyToWithDefaults() *TemplateContentEmailAllOfReplyTo`

NewTemplateContentEmailAllOfReplyToWithDefaults instantiates a new TemplateContentEmailAllOfReplyTo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TemplateContentEmailAllOfReplyTo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateContentEmailAllOfReplyTo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateContentEmailAllOfReplyTo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplateContentEmailAllOfReplyTo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *TemplateContentEmailAllOfReplyTo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TemplateContentEmailAllOfReplyTo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TemplateContentEmailAllOfReplyTo) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TemplateContentEmailAllOfReplyTo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


