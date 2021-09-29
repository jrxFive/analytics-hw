/*
Bitly API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.0
Contact: api@bitly.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// BitlinkHistory struct for BitlinkHistory
type BitlinkHistory struct {
	References *map[string]string `json:"references,omitempty"`
	// The backhalf of the underlying hash link
	Hash          *string `json:"hash,omitempty"`
	GroupGuid     *string `json:"group_guid,omitempty"`
	CreatedAt     *string `json:"created_at,omitempty"`
	IsActive      *bool   `json:"is_active,omitempty"`
	LongUrl       *string `json:"long_url,omitempty"`
	DeactivatedAt *string `json:"deactivated_at,omitempty"`
	Login         *string `json:"login,omitempty"`
	// The domain and backhalf of the underlying hash link
	TargetBitlinkId *string `json:"target_bitlink_id,omitempty"`
	// The domain and backhalf of the parent override
	Id *string `json:"id,omitempty"`
}

// NewBitlinkHistory instantiates a new BitlinkHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBitlinkHistory() *BitlinkHistory {
	this := BitlinkHistory{}
	return &this
}

// NewBitlinkHistoryWithDefaults instantiates a new BitlinkHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBitlinkHistoryWithDefaults() *BitlinkHistory {
	this := BitlinkHistory{}
	return &this
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *BitlinkHistory) GetReferences() map[string]string {
	if o == nil || o.References == nil {
		var ret map[string]string
		return ret
	}
	return *o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BitlinkHistory) GetReferencesOk() (*map[string]string, bool) {
	if o == nil || o.References == nil {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *BitlinkHistory) HasReferences() bool {
	if o != nil && o.References != nil {
		return true
	}

	return false
}

// SetReferences gets a reference to the given map[string]string and assigns it to the References field.
func (o *BitlinkHistory) SetReferences(v map[string]string) {
	o.References = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *BitlinkHistory) GetHash() string {
	if o == nil || o.Hash == nil {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BitlinkHistory) GetHashOk() (*string, bool) {
	if o == nil || o.Hash == nil {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *BitlinkHistory) HasHash() bool {
	if o != nil && o.Hash != nil {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *BitlinkHistory) SetHash(v string) {
	o.Hash = &v
}

// GetGroupGuid returns the GroupGuid field value if set, zero value otherwise.
func (o *BitlinkHistory) GetGroupGuid() string {
	if o == nil || o.GroupGuid == nil {
		var ret string
		return ret
	}
	return *o.GroupGuid
}

// GetGroupGuidOk returns a tuple with the GroupGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BitlinkHistory) GetGroupGuidOk() (*string, bool) {
	if o == nil || o.GroupGuid == nil {
		return nil, false
	}
	return o.GroupGuid, true
}

// HasGroupGuid returns a boolean if a field has been set.
func (o *BitlinkHistory) HasGroupGuid() bool {
	if o != nil && o.GroupGuid != nil {
		return true
	}

	return false
}

// SetGroupGuid gets a reference to the given string and assigns it to the GroupGuid field.
func (o *BitlinkHistory) SetGroupGuid(v string) {
	o.GroupGuid = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BitlinkHistory) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BitlinkHistory) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BitlinkHistory) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *BitlinkHistory) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *BitlinkHistory) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BitlinkHistory) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *BitlinkHistory) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *BitlinkHistory) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetLongUrl returns the LongUrl field value if set, zero value otherwise.
func (o *BitlinkHistory) GetLongUrl() string {
	if o == nil || o.LongUrl == nil {
		var ret string
		return ret
	}
	return *o.LongUrl
}

// GetLongUrlOk returns a tuple with the LongUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BitlinkHistory) GetLongUrlOk() (*string, bool) {
	if o == nil || o.LongUrl == nil {
		return nil, false
	}
	return o.LongUrl, true
}

// HasLongUrl returns a boolean if a field has been set.
func (o *BitlinkHistory) HasLongUrl() bool {
	if o != nil && o.LongUrl != nil {
		return true
	}

	return false
}

// SetLongUrl gets a reference to the given string and assigns it to the LongUrl field.
func (o *BitlinkHistory) SetLongUrl(v string) {
	o.LongUrl = &v
}

// GetDeactivatedAt returns the DeactivatedAt field value if set, zero value otherwise.
func (o *BitlinkHistory) GetDeactivatedAt() string {
	if o == nil || o.DeactivatedAt == nil {
		var ret string
		return ret
	}
	return *o.DeactivatedAt
}

// GetDeactivatedAtOk returns a tuple with the DeactivatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BitlinkHistory) GetDeactivatedAtOk() (*string, bool) {
	if o == nil || o.DeactivatedAt == nil {
		return nil, false
	}
	return o.DeactivatedAt, true
}

// HasDeactivatedAt returns a boolean if a field has been set.
func (o *BitlinkHistory) HasDeactivatedAt() bool {
	if o != nil && o.DeactivatedAt != nil {
		return true
	}

	return false
}

// SetDeactivatedAt gets a reference to the given string and assigns it to the DeactivatedAt field.
func (o *BitlinkHistory) SetDeactivatedAt(v string) {
	o.DeactivatedAt = &v
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *BitlinkHistory) GetLogin() string {
	if o == nil || o.Login == nil {
		var ret string
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BitlinkHistory) GetLoginOk() (*string, bool) {
	if o == nil || o.Login == nil {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *BitlinkHistory) HasLogin() bool {
	if o != nil && o.Login != nil {
		return true
	}

	return false
}

// SetLogin gets a reference to the given string and assigns it to the Login field.
func (o *BitlinkHistory) SetLogin(v string) {
	o.Login = &v
}

// GetTargetBitlinkId returns the TargetBitlinkId field value if set, zero value otherwise.
func (o *BitlinkHistory) GetTargetBitlinkId() string {
	if o == nil || o.TargetBitlinkId == nil {
		var ret string
		return ret
	}
	return *o.TargetBitlinkId
}

// GetTargetBitlinkIdOk returns a tuple with the TargetBitlinkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BitlinkHistory) GetTargetBitlinkIdOk() (*string, bool) {
	if o == nil || o.TargetBitlinkId == nil {
		return nil, false
	}
	return o.TargetBitlinkId, true
}

// HasTargetBitlinkId returns a boolean if a field has been set.
func (o *BitlinkHistory) HasTargetBitlinkId() bool {
	if o != nil && o.TargetBitlinkId != nil {
		return true
	}

	return false
}

// SetTargetBitlinkId gets a reference to the given string and assigns it to the TargetBitlinkId field.
func (o *BitlinkHistory) SetTargetBitlinkId(v string) {
	o.TargetBitlinkId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BitlinkHistory) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BitlinkHistory) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BitlinkHistory) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BitlinkHistory) SetId(v string) {
	o.Id = &v
}

func (o BitlinkHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.References != nil {
		toSerialize["references"] = o.References
	}
	if o.Hash != nil {
		toSerialize["hash"] = o.Hash
	}
	if o.GroupGuid != nil {
		toSerialize["group_guid"] = o.GroupGuid
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.IsActive != nil {
		toSerialize["is_active"] = o.IsActive
	}
	if o.LongUrl != nil {
		toSerialize["long_url"] = o.LongUrl
	}
	if o.DeactivatedAt != nil {
		toSerialize["deactivated_at"] = o.DeactivatedAt
	}
	if o.Login != nil {
		toSerialize["login"] = o.Login
	}
	if o.TargetBitlinkId != nil {
		toSerialize["target_bitlink_id"] = o.TargetBitlinkId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableBitlinkHistory struct {
	value *BitlinkHistory
	isSet bool
}

func (v NullableBitlinkHistory) Get() *BitlinkHistory {
	return v.value
}

func (v *NullableBitlinkHistory) Set(val *BitlinkHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableBitlinkHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableBitlinkHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBitlinkHistory(val *BitlinkHistory) *NullableBitlinkHistory {
	return &NullableBitlinkHistory{value: val, isSet: true}
}

func (v NullableBitlinkHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBitlinkHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
