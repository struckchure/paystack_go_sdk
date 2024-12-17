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
	"time"
	"bytes"
	"fmt"
)

// checks if the ChargeCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargeCreateRequest{}

// ChargeCreateRequest struct for ChargeCreateRequest
type ChargeCreateRequest struct {
	// Customer's email address
	Email string `json:"email"`
	// Amount should be in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR
	Amount string `json:"amount"`
	// An authorization code to charge.
	AuthorizationCode *string `json:"authorization_code,omitempty"`
	// 4-digit PIN (send with a non-reusable authorization code)
	Pin *string `json:"pin,omitempty"`
	// Unique transaction reference. Only -, .`, = and alphanumeric characters allowed.
	Reference *string `json:"reference,omitempty"`
	// The customer's birthday in the format YYYY-MM-DD e.g 2017-05-16
	Birthday *time.Time `json:"birthday,omitempty"`
	// This is the unique identifier of the device a user uses in making payment.  Only -, .`, = and alphanumeric characters are allowed.
	DeviceId *string `json:"device_id,omitempty"`
	// Stringified JSON object of custom data
	Metadata *string `json:"metadata,omitempty"`
	Bank *Bank `json:"bank,omitempty"`
	MobileMoney *MobileMoney `json:"mobile_money,omitempty"`
	Ussd *USSD `json:"ussd,omitempty"`
	Eft *EFT `json:"eft,omitempty"`
}

type _ChargeCreateRequest ChargeCreateRequest

// NewChargeCreateRequest instantiates a new ChargeCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargeCreateRequest(email string, amount string) *ChargeCreateRequest {
	this := ChargeCreateRequest{}
	this.Email = email
	this.Amount = amount
	return &this
}

// NewChargeCreateRequestWithDefaults instantiates a new ChargeCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargeCreateRequestWithDefaults() *ChargeCreateRequest {
	this := ChargeCreateRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *ChargeCreateRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ChargeCreateRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ChargeCreateRequest) SetEmail(v string) {
	o.Email = v
}

// GetAmount returns the Amount field value
func (o *ChargeCreateRequest) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ChargeCreateRequest) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ChargeCreateRequest) SetAmount(v string) {
	o.Amount = v
}

// GetAuthorizationCode returns the AuthorizationCode field value if set, zero value otherwise.
func (o *ChargeCreateRequest) GetAuthorizationCode() string {
	if o == nil || IsNil(o.AuthorizationCode) {
		var ret string
		return ret
	}
	return *o.AuthorizationCode
}

// GetAuthorizationCodeOk returns a tuple with the AuthorizationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeCreateRequest) GetAuthorizationCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorizationCode) {
		return nil, false
	}
	return o.AuthorizationCode, true
}

// HasAuthorizationCode returns a boolean if a field has been set.
func (o *ChargeCreateRequest) HasAuthorizationCode() bool {
	if o != nil && !IsNil(o.AuthorizationCode) {
		return true
	}

	return false
}

// SetAuthorizationCode gets a reference to the given string and assigns it to the AuthorizationCode field.
func (o *ChargeCreateRequest) SetAuthorizationCode(v string) {
	o.AuthorizationCode = &v
}

// GetPin returns the Pin field value if set, zero value otherwise.
func (o *ChargeCreateRequest) GetPin() string {
	if o == nil || IsNil(o.Pin) {
		var ret string
		return ret
	}
	return *o.Pin
}

// GetPinOk returns a tuple with the Pin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeCreateRequest) GetPinOk() (*string, bool) {
	if o == nil || IsNil(o.Pin) {
		return nil, false
	}
	return o.Pin, true
}

// HasPin returns a boolean if a field has been set.
func (o *ChargeCreateRequest) HasPin() bool {
	if o != nil && !IsNil(o.Pin) {
		return true
	}

	return false
}

// SetPin gets a reference to the given string and assigns it to the Pin field.
func (o *ChargeCreateRequest) SetPin(v string) {
	o.Pin = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *ChargeCreateRequest) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeCreateRequest) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *ChargeCreateRequest) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *ChargeCreateRequest) SetReference(v string) {
	o.Reference = &v
}

// GetBirthday returns the Birthday field value if set, zero value otherwise.
func (o *ChargeCreateRequest) GetBirthday() time.Time {
	if o == nil || IsNil(o.Birthday) {
		var ret time.Time
		return ret
	}
	return *o.Birthday
}

// GetBirthdayOk returns a tuple with the Birthday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeCreateRequest) GetBirthdayOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Birthday) {
		return nil, false
	}
	return o.Birthday, true
}

// HasBirthday returns a boolean if a field has been set.
func (o *ChargeCreateRequest) HasBirthday() bool {
	if o != nil && !IsNil(o.Birthday) {
		return true
	}

	return false
}

// SetBirthday gets a reference to the given time.Time and assigns it to the Birthday field.
func (o *ChargeCreateRequest) SetBirthday(v time.Time) {
	o.Birthday = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *ChargeCreateRequest) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeCreateRequest) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *ChargeCreateRequest) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *ChargeCreateRequest) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ChargeCreateRequest) GetMetadata() string {
	if o == nil || IsNil(o.Metadata) {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeCreateRequest) GetMetadataOk() (*string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ChargeCreateRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *ChargeCreateRequest) SetMetadata(v string) {
	o.Metadata = &v
}

// GetBank returns the Bank field value if set, zero value otherwise.
func (o *ChargeCreateRequest) GetBank() Bank {
	if o == nil || IsNil(o.Bank) {
		var ret Bank
		return ret
	}
	return *o.Bank
}

// GetBankOk returns a tuple with the Bank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeCreateRequest) GetBankOk() (*Bank, bool) {
	if o == nil || IsNil(o.Bank) {
		return nil, false
	}
	return o.Bank, true
}

// HasBank returns a boolean if a field has been set.
func (o *ChargeCreateRequest) HasBank() bool {
	if o != nil && !IsNil(o.Bank) {
		return true
	}

	return false
}

// SetBank gets a reference to the given Bank and assigns it to the Bank field.
func (o *ChargeCreateRequest) SetBank(v Bank) {
	o.Bank = &v
}

// GetMobileMoney returns the MobileMoney field value if set, zero value otherwise.
func (o *ChargeCreateRequest) GetMobileMoney() MobileMoney {
	if o == nil || IsNil(o.MobileMoney) {
		var ret MobileMoney
		return ret
	}
	return *o.MobileMoney
}

// GetMobileMoneyOk returns a tuple with the MobileMoney field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeCreateRequest) GetMobileMoneyOk() (*MobileMoney, bool) {
	if o == nil || IsNil(o.MobileMoney) {
		return nil, false
	}
	return o.MobileMoney, true
}

// HasMobileMoney returns a boolean if a field has been set.
func (o *ChargeCreateRequest) HasMobileMoney() bool {
	if o != nil && !IsNil(o.MobileMoney) {
		return true
	}

	return false
}

// SetMobileMoney gets a reference to the given MobileMoney and assigns it to the MobileMoney field.
func (o *ChargeCreateRequest) SetMobileMoney(v MobileMoney) {
	o.MobileMoney = &v
}

// GetUssd returns the Ussd field value if set, zero value otherwise.
func (o *ChargeCreateRequest) GetUssd() USSD {
	if o == nil || IsNil(o.Ussd) {
		var ret USSD
		return ret
	}
	return *o.Ussd
}

// GetUssdOk returns a tuple with the Ussd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeCreateRequest) GetUssdOk() (*USSD, bool) {
	if o == nil || IsNil(o.Ussd) {
		return nil, false
	}
	return o.Ussd, true
}

// HasUssd returns a boolean if a field has been set.
func (o *ChargeCreateRequest) HasUssd() bool {
	if o != nil && !IsNil(o.Ussd) {
		return true
	}

	return false
}

// SetUssd gets a reference to the given USSD and assigns it to the Ussd field.
func (o *ChargeCreateRequest) SetUssd(v USSD) {
	o.Ussd = &v
}

// GetEft returns the Eft field value if set, zero value otherwise.
func (o *ChargeCreateRequest) GetEft() EFT {
	if o == nil || IsNil(o.Eft) {
		var ret EFT
		return ret
	}
	return *o.Eft
}

// GetEftOk returns a tuple with the Eft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeCreateRequest) GetEftOk() (*EFT, bool) {
	if o == nil || IsNil(o.Eft) {
		return nil, false
	}
	return o.Eft, true
}

// HasEft returns a boolean if a field has been set.
func (o *ChargeCreateRequest) HasEft() bool {
	if o != nil && !IsNil(o.Eft) {
		return true
	}

	return false
}

// SetEft gets a reference to the given EFT and assigns it to the Eft field.
func (o *ChargeCreateRequest) SetEft(v EFT) {
	o.Eft = &v
}

func (o ChargeCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargeCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["amount"] = o.Amount
	if !IsNil(o.AuthorizationCode) {
		toSerialize["authorization_code"] = o.AuthorizationCode
	}
	if !IsNil(o.Pin) {
		toSerialize["pin"] = o.Pin
	}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !IsNil(o.Birthday) {
		toSerialize["birthday"] = o.Birthday
	}
	if !IsNil(o.DeviceId) {
		toSerialize["device_id"] = o.DeviceId
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Bank) {
		toSerialize["bank"] = o.Bank
	}
	if !IsNil(o.MobileMoney) {
		toSerialize["mobile_money"] = o.MobileMoney
	}
	if !IsNil(o.Ussd) {
		toSerialize["ussd"] = o.Ussd
	}
	if !IsNil(o.Eft) {
		toSerialize["eft"] = o.Eft
	}
	return toSerialize, nil
}

func (o *ChargeCreateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"amount",
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

	varChargeCreateRequest := _ChargeCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChargeCreateRequest)

	if err != nil {
		return err
	}

	*o = ChargeCreateRequest(varChargeCreateRequest)

	return err
}

type NullableChargeCreateRequest struct {
	value *ChargeCreateRequest
	isSet bool
}

func (v NullableChargeCreateRequest) Get() *ChargeCreateRequest {
	return v.value
}

func (v *NullableChargeCreateRequest) Set(val *ChargeCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChargeCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChargeCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargeCreateRequest(val *ChargeCreateRequest) *NullableChargeCreateRequest {
	return &NullableChargeCreateRequest{value: val, isSet: true}
}

func (v NullableChargeCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargeCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


