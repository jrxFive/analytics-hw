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

// ClickLink struct for ClickLink
type ClickLink struct {
	Clicks *int32  `json:"clicks,omitempty"`
	Id     *string `json:"id,omitempty"`
}

// NewClickLink instantiates a new ClickLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClickLink() *ClickLink {
	this := ClickLink{}
	return &this
}

// NewClickLinkWithDefaults instantiates a new ClickLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClickLinkWithDefaults() *ClickLink {
	this := ClickLink{}
	return &this
}

// GetClicks returns the Clicks field value if set, zero value otherwise.
func (o *ClickLink) GetClicks() int32 {
	if o == nil || o.Clicks == nil {
		var ret int32
		return ret
	}
	return *o.Clicks
}

// GetClicksOk returns a tuple with the Clicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickLink) GetClicksOk() (*int32, bool) {
	if o == nil || o.Clicks == nil {
		return nil, false
	}
	return o.Clicks, true
}

// HasClicks returns a boolean if a field has been set.
func (o *ClickLink) HasClicks() bool {
	if o != nil && o.Clicks != nil {
		return true
	}

	return false
}

// SetClicks gets a reference to the given int32 and assigns it to the Clicks field.
func (o *ClickLink) SetClicks(v int32) {
	o.Clicks = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClickLink) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickLink) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClickLink) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ClickLink) SetId(v string) {
	o.Id = &v
}

func (o ClickLink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Clicks != nil {
		toSerialize["clicks"] = o.Clicks
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableClickLink struct {
	value *ClickLink
	isSet bool
}

func (v NullableClickLink) Get() *ClickLink {
	return v.value
}

func (v *NullableClickLink) Set(val *ClickLink) {
	v.value = val
	v.isSet = true
}

func (v NullableClickLink) IsSet() bool {
	return v.isSet
}

func (v *NullableClickLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClickLink(val *ClickLink) *NullableClickLink {
	return &NullableClickLink{value: val, isSet: true}
}

func (v NullableClickLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClickLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}