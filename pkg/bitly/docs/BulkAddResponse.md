# BulkAddResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**BulkAddResponseData**](BulkAddResponseData.md) |  | [optional] 

## Methods

### NewBulkAddResponse

`func NewBulkAddResponse() *BulkAddResponse`

NewBulkAddResponse instantiates a new BulkAddResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkAddResponseWithDefaults

`func NewBulkAddResponseWithDefaults() *BulkAddResponse`

NewBulkAddResponseWithDefaults instantiates a new BulkAddResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *BulkAddResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BulkAddResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BulkAddResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BulkAddResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *BulkAddResponse) GetData() BulkAddResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BulkAddResponse) GetDataOk() (*BulkAddResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BulkAddResponse) SetData(v BulkAddResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *BulkAddResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


