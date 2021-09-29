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

// FacetCountItem struct for FacetCountItem
type FacetCountItem struct {
	Count *int32  `json:"count,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewFacetCountItem instantiates a new FacetCountItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFacetCountItem() *FacetCountItem {
	this := FacetCountItem{}
	return &this
}

// NewFacetCountItemWithDefaults instantiates a new FacetCountItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFacetCountItemWithDefaults() *FacetCountItem {
	this := FacetCountItem{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *FacetCountItem) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetCountItem) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *FacetCountItem) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *FacetCountItem) SetCount(v int32) {
	o.Count = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FacetCountItem) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetCountItem) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FacetCountItem) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *FacetCountItem) SetValue(v string) {
	o.Value = &v
}

func (o FacetCountItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableFacetCountItem struct {
	value *FacetCountItem
	isSet bool
}

func (v NullableFacetCountItem) Get() *FacetCountItem {
	return v.value
}

func (v *NullableFacetCountItem) Set(val *FacetCountItem) {
	v.value = val
	v.isSet = true
}

func (v NullableFacetCountItem) IsSet() bool {
	return v.isSet
}

func (v *NullableFacetCountItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFacetCountItem(val *FacetCountItem) *NullableFacetCountItem {
	return &NullableFacetCountItem{value: val, isSet: true}
}

func (v NullableFacetCountItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFacetCountItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}