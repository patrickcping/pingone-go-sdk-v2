# EntityArrayEmbedded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerifyPolicies** | Pointer to [**[]VerifyPolicy**](VerifyPolicy.md) |  | [optional] 
**VoicePhrases** | Pointer to [**[]VoicePhrase**](VoicePhrase.md) |  | [optional] 
**Contents** | Pointer to [**[]VoicePhraseContents**](VoicePhraseContents.md) |  | [optional] 

## Methods

### NewEntityArrayEmbedded

`func NewEntityArrayEmbedded() *EntityArrayEmbedded`

NewEntityArrayEmbedded instantiates a new EntityArrayEmbedded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityArrayEmbeddedWithDefaults

`func NewEntityArrayEmbeddedWithDefaults() *EntityArrayEmbedded`

NewEntityArrayEmbeddedWithDefaults instantiates a new EntityArrayEmbedded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerifyPolicies

`func (o *EntityArrayEmbedded) GetVerifyPolicies() []VerifyPolicy`

GetVerifyPolicies returns the VerifyPolicies field if non-nil, zero value otherwise.

### GetVerifyPoliciesOk

`func (o *EntityArrayEmbedded) GetVerifyPoliciesOk() (*[]VerifyPolicy, bool)`

GetVerifyPoliciesOk returns a tuple with the VerifyPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyPolicies

`func (o *EntityArrayEmbedded) SetVerifyPolicies(v []VerifyPolicy)`

SetVerifyPolicies sets VerifyPolicies field to given value.

### HasVerifyPolicies

`func (o *EntityArrayEmbedded) HasVerifyPolicies() bool`

HasVerifyPolicies returns a boolean if a field has been set.

### GetVoicePhrases

`func (o *EntityArrayEmbedded) GetVoicePhrases() []VoicePhrase`

GetVoicePhrases returns the VoicePhrases field if non-nil, zero value otherwise.

### GetVoicePhrasesOk

`func (o *EntityArrayEmbedded) GetVoicePhrasesOk() (*[]VoicePhrase, bool)`

GetVoicePhrasesOk returns a tuple with the VoicePhrases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoicePhrases

`func (o *EntityArrayEmbedded) SetVoicePhrases(v []VoicePhrase)`

SetVoicePhrases sets VoicePhrases field to given value.

### HasVoicePhrases

`func (o *EntityArrayEmbedded) HasVoicePhrases() bool`

HasVoicePhrases returns a boolean if a field has been set.

### GetContents

`func (o *EntityArrayEmbedded) GetContents() []VoicePhraseContents`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *EntityArrayEmbedded) GetContentsOk() (*[]VoicePhraseContents, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *EntityArrayEmbedded) SetContents(v []VoicePhraseContents)`

SetContents sets Contents field to given value.

### HasContents

`func (o *EntityArrayEmbedded) HasContents() bool`

HasContents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


