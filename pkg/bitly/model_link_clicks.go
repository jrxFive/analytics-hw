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

// LinkClicks struct for LinkClicks
type LinkClicks struct {
	Date   *string `json:"date,omitempty"`
	Clicks *int32  `json:"clicks,omitempty"`
}

// NewLinkClicks instantiates a new LinkClicks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkClicks() *LinkClicks {
	this := LinkClicks{}
	return &this
}

// NewLinkClicksWithDefaults instantiates a new LinkClicks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkClicksWithDefaults() *LinkClicks {
	this := LinkClicks{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *LinkClicks) GetDate() string {
	if o == nil || o.Date == nil {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkClicks) GetDateOk() (*string, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *LinkClicks) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *LinkClicks) SetDate(v string) {
	o.Date = &v
}

// GetClicks returns the Clicks field value if set, zero value otherwise.
func (o *LinkClicks) GetClicks() int32 {
	if o == nil || o.Clicks == nil {
		var ret int32
		return ret
	}
	return *o.Clicks
}

// GetClicksOk returns a tuple with the Clicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkClicks) GetClicksOk() (*int32, bool) {
	if o == nil || o.Clicks == nil {
		return nil, false
	}
	return o.Clicks, true
}

// HasClicks returns a boolean if a field has been set.
func (o *LinkClicks) HasClicks() bool {
	if o != nil && o.Clicks != nil {
		return true
	}

	return false
}

// SetClicks gets a reference to the given int32 and assigns it to the Clicks field.
func (o *LinkClicks) SetClicks(v int32) {
	o.Clicks = &v
}

func (o LinkClicks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.Clicks != nil {
		toSerialize["clicks"] = o.Clicks
	}
	return json.Marshal(toSerialize)
}

type NullableLinkClicks struct {
	value *LinkClicks
	isSet bool
}

func (v NullableLinkClicks) Get() *LinkClicks {
	return v.value
}

func (v *NullableLinkClicks) Set(val *LinkClicks) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkClicks) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkClicks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkClicks(val *LinkClicks) *NullableLinkClicks {
	return &NullableLinkClicks{value: val, isSet: true}
}

func (v NullableLinkClicks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkClicks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
