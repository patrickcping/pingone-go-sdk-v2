# EmailDomainOwnershipStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A string that specifies the type of DNS record, with the value \&quot;TXT\&quot;. | [optional] [readonly] 
**Regions** | Pointer to [**[]EmailDomainOwnershipStatusRegionsInner**](EmailDomainOwnershipStatusRegionsInner.md) | The regions collection specifies the properties for the 4 AWS SES regions that are used for sending email for the environment. The regions are determined by the geography where this environment was provisioned (North America, Canada, Europe &amp; Asia-Pacific). | [optional] [readonly] 

## Methods

### NewEmailDomainOwnershipStatus

`func NewEmailDomainOwnershipStatus() *EmailDomainOwnershipStatus`

NewEmailDomainOwnershipStatus instantiates a new EmailDomainOwnershipStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailDomainOwnershipStatusWithDefaults

`func NewEmailDomainOwnershipStatusWithDefaults() *EmailDomainOwnershipStatus`

NewEmailDomainOwnershipStatusWithDefaults instantiates a new EmailDomainOwnershipStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EmailDomainOwnershipStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EmailDomainOwnershipStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EmailDomainOwnershipStatus) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EmailDomainOwnershipStatus) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRegions

`func (o *EmailDomainOwnershipStatus) GetRegions() []EmailDomainOwnershipStatusRegionsInner`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *EmailDomainOwnershipStatus) GetRegionsOk() (*[]EmailDomainOwnershipStatusRegionsInner, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *EmailDomainOwnershipStatus) SetRegions(v []EmailDomainOwnershipStatusRegionsInner)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *EmailDomainOwnershipStatus) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


