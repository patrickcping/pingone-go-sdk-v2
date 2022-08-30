# UserAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **string** | A string that specifies the country name component. When specified, the value must be in ISO 3166-1 &#x60;alpha-2&#x60; code format [ISO3166]. For example, the country codes for the United States and Sweden are &#x60;US&#x60; and \&quot;SE\&quot;, respectively. Valid characters consist of two upper-case letters (regex &#x60;[A-Z]{2}&#x60;). | [optional] 
**Locality** | Pointer to **string** | A string that specifies the city or locality component of the address. The string can contain any letters, numbers, combining characters, math and currency symbols, dingbats and drawing characters, and invisible whitespace (regex &#x60;^[\\p{L}\\p{M}\\p{Zs}\\p{S}\\p{N}\\p{P}]*$&#x60;). It can have a length of no more than 256 characters (min/max&#x3D;1/256). | [optional] 
**PostalCode** | Pointer to **string** | A string that specifies the zip code or postal code component of the address. The string can contain any letters, numbers, combining characters, math and currency symbols, dingbats and drawing characters, and invisible whitespace (regex &#x60;^[\\p{L}\\p{M}\\p{Zs}\\p{S}\\p{N}\\p{P}]*$&#x60;). It can have a length of no more than 40 characters (min/max&#x3D;1/40). | [optional] 
**Region** | Pointer to **string** | A string that specifies the state or region component of the address. The string can contain any letters, numbers, combining characters, math and currency symbols, dingbats and drawing characters, and invisible whitespace (regex &#x60;^[\\p{L}\\p{M}\\p{Zs}\\p{S}\\p{N}\\p{P}]*$&#x60;). It can have a length of no more than 256 characters (min/max&#x3D;1/256). | [optional] 
**StreetAddress** | Pointer to **string** | A string that specifies the full street address component, which may include house number, street name, P.O. box, and multi-line extended street address information. This attribute may contain newlines (regex &#x60;^[\\p{L}\\p{M}\\p{N}\\p{Zs}\\p{P}\\n\\r]*$&#x60;). It can have a length of no more than 256 characters (min/max&#x3D;1/256). | [optional] 

## Methods

### NewUserAddress

`func NewUserAddress() *UserAddress`

NewUserAddress instantiates a new UserAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAddressWithDefaults

`func NewUserAddressWithDefaults() *UserAddress`

NewUserAddressWithDefaults instantiates a new UserAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *UserAddress) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UserAddress) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UserAddress) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UserAddress) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetLocality

`func (o *UserAddress) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *UserAddress) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *UserAddress) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *UserAddress) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetPostalCode

`func (o *UserAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *UserAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *UserAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *UserAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetRegion

`func (o *UserAddress) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *UserAddress) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *UserAddress) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *UserAddress) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStreetAddress

`func (o *UserAddress) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *UserAddress) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *UserAddress) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *UserAddress) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


