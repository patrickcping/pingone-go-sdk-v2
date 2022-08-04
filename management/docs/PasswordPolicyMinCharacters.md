# PasswordPolicyMinCharacters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ABCDEFGHIJKLMNOPQRSTUVWXYZ** | Pointer to **int32** | Count of alphabetical uppercase characters (&#x60;ABCDEFGHIJKLMNOPQRSTUVWXYZ&#x60;) that should feature in the user&#39;s password. | [optional] 
**Abcdefghijklmnopqrstuvwxyz** | Pointer to **int32** | Count of alphabetical uppercase characters (&#x60;abcdefghijklmnopqrstuvwxyz&#x60;) that should feature in the user&#39;s password. | [optional] 
**Var0123456789** | Pointer to **int32** | Count of numeric characters (&#x60;0123456789&#x60;) that should feature in the user&#39;s password. | [optional] 
**SpecialChar** | Pointer to **int32** | Count of special characters (&#x60;~!@#$%^&amp;*()-_&#x3D;+[]{}\\\\|;:,.&lt;&gt;/?&#x60;) that should feature in the user&#39;s password. | [optional] 

## Methods

### NewPasswordPolicyMinCharacters

`func NewPasswordPolicyMinCharacters() *PasswordPolicyMinCharacters`

NewPasswordPolicyMinCharacters instantiates a new PasswordPolicyMinCharacters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyMinCharactersWithDefaults

`func NewPasswordPolicyMinCharactersWithDefaults() *PasswordPolicyMinCharacters`

NewPasswordPolicyMinCharactersWithDefaults instantiates a new PasswordPolicyMinCharacters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetABCDEFGHIJKLMNOPQRSTUVWXYZ

`func (o *PasswordPolicyMinCharacters) GetABCDEFGHIJKLMNOPQRSTUVWXYZ() int32`

GetABCDEFGHIJKLMNOPQRSTUVWXYZ returns the ABCDEFGHIJKLMNOPQRSTUVWXYZ field if non-nil, zero value otherwise.

### GetABCDEFGHIJKLMNOPQRSTUVWXYZOk

`func (o *PasswordPolicyMinCharacters) GetABCDEFGHIJKLMNOPQRSTUVWXYZOk() (*int32, bool)`

GetABCDEFGHIJKLMNOPQRSTUVWXYZOk returns a tuple with the ABCDEFGHIJKLMNOPQRSTUVWXYZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetABCDEFGHIJKLMNOPQRSTUVWXYZ

`func (o *PasswordPolicyMinCharacters) SetABCDEFGHIJKLMNOPQRSTUVWXYZ(v int32)`

SetABCDEFGHIJKLMNOPQRSTUVWXYZ sets ABCDEFGHIJKLMNOPQRSTUVWXYZ field to given value.

### HasABCDEFGHIJKLMNOPQRSTUVWXYZ

`func (o *PasswordPolicyMinCharacters) HasABCDEFGHIJKLMNOPQRSTUVWXYZ() bool`

HasABCDEFGHIJKLMNOPQRSTUVWXYZ returns a boolean if a field has been set.

### GetAbcdefghijklmnopqrstuvwxyz

`func (o *PasswordPolicyMinCharacters) GetAbcdefghijklmnopqrstuvwxyz() int32`

GetAbcdefghijklmnopqrstuvwxyz returns the Abcdefghijklmnopqrstuvwxyz field if non-nil, zero value otherwise.

### GetAbcdefghijklmnopqrstuvwxyzOk

`func (o *PasswordPolicyMinCharacters) GetAbcdefghijklmnopqrstuvwxyzOk() (*int32, bool)`

GetAbcdefghijklmnopqrstuvwxyzOk returns a tuple with the Abcdefghijklmnopqrstuvwxyz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbcdefghijklmnopqrstuvwxyz

`func (o *PasswordPolicyMinCharacters) SetAbcdefghijklmnopqrstuvwxyz(v int32)`

SetAbcdefghijklmnopqrstuvwxyz sets Abcdefghijklmnopqrstuvwxyz field to given value.

### HasAbcdefghijklmnopqrstuvwxyz

`func (o *PasswordPolicyMinCharacters) HasAbcdefghijklmnopqrstuvwxyz() bool`

HasAbcdefghijklmnopqrstuvwxyz returns a boolean if a field has been set.

### GetVar0123456789

`func (o *PasswordPolicyMinCharacters) GetVar0123456789() int32`

GetVar0123456789 returns the Var0123456789 field if non-nil, zero value otherwise.

### GetVar0123456789Ok

`func (o *PasswordPolicyMinCharacters) GetVar0123456789Ok() (*int32, bool)`

GetVar0123456789Ok returns a tuple with the Var0123456789 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar0123456789

`func (o *PasswordPolicyMinCharacters) SetVar0123456789(v int32)`

SetVar0123456789 sets Var0123456789 field to given value.

### HasVar0123456789

`func (o *PasswordPolicyMinCharacters) HasVar0123456789() bool`

HasVar0123456789 returns a boolean if a field has been set.

### GetSpecialChar

`func (o *PasswordPolicyMinCharacters) GetSpecialChar() int32`

GetSpecialChar returns the SpecialChar field if non-nil, zero value otherwise.

### GetSpecialCharOk

`func (o *PasswordPolicyMinCharacters) GetSpecialCharOk() (*int32, bool)`

GetSpecialCharOk returns a tuple with the SpecialChar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialChar

`func (o *PasswordPolicyMinCharacters) SetSpecialChar(v int32)`

SetSpecialChar sets SpecialChar field to given value.

### HasSpecialChar

`func (o *PasswordPolicyMinCharacters) HasSpecialChar() bool`

HasSpecialChar returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


