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

// UserOnboardingSurvey struct for UserOnboardingSurvey
type UserOnboardingSurvey struct {
	FirstName        *string   `json:"first_name,omitempty"`
	LastName         *string   `json:"last_name,omitempty"`
	Created          *string   `json:"created,omitempty"`
	Modified         *string   `json:"modified,omitempty"`
	CompanySize      *string   `json:"company_size,omitempty"`
	OrganizationName *string   `json:"organization_name,omitempty"`
	UseCases         *[]string `json:"use_cases,omitempty"`
	Version          *float32  `json:"version,omitempty"`
	UseCasesOther    *[]string `json:"use_cases_other,omitempty"`
	Department       *string   `json:"department,omitempty"`
	Login            *string   `json:"login,omitempty"`
	JobTitle         *string   `json:"job_title,omitempty"`
	DepartmentOther  *string   `json:"department_other,omitempty"`
	TeamType         *string   `json:"team_type,omitempty"`
}

// NewUserOnboardingSurvey instantiates a new UserOnboardingSurvey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserOnboardingSurvey() *UserOnboardingSurvey {
	this := UserOnboardingSurvey{}
	return &this
}

// NewUserOnboardingSurveyWithDefaults instantiates a new UserOnboardingSurvey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserOnboardingSurveyWithDefaults() *UserOnboardingSurvey {
	this := UserOnboardingSurvey{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UserOnboardingSurvey) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserOnboardingSurvey) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UserOnboardingSurvey) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UserOnboardingSurvey) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UserOnboardingSurvey) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserOnboardingSurvey) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UserOnboardingSurvey) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *UserOnboardingSurvey) SetLastName(v string) {
	o.LastName = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *UserOnboardingSurvey) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserOnboardingSurvey) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *UserOnboardingSurvey) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *UserOnboardingSurvey) SetCreated(v string) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *UserOnboardingSurvey) GetModified() string {
	if o == nil || o.Modified == nil {
		var ret string
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserOnboardingSurvey) GetModifiedOk() (*string, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *UserOnboardingSurvey) HasModified() bool {
	if o != nil && o.Modified != nil {
		return true
	}

	return false
}

// SetModified gets a reference to the given string and assigns it to the Modified field.
func (o *UserOnboardingSurvey) SetModified(v string) {
	o.Modified = &v
}

// GetCompanySize returns the CompanySize field value if set, zero value otherwise.
func (o *UserOnboardingSurvey) GetCompanySize() string {
	if o == nil || o.CompanySize == nil {
		var ret string
		return ret
	}
	return *o.CompanySize
}

// GetCompanySizeOk returns a tuple with the CompanySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserOnboardingSurvey) GetCompanySizeOk() (*string, bool) {
	if o == nil || o.CompanySize == nil {
		return nil, false
	}
	return o.CompanySize, true
}

// HasCompanySize returns a boolean if a field has been set.
func (o *UserOnboardingSurvey) HasCompanySize() bool {
	if o != nil && o.CompanySize != nil {
		return true
	}

	return false
}

// SetCompanySize gets a reference to the given string and assigns it to the CompanySize field.
func (o *UserOnboardingSurvey) SetCompanySize(v string) {
	o.CompanySize = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise.
func (o *UserOnboardingSurvey) GetOrganizationName() string {
	if o == nil || o.OrganizationName == nil {
		var ret string
		return ret
	}
	return *o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserOnboardingSurvey) GetOrganizationNameOk() (*string, bool) {
	if o == nil || o.OrganizationName == nil {
		return nil, false
	}
	return o.OrganizationName, true
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *UserOnboardingSurvey) HasOrganizationName() bool {
	if o != nil && o.OrganizationName != nil {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given string and assigns it to the OrganizationName field.
func (o *UserOnboardingSurvey) SetOrganizationName(v string) {
	o.OrganizationName = &v
}

// GetUseCases returns the UseCases field value if set, zero value otherwise.
func (o *UserOnboardingSurvey) GetUseCases() []string {
	if o == nil || o.UseCases == nil {
		var ret []string
		return ret
	}
	return *o.UseCases
}

// GetUseCasesOk returns a tuple with the UseCases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserOnboardingSurvey) GetUseCasesOk() (*[]string, bool) {
	if o == nil || o.UseCases == nil {
		return nil, false
	}
	return o.UseCases, true
}

// HasUseCases returns a boolean if a field has been set.
func (o *UserOnboardingSurvey) HasUseCases() bool {
	if o != nil && o.UseCases != nil {
		return true
	}

	return false
}

// SetUseCases gets a reference to the given []string and assigns it to the UseCases field.
func (o *UserOnboardingSurvey) SetUseCases(v []string) {
	o.UseCases = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *UserOnboardingSurvey) GetVersion() float32 {
	if o == nil || o.Version == nil {
		var ret float32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserOnboardingSurvey) GetVersionOk() (*float32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *UserOnboardingSurvey) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given float32 and assigns it to the Version field.
func (o *UserOnboardingSurvey) SetVersion(v float32) {
	o.Version = &v
}

// GetUseCasesOther returns the UseCasesOther field value if set, zero value otherwise.
func (o *UserOnboardingSurvey) GetUseCasesOther() []string {
	if o == nil || o.UseCasesOther == nil {
		var ret []string
		return ret
	}
	return *o.UseCasesOther
}

// GetUseCasesOtherOk returns a tuple with the UseCasesOther field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserOnboardingSurvey) GetUseCasesOtherOk() (*[]string, bool) {
	if o == nil || o.UseCasesOther == nil {
		return nil, false
	}
	return o.UseCasesOther, true
}

// HasUseCasesOther returns a boolean if a field has been set.
func (o *UserOnboardingSurvey) HasUseCasesOther() bool {
	if o != nil && o.UseCasesOther != nil {
		return true
	}

	return false
}

// SetUseCasesOther gets a reference to the given []string and assigns it to the UseCasesOther field.
func (o *UserOnboardingSurvey) SetUseCasesOther(v []string) {
	o.UseCasesOther = &v
}

// GetDepartment returns the Department field value if set, zero value otherwise.
func (o *UserOnboardingSurvey) GetDepartment() string {
	if o == nil || o.Department == nil {
		var ret string
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserOnboardingSurvey) GetDepartmentOk() (*string, bool) {
	if o == nil || o.Department == nil {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *UserOnboardingSurvey) HasDepartment() bool {
	if o != nil && o.Department != nil {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given string and assigns it to the Department field.
func (o *UserOnboardingSurvey) SetDepartment(v string) {
	o.Department = &v
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *UserOnboardingSurvey) GetLogin() string {
	if o == nil || o.Login == nil {
		var ret string
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserOnboardingSurvey) GetLoginOk() (*string, bool) {
	if o == nil || o.Login == nil {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *UserOnboardingSurvey) HasLogin() bool {
	if o != nil && o.Login != nil {
		return true
	}

	return false
}

// SetLogin gets a reference to the given string and assigns it to the Login field.
func (o *UserOnboardingSurvey) SetLogin(v string) {
	o.Login = &v
}

// GetJobTitle returns the JobTitle field value if set, zero value otherwise.
func (o *UserOnboardingSurvey) GetJobTitle() string {
	if o == nil || o.JobTitle == nil {
		var ret string
		return ret
	}
	return *o.JobTitle
}

// GetJobTitleOk returns a tuple with the JobTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserOnboardingSurvey) GetJobTitleOk() (*string, bool) {
	if o == nil || o.JobTitle == nil {
		return nil, false
	}
	return o.JobTitle, true
}

// HasJobTitle returns a boolean if a field has been set.
func (o *UserOnboardingSurvey) HasJobTitle() bool {
	if o != nil && o.JobTitle != nil {
		return true
	}

	return false
}

// SetJobTitle gets a reference to the given string and assigns it to the JobTitle field.
func (o *UserOnboardingSurvey) SetJobTitle(v string) {
	o.JobTitle = &v
}

// GetDepartmentOther returns the DepartmentOther field value if set, zero value otherwise.
func (o *UserOnboardingSurvey) GetDepartmentOther() string {
	if o == nil || o.DepartmentOther == nil {
		var ret string
		return ret
	}
	return *o.DepartmentOther
}

// GetDepartmentOtherOk returns a tuple with the DepartmentOther field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserOnboardingSurvey) GetDepartmentOtherOk() (*string, bool) {
	if o == nil || o.DepartmentOther == nil {
		return nil, false
	}
	return o.DepartmentOther, true
}

// HasDepartmentOther returns a boolean if a field has been set.
func (o *UserOnboardingSurvey) HasDepartmentOther() bool {
	if o != nil && o.DepartmentOther != nil {
		return true
	}

	return false
}

// SetDepartmentOther gets a reference to the given string and assigns it to the DepartmentOther field.
func (o *UserOnboardingSurvey) SetDepartmentOther(v string) {
	o.DepartmentOther = &v
}

// GetTeamType returns the TeamType field value if set, zero value otherwise.
func (o *UserOnboardingSurvey) GetTeamType() string {
	if o == nil || o.TeamType == nil {
		var ret string
		return ret
	}
	return *o.TeamType
}

// GetTeamTypeOk returns a tuple with the TeamType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserOnboardingSurvey) GetTeamTypeOk() (*string, bool) {
	if o == nil || o.TeamType == nil {
		return nil, false
	}
	return o.TeamType, true
}

// HasTeamType returns a boolean if a field has been set.
func (o *UserOnboardingSurvey) HasTeamType() bool {
	if o != nil && o.TeamType != nil {
		return true
	}

	return false
}

// SetTeamType gets a reference to the given string and assigns it to the TeamType field.
func (o *UserOnboardingSurvey) SetTeamType(v string) {
	o.TeamType = &v
}

func (o UserOnboardingSurvey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Modified != nil {
		toSerialize["modified"] = o.Modified
	}
	if o.CompanySize != nil {
		toSerialize["company_size"] = o.CompanySize
	}
	if o.OrganizationName != nil {
		toSerialize["organization_name"] = o.OrganizationName
	}
	if o.UseCases != nil {
		toSerialize["use_cases"] = o.UseCases
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.UseCasesOther != nil {
		toSerialize["use_cases_other"] = o.UseCasesOther
	}
	if o.Department != nil {
		toSerialize["department"] = o.Department
	}
	if o.Login != nil {
		toSerialize["login"] = o.Login
	}
	if o.JobTitle != nil {
		toSerialize["job_title"] = o.JobTitle
	}
	if o.DepartmentOther != nil {
		toSerialize["department_other"] = o.DepartmentOther
	}
	if o.TeamType != nil {
		toSerialize["team_type"] = o.TeamType
	}
	return json.Marshal(toSerialize)
}

type NullableUserOnboardingSurvey struct {
	value *UserOnboardingSurvey
	isSet bool
}

func (v NullableUserOnboardingSurvey) Get() *UserOnboardingSurvey {
	return v.value
}

func (v *NullableUserOnboardingSurvey) Set(val *UserOnboardingSurvey) {
	v.value = val
	v.isSet = true
}

func (v NullableUserOnboardingSurvey) IsSet() bool {
	return v.isSet
}

func (v *NullableUserOnboardingSurvey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserOnboardingSurvey(val *UserOnboardingSurvey) *NullableUserOnboardingSurvey {
	return &NullableUserOnboardingSurvey{value: val, isSet: true}
}

func (v NullableUserOnboardingSurvey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserOnboardingSurvey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
