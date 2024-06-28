# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Account** | Pointer to [**UserAccount**](UserAccount.md) |  | [optional] 
**Address** | Pointer to [**UserAddress**](UserAddress.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time the resource was created. | [optional] [readonly] 
**Email** | **string** | A string that specifies the user’s email address, which must be provided and valid. For more information about email address formatting, see section 3.4 of RFC 2822, Internet Message Format. | 
**EmailVerified** | Pointer to **bool** | Whether the user’s email is verified. An email address can be verified during account verification. If the email address used to request the verification code is the same as the user’s email at verification time (and the verification code is valid), then the email is verified. The value of this property can be set on user import. | [optional] [readonly] 
**Enabled** | Pointer to **bool** | A read-only boolean attribute that specifies whether the user is enabled. This attribute is set to ‘true’ by default when the user is created. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**ExternalId** | Pointer to **string** | A string that specifies an identifier for the user resource as defined by the provisioning client. This is optional. This may be explicitly set to null when updating a user to unset it. The externalId attribute simplifies the correlation of the user in PingOne with the user’s account in another system of record. The platform does not use this attribute directly in any way, but it is used by Ping Identity’s Data Sync product. It can have a length of no more than 1024 characters (min/max&#x3D;1/1024). | [optional] 
**Id** | Pointer to **string** | A string that specifies the user resource’s unique identifier. | [optional] [readonly] 
**IdentityProvider** | Pointer to [**UserIdentityProvider**](UserIdentityProvider.md) |  | [optional] 
**LastSignOn** | Pointer to [**UserLastSignOn**](UserLastSignOn.md) |  | [optional] 
**Lifecycle** | Pointer to [**UserLifecycle**](UserLifecycle.md) |  | [optional] 
**Locale** | Pointer to **string** | A string that specifies the user’s default location, which is optional. This may be explicitly set to null when updating a user to unset it. This is used for purposes of localizing such items as currency, date time format, or numerical representations. If provided, it must be a valid language tag as defined in RFC 5646. The following are example tags fr, en-US, es-419, az-Arab, man-Nkoo-GN. The string can contain any letters, numbers, combining characters, math and currency symbols, dingbats and drawing characters, and invisible whitespace (regex &#x60;^[\\p{L}\\p{M}\\p{Zs}\\p{S}\\p{N}\\p{P}]*$&#x60;). It can have a length of no more than 256 characters (min/max&#x3D;1/256). | [optional] 
**MemberOfGroupIDs** | Pointer to **[]string** | A read-only array of IDs for the groups the user is a member of. This property is returned for GET /environments/{environmentID}/users/{userID} when include&#x3D;memberOfGroupIDs is appended to the request. This property is not returned with a list of users. | [optional] [readonly] 
**MemberOfGroupNames** | Pointer to **[]string** | A read-only array of names for the groups the user is a member of. This property is returned for GET /environments/{environmentID}/users/{userID} when include&#x3D;memberOfGroupNames is appended to the request. This property is not returned with a list of users. | [optional] [readonly] 
**MfaEnabled** | Pointer to **bool** | A boolean attribute that specifies whether multi-factor authentication is enabled. This attribute is set to false by default when the user is created. You can set mfaEnabled to true with POST CREATE User, POST CREATE User (Import), or PUT UPDATE User MFA Enabled. You cannot update mfaEnabled with PUT UPDATE User or PATCH UPDATE User. | [optional] 
**MobilePhone** | Pointer to **string** | A string that specifies the user’s native phone number, which is optional. This might also match the primaryPhone attribute. This may be explicitly set to null when updating a user to unset it. Valid phone numbers must have at least one number and a maximum character length of 32. | [optional] 
**Name** | Pointer to [**UserName**](UserName.md) |  | [optional] 
**Nickname** | Pointer to **string** | A string that specifies the user’s nickname, which is optional. This can be explicitly set to null when updating a user to unset it. The string can contain any letters, numbers, combining characters, math and currency symbols, dingbats and drawing characters, and invisible whitespace (regex &#x60;^[\\p{L}\\p{M}\\p{Zs}\\p{S}\\p{N}\\p{P}]*$&#x60;). It can have a length of no more than 256 characters (min/max&#x3D;1/256). | [optional] 
**Password** | Pointer to [**UserPassword**](UserPassword.md) |  | [optional] 
**Photo** | Pointer to [**UserPhoto**](UserPhoto.md) |  | [optional] 
**Population** | Pointer to [**UserPopulation**](UserPopulation.md) |  | [optional] 
**PreferredLanguage** | Pointer to **string** | A string that specifies the user’s preferred written or spoken languages, which are optional. This may be explicitly set to null when updating a user to unset it. If provided, the format of the value must be a valid language range and the same as the HTTP Accept-Language header field (not including Accept-Language:) and is specified in Section 5.3.5 of RFC 7231. For example en-US, en-gb;q&#x3D;0.8, en;q&#x3D;0.7. | [optional] 
**PrimaryPhone** | Pointer to **string** | A string that specifies the user’s primary phone number, which is optional. This might also match the mobilePhone attribute. This may be explicitly set to null when updating a user to unset it. Valid phone numbers must have at least one number and a maximum character length of 32. | [optional] 
**Timezone** | Pointer to **string** | A string that specifies the user’s time zone, which is optional. This can be explicitly set to null when updating a user to unset it. If provided, it must conform with the IANA Time Zone database format [RFC6557], also known as the &#x60;Olson&#x60; time zone database format [Olson-TZ] for example, &#x60;America/Los_Angeles&#x60; (regex &#x60;^\\w+/\\w+$&#x60;). | [optional] 
**Title** | Pointer to **string** | A string that specifies the user’s title, which is optional, such as \&quot;Vice President\&quot;. This can be explicitly set to null when updating a user to unset it. The string can contain any letters, numbers, combining characters, math and currency symbols, dingbats and drawing characters, and invisible whitespace (regex &#x60;^[\\p{L}\\p{M}\\p{Zs}\\p{S}\\p{N}\\p{P}]*$&#x60;). It can have a length of no more than 256 characters (min/max&#x3D;1/256). | [optional] 
**Type** | Pointer to **string** | A string that specifies the user’s type, which is optional. This can be explicitly set to null when updating a user to unset it. This attribute is organization-specific and has no special meaning within the PingOne platform. It could have values of \&quot;Contractor\&quot;, \&quot;Employee\&quot;, \&quot;Intern\&quot;, \&quot;Temp\&quot;, \&quot;External\&quot;, and &#x60;Unknown&#x60;. The string can contain any letters, numbers, combining characters, math and currency symbols, dingbats and drawing characters, and invisible whitespace (regex &#x60;^[\\p{L}\\p{M}\\p{Zs}\\p{S}\\p{N}\\p{P}]*$&#x60;). It can have a length of no more than 256 characters (min/max&#x3D;1/256). | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The time the resource was last updated. | [optional] [readonly] 
**Username** | **string** | A string that specifies the user name, which must be provided and must be unique within an environment. The username must either be a well-formed email address or a string. The string can contain any letters, numbers, combining characters, math and currency symbols, dingbats and drawing characters, and invisible whitespace (regex &#x60;^[\\p{L}\\p{M}\\p{Zs}\\p{S}\\p{N}\\p{P}]*$&#x60;). It can have a length of no more than 128 characters (min/max&#x3D;1/128). | 
**VerifyStatus** | Pointer to [**EnumUserVerifyStatus**](EnumUserVerifyStatus.md) |  | [optional] 

## Methods

### NewUser

`func NewUser(email string, username string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *User) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *User) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *User) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *User) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAccount

`func (o *User) GetAccount() UserAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *User) GetAccountOk() (*UserAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *User) SetAccount(v UserAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *User) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAddress

`func (o *User) GetAddress() UserAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *User) GetAddressOk() (*UserAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *User) SetAddress(v UserAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *User) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCreatedAt

`func (o *User) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *User) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *User) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *User) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEmailVerified

`func (o *User) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *User) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *User) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *User) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### GetEnabled

`func (o *User) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *User) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *User) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *User) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEnvironment

`func (o *User) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *User) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *User) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *User) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetExternalId

`func (o *User) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *User) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *User) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *User) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *User) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentityProvider

`func (o *User) GetIdentityProvider() UserIdentityProvider`

GetIdentityProvider returns the IdentityProvider field if non-nil, zero value otherwise.

### GetIdentityProviderOk

`func (o *User) GetIdentityProviderOk() (*UserIdentityProvider, bool)`

GetIdentityProviderOk returns a tuple with the IdentityProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProvider

`func (o *User) SetIdentityProvider(v UserIdentityProvider)`

SetIdentityProvider sets IdentityProvider field to given value.

### HasIdentityProvider

`func (o *User) HasIdentityProvider() bool`

HasIdentityProvider returns a boolean if a field has been set.

### GetLastSignOn

`func (o *User) GetLastSignOn() UserLastSignOn`

GetLastSignOn returns the LastSignOn field if non-nil, zero value otherwise.

### GetLastSignOnOk

`func (o *User) GetLastSignOnOk() (*UserLastSignOn, bool)`

GetLastSignOnOk returns a tuple with the LastSignOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSignOn

`func (o *User) SetLastSignOn(v UserLastSignOn)`

SetLastSignOn sets LastSignOn field to given value.

### HasLastSignOn

`func (o *User) HasLastSignOn() bool`

HasLastSignOn returns a boolean if a field has been set.

### GetLifecycle

`func (o *User) GetLifecycle() UserLifecycle`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *User) GetLifecycleOk() (*UserLifecycle, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *User) SetLifecycle(v UserLifecycle)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *User) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetLocale

`func (o *User) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *User) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *User) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *User) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetMemberOfGroupIDs

`func (o *User) GetMemberOfGroupIDs() []string`

GetMemberOfGroupIDs returns the MemberOfGroupIDs field if non-nil, zero value otherwise.

### GetMemberOfGroupIDsOk

`func (o *User) GetMemberOfGroupIDsOk() (*[]string, bool)`

GetMemberOfGroupIDsOk returns a tuple with the MemberOfGroupIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOfGroupIDs

`func (o *User) SetMemberOfGroupIDs(v []string)`

SetMemberOfGroupIDs sets MemberOfGroupIDs field to given value.

### HasMemberOfGroupIDs

`func (o *User) HasMemberOfGroupIDs() bool`

HasMemberOfGroupIDs returns a boolean if a field has been set.

### GetMemberOfGroupNames

`func (o *User) GetMemberOfGroupNames() []string`

GetMemberOfGroupNames returns the MemberOfGroupNames field if non-nil, zero value otherwise.

### GetMemberOfGroupNamesOk

`func (o *User) GetMemberOfGroupNamesOk() (*[]string, bool)`

GetMemberOfGroupNamesOk returns a tuple with the MemberOfGroupNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOfGroupNames

`func (o *User) SetMemberOfGroupNames(v []string)`

SetMemberOfGroupNames sets MemberOfGroupNames field to given value.

### HasMemberOfGroupNames

`func (o *User) HasMemberOfGroupNames() bool`

HasMemberOfGroupNames returns a boolean if a field has been set.

### GetMfaEnabled

`func (o *User) GetMfaEnabled() bool`

GetMfaEnabled returns the MfaEnabled field if non-nil, zero value otherwise.

### GetMfaEnabledOk

`func (o *User) GetMfaEnabledOk() (*bool, bool)`

GetMfaEnabledOk returns a tuple with the MfaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaEnabled

`func (o *User) SetMfaEnabled(v bool)`

SetMfaEnabled sets MfaEnabled field to given value.

### HasMfaEnabled

`func (o *User) HasMfaEnabled() bool`

HasMfaEnabled returns a boolean if a field has been set.

### GetMobilePhone

`func (o *User) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *User) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *User) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *User) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetName

`func (o *User) GetName() UserName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *User) GetNameOk() (*UserName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *User) SetName(v UserName)`

SetName sets Name field to given value.

### HasName

`func (o *User) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNickname

`func (o *User) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *User) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *User) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *User) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetPassword

`func (o *User) GetPassword() UserPassword`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *User) GetPasswordOk() (*UserPassword, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *User) SetPassword(v UserPassword)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *User) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPhoto

`func (o *User) GetPhoto() UserPhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *User) GetPhotoOk() (*UserPhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *User) SetPhoto(v UserPhoto)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *User) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### GetPopulation

`func (o *User) GetPopulation() UserPopulation`

GetPopulation returns the Population field if non-nil, zero value otherwise.

### GetPopulationOk

`func (o *User) GetPopulationOk() (*UserPopulation, bool)`

GetPopulationOk returns a tuple with the Population field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulation

`func (o *User) SetPopulation(v UserPopulation)`

SetPopulation sets Population field to given value.

### HasPopulation

`func (o *User) HasPopulation() bool`

HasPopulation returns a boolean if a field has been set.

### GetPreferredLanguage

`func (o *User) GetPreferredLanguage() string`

GetPreferredLanguage returns the PreferredLanguage field if non-nil, zero value otherwise.

### GetPreferredLanguageOk

`func (o *User) GetPreferredLanguageOk() (*string, bool)`

GetPreferredLanguageOk returns a tuple with the PreferredLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLanguage

`func (o *User) SetPreferredLanguage(v string)`

SetPreferredLanguage sets PreferredLanguage field to given value.

### HasPreferredLanguage

`func (o *User) HasPreferredLanguage() bool`

HasPreferredLanguage returns a boolean if a field has been set.

### GetPrimaryPhone

`func (o *User) GetPrimaryPhone() string`

GetPrimaryPhone returns the PrimaryPhone field if non-nil, zero value otherwise.

### GetPrimaryPhoneOk

`func (o *User) GetPrimaryPhoneOk() (*string, bool)`

GetPrimaryPhoneOk returns a tuple with the PrimaryPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryPhone

`func (o *User) SetPrimaryPhone(v string)`

SetPrimaryPhone sets PrimaryPhone field to given value.

### HasPrimaryPhone

`func (o *User) HasPrimaryPhone() bool`

HasPrimaryPhone returns a boolean if a field has been set.

### GetTimezone

`func (o *User) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *User) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *User) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *User) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetTitle

`func (o *User) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *User) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *User) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *User) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *User) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *User) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *User) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *User) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *User) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *User) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *User) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *User) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetVerifyStatus

`func (o *User) GetVerifyStatus() EnumUserVerifyStatus`

GetVerifyStatus returns the VerifyStatus field if non-nil, zero value otherwise.

### GetVerifyStatusOk

`func (o *User) GetVerifyStatusOk() (*EnumUserVerifyStatus, bool)`

GetVerifyStatusOk returns a tuple with the VerifyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyStatus

`func (o *User) SetVerifyStatus(v EnumUserVerifyStatus)`

SetVerifyStatus sets VerifyStatus field to given value.

### HasVerifyStatus

`func (o *User) HasVerifyStatus() bool`

HasVerifyStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


