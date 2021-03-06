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

// UnprocessableEntity UNPROCESSABLE_ENTITY
type UnprocessableEntity struct {
	Message     *string       `json:"message,omitempty"`
	Errors      *[]FieldError `json:"errors,omitempty"`
	Resource    *string       `json:"resource,omitempty"`
	Description *string       `json:"description,omitempty"`
}

// NewUnprocessableEntity instantiates a new UnprocessableEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnprocessableEntity() *UnprocessableEntity {
	this := UnprocessableEntity{}
	return &this
}

// NewUnprocessableEntityWithDefaults instantiates a new UnprocessableEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnprocessableEntityWithDefaults() *UnprocessableEntity {
	this := UnprocessableEntity{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *UnprocessableEntity) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnprocessableEntity) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *UnprocessableEntity) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *UnprocessableEntity) SetMessage(v string) {
	o.Message = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *UnprocessableEntity) GetErrors() []FieldError {
	if o == nil || o.Errors == nil {
		var ret []FieldError
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnprocessableEntity) GetErrorsOk() (*[]FieldError, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *UnprocessableEntity) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []FieldError and assigns it to the Errors field.
func (o *UnprocessableEntity) SetErrors(v []FieldError) {
	o.Errors = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *UnprocessableEntity) GetResource() string {
	if o == nil || o.Resource == nil {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnprocessableEntity) GetResourceOk() (*string, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *UnprocessableEntity) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *UnprocessableEntity) SetResource(v string) {
	o.Resource = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UnprocessableEntity) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnprocessableEntity) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UnprocessableEntity) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UnprocessableEntity) SetDescription(v string) {
	o.Description = &v
}

func (o UnprocessableEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableUnprocessableEntity struct {
	value *UnprocessableEntity
	isSet bool
}

func (v NullableUnprocessableEntity) Get() *UnprocessableEntity {
	return v.value
}

func (v *NullableUnprocessableEntity) Set(val *UnprocessableEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableUnprocessableEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableUnprocessableEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnprocessableEntity(val *UnprocessableEntity) *NullableUnprocessableEntity {
	return &NullableUnprocessableEntity{value: val, isSet: true}
}

func (v NullableUnprocessableEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnprocessableEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
