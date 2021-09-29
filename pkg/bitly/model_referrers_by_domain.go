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

// ReferrersByDomain struct for ReferrersByDomain
type ReferrersByDomain struct {
	Referrers *[]Metric `json:"referrers,omitempty"`
	Network   *string   `json:"network,omitempty"`
}

// NewReferrersByDomain instantiates a new ReferrersByDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferrersByDomain() *ReferrersByDomain {
	this := ReferrersByDomain{}
	return &this
}

// NewReferrersByDomainWithDefaults instantiates a new ReferrersByDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferrersByDomainWithDefaults() *ReferrersByDomain {
	this := ReferrersByDomain{}
	return &this
}

// GetReferrers returns the Referrers field value if set, zero value otherwise.
func (o *ReferrersByDomain) GetReferrers() []Metric {
	if o == nil || o.Referrers == nil {
		var ret []Metric
		return ret
	}
	return *o.Referrers
}

// GetReferrersOk returns a tuple with the Referrers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferrersByDomain) GetReferrersOk() (*[]Metric, bool) {
	if o == nil || o.Referrers == nil {
		return nil, false
	}
	return o.Referrers, true
}

// HasReferrers returns a boolean if a field has been set.
func (o *ReferrersByDomain) HasReferrers() bool {
	if o != nil && o.Referrers != nil {
		return true
	}

	return false
}

// SetReferrers gets a reference to the given []Metric and assigns it to the Referrers field.
func (o *ReferrersByDomain) SetReferrers(v []Metric) {
	o.Referrers = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *ReferrersByDomain) GetNetwork() string {
	if o == nil || o.Network == nil {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferrersByDomain) GetNetworkOk() (*string, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *ReferrersByDomain) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *ReferrersByDomain) SetNetwork(v string) {
	o.Network = &v
}

func (o ReferrersByDomain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Referrers != nil {
		toSerialize["referrers"] = o.Referrers
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	return json.Marshal(toSerialize)
}

type NullableReferrersByDomain struct {
	value *ReferrersByDomain
	isSet bool
}

func (v NullableReferrersByDomain) Get() *ReferrersByDomain {
	return v.value
}

func (v *NullableReferrersByDomain) Set(val *ReferrersByDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableReferrersByDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableReferrersByDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferrersByDomain(val *ReferrersByDomain) *NullableReferrersByDomain {
	return &NullableReferrersByDomain{value: val, isSet: true}
}

func (v NullableReferrersByDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferrersByDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}