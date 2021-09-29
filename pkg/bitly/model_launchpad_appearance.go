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

// LaunchpadAppearance struct for LaunchpadAppearance
type LaunchpadAppearance struct {
	LaunchpadId           *string                    `json:"launchpad_id,omitempty"`
	AvatarPreference      *string                    `json:"avatar_preference,omitempty"`
	HideBitlyLogo         bool                       `json:"hide_bitly_logo"`
	DisplayName           *string                    `json:"display_name,omitempty"`
	AvatarBackgroundColor *string                    `json:"avatar_background_color,omitempty"`
	LaunchpadTextColor    *string                    `json:"launchpad_text_color,omitempty"`
	Description           *string                    `json:"description,omitempty"`
	ButtonAppearance      *LaunchpadButtonAppearance `json:"button_appearance,omitempty"`
	AvatarImageUrl        *string                    `json:"avatar_image_url,omitempty"`
	AvatarUploadImage     *string                    `json:"avatar_upload_image,omitempty"`
	StylePreference       *string                    `json:"style_preference,omitempty"`
	Background            *string                    `json:"background,omitempty"`
	ThemeId               *int32                     `json:"theme_id,omitempty"`
	AvatarTextColor       *string                    `json:"avatar_text_color,omitempty"`
	Layout                *string                    `json:"layout,omitempty"`
}

// NewLaunchpadAppearance instantiates a new LaunchpadAppearance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLaunchpadAppearance(hideBitlyLogo bool) *LaunchpadAppearance {
	this := LaunchpadAppearance{}
	this.HideBitlyLogo = hideBitlyLogo
	return &this
}

// NewLaunchpadAppearanceWithDefaults instantiates a new LaunchpadAppearance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLaunchpadAppearanceWithDefaults() *LaunchpadAppearance {
	this := LaunchpadAppearance{}
	return &this
}

// GetLaunchpadId returns the LaunchpadId field value if set, zero value otherwise.
func (o *LaunchpadAppearance) GetLaunchpadId() string {
	if o == nil || o.LaunchpadId == nil {
		var ret string
		return ret
	}
	return *o.LaunchpadId
}

// GetLaunchpadIdOk returns a tuple with the LaunchpadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadAppearance) GetLaunchpadIdOk() (*string, bool) {
	if o == nil || o.LaunchpadId == nil {
		return nil, false
	}
	return o.LaunchpadId, true
}

// HasLaunchpadId returns a boolean if a field has been set.
func (o *LaunchpadAppearance) HasLaunchpadId() bool {
	if o != nil && o.LaunchpadId != nil {
		return true
	}

	return false
}

// SetLaunchpadId gets a reference to the given string and assigns it to the LaunchpadId field.
func (o *LaunchpadAppearance) SetLaunchpadId(v string) {
	o.LaunchpadId = &v
}

// GetAvatarPreference returns the AvatarPreference field value if set, zero value otherwise.
func (o *LaunchpadAppearance) GetAvatarPreference() string {
	if o == nil || o.AvatarPreference == nil {
		var ret string
		return ret
	}
	return *o.AvatarPreference
}

// GetAvatarPreferenceOk returns a tuple with the AvatarPreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadAppearance) GetAvatarPreferenceOk() (*string, bool) {
	if o == nil || o.AvatarPreference == nil {
		return nil, false
	}
	return o.AvatarPreference, true
}

// HasAvatarPreference returns a boolean if a field has been set.
func (o *LaunchpadAppearance) HasAvatarPreference() bool {
	if o != nil && o.AvatarPreference != nil {
		return true
	}

	return false
}

// SetAvatarPreference gets a reference to the given string and assigns it to the AvatarPreference field.
func (o *LaunchpadAppearance) SetAvatarPreference(v string) {
	o.AvatarPreference = &v
}

// GetHideBitlyLogo returns the HideBitlyLogo field value
func (o *LaunchpadAppearance) GetHideBitlyLogo() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HideBitlyLogo
}

// GetHideBitlyLogoOk returns a tuple with the HideBitlyLogo field value
// and a boolean to check if the value has been set.
func (o *LaunchpadAppearance) GetHideBitlyLogoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HideBitlyLogo, true
}

// SetHideBitlyLogo sets field value
func (o *LaunchpadAppearance) SetHideBitlyLogo(v bool) {
	o.HideBitlyLogo = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *LaunchpadAppearance) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadAppearance) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *LaunchpadAppearance) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *LaunchpadAppearance) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetAvatarBackgroundColor returns the AvatarBackgroundColor field value if set, zero value otherwise.
func (o *LaunchpadAppearance) GetAvatarBackgroundColor() string {
	if o == nil || o.AvatarBackgroundColor == nil {
		var ret string
		return ret
	}
	return *o.AvatarBackgroundColor
}

// GetAvatarBackgroundColorOk returns a tuple with the AvatarBackgroundColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadAppearance) GetAvatarBackgroundColorOk() (*string, bool) {
	if o == nil || o.AvatarBackgroundColor == nil {
		return nil, false
	}
	return o.AvatarBackgroundColor, true
}

// HasAvatarBackgroundColor returns a boolean if a field has been set.
func (o *LaunchpadAppearance) HasAvatarBackgroundColor() bool {
	if o != nil && o.AvatarBackgroundColor != nil {
		return true
	}

	return false
}

// SetAvatarBackgroundColor gets a reference to the given string and assigns it to the AvatarBackgroundColor field.
func (o *LaunchpadAppearance) SetAvatarBackgroundColor(v string) {
	o.AvatarBackgroundColor = &v
}

// GetLaunchpadTextColor returns the LaunchpadTextColor field value if set, zero value otherwise.
func (o *LaunchpadAppearance) GetLaunchpadTextColor() string {
	if o == nil || o.LaunchpadTextColor == nil {
		var ret string
		return ret
	}
	return *o.LaunchpadTextColor
}

// GetLaunchpadTextColorOk returns a tuple with the LaunchpadTextColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadAppearance) GetLaunchpadTextColorOk() (*string, bool) {
	if o == nil || o.LaunchpadTextColor == nil {
		return nil, false
	}
	return o.LaunchpadTextColor, true
}

// HasLaunchpadTextColor returns a boolean if a field has been set.
func (o *LaunchpadAppearance) HasLaunchpadTextColor() bool {
	if o != nil && o.LaunchpadTextColor != nil {
		return true
	}

	return false
}

// SetLaunchpadTextColor gets a reference to the given string and assigns it to the LaunchpadTextColor field.
func (o *LaunchpadAppearance) SetLaunchpadTextColor(v string) {
	o.LaunchpadTextColor = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LaunchpadAppearance) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadAppearance) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LaunchpadAppearance) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LaunchpadAppearance) SetDescription(v string) {
	o.Description = &v
}

// GetButtonAppearance returns the ButtonAppearance field value if set, zero value otherwise.
func (o *LaunchpadAppearance) GetButtonAppearance() LaunchpadButtonAppearance {
	if o == nil || o.ButtonAppearance == nil {
		var ret LaunchpadButtonAppearance
		return ret
	}
	return *o.ButtonAppearance
}

// GetButtonAppearanceOk returns a tuple with the ButtonAppearance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadAppearance) GetButtonAppearanceOk() (*LaunchpadButtonAppearance, bool) {
	if o == nil || o.ButtonAppearance == nil {
		return nil, false
	}
	return o.ButtonAppearance, true
}

// HasButtonAppearance returns a boolean if a field has been set.
func (o *LaunchpadAppearance) HasButtonAppearance() bool {
	if o != nil && o.ButtonAppearance != nil {
		return true
	}

	return false
}

// SetButtonAppearance gets a reference to the given LaunchpadButtonAppearance and assigns it to the ButtonAppearance field.
func (o *LaunchpadAppearance) SetButtonAppearance(v LaunchpadButtonAppearance) {
	o.ButtonAppearance = &v
}

// GetAvatarImageUrl returns the AvatarImageUrl field value if set, zero value otherwise.
func (o *LaunchpadAppearance) GetAvatarImageUrl() string {
	if o == nil || o.AvatarImageUrl == nil {
		var ret string
		return ret
	}
	return *o.AvatarImageUrl
}

// GetAvatarImageUrlOk returns a tuple with the AvatarImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadAppearance) GetAvatarImageUrlOk() (*string, bool) {
	if o == nil || o.AvatarImageUrl == nil {
		return nil, false
	}
	return o.AvatarImageUrl, true
}

// HasAvatarImageUrl returns a boolean if a field has been set.
func (o *LaunchpadAppearance) HasAvatarImageUrl() bool {
	if o != nil && o.AvatarImageUrl != nil {
		return true
	}

	return false
}

// SetAvatarImageUrl gets a reference to the given string and assigns it to the AvatarImageUrl field.
func (o *LaunchpadAppearance) SetAvatarImageUrl(v string) {
	o.AvatarImageUrl = &v
}

// GetAvatarUploadImage returns the AvatarUploadImage field value if set, zero value otherwise.
func (o *LaunchpadAppearance) GetAvatarUploadImage() string {
	if o == nil || o.AvatarUploadImage == nil {
		var ret string
		return ret
	}
	return *o.AvatarUploadImage
}

// GetAvatarUploadImageOk returns a tuple with the AvatarUploadImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadAppearance) GetAvatarUploadImageOk() (*string, bool) {
	if o == nil || o.AvatarUploadImage == nil {
		return nil, false
	}
	return o.AvatarUploadImage, true
}

// HasAvatarUploadImage returns a boolean if a field has been set.
func (o *LaunchpadAppearance) HasAvatarUploadImage() bool {
	if o != nil && o.AvatarUploadImage != nil {
		return true
	}

	return false
}

// SetAvatarUploadImage gets a reference to the given string and assigns it to the AvatarUploadImage field.
func (o *LaunchpadAppearance) SetAvatarUploadImage(v string) {
	o.AvatarUploadImage = &v
}

// GetStylePreference returns the StylePreference field value if set, zero value otherwise.
func (o *LaunchpadAppearance) GetStylePreference() string {
	if o == nil || o.StylePreference == nil {
		var ret string
		return ret
	}
	return *o.StylePreference
}

// GetStylePreferenceOk returns a tuple with the StylePreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadAppearance) GetStylePreferenceOk() (*string, bool) {
	if o == nil || o.StylePreference == nil {
		return nil, false
	}
	return o.StylePreference, true
}

// HasStylePreference returns a boolean if a field has been set.
func (o *LaunchpadAppearance) HasStylePreference() bool {
	if o != nil && o.StylePreference != nil {
		return true
	}

	return false
}

// SetStylePreference gets a reference to the given string and assigns it to the StylePreference field.
func (o *LaunchpadAppearance) SetStylePreference(v string) {
	o.StylePreference = &v
}

// GetBackground returns the Background field value if set, zero value otherwise.
func (o *LaunchpadAppearance) GetBackground() string {
	if o == nil || o.Background == nil {
		var ret string
		return ret
	}
	return *o.Background
}

// GetBackgroundOk returns a tuple with the Background field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadAppearance) GetBackgroundOk() (*string, bool) {
	if o == nil || o.Background == nil {
		return nil, false
	}
	return o.Background, true
}

// HasBackground returns a boolean if a field has been set.
func (o *LaunchpadAppearance) HasBackground() bool {
	if o != nil && o.Background != nil {
		return true
	}

	return false
}

// SetBackground gets a reference to the given string and assigns it to the Background field.
func (o *LaunchpadAppearance) SetBackground(v string) {
	o.Background = &v
}

// GetThemeId returns the ThemeId field value if set, zero value otherwise.
func (o *LaunchpadAppearance) GetThemeId() int32 {
	if o == nil || o.ThemeId == nil {
		var ret int32
		return ret
	}
	return *o.ThemeId
}

// GetThemeIdOk returns a tuple with the ThemeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadAppearance) GetThemeIdOk() (*int32, bool) {
	if o == nil || o.ThemeId == nil {
		return nil, false
	}
	return o.ThemeId, true
}

// HasThemeId returns a boolean if a field has been set.
func (o *LaunchpadAppearance) HasThemeId() bool {
	if o != nil && o.ThemeId != nil {
		return true
	}

	return false
}

// SetThemeId gets a reference to the given int32 and assigns it to the ThemeId field.
func (o *LaunchpadAppearance) SetThemeId(v int32) {
	o.ThemeId = &v
}

// GetAvatarTextColor returns the AvatarTextColor field value if set, zero value otherwise.
func (o *LaunchpadAppearance) GetAvatarTextColor() string {
	if o == nil || o.AvatarTextColor == nil {
		var ret string
		return ret
	}
	return *o.AvatarTextColor
}

// GetAvatarTextColorOk returns a tuple with the AvatarTextColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadAppearance) GetAvatarTextColorOk() (*string, bool) {
	if o == nil || o.AvatarTextColor == nil {
		return nil, false
	}
	return o.AvatarTextColor, true
}

// HasAvatarTextColor returns a boolean if a field has been set.
func (o *LaunchpadAppearance) HasAvatarTextColor() bool {
	if o != nil && o.AvatarTextColor != nil {
		return true
	}

	return false
}

// SetAvatarTextColor gets a reference to the given string and assigns it to the AvatarTextColor field.
func (o *LaunchpadAppearance) SetAvatarTextColor(v string) {
	o.AvatarTextColor = &v
}

// GetLayout returns the Layout field value if set, zero value otherwise.
func (o *LaunchpadAppearance) GetLayout() string {
	if o == nil || o.Layout == nil {
		var ret string
		return ret
	}
	return *o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LaunchpadAppearance) GetLayoutOk() (*string, bool) {
	if o == nil || o.Layout == nil {
		return nil, false
	}
	return o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *LaunchpadAppearance) HasLayout() bool {
	if o != nil && o.Layout != nil {
		return true
	}

	return false
}

// SetLayout gets a reference to the given string and assigns it to the Layout field.
func (o *LaunchpadAppearance) SetLayout(v string) {
	o.Layout = &v
}

func (o LaunchpadAppearance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LaunchpadId != nil {
		toSerialize["launchpad_id"] = o.LaunchpadId
	}
	if o.AvatarPreference != nil {
		toSerialize["avatar_preference"] = o.AvatarPreference
	}
	if true {
		toSerialize["hide_bitly_logo"] = o.HideBitlyLogo
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.AvatarBackgroundColor != nil {
		toSerialize["avatar_background_color"] = o.AvatarBackgroundColor
	}
	if o.LaunchpadTextColor != nil {
		toSerialize["launchpad_text_color"] = o.LaunchpadTextColor
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ButtonAppearance != nil {
		toSerialize["button_appearance"] = o.ButtonAppearance
	}
	if o.AvatarImageUrl != nil {
		toSerialize["avatar_image_url"] = o.AvatarImageUrl
	}
	if o.AvatarUploadImage != nil {
		toSerialize["avatar_upload_image"] = o.AvatarUploadImage
	}
	if o.StylePreference != nil {
		toSerialize["style_preference"] = o.StylePreference
	}
	if o.Background != nil {
		toSerialize["background"] = o.Background
	}
	if o.ThemeId != nil {
		toSerialize["theme_id"] = o.ThemeId
	}
	if o.AvatarTextColor != nil {
		toSerialize["avatar_text_color"] = o.AvatarTextColor
	}
	if o.Layout != nil {
		toSerialize["layout"] = o.Layout
	}
	return json.Marshal(toSerialize)
}

type NullableLaunchpadAppearance struct {
	value *LaunchpadAppearance
	isSet bool
}

func (v NullableLaunchpadAppearance) Get() *LaunchpadAppearance {
	return v.value
}

func (v *NullableLaunchpadAppearance) Set(val *LaunchpadAppearance) {
	v.value = val
	v.isSet = true
}

func (v NullableLaunchpadAppearance) IsSet() bool {
	return v.isSet
}

func (v *NullableLaunchpadAppearance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLaunchpadAppearance(val *LaunchpadAppearance) *NullableLaunchpadAppearance {
	return &NullableLaunchpadAppearance{value: val, isSet: true}
}

func (v NullableLaunchpadAppearance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLaunchpadAppearance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
