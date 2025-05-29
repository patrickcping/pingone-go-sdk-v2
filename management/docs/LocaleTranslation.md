# LocaleTranslation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The translation ID for a specific string. | [optional] [readonly] 
**Key** | **string** | The page and name of the interface element to be localized (for example, &#x60;flow-ui.button.cancel&#x60;). | 
**ShortKey** | Pointer to **string** | The interace element only (for example, &#x60;button.cancel&#x60;). | [optional] [readonly] 
**ReferenceText** | Pointer to **string** | The English string text associated with the interface element. | [optional] [readonly] 
**TranslatedText** | **string** | The translated string text associated with the interface element. | 

## Methods

### NewLocaleTranslation

`func NewLocaleTranslation(key string, translatedText string, ) *LocaleTranslation`

NewLocaleTranslation instantiates a new LocaleTranslation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocaleTranslationWithDefaults

`func NewLocaleTranslationWithDefaults() *LocaleTranslation`

NewLocaleTranslationWithDefaults instantiates a new LocaleTranslation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LocaleTranslation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocaleTranslation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocaleTranslation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LocaleTranslation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *LocaleTranslation) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *LocaleTranslation) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *LocaleTranslation) SetKey(v string)`

SetKey sets Key field to given value.


### GetShortKey

`func (o *LocaleTranslation) GetShortKey() string`

GetShortKey returns the ShortKey field if non-nil, zero value otherwise.

### GetShortKeyOk

`func (o *LocaleTranslation) GetShortKeyOk() (*string, bool)`

GetShortKeyOk returns a tuple with the ShortKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortKey

`func (o *LocaleTranslation) SetShortKey(v string)`

SetShortKey sets ShortKey field to given value.

### HasShortKey

`func (o *LocaleTranslation) HasShortKey() bool`

HasShortKey returns a boolean if a field has been set.

### GetReferenceText

`func (o *LocaleTranslation) GetReferenceText() string`

GetReferenceText returns the ReferenceText field if non-nil, zero value otherwise.

### GetReferenceTextOk

`func (o *LocaleTranslation) GetReferenceTextOk() (*string, bool)`

GetReferenceTextOk returns a tuple with the ReferenceText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceText

`func (o *LocaleTranslation) SetReferenceText(v string)`

SetReferenceText sets ReferenceText field to given value.

### HasReferenceText

`func (o *LocaleTranslation) HasReferenceText() bool`

HasReferenceText returns a boolean if a field has been set.

### GetTranslatedText

`func (o *LocaleTranslation) GetTranslatedText() string`

GetTranslatedText returns the TranslatedText field if non-nil, zero value otherwise.

### GetTranslatedTextOk

`func (o *LocaleTranslation) GetTranslatedTextOk() (*string, bool)`

GetTranslatedTextOk returns a tuple with the TranslatedText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedText

`func (o *LocaleTranslation) SetTranslatedText(v string)`

SetTranslatedText sets TranslatedText field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


