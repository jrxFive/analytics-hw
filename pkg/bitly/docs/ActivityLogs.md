# ActivityLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPage** | Pointer to **string** |  | [optional] 
**Activity** | Pointer to [**[]ActivityLog**](ActivityLog.md) |  | [optional] 

## Methods

### NewActivityLogs

`func NewActivityLogs() *ActivityLogs`

NewActivityLogs instantiates a new ActivityLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityLogsWithDefaults

`func NewActivityLogsWithDefaults() *ActivityLogs`

NewActivityLogsWithDefaults instantiates a new ActivityLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPage

`func (o *ActivityLogs) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *ActivityLogs) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *ActivityLogs) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *ActivityLogs) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### GetActivity

`func (o *ActivityLogs) GetActivity() []ActivityLog`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *ActivityLogs) GetActivityOk() (*[]ActivityLog, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *ActivityLogs) SetActivity(v []ActivityLog)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *ActivityLogs) HasActivity() bool`

HasActivity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


