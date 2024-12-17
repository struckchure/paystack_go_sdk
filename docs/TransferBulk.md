# TransferBulk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to **string** | Where should we transfer from? Only balance is allowed for now | [optional] 
**Transfers** | Pointer to [**[]TransferInitiate**](TransferInitiate.md) | A list of transfer object. Each object should contain amount, recipient, and reference | [optional] 

## Methods

### NewTransferBulk

`func NewTransferBulk() *TransferBulk`

NewTransferBulk instantiates a new TransferBulk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferBulkWithDefaults

`func NewTransferBulkWithDefaults() *TransferBulk`

NewTransferBulkWithDefaults instantiates a new TransferBulk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *TransferBulk) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TransferBulk) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TransferBulk) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *TransferBulk) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTransfers

`func (o *TransferBulk) GetTransfers() []TransferInitiate`

GetTransfers returns the Transfers field if non-nil, zero value otherwise.

### GetTransfersOk

`func (o *TransferBulk) GetTransfersOk() (*[]TransferInitiate, bool)`

GetTransfersOk returns a tuple with the Transfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfers

`func (o *TransferBulk) SetTransfers(v []TransferInitiate)`

SetTransfers sets Transfers field to given value.

### HasTransfers

`func (o *TransferBulk) HasTransfers() bool`

HasTransfers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


