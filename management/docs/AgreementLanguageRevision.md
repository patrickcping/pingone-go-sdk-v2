# AgreementLanguageRevision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agreement** | Pointer to [**AgreementLanguageAgreement**](AgreementLanguageAgreement.md) |  | [optional] 
**ContentType** | Pointer to **string** | An immutable string that specifies the content type of text. Options are text/html and text/plain, as defined by rfc-6838 and Media Types/text. This attribute is supported in POST requests only. | [optional] 
**EffectiveAt** | Pointer to **string** | A date that specifies the start date that the revision is presented to users. This property value can be modified only if the current value is a date that has not already passed. The effective date must be unique for each language agreement, and the property value can be the present date or a future date only. | [optional] 
**Id** | Pointer to **string** | A read-only string that specifies the revision ID. | [optional] [readonly] 
**Language** | Pointer to [**AgreementLanguageRevisionLanguage**](AgreementLanguageRevisionLanguage.md) |  | [optional] 
**NotValidAfter** | Pointer to **string** | A date that specifies whether the revision is still valid in the context of all revisions for a language. This property is calculated dynamically at read time, taking into consideration the agreement language, the language enabled property, and the agreement enabled property. When a new revision is added, the notValidAfter property values for all other previous revisions might be impacted. For example, if a new revision becomes effective and it forces reconsent, then all older revisions are no longer valid. | [optional] [readonly] 
**RequiresReconsent** | Pointer to **bool** | A boolean that specifies whether the user is required to provide consent to the language revision after it becomes effective. | [optional] 
**Text** | Pointer to **string** | An immutable string that specifies text or HTML for the revision. This attribute is supported in POST requests only. For more information, see contentType. | [optional] 

## Methods

### NewAgreementLanguageRevision

`func NewAgreementLanguageRevision() *AgreementLanguageRevision`

NewAgreementLanguageRevision instantiates a new AgreementLanguageRevision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgreementLanguageRevisionWithDefaults

`func NewAgreementLanguageRevisionWithDefaults() *AgreementLanguageRevision`

NewAgreementLanguageRevisionWithDefaults instantiates a new AgreementLanguageRevision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgreement

`func (o *AgreementLanguageRevision) GetAgreement() AgreementLanguageAgreement`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *AgreementLanguageRevision) GetAgreementOk() (*AgreementLanguageAgreement, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *AgreementLanguageRevision) SetAgreement(v AgreementLanguageAgreement)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *AgreementLanguageRevision) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetContentType

`func (o *AgreementLanguageRevision) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *AgreementLanguageRevision) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *AgreementLanguageRevision) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *AgreementLanguageRevision) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetEffectiveAt

`func (o *AgreementLanguageRevision) GetEffectiveAt() string`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *AgreementLanguageRevision) GetEffectiveAtOk() (*string, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *AgreementLanguageRevision) SetEffectiveAt(v string)`

SetEffectiveAt sets EffectiveAt field to given value.

### HasEffectiveAt

`func (o *AgreementLanguageRevision) HasEffectiveAt() bool`

HasEffectiveAt returns a boolean if a field has been set.

### GetId

`func (o *AgreementLanguageRevision) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgreementLanguageRevision) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgreementLanguageRevision) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgreementLanguageRevision) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLanguage

`func (o *AgreementLanguageRevision) GetLanguage() AgreementLanguageRevisionLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *AgreementLanguageRevision) GetLanguageOk() (*AgreementLanguageRevisionLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *AgreementLanguageRevision) SetLanguage(v AgreementLanguageRevisionLanguage)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *AgreementLanguageRevision) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetNotValidAfter

`func (o *AgreementLanguageRevision) GetNotValidAfter() string`

GetNotValidAfter returns the NotValidAfter field if non-nil, zero value otherwise.

### GetNotValidAfterOk

`func (o *AgreementLanguageRevision) GetNotValidAfterOk() (*string, bool)`

GetNotValidAfterOk returns a tuple with the NotValidAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotValidAfter

`func (o *AgreementLanguageRevision) SetNotValidAfter(v string)`

SetNotValidAfter sets NotValidAfter field to given value.

### HasNotValidAfter

`func (o *AgreementLanguageRevision) HasNotValidAfter() bool`

HasNotValidAfter returns a boolean if a field has been set.

### GetRequiresReconsent

`func (o *AgreementLanguageRevision) GetRequiresReconsent() bool`

GetRequiresReconsent returns the RequiresReconsent field if non-nil, zero value otherwise.

### GetRequiresReconsentOk

`func (o *AgreementLanguageRevision) GetRequiresReconsentOk() (*bool, bool)`

GetRequiresReconsentOk returns a tuple with the RequiresReconsent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresReconsent

`func (o *AgreementLanguageRevision) SetRequiresReconsent(v bool)`

SetRequiresReconsent sets RequiresReconsent field to given value.

### HasRequiresReconsent

`func (o *AgreementLanguageRevision) HasRequiresReconsent() bool`

HasRequiresReconsent returns a boolean if a field has been set.

### GetText

`func (o *AgreementLanguageRevision) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *AgreementLanguageRevision) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *AgreementLanguageRevision) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *AgreementLanguageRevision) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


