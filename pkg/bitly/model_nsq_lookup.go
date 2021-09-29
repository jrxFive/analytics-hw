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

// NSQLookup struct for NSQLookup
type NSQLookup struct {
	Producers *[]NSQProducer `json:"producers,omitempty"`
}

// NewNSQLookup instantiates a new NSQLookup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNSQLookup() *NSQLookup {
	this := NSQLookup{}
	return &this
}

// NewNSQLookupWithDefaults instantiates a new NSQLookup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNSQLookupWithDefaults() *NSQLookup {
	this := NSQLookup{}
	return &this
}

// GetProducers returns the Producers field value if set, zero value otherwise.
func (o *NSQLookup) GetProducers() []NSQProducer {
	if o == nil || o.Producers == nil {
		var ret []NSQProducer
		return ret
	}
	return *o.Producers
}

// GetProducersOk returns a tuple with the Producers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NSQLookup) GetProducersOk() (*[]NSQProducer, bool) {
	if o == nil || o.Producers == nil {
		return nil, false
	}
	return o.Producers, true
}

// HasProducers returns a boolean if a field has been set.
func (o *NSQLookup) HasProducers() bool {
	if o != nil && o.Producers != nil {
		return true
	}

	return false
}

// SetProducers gets a reference to the given []NSQProducer and assigns it to the Producers field.
func (o *NSQLookup) SetProducers(v []NSQProducer) {
	o.Producers = &v
}

func (o NSQLookup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Producers != nil {
		toSerialize["producers"] = o.Producers
	}
	return json.Marshal(toSerialize)
}

type NullableNSQLookup struct {
	value *NSQLookup
	isSet bool
}

func (v NullableNSQLookup) Get() *NSQLookup {
	return v.value
}

func (v *NullableNSQLookup) Set(val *NSQLookup) {
	v.value = val
	v.isSet = true
}

func (v NullableNSQLookup) IsSet() bool {
	return v.isSet
}

func (v *NullableNSQLookup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNSQLookup(val *NSQLookup) *NullableNSQLookup {
	return &NullableNSQLookup{value: val, isSet: true}
}

func (v NullableNSQLookup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNSQLookup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
