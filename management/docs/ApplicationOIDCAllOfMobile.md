# ApplicationOIDCAllOfMobile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleId** | Pointer to **string** | A string that specifies the bundle associated with the application, for push notifications in native apps. The value of the bundleId property is unique per environment, and once defined, is immutable.  this setting overrides the top-level bundleId field | [optional] 
**PackageName** | Pointer to **string** | A string that specifies the package name associated with the application, for push notifications in native apps. The value of the mobile.packageName property is unique per environment, and once defined, is immutable.  this setting overrides the top-level packageName field. | [optional] 
**PasscodeRefreshDuration** | Pointer to [**ApplicationOIDCAllOfMobilePasscodeRefreshDuration**](ApplicationOIDCAllOfMobilePasscodeRefreshDuration.md) |  | [optional] 
**IntegrityDetection** | Pointer to [**ApplicationOIDCAllOfMobileIntegrityDetection**](ApplicationOIDCAllOfMobileIntegrityDetection.md) |  | [optional] 
**UriPrefix** | Pointer to **string** | A string that specifies a URI prefix that enables direct triggering of the mobile application when scanning a QR code. The URI prefix can be set to a universal link with a valid value (which can be a URL address that starts with &#x60;HTTP://&#x60; or &#x60;HTTPS://&#x60;, such as &#x60;https://www.acme.com&#x60;), or an app schema, which is just a string and requires no special validation. | [optional] 

## Methods

### NewApplicationOIDCAllOfMobile

`func NewApplicationOIDCAllOfMobile() *ApplicationOIDCAllOfMobile`

NewApplicationOIDCAllOfMobile instantiates a new ApplicationOIDCAllOfMobile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationOIDCAllOfMobileWithDefaults

`func NewApplicationOIDCAllOfMobileWithDefaults() *ApplicationOIDCAllOfMobile`

NewApplicationOIDCAllOfMobileWithDefaults instantiates a new ApplicationOIDCAllOfMobile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleId

`func (o *ApplicationOIDCAllOfMobile) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *ApplicationOIDCAllOfMobile) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *ApplicationOIDCAllOfMobile) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.

### HasBundleId

`func (o *ApplicationOIDCAllOfMobile) HasBundleId() bool`

HasBundleId returns a boolean if a field has been set.

### GetPackageName

`func (o *ApplicationOIDCAllOfMobile) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *ApplicationOIDCAllOfMobile) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *ApplicationOIDCAllOfMobile) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *ApplicationOIDCAllOfMobile) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### GetPasscodeRefreshDuration

`func (o *ApplicationOIDCAllOfMobile) GetPasscodeRefreshDuration() ApplicationOIDCAllOfMobilePasscodeRefreshDuration`

GetPasscodeRefreshDuration returns the PasscodeRefreshDuration field if non-nil, zero value otherwise.

### GetPasscodeRefreshDurationOk

`func (o *ApplicationOIDCAllOfMobile) GetPasscodeRefreshDurationOk() (*ApplicationOIDCAllOfMobilePasscodeRefreshDuration, bool)`

GetPasscodeRefreshDurationOk returns a tuple with the PasscodeRefreshDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscodeRefreshDuration

`func (o *ApplicationOIDCAllOfMobile) SetPasscodeRefreshDuration(v ApplicationOIDCAllOfMobilePasscodeRefreshDuration)`

SetPasscodeRefreshDuration sets PasscodeRefreshDuration field to given value.

### HasPasscodeRefreshDuration

`func (o *ApplicationOIDCAllOfMobile) HasPasscodeRefreshDuration() bool`

HasPasscodeRefreshDuration returns a boolean if a field has been set.

### GetIntegrityDetection

`func (o *ApplicationOIDCAllOfMobile) GetIntegrityDetection() ApplicationOIDCAllOfMobileIntegrityDetection`

GetIntegrityDetection returns the IntegrityDetection field if non-nil, zero value otherwise.

### GetIntegrityDetectionOk

`func (o *ApplicationOIDCAllOfMobile) GetIntegrityDetectionOk() (*ApplicationOIDCAllOfMobileIntegrityDetection, bool)`

GetIntegrityDetectionOk returns a tuple with the IntegrityDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrityDetection

`func (o *ApplicationOIDCAllOfMobile) SetIntegrityDetection(v ApplicationOIDCAllOfMobileIntegrityDetection)`

SetIntegrityDetection sets IntegrityDetection field to given value.

### HasIntegrityDetection

`func (o *ApplicationOIDCAllOfMobile) HasIntegrityDetection() bool`

HasIntegrityDetection returns a boolean if a field has been set.

### GetUriPrefix

`func (o *ApplicationOIDCAllOfMobile) GetUriPrefix() string`

GetUriPrefix returns the UriPrefix field if non-nil, zero value otherwise.

### GetUriPrefixOk

`func (o *ApplicationOIDCAllOfMobile) GetUriPrefixOk() (*string, bool)`

GetUriPrefixOk returns a tuple with the UriPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUriPrefix

`func (o *ApplicationOIDCAllOfMobile) SetUriPrefix(v string)`

SetUriPrefix sets UriPrefix field to given value.

### HasUriPrefix

`func (o *ApplicationOIDCAllOfMobile) HasUriPrefix() bool`

HasUriPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


