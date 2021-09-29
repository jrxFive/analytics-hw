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

// LaunchpadClicks struct for LaunchpadClicks
type LaunchpadClicks struct {
	Domain    *string `json:"domain,omitempty"`
	BitlinkId *string `json:"bitlink_id,omitempty"`
	Keyword   *string `json:"keyword,omitempty"`
	Title     *string `json:"title,omitempty"`
	IsActive  *bool   `json:"is_active,omitempty"`
	LongUrl   *string `json:"long_url,omitempty"`
	Date      *string `json:"date,omitempty"`
	Clicks    *int32  `json:"clicks,omitempty"`
}

// NewLaunchpadClicks instantiates a new LaunchpadClicks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLaunchpadClicks() *LaunchpadClicks {
	this := LaunchpadClicks{}
	return &this
}

// NewLaunchpadClicksWithDefaults instantiates a new LaunchpadClicks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLaunchpadClicksWithDefaults() *LaunchpadClicks {
	this := LaunchpadClicks{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *LaunchpadClicks) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadClicks) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *LaunchpadClicks) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *LaunchpadClicks) SetDomain(v string) {
	o.Domain = &v
}

// GetBitlinkId returns the BitlinkId field value if set, zero value otherwise.
func (o *LaunchpadClicks) GetBitlinkId() string {
	if o == nil || o.BitlinkId == nil {
		var ret string
		return ret
	}
	return *o.BitlinkId
}

// GetBitlinkIdOk returns a tuple with the BitlinkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadClicks) GetBitlinkIdOk() (*string, bool) {
	if o == nil || o.BitlinkId == nil {
		return nil, false
	}
	return o.BitlinkId, true
}

// HasBitlinkId returns a boolean if a field has been set.
func (o *LaunchpadClicks) HasBitlinkId() bool {
	if o != nil && o.BitlinkId != nil {
		return true
	}

	return false
}

// SetBitlinkId gets a reference to the given string and assigns it to the BitlinkId field.
func (o *LaunchpadClicks) SetBitlinkId(v string) {
	o.BitlinkId = &v
}

// GetKeyword returns the Keyword field value if set, zero value otherwise.
func (o *LaunchpadClicks) GetKeyword() string {
	if o == nil || o.Keyword == nil {
		var ret string
		return ret
	}
	return *o.Keyword
}

// GetKeywordOk returns a tuple with the Keyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadClicks) GetKeywordOk() (*string, bool) {
	if o == nil || o.Keyword == nil {
		return nil, false
	}
	return o.Keyword, true
}

// HasKeyword returns a boolean if a field has been set.
func (o *LaunchpadClicks) HasKeyword() bool {
	if o != nil && o.Keyword != nil {
		return true
	}

	return false
}

// SetKeyword gets a reference to the given string and assigns it to the Keyword field.
func (o *LaunchpadClicks) SetKeyword(v string) {
	o.Keyword = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *LaunchpadClicks) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadClicks) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *LaunchpadClicks) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *LaunchpadClicks) SetTitle(v string) {
	o.Title = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *LaunchpadClicks) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadClicks) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *LaunchpadClicks) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *LaunchpadClicks) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetLongUrl returns the LongUrl field value if set, zero value otherwise.
func (o *LaunchpadClicks) GetLongUrl() string {
	if o == nil || o.LongUrl == nil {
		var ret string
		return ret
	}
	return *o.LongUrl
}

// GetLongUrlOk returns a tuple with the LongUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadClicks) GetLongUrlOk() (*string, bool) {
	if o == nil || o.LongUrl == nil {
		return nil, false
	}
	return o.LongUrl, true
}

// HasLongUrl returns a boolean if a field has been set.
func (o *LaunchpadClicks) HasLongUrl() bool {
	if o != nil && o.LongUrl != nil {
		return true
	}

	return false
}

// SetLongUrl gets a reference to the given string and assigns it to the LongUrl field.
func (o *LaunchpadClicks) SetLongUrl(v string) {
	o.LongUrl = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *LaunchpadClicks) GetDate() string {
	if o == nil || o.Date == nil {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadClicks) GetDateOk() (*string, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *LaunchpadClicks) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *LaunchpadClicks) SetDate(v string) {
	o.Date = &v
}

// GetClicks returns the Clicks field value if set, zero value otherwise.
func (o *LaunchpadClicks) GetClicks() int32 {
	if o == nil || o.Clicks == nil {
		var ret int32
		return ret
	}
	return *o.Clicks
}

// GetClicksOk returns a tuple with the Clicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadClicks) GetClicksOk() (*int32, bool) {
	if o == nil || o.Clicks == nil {
		return nil, false
	}
	return o.Clicks, true
}

// HasClicks returns a boolean if a field has been set.
func (o *LaunchpadClicks) HasClicks() bool {
	if o != nil && o.Clicks != nil {
		return true
	}

	return false
}

// SetClicks gets a reference to the given int32 and assigns it to the Clicks field.
func (o *LaunchpadClicks) SetClicks(v int32) {
	o.Clicks = &v
}

func (o LaunchpadClicks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.BitlinkId != nil {
		toSerialize["bitlink_id"] = o.BitlinkId
	}
	if o.Keyword != nil {
		toSerialize["keyword"] = o.Keyword
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.IsActive != nil {
		toSerialize["is_active"] = o.IsActive
	}
	if o.LongUrl != nil {
		toSerialize["long_url"] = o.LongUrl
	}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.Clicks != nil {
		toSerialize["clicks"] = o.Clicks
	}
	return json.Marshal(toSerialize)
}

type NullableLaunchpadClicks struct {
	value *LaunchpadClicks
	isSet bool
}

func (v NullableLaunchpadClicks) Get() *LaunchpadClicks {
	return v.value
}

func (v *NullableLaunchpadClicks) Set(val *LaunchpadClicks) {
	v.value = val
	v.isSet = true
}

func (v NullableLaunchpadClicks) IsSet() bool {
	return v.isSet
}

func (v *NullableLaunchpadClicks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLaunchpadClicks(val *LaunchpadClicks) *NullableLaunchpadClicks {
	return &NullableLaunchpadClicks{value: val, isSet: true}
}

func (v NullableLaunchpadClicks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLaunchpadClicks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
