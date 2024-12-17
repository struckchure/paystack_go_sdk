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

// checks if the CustomerDeactivateAuthorization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerDeactivateAuthorization{}

// CustomerDeactivateAuthorization struct for CustomerDeactivateAuthorization
type CustomerDeactivateAuthorization struct {
	// Authorization code to be deactivated
	AuthorizationCode string `json:"authorization_code"`
}

type _CustomerDeactivateAuthorization CustomerDeactivateAuthorization

// NewCustomerDeactivateAuthorization instantiates a new CustomerDeactivateAuthorization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerDeactivateAuthorization(authorizationCode string) *CustomerDeactivateAuthorization {
	this := CustomerDeactivateAuthorization{}
	this.AuthorizationCode = authorizationCode
	return &this
}

// NewCustomerDeactivateAuthorizationWithDefaults instantiates a new CustomerDeactivateAuthorization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerDeactivateAuthorizationWithDefaults() *CustomerDeactivateAuthorization {
	this := CustomerDeactivateAuthorization{}
	return &this
}

// GetAuthorizationCode returns the AuthorizationCode field value
func (o *CustomerDeactivateAuthorization) GetAuthorizationCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorizationCode
}

// GetAuthorizationCodeOk returns a tuple with the AuthorizationCode field value
// and a boolean to check if the value has been set.
func (o *CustomerDeactivateAuthorization) GetAuthorizationCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationCode, true
}

// SetAuthorizationCode sets field value
func (o *CustomerDeactivateAuthorization) SetAuthorizationCode(v string) {
	o.AuthorizationCode = v
}

func (o CustomerDeactivateAuthorization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerDeactivateAuthorization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authorization_code"] = o.AuthorizationCode
	return toSerialize, nil
}

func (o *CustomerDeactivateAuthorization) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authorization_code",
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

	varCustomerDeactivateAuthorization := _CustomerDeactivateAuthorization{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomerDeactivateAuthorization)

	if err != nil {
		return err
	}

	*o = CustomerDeactivateAuthorization(varCustomerDeactivateAuthorization)

	return err
}

type NullableCustomerDeactivateAuthorization struct {
	value *CustomerDeactivateAuthorization
	isSet bool
}

func (v NullableCustomerDeactivateAuthorization) Get() *CustomerDeactivateAuthorization {
	return v.value
}

func (v *NullableCustomerDeactivateAuthorization) Set(val *CustomerDeactivateAuthorization) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerDeactivateAuthorization) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerDeactivateAuthorization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerDeactivateAuthorization(val *CustomerDeactivateAuthorization) *NullableCustomerDeactivateAuthorization {
	return &NullableCustomerDeactivateAuthorization{value: val, isSet: true}
}

func (v NullableCustomerDeactivateAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerDeactivateAuthorization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


