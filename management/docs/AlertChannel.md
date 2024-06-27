# AlertChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Id** | Pointer to **string** | Unique ID of the alert channel. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**AlertName** | Pointer to **string** | The name to assign to the alert channel. | [optional] 
**ChannelType** | [**EnumAlertChannelType**](EnumAlertChannelType.md) |  | 
**Addresses** | **[]string** | The email addresses to send the alert to. | 
**IncludeSeverities** | Pointer to [**[]EnumAlertChannelSeverity**](EnumAlertChannelSeverity.md) | Filters alerts by severity. If empty, all severities are included. Possible values are &#x60;INFO&#x60;, &#x60;WARNING&#x60;, and &#x60;ERROR&#x60;&#x60;. | [optional] 
**IncludeAlertTypes** | Pointer to [**[]EnumAlertChannelAlertType**](EnumAlertChannelAlertType.md) | Filters alerts by alert type. If empty, all alert types are included. Possible values are CERTIFICATE_EXPIRED, CERTIFICATE_EXPIRING, KEY_PAIR_EXPIRED, GATEWAY_VERSION_DEPRECATED, KEY_PAIR_EXPIRING, and GATEWAY_VERSION_DEPRECATING. | [optional] 
**ExcludeAlertTypes** | Pointer to [**[]EnumAlertChannelAlertType**](EnumAlertChannelAlertType.md) | Administrators will not be emailed alerts of these types. If empty, no alert types are excluded. Possible values are CERTIFICATE_EXPIRED, CERTIFICATE_EXPIRING, KEY_PAIR_EXPIRED, GATEWAY_VERSION_DEPRECATED, KEY_PAIR_EXPIRING, and GATEWAY_VERSION_DEPRECATING. | [optional] 

## Methods

### NewAlertChannel

`func NewAlertChannel(channelType EnumAlertChannelType, addresses []string, ) *AlertChannel`

NewAlertChannel instantiates a new AlertChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertChannelWithDefaults

`func NewAlertChannelWithDefaults() *AlertChannel`

NewAlertChannelWithDefaults instantiates a new AlertChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AlertChannel) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AlertChannel) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AlertChannel) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AlertChannel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *AlertChannel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertChannel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertChannel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AlertChannel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *AlertChannel) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AlertChannel) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AlertChannel) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *AlertChannel) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetAlertName

`func (o *AlertChannel) GetAlertName() string`

GetAlertName returns the AlertName field if non-nil, zero value otherwise.

### GetAlertNameOk

`func (o *AlertChannel) GetAlertNameOk() (*string, bool)`

GetAlertNameOk returns a tuple with the AlertName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertName

`func (o *AlertChannel) SetAlertName(v string)`

SetAlertName sets AlertName field to given value.

### HasAlertName

`func (o *AlertChannel) HasAlertName() bool`

HasAlertName returns a boolean if a field has been set.

### GetChannelType

`func (o *AlertChannel) GetChannelType() EnumAlertChannelType`

GetChannelType returns the ChannelType field if non-nil, zero value otherwise.

### GetChannelTypeOk

`func (o *AlertChannel) GetChannelTypeOk() (*EnumAlertChannelType, bool)`

GetChannelTypeOk returns a tuple with the ChannelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelType

`func (o *AlertChannel) SetChannelType(v EnumAlertChannelType)`

SetChannelType sets ChannelType field to given value.


### GetAddresses

`func (o *AlertChannel) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *AlertChannel) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *AlertChannel) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetIncludeSeverities

`func (o *AlertChannel) GetIncludeSeverities() []EnumAlertChannelSeverity`

GetIncludeSeverities returns the IncludeSeverities field if non-nil, zero value otherwise.

### GetIncludeSeveritiesOk

`func (o *AlertChannel) GetIncludeSeveritiesOk() (*[]EnumAlertChannelSeverity, bool)`

GetIncludeSeveritiesOk returns a tuple with the IncludeSeverities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSeverities

`func (o *AlertChannel) SetIncludeSeverities(v []EnumAlertChannelSeverity)`

SetIncludeSeverities sets IncludeSeverities field to given value.

### HasIncludeSeverities

`func (o *AlertChannel) HasIncludeSeverities() bool`

HasIncludeSeverities returns a boolean if a field has been set.

### GetIncludeAlertTypes

`func (o *AlertChannel) GetIncludeAlertTypes() []EnumAlertChannelAlertType`

GetIncludeAlertTypes returns the IncludeAlertTypes field if non-nil, zero value otherwise.

### GetIncludeAlertTypesOk

`func (o *AlertChannel) GetIncludeAlertTypesOk() (*[]EnumAlertChannelAlertType, bool)`

GetIncludeAlertTypesOk returns a tuple with the IncludeAlertTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAlertTypes

`func (o *AlertChannel) SetIncludeAlertTypes(v []EnumAlertChannelAlertType)`

SetIncludeAlertTypes sets IncludeAlertTypes field to given value.

### HasIncludeAlertTypes

`func (o *AlertChannel) HasIncludeAlertTypes() bool`

HasIncludeAlertTypes returns a boolean if a field has been set.

### GetExcludeAlertTypes

`func (o *AlertChannel) GetExcludeAlertTypes() []EnumAlertChannelAlertType`

GetExcludeAlertTypes returns the ExcludeAlertTypes field if non-nil, zero value otherwise.

### GetExcludeAlertTypesOk

`func (o *AlertChannel) GetExcludeAlertTypesOk() (*[]EnumAlertChannelAlertType, bool)`

GetExcludeAlertTypesOk returns a tuple with the ExcludeAlertTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeAlertTypes

`func (o *AlertChannel) SetExcludeAlertTypes(v []EnumAlertChannelAlertType)`

SetExcludeAlertTypes sets ExcludeAlertTypes field to given value.

### HasExcludeAlertTypes

`func (o *AlertChannel) HasExcludeAlertTypes() bool`

HasExcludeAlertTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


