# UserName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Family** | Pointer to **string** | A string that specifies the family name of the user, or Last in most Western languages (for example, ‘Jensen’ given the full name ‘Ms. Barbara J Jensen, III’). This may be explicitly set to null when updating a name to unset it. Valid characters consist of any Unicode letter, mark (for example, accent, umlaut), space, dot, apostrophe, or hyphen (regex &#x60;^[\\p{L}\\p{M}\\p{N}&#39; .-]*$&#x60;). It can have a length of no more than 256 characters (min/max&#x3D;1/256). | [optional] 
**Formatted** | Pointer to **string** | A string that specifies the fully formatted name of the user (for example ‘Ms. Barbara J Jensen, III’). This can be explicitly set to null when updating a name to unset it. Valid characters consist of any Unicode letter, mark (for example, accent, umlaut), space, dot, apostrophe, or hyphen (regex &#x60;^[\\p{L}\\p{M}\\p{N}&#39; .-]*$&#x60;). It can have a length of no more than 256 characters (min/max&#x3D;1/256). | [optional] 
**Given** | Pointer to **string** | A string that specifies the given name of the user, or First in most Western languages (for example, ‘Barbara’ given the full name ‘Ms. Barbara J Jensen, III’). This may be explicitly set to null when updating a name to unset it. The string can contain any letters, numbers, combining characters, math and currency symbols, dingbats and drawing characters, and invisible whitespace (regex &#x60;^[\\p{L}\\p{M}\\p{Zs}\\p{S}\\p{N}\\p{P}]*$&#x60;). It can have a length of no more than 256 characters (min/max&#x3D;1/256). | [optional] 
**HonorificPrefix** | Pointer to **string** | A string that specifies the honorific prefix(es) of the user, or title in most Western languages (for example, ‘Ms.’ given the full name ‘Ms. Barbara Jane Jensen, III’). This can be explicitly set to null when updating a name to unset it. | [optional] 
**HonorificSuffix** | Pointer to **string** | A string that specifies the honorific suffix(es) of the user, or suffix in most Western languages (for example, ‘III’ given the full name ‘Ms. Barbara Jane Jensen, III’). This can be explicitly set to null when updating a name to unset it. | [optional] 
**Middle** | Pointer to **string** | A string that specifies the the middle name(s) of the user (for exmple, ‘Jane’ given the full name ‘Ms. Barbara Jane Jensen, III’). This can be explicitly set to null when updating a name to unset it. The string can contain any letters, numbers, combining characters, math and currency symbols, dingbats and drawing characters, and invisible whitespace (regex &#x60;^[\\p{L}\\p{M}\\p{Zs}\\p{S}\\p{N}\\p{P}]*$&#x60;). It can have a length of no more than 256 characters (min/max&#x3D;1/256). | [optional] 

## Methods

### NewUserName

`func NewUserName() *UserName`

NewUserName instantiates a new UserName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserNameWithDefaults

`func NewUserNameWithDefaults() *UserName`

NewUserNameWithDefaults instantiates a new UserName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFamily

`func (o *UserName) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *UserName) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *UserName) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *UserName) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetFormatted

`func (o *UserName) GetFormatted() string`

GetFormatted returns the Formatted field if non-nil, zero value otherwise.

### GetFormattedOk

`func (o *UserName) GetFormattedOk() (*string, bool)`

GetFormattedOk returns a tuple with the Formatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatted

`func (o *UserName) SetFormatted(v string)`

SetFormatted sets Formatted field to given value.

### HasFormatted

`func (o *UserName) HasFormatted() bool`

HasFormatted returns a boolean if a field has been set.

### GetGiven

`func (o *UserName) GetGiven() string`

GetGiven returns the Given field if non-nil, zero value otherwise.

### GetGivenOk

`func (o *UserName) GetGivenOk() (*string, bool)`

GetGivenOk returns a tuple with the Given field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiven

`func (o *UserName) SetGiven(v string)`

SetGiven sets Given field to given value.

### HasGiven

`func (o *UserName) HasGiven() bool`

HasGiven returns a boolean if a field has been set.

### GetHonorificPrefix

`func (o *UserName) GetHonorificPrefix() string`

GetHonorificPrefix returns the HonorificPrefix field if non-nil, zero value otherwise.

### GetHonorificPrefixOk

`func (o *UserName) GetHonorificPrefixOk() (*string, bool)`

GetHonorificPrefixOk returns a tuple with the HonorificPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHonorificPrefix

`func (o *UserName) SetHonorificPrefix(v string)`

SetHonorificPrefix sets HonorificPrefix field to given value.

### HasHonorificPrefix

`func (o *UserName) HasHonorificPrefix() bool`

HasHonorificPrefix returns a boolean if a field has been set.

### GetHonorificSuffix

`func (o *UserName) GetHonorificSuffix() string`

GetHonorificSuffix returns the HonorificSuffix field if non-nil, zero value otherwise.

### GetHonorificSuffixOk

`func (o *UserName) GetHonorificSuffixOk() (*string, bool)`

GetHonorificSuffixOk returns a tuple with the HonorificSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHonorificSuffix

`func (o *UserName) SetHonorificSuffix(v string)`

SetHonorificSuffix sets HonorificSuffix field to given value.

### HasHonorificSuffix

`func (o *UserName) HasHonorificSuffix() bool`

HasHonorificSuffix returns a boolean if a field has been set.

### GetMiddle

`func (o *UserName) GetMiddle() string`

GetMiddle returns the Middle field if non-nil, zero value otherwise.

### GetMiddleOk

`func (o *UserName) GetMiddleOk() (*string, bool)`

GetMiddleOk returns a tuple with the Middle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddle

`func (o *UserName) SetMiddle(v string)`

SetMiddle sets Middle field to given value.

### HasMiddle

`func (o *UserName) HasMiddle() bool`

HasMiddle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


