# CustomerValidation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** | Customer&#39;s first name | 
**LastName** | **string** | Customer&#39;s last name | 
**Type** | **string** | Predefined types of identification. | 
**Country** | **string** | Two-letter country code of identification issuer | 
**Bvn** | **string** | Customer&#39;s Bank Verification Number | 
**BankCode** | **string** | You can get the list of bank codes by calling the List Banks endpoint (https://api.paystack.co/bank). | 
**AccountNumber** | **string** | Customer&#39;s bank account number. | 
**Value** | Pointer to **string** | Customer&#39;s identification number. Required if type is bvn | [optional] 

## Methods

### NewCustomerValidation

`func NewCustomerValidation(firstName string, lastName string, type_ string, country string, bvn string, bankCode string, accountNumber string, ) *CustomerValidation`

NewCustomerValidation instantiates a new CustomerValidation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerValidationWithDefaults

`func NewCustomerValidationWithDefaults() *CustomerValidation`

NewCustomerValidationWithDefaults instantiates a new CustomerValidation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *CustomerValidation) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CustomerValidation) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CustomerValidation) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *CustomerValidation) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CustomerValidation) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CustomerValidation) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetType

`func (o *CustomerValidation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerValidation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerValidation) SetType(v string)`

SetType sets Type field to given value.


### GetCountry

`func (o *CustomerValidation) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CustomerValidation) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CustomerValidation) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetBvn

`func (o *CustomerValidation) GetBvn() string`

GetBvn returns the Bvn field if non-nil, zero value otherwise.

### GetBvnOk

`func (o *CustomerValidation) GetBvnOk() (*string, bool)`

GetBvnOk returns a tuple with the Bvn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBvn

`func (o *CustomerValidation) SetBvn(v string)`

SetBvn sets Bvn field to given value.


### GetBankCode

`func (o *CustomerValidation) GetBankCode() string`

GetBankCode returns the BankCode field if non-nil, zero value otherwise.

### GetBankCodeOk

`func (o *CustomerValidation) GetBankCodeOk() (*string, bool)`

GetBankCodeOk returns a tuple with the BankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCode

`func (o *CustomerValidation) SetBankCode(v string)`

SetBankCode sets BankCode field to given value.


### GetAccountNumber

`func (o *CustomerValidation) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *CustomerValidation) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *CustomerValidation) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetValue

`func (o *CustomerValidation) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomerValidation) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomerValidation) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CustomerValidation) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


