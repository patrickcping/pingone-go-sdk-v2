# ApplicationOIDCAllOfMobileIntegrityDetection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludedPlatforms** | Pointer to [**[]EnumMobileIntegrityDetectionPlatform**](EnumMobileIntegrityDetectionPlatform.md) | You can enable device integrity checking separately for Android and iOS by setting &#x60;mobile.integrityDetection.mode&#x60; to &#x60;ENABLED&#x60; and then using &#x60;mobile.integrityDetection.excludedPlatforms&#x60; to specify the OS where you do not want to use device integrity checking. The values to use are &#x60;GOOGLE&#x60; and &#x60;IOS&#x60; (all upper case). Note that this is implemented as an array even though currently you can only include a single value.  If &#x60;GOOGLE&#x60; is not included as a value, &#x60;googlePlay&#x60; is required to be set. | [optional] 
**Mode** | Pointer to [**EnumEnabledStatus**](EnumEnabledStatus.md) |  | [optional] 
**CacheDuration** | Pointer to [**ApplicationOIDCAllOfMobileIntegrityDetectionCacheDuration**](ApplicationOIDCAllOfMobileIntegrityDetectionCacheDuration.md) |  | [optional] 
**GooglePlay** | Pointer to [**ApplicationOIDCAllOfMobileIntegrityDetectionGooglePlay**](ApplicationOIDCAllOfMobileIntegrityDetectionGooglePlay.md) |  | [optional] 

## Methods

### NewApplicationOIDCAllOfMobileIntegrityDetection

`func NewApplicationOIDCAllOfMobileIntegrityDetection() *ApplicationOIDCAllOfMobileIntegrityDetection`

NewApplicationOIDCAllOfMobileIntegrityDetection instantiates a new ApplicationOIDCAllOfMobileIntegrityDetection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationOIDCAllOfMobileIntegrityDetectionWithDefaults

`func NewApplicationOIDCAllOfMobileIntegrityDetectionWithDefaults() *ApplicationOIDCAllOfMobileIntegrityDetection`

NewApplicationOIDCAllOfMobileIntegrityDetectionWithDefaults instantiates a new ApplicationOIDCAllOfMobileIntegrityDetection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludedPlatforms

`func (o *ApplicationOIDCAllOfMobileIntegrityDetection) GetExcludedPlatforms() []EnumMobileIntegrityDetectionPlatform`

GetExcludedPlatforms returns the ExcludedPlatforms field if non-nil, zero value otherwise.

### GetExcludedPlatformsOk

`func (o *ApplicationOIDCAllOfMobileIntegrityDetection) GetExcludedPlatformsOk() (*[]EnumMobileIntegrityDetectionPlatform, bool)`

GetExcludedPlatformsOk returns a tuple with the ExcludedPlatforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedPlatforms

`func (o *ApplicationOIDCAllOfMobileIntegrityDetection) SetExcludedPlatforms(v []EnumMobileIntegrityDetectionPlatform)`

SetExcludedPlatforms sets ExcludedPlatforms field to given value.

### HasExcludedPlatforms

`func (o *ApplicationOIDCAllOfMobileIntegrityDetection) HasExcludedPlatforms() bool`

HasExcludedPlatforms returns a boolean if a field has been set.

### GetMode

`func (o *ApplicationOIDCAllOfMobileIntegrityDetection) GetMode() EnumEnabledStatus`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ApplicationOIDCAllOfMobileIntegrityDetection) GetModeOk() (*EnumEnabledStatus, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ApplicationOIDCAllOfMobileIntegrityDetection) SetMode(v EnumEnabledStatus)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ApplicationOIDCAllOfMobileIntegrityDetection) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetCacheDuration

`func (o *ApplicationOIDCAllOfMobileIntegrityDetection) GetCacheDuration() ApplicationOIDCAllOfMobileIntegrityDetectionCacheDuration`

GetCacheDuration returns the CacheDuration field if non-nil, zero value otherwise.

### GetCacheDurationOk

`func (o *ApplicationOIDCAllOfMobileIntegrityDetection) GetCacheDurationOk() (*ApplicationOIDCAllOfMobileIntegrityDetectionCacheDuration, bool)`

GetCacheDurationOk returns a tuple with the CacheDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheDuration

`func (o *ApplicationOIDCAllOfMobileIntegrityDetection) SetCacheDuration(v ApplicationOIDCAllOfMobileIntegrityDetectionCacheDuration)`

SetCacheDuration sets CacheDuration field to given value.

### HasCacheDuration

`func (o *ApplicationOIDCAllOfMobileIntegrityDetection) HasCacheDuration() bool`

HasCacheDuration returns a boolean if a field has been set.

### GetGooglePlay

`func (o *ApplicationOIDCAllOfMobileIntegrityDetection) GetGooglePlay() ApplicationOIDCAllOfMobileIntegrityDetectionGooglePlay`

GetGooglePlay returns the GooglePlay field if non-nil, zero value otherwise.

### GetGooglePlayOk

`func (o *ApplicationOIDCAllOfMobileIntegrityDetection) GetGooglePlayOk() (*ApplicationOIDCAllOfMobileIntegrityDetectionGooglePlay, bool)`

GetGooglePlayOk returns a tuple with the GooglePlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGooglePlay

`func (o *ApplicationOIDCAllOfMobileIntegrityDetection) SetGooglePlay(v ApplicationOIDCAllOfMobileIntegrityDetectionGooglePlay)`

SetGooglePlay sets GooglePlay field to given value.

### HasGooglePlay

`func (o *ApplicationOIDCAllOfMobileIntegrityDetection) HasGooglePlay() bool`

HasGooglePlay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


