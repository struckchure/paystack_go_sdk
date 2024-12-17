# CustomerValidate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** | Customer&#39;s first name | 
**LastName** | **string** | Customer&#39;s last name | 
**Type** | **string** | Predefined types of identification. e.g. (BVN) | 
**Value** | **string** | Customer&#39;s identification number | 
**Country** | **string** | 2 letter country code of identification issuer | 

## Methods

### NewCustomerValidate

`func NewCustomerValidate(firstName string, lastName string, type_ string, value string, country string, ) *CustomerValidate`

NewCustomerValidate instantiates a new CustomerValidate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerValidateWithDefaults

`func NewCustomerValidateWithDefaults() *CustomerValidate`

NewCustomerValidateWithDefaults instantiates a new CustomerValidate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *CustomerValidate) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CustomerValidate) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CustomerValidate) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *CustomerValidate) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CustomerValidate) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CustomerValidate) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetType

`func (o *CustomerValidate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerValidate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerValidate) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *CustomerValidate) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomerValidate) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomerValidate) SetValue(v string)`

SetValue sets Value field to given value.


### GetCountry

`func (o *CustomerValidate) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CustomerValidate) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CustomerValidate) SetCountry(v string)`

SetCountry sets Country field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


