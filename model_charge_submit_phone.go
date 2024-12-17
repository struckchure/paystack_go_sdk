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

// checks if the ChargeSubmitPhone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargeSubmitPhone{}

// ChargeSubmitPhone struct for ChargeSubmitPhone
type ChargeSubmitPhone struct {
	// Customer's mobile number
	Phone string `json:"phone"`
	// The reference of the ongoing transaction
	Reference string `json:"reference"`
}

type _ChargeSubmitPhone ChargeSubmitPhone

// NewChargeSubmitPhone instantiates a new ChargeSubmitPhone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargeSubmitPhone(phone string, reference string) *ChargeSubmitPhone {
	this := ChargeSubmitPhone{}
	this.Phone = phone
	this.Reference = reference
	return &this
}

// NewChargeSubmitPhoneWithDefaults instantiates a new ChargeSubmitPhone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargeSubmitPhoneWithDefaults() *ChargeSubmitPhone {
	this := ChargeSubmitPhone{}
	return &this
}

// GetPhone returns the Phone field value
func (o *ChargeSubmitPhone) GetPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
func (o *ChargeSubmitPhone) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phone, true
}

// SetPhone sets field value
func (o *ChargeSubmitPhone) SetPhone(v string) {
	o.Phone = v
}

// GetReference returns the Reference field value
func (o *ChargeSubmitPhone) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *ChargeSubmitPhone) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *ChargeSubmitPhone) SetReference(v string) {
	o.Reference = v
}

func (o ChargeSubmitPhone) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargeSubmitPhone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["phone"] = o.Phone
	toSerialize["reference"] = o.Reference
	return toSerialize, nil
}

func (o *ChargeSubmitPhone) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"phone",
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

	varChargeSubmitPhone := _ChargeSubmitPhone{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChargeSubmitPhone)

	if err != nil {
		return err
	}

	*o = ChargeSubmitPhone(varChargeSubmitPhone)

	return err
}

type NullableChargeSubmitPhone struct {
	value *ChargeSubmitPhone
	isSet bool
}

func (v NullableChargeSubmitPhone) Get() *ChargeSubmitPhone {
	return v.value
}

func (v *NullableChargeSubmitPhone) Set(val *ChargeSubmitPhone) {
	v.value = val
	v.isSet = true
}

func (v NullableChargeSubmitPhone) IsSet() bool {
	return v.isSet
}

func (v *NullableChargeSubmitPhone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargeSubmitPhone(val *ChargeSubmitPhone) *NullableChargeSubmitPhone {
	return &NullableChargeSubmitPhone{value: val, isSet: true}
}

func (v NullableChargeSubmitPhone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargeSubmitPhone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


