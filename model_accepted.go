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

// checks if the Accepted type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Accepted{}

// Accepted struct for Accepted
type Accepted struct {
	Status *bool `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewAccepted instantiates a new Accepted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccepted() *Accepted {
	this := Accepted{}
	return &this
}

// NewAcceptedWithDefaults instantiates a new Accepted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcceptedWithDefaults() *Accepted {
	this := Accepted{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Accepted) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Accepted) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Accepted) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *Accepted) SetStatus(v bool) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Accepted) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Accepted) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Accepted) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Accepted) SetMessage(v string) {
	o.Message = &v
}

func (o Accepted) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Accepted) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableAccepted struct {
	value *Accepted
	isSet bool
}

func (v NullableAccepted) Get() *Accepted {
	return v.value
}

func (v *NullableAccepted) Set(val *Accepted) {
	v.value = val
	v.isSet = true
}

func (v NullableAccepted) IsSet() bool {
	return v.isSet
}

func (v *NullableAccepted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccepted(val *Accepted) *NullableAccepted {
	return &NullableAccepted{value: val, isSet: true}
}

func (v NullableAccepted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccepted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

