# EmailDomainTrustedEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A string that specifies the auto generated ID of the trusted email address. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**EmailDomain** | Pointer to [**EmailDomainTrustedEmailEmailDomain**](EmailDomainTrustedEmailEmailDomain.md) |  | [optional] 
**EmailAddress** | **string** | A string that specifies the trusted email address, for example john.smith@shopco.com. | 
**Status** | Pointer to [**EnumTrustedEmailStatus**](EnumTrustedEmailStatus.md) |  | [optional] 

## Methods

### NewEmailDomainTrustedEmail

`func NewEmailDomainTrustedEmail(emailAddress string, ) *EmailDomainTrustedEmail`

NewEmailDomainTrustedEmail instantiates a new EmailDomainTrustedEmail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailDomainTrustedEmailWithDefaults

`func NewEmailDomainTrustedEmailWithDefaults() *EmailDomainTrustedEmail`

NewEmailDomainTrustedEmailWithDefaults instantiates a new EmailDomainTrustedEmail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailDomainTrustedEmail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailDomainTrustedEmail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailDomainTrustedEmail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmailDomainTrustedEmail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *EmailDomainTrustedEmail) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *EmailDomainTrustedEmail) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *EmailDomainTrustedEmail) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *EmailDomainTrustedEmail) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetEmailDomain

`func (o *EmailDomainTrustedEmail) GetEmailDomain() EmailDomainTrustedEmailEmailDomain`

GetEmailDomain returns the EmailDomain field if non-nil, zero value otherwise.

### GetEmailDomainOk

`func (o *EmailDomainTrustedEmail) GetEmailDomainOk() (*EmailDomainTrustedEmailEmailDomain, bool)`

GetEmailDomainOk returns a tuple with the EmailDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailDomain

`func (o *EmailDomainTrustedEmail) SetEmailDomain(v EmailDomainTrustedEmailEmailDomain)`

SetEmailDomain sets EmailDomain field to given value.

### HasEmailDomain

`func (o *EmailDomainTrustedEmail) HasEmailDomain() bool`

HasEmailDomain returns a boolean if a field has been set.

### GetEmailAddress

`func (o *EmailDomainTrustedEmail) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *EmailDomainTrustedEmail) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *EmailDomainTrustedEmail) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetStatus

`func (o *EmailDomainTrustedEmail) GetStatus() EnumTrustedEmailStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmailDomainTrustedEmail) GetStatusOk() (*EnumTrustedEmailStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmailDomainTrustedEmail) SetStatus(v EnumTrustedEmailStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EmailDomainTrustedEmail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


