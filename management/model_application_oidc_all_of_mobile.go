/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// checks if the ApplicationOIDCAllOfMobile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationOIDCAllOfMobile{}

// ApplicationOIDCAllOfMobile struct for ApplicationOIDCAllOfMobile
type ApplicationOIDCAllOfMobile struct {
	// A string that specifies the bundle associated with the application, for push notifications in native apps. The value of the bundleId property is unique per environment, and once defined, is immutable.  this setting overrides the top-level bundleId field
	BundleId *string `json:"bundleId,omitempty"`
	// A string that specifies the package name associated with the application, for push notifications in native apps. The value of the mobile.packageName property is unique per environment, and once defined, is immutable.  this setting overrides the top-level packageName field.
	PackageName *string `json:"packageName,omitempty"`
	// The unique identifier for the app on the device and in the Huawei Mobile Service AppGallery. The value of the `huaweiAppId` property is unique per environment, and once defined, is immutable. Used only for applications for the Huawei ecosystem.
	HuaweiAppId *string `json:"huaweiAppId,omitempty"`
	// The package name associated with the application, for push notifications in native apps. The value of the `huaweiPackageName` property is unique per environment, and once defined, is immutable. Used only for applications for the Huawei ecosystem.
	HuaweiPackageName *string `json:"huaweiPackageName,omitempty"`
	PasscodeRefreshDuration *ApplicationOIDCAllOfMobilePasscodeRefreshDuration `json:"passcodeRefreshDuration,omitempty"`
	IntegrityDetection *ApplicationOIDCAllOfMobileIntegrityDetection `json:"integrityDetection,omitempty"`
	// A string that specifies a URI prefix that enables direct triggering of the mobile application when scanning a QR code. The URI prefix can be set to a universal link with a valid value (which can be a URL address that starts with `HTTP://` or `HTTPS://`, such as `https://www.acme.com`), or an app schema, which is just a string and requires no special validation.
	UriPrefix *string `json:"uriPrefix,omitempty"`
}

// NewApplicationOIDCAllOfMobile instantiates a new ApplicationOIDCAllOfMobile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationOIDCAllOfMobile() *ApplicationOIDCAllOfMobile {
	this := ApplicationOIDCAllOfMobile{}
	return &this
}

// NewApplicationOIDCAllOfMobileWithDefaults instantiates a new ApplicationOIDCAllOfMobile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationOIDCAllOfMobileWithDefaults() *ApplicationOIDCAllOfMobile {
	this := ApplicationOIDCAllOfMobile{}
	return &this
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOfMobile) GetBundleId() string {
	if o == nil || IsNil(o.BundleId) {
		var ret string
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOfMobile) GetBundleIdOk() (*string, bool) {
	if o == nil || IsNil(o.BundleId) {
		return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOfMobile) HasBundleId() bool {
	if o != nil && !IsNil(o.BundleId) {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given string and assigns it to the BundleId field.
func (o *ApplicationOIDCAllOfMobile) SetBundleId(v string) {
	o.BundleId = &v
}

// GetPackageName returns the PackageName field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOfMobile) GetPackageName() string {
	if o == nil || IsNil(o.PackageName) {
		var ret string
		return ret
	}
	return *o.PackageName
}

// GetPackageNameOk returns a tuple with the PackageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOfMobile) GetPackageNameOk() (*string, bool) {
	if o == nil || IsNil(o.PackageName) {
		return nil, false
	}
	return o.PackageName, true
}

// HasPackageName returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOfMobile) HasPackageName() bool {
	if o != nil && !IsNil(o.PackageName) {
		return true
	}

	return false
}

// SetPackageName gets a reference to the given string and assigns it to the PackageName field.
func (o *ApplicationOIDCAllOfMobile) SetPackageName(v string) {
	o.PackageName = &v
}

// GetHuaweiAppId returns the HuaweiAppId field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOfMobile) GetHuaweiAppId() string {
	if o == nil || IsNil(o.HuaweiAppId) {
		var ret string
		return ret
	}
	return *o.HuaweiAppId
}

// GetHuaweiAppIdOk returns a tuple with the HuaweiAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOfMobile) GetHuaweiAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.HuaweiAppId) {
		return nil, false
	}
	return o.HuaweiAppId, true
}

// HasHuaweiAppId returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOfMobile) HasHuaweiAppId() bool {
	if o != nil && !IsNil(o.HuaweiAppId) {
		return true
	}

	return false
}

// SetHuaweiAppId gets a reference to the given string and assigns it to the HuaweiAppId field.
func (o *ApplicationOIDCAllOfMobile) SetHuaweiAppId(v string) {
	o.HuaweiAppId = &v
}

// GetHuaweiPackageName returns the HuaweiPackageName field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOfMobile) GetHuaweiPackageName() string {
	if o == nil || IsNil(o.HuaweiPackageName) {
		var ret string
		return ret
	}
	return *o.HuaweiPackageName
}

// GetHuaweiPackageNameOk returns a tuple with the HuaweiPackageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOfMobile) GetHuaweiPackageNameOk() (*string, bool) {
	if o == nil || IsNil(o.HuaweiPackageName) {
		return nil, false
	}
	return o.HuaweiPackageName, true
}

// HasHuaweiPackageName returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOfMobile) HasHuaweiPackageName() bool {
	if o != nil && !IsNil(o.HuaweiPackageName) {
		return true
	}

	return false
}

// SetHuaweiPackageName gets a reference to the given string and assigns it to the HuaweiPackageName field.
func (o *ApplicationOIDCAllOfMobile) SetHuaweiPackageName(v string) {
	o.HuaweiPackageName = &v
}

// GetPasscodeRefreshDuration returns the PasscodeRefreshDuration field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOfMobile) GetPasscodeRefreshDuration() ApplicationOIDCAllOfMobilePasscodeRefreshDuration {
	if o == nil || IsNil(o.PasscodeRefreshDuration) {
		var ret ApplicationOIDCAllOfMobilePasscodeRefreshDuration
		return ret
	}
	return *o.PasscodeRefreshDuration
}

// GetPasscodeRefreshDurationOk returns a tuple with the PasscodeRefreshDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOfMobile) GetPasscodeRefreshDurationOk() (*ApplicationOIDCAllOfMobilePasscodeRefreshDuration, bool) {
	if o == nil || IsNil(o.PasscodeRefreshDuration) {
		return nil, false
	}
	return o.PasscodeRefreshDuration, true
}

// HasPasscodeRefreshDuration returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOfMobile) HasPasscodeRefreshDuration() bool {
	if o != nil && !IsNil(o.PasscodeRefreshDuration) {
		return true
	}

	return false
}

// SetPasscodeRefreshDuration gets a reference to the given ApplicationOIDCAllOfMobilePasscodeRefreshDuration and assigns it to the PasscodeRefreshDuration field.
func (o *ApplicationOIDCAllOfMobile) SetPasscodeRefreshDuration(v ApplicationOIDCAllOfMobilePasscodeRefreshDuration) {
	o.PasscodeRefreshDuration = &v
}

// GetIntegrityDetection returns the IntegrityDetection field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOfMobile) GetIntegrityDetection() ApplicationOIDCAllOfMobileIntegrityDetection {
	if o == nil || IsNil(o.IntegrityDetection) {
		var ret ApplicationOIDCAllOfMobileIntegrityDetection
		return ret
	}
	return *o.IntegrityDetection
}

// GetIntegrityDetectionOk returns a tuple with the IntegrityDetection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOfMobile) GetIntegrityDetectionOk() (*ApplicationOIDCAllOfMobileIntegrityDetection, bool) {
	if o == nil || IsNil(o.IntegrityDetection) {
		return nil, false
	}
	return o.IntegrityDetection, true
}

// HasIntegrityDetection returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOfMobile) HasIntegrityDetection() bool {
	if o != nil && !IsNil(o.IntegrityDetection) {
		return true
	}

	return false
}

// SetIntegrityDetection gets a reference to the given ApplicationOIDCAllOfMobileIntegrityDetection and assigns it to the IntegrityDetection field.
func (o *ApplicationOIDCAllOfMobile) SetIntegrityDetection(v ApplicationOIDCAllOfMobileIntegrityDetection) {
	o.IntegrityDetection = &v
}

// GetUriPrefix returns the UriPrefix field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOfMobile) GetUriPrefix() string {
	if o == nil || IsNil(o.UriPrefix) {
		var ret string
		return ret
	}
	return *o.UriPrefix
}

// GetUriPrefixOk returns a tuple with the UriPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOfMobile) GetUriPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.UriPrefix) {
		return nil, false
	}
	return o.UriPrefix, true
}

// HasUriPrefix returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOfMobile) HasUriPrefix() bool {
	if o != nil && !IsNil(o.UriPrefix) {
		return true
	}

	return false
}

// SetUriPrefix gets a reference to the given string and assigns it to the UriPrefix field.
func (o *ApplicationOIDCAllOfMobile) SetUriPrefix(v string) {
	o.UriPrefix = &v
}

func (o ApplicationOIDCAllOfMobile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationOIDCAllOfMobile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BundleId) {
		toSerialize["bundleId"] = o.BundleId
	}
	if !IsNil(o.PackageName) {
		toSerialize["packageName"] = o.PackageName
	}
	if !IsNil(o.HuaweiAppId) {
		toSerialize["huaweiAppId"] = o.HuaweiAppId
	}
	if !IsNil(o.HuaweiPackageName) {
		toSerialize["huaweiPackageName"] = o.HuaweiPackageName
	}
	if !IsNil(o.PasscodeRefreshDuration) {
		toSerialize["passcodeRefreshDuration"] = o.PasscodeRefreshDuration
	}
	if !IsNil(o.IntegrityDetection) {
		toSerialize["integrityDetection"] = o.IntegrityDetection
	}
	if !IsNil(o.UriPrefix) {
		toSerialize["uriPrefix"] = o.UriPrefix
	}
	return toSerialize, nil
}

type NullableApplicationOIDCAllOfMobile struct {
	value *ApplicationOIDCAllOfMobile
	isSet bool
}

func (v NullableApplicationOIDCAllOfMobile) Get() *ApplicationOIDCAllOfMobile {
	return v.value
}

func (v *NullableApplicationOIDCAllOfMobile) Set(val *ApplicationOIDCAllOfMobile) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationOIDCAllOfMobile) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationOIDCAllOfMobile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationOIDCAllOfMobile(val *ApplicationOIDCAllOfMobile) *NullableApplicationOIDCAllOfMobile {
	return &NullableApplicationOIDCAllOfMobile{value: val, isSet: true}
}

func (v NullableApplicationOIDCAllOfMobile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationOIDCAllOfMobile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


