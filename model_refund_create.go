/*
Paystack

The OpenAPI specification of the Paystack API that merchants and developers can harness to build financial solutions in Africa.

API version: 1.0.0
Contact: techsupport@paystack.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package paystack

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RefundCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefundCreate{}

// RefundCreate struct for RefundCreate
type RefundCreate struct {
	// Transaction reference or id
	Transaction string `json:"transaction"`
	// Amount ( in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR ) to be refunded to the customer.  Amount cannot be more than the original transaction amount
	Amount *int32 `json:"amount,omitempty"`
	// Three-letter ISO currency. Allowed values are NGN, GHS, ZAR or USD
	Currency *string `json:"currency,omitempty"`
	// Customer reason
	CustomerNote *string `json:"customer_note,omitempty"`
	// Merchant reason
	MerchantNote *string `json:"merchant_note,omitempty"`
}

type _RefundCreate RefundCreate

// NewRefundCreate instantiates a new RefundCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefundCreate(transaction string) *RefundCreate {
	this := RefundCreate{}
	this.Transaction = transaction
	return &this
}

// NewRefundCreateWithDefaults instantiates a new RefundCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefundCreateWithDefaults() *RefundCreate {
	this := RefundCreate{}
	return &this
}

// GetTransaction returns the Transaction field value
func (o *RefundCreate) GetTransaction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value
// and a boolean to check if the value has been set.
func (o *RefundCreate) GetTransactionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transaction, true
}

// SetTransaction sets field value
func (o *RefundCreate) SetTransaction(v string) {
	o.Transaction = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *RefundCreate) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundCreate) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *RefundCreate) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *RefundCreate) SetAmount(v int32) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *RefundCreate) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundCreate) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *RefundCreate) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *RefundCreate) SetCurrency(v string) {
	o.Currency = &v
}

// GetCustomerNote returns the CustomerNote field value if set, zero value otherwise.
func (o *RefundCreate) GetCustomerNote() string {
	if o == nil || IsNil(o.CustomerNote) {
		var ret string
		return ret
	}
	return *o.CustomerNote
}

// GetCustomerNoteOk returns a tuple with the CustomerNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundCreate) GetCustomerNoteOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerNote) {
		return nil, false
	}
	return o.CustomerNote, true
}

// HasCustomerNote returns a boolean if a field has been set.
func (o *RefundCreate) HasCustomerNote() bool {
	if o != nil && !IsNil(o.CustomerNote) {
		return true
	}

	return false
}

// SetCustomerNote gets a reference to the given string and assigns it to the CustomerNote field.
func (o *RefundCreate) SetCustomerNote(v string) {
	o.CustomerNote = &v
}

// GetMerchantNote returns the MerchantNote field value if set, zero value otherwise.
func (o *RefundCreate) GetMerchantNote() string {
	if o == nil || IsNil(o.MerchantNote) {
		var ret string
		return ret
	}
	return *o.MerchantNote
}

// GetMerchantNoteOk returns a tuple with the MerchantNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundCreate) GetMerchantNoteOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantNote) {
		return nil, false
	}
	return o.MerchantNote, true
}

// HasMerchantNote returns a boolean if a field has been set.
func (o *RefundCreate) HasMerchantNote() bool {
	if o != nil && !IsNil(o.MerchantNote) {
		return true
	}

	return false
}

// SetMerchantNote gets a reference to the given string and assigns it to the MerchantNote field.
func (o *RefundCreate) SetMerchantNote(v string) {
	o.MerchantNote = &v
}

func (o RefundCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefundCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transaction"] = o.Transaction
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.CustomerNote) {
		toSerialize["customer_note"] = o.CustomerNote
	}
	if !IsNil(o.MerchantNote) {
		toSerialize["merchant_note"] = o.MerchantNote
	}
	return toSerialize, nil
}

func (o *RefundCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transaction",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRefundCreate := _RefundCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRefundCreate)

	if err != nil {
		return err
	}

	*o = RefundCreate(varRefundCreate)

	return err
}

type NullableRefundCreate struct {
	value *RefundCreate
	isSet bool
}

func (v NullableRefundCreate) Get() *RefundCreate {
	return v.value
}

func (v *NullableRefundCreate) Set(val *RefundCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableRefundCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableRefundCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefundCreate(val *RefundCreate) *NullableRefundCreate {
	return &NullableRefundCreate{value: val, isSet: true}
}

func (v NullableRefundCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefundCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

