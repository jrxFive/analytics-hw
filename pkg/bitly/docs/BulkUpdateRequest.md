# BulkUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**RemoveTags** | Pointer to **[]string** |  | [optional] 
**Archive** | Pointer to **bool** |  | [optional] 
**Links** | Pointer to **[]string** | this is limited to 100 bitlink ids | [optional] 
**AddTags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBulkUpdateRequest

`func NewBulkUpdateRequest() *BulkUpdateRequest`

NewBulkUpdateRequest instantiates a new BulkUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUpdateRequestWithDefaults

`func NewBulkUpdateRequestWithDefaults() *BulkUpdateRequest`

NewBulkUpdateRequestWithDefaults instantiates a new BulkUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *BulkUpdateRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BulkUpdateRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BulkUpdateRequest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *BulkUpdateRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetRemoveTags

`func (o *BulkUpdateRequest) GetRemoveTags() []string`

GetRemoveTags returns the RemoveTags field if non-nil, zero value otherwise.

### GetRemoveTagsOk

`func (o *BulkUpdateRequest) GetRemoveTagsOk() (*[]string, bool)`

GetRemoveTagsOk returns a tuple with the RemoveTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveTags

`func (o *BulkUpdateRequest) SetRemoveTags(v []string)`

SetRemoveTags sets RemoveTags field to given value.

### HasRemoveTags

`func (o *BulkUpdateRequest) HasRemoveTags() bool`

HasRemoveTags returns a boolean if a field has been set.

### GetArchive

`func (o *BulkUpdateRequest) GetArchive() bool`

GetArchive returns the Archive field if non-nil, zero value otherwise.

### GetArchiveOk

`func (o *BulkUpdateRequest) GetArchiveOk() (*bool, bool)`

GetArchiveOk returns a tuple with the Archive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchive

`func (o *BulkUpdateRequest) SetArchive(v bool)`

SetArchive sets Archive field to given value.

### HasArchive

`func (o *BulkUpdateRequest) HasArchive() bool`

HasArchive returns a boolean if a field has been set.

### GetLinks

`func (o *BulkUpdateRequest) GetLinks() []string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BulkUpdateRequest) GetLinksOk() (*[]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BulkUpdateRequest) SetLinks(v []string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BulkUpdateRequest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAddTags

`func (o *BulkUpdateRequest) GetAddTags() []string`

GetAddTags returns the AddTags field if non-nil, zero value otherwise.

### GetAddTagsOk

`func (o *BulkUpdateRequest) GetAddTagsOk() (*[]string, bool)`

GetAddTagsOk returns a tuple with the AddTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTags

`func (o *BulkUpdateRequest) SetAddTags(v []string)`

SetAddTags sets AddTags field to given value.

### HasAddTags

`func (o *BulkUpdateRequest) HasAddTags() bool`

HasAddTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


