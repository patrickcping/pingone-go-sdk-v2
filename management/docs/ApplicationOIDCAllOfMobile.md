# ApplicationOIDCAllOfMobile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleId** | Pointer to **string** | A string that specifies the bundle associated with the application, for push notifications in native apps. The value of the bundleId property is unique per environment, and once defined, is immutable.  this setting overrides the top-level bundleId field | [optional] 
**PackageName** | Pointer to **string** | A string that specifies the package name associated with the application, for push notifications in native apps. The value of the mobile.packageName property is unique per environment, and once defined, is immutable.  this setting overrides the top-level packageName field. | [optional] 
**IntegrityDetection** | Pointer to [**ApplicationOIDCAllOfMobileIntegrityDetection**](ApplicationOIDCAllOfMobileIntegrityDetection.md) |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


