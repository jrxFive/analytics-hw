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

// AddCustomBitlink struct for AddCustomBitlink
type AddCustomBitlink struct {
	BitlinkId     *string `json:"bitlink_id,omitempty"`
	CustomBitlink *string `json:"custom_bitlink,omitempty"`
}

// NewAddCustomBitlink instantiates a new AddCustomBitlink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddCustomBitlink() *AddCustomBitlink {
	this := AddCustomBitlink{}
	return &this
}

// NewAddCustomBitlinkWithDefaults instantiates a new AddCustomBitlink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCustomBitlinkWithDefaults() *AddCustomBitlink {
	this := AddCustomBitlink{}
	return &this
}

// GetBitlinkId returns the BitlinkId field value if set, zero value otherwise.
func (o *AddCustomBitlink) GetBitlinkId() string {
	if o == nil || o.BitlinkId == nil {
		var ret string
		return ret
	}
	return *o.BitlinkId
}

// GetBitlinkIdOk returns a tuple with the BitlinkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCustomBitlink) GetBitlinkIdOk() (*string, bool) {
	if o == nil || o.BitlinkId == nil {
		return nil, false
	}
	return o.BitlinkId, true
}

// HasBitlinkId returns a boolean if a field has been set.
func (o *AddCustomBitlink) HasBitlinkId() bool {
	if o != nil && o.BitlinkId != nil {
		return true
	}

	return false
}

// SetBitlinkId gets a reference to the given string and assigns it to the BitlinkId field.
func (o *AddCustomBitlink) SetBitlinkId(v string) {
	o.BitlinkId = &v
}

// GetCustomBitlink returns the CustomBitlink field value if set, zero value otherwise.
func (o *AddCustomBitlink) GetCustomBitlink() string {
	if o == nil || o.CustomBitlink == nil {
		var ret string
		return ret
	}
	return *o.CustomBitlink
}

// GetCustomBitlinkOk returns a tuple with the CustomBitlink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCustomBitlink) GetCustomBitlinkOk() (*string, bool) {
	if o == nil || o.CustomBitlink == nil {
		return nil, false
	}
	return o.CustomBitlink, true
}

// HasCustomBitlink returns a boolean if a field has been set.
func (o *AddCustomBitlink) HasCustomBitlink() bool {
	if o != nil && o.CustomBitlink != nil {
		return true
	}

	return false
}

// SetCustomBitlink gets a reference to the given string and assigns it to the CustomBitlink field.
func (o *AddCustomBitlink) SetCustomBitlink(v string) {
	o.CustomBitlink = &v
}

func (o AddCustomBitlink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BitlinkId != nil {
		toSerialize["bitlink_id"] = o.BitlinkId
	}
	if o.CustomBitlink != nil {
		toSerialize["custom_bitlink"] = o.CustomBitlink
	}
	return json.Marshal(toSerialize)
}

type NullableAddCustomBitlink struct {
	value *AddCustomBitlink
	isSet bool
}

func (v NullableAddCustomBitlink) Get() *AddCustomBitlink {
	return v.value
}

func (v *NullableAddCustomBitlink) Set(val *AddCustomBitlink) {
	v.value = val
	v.isSet = true
}

func (v NullableAddCustomBitlink) IsSet() bool {
	return v.isSet
}

func (v *NullableAddCustomBitlink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddCustomBitlink(val *AddCustomBitlink) *NullableAddCustomBitlink {
	return &NullableAddCustomBitlink{value: val, isSet: true}
}

func (v NullableAddCustomBitlink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddCustomBitlink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
