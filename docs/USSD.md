# USSD

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The three-digit USSD code. One of, 737, 919, 822, 966 | [optional] 

## Methods

### NewUSSD

`func NewUSSD() *USSD`

NewUSSD instantiates a new USSD object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUSSDWithDefaults

`func NewUSSDWithDefaults() *USSD`

NewUSSDWithDefaults instantiates a new USSD object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *USSD) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *USSD) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *USSD) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *USSD) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


