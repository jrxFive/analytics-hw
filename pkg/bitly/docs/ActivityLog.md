# ActivityLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionDescription** | Pointer to **string** |  | [optional] 
**Ts** | Pointer to **string** | ISO timestamp | [optional] 
**User** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**OrgGuid** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewActivityLog

`func NewActivityLog() *ActivityLog`

NewActivityLog instantiates a new ActivityLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityLogWithDefaults

`func NewActivityLogWithDefaults() *ActivityLog`

NewActivityLogWithDefaults instantiates a new ActivityLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionDescription

`func (o *ActivityLog) GetActionDescription() string`

GetActionDescription returns the ActionDescription field if non-nil, zero value otherwise.

### GetActionDescriptionOk

`func (o *ActivityLog) GetActionDescriptionOk() (*string, bool)`

GetActionDescriptionOk returns a tuple with the ActionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDescription

`func (o *ActivityLog) SetActionDescription(v string)`

SetActionDescription sets ActionDescription field to given value.

### HasActionDescription

`func (o *ActivityLog) HasActionDescription() bool`

HasActionDescription returns a boolean if a field has been set.

### GetTs

`func (o *ActivityLog) GetTs() string`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *ActivityLog) GetTsOk() (*string, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *ActivityLog) SetTs(v string)`

SetTs sets Ts field to given value.

### HasTs

`func (o *ActivityLog) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetUser

`func (o *ActivityLog) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ActivityLog) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ActivityLog) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ActivityLog) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetAction

`func (o *ActivityLog) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ActivityLog) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ActivityLog) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ActivityLog) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetOrgGuid

`func (o *ActivityLog) GetOrgGuid() string`

GetOrgGuid returns the OrgGuid field if non-nil, zero value otherwise.

### GetOrgGuidOk

`func (o *ActivityLog) GetOrgGuidOk() (*string, bool)`

GetOrgGuidOk returns a tuple with the OrgGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgGuid

`func (o *ActivityLog) SetOrgGuid(v string)`

SetOrgGuid sets OrgGuid field to given value.

### HasOrgGuid

`func (o *ActivityLog) HasOrgGuid() bool`

HasOrgGuid returns a boolean if a field has been set.

### GetIpAddress

`func (o *ActivityLog) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ActivityLog) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ActivityLog) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ActivityLog) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetId

`func (o *ActivityLog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivityLog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivityLog) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ActivityLog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *ActivityLog) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ActivityLog) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ActivityLog) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ActivityLog) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


