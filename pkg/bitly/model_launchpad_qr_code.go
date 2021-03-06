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

// LaunchpadQRCode struct for LaunchpadQRCode
type LaunchpadQRCode struct {
	LaunchpadId *string `json:"launchpad_id,omitempty"`
	QrCode      *string `json:"qr_code,omitempty"`
	Link        *string `json:"link,omitempty"`
}

// NewLaunchpadQRCode instantiates a new LaunchpadQRCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLaunchpadQRCode() *LaunchpadQRCode {
	this := LaunchpadQRCode{}
	return &this
}

// NewLaunchpadQRCodeWithDefaults instantiates a new LaunchpadQRCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLaunchpadQRCodeWithDefaults() *LaunchpadQRCode {
	this := LaunchpadQRCode{}
	return &this
}

// GetLaunchpadId returns the LaunchpadId field value if set, zero value otherwise.
func (o *LaunchpadQRCode) GetLaunchpadId() string {
	if o == nil || o.LaunchpadId == nil {
		var ret string
		return ret
	}
	return *o.LaunchpadId
}

// GetLaunchpadIdOk returns a tuple with the LaunchpadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadQRCode) GetLaunchpadIdOk() (*string, bool) {
	if o == nil || o.LaunchpadId == nil {
		return nil, false
	}
	return o.LaunchpadId, true
}

// HasLaunchpadId returns a boolean if a field has been set.
func (o *LaunchpadQRCode) HasLaunchpadId() bool {
	if o != nil && o.LaunchpadId != nil {
		return true
	}

	return false
}

// SetLaunchpadId gets a reference to the given string and assigns it to the LaunchpadId field.
func (o *LaunchpadQRCode) SetLaunchpadId(v string) {
	o.LaunchpadId = &v
}

// GetQrCode returns the QrCode field value if set, zero value otherwise.
func (o *LaunchpadQRCode) GetQrCode() string {
	if o == nil || o.QrCode == nil {
		var ret string
		return ret
	}
	return *o.QrCode
}

// GetQrCodeOk returns a tuple with the QrCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadQRCode) GetQrCodeOk() (*string, bool) {
	if o == nil || o.QrCode == nil {
		return nil, false
	}
	return o.QrCode, true
}

// HasQrCode returns a boolean if a field has been set.
func (o *LaunchpadQRCode) HasQrCode() bool {
	if o != nil && o.QrCode != nil {
		return true
	}

	return false
}

// SetQrCode gets a reference to the given string and assigns it to the QrCode field.
func (o *LaunchpadQRCode) SetQrCode(v string) {
	o.QrCode = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *LaunchpadQRCode) GetLink() string {
	if o == nil || o.Link == nil {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadQRCode) GetLinkOk() (*string, bool) {
	if o == nil || o.Link == nil {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *LaunchpadQRCode) HasLink() bool {
	if o != nil && o.Link != nil {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *LaunchpadQRCode) SetLink(v string) {
	o.Link = &v
}

func (o LaunchpadQRCode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LaunchpadId != nil {
		toSerialize["launchpad_id"] = o.LaunchpadId
	}
	if o.QrCode != nil {
		toSerialize["qr_code"] = o.QrCode
	}
	if o.Link != nil {
		toSerialize["link"] = o.Link
	}
	return json.Marshal(toSerialize)
}

type NullableLaunchpadQRCode struct {
	value *LaunchpadQRCode
	isSet bool
}

func (v NullableLaunchpadQRCode) Get() *LaunchpadQRCode {
	return v.value
}

func (v *NullableLaunchpadQRCode) Set(val *LaunchpadQRCode) {
	v.value = val
	v.isSet = true
}

func (v NullableLaunchpadQRCode) IsSet() bool {
	return v.isSet
}

func (v *NullableLaunchpadQRCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLaunchpadQRCode(val *LaunchpadQRCode) *NullableLaunchpadQRCode {
	return &NullableLaunchpadQRCode{value: val, isSet: true}
}

func (v NullableLaunchpadQRCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLaunchpadQRCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
