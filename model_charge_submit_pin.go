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

// checks if the ChargeSubmitPin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargeSubmitPin{}

// ChargeSubmitPin struct for ChargeSubmitPin
type ChargeSubmitPin struct {
	// Customer's PIN
	Pin string `json:"pin"`
	// Transaction reference that requires the PIN
	Reference string `json:"reference"`
}

type _ChargeSubmitPin ChargeSubmitPin

// NewChargeSubmitPin instantiates a new ChargeSubmitPin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargeSubmitPin(pin string, reference string) *ChargeSubmitPin {
	this := ChargeSubmitPin{}
	this.Pin = pin
	this.Reference = reference
	return &this
}

// NewChargeSubmitPinWithDefaults instantiates a new ChargeSubmitPin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargeSubmitPinWithDefaults() *ChargeSubmitPin {
	this := ChargeSubmitPin{}
	return &this
}

// GetPin returns the Pin field value
func (o *ChargeSubmitPin) GetPin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pin
}

// GetPinOk returns a tuple with the Pin field value
// and a boolean to check if the value has been set.
func (o *ChargeSubmitPin) GetPinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pin, true
}

// SetPin sets field value
func (o *ChargeSubmitPin) SetPin(v string) {
	o.Pin = v
}

// GetReference returns the Reference field value
func (o *ChargeSubmitPin) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *ChargeSubmitPin) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *ChargeSubmitPin) SetReference(v string) {
	o.Reference = v
}

func (o ChargeSubmitPin) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargeSubmitPin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pin"] = o.Pin
	toSerialize["reference"] = o.Reference
	return toSerialize, nil
}

func (o *ChargeSubmitPin) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pin",
		"reference",
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

	varChargeSubmitPin := _ChargeSubmitPin{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChargeSubmitPin)

	if err != nil {
		return err
	}

	*o = ChargeSubmitPin(varChargeSubmitPin)

	return err
}

type NullableChargeSubmitPin struct {
	value *ChargeSubmitPin
	isSet bool
}

func (v NullableChargeSubmitPin) Get() *ChargeSubmitPin {
	return v.value
}

func (v *NullableChargeSubmitPin) Set(val *ChargeSubmitPin) {
	v.value = val
	v.isSet = true
}

func (v NullableChargeSubmitPin) IsSet() bool {
	return v.isSet
}

func (v *NullableChargeSubmitPin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargeSubmitPin(val *ChargeSubmitPin) *NullableChargeSubmitPin {
	return &NullableChargeSubmitPin{value: val, isSet: true}
}

func (v NullableChargeSubmitPin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargeSubmitPin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

