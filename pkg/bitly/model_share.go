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

// Share struct for Share
type Share struct {
	Text *string `json:"text,omitempty"`
	Link *string `json:"link,omitempty"`
}

// NewShare instantiates a new Share object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShare() *Share {
	this := Share{}
	return &this
}

// NewShareWithDefaults instantiates a new Share object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShareWithDefaults() *Share {
	this := Share{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *Share) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *Share) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *Share) SetText(v string) {
	o.Text = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *Share) GetLink() string {
	if o == nil || o.Link == nil {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetLinkOk() (*string, bool) {
	if o == nil || o.Link == nil {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *Share) HasLink() bool {
	if o != nil && o.Link != nil {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *Share) SetLink(v string) {
	o.Link = &v
}

func (o Share) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	if o.Link != nil {
		toSerialize["link"] = o.Link
	}
	return json.Marshal(toSerialize)
}

type NullableShare struct {
	value *Share
	isSet bool
}

func (v NullableShare) Get() *Share {
	return v.value
}

func (v *NullableShare) Set(val *Share) {
	v.value = val
	v.isSet = true
}

func (v NullableShare) IsSet() bool {
	return v.isSet
}

func (v *NullableShare) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShare(val *Share) *NullableShare {
	return &NullableShare{value: val, isSet: true}
}

func (v NullableShare) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShare) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
