# TemplateContentEmailAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to [**TemplateContentEmailAllOfFrom**](TemplateContentEmailAllOfFrom.md) |  | [optional] 
**Subject** | Pointer to **string** | The email&#39;s subject line. Cannot exceed 256 characters. If supported, can include variables. | [optional] 
**ReplyTo** | Pointer to [**TemplateContentEmailAllOfReplyTo**](TemplateContentEmailAllOfReplyTo.md) |  | [optional] 
**Charset** | Pointer to **string** | If not specified, &#x60;UTF-8&#x60; is the default value. | [optional] [default to "UTF-8"]
**EmailContentType** | Pointer to **string** | If not specified, &#x60;text/html&#x60; is the default value. | [optional] [default to "text/html"]

## Methods

### NewTemplateContentEmailAllOf

`func NewTemplateContentEmailAllOf() *TemplateContentEmailAllOf`

NewTemplateContentEmailAllOf instantiates a new TemplateContentEmailAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateContentEmailAllOfWithDefaults

`func NewTemplateContentEmailAllOfWithDefaults() *TemplateContentEmailAllOf`

NewTemplateContentEmailAllOfWithDefaults instantiates a new TemplateContentEmailAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *TemplateContentEmailAllOf) GetFrom() TemplateContentEmailAllOfFrom`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *TemplateContentEmailAllOf) GetFromOk() (*TemplateContentEmailAllOfFrom, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *TemplateContentEmailAllOf) SetFrom(v TemplateContentEmailAllOfFrom)`

SetFrom sets From field to given value.

### HasFrom

`func (o *TemplateContentEmailAllOf) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetSubject

`func (o *TemplateContentEmailAllOf) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TemplateContentEmailAllOf) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TemplateContentEmailAllOf) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *TemplateContentEmailAllOf) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetReplyTo

`func (o *TemplateContentEmailAllOf) GetReplyTo() TemplateContentEmailAllOfReplyTo`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *TemplateContentEmailAllOf) GetReplyToOk() (*TemplateContentEmailAllOfReplyTo, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *TemplateContentEmailAllOf) SetReplyTo(v TemplateContentEmailAllOfReplyTo)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *TemplateContentEmailAllOf) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetCharset

`func (o *TemplateContentEmailAllOf) GetCharset() string`

GetCharset returns the Charset field if non-nil, zero value otherwise.

### GetCharsetOk

`func (o *TemplateContentEmailAllOf) GetCharsetOk() (*string, bool)`

GetCharsetOk returns a tuple with the Charset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharset

`func (o *TemplateContentEmailAllOf) SetCharset(v string)`

SetCharset sets Charset field to given value.

### HasCharset

`func (o *TemplateContentEmailAllOf) HasCharset() bool`

HasCharset returns a boolean if a field has been set.

### GetEmailContentType

`func (o *TemplateContentEmailAllOf) GetEmailContentType() string`

GetEmailContentType returns the EmailContentType field if non-nil, zero value otherwise.

### GetEmailContentTypeOk

`func (o *TemplateContentEmailAllOf) GetEmailContentTypeOk() (*string, bool)`

GetEmailContentTypeOk returns a tuple with the EmailContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailContentType

`func (o *TemplateContentEmailAllOf) SetEmailContentType(v string)`

SetEmailContentType sets EmailContentType field to given value.

### HasEmailContentType

`func (o *TemplateContentEmailAllOf) HasEmailContentType() bool`

HasEmailContentType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


