# CreateOauthAppReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RedirectUris** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrganizationGuid** | Pointer to **string** |  | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**IpAllowlist** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateOauthAppReq

`func NewCreateOauthAppReq() *CreateOauthAppReq`

NewCreateOauthAppReq instantiates a new CreateOauthAppReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOauthAppReqWithDefaults

`func NewCreateOauthAppReqWithDefaults() *CreateOauthAppReq`

NewCreateOauthAppReqWithDefaults instantiates a new CreateOauthAppReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedirectUris

`func (o *CreateOauthAppReq) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *CreateOauthAppReq) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *CreateOauthAppReq) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *CreateOauthAppReq) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetName

`func (o *CreateOauthAppReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOauthAppReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOauthAppReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateOauthAppReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationGuid

`func (o *CreateOauthAppReq) GetOrganizationGuid() string`

GetOrganizationGuid returns the OrganizationGuid field if non-nil, zero value otherwise.

### GetOrganizationGuidOk

`func (o *CreateOauthAppReq) GetOrganizationGuidOk() (*string, bool)`

GetOrganizationGuidOk returns a tuple with the OrganizationGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationGuid

`func (o *CreateOauthAppReq) SetOrganizationGuid(v string)`

SetOrganizationGuid sets OrganizationGuid field to given value.

### HasOrganizationGuid

`func (o *CreateOauthAppReq) HasOrganizationGuid() bool`

HasOrganizationGuid returns a boolean if a field has been set.

### GetLink

`func (o *CreateOauthAppReq) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *CreateOauthAppReq) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *CreateOauthAppReq) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *CreateOauthAppReq) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetIpAllowlist

`func (o *CreateOauthAppReq) GetIpAllowlist() []string`

GetIpAllowlist returns the IpAllowlist field if non-nil, zero value otherwise.

### GetIpAllowlistOk

`func (o *CreateOauthAppReq) GetIpAllowlistOk() (*[]string, bool)`

GetIpAllowlistOk returns a tuple with the IpAllowlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAllowlist

`func (o *CreateOauthAppReq) SetIpAllowlist(v []string)`

SetIpAllowlist sets IpAllowlist field to given value.

### HasIpAllowlist

`func (o *CreateOauthAppReq) HasIpAllowlist() bool`

HasIpAllowlist returns a boolean if a field has been set.

### GetDescription

`func (o *CreateOauthAppReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOauthAppReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOauthAppReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateOauthAppReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


