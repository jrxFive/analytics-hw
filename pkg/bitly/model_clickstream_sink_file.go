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
	"time"
)

// ClickstreamSinkFile struct for ClickstreamSinkFile
type ClickstreamSinkFile struct {
	// start of period
	DateFrom *time.Time `json:"date_from,omitempty"`
	// end of period
	DateTo *time.Time `json:"date_to,omitempty"`
	// Unique id of the file
	FileId *string `json:"file_id,omitempty"`
	// use this url with follow redirects to download the file
	DownloadUrl *string `json:"download_url,omitempty"`
	// time of creation
	Created *time.Time `json:"created,omitempty"`
}

// NewClickstreamSinkFile instantiates a new ClickstreamSinkFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClickstreamSinkFile() *ClickstreamSinkFile {
	this := ClickstreamSinkFile{}
	return &this
}

// NewClickstreamSinkFileWithDefaults instantiates a new ClickstreamSinkFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClickstreamSinkFileWithDefaults() *ClickstreamSinkFile {
	this := ClickstreamSinkFile{}
	return &this
}

// GetDateFrom returns the DateFrom field value if set, zero value otherwise.
func (o *ClickstreamSinkFile) GetDateFrom() time.Time {
	if o == nil || o.DateFrom == nil {
		var ret time.Time
		return ret
	}
	return *o.DateFrom
}

// GetDateFromOk returns a tuple with the DateFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickstreamSinkFile) GetDateFromOk() (*time.Time, bool) {
	if o == nil || o.DateFrom == nil {
		return nil, false
	}
	return o.DateFrom, true
}

// HasDateFrom returns a boolean if a field has been set.
func (o *ClickstreamSinkFile) HasDateFrom() bool {
	if o != nil && o.DateFrom != nil {
		return true
	}

	return false
}

// SetDateFrom gets a reference to the given time.Time and assigns it to the DateFrom field.
func (o *ClickstreamSinkFile) SetDateFrom(v time.Time) {
	o.DateFrom = &v
}

// GetDateTo returns the DateTo field value if set, zero value otherwise.
func (o *ClickstreamSinkFile) GetDateTo() time.Time {
	if o == nil || o.DateTo == nil {
		var ret time.Time
		return ret
	}
	return *o.DateTo
}

// GetDateToOk returns a tuple with the DateTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickstreamSinkFile) GetDateToOk() (*time.Time, bool) {
	if o == nil || o.DateTo == nil {
		return nil, false
	}
	return o.DateTo, true
}

// HasDateTo returns a boolean if a field has been set.
func (o *ClickstreamSinkFile) HasDateTo() bool {
	if o != nil && o.DateTo != nil {
		return true
	}

	return false
}

// SetDateTo gets a reference to the given time.Time and assigns it to the DateTo field.
func (o *ClickstreamSinkFile) SetDateTo(v time.Time) {
	o.DateTo = &v
}

// GetFileId returns the FileId field value if set, zero value otherwise.
func (o *ClickstreamSinkFile) GetFileId() string {
	if o == nil || o.FileId == nil {
		var ret string
		return ret
	}
	return *o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickstreamSinkFile) GetFileIdOk() (*string, bool) {
	if o == nil || o.FileId == nil {
		return nil, false
	}
	return o.FileId, true
}

// HasFileId returns a boolean if a field has been set.
func (o *ClickstreamSinkFile) HasFileId() bool {
	if o != nil && o.FileId != nil {
		return true
	}

	return false
}

// SetFileId gets a reference to the given string and assigns it to the FileId field.
func (o *ClickstreamSinkFile) SetFileId(v string) {
	o.FileId = &v
}

// GetDownloadUrl returns the DownloadUrl field value if set, zero value otherwise.
func (o *ClickstreamSinkFile) GetDownloadUrl() string {
	if o == nil || o.DownloadUrl == nil {
		var ret string
		return ret
	}
	return *o.DownloadUrl
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickstreamSinkFile) GetDownloadUrlOk() (*string, bool) {
	if o == nil || o.DownloadUrl == nil {
		return nil, false
	}
	return o.DownloadUrl, true
}

// HasDownloadUrl returns a boolean if a field has been set.
func (o *ClickstreamSinkFile) HasDownloadUrl() bool {
	if o != nil && o.DownloadUrl != nil {
		return true
	}

	return false
}

// SetDownloadUrl gets a reference to the given string and assigns it to the DownloadUrl field.
func (o *ClickstreamSinkFile) SetDownloadUrl(v string) {
	o.DownloadUrl = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ClickstreamSinkFile) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickstreamSinkFile) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ClickstreamSinkFile) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ClickstreamSinkFile) SetCreated(v time.Time) {
	o.Created = &v
}

func (o ClickstreamSinkFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DateFrom != nil {
		toSerialize["date_from"] = o.DateFrom
	}
	if o.DateTo != nil {
		toSerialize["date_to"] = o.DateTo
	}
	if o.FileId != nil {
		toSerialize["file_id"] = o.FileId
	}
	if o.DownloadUrl != nil {
		toSerialize["download_url"] = o.DownloadUrl
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	return json.Marshal(toSerialize)
}

type NullableClickstreamSinkFile struct {
	value *ClickstreamSinkFile
	isSet bool
}

func (v NullableClickstreamSinkFile) Get() *ClickstreamSinkFile {
	return v.value
}

func (v *NullableClickstreamSinkFile) Set(val *ClickstreamSinkFile) {
	v.value = val
	v.isSet = true
}

func (v NullableClickstreamSinkFile) IsSet() bool {
	return v.isSet
}

func (v *NullableClickstreamSinkFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClickstreamSinkFile(val *ClickstreamSinkFile) *NullableClickstreamSinkFile {
	return &NullableClickstreamSinkFile{value: val, isSet: true}
}

func (v NullableClickstreamSinkFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClickstreamSinkFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
