# ApplicationCorsSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Behavior** | [**EnumApplicationCorsSettingsBehavior**](EnumApplicationCorsSettingsBehavior.md) |  | 
**Origins** | Pointer to **[]string** | Must be non-empty when &#x60;corsSettings.behavior&#x60; is &#x60;ALLOW_SPECIFIC_ORIGINS&#x60; and must be omitted or empty when &#x60;corsSettings.behavior&#x60; is &#x60;ALLOW_NO_ORIGINS&#x60;.  Limited to 20 values.  Values are the origins from which CORS requests to the Authorization and Authentication APIs are allowed.  Each value is an &#x60;http&#x60; or &#x60;https&#x60; URL without a path.  The host may be a domain name (including &#x60;localhost&#x60;), or an IPv4 or IPv6 address.  Subdomains may use the wildcard (*) to match any string. | [optional] 

## Methods

### NewApplicationCorsSettings

`func NewApplicationCorsSettings(behavior EnumApplicationCorsSettingsBehavior, ) *ApplicationCorsSettings`

NewApplicationCorsSettings instantiates a new ApplicationCorsSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCorsSettingsWithDefaults

`func NewApplicationCorsSettingsWithDefaults() *ApplicationCorsSettings`

NewApplicationCorsSettingsWithDefaults instantiates a new ApplicationCorsSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBehavior

`func (o *ApplicationCorsSettings) GetBehavior() EnumApplicationCorsSettingsBehavior`

GetBehavior returns the Behavior field if non-nil, zero value otherwise.

### GetBehaviorOk

`func (o *ApplicationCorsSettings) GetBehaviorOk() (*EnumApplicationCorsSettingsBehavior, bool)`

GetBehaviorOk returns a tuple with the Behavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehavior

`func (o *ApplicationCorsSettings) SetBehavior(v EnumApplicationCorsSettingsBehavior)`

SetBehavior sets Behavior field to given value.


### GetOrigins

`func (o *ApplicationCorsSettings) GetOrigins() []string`

GetOrigins returns the Origins field if non-nil, zero value otherwise.

### GetOriginsOk

`func (o *ApplicationCorsSettings) GetOriginsOk() (*[]string, bool)`

GetOriginsOk returns a tuple with the Origins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigins

`func (o *ApplicationCorsSettings) SetOrigins(v []string)`

SetOrigins sets Origins field to given value.

### HasOrigins

`func (o *ApplicationCorsSettings) HasOrigins() bool`

HasOrigins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


