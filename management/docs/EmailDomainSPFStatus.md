# EmailDomainSPFStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A string that specifies the type of DNS record, with the value \&quot;TXT\&quot;. | [optional] [readonly] 
**Status** | Pointer to [**EnumEmailDomainStatus**](EnumEmailDomainStatus.md) |  | [optional] 
**Key** | Pointer to **string** | Record name. | [optional] 
**Value** | Pointer to **string** | Record value. | [optional] 

## Methods

### NewEmailDomainSPFStatus

`func NewEmailDomainSPFStatus() *EmailDomainSPFStatus`

NewEmailDomainSPFStatus instantiates a new EmailDomainSPFStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailDomainSPFStatusWithDefaults

`func NewEmailDomainSPFStatusWithDefaults() *EmailDomainSPFStatus`

NewEmailDomainSPFStatusWithDefaults instantiates a new EmailDomainSPFStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EmailDomainSPFStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EmailDomainSPFStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EmailDomainSPFStatus) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EmailDomainSPFStatus) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *EmailDomainSPFStatus) GetStatus() EnumEmailDomainStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmailDomainSPFStatus) GetStatusOk() (*EnumEmailDomainStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmailDomainSPFStatus) SetStatus(v EnumEmailDomainStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EmailDomainSPFStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetKey

`func (o *EmailDomainSPFStatus) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EmailDomainSPFStatus) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EmailDomainSPFStatus) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *EmailDomainSPFStatus) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *EmailDomainSPFStatus) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EmailDomainSPFStatus) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EmailDomainSPFStatus) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *EmailDomainSPFStatus) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


