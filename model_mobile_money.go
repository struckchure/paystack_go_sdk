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
)

// checks if the MobileMoney type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MobileMoney{}

// MobileMoney struct for MobileMoney
type MobileMoney struct {
	// Customer's phone number
	Phone *string `json:"phone,omitempty"`
	// The telco provider of customer's phone number. This can be fetched from the List Bank endpoint
	Provider *string `json:"provider,omitempty"`
}

// NewMobileMoney instantiates a new MobileMoney object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobileMoney() *MobileMoney {
	this := MobileMoney{}
	return &this
}

// NewMobileMoneyWithDefaults instantiates a new MobileMoney object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobileMoneyWithDefaults() *MobileMoney {
	this := MobileMoney{}
	return &this
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *MobileMoney) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileMoney) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *MobileMoney) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *MobileMoney) SetPhone(v string) {
	o.Phone = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *MobileMoney) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileMoney) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *MobileMoney) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *MobileMoney) SetProvider(v string) {
	o.Provider = &v
}

func (o MobileMoney) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MobileMoney) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	return toSerialize, nil
}

type NullableMobileMoney struct {
	value *MobileMoney
	isSet bool
}

func (v NullableMobileMoney) Get() *MobileMoney {
	return v.value
}

func (v *NullableMobileMoney) Set(val *MobileMoney) {
	v.value = val
	v.isSet = true
}

func (v NullableMobileMoney) IsSet() bool {
	return v.isSet
}

func (v *NullableMobileMoney) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobileMoney(val *MobileMoney) *NullableMobileMoney {
	return &NullableMobileMoney{value: val, isSet: true}
}

func (v NullableMobileMoney) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobileMoney) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

