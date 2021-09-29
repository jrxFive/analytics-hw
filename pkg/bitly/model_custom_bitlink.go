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

// CustomBitlink struct for CustomBitlink
type CustomBitlink struct {
	Bitlink        *BitlinkBody            `json:"bitlink,omitempty"`
	BitlinkHistory *[]CustomBitlinkHistory `json:"bitlink_history,omitempty"`
	CustomBitlink  *string                 `json:"custom_bitlink,omitempty"`
}

// NewCustomBitlink instantiates a new CustomBitlink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomBitlink() *CustomBitlink {
	this := CustomBitlink{}
	return &this
}

// NewCustomBitlinkWithDefaults instantiates a new CustomBitlink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomBitlinkWithDefaults() *CustomBitlink {
	this := CustomBitlink{}
	return &this
}

// GetBitlink returns the Bitlink field value if set, zero value otherwise.
func (o *CustomBitlink) GetBitlink() BitlinkBody {
	if o == nil || o.Bitlink == nil {
		var ret BitlinkBody
		return ret
	}
	return *o.Bitlink
}

// GetBitlinkOk returns a tuple with the Bitlink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBitlink) GetBitlinkOk() (*BitlinkBody, bool) {
	if o == nil || o.Bitlink == nil {
		return nil, false
	}
	return o.Bitlink, true
}

// HasBitlink returns a boolean if a field has been set.
func (o *CustomBitlink) HasBitlink() bool {
	if o != nil && o.Bitlink != nil {
		return true
	}

	return false
}

// SetBitlink gets a reference to the given BitlinkBody and assigns it to the Bitlink field.
func (o *CustomBitlink) SetBitlink(v BitlinkBody) {
	o.Bitlink = &v
}

// GetBitlinkHistory returns the BitlinkHistory field value if set, zero value otherwise.
func (o *CustomBitlink) GetBitlinkHistory() []CustomBitlinkHistory {
	if o == nil || o.BitlinkHistory == nil {
		var ret []CustomBitlinkHistory
		return ret
	}
	return *o.BitlinkHistory
}

// GetBitlinkHistoryOk returns a tuple with the BitlinkHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBitlink) GetBitlinkHistoryOk() (*[]CustomBitlinkHistory, bool) {
	if o == nil || o.BitlinkHistory == nil {
		return nil, false
	}
	return o.BitlinkHistory, true
}

// HasBitlinkHistory returns a boolean if a field has been set.
func (o *CustomBitlink) HasBitlinkHistory() bool {
	if o != nil && o.BitlinkHistory != nil {
		return true
	}

	return false
}

// SetBitlinkHistory gets a reference to the given []CustomBitlinkHistory and assigns it to the BitlinkHistory field.
func (o *CustomBitlink) SetBitlinkHistory(v []CustomBitlinkHistory) {
	o.BitlinkHistory = &v
}

// GetCustomBitlink returns the CustomBitlink field value if set, zero value otherwise.
func (o *CustomBitlink) GetCustomBitlink() string {
	if o == nil || o.CustomBitlink == nil {
		var ret string
		return ret
	}
	return *o.CustomBitlink
}

// GetCustomBitlinkOk returns a tuple with the CustomBitlink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomBitlink) GetCustomBitlinkOk() (*string, bool) {
	if o == nil || o.CustomBitlink == nil {
		return nil, false
	}
	return o.CustomBitlink, true
}

// HasCustomBitlink returns a boolean if a field has been set.
func (o *CustomBitlink) HasCustomBitlink() bool {
	if o != nil && o.CustomBitlink != nil {
		return true
	}

	return false
}

// SetCustomBitlink gets a reference to the given string and assigns it to the CustomBitlink field.
func (o *CustomBitlink) SetCustomBitlink(v string) {
	o.CustomBitlink = &v
}

func (o CustomBitlink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bitlink != nil {
		toSerialize["bitlink"] = o.Bitlink
	}
	if o.BitlinkHistory != nil {
		toSerialize["bitlink_history"] = o.BitlinkHistory
	}
	if o.CustomBitlink != nil {
		toSerialize["custom_bitlink"] = o.CustomBitlink
	}
	return json.Marshal(toSerialize)
}

type NullableCustomBitlink struct {
	value *CustomBitlink
	isSet bool
}

func (v NullableCustomBitlink) Get() *CustomBitlink {
	return v.value
}

func (v *NullableCustomBitlink) Set(val *CustomBitlink) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomBitlink) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomBitlink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomBitlink(val *CustomBitlink) *NullableCustomBitlink {
	return &NullableCustomBitlink{value: val, isSet: true}
}

func (v NullableCustomBitlink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomBitlink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}