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

// LaunchpadButtonRequest struct for LaunchpadButtonRequest
type LaunchpadButtonRequest struct {
	Bitlink       *string `json:"bitlink,omitempty"`
	Title         *string `json:"title,omitempty"`
	IsActive      *bool   `json:"is_active,omitempty"`
	IsPinned      *bool   `json:"is_pinned,omitempty"`
	ScheduleEnd   *string `json:"schedule_end,omitempty"`
	LongUrl       *string `json:"long_url,omitempty"`
	Domain        *string `json:"domain,omitempty"`
	ScheduleStart *string `json:"schedule_start,omitempty"`
}

// NewLaunchpadButtonRequest instantiates a new LaunchpadButtonRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLaunchpadButtonRequest() *LaunchpadButtonRequest {
	this := LaunchpadButtonRequest{}
	return &this
}

// NewLaunchpadButtonRequestWithDefaults instantiates a new LaunchpadButtonRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLaunchpadButtonRequestWithDefaults() *LaunchpadButtonRequest {
	this := LaunchpadButtonRequest{}
	return &this
}

// GetBitlink returns the Bitlink field value if set, zero value otherwise.
func (o *LaunchpadButtonRequest) GetBitlink() string {
	if o == nil || o.Bitlink == nil {
		var ret string
		return ret
	}
	return *o.Bitlink
}

// GetBitlinkOk returns a tuple with the Bitlink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadButtonRequest) GetBitlinkOk() (*string, bool) {
	if o == nil || o.Bitlink == nil {
		return nil, false
	}
	return o.Bitlink, true
}

// HasBitlink returns a boolean if a field has been set.
func (o *LaunchpadButtonRequest) HasBitlink() bool {
	if o != nil && o.Bitlink != nil {
		return true
	}

	return false
}

// SetBitlink gets a reference to the given string and assigns it to the Bitlink field.
func (o *LaunchpadButtonRequest) SetBitlink(v string) {
	o.Bitlink = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *LaunchpadButtonRequest) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadButtonRequest) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *LaunchpadButtonRequest) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *LaunchpadButtonRequest) SetTitle(v string) {
	o.Title = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *LaunchpadButtonRequest) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadButtonRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *LaunchpadButtonRequest) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *LaunchpadButtonRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsPinned returns the IsPinned field value if set, zero value otherwise.
func (o *LaunchpadButtonRequest) GetIsPinned() bool {
	if o == nil || o.IsPinned == nil {
		var ret bool
		return ret
	}
	return *o.IsPinned
}

// GetIsPinnedOk returns a tuple with the IsPinned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadButtonRequest) GetIsPinnedOk() (*bool, bool) {
	if o == nil || o.IsPinned == nil {
		return nil, false
	}
	return o.IsPinned, true
}

// HasIsPinned returns a boolean if a field has been set.
func (o *LaunchpadButtonRequest) HasIsPinned() bool {
	if o != nil && o.IsPinned != nil {
		return true
	}

	return false
}

// SetIsPinned gets a reference to the given bool and assigns it to the IsPinned field.
func (o *LaunchpadButtonRequest) SetIsPinned(v bool) {
	o.IsPinned = &v
}

// GetScheduleEnd returns the ScheduleEnd field value if set, zero value otherwise.
func (o *LaunchpadButtonRequest) GetScheduleEnd() string {
	if o == nil || o.ScheduleEnd == nil {
		var ret string
		return ret
	}
	return *o.ScheduleEnd
}

// GetScheduleEndOk returns a tuple with the ScheduleEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadButtonRequest) GetScheduleEndOk() (*string, bool) {
	if o == nil || o.ScheduleEnd == nil {
		return nil, false
	}
	return o.ScheduleEnd, true
}

// HasScheduleEnd returns a boolean if a field has been set.
func (o *LaunchpadButtonRequest) HasScheduleEnd() bool {
	if o != nil && o.ScheduleEnd != nil {
		return true
	}

	return false
}

// SetScheduleEnd gets a reference to the given string and assigns it to the ScheduleEnd field.
func (o *LaunchpadButtonRequest) SetScheduleEnd(v string) {
	o.ScheduleEnd = &v
}

// GetLongUrl returns the LongUrl field value if set, zero value otherwise.
func (o *LaunchpadButtonRequest) GetLongUrl() string {
	if o == nil || o.LongUrl == nil {
		var ret string
		return ret
	}
	return *o.LongUrl
}

// GetLongUrlOk returns a tuple with the LongUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadButtonRequest) GetLongUrlOk() (*string, bool) {
	if o == nil || o.LongUrl == nil {
		return nil, false
	}
	return o.LongUrl, true
}

// HasLongUrl returns a boolean if a field has been set.
func (o *LaunchpadButtonRequest) HasLongUrl() bool {
	if o != nil && o.LongUrl != nil {
		return true
	}

	return false
}

// SetLongUrl gets a reference to the given string and assigns it to the LongUrl field.
func (o *LaunchpadButtonRequest) SetLongUrl(v string) {
	o.LongUrl = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *LaunchpadButtonRequest) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadButtonRequest) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *LaunchpadButtonRequest) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *LaunchpadButtonRequest) SetDomain(v string) {
	o.Domain = &v
}

// GetScheduleStart returns the ScheduleStart field value if set, zero value otherwise.
func (o *LaunchpadButtonRequest) GetScheduleStart() string {
	if o == nil || o.ScheduleStart == nil {
		var ret string
		return ret
	}
	return *o.ScheduleStart
}

// GetScheduleStartOk returns a tuple with the ScheduleStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadButtonRequest) GetScheduleStartOk() (*string, bool) {
	if o == nil || o.ScheduleStart == nil {
		return nil, false
	}
	return o.ScheduleStart, true
}

// HasScheduleStart returns a boolean if a field has been set.
func (o *LaunchpadButtonRequest) HasScheduleStart() bool {
	if o != nil && o.ScheduleStart != nil {
		return true
	}

	return false
}

// SetScheduleStart gets a reference to the given string and assigns it to the ScheduleStart field.
func (o *LaunchpadButtonRequest) SetScheduleStart(v string) {
	o.ScheduleStart = &v
}

func (o LaunchpadButtonRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bitlink != nil {
		toSerialize["bitlink"] = o.Bitlink
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.IsActive != nil {
		toSerialize["is_active"] = o.IsActive
	}
	if o.IsPinned != nil {
		toSerialize["is_pinned"] = o.IsPinned
	}
	if o.ScheduleEnd != nil {
		toSerialize["schedule_end"] = o.ScheduleEnd
	}
	if o.LongUrl != nil {
		toSerialize["long_url"] = o.LongUrl
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.ScheduleStart != nil {
		toSerialize["schedule_start"] = o.ScheduleStart
	}
	return json.Marshal(toSerialize)
}

type NullableLaunchpadButtonRequest struct {
	value *LaunchpadButtonRequest
	isSet bool
}

func (v NullableLaunchpadButtonRequest) Get() *LaunchpadButtonRequest {
	return v.value
}

func (v *NullableLaunchpadButtonRequest) Set(val *LaunchpadButtonRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLaunchpadButtonRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLaunchpadButtonRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLaunchpadButtonRequest(val *LaunchpadButtonRequest) *NullableLaunchpadButtonRequest {
	return &NullableLaunchpadButtonRequest{value: val, isSet: true}
}

func (v NullableLaunchpadButtonRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLaunchpadButtonRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
