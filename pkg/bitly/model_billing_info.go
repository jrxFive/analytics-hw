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

// BillingInfo struct for BillingInfo
type BillingInfo struct {
	BillingAddress *BillingAddress `json:"billing_address,omitempty"`
	BillingContact *BillingContact `json:"billing_contact,omitempty"`
}

// NewBillingInfo instantiates a new BillingInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingInfo() *BillingInfo {
	this := BillingInfo{}
	return &this
}

// NewBillingInfoWithDefaults instantiates a new BillingInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingInfoWithDefaults() *BillingInfo {
	this := BillingInfo{}
	return &this
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *BillingInfo) GetBillingAddress() BillingAddress {
	if o == nil || o.BillingAddress == nil {
		var ret BillingAddress
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInfo) GetBillingAddressOk() (*BillingAddress, bool) {
	if o == nil || o.BillingAddress == nil {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *BillingInfo) HasBillingAddress() bool {
	if o != nil && o.BillingAddress != nil {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given BillingAddress and assigns it to the BillingAddress field.
func (o *BillingInfo) SetBillingAddress(v BillingAddress) {
	o.BillingAddress = &v
}

// GetBillingContact returns the BillingContact field value if set, zero value otherwise.
func (o *BillingInfo) GetBillingContact() BillingContact {
	if o == nil || o.BillingContact == nil {
		var ret BillingContact
		return ret
	}
	return *o.BillingContact
}

// GetBillingContactOk returns a tuple with the BillingContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInfo) GetBillingContactOk() (*BillingContact, bool) {
	if o == nil || o.BillingContact == nil {
		return nil, false
	}
	return o.BillingContact, true
}

// HasBillingContact returns a boolean if a field has been set.
func (o *BillingInfo) HasBillingContact() bool {
	if o != nil && o.BillingContact != nil {
		return true
	}

	return false
}

// SetBillingContact gets a reference to the given BillingContact and assigns it to the BillingContact field.
func (o *BillingInfo) SetBillingContact(v BillingContact) {
	o.BillingContact = &v
}

func (o BillingInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BillingAddress != nil {
		toSerialize["billing_address"] = o.BillingAddress
	}
	if o.BillingContact != nil {
		toSerialize["billing_contact"] = o.BillingContact
	}
	return json.Marshal(toSerialize)
}

type NullableBillingInfo struct {
	value *BillingInfo
	isSet bool
}

func (v NullableBillingInfo) Get() *BillingInfo {
	return v.value
}

func (v *NullableBillingInfo) Set(val *BillingInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingInfo(val *BillingInfo) *NullableBillingInfo {
	return &NullableBillingInfo{value: val, isSet: true}
}

func (v NullableBillingInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
