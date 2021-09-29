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

// BulkAddResponse struct for BulkAddResponse
type BulkAddResponse struct {
	Message *string              `json:"message,omitempty"`
	Data    *BulkAddResponseData `json:"data,omitempty"`
}

// NewBulkAddResponse instantiates a new BulkAddResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkAddResponse() *BulkAddResponse {
	this := BulkAddResponse{}
	return &this
}

// NewBulkAddResponseWithDefaults instantiates a new BulkAddResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkAddResponseWithDefaults() *BulkAddResponse {
	this := BulkAddResponse{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *BulkAddResponse) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkAddResponse) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *BulkAddResponse) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *BulkAddResponse) SetMessage(v string) {
	o.Message = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BulkAddResponse) GetData() BulkAddResponseData {
	if o == nil || o.Data == nil {
		var ret BulkAddResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkAddResponse) GetDataOk() (*BulkAddResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BulkAddResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given BulkAddResponseData and assigns it to the Data field.
func (o *BulkAddResponse) SetData(v BulkAddResponseData) {
	o.Data = &v
}

func (o BulkAddResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableBulkAddResponse struct {
	value *BulkAddResponse
	isSet bool
}

func (v NullableBulkAddResponse) Get() *BulkAddResponse {
	return v.value
}

func (v *NullableBulkAddResponse) Set(val *BulkAddResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkAddResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkAddResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkAddResponse(val *BulkAddResponse) *NullableBulkAddResponse {
	return &NullableBulkAddResponse{value: val, isSet: true}
}

func (v NullableBulkAddResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkAddResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
