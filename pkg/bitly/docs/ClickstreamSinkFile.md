# ClickstreamSinkFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateFrom** | Pointer to **time.Time** | start of period | [optional] 
**DateTo** | Pointer to **time.Time** | end of period | [optional] 
**FileId** | Pointer to **string** | Unique id of the file | [optional] 
**DownloadUrl** | Pointer to **string** | use this url with follow redirects to download the file | [optional] 
**Created** | Pointer to **time.Time** | time of creation | [optional] 

## Methods

### NewClickstreamSinkFile

`func NewClickstreamSinkFile() *ClickstreamSinkFile`

NewClickstreamSinkFile instantiates a new ClickstreamSinkFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClickstreamSinkFileWithDefaults

`func NewClickstreamSinkFileWithDefaults() *ClickstreamSinkFile`

NewClickstreamSinkFileWithDefaults instantiates a new ClickstreamSinkFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateFrom

`func (o *ClickstreamSinkFile) GetDateFrom() time.Time`

GetDateFrom returns the DateFrom field if non-nil, zero value otherwise.

### GetDateFromOk

`func (o *ClickstreamSinkFile) GetDateFromOk() (*time.Time, bool)`

GetDateFromOk returns a tuple with the DateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFrom

`func (o *ClickstreamSinkFile) SetDateFrom(v time.Time)`

SetDateFrom sets DateFrom field to given value.

### HasDateFrom

`func (o *ClickstreamSinkFile) HasDateFrom() bool`

HasDateFrom returns a boolean if a field has been set.

### GetDateTo

`func (o *ClickstreamSinkFile) GetDateTo() time.Time`

GetDateTo returns the DateTo field if non-nil, zero value otherwise.

### GetDateToOk

`func (o *ClickstreamSinkFile) GetDateToOk() (*time.Time, bool)`

GetDateToOk returns a tuple with the DateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTo

`func (o *ClickstreamSinkFile) SetDateTo(v time.Time)`

SetDateTo sets DateTo field to given value.

### HasDateTo

`func (o *ClickstreamSinkFile) HasDateTo() bool`

HasDateTo returns a boolean if a field has been set.

### GetFileId

`func (o *ClickstreamSinkFile) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *ClickstreamSinkFile) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *ClickstreamSinkFile) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *ClickstreamSinkFile) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *ClickstreamSinkFile) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *ClickstreamSinkFile) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *ClickstreamSinkFile) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *ClickstreamSinkFile) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetCreated

`func (o *ClickstreamSinkFile) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ClickstreamSinkFile) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ClickstreamSinkFile) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ClickstreamSinkFile) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


