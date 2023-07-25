# VoicePhraseContents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**VoicePhrase** | [**VoicePhraseContentsVoicePhrase**](VoicePhraseContentsVoicePhrase.md) |  | 
**Locale** | **string** |  | 
**Content** | **string** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewVoicePhraseContents

`func NewVoicePhraseContents(voicePhrase VoicePhraseContentsVoicePhrase, locale string, content string, ) *VoicePhraseContents`

NewVoicePhraseContents instantiates a new VoicePhraseContents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoicePhraseContentsWithDefaults

`func NewVoicePhraseContentsWithDefaults() *VoicePhraseContents`

NewVoicePhraseContentsWithDefaults instantiates a new VoicePhraseContents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VoicePhraseContents) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VoicePhraseContents) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VoicePhraseContents) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VoicePhraseContents) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *VoicePhraseContents) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *VoicePhraseContents) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *VoicePhraseContents) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *VoicePhraseContents) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetVoicePhrase

`func (o *VoicePhraseContents) GetVoicePhrase() VoicePhraseContentsVoicePhrase`

GetVoicePhrase returns the VoicePhrase field if non-nil, zero value otherwise.

### GetVoicePhraseOk

`func (o *VoicePhraseContents) GetVoicePhraseOk() (*VoicePhraseContentsVoicePhrase, bool)`

GetVoicePhraseOk returns a tuple with the VoicePhrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoicePhrase

`func (o *VoicePhraseContents) SetVoicePhrase(v VoicePhraseContentsVoicePhrase)`

SetVoicePhrase sets VoicePhrase field to given value.


### GetLocale

`func (o *VoicePhraseContents) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *VoicePhraseContents) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *VoicePhraseContents) SetLocale(v string)`

SetLocale sets Locale field to given value.


### GetContent

`func (o *VoicePhraseContents) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *VoicePhraseContents) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *VoicePhraseContents) SetContent(v string)`

SetContent sets Content field to given value.


### GetCreatedAt

`func (o *VoicePhraseContents) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VoicePhraseContents) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VoicePhraseContents) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VoicePhraseContents) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VoicePhraseContents) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VoicePhraseContents) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VoicePhraseContents) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VoicePhraseContents) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


