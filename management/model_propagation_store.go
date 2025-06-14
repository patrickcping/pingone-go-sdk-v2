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

// checks if the PropagationStore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PropagationStore{}

// PropagationStore struct for PropagationStore
type PropagationStore struct {
	Links *map[string]LinksHATEOASValue `json:"_links,omitempty"`
	// Configuration properties specific to each identity propagation store.
	Configuration map[string]interface{} `json:"configuration"`
	// A description for the identity propagation store.
	Description *string            `json:"description,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// A string that specifies the resource’s unique identifier.
	Id    *string                `json:"id,omitempty"`
	Image *PropagationStoreImage `json:"image,omitempty"`
	// Indicates whether or not to enable deprovisioning of users for a store when it is deleted. The deprovisioning occurs when a new revision is created (`POST {{apiPath}}/environments/{{envID}}/propagation/revisions`).
	Managed *bool `json:"managed,omitempty"`
	// The name of the identity store.
	Name       string                      `json:"name"`
	Status     *EnumPropagationStoreStatus `json:"status,omitempty"`
	SyncStatus *PropagationStoreSyncStatus `json:"syncStatus,omitempty"`
	Type       EnumPropagationStoreType    `json:"type"`
}

// NewPropagationStore instantiates a new PropagationStore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropagationStore(configuration map[string]interface{}, name string, type_ EnumPropagationStoreType) *PropagationStore {
	this := PropagationStore{}
	this.Configuration = configuration
	this.Name = name
	this.Type = type_
	return &this
}

// NewPropagationStoreWithDefaults instantiates a new PropagationStore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropagationStoreWithDefaults() *PropagationStore {
	this := PropagationStore{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PropagationStore) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationStore) GetLinksOk() (*map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PropagationStore) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *PropagationStore) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = &v
}

// GetConfiguration returns the Configuration field value
func (o *PropagationStore) GetConfiguration() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value
// and a boolean to check if the value has been set.
func (o *PropagationStore) GetConfigurationOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Configuration, true
}

// SetConfiguration sets field value
func (o *PropagationStore) SetConfiguration(v map[string]interface{}) {
	o.Configuration = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PropagationStore) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationStore) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PropagationStore) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PropagationStore) SetDescription(v string) {
	o.Description = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *PropagationStore) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationStore) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *PropagationStore) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *PropagationStore) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PropagationStore) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationStore) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PropagationStore) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PropagationStore) SetId(v string) {
	o.Id = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *PropagationStore) GetImage() PropagationStoreImage {
	if o == nil || IsNil(o.Image) {
		var ret PropagationStoreImage
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationStore) GetImageOk() (*PropagationStoreImage, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *PropagationStore) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given PropagationStoreImage and assigns it to the Image field.
func (o *PropagationStore) SetImage(v PropagationStoreImage) {
	o.Image = &v
}

// GetManaged returns the Managed field value if set, zero value otherwise.
func (o *PropagationStore) GetManaged() bool {
	if o == nil || IsNil(o.Managed) {
		var ret bool
		return ret
	}
	return *o.Managed
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationStore) GetManagedOk() (*bool, bool) {
	if o == nil || IsNil(o.Managed) {
		return nil, false
	}
	return o.Managed, true
}

// HasManaged returns a boolean if a field has been set.
func (o *PropagationStore) HasManaged() bool {
	if o != nil && !IsNil(o.Managed) {
		return true
	}

	return false
}

// SetManaged gets a reference to the given bool and assigns it to the Managed field.
func (o *PropagationStore) SetManaged(v bool) {
	o.Managed = &v
}

// GetName returns the Name field value
func (o *PropagationStore) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PropagationStore) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PropagationStore) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PropagationStore) GetStatus() EnumPropagationStoreStatus {
	if o == nil || IsNil(o.Status) {
		var ret EnumPropagationStoreStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationStore) GetStatusOk() (*EnumPropagationStoreStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PropagationStore) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given EnumPropagationStoreStatus and assigns it to the Status field.
func (o *PropagationStore) SetStatus(v EnumPropagationStoreStatus) {
	o.Status = &v
}

// GetSyncStatus returns the SyncStatus field value if set, zero value otherwise.
func (o *PropagationStore) GetSyncStatus() PropagationStoreSyncStatus {
	if o == nil || IsNil(o.SyncStatus) {
		var ret PropagationStoreSyncStatus
		return ret
	}
	return *o.SyncStatus
}

// GetSyncStatusOk returns a tuple with the SyncStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationStore) GetSyncStatusOk() (*PropagationStoreSyncStatus, bool) {
	if o == nil || IsNil(o.SyncStatus) {
		return nil, false
	}
	return o.SyncStatus, true
}

// HasSyncStatus returns a boolean if a field has been set.
func (o *PropagationStore) HasSyncStatus() bool {
	if o != nil && !IsNil(o.SyncStatus) {
		return true
	}

	return false
}

// SetSyncStatus gets a reference to the given PropagationStoreSyncStatus and assigns it to the SyncStatus field.
func (o *PropagationStore) SetSyncStatus(v PropagationStoreSyncStatus) {
	o.SyncStatus = &v
}

// GetType returns the Type field value
func (o *PropagationStore) GetType() EnumPropagationStoreType {
	if o == nil {
		var ret EnumPropagationStoreType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PropagationStore) GetTypeOk() (*EnumPropagationStoreType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PropagationStore) SetType(v EnumPropagationStoreType) {
	o.Type = v
}

func (o PropagationStore) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PropagationStore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	toSerialize["configuration"] = o.Configuration
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.Managed) {
		toSerialize["managed"] = o.Managed
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SyncStatus) {
		toSerialize["syncStatus"] = o.SyncStatus
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullablePropagationStore struct {
	value *PropagationStore
	isSet bool
}

func (v NullablePropagationStore) Get() *PropagationStore {
	return v.value
}

func (v *NullablePropagationStore) Set(val *PropagationStore) {
	v.value = val
	v.isSet = true
}

func (v NullablePropagationStore) IsSet() bool {
	return v.isSet
}

func (v *NullablePropagationStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropagationStore(val *PropagationStore) *NullablePropagationStore {
	return &NullablePropagationStore{value: val, isSet: true}
}

func (v NullablePropagationStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropagationStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
