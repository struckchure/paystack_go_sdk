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

// checks if the DisputeUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DisputeUpdate{}

// DisputeUpdate struct for DisputeUpdate
type DisputeUpdate struct {
	// The amount to refund, in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR
	RefundAmount string `json:"refund_amount"`
	// Filename of attachment returned via response from the Dispute upload URL
	UploadedFilename *string `json:"uploaded_filename,omitempty"`
}

type _DisputeUpdate DisputeUpdate

// NewDisputeUpdate instantiates a new DisputeUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisputeUpdate(refundAmount string) *DisputeUpdate {
	this := DisputeUpdate{}
	this.RefundAmount = refundAmount
	return &this
}

// NewDisputeUpdateWithDefaults instantiates a new DisputeUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisputeUpdateWithDefaults() *DisputeUpdate {
	this := DisputeUpdate{}
	return &this
}

// GetRefundAmount returns the RefundAmount field value
func (o *DisputeUpdate) GetRefundAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefundAmount
}

// GetRefundAmountOk returns a tuple with the RefundAmount field value
// and a boolean to check if the value has been set.
func (o *DisputeUpdate) GetRefundAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefundAmount, true
}

// SetRefundAmount sets field value
func (o *DisputeUpdate) SetRefundAmount(v string) {
	o.RefundAmount = v
}

// GetUploadedFilename returns the UploadedFilename field value if set, zero value otherwise.
func (o *DisputeUpdate) GetUploadedFilename() string {
	if o == nil || IsNil(o.UploadedFilename) {
		var ret string
		return ret
	}
	return *o.UploadedFilename
}

// GetUploadedFilenameOk returns a tuple with the UploadedFilename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisputeUpdate) GetUploadedFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.UploadedFilename) {
		return nil, false
	}
	return o.UploadedFilename, true
}

// HasUploadedFilename returns a boolean if a field has been set.
func (o *DisputeUpdate) HasUploadedFilename() bool {
	if o != nil && !IsNil(o.UploadedFilename) {
		return true
	}

	return false
}

// SetUploadedFilename gets a reference to the given string and assigns it to the UploadedFilename field.
func (o *DisputeUpdate) SetUploadedFilename(v string) {
	o.UploadedFilename = &v
}

func (o DisputeUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DisputeUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["refund_amount"] = o.RefundAmount
	if !IsNil(o.UploadedFilename) {
		toSerialize["uploaded_filename"] = o.UploadedFilename
	}
	return toSerialize, nil
}

func (o *DisputeUpdate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"refund_amount",
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

	varDisputeUpdate := _DisputeUpdate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDisputeUpdate)

	if err != nil {
		return err
	}

	*o = DisputeUpdate(varDisputeUpdate)

	return err
}

type NullableDisputeUpdate struct {
	value *DisputeUpdate
	isSet bool
}

func (v NullableDisputeUpdate) Get() *DisputeUpdate {
	return v.value
}

func (v *NullableDisputeUpdate) Set(val *DisputeUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableDisputeUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableDisputeUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisputeUpdate(val *DisputeUpdate) *NullableDisputeUpdate {
	return &NullableDisputeUpdate{value: val, isSet: true}
}

func (v NullableDisputeUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisputeUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


