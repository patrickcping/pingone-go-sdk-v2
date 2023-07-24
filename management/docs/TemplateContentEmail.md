# TemplateContentEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**LinksHATEOAS**](LinksHATEOAS.md) |  | [optional] 
**Id** | Pointer to **string** | The template id. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | The time the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | The time the resource was last updated. | [optional] [readonly] 
**Default** | Pointer to **bool** | Specifies whether the template is a predefined default template. | [optional] 
**Locale** | **string** | A valid case-insensitive locale, complying with the ISO-639 language code and ISO-3166 country code standards: Two-character language code, for example, \&quot;en\&quot;. Two-character language code followed by a two-character country code, separated by an underscore or dash, for example: \&quot;en_GB\&quot;, \&quot;en-GB\&quot;. Cannot be changed after it is initially set in &#x60;POST /environments/{{envID}}/templates/{{templateName}}/contents&#x60;.  | 
**DeliveryMethod** | [**EnumTemplateContentDeliveryMethod**](EnumTemplateContentDeliveryMethod.md) |  | 
**Variant** | Pointer to **string** | Holds the unique user-defined name for each content variant that uses the same template + &#x60;deliveryMethod&#x60; + &#x60;locale&#x60; combination. This property is case insensitive and has a limit of 100 characters. For more information, see [Creating custom contents](https://apidocs.pingidentity.com/pingone/platform/v1/api/#notifications-templates-creating-custom-contents). | [optional] 
**Body** | **string** | The email text. Email text cannot be larger than 100 kB. Email text can contain HTML. If supported, this can include variables. | 
**From** | Pointer to [**TemplateContentEmailAllOfFrom**](TemplateContentEmailAllOfFrom.md) |  | [optional] 
**Subject** | Pointer to **string** | The email&#39;s subject line. Cannot exceed 256 characters. If supported, can include variables. | [optional] 
**ReplyTo** | Pointer to [**TemplateContentEmailAllOfReplyTo**](TemplateContentEmailAllOfReplyTo.md) |  | [optional] 
**Charset** | Pointer to **string** | If not specified, &#x60;UTF-8&#x60; is the default value. | [optional] [default to "UTF-8"]
**EmailContentType** | Pointer to **string** | If not specified, &#x60;text/html&#x60; is the default value. | [optional] [default to "text/html"]

## Methods

### NewTemplateContentEmail

`func NewTemplateContentEmail(locale string, deliveryMethod EnumTemplateContentDeliveryMethod, body string, ) *TemplateContentEmail`

NewTemplateContentEmail instantiates a new TemplateContentEmail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateContentEmailWithDefaults

`func NewTemplateContentEmailWithDefaults() *TemplateContentEmail`

NewTemplateContentEmailWithDefaults instantiates a new TemplateContentEmail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *TemplateContentEmail) GetLinks() LinksHATEOAS`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TemplateContentEmail) GetLinksOk() (*LinksHATEOAS, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TemplateContentEmail) SetLinks(v LinksHATEOAS)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TemplateContentEmail) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *TemplateContentEmail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateContentEmail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateContentEmail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateContentEmail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TemplateContentEmail) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TemplateContentEmail) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TemplateContentEmail) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TemplateContentEmail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TemplateContentEmail) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TemplateContentEmail) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TemplateContentEmail) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TemplateContentEmail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDefault

`func (o *TemplateContentEmail) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *TemplateContentEmail) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *TemplateContentEmail) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *TemplateContentEmail) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetLocale

`func (o *TemplateContentEmail) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *TemplateContentEmail) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *TemplateContentEmail) SetLocale(v string)`

SetLocale sets Locale field to given value.


### GetDeliveryMethod

`func (o *TemplateContentEmail) GetDeliveryMethod() EnumTemplateContentDeliveryMethod`

GetDeliveryMethod returns the DeliveryMethod field if non-nil, zero value otherwise.

### GetDeliveryMethodOk

`func (o *TemplateContentEmail) GetDeliveryMethodOk() (*EnumTemplateContentDeliveryMethod, bool)`

GetDeliveryMethodOk returns a tuple with the DeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethod

`func (o *TemplateContentEmail) SetDeliveryMethod(v EnumTemplateContentDeliveryMethod)`

SetDeliveryMethod sets DeliveryMethod field to given value.


### GetVariant

`func (o *TemplateContentEmail) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *TemplateContentEmail) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *TemplateContentEmail) SetVariant(v string)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *TemplateContentEmail) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### GetBody

`func (o *TemplateContentEmail) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *TemplateContentEmail) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *TemplateContentEmail) SetBody(v string)`

SetBody sets Body field to given value.


### GetFrom

`func (o *TemplateContentEmail) GetFrom() TemplateContentEmailAllOfFrom`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *TemplateContentEmail) GetFromOk() (*TemplateContentEmailAllOfFrom, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *TemplateContentEmail) SetFrom(v TemplateContentEmailAllOfFrom)`

SetFrom sets From field to given value.

### HasFrom

`func (o *TemplateContentEmail) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetSubject

`func (o *TemplateContentEmail) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TemplateContentEmail) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TemplateContentEmail) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *TemplateContentEmail) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetReplyTo

`func (o *TemplateContentEmail) GetReplyTo() TemplateContentEmailAllOfReplyTo`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *TemplateContentEmail) GetReplyToOk() (*TemplateContentEmailAllOfReplyTo, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *TemplateContentEmail) SetReplyTo(v TemplateContentEmailAllOfReplyTo)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *TemplateContentEmail) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetCharset

`func (o *TemplateContentEmail) GetCharset() string`

GetCharset returns the Charset field if non-nil, zero value otherwise.

### GetCharsetOk

`func (o *TemplateContentEmail) GetCharsetOk() (*string, bool)`

GetCharsetOk returns a tuple with the Charset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharset

`func (o *TemplateContentEmail) SetCharset(v string)`

SetCharset sets Charset field to given value.

### HasCharset

`func (o *TemplateContentEmail) HasCharset() bool`

HasCharset returns a boolean if a field has been set.

### GetEmailContentType

`func (o *TemplateContentEmail) GetEmailContentType() string`

GetEmailContentType returns the EmailContentType field if non-nil, zero value otherwise.

### GetEmailContentTypeOk

`func (o *TemplateContentEmail) GetEmailContentTypeOk() (*string, bool)`

GetEmailContentTypeOk returns a tuple with the EmailContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailContentType

`func (o *TemplateContentEmail) SetEmailContentType(v string)`

SetEmailContentType sets EmailContentType field to given value.

### HasEmailContentType

`func (o *TemplateContentEmail) HasEmailContentType() bool`

HasEmailContentType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


