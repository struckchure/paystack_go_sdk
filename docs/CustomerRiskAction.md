# CustomerRiskAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customer** | **string** | Customer&#39;s code, or email address | 
**RiskAction** | Pointer to **string** | One of the possible risk actions [ default, allow, deny ]. allow to whitelist.  deny to blacklist. Customers start with a default risk action.  | [optional] 

## Methods

### NewCustomerRiskAction

`func NewCustomerRiskAction(customer string, ) *CustomerRiskAction`

NewCustomerRiskAction instantiates a new CustomerRiskAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerRiskActionWithDefaults

`func NewCustomerRiskActionWithDefaults() *CustomerRiskAction`

NewCustomerRiskActionWithDefaults instantiates a new CustomerRiskAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomer

`func (o *CustomerRiskAction) GetCustomer() string`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *CustomerRiskAction) GetCustomerOk() (*string, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *CustomerRiskAction) SetCustomer(v string)`

SetCustomer sets Customer field to given value.


### GetRiskAction

`func (o *CustomerRiskAction) GetRiskAction() string`

GetRiskAction returns the RiskAction field if non-nil, zero value otherwise.

### GetRiskActionOk

`func (o *CustomerRiskAction) GetRiskActionOk() (*string, bool)`

GetRiskActionOk returns a tuple with the RiskAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskAction

`func (o *CustomerRiskAction) SetRiskAction(v string)`

SetRiskAction sets RiskAction field to given value.

### HasRiskAction

`func (o *CustomerRiskAction) HasRiskAction() bool`

HasRiskAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


