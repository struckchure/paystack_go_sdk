# VerificationBVNMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | Bank Account Number | 
**BankCode** | **int32** | You can get the list of banks codes by calling the List Bank endpoint | 
**Bvn** | **string** | 11 digits Bank Verification Number | 
**FirstName** | Pointer to **string** | Customer&#39;s first name | [optional] 
**MiddleName** | Pointer to **string** | Customer&#39;s middle name | [optional] 
**LastName** | Pointer to **string** | Customer&#39;s last name | [optional] 

## Methods

### NewVerificationBVNMatch

`func NewVerificationBVNMatch(accountNumber string, bankCode int32, bvn string, ) *VerificationBVNMatch`

NewVerificationBVNMatch instantiates a new VerificationBVNMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationBVNMatchWithDefaults

`func NewVerificationBVNMatchWithDefaults() *VerificationBVNMatch`

NewVerificationBVNMatchWithDefaults instantiates a new VerificationBVNMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *VerificationBVNMatch) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *VerificationBVNMatch) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *VerificationBVNMatch) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetBankCode

`func (o *VerificationBVNMatch) GetBankCode() int32`

GetBankCode returns the BankCode field if non-nil, zero value otherwise.

### GetBankCodeOk

`func (o *VerificationBVNMatch) GetBankCodeOk() (*int32, bool)`

GetBankCodeOk returns a tuple with the BankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCode

`func (o *VerificationBVNMatch) SetBankCode(v int32)`

SetBankCode sets BankCode field to given value.


### GetBvn

`func (o *VerificationBVNMatch) GetBvn() string`

GetBvn returns the Bvn field if non-nil, zero value otherwise.

### GetBvnOk

`func (o *VerificationBVNMatch) GetBvnOk() (*string, bool)`

GetBvnOk returns a tuple with the Bvn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBvn

`func (o *VerificationBVNMatch) SetBvn(v string)`

SetBvn sets Bvn field to given value.


### GetFirstName

`func (o *VerificationBVNMatch) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *VerificationBVNMatch) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *VerificationBVNMatch) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *VerificationBVNMatch) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetMiddleName

`func (o *VerificationBVNMatch) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *VerificationBVNMatch) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *VerificationBVNMatch) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *VerificationBVNMatch) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetLastName

`func (o *VerificationBVNMatch) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *VerificationBVNMatch) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *VerificationBVNMatch) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *VerificationBVNMatch) HasLastName() bool`

HasLastName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


