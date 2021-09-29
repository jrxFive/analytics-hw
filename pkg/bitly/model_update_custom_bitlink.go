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

// UpdateCustomBitlink struct for UpdateCustomBitlink
type UpdateCustomBitlink struct {
	BitlinkId *string `json:"bitlink_id,omitempty"`
}

// NewUpdateCustomBitlink instantiates a new UpdateCustomBitlink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCustomBitlink() *UpdateCustomBitlink {
	this := UpdateCustomBitlink{}
	return &this
}

// NewUpdateCustomBitlinkWithDefaults instantiates a new UpdateCustomBitlink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCustomBitlinkWithDefaults() *UpdateCustomBitlink {
	this := UpdateCustomBitlink{}
	return &this
}

// GetBitlinkId returns the BitlinkId field value if set, zero value otherwise.
func (o *UpdateCustomBitlink) GetBitlinkId() string {
	if o == nil || o.BitlinkId == nil {
		var ret string
		return ret
	}
	return *o.BitlinkId
}

// GetBitlinkIdOk returns a tuple with the BitlinkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomBitlink) GetBitlinkIdOk() (*string, bool) {
	if o == nil || o.BitlinkId == nil {
		return nil, false
	}
	return o.BitlinkId, true
}

// HasBitlinkId returns a boolean if a field has been set.
func (o *UpdateCustomBitlink) HasBitlinkId() bool {
	if o != nil && o.BitlinkId != nil {
		return true
	}

	return false
}

// SetBitlinkId gets a reference to the given string and assigns it to the BitlinkId field.
func (o *UpdateCustomBitlink) SetBitlinkId(v string) {
	o.BitlinkId = &v
}

func (o UpdateCustomBitlink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BitlinkId != nil {
		toSerialize["bitlink_id"] = o.BitlinkId
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateCustomBitlink struct {
	value *UpdateCustomBitlink
	isSet bool
}

func (v NullableUpdateCustomBitlink) Get() *UpdateCustomBitlink {
	return v.value
}

func (v *NullableUpdateCustomBitlink) Set(val *UpdateCustomBitlink) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCustomBitlink) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCustomBitlink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCustomBitlink(val *UpdateCustomBitlink) *NullableUpdateCustomBitlink {
	return &NullableUpdateCustomBitlink{value: val, isSet: true}
}

func (v NullableUpdateCustomBitlink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCustomBitlink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}