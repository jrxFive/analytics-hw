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

// BulkUpdate struct for BulkUpdate
type BulkUpdate struct {
	Links *[]string `json:"links,omitempty"`
}

// NewBulkUpdate instantiates a new BulkUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkUpdate() *BulkUpdate {
	this := BulkUpdate{}
	return &this
}

// NewBulkUpdateWithDefaults instantiates a new BulkUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkUpdateWithDefaults() *BulkUpdate {
	this := BulkUpdate{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BulkUpdate) GetLinks() []string {
	if o == nil || o.Links == nil {
		var ret []string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdate) GetLinksOk() (*[]string, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BulkUpdate) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []string and assigns it to the Links field.
func (o *BulkUpdate) SetLinks(v []string) {
	o.Links = &v
}

func (o BulkUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableBulkUpdate struct {
	value *BulkUpdate
	isSet bool
}

func (v NullableBulkUpdate) Get() *BulkUpdate {
	return v.value
}

func (v *NullableBulkUpdate) Set(val *BulkUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkUpdate(val *BulkUpdate) *NullableBulkUpdate {
	return &NullableBulkUpdate{value: val, isSet: true}
}

func (v NullableBulkUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
