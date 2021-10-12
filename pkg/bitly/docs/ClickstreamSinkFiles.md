# ClickstreamSinkFiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateFrom** | Pointer to **string** | start of period | [optional] 
**Objects** | Pointer to [**[]ClickstreamSinkFile**](ClickstreamSinkFile.md) |  | [optional] 
**Id** | Pointer to **string** | Unique short id of sink file | [optional] 
**ClickstreamId** | Pointer to **string** | Unique short id of clickstream | [optional] 
**DateTo** | Pointer to **string** | end of period | [optional] 

## Methods

### NewClickstreamSinkFiles

`func NewClickstreamSinkFiles() *ClickstreamSinkFiles`

NewClickstreamSinkFiles instantiates a new ClickstreamSinkFiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClickstreamSinkFilesWithDefaults

`func NewClickstreamSinkFilesWithDefaults() *ClickstreamSinkFiles`

NewClickstreamSinkFilesWithDefaults instantiates a new ClickstreamSinkFiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateFrom

`func (o *ClickstreamSinkFiles) GetDateFrom() string`

GetDateFrom returns the DateFrom field if non-nil, zero value otherwise.

### GetDateFromOk

`func (o *ClickstreamSinkFiles) GetDateFromOk() (*string, bool)`

GetDateFromOk returns a tuple with the DateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFrom

`func (o *ClickstreamSinkFiles) SetDateFrom(v string)`

SetDateFrom sets DateFrom field to given value.

### HasDateFrom

`func (o *ClickstreamSinkFiles) HasDateFrom() bool`

HasDateFrom returns a boolean if a field has been set.

### GetObjects

`func (o *ClickstreamSinkFiles) GetObjects() []ClickstreamSinkFile`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *ClickstreamSinkFiles) GetObjectsOk() (*[]ClickstreamSinkFile, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *ClickstreamSinkFiles) SetObjects(v []ClickstreamSinkFile)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *ClickstreamSinkFiles) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetId

`func (o *ClickstreamSinkFiles) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClickstreamSinkFiles) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClickstreamSinkFiles) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClickstreamSinkFiles) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClickstreamId

`func (o *ClickstreamSinkFiles) GetClickstreamId() string`

GetClickstreamId returns the ClickstreamId field if non-nil, zero value otherwise.

### GetClickstreamIdOk

`func (o *ClickstreamSinkFiles) GetClickstreamIdOk() (*string, bool)`

GetClickstreamIdOk returns a tuple with the ClickstreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickstreamId

`func (o *ClickstreamSinkFiles) SetClickstreamId(v string)`

SetClickstreamId sets ClickstreamId field to given value.

### HasClickstreamId

`func (o *ClickstreamSinkFiles) HasClickstreamId() bool`

HasClickstreamId returns a boolean if a field has been set.

### GetDateTo

`func (o *ClickstreamSinkFiles) GetDateTo() string`

GetDateTo returns the DateTo field if non-nil, zero value otherwise.

### GetDateToOk

`func (o *ClickstreamSinkFiles) GetDateToOk() (*string, bool)`

GetDateToOk returns a tuple with the DateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTo

`func (o *ClickstreamSinkFiles) SetDateTo(v string)`

SetDateTo sets DateTo field to given value.

### HasDateTo

`func (o *ClickstreamSinkFiles) HasDateTo() bool`

HasDateTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


