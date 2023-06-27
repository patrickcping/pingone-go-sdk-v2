/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"time"
)

// checks if the NotificationsSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationsSettings{}

// NotificationsSettings struct for NotificationsSettings
type NotificationsSettings struct {
	// The time the resource was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	DeliveryMode *EnumNotificationsSettingsDeliveryMode `json:"deliveryMode,omitempty"`
	Restrictions *NotificationsSettingsRestrictions `json:"restrictions,omitempty"`
	// A string that specifies the resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	// A list which represents the execution order of different SMS/Voice providers configured for the environment. The providers and their accounts’ configurations are represented in the list by the ID of the corresponding `PhoneDeliverySettings` resource. The only provider which is not represented by a `phoneDeliverySettingsID` is the PingOne Twilio provider. The PingOne Twilio provider is represented by the `PINGONE_TWILIO` string. If the `smsProvidersFallbackChain` list is empty, an SMS or voice message will be sent using the default Ping Twilio account. Otherwise, an SMS or voice message will be sent using the first provider in the list. If the server fails to queue the message using that provider, it will use the next provider in the list to try to send the message. This process will go on until there are no more providers in the list. If the server failed to send the message using all providers, the notification status is set to `FAILED`.
	SmsProvidersFallbackChain []string `json:"smsProvidersFallbackChain,omitempty"`
	From *NotificationsSettingsFrom `json:"from,omitempty"`
	ReplyTo *NotificationsSettingsReplyTo `json:"replyTo,omitempty"`
	Whitelist []NotificationsSettingsWhitelistInner `json:"whitelist,omitempty"`
}

// NewNotificationsSettings instantiates a new NotificationsSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationsSettings() *NotificationsSettings {
	this := NotificationsSettings{}
	return &this
}

// NewNotificationsSettingsWithDefaults instantiates a new NotificationsSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationsSettingsWithDefaults() *NotificationsSettings {
	this := NotificationsSettings{}
	return &this
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *NotificationsSettings) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettings) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *NotificationsSettings) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *NotificationsSettings) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *NotificationsSettings) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettings) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *NotificationsSettings) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *NotificationsSettings) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetDeliveryMode returns the DeliveryMode field value if set, zero value otherwise.
func (o *NotificationsSettings) GetDeliveryMode() EnumNotificationsSettingsDeliveryMode {
	if o == nil || IsNil(o.DeliveryMode) {
		var ret EnumNotificationsSettingsDeliveryMode
		return ret
	}
	return *o.DeliveryMode
}

// GetDeliveryModeOk returns a tuple with the DeliveryMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettings) GetDeliveryModeOk() (*EnumNotificationsSettingsDeliveryMode, bool) {
	if o == nil || IsNil(o.DeliveryMode) {
		return nil, false
	}
	return o.DeliveryMode, true
}

// HasDeliveryMode returns a boolean if a field has been set.
func (o *NotificationsSettings) HasDeliveryMode() bool {
	if o != nil && !IsNil(o.DeliveryMode) {
		return true
	}

	return false
}

// SetDeliveryMode gets a reference to the given EnumNotificationsSettingsDeliveryMode and assigns it to the DeliveryMode field.
func (o *NotificationsSettings) SetDeliveryMode(v EnumNotificationsSettingsDeliveryMode) {
	o.DeliveryMode = &v
}

// GetRestrictions returns the Restrictions field value if set, zero value otherwise.
func (o *NotificationsSettings) GetRestrictions() NotificationsSettingsRestrictions {
	if o == nil || IsNil(o.Restrictions) {
		var ret NotificationsSettingsRestrictions
		return ret
	}
	return *o.Restrictions
}

// GetRestrictionsOk returns a tuple with the Restrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettings) GetRestrictionsOk() (*NotificationsSettingsRestrictions, bool) {
	if o == nil || IsNil(o.Restrictions) {
		return nil, false
	}
	return o.Restrictions, true
}

// HasRestrictions returns a boolean if a field has been set.
func (o *NotificationsSettings) HasRestrictions() bool {
	if o != nil && !IsNil(o.Restrictions) {
		return true
	}

	return false
}

// SetRestrictions gets a reference to the given NotificationsSettingsRestrictions and assigns it to the Restrictions field.
func (o *NotificationsSettings) SetRestrictions(v NotificationsSettingsRestrictions) {
	o.Restrictions = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NotificationsSettings) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettings) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NotificationsSettings) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NotificationsSettings) SetId(v string) {
	o.Id = &v
}

// GetSmsProvidersFallbackChain returns the SmsProvidersFallbackChain field value if set, zero value otherwise.
func (o *NotificationsSettings) GetSmsProvidersFallbackChain() []string {
	if o == nil || IsNil(o.SmsProvidersFallbackChain) {
		var ret []string
		return ret
	}
	return o.SmsProvidersFallbackChain
}

// GetSmsProvidersFallbackChainOk returns a tuple with the SmsProvidersFallbackChain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettings) GetSmsProvidersFallbackChainOk() ([]string, bool) {
	if o == nil || IsNil(o.SmsProvidersFallbackChain) {
		return nil, false
	}
	return o.SmsProvidersFallbackChain, true
}

// HasSmsProvidersFallbackChain returns a boolean if a field has been set.
func (o *NotificationsSettings) HasSmsProvidersFallbackChain() bool {
	if o != nil && !IsNil(o.SmsProvidersFallbackChain) {
		return true
	}

	return false
}

// SetSmsProvidersFallbackChain gets a reference to the given []string and assigns it to the SmsProvidersFallbackChain field.
func (o *NotificationsSettings) SetSmsProvidersFallbackChain(v []string) {
	o.SmsProvidersFallbackChain = v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *NotificationsSettings) GetFrom() NotificationsSettingsFrom {
	if o == nil || IsNil(o.From) {
		var ret NotificationsSettingsFrom
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettings) GetFromOk() (*NotificationsSettingsFrom, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *NotificationsSettings) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given NotificationsSettingsFrom and assigns it to the From field.
func (o *NotificationsSettings) SetFrom(v NotificationsSettingsFrom) {
	o.From = &v
}

// GetReplyTo returns the ReplyTo field value if set, zero value otherwise.
func (o *NotificationsSettings) GetReplyTo() NotificationsSettingsReplyTo {
	if o == nil || IsNil(o.ReplyTo) {
		var ret NotificationsSettingsReplyTo
		return ret
	}
	return *o.ReplyTo
}

// GetReplyToOk returns a tuple with the ReplyTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettings) GetReplyToOk() (*NotificationsSettingsReplyTo, bool) {
	if o == nil || IsNil(o.ReplyTo) {
		return nil, false
	}
	return o.ReplyTo, true
}

// HasReplyTo returns a boolean if a field has been set.
func (o *NotificationsSettings) HasReplyTo() bool {
	if o != nil && !IsNil(o.ReplyTo) {
		return true
	}

	return false
}

// SetReplyTo gets a reference to the given NotificationsSettingsReplyTo and assigns it to the ReplyTo field.
func (o *NotificationsSettings) SetReplyTo(v NotificationsSettingsReplyTo) {
	o.ReplyTo = &v
}

// GetWhitelist returns the Whitelist field value if set, zero value otherwise.
func (o *NotificationsSettings) GetWhitelist() []NotificationsSettingsWhitelistInner {
	if o == nil || IsNil(o.Whitelist) {
		var ret []NotificationsSettingsWhitelistInner
		return ret
	}
	return o.Whitelist
}

// GetWhitelistOk returns a tuple with the Whitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettings) GetWhitelistOk() ([]NotificationsSettingsWhitelistInner, bool) {
	if o == nil || IsNil(o.Whitelist) {
		return nil, false
	}
	return o.Whitelist, true
}

// HasWhitelist returns a boolean if a field has been set.
func (o *NotificationsSettings) HasWhitelist() bool {
	if o != nil && !IsNil(o.Whitelist) {
		return true
	}

	return false
}

// SetWhitelist gets a reference to the given []NotificationsSettingsWhitelistInner and assigns it to the Whitelist field.
func (o *NotificationsSettings) SetWhitelist(v []NotificationsSettingsWhitelistInner) {
	o.Whitelist = v
}

func (o NotificationsSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationsSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: updatedAt is readOnly
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.DeliveryMode) {
		toSerialize["deliveryMode"] = o.DeliveryMode
	}
	if !IsNil(o.Restrictions) {
		toSerialize["restrictions"] = o.Restrictions
	}
	// skip: id is readOnly
	if !IsNil(o.SmsProvidersFallbackChain) {
		toSerialize["smsProvidersFallbackChain"] = o.SmsProvidersFallbackChain
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.ReplyTo) {
		toSerialize["replyTo"] = o.ReplyTo
	}
	if !IsNil(o.Whitelist) {
		toSerialize["whitelist"] = o.Whitelist
	}
	return toSerialize, nil
}

type NullableNotificationsSettings struct {
	value *NotificationsSettings
	isSet bool
}

func (v NullableNotificationsSettings) Get() *NotificationsSettings {
	return v.value
}

func (v *NullableNotificationsSettings) Set(val *NotificationsSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationsSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationsSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationsSettings(val *NotificationsSettings) *NullableNotificationsSettings {
	return &NullableNotificationsSettings{value: val, isSet: true}
}

func (v NullableNotificationsSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationsSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


