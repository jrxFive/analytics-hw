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

// BitlinkBody struct for BitlinkBody
type BitlinkBody struct {
	References     *map[string]string `json:"references,omitempty"`
	Archived       *bool              `json:"archived,omitempty"`
	Tags           *[]string          `json:"tags,omitempty"`
	CreatedAt      *string            `json:"created_at,omitempty"`
	Title          *string            `json:"title,omitempty"`
	Deeplinks      *[]DeeplinkRule    `json:"deeplinks,omitempty"`
	CreatedBy      *string            `json:"created_by,omitempty"`
	LongUrl        *string            `json:"long_url,omitempty"`
	ClientId       *string            `json:"client_id,omitempty"`
	CustomBitlinks *[]string          `json:"custom_bitlinks,omitempty"`
	Link           *string            `json:"link,omitempty"`
	LaunchpadIds   *[]string          `json:"launchpad_ids,omitempty"`
	Id             *string            `json:"id,omitempty"`
}

// NewBitlinkBody instantiates a new BitlinkBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBitlinkBody() *BitlinkBody {
	this := BitlinkBody{}
	return &this
}

// NewBitlinkBodyWithDefaults instantiates a new BitlinkBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBitlinkBodyWithDefaults() *BitlinkBody {
	this := BitlinkBody{}
	return &this
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *BitlinkBody) GetReferences() map[string]string {
	if o == nil || o.References == nil {
		var ret map[string]string
		return ret
	}
	return *o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BitlinkBody) GetReferencesOk() (*map[string]string, bool) {
	if o == nil || o.References == nil {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *BitlinkBody) HasReferences() bool {
	if o != nil && o.References != nil {
		return true
	}

	return false
}

// SetReferences gets a reference to the given map[string]string and assigns it to the References field.
func (o *BitlinkBody) SetReferences(v map[string]string) {
	o.References = &v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *BitlinkBody) GetArchived() bool {
	if o == nil || o.Archived == nil {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BitlinkBody) GetArchivedOk() (*bool, bool) {
	if o == nil || o.Archived == nil {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *BitlinkBody) HasArchived() bool {
	if o != nil && o.Archived != nil {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *BitlinkBody) SetArchived(v bool) {
	o.Archived = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *BitlinkBody) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BitlinkBody) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BitlinkBody) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *BitlinkBody) SetTags(v []string) {
	o.Tags = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BitlinkBody) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BitlinkBody) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BitlinkBody) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *BitlinkBody) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *BitlinkBody) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BitlinkBody) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *BitlinkBody) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *BitlinkBody) SetTitle(v string) {
	o.Title = &v
}

// GetDeeplinks returns the Deeplinks field value if set, zero value otherwise.
func (o *BitlinkBody) GetDeeplinks() []DeeplinkRule {
	if o == nil || o.Deeplinks == nil {
		var ret []DeeplinkRule
		return ret
	}
	return *o.Deeplinks
}

// GetDeeplinksOk returns a tuple with the Deeplinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BitlinkBody) GetDeeplinksOk() (*[]DeeplinkRule, bool) {
	if o == nil || o.Deeplinks == nil {
		return nil, false
	}
	return o.Deeplinks, true
}

// HasDeeplinks returns a boolean if a field has been set.
func (o *BitlinkBody) HasDeeplinks() bool {
	if o != nil && o.Deeplinks != nil {
		return true
	}

	return false
}

// SetDeeplinks gets a reference to the given []DeeplinkRule and assigns it to the Deeplinks field.
func (o *BitlinkBody) SetDeeplinks(v []DeeplinkRule) {
	o.Deeplinks = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *BitlinkBody) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BitlinkBody) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *BitlinkBody) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *BitlinkBody) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetLongUrl returns the LongUrl field value if set, zero value otherwise.
func (o *BitlinkBody) GetLongUrl() string {
	if o == nil || o.LongUrl == nil {
		var ret string
		return ret
	}
	return *o.LongUrl
}

// GetLongUrlOk returns a tuple with the LongUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BitlinkBody) GetLongUrlOk() (*string, bool) {
	if o == nil || o.LongUrl == nil {
		return nil, false
	}
	return o.LongUrl, true
}

// HasLongUrl returns a boolean if a field has been set.
func (o *BitlinkBody) HasLongUrl() bool {
	if o != nil && o.LongUrl != nil {
		return true
	}

	return false
}

// SetLongUrl gets a reference to the given string and assigns it to the LongUrl field.
func (o *BitlinkBody) SetLongUrl(v string) {
	o.LongUrl = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *BitlinkBody) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BitlinkBody) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *BitlinkBody) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *BitlinkBody) SetClientId(v string) {
	o.ClientId = &v
}

// GetCustomBitlinks returns the CustomBitlinks field value if set, zero value otherwise.
func (o *BitlinkBody) GetCustomBitlinks() []string {
	if o == nil || o.CustomBitlinks == nil {
		var ret []string
		return ret
	}
	return *o.CustomBitlinks
}

// GetCustomBitlinksOk returns a tuple with the CustomBitlinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BitlinkBody) GetCustomBitlinksOk() (*[]string, bool) {
	if o == nil || o.CustomBitlinks == nil {
		return nil, false
	}
	return o.CustomBitlinks, true
}

// HasCustomBitlinks returns a boolean if a field has been set.
func (o *BitlinkBody) HasCustomBitlinks() bool {
	if o != nil && o.CustomBitlinks != nil {
		return true
	}

	return false
}

// SetCustomBitlinks gets a reference to the given []string and assigns it to the CustomBitlinks field.
func (o *BitlinkBody) SetCustomBitlinks(v []string) {
	o.CustomBitlinks = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *BitlinkBody) GetLink() string {
	if o == nil || o.Link == nil {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BitlinkBody) GetLinkOk() (*string, bool) {
	if o == nil || o.Link == nil {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *BitlinkBody) HasLink() bool {
	if o != nil && o.Link != nil {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *BitlinkBody) SetLink(v string) {
	o.Link = &v
}

// GetLaunchpadIds returns the LaunchpadIds field value if set, zero value otherwise.
func (o *BitlinkBody) GetLaunchpadIds() []string {
	if o == nil || o.LaunchpadIds == nil {
		var ret []string
		return ret
	}
	return *o.LaunchpadIds
}

// GetLaunchpadIdsOk returns a tuple with the LaunchpadIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BitlinkBody) GetLaunchpadIdsOk() (*[]string, bool) {
	if o == nil || o.LaunchpadIds == nil {
		return nil, false
	}
	return o.LaunchpadIds, true
}

// HasLaunchpadIds returns a boolean if a field has been set.
func (o *BitlinkBody) HasLaunchpadIds() bool {
	if o != nil && o.LaunchpadIds != nil {
		return true
	}

	return false
}

// SetLaunchpadIds gets a reference to the given []string and assigns it to the LaunchpadIds field.
func (o *BitlinkBody) SetLaunchpadIds(v []string) {
	o.LaunchpadIds = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BitlinkBody) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BitlinkBody) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BitlinkBody) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BitlinkBody) SetId(v string) {
	o.Id = &v
}

func (o BitlinkBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.References != nil {
		toSerialize["references"] = o.References
	}
	if o.Archived != nil {
		toSerialize["archived"] = o.Archived
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Deeplinks != nil {
		toSerialize["deeplinks"] = o.Deeplinks
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.LongUrl != nil {
		toSerialize["long_url"] = o.LongUrl
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.CustomBitlinks != nil {
		toSerialize["custom_bitlinks"] = o.CustomBitlinks
	}
	if o.Link != nil {
		toSerialize["link"] = o.Link
	}
	if o.LaunchpadIds != nil {
		toSerialize["launchpad_ids"] = o.LaunchpadIds
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableBitlinkBody struct {
	value *BitlinkBody
	isSet bool
}

func (v NullableBitlinkBody) Get() *BitlinkBody {
	return v.value
}

func (v *NullableBitlinkBody) Set(val *BitlinkBody) {
	v.value = val
	v.isSet = true
}

func (v NullableBitlinkBody) IsSet() bool {
	return v.isSet
}

func (v *NullableBitlinkBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBitlinkBody(val *BitlinkBody) *NullableBitlinkBody {
	return &NullableBitlinkBody{value: val, isSet: true}
}

func (v NullableBitlinkBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBitlinkBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
