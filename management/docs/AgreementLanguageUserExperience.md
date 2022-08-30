# AgreementLanguageUserExperience

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptCheckboxText** | Pointer to **string** | A string that specifies the text next to the \&quot;accept\&quot; checkbox in the end user interface. Accepted character are unicode letters, combining marks, numeric characters, whitespace, and punctuation characters (regex &#x60;^[\\p{L}\\p{M}\\p{N}\\p{Zs}\\p{P}]+$&#x60;). | [optional] 
**ContinueButtonText** | Pointer to **string** | A string that specifies the text of the \&quot;continue\&quot; button in the end user interface. Accepted character are unicode letters, combining marks, numeric characters, whitespace, and punctuation characters (regex &#x60;^[\\p{L}\\p{M}\\p{N}\\p{Zs}\\p{P}]+$&#x60;). | [optional] 
**DeclineButtonText** | Pointer to **string** | A string that specifies the text of the \&quot;decline\&quot; button in the end user interface. Accepted character are unicode letters, combining marks, numeric characters, whitespace, and punctuation characters (regex &#x60;^[\\p{L}\\p{M}\\p{N}\\p{Zs}\\p{P}]+$&#x60;). | [optional] 

## Methods

### NewAgreementLanguageUserExperience

`func NewAgreementLanguageUserExperience() *AgreementLanguageUserExperience`

NewAgreementLanguageUserExperience instantiates a new AgreementLanguageUserExperience object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgreementLanguageUserExperienceWithDefaults

`func NewAgreementLanguageUserExperienceWithDefaults() *AgreementLanguageUserExperience`

NewAgreementLanguageUserExperienceWithDefaults instantiates a new AgreementLanguageUserExperience object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptCheckboxText

`func (o *AgreementLanguageUserExperience) GetAcceptCheckboxText() string`

GetAcceptCheckboxText returns the AcceptCheckboxText field if non-nil, zero value otherwise.

### GetAcceptCheckboxTextOk

`func (o *AgreementLanguageUserExperience) GetAcceptCheckboxTextOk() (*string, bool)`

GetAcceptCheckboxTextOk returns a tuple with the AcceptCheckboxText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptCheckboxText

`func (o *AgreementLanguageUserExperience) SetAcceptCheckboxText(v string)`

SetAcceptCheckboxText sets AcceptCheckboxText field to given value.

### HasAcceptCheckboxText

`func (o *AgreementLanguageUserExperience) HasAcceptCheckboxText() bool`

HasAcceptCheckboxText returns a boolean if a field has been set.

### GetContinueButtonText

`func (o *AgreementLanguageUserExperience) GetContinueButtonText() string`

GetContinueButtonText returns the ContinueButtonText field if non-nil, zero value otherwise.

### GetContinueButtonTextOk

`func (o *AgreementLanguageUserExperience) GetContinueButtonTextOk() (*string, bool)`

GetContinueButtonTextOk returns a tuple with the ContinueButtonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueButtonText

`func (o *AgreementLanguageUserExperience) SetContinueButtonText(v string)`

SetContinueButtonText sets ContinueButtonText field to given value.

### HasContinueButtonText

`func (o *AgreementLanguageUserExperience) HasContinueButtonText() bool`

HasContinueButtonText returns a boolean if a field has been set.

### GetDeclineButtonText

`func (o *AgreementLanguageUserExperience) GetDeclineButtonText() string`

GetDeclineButtonText returns the DeclineButtonText field if non-nil, zero value otherwise.

### GetDeclineButtonTextOk

`func (o *AgreementLanguageUserExperience) GetDeclineButtonTextOk() (*string, bool)`

GetDeclineButtonTextOk returns a tuple with the DeclineButtonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclineButtonText

`func (o *AgreementLanguageUserExperience) SetDeclineButtonText(v string)`

SetDeclineButtonText sets DeclineButtonText field to given value.

### HasDeclineButtonText

`func (o *AgreementLanguageUserExperience) HasDeclineButtonText() bool`

HasDeclineButtonText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


