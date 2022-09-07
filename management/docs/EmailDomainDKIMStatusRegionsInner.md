# EmailDomainDKIMStatusRegionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the region. | [optional] 
**Status** | Pointer to [**EnumEmailDomainStatus**](EnumEmailDomainStatus.md) |  | [optional] 
**Tokens** | Pointer to [**[]EmailDomainDKIMStatusRegionsInnerTokensInner**](EmailDomainDKIMStatusRegionsInnerTokensInner.md) | A collection of key and value pairs. | [optional] 

## Methods

### NewEmailDomainDKIMStatusRegionsInner

`func NewEmailDomainDKIMStatusRegionsInner() *EmailDomainDKIMStatusRegionsInner`

NewEmailDomainDKIMStatusRegionsInner instantiates a new EmailDomainDKIMStatusRegionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailDomainDKIMStatusRegionsInnerWithDefaults

`func NewEmailDomainDKIMStatusRegionsInnerWithDefaults() *EmailDomainDKIMStatusRegionsInner`

NewEmailDomainDKIMStatusRegionsInnerWithDefaults instantiates a new EmailDomainDKIMStatusRegionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EmailDomainDKIMStatusRegionsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailDomainDKIMStatusRegionsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailDomainDKIMStatusRegionsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EmailDomainDKIMStatusRegionsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *EmailDomainDKIMStatusRegionsInner) GetStatus() EnumEmailDomainStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmailDomainDKIMStatusRegionsInner) GetStatusOk() (*EnumEmailDomainStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmailDomainDKIMStatusRegionsInner) SetStatus(v EnumEmailDomainStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EmailDomainDKIMStatusRegionsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTokens

`func (o *EmailDomainDKIMStatusRegionsInner) GetTokens() []EmailDomainDKIMStatusRegionsInnerTokensInner`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *EmailDomainDKIMStatusRegionsInner) GetTokensOk() (*[]EmailDomainDKIMStatusRegionsInnerTokensInner, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *EmailDomainDKIMStatusRegionsInner) SetTokens(v []EmailDomainDKIMStatusRegionsInnerTokensInner)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *EmailDomainDKIMStatusRegionsInner) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


