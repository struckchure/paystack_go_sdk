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

// checks if the DedicatedVirtualAccountSplit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DedicatedVirtualAccountSplit{}

// DedicatedVirtualAccountSplit struct for DedicatedVirtualAccountSplit
type DedicatedVirtualAccountSplit struct {
	// Valid Dedicated virtual account
	AccountNumber string `json:"account_number"`
	// Subaccount code of the account you want to split the transaction with
	Subaccount *string `json:"subaccount,omitempty"`
	// Split code consisting of the lists of accounts you want to split the transaction with
	SplitCode *string `json:"split_code,omitempty"`
}

type _DedicatedVirtualAccountSplit DedicatedVirtualAccountSplit

// NewDedicatedVirtualAccountSplit instantiates a new DedicatedVirtualAccountSplit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDedicatedVirtualAccountSplit(accountNumber string) *DedicatedVirtualAccountSplit {
	this := DedicatedVirtualAccountSplit{}
	this.AccountNumber = accountNumber
	return &this
}

// NewDedicatedVirtualAccountSplitWithDefaults instantiates a new DedicatedVirtualAccountSplit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDedicatedVirtualAccountSplitWithDefaults() *DedicatedVirtualAccountSplit {
	this := DedicatedVirtualAccountSplit{}
	return &this
}

// GetAccountNumber returns the AccountNumber field value
func (o *DedicatedVirtualAccountSplit) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *DedicatedVirtualAccountSplit) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *DedicatedVirtualAccountSplit) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetSubaccount returns the Subaccount field value if set, zero value otherwise.
func (o *DedicatedVirtualAccountSplit) GetSubaccount() string {
	if o == nil || IsNil(o.Subaccount) {
		var ret string
		return ret
	}
	return *o.Subaccount
}

// GetSubaccountOk returns a tuple with the Subaccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DedicatedVirtualAccountSplit) GetSubaccountOk() (*string, bool) {
	if o == nil || IsNil(o.Subaccount) {
		return nil, false
	}
	return o.Subaccount, true
}

// HasSubaccount returns a boolean if a field has been set.
func (o *DedicatedVirtualAccountSplit) HasSubaccount() bool {
	if o != nil && !IsNil(o.Subaccount) {
		return true
	}

	return false
}

// SetSubaccount gets a reference to the given string and assigns it to the Subaccount field.
func (o *DedicatedVirtualAccountSplit) SetSubaccount(v string) {
	o.Subaccount = &v
}

// GetSplitCode returns the SplitCode field value if set, zero value otherwise.
func (o *DedicatedVirtualAccountSplit) GetSplitCode() string {
	if o == nil || IsNil(o.SplitCode) {
		var ret string
		return ret
	}
	return *o.SplitCode
}

// GetSplitCodeOk returns a tuple with the SplitCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DedicatedVirtualAccountSplit) GetSplitCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SplitCode) {
		return nil, false
	}
	return o.SplitCode, true
}

// HasSplitCode returns a boolean if a field has been set.
func (o *DedicatedVirtualAccountSplit) HasSplitCode() bool {
	if o != nil && !IsNil(o.SplitCode) {
		return true
	}

	return false
}

// SetSplitCode gets a reference to the given string and assigns it to the SplitCode field.
func (o *DedicatedVirtualAccountSplit) SetSplitCode(v string) {
	o.SplitCode = &v
}

func (o DedicatedVirtualAccountSplit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DedicatedVirtualAccountSplit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_number"] = o.AccountNumber
	if !IsNil(o.Subaccount) {
		toSerialize["subaccount"] = o.Subaccount
	}
	if !IsNil(o.SplitCode) {
		toSerialize["split_code"] = o.SplitCode
	}
	return toSerialize, nil
}

func (o *DedicatedVirtualAccountSplit) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_number",
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

	varDedicatedVirtualAccountSplit := _DedicatedVirtualAccountSplit{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDedicatedVirtualAccountSplit)

	if err != nil {
		return err
	}

	*o = DedicatedVirtualAccountSplit(varDedicatedVirtualAccountSplit)

	return err
}

type NullableDedicatedVirtualAccountSplit struct {
	value *DedicatedVirtualAccountSplit
	isSet bool
}

func (v NullableDedicatedVirtualAccountSplit) Get() *DedicatedVirtualAccountSplit {
	return v.value
}

func (v *NullableDedicatedVirtualAccountSplit) Set(val *DedicatedVirtualAccountSplit) {
	v.value = val
	v.isSet = true
}

func (v NullableDedicatedVirtualAccountSplit) IsSet() bool {
	return v.isSet
}

func (v *NullableDedicatedVirtualAccountSplit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDedicatedVirtualAccountSplit(val *DedicatedVirtualAccountSplit) *NullableDedicatedVirtualAccountSplit {
	return &NullableDedicatedVirtualAccountSplit{value: val, isSet: true}
}

func (v NullableDedicatedVirtualAccountSplit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDedicatedVirtualAccountSplit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


