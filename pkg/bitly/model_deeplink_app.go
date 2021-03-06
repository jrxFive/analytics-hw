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

// DeeplinkApp struct for DeeplinkApp
type DeeplinkApp struct {
	ThirdPartyAppId       string    `json:"third_party_app_id"`
	CreatedTs             *int32    `json:"created_ts,omitempty"`
	Scheme                *string   `json:"scheme,omitempty"`
	InstallUrl            string    `json:"install_url"`
	Name                  string    `json:"name"`
	LegacyAppId           *string   `json:"legacy_app_id,omitempty"`
	IconUrl               *string   `json:"icon_url,omitempty"`
	OrganizationGuid      *string   `json:"organization_guid,omitempty"`
	ModifiedTs            *int32    `json:"modified_ts,omitempty"`
	Guid                  *string   `json:"guid,omitempty"`
	AppleAppEntitlementId *string   `json:"apple_app_entitlement_id,omitempty"`
	AndroidSha256         *[]string `json:"android_sha256,omitempty"`
	Os                    string    `json:"os"`
}

// NewDeeplinkApp instantiates a new DeeplinkApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeeplinkApp(thirdPartyAppId string, installUrl string, name string, os string) *DeeplinkApp {
	this := DeeplinkApp{}
	this.ThirdPartyAppId = thirdPartyAppId
	this.InstallUrl = installUrl
	this.Name = name
	this.Os = os
	return &this
}

// NewDeeplinkAppWithDefaults instantiates a new DeeplinkApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeeplinkAppWithDefaults() *DeeplinkApp {
	this := DeeplinkApp{}
	return &this
}

// GetThirdPartyAppId returns the ThirdPartyAppId field value
func (o *DeeplinkApp) GetThirdPartyAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThirdPartyAppId
}

// GetThirdPartyAppIdOk returns a tuple with the ThirdPartyAppId field value
// and a boolean to check if the value has been set.
func (o *DeeplinkApp) GetThirdPartyAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThirdPartyAppId, true
}

// SetThirdPartyAppId sets field value
func (o *DeeplinkApp) SetThirdPartyAppId(v string) {
	o.ThirdPartyAppId = v
}

// GetCreatedTs returns the CreatedTs field value if set, zero value otherwise.
func (o *DeeplinkApp) GetCreatedTs() int32 {
	if o == nil || o.CreatedTs == nil {
		var ret int32
		return ret
	}
	return *o.CreatedTs
}

// GetCreatedTsOk returns a tuple with the CreatedTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeeplinkApp) GetCreatedTsOk() (*int32, bool) {
	if o == nil || o.CreatedTs == nil {
		return nil, false
	}
	return o.CreatedTs, true
}

// HasCreatedTs returns a boolean if a field has been set.
func (o *DeeplinkApp) HasCreatedTs() bool {
	if o != nil && o.CreatedTs != nil {
		return true
	}

	return false
}

// SetCreatedTs gets a reference to the given int32 and assigns it to the CreatedTs field.
func (o *DeeplinkApp) SetCreatedTs(v int32) {
	o.CreatedTs = &v
}

// GetScheme returns the Scheme field value if set, zero value otherwise.
func (o *DeeplinkApp) GetScheme() string {
	if o == nil || o.Scheme == nil {
		var ret string
		return ret
	}
	return *o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeeplinkApp) GetSchemeOk() (*string, bool) {
	if o == nil || o.Scheme == nil {
		return nil, false
	}
	return o.Scheme, true
}

// HasScheme returns a boolean if a field has been set.
func (o *DeeplinkApp) HasScheme() bool {
	if o != nil && o.Scheme != nil {
		return true
	}

	return false
}

// SetScheme gets a reference to the given string and assigns it to the Scheme field.
func (o *DeeplinkApp) SetScheme(v string) {
	o.Scheme = &v
}

// GetInstallUrl returns the InstallUrl field value
func (o *DeeplinkApp) GetInstallUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstallUrl
}

// GetInstallUrlOk returns a tuple with the InstallUrl field value
// and a boolean to check if the value has been set.
func (o *DeeplinkApp) GetInstallUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstallUrl, true
}

// SetInstallUrl sets field value
func (o *DeeplinkApp) SetInstallUrl(v string) {
	o.InstallUrl = v
}

// GetName returns the Name field value
func (o *DeeplinkApp) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeeplinkApp) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeeplinkApp) SetName(v string) {
	o.Name = v
}

// GetLegacyAppId returns the LegacyAppId field value if set, zero value otherwise.
func (o *DeeplinkApp) GetLegacyAppId() string {
	if o == nil || o.LegacyAppId == nil {
		var ret string
		return ret
	}
	return *o.LegacyAppId
}

// GetLegacyAppIdOk returns a tuple with the LegacyAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeeplinkApp) GetLegacyAppIdOk() (*string, bool) {
	if o == nil || o.LegacyAppId == nil {
		return nil, false
	}
	return o.LegacyAppId, true
}

// HasLegacyAppId returns a boolean if a field has been set.
func (o *DeeplinkApp) HasLegacyAppId() bool {
	if o != nil && o.LegacyAppId != nil {
		return true
	}

	return false
}

// SetLegacyAppId gets a reference to the given string and assigns it to the LegacyAppId field.
func (o *DeeplinkApp) SetLegacyAppId(v string) {
	o.LegacyAppId = &v
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise.
func (o *DeeplinkApp) GetIconUrl() string {
	if o == nil || o.IconUrl == nil {
		var ret string
		return ret
	}
	return *o.IconUrl
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeeplinkApp) GetIconUrlOk() (*string, bool) {
	if o == nil || o.IconUrl == nil {
		return nil, false
	}
	return o.IconUrl, true
}

// HasIconUrl returns a boolean if a field has been set.
func (o *DeeplinkApp) HasIconUrl() bool {
	if o != nil && o.IconUrl != nil {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given string and assigns it to the IconUrl field.
func (o *DeeplinkApp) SetIconUrl(v string) {
	o.IconUrl = &v
}

// GetOrganizationGuid returns the OrganizationGuid field value if set, zero value otherwise.
func (o *DeeplinkApp) GetOrganizationGuid() string {
	if o == nil || o.OrganizationGuid == nil {
		var ret string
		return ret
	}
	return *o.OrganizationGuid
}

// GetOrganizationGuidOk returns a tuple with the OrganizationGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeeplinkApp) GetOrganizationGuidOk() (*string, bool) {
	if o == nil || o.OrganizationGuid == nil {
		return nil, false
	}
	return o.OrganizationGuid, true
}

// HasOrganizationGuid returns a boolean if a field has been set.
func (o *DeeplinkApp) HasOrganizationGuid() bool {
	if o != nil && o.OrganizationGuid != nil {
		return true
	}

	return false
}

// SetOrganizationGuid gets a reference to the given string and assigns it to the OrganizationGuid field.
func (o *DeeplinkApp) SetOrganizationGuid(v string) {
	o.OrganizationGuid = &v
}

// GetModifiedTs returns the ModifiedTs field value if set, zero value otherwise.
func (o *DeeplinkApp) GetModifiedTs() int32 {
	if o == nil || o.ModifiedTs == nil {
		var ret int32
		return ret
	}
	return *o.ModifiedTs
}

// GetModifiedTsOk returns a tuple with the ModifiedTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeeplinkApp) GetModifiedTsOk() (*int32, bool) {
	if o == nil || o.ModifiedTs == nil {
		return nil, false
	}
	return o.ModifiedTs, true
}

// HasModifiedTs returns a boolean if a field has been set.
func (o *DeeplinkApp) HasModifiedTs() bool {
	if o != nil && o.ModifiedTs != nil {
		return true
	}

	return false
}

// SetModifiedTs gets a reference to the given int32 and assigns it to the ModifiedTs field.
func (o *DeeplinkApp) SetModifiedTs(v int32) {
	o.ModifiedTs = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *DeeplinkApp) GetGuid() string {
	if o == nil || o.Guid == nil {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeeplinkApp) GetGuidOk() (*string, bool) {
	if o == nil || o.Guid == nil {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *DeeplinkApp) HasGuid() bool {
	if o != nil && o.Guid != nil {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *DeeplinkApp) SetGuid(v string) {
	o.Guid = &v
}

// GetAppleAppEntitlementId returns the AppleAppEntitlementId field value if set, zero value otherwise.
func (o *DeeplinkApp) GetAppleAppEntitlementId() string {
	if o == nil || o.AppleAppEntitlementId == nil {
		var ret string
		return ret
	}
	return *o.AppleAppEntitlementId
}

// GetAppleAppEntitlementIdOk returns a tuple with the AppleAppEntitlementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeeplinkApp) GetAppleAppEntitlementIdOk() (*string, bool) {
	if o == nil || o.AppleAppEntitlementId == nil {
		return nil, false
	}
	return o.AppleAppEntitlementId, true
}

// HasAppleAppEntitlementId returns a boolean if a field has been set.
func (o *DeeplinkApp) HasAppleAppEntitlementId() bool {
	if o != nil && o.AppleAppEntitlementId != nil {
		return true
	}

	return false
}

// SetAppleAppEntitlementId gets a reference to the given string and assigns it to the AppleAppEntitlementId field.
func (o *DeeplinkApp) SetAppleAppEntitlementId(v string) {
	o.AppleAppEntitlementId = &v
}

// GetAndroidSha256 returns the AndroidSha256 field value if set, zero value otherwise.
func (o *DeeplinkApp) GetAndroidSha256() []string {
	if o == nil || o.AndroidSha256 == nil {
		var ret []string
		return ret
	}
	return *o.AndroidSha256
}

// GetAndroidSha256Ok returns a tuple with the AndroidSha256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeeplinkApp) GetAndroidSha256Ok() (*[]string, bool) {
	if o == nil || o.AndroidSha256 == nil {
		return nil, false
	}
	return o.AndroidSha256, true
}

// HasAndroidSha256 returns a boolean if a field has been set.
func (o *DeeplinkApp) HasAndroidSha256() bool {
	if o != nil && o.AndroidSha256 != nil {
		return true
	}

	return false
}

// SetAndroidSha256 gets a reference to the given []string and assigns it to the AndroidSha256 field.
func (o *DeeplinkApp) SetAndroidSha256(v []string) {
	o.AndroidSha256 = &v
}

// GetOs returns the Os field value
func (o *DeeplinkApp) GetOs() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Os
}

// GetOsOk returns a tuple with the Os field value
// and a boolean to check if the value has been set.
func (o *DeeplinkApp) GetOsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Os, true
}

// SetOs sets field value
func (o *DeeplinkApp) SetOs(v string) {
	o.Os = v
}

func (o DeeplinkApp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["third_party_app_id"] = o.ThirdPartyAppId
	}
	if o.CreatedTs != nil {
		toSerialize["created_ts"] = o.CreatedTs
	}
	if o.Scheme != nil {
		toSerialize["scheme"] = o.Scheme
	}
	if true {
		toSerialize["install_url"] = o.InstallUrl
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.LegacyAppId != nil {
		toSerialize["legacy_app_id"] = o.LegacyAppId
	}
	if o.IconUrl != nil {
		toSerialize["icon_url"] = o.IconUrl
	}
	if o.OrganizationGuid != nil {
		toSerialize["organization_guid"] = o.OrganizationGuid
	}
	if o.ModifiedTs != nil {
		toSerialize["modified_ts"] = o.ModifiedTs
	}
	if o.Guid != nil {
		toSerialize["guid"] = o.Guid
	}
	if o.AppleAppEntitlementId != nil {
		toSerialize["apple_app_entitlement_id"] = o.AppleAppEntitlementId
	}
	if o.AndroidSha256 != nil {
		toSerialize["android_sha256"] = o.AndroidSha256
	}
	if true {
		toSerialize["os"] = o.Os
	}
	return json.Marshal(toSerialize)
}

type NullableDeeplinkApp struct {
	value *DeeplinkApp
	isSet bool
}

func (v NullableDeeplinkApp) Get() *DeeplinkApp {
	return v.value
}

func (v *NullableDeeplinkApp) Set(val *DeeplinkApp) {
	v.value = val
	v.isSet = true
}

func (v NullableDeeplinkApp) IsSet() bool {
	return v.isSet
}

func (v *NullableDeeplinkApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeeplinkApp(val *DeeplinkApp) *NullableDeeplinkApp {
	return &NullableDeeplinkApp{value: val, isSet: true}
}

func (v NullableDeeplinkApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeeplinkApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
