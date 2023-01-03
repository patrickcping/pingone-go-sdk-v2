# TemplateContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The template id. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | The time the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | The time the resource was last updated. | [optional] [readonly] 
**Default** | Pointer to **bool** | Specifies whether the template is a predefined default template. | [optional] 
**Locale** | **string** | A valid case-insensitive locale, complying with the ISO-639 language code and ISO-3166 country code standards: Two-character language code, for example, \&quot;en\&quot;. Two-character language code followed by a two-character country code, separated by an underscore or dash, for example: \&quot;en_GB\&quot;, \&quot;en-GB\&quot;. Cannot be changed after it is initially set in &#x60;POST /environments/{{envID}}/templates/{{templateName}}/contents&#x60;.  | 
**DeliveryMethod** | **string** | The content&#39;s delivery method. Possible values are &#x60;Email&#x60;, &#x60;SMS&#x60;, &#x60;Voice&#x60; or &#x60;Push&#x60;. Cannot be changed after it is initially set in &#x60;POST /environments/{{envID}}/templates/{{templateName}}/contents&#x60;. | 
**Variant** | Pointer to **string** | Holds the unique user-defined name for each content variant that uses the same template + &#x60;deliveryMethod&#x60; + &#x60;locale&#x60; combination. This property is case insensitive and has a limit of 100 characters. For more information, see [Creating custom contents](https://apidocs.pingidentity.com/pingone/platform/v1/api/#notifications-templates-creating-custom-contents). | [optional] 

## Methods

### NewTemplateContent

`func NewTemplateContent(locale string, deliveryMethod string, ) *TemplateContent`

NewTemplateContent instantiates a new TemplateContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateContentWithDefaults

`func NewTemplateContentWithDefaults() *TemplateContent`

NewTemplateContentWithDefaults instantiates a new TemplateContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateContent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateContent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateContent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateContent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TemplateContent) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TemplateContent) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TemplateContent) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TemplateContent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TemplateContent) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TemplateContent) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TemplateContent) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TemplateContent) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDefault

`func (o *TemplateContent) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *TemplateContent) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *TemplateContent) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *TemplateContent) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetLocale

`func (o *TemplateContent) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *TemplateContent) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *TemplateContent) SetLocale(v string)`

SetLocale sets Locale field to given value.


### GetDeliveryMethod

`func (o *TemplateContent) GetDeliveryMethod() string`

GetDeliveryMethod returns the DeliveryMethod field if non-nil, zero value otherwise.

### GetDeliveryMethodOk

`func (o *TemplateContent) GetDeliveryMethodOk() (*string, bool)`

GetDeliveryMethodOk returns a tuple with the DeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethod

`func (o *TemplateContent) SetDeliveryMethod(v string)`

SetDeliveryMethod sets DeliveryMethod field to given value.


### GetVariant

`func (o *TemplateContent) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *TemplateContent) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *TemplateContent) SetVariant(v string)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *TemplateContent) HasVariant() bool`

HasVariant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


