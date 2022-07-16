# RiskEvaluationEventBrowser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserAgent** | Pointer to **string** |  | [optional] 
**Cookie** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**ColorDepth** | Pointer to **float32** |  | [optional] 
**DeviceMemory** | Pointer to **float32** |  | [optional] 
**HardwareConcurrency** | Pointer to **float32** |  | [optional] 
**ScreenResolution** | Pointer to **[]float32** |  | [optional] 
**AvailableScreenResolution** | Pointer to **[]float32** |  | [optional] 
**TimezoneOffset** | Pointer to **float32** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**SessionStorage** | Pointer to **bool** |  | [optional] 
**LocalStorage** | Pointer to **bool** |  | [optional] 
**IndexedDb** | Pointer to **bool** |  | [optional] 
**AddBehaviour** | Pointer to **map[string]interface{}** |  | [optional] 
**OpenDatabase** | Pointer to **bool** |  | [optional] 
**CpuClass** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Plugins** | Pointer to **[]map[string]interface{}** |  | [optional] 
**WebglVendorAndRenderer** | Pointer to **string** |  | [optional] 
**Webgl** | Pointer to **[]string** |  | [optional] 
**AdBlock** | Pointer to **bool** |  | [optional] 
**HasLiedLanguages** | Pointer to **bool** |  | [optional] 
**HasLiedResolution** | Pointer to **bool** |  | [optional] 
**HasLiedOs** | Pointer to **bool** |  | [optional] 
**HasLiedBrowser** | Pointer to **bool** |  | [optional] 
**TouchSupport** | Pointer to **[]string** |  | [optional] 
**Fonts** | Pointer to **[]string** |  | [optional] 
**Audio** | Pointer to **string** |  | [optional] 

## Methods

### NewRiskEvaluationEventBrowser

`func NewRiskEvaluationEventBrowser() *RiskEvaluationEventBrowser`

NewRiskEvaluationEventBrowser instantiates a new RiskEvaluationEventBrowser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskEvaluationEventBrowserWithDefaults

`func NewRiskEvaluationEventBrowserWithDefaults() *RiskEvaluationEventBrowser`

NewRiskEvaluationEventBrowserWithDefaults instantiates a new RiskEvaluationEventBrowser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserAgent

`func (o *RiskEvaluationEventBrowser) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *RiskEvaluationEventBrowser) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *RiskEvaluationEventBrowser) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *RiskEvaluationEventBrowser) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetCookie

`func (o *RiskEvaluationEventBrowser) GetCookie() string`

GetCookie returns the Cookie field if non-nil, zero value otherwise.

### GetCookieOk

`func (o *RiskEvaluationEventBrowser) GetCookieOk() (*string, bool)`

GetCookieOk returns a tuple with the Cookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookie

`func (o *RiskEvaluationEventBrowser) SetCookie(v string)`

SetCookie sets Cookie field to given value.

### HasCookie

`func (o *RiskEvaluationEventBrowser) HasCookie() bool`

HasCookie returns a boolean if a field has been set.

### GetLanguage

`func (o *RiskEvaluationEventBrowser) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *RiskEvaluationEventBrowser) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *RiskEvaluationEventBrowser) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *RiskEvaluationEventBrowser) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetColorDepth

`func (o *RiskEvaluationEventBrowser) GetColorDepth() float32`

GetColorDepth returns the ColorDepth field if non-nil, zero value otherwise.

### GetColorDepthOk

`func (o *RiskEvaluationEventBrowser) GetColorDepthOk() (*float32, bool)`

GetColorDepthOk returns a tuple with the ColorDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorDepth

`func (o *RiskEvaluationEventBrowser) SetColorDepth(v float32)`

SetColorDepth sets ColorDepth field to given value.

### HasColorDepth

`func (o *RiskEvaluationEventBrowser) HasColorDepth() bool`

HasColorDepth returns a boolean if a field has been set.

### GetDeviceMemory

`func (o *RiskEvaluationEventBrowser) GetDeviceMemory() float32`

GetDeviceMemory returns the DeviceMemory field if non-nil, zero value otherwise.

### GetDeviceMemoryOk

`func (o *RiskEvaluationEventBrowser) GetDeviceMemoryOk() (*float32, bool)`

GetDeviceMemoryOk returns a tuple with the DeviceMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMemory

`func (o *RiskEvaluationEventBrowser) SetDeviceMemory(v float32)`

SetDeviceMemory sets DeviceMemory field to given value.

### HasDeviceMemory

`func (o *RiskEvaluationEventBrowser) HasDeviceMemory() bool`

HasDeviceMemory returns a boolean if a field has been set.

### GetHardwareConcurrency

`func (o *RiskEvaluationEventBrowser) GetHardwareConcurrency() float32`

GetHardwareConcurrency returns the HardwareConcurrency field if non-nil, zero value otherwise.

### GetHardwareConcurrencyOk

`func (o *RiskEvaluationEventBrowser) GetHardwareConcurrencyOk() (*float32, bool)`

GetHardwareConcurrencyOk returns a tuple with the HardwareConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareConcurrency

`func (o *RiskEvaluationEventBrowser) SetHardwareConcurrency(v float32)`

SetHardwareConcurrency sets HardwareConcurrency field to given value.

### HasHardwareConcurrency

`func (o *RiskEvaluationEventBrowser) HasHardwareConcurrency() bool`

HasHardwareConcurrency returns a boolean if a field has been set.

### GetScreenResolution

`func (o *RiskEvaluationEventBrowser) GetScreenResolution() []float32`

GetScreenResolution returns the ScreenResolution field if non-nil, zero value otherwise.

### GetScreenResolutionOk

`func (o *RiskEvaluationEventBrowser) GetScreenResolutionOk() (*[]float32, bool)`

GetScreenResolutionOk returns a tuple with the ScreenResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenResolution

`func (o *RiskEvaluationEventBrowser) SetScreenResolution(v []float32)`

SetScreenResolution sets ScreenResolution field to given value.

### HasScreenResolution

`func (o *RiskEvaluationEventBrowser) HasScreenResolution() bool`

HasScreenResolution returns a boolean if a field has been set.

### GetAvailableScreenResolution

`func (o *RiskEvaluationEventBrowser) GetAvailableScreenResolution() []float32`

GetAvailableScreenResolution returns the AvailableScreenResolution field if non-nil, zero value otherwise.

### GetAvailableScreenResolutionOk

`func (o *RiskEvaluationEventBrowser) GetAvailableScreenResolutionOk() (*[]float32, bool)`

GetAvailableScreenResolutionOk returns a tuple with the AvailableScreenResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableScreenResolution

`func (o *RiskEvaluationEventBrowser) SetAvailableScreenResolution(v []float32)`

SetAvailableScreenResolution sets AvailableScreenResolution field to given value.

### HasAvailableScreenResolution

`func (o *RiskEvaluationEventBrowser) HasAvailableScreenResolution() bool`

HasAvailableScreenResolution returns a boolean if a field has been set.

### GetTimezoneOffset

`func (o *RiskEvaluationEventBrowser) GetTimezoneOffset() float32`

GetTimezoneOffset returns the TimezoneOffset field if non-nil, zero value otherwise.

### GetTimezoneOffsetOk

`func (o *RiskEvaluationEventBrowser) GetTimezoneOffsetOk() (*float32, bool)`

GetTimezoneOffsetOk returns a tuple with the TimezoneOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneOffset

`func (o *RiskEvaluationEventBrowser) SetTimezoneOffset(v float32)`

SetTimezoneOffset sets TimezoneOffset field to given value.

### HasTimezoneOffset

`func (o *RiskEvaluationEventBrowser) HasTimezoneOffset() bool`

HasTimezoneOffset returns a boolean if a field has been set.

### GetTimezone

`func (o *RiskEvaluationEventBrowser) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *RiskEvaluationEventBrowser) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *RiskEvaluationEventBrowser) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *RiskEvaluationEventBrowser) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetSessionStorage

`func (o *RiskEvaluationEventBrowser) GetSessionStorage() bool`

GetSessionStorage returns the SessionStorage field if non-nil, zero value otherwise.

### GetSessionStorageOk

`func (o *RiskEvaluationEventBrowser) GetSessionStorageOk() (*bool, bool)`

GetSessionStorageOk returns a tuple with the SessionStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionStorage

`func (o *RiskEvaluationEventBrowser) SetSessionStorage(v bool)`

SetSessionStorage sets SessionStorage field to given value.

### HasSessionStorage

`func (o *RiskEvaluationEventBrowser) HasSessionStorage() bool`

HasSessionStorage returns a boolean if a field has been set.

### GetLocalStorage

`func (o *RiskEvaluationEventBrowser) GetLocalStorage() bool`

GetLocalStorage returns the LocalStorage field if non-nil, zero value otherwise.

### GetLocalStorageOk

`func (o *RiskEvaluationEventBrowser) GetLocalStorageOk() (*bool, bool)`

GetLocalStorageOk returns a tuple with the LocalStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStorage

`func (o *RiskEvaluationEventBrowser) SetLocalStorage(v bool)`

SetLocalStorage sets LocalStorage field to given value.

### HasLocalStorage

`func (o *RiskEvaluationEventBrowser) HasLocalStorage() bool`

HasLocalStorage returns a boolean if a field has been set.

### GetIndexedDb

`func (o *RiskEvaluationEventBrowser) GetIndexedDb() bool`

GetIndexedDb returns the IndexedDb field if non-nil, zero value otherwise.

### GetIndexedDbOk

`func (o *RiskEvaluationEventBrowser) GetIndexedDbOk() (*bool, bool)`

GetIndexedDbOk returns a tuple with the IndexedDb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexedDb

`func (o *RiskEvaluationEventBrowser) SetIndexedDb(v bool)`

SetIndexedDb sets IndexedDb field to given value.

### HasIndexedDb

`func (o *RiskEvaluationEventBrowser) HasIndexedDb() bool`

HasIndexedDb returns a boolean if a field has been set.

### GetAddBehaviour

`func (o *RiskEvaluationEventBrowser) GetAddBehaviour() map[string]interface{}`

GetAddBehaviour returns the AddBehaviour field if non-nil, zero value otherwise.

### GetAddBehaviourOk

`func (o *RiskEvaluationEventBrowser) GetAddBehaviourOk() (*map[string]interface{}, bool)`

GetAddBehaviourOk returns a tuple with the AddBehaviour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddBehaviour

`func (o *RiskEvaluationEventBrowser) SetAddBehaviour(v map[string]interface{})`

SetAddBehaviour sets AddBehaviour field to given value.

### HasAddBehaviour

`func (o *RiskEvaluationEventBrowser) HasAddBehaviour() bool`

HasAddBehaviour returns a boolean if a field has been set.

### GetOpenDatabase

`func (o *RiskEvaluationEventBrowser) GetOpenDatabase() bool`

GetOpenDatabase returns the OpenDatabase field if non-nil, zero value otherwise.

### GetOpenDatabaseOk

`func (o *RiskEvaluationEventBrowser) GetOpenDatabaseOk() (*bool, bool)`

GetOpenDatabaseOk returns a tuple with the OpenDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenDatabase

`func (o *RiskEvaluationEventBrowser) SetOpenDatabase(v bool)`

SetOpenDatabase sets OpenDatabase field to given value.

### HasOpenDatabase

`func (o *RiskEvaluationEventBrowser) HasOpenDatabase() bool`

HasOpenDatabase returns a boolean if a field has been set.

### GetCpuClass

`func (o *RiskEvaluationEventBrowser) GetCpuClass() string`

GetCpuClass returns the CpuClass field if non-nil, zero value otherwise.

### GetCpuClassOk

`func (o *RiskEvaluationEventBrowser) GetCpuClassOk() (*string, bool)`

GetCpuClassOk returns a tuple with the CpuClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuClass

`func (o *RiskEvaluationEventBrowser) SetCpuClass(v string)`

SetCpuClass sets CpuClass field to given value.

### HasCpuClass

`func (o *RiskEvaluationEventBrowser) HasCpuClass() bool`

HasCpuClass returns a boolean if a field has been set.

### GetPlatform

`func (o *RiskEvaluationEventBrowser) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *RiskEvaluationEventBrowser) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *RiskEvaluationEventBrowser) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *RiskEvaluationEventBrowser) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPlugins

`func (o *RiskEvaluationEventBrowser) GetPlugins() []map[string]interface{}`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *RiskEvaluationEventBrowser) GetPluginsOk() (*[]map[string]interface{}, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *RiskEvaluationEventBrowser) SetPlugins(v []map[string]interface{})`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *RiskEvaluationEventBrowser) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.

### GetWebglVendorAndRenderer

`func (o *RiskEvaluationEventBrowser) GetWebglVendorAndRenderer() string`

GetWebglVendorAndRenderer returns the WebglVendorAndRenderer field if non-nil, zero value otherwise.

### GetWebglVendorAndRendererOk

`func (o *RiskEvaluationEventBrowser) GetWebglVendorAndRendererOk() (*string, bool)`

GetWebglVendorAndRendererOk returns a tuple with the WebglVendorAndRenderer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebglVendorAndRenderer

`func (o *RiskEvaluationEventBrowser) SetWebglVendorAndRenderer(v string)`

SetWebglVendorAndRenderer sets WebglVendorAndRenderer field to given value.

### HasWebglVendorAndRenderer

`func (o *RiskEvaluationEventBrowser) HasWebglVendorAndRenderer() bool`

HasWebglVendorAndRenderer returns a boolean if a field has been set.

### GetWebgl

`func (o *RiskEvaluationEventBrowser) GetWebgl() []string`

GetWebgl returns the Webgl field if non-nil, zero value otherwise.

### GetWebglOk

`func (o *RiskEvaluationEventBrowser) GetWebglOk() (*[]string, bool)`

GetWebglOk returns a tuple with the Webgl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebgl

`func (o *RiskEvaluationEventBrowser) SetWebgl(v []string)`

SetWebgl sets Webgl field to given value.

### HasWebgl

`func (o *RiskEvaluationEventBrowser) HasWebgl() bool`

HasWebgl returns a boolean if a field has been set.

### GetAdBlock

`func (o *RiskEvaluationEventBrowser) GetAdBlock() bool`

GetAdBlock returns the AdBlock field if non-nil, zero value otherwise.

### GetAdBlockOk

`func (o *RiskEvaluationEventBrowser) GetAdBlockOk() (*bool, bool)`

GetAdBlockOk returns a tuple with the AdBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdBlock

`func (o *RiskEvaluationEventBrowser) SetAdBlock(v bool)`

SetAdBlock sets AdBlock field to given value.

### HasAdBlock

`func (o *RiskEvaluationEventBrowser) HasAdBlock() bool`

HasAdBlock returns a boolean if a field has been set.

### GetHasLiedLanguages

`func (o *RiskEvaluationEventBrowser) GetHasLiedLanguages() bool`

GetHasLiedLanguages returns the HasLiedLanguages field if non-nil, zero value otherwise.

### GetHasLiedLanguagesOk

`func (o *RiskEvaluationEventBrowser) GetHasLiedLanguagesOk() (*bool, bool)`

GetHasLiedLanguagesOk returns a tuple with the HasLiedLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLiedLanguages

`func (o *RiskEvaluationEventBrowser) SetHasLiedLanguages(v bool)`

SetHasLiedLanguages sets HasLiedLanguages field to given value.

### HasHasLiedLanguages

`func (o *RiskEvaluationEventBrowser) HasHasLiedLanguages() bool`

HasHasLiedLanguages returns a boolean if a field has been set.

### GetHasLiedResolution

`func (o *RiskEvaluationEventBrowser) GetHasLiedResolution() bool`

GetHasLiedResolution returns the HasLiedResolution field if non-nil, zero value otherwise.

### GetHasLiedResolutionOk

`func (o *RiskEvaluationEventBrowser) GetHasLiedResolutionOk() (*bool, bool)`

GetHasLiedResolutionOk returns a tuple with the HasLiedResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLiedResolution

`func (o *RiskEvaluationEventBrowser) SetHasLiedResolution(v bool)`

SetHasLiedResolution sets HasLiedResolution field to given value.

### HasHasLiedResolution

`func (o *RiskEvaluationEventBrowser) HasHasLiedResolution() bool`

HasHasLiedResolution returns a boolean if a field has been set.

### GetHasLiedOs

`func (o *RiskEvaluationEventBrowser) GetHasLiedOs() bool`

GetHasLiedOs returns the HasLiedOs field if non-nil, zero value otherwise.

### GetHasLiedOsOk

`func (o *RiskEvaluationEventBrowser) GetHasLiedOsOk() (*bool, bool)`

GetHasLiedOsOk returns a tuple with the HasLiedOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLiedOs

`func (o *RiskEvaluationEventBrowser) SetHasLiedOs(v bool)`

SetHasLiedOs sets HasLiedOs field to given value.

### HasHasLiedOs

`func (o *RiskEvaluationEventBrowser) HasHasLiedOs() bool`

HasHasLiedOs returns a boolean if a field has been set.

### GetHasLiedBrowser

`func (o *RiskEvaluationEventBrowser) GetHasLiedBrowser() bool`

GetHasLiedBrowser returns the HasLiedBrowser field if non-nil, zero value otherwise.

### GetHasLiedBrowserOk

`func (o *RiskEvaluationEventBrowser) GetHasLiedBrowserOk() (*bool, bool)`

GetHasLiedBrowserOk returns a tuple with the HasLiedBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLiedBrowser

`func (o *RiskEvaluationEventBrowser) SetHasLiedBrowser(v bool)`

SetHasLiedBrowser sets HasLiedBrowser field to given value.

### HasHasLiedBrowser

`func (o *RiskEvaluationEventBrowser) HasHasLiedBrowser() bool`

HasHasLiedBrowser returns a boolean if a field has been set.

### GetTouchSupport

`func (o *RiskEvaluationEventBrowser) GetTouchSupport() []string`

GetTouchSupport returns the TouchSupport field if non-nil, zero value otherwise.

### GetTouchSupportOk

`func (o *RiskEvaluationEventBrowser) GetTouchSupportOk() (*[]string, bool)`

GetTouchSupportOk returns a tuple with the TouchSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTouchSupport

`func (o *RiskEvaluationEventBrowser) SetTouchSupport(v []string)`

SetTouchSupport sets TouchSupport field to given value.

### HasTouchSupport

`func (o *RiskEvaluationEventBrowser) HasTouchSupport() bool`

HasTouchSupport returns a boolean if a field has been set.

### GetFonts

`func (o *RiskEvaluationEventBrowser) GetFonts() []string`

GetFonts returns the Fonts field if non-nil, zero value otherwise.

### GetFontsOk

`func (o *RiskEvaluationEventBrowser) GetFontsOk() (*[]string, bool)`

GetFontsOk returns a tuple with the Fonts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFonts

`func (o *RiskEvaluationEventBrowser) SetFonts(v []string)`

SetFonts sets Fonts field to given value.

### HasFonts

`func (o *RiskEvaluationEventBrowser) HasFonts() bool`

HasFonts returns a boolean if a field has been set.

### GetAudio

`func (o *RiskEvaluationEventBrowser) GetAudio() string`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *RiskEvaluationEventBrowser) GetAudioOk() (*string, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *RiskEvaluationEventBrowser) SetAudio(v string)`

SetAudio sets Audio field to given value.

### HasAudio

`func (o *RiskEvaluationEventBrowser) HasAudio() bool`

HasAudio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


