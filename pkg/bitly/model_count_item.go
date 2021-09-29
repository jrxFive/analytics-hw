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

// CountItem struct for CountItem
type CountItem struct {
	Count *int32  `json:"count,omitempty"`
	Ts    *string `json:"ts,omitempty"`
}

// NewCountItem instantiates a new CountItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountItem() *CountItem {
	this := CountItem{}
	return &this
}

// NewCountItemWithDefaults instantiates a new CountItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountItemWithDefaults() *CountItem {
	this := CountItem{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *CountItem) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountItem) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *CountItem) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *CountItem) SetCount(v int32) {
	o.Count = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *CountItem) GetTs() string {
	if o == nil || o.Ts == nil {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountItem) GetTsOk() (*string, bool) {
	if o == nil || o.Ts == nil {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *CountItem) HasTs() bool {
	if o != nil && o.Ts != nil {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *CountItem) SetTs(v string) {
	o.Ts = &v
}

func (o CountItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Ts != nil {
		toSerialize["ts"] = o.Ts
	}
	return json.Marshal(toSerialize)
}

type NullableCountItem struct {
	value *CountItem
	isSet bool
}

func (v NullableCountItem) Get() *CountItem {
	return v.value
}

func (v *NullableCountItem) Set(val *CountItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCountItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCountItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountItem(val *CountItem) *NullableCountItem {
	return &NullableCountItem{value: val, isSet: true}
}

func (v NullableCountItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}