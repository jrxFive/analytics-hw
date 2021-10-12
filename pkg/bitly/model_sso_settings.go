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

// SSOSettings struct for SSOSettings
type SSOSettings struct {
	SloEndpoint      *string   `json:"slo_endpoint,omitempty"`
	DefaultGroupGuid *string   `json:"default_group_guid,omitempty"`
	IdpUrl           *string   `json:"idp_url,omitempty"`
	Certificate      *string   `json:"certificate,omitempty"`
	OrganizationGuid *string   `json:"organization_guid,omitempty"`
	SamlEndpoint     *string   `json:"saml_endpoint,omitempty"`
	IdentityProvider *string   `json:"identity_provider,omitempty"`
	Domains          *[]string `json:"domains,omitempty"`
	IssuerUrl        *string   `json:"issuer_url,omitempty"`
	UrlSlug          *string   `json:"url_slug,omitempty"`
}

// NewSSOSettings instantiates a new SSOSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSSOSettings() *SSOSettings {
	this := SSOSettings{}
	return &this
}

// NewSSOSettingsWithDefaults instantiates a new SSOSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSOSettingsWithDefaults() *SSOSettings {
	this := SSOSettings{}
	return &this
}

// GetSloEndpoint returns the SloEndpoint field value if set, zero value otherwise.
func (o *SSOSettings) GetSloEndpoint() string {
	if o == nil || o.SloEndpoint == nil {
		var ret string
		return ret
	}
	return *o.SloEndpoint
}

// GetSloEndpointOk returns a tuple with the SloEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOSettings) GetSloEndpointOk() (*string, bool) {
	if o == nil || o.SloEndpoint == nil {
		return nil, false
	}
	return o.SloEndpoint, true
}

// HasSloEndpoint returns a boolean if a field has been set.
func (o *SSOSettings) HasSloEndpoint() bool {
	if o != nil && o.SloEndpoint != nil {
		return true
	}

	return false
}

// SetSloEndpoint gets a reference to the given string and assigns it to the SloEndpoint field.
func (o *SSOSettings) SetSloEndpoint(v string) {
	o.SloEndpoint = &v
}

// GetDefaultGroupGuid returns the DefaultGroupGuid field value if set, zero value otherwise.
func (o *SSOSettings) GetDefaultGroupGuid() string {
	if o == nil || o.DefaultGroupGuid == nil {
		var ret string
		return ret
	}
	return *o.DefaultGroupGuid
}

// GetDefaultGroupGuidOk returns a tuple with the DefaultGroupGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOSettings) GetDefaultGroupGuidOk() (*string, bool) {
	if o == nil || o.DefaultGroupGuid == nil {
		return nil, false
	}
	return o.DefaultGroupGuid, true
}

// HasDefaultGroupGuid returns a boolean if a field has been set.
func (o *SSOSettings) HasDefaultGroupGuid() bool {
	if o != nil && o.DefaultGroupGuid != nil {
		return true
	}

	return false
}

// SetDefaultGroupGuid gets a reference to the given string and assigns it to the DefaultGroupGuid field.
func (o *SSOSettings) SetDefaultGroupGuid(v string) {
	o.DefaultGroupGuid = &v
}

// GetIdpUrl returns the IdpUrl field value if set, zero value otherwise.
func (o *SSOSettings) GetIdpUrl() string {
	if o == nil || o.IdpUrl == nil {
		var ret string
		return ret
	}
	return *o.IdpUrl
}

// GetIdpUrlOk returns a tuple with the IdpUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOSettings) GetIdpUrlOk() (*string, bool) {
	if o == nil || o.IdpUrl == nil {
		return nil, false
	}
	return o.IdpUrl, true
}

// HasIdpUrl returns a boolean if a field has been set.
func (o *SSOSettings) HasIdpUrl() bool {
	if o != nil && o.IdpUrl != nil {
		return true
	}

	return false
}

// SetIdpUrl gets a reference to the given string and assigns it to the IdpUrl field.
func (o *SSOSettings) SetIdpUrl(v string) {
	o.IdpUrl = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *SSOSettings) GetCertificate() string {
	if o == nil || o.Certificate == nil {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOSettings) GetCertificateOk() (*string, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *SSOSettings) HasCertificate() bool {
	if o != nil && o.Certificate != nil {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *SSOSettings) SetCertificate(v string) {
	o.Certificate = &v
}

// GetOrganizationGuid returns the OrganizationGuid field value if set, zero value otherwise.
func (o *SSOSettings) GetOrganizationGuid() string {
	if o == nil || o.OrganizationGuid == nil {
		var ret string
		return ret
	}
	return *o.OrganizationGuid
}

// GetOrganizationGuidOk returns a tuple with the OrganizationGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOSettings) GetOrganizationGuidOk() (*string, bool) {
	if o == nil || o.OrganizationGuid == nil {
		return nil, false
	}
	return o.OrganizationGuid, true
}

// HasOrganizationGuid returns a boolean if a field has been set.
func (o *SSOSettings) HasOrganizationGuid() bool {
	if o != nil && o.OrganizationGuid != nil {
		return true
	}

	return false
}

// SetOrganizationGuid gets a reference to the given string and assigns it to the OrganizationGuid field.
func (o *SSOSettings) SetOrganizationGuid(v string) {
	o.OrganizationGuid = &v
}

// GetSamlEndpoint returns the SamlEndpoint field value if set, zero value otherwise.
func (o *SSOSettings) GetSamlEndpoint() string {
	if o == nil || o.SamlEndpoint == nil {
		var ret string
		return ret
	}
	return *o.SamlEndpoint
}

// GetSamlEndpointOk returns a tuple with the SamlEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOSettings) GetSamlEndpointOk() (*string, bool) {
	if o == nil || o.SamlEndpoint == nil {
		return nil, false
	}
	return o.SamlEndpoint, true
}

// HasSamlEndpoint returns a boolean if a field has been set.
func (o *SSOSettings) HasSamlEndpoint() bool {
	if o != nil && o.SamlEndpoint != nil {
		return true
	}

	return false
}

// SetSamlEndpoint gets a reference to the given string and assigns it to the SamlEndpoint field.
func (o *SSOSettings) SetSamlEndpoint(v string) {
	o.SamlEndpoint = &v
}

// GetIdentityProvider returns the IdentityProvider field value if set, zero value otherwise.
func (o *SSOSettings) GetIdentityProvider() string {
	if o == nil || o.IdentityProvider == nil {
		var ret string
		return ret
	}
	return *o.IdentityProvider
}

// GetIdentityProviderOk returns a tuple with the IdentityProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOSettings) GetIdentityProviderOk() (*string, bool) {
	if o == nil || o.IdentityProvider == nil {
		return nil, false
	}
	return o.IdentityProvider, true
}

// HasIdentityProvider returns a boolean if a field has been set.
func (o *SSOSettings) HasIdentityProvider() bool {
	if o != nil && o.IdentityProvider != nil {
		return true
	}

	return false
}

// SetIdentityProvider gets a reference to the given string and assigns it to the IdentityProvider field.
func (o *SSOSettings) SetIdentityProvider(v string) {
	o.IdentityProvider = &v
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *SSOSettings) GetDomains() []string {
	if o == nil || o.Domains == nil {
		var ret []string
		return ret
	}
	return *o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOSettings) GetDomainsOk() (*[]string, bool) {
	if o == nil || o.Domains == nil {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *SSOSettings) HasDomains() bool {
	if o != nil && o.Domains != nil {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []string and assigns it to the Domains field.
func (o *SSOSettings) SetDomains(v []string) {
	o.Domains = &v
}

// GetIssuerUrl returns the IssuerUrl field value if set, zero value otherwise.
func (o *SSOSettings) GetIssuerUrl() string {
	if o == nil || o.IssuerUrl == nil {
		var ret string
		return ret
	}
	return *o.IssuerUrl
}

// GetIssuerUrlOk returns a tuple with the IssuerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOSettings) GetIssuerUrlOk() (*string, bool) {
	if o == nil || o.IssuerUrl == nil {
		return nil, false
	}
	return o.IssuerUrl, true
}

// HasIssuerUrl returns a boolean if a field has been set.
func (o *SSOSettings) HasIssuerUrl() bool {
	if o != nil && o.IssuerUrl != nil {
		return true
	}

	return false
}

// SetIssuerUrl gets a reference to the given string and assigns it to the IssuerUrl field.
func (o *SSOSettings) SetIssuerUrl(v string) {
	o.IssuerUrl = &v
}

// GetUrlSlug returns the UrlSlug field value if set, zero value otherwise.
func (o *SSOSettings) GetUrlSlug() string {
	if o == nil || o.UrlSlug == nil {
		var ret string
		return ret
	}
	return *o.UrlSlug
}

// GetUrlSlugOk returns a tuple with the UrlSlug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSOSettings) GetUrlSlugOk() (*string, bool) {
	if o == nil || o.UrlSlug == nil {
		return nil, false
	}
	return o.UrlSlug, true
}

// HasUrlSlug returns a boolean if a field has been set.
func (o *SSOSettings) HasUrlSlug() bool {
	if o != nil && o.UrlSlug != nil {
		return true
	}

	return false
}

// SetUrlSlug gets a reference to the given string and assigns it to the UrlSlug field.
func (o *SSOSettings) SetUrlSlug(v string) {
	o.UrlSlug = &v
}

func (o SSOSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SloEndpoint != nil {
		toSerialize["slo_endpoint"] = o.SloEndpoint
	}
	if o.DefaultGroupGuid != nil {
		toSerialize["default_group_guid"] = o.DefaultGroupGuid
	}
	if o.IdpUrl != nil {
		toSerialize["idp_url"] = o.IdpUrl
	}
	if o.Certificate != nil {
		toSerialize["certificate"] = o.Certificate
	}
	if o.OrganizationGuid != nil {
		toSerialize["organization_guid"] = o.OrganizationGuid
	}
	if o.SamlEndpoint != nil {
		toSerialize["saml_endpoint"] = o.SamlEndpoint
	}
	if o.IdentityProvider != nil {
		toSerialize["identity_provider"] = o.IdentityProvider
	}
	if o.Domains != nil {
		toSerialize["domains"] = o.Domains
	}
	if o.IssuerUrl != nil {
		toSerialize["issuer_url"] = o.IssuerUrl
	}
	if o.UrlSlug != nil {
		toSerialize["url_slug"] = o.UrlSlug
	}
	return json.Marshal(toSerialize)
}

type NullableSSOSettings struct {
	value *SSOSettings
	isSet bool
}

func (v NullableSSOSettings) Get() *SSOSettings {
	return v.value
}

func (v *NullableSSOSettings) Set(val *SSOSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSSOSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSSOSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSSOSettings(val *SSOSettings) *NullableSSOSettings {
	return &NullableSSOSettings{value: val, isSet: true}
}

func (v NullableSSOSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSSOSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
