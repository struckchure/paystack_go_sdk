# ProductCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of product | 
**Description** | **string** | The description of the product | 
**Price** | **int32** | Price should be in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR | 
**Currency** | **string** | Currency in which price is set. Allowed values are: NGN, GHS, ZAR or USD | 
**Limited** | Pointer to **bool** | Set to true if the product has limited stock. Leave as false if the product has unlimited stock | [optional] 
**Quantity** | Pointer to **int32** | Number of products in stock. Use if limited is true | [optional] 

## Methods

### NewProductCreate

`func NewProductCreate(name string, description string, price int32, currency string, ) *ProductCreate`

NewProductCreate instantiates a new ProductCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductCreateWithDefaults

`func NewProductCreateWithDefaults() *ProductCreate`

NewProductCreateWithDefaults instantiates a new ProductCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProductCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProductCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductCreate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPrice

`func (o *ProductCreate) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductCreate) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductCreate) SetPrice(v int32)`

SetPrice sets Price field to given value.


### GetCurrency

`func (o *ProductCreate) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ProductCreate) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ProductCreate) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetLimited

`func (o *ProductCreate) GetLimited() bool`

GetLimited returns the Limited field if non-nil, zero value otherwise.

### GetLimitedOk

`func (o *ProductCreate) GetLimitedOk() (*bool, bool)`

GetLimitedOk returns a tuple with the Limited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimited

`func (o *ProductCreate) SetLimited(v bool)`

SetLimited sets Limited field to given value.

### HasLimited

`func (o *ProductCreate) HasLimited() bool`

HasLimited returns a boolean if a field has been set.

### GetQuantity

`func (o *ProductCreate) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ProductCreate) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ProductCreate) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ProductCreate) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


