# BulkChargeInitiate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authorization** | **string** | Customer&#39;s card authorization code | 
**Amount** | **string** | Amount to charge on the authorization | 

## Methods

### NewBulkChargeInitiate

`func NewBulkChargeInitiate(authorization string, amount string, ) *BulkChargeInitiate`

NewBulkChargeInitiate instantiates a new BulkChargeInitiate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkChargeInitiateWithDefaults

`func NewBulkChargeInitiateWithDefaults() *BulkChargeInitiate`

NewBulkChargeInitiateWithDefaults instantiates a new BulkChargeInitiate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorization

`func (o *BulkChargeInitiate) GetAuthorization() string`

GetAuthorization returns the Authorization field if non-nil, zero value otherwise.

### GetAuthorizationOk

`func (o *BulkChargeInitiate) GetAuthorizationOk() (*string, bool)`

GetAuthorizationOk returns a tuple with the Authorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorization

`func (o *BulkChargeInitiate) SetAuthorization(v string)`

SetAuthorization sets Authorization field to given value.


### GetAmount

`func (o *BulkChargeInitiate) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BulkChargeInitiate) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BulkChargeInitiate) SetAmount(v string)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


