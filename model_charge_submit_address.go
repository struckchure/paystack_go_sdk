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

// checks if the ChargeSubmitAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargeSubmitAddress{}

// ChargeSubmitAddress struct for ChargeSubmitAddress
type ChargeSubmitAddress struct {
	// Customer's address
	Address string `json:"address"`
	// Customer's city
	City string `json:"city"`
	// Customer's state
	State string `json:"state"`
	// Customer's zipcode
	Zipcode string `json:"zipcode"`
	// The reference of the ongoing transaction
	Reference string `json:"reference"`
}

type _ChargeSubmitAddress ChargeSubmitAddress

// NewChargeSubmitAddress instantiates a new ChargeSubmitAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargeSubmitAddress(address string, city string, state string, zipcode string, reference string) *ChargeSubmitAddress {
	this := ChargeSubmitAddress{}
	this.Address = address
	this.City = city
	this.State = state
	this.Zipcode = zipcode
	this.Reference = reference
	return &this
}

// NewChargeSubmitAddressWithDefaults instantiates a new ChargeSubmitAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargeSubmitAddressWithDefaults() *ChargeSubmitAddress {
	this := ChargeSubmitAddress{}
	return &this
}

// GetAddress returns the Address field value
func (o *ChargeSubmitAddress) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *ChargeSubmitAddress) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *ChargeSubmitAddress) SetAddress(v string) {
	o.Address = v
}

// GetCity returns the City field value
func (o *ChargeSubmitAddress) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *ChargeSubmitAddress) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *ChargeSubmitAddress) SetCity(v string) {
	o.City = v
}

// GetState returns the State field value
func (o *ChargeSubmitAddress) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ChargeSubmitAddress) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ChargeSubmitAddress) SetState(v string) {
	o.State = v
}

// GetZipcode returns the Zipcode field value
func (o *ChargeSubmitAddress) GetZipcode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Zipcode
}

// GetZipcodeOk returns a tuple with the Zipcode field value
// and a boolean to check if the value has been set.
func (o *ChargeSubmitAddress) GetZipcodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Zipcode, true
}

// SetZipcode sets field value
func (o *ChargeSubmitAddress) SetZipcode(v string) {
	o.Zipcode = v
}

// GetReference returns the Reference field value
func (o *ChargeSubmitAddress) GetReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
func (o *ChargeSubmitAddress) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *ChargeSubmitAddress) SetReference(v string) {
	o.Reference = v
}

func (o ChargeSubmitAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargeSubmitAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["city"] = o.City
	toSerialize["state"] = o.State
	toSerialize["zipcode"] = o.Zipcode
	toSerialize["reference"] = o.Reference
	return toSerialize, nil
}

func (o *ChargeSubmitAddress) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
		"city",
		"state",
		"zipcode",
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

	varChargeSubmitAddress := _ChargeSubmitAddress{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChargeSubmitAddress)

	if err != nil {
		return err
	}

	*o = ChargeSubmitAddress(varChargeSubmitAddress)

	return err
}

type NullableChargeSubmitAddress struct {
	value *ChargeSubmitAddress
	isSet bool
}

func (v NullableChargeSubmitAddress) Get() *ChargeSubmitAddress {
	return v.value
}

func (v *NullableChargeSubmitAddress) Set(val *ChargeSubmitAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableChargeSubmitAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableChargeSubmitAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargeSubmitAddress(val *ChargeSubmitAddress) *NullableChargeSubmitAddress {
	return &NullableChargeSubmitAddress{value: val, isSet: true}
}

func (v NullableChargeSubmitAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargeSubmitAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


