# AgreementLanguageRevision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agreement** | Pointer to [**AgreementLanguageAgreement**](AgreementLanguageAgreement.md) |  | [optional] 
**ContentType** | [**EnumAgreementRevisionContentType**](EnumAgreementRevisionContentType.md) |  | 
**EffectiveAt** | **time.Time** | A date that specifies the start date that the revision is presented to users. This property value can be modified only if the current value is a date that has not already passed. The effective date must be unique for each language agreement, and the property value can be the present date or a future date only. | 
**Id** | Pointer to **string** | A read-only string that specifies the revision ID. | [optional] [readonly] 
**Language** | Pointer to [**AgreementLanguageRevisionLanguage**](AgreementLanguageRevisionLanguage.md) |  | [optional] 
**NotValidAfter** | Pointer to **time.Time** | A date that specifies whether the revision is still valid in the context of all revisions for a language. This property is calculated dynamically at read time, taking into consideration the agreement language, the language enabled property, and the agreement enabled property. When a new revision is added, the notValidAfter property values for all other previous revisions might be impacted. For example, if a new revision becomes effective and it forces reconsent, then all older revisions are no longer valid. | [optional] [readonly] 
**RequireReconsent** | **bool** | A boolean that specifies whether the user is required to provide consent to the language revision after it becomes effective. | 
**Text** | **string** | An immutable string that specifies text or HTML for the revision. This attribute is supported in POST requests only. For more information, see contentType. | 

## Methods

### NewAgreementLanguageRevision

`func NewAgreementLanguageRevision(contentType EnumAgreementRevisionContentType, effectiveAt time.Time, requireReconsent bool, text string, ) *AgreementLanguageRevision`

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

`func (o *AgreementLanguageRevision) GetContentType() EnumAgreementRevisionContentType`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *AgreementLanguageRevision) GetContentTypeOk() (*EnumAgreementRevisionContentType, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *AgreementLanguageRevision) SetContentType(v EnumAgreementRevisionContentType)`

SetContentType sets ContentType field to given value.


### GetEffectiveAt

`func (o *AgreementLanguageRevision) GetEffectiveAt() time.Time`

GetEffectiveAt returns the EffectiveAt field if non-nil, zero value otherwise.

### GetEffectiveAtOk

`func (o *AgreementLanguageRevision) GetEffectiveAtOk() (*time.Time, bool)`

GetEffectiveAtOk returns a tuple with the EffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAt

`func (o *AgreementLanguageRevision) SetEffectiveAt(v time.Time)`

SetEffectiveAt sets EffectiveAt field to given value.


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

`func (o *AgreementLanguageRevision) GetNotValidAfter() time.Time`

GetNotValidAfter returns the NotValidAfter field if non-nil, zero value otherwise.

### GetNotValidAfterOk

`func (o *AgreementLanguageRevision) GetNotValidAfterOk() (*time.Time, bool)`

GetNotValidAfterOk returns a tuple with the NotValidAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotValidAfter

`func (o *AgreementLanguageRevision) SetNotValidAfter(v time.Time)`

SetNotValidAfter sets NotValidAfter field to given value.

### HasNotValidAfter

`func (o *AgreementLanguageRevision) HasNotValidAfter() bool`

HasNotValidAfter returns a boolean if a field has been set.

### GetRequireReconsent

`func (o *AgreementLanguageRevision) GetRequireReconsent() bool`

GetRequireReconsent returns the RequireReconsent field if non-nil, zero value otherwise.

### GetRequireReconsentOk

`func (o *AgreementLanguageRevision) GetRequireReconsentOk() (*bool, bool)`

GetRequireReconsentOk returns a tuple with the RequireReconsent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireReconsent

`func (o *AgreementLanguageRevision) SetRequireReconsent(v bool)`

SetRequireReconsent sets RequireReconsent field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


