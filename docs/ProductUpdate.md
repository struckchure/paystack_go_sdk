# ProductUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of product | [optional] 
**Description** | Pointer to **string** | The description of the product | [optional] 
**Price** | Pointer to **int32** | Price should be in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR | [optional] 
**Currency** | Pointer to **string** | Currency in which price is set. Allowed values are: NGN, GHS, ZAR or USD | [optional] 
**Limited** | Pointer to **bool** | Set to true if the product has limited stock. Leave as false if the product has unlimited stock | [optional] 
**Quantity** | Pointer to **int32** | Number of products in stock. Use if limited is true | [optional] 

## Methods

### NewProductUpdate

`func NewProductUpdate() *ProductUpdate`

NewProductUpdate instantiates a new ProductUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductUpdateWithDefaults

`func NewProductUpdateWithDefaults() *ProductUpdate`

NewProductUpdateWithDefaults instantiates a new ProductUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProductUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ProductUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrice

`func (o *ProductUpdate) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductUpdate) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductUpdate) SetPrice(v int32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ProductUpdate) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *ProductUpdate) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ProductUpdate) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ProductUpdate) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ProductUpdate) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetLimited

`func (o *ProductUpdate) GetLimited() bool`

GetLimited returns the Limited field if non-nil, zero value otherwise.

### GetLimitedOk

`func (o *ProductUpdate) GetLimitedOk() (*bool, bool)`

GetLimitedOk returns a tuple with the Limited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimited

`func (o *ProductUpdate) SetLimited(v bool)`

SetLimited sets Limited field to given value.

### HasLimited

`func (o *ProductUpdate) HasLimited() bool`

HasLimited returns a boolean if a field has been set.

### GetQuantity

`func (o *ProductUpdate) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ProductUpdate) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ProductUpdate) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ProductUpdate) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


