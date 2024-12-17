# \ChargeAPI

All URIs are relative to *https://api.paystack.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChargeCheck**](ChargeAPI.md#ChargeCheck) | **Get** /charge/{reference} | Check pending charge
[**ChargeCreate**](ChargeAPI.md#ChargeCreate) | **Post** /charge | Create Charge
[**ChargeSubmitAddress**](ChargeAPI.md#ChargeSubmitAddress) | **Post** /charge/submit_address | Submit Address
[**ChargeSubmitBirthday**](ChargeAPI.md#ChargeSubmitBirthday) | **Post** /charge/submit_birthday | Submit Birthday
[**ChargeSubmitOtp**](ChargeAPI.md#ChargeSubmitOtp) | **Post** /charge/submit_otp | Submit OTP
[**ChargeSubmitPhone**](ChargeAPI.md#ChargeSubmitPhone) | **Post** /charge/submit_phone | Submit Phone
[**ChargeSubmitPin**](ChargeAPI.md#ChargeSubmitPin) | **Post** /charge/submit_pin | Submit PIN



## ChargeCheck

> Response ChargeCheck(ctx, reference).Execute()

Check pending charge

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	reference := "reference_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChargeAPI.ChargeCheck(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChargeAPI.ChargeCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChargeCheck`: Response
	fmt.Fprintf(os.Stdout, "Response from `ChargeAPI.ChargeCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChargeCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Response**](Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChargeCreate

> Response ChargeCreate(ctx).Email(email).Amount(amount).AuthorizationCode(authorizationCode).Pin(pin).Reference(reference).Birthday(birthday).DeviceId(deviceId).Metadata(metadata).Bank(bank).MobileMoney(mobileMoney).Ussd(ussd).Eft(eft).Execute()

Create Charge

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	email := "email_example" // string | Customer's email address
	amount := "amount_example" // string | Amount should be in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR
	authorizationCode := "authorizationCode_example" // string | An authorization code to charge. (optional)
	pin := "pin_example" // string | 4-digit PIN (send with a non-reusable authorization code) (optional)
	reference := "reference_example" // string | Unique transaction reference. Only -, .`, = and alphanumeric characters allowed. (optional)
	birthday := time.Now() // time.Time | The customer's birthday in the format YYYY-MM-DD e.g 2017-05-16 (optional)
	deviceId := "deviceId_example" // string | This is the unique identifier of the device a user uses in making payment.  Only -, .`, = and alphanumeric characters are allowed. (optional)
	metadata := "metadata_example" // string | Stringified JSON object of custom data (optional)
	bank := *openapiclient.NewBank() // Bank |  (optional)
	mobileMoney := *openapiclient.NewMobileMoney() // MobileMoney |  (optional)
	ussd := *openapiclient.NewUSSD() // USSD |  (optional)
	eft := *openapiclient.NewEFT() // EFT |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChargeAPI.ChargeCreate(context.Background()).Email(email).Amount(amount).AuthorizationCode(authorizationCode).Pin(pin).Reference(reference).Birthday(birthday).DeviceId(deviceId).Metadata(metadata).Bank(bank).MobileMoney(mobileMoney).Ussd(ussd).Eft(eft).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChargeAPI.ChargeCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChargeCreate`: Response
	fmt.Fprintf(os.Stdout, "Response from `ChargeAPI.ChargeCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChargeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Customer&#39;s email address | 
 **amount** | **string** | Amount should be in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR | 
 **authorizationCode** | **string** | An authorization code to charge. | 
 **pin** | **string** | 4-digit PIN (send with a non-reusable authorization code) | 
 **reference** | **string** | Unique transaction reference. Only -, .&#x60;, &#x3D; and alphanumeric characters allowed. | 
 **birthday** | **time.Time** | The customer&#39;s birthday in the format YYYY-MM-DD e.g 2017-05-16 | 
 **deviceId** | **string** | This is the unique identifier of the device a user uses in making payment.  Only -, .&#x60;, &#x3D; and alphanumeric characters are allowed. | 
 **metadata** | **string** | Stringified JSON object of custom data | 
 **bank** | [**Bank**](Bank.md) |  | 
 **mobileMoney** | [**MobileMoney**](MobileMoney.md) |  | 
 **ussd** | [**USSD**](USSD.md) |  | 
 **eft** | [**EFT**](EFT.md) |  | 

### Return type

[**Response**](Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChargeSubmitAddress

> Response ChargeSubmitAddress(ctx).Address(address).City(city).State(state).Zipcode(zipcode).Reference(reference).Execute()

Submit Address

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	address := "address_example" // string | Customer's address
	city := "city_example" // string | Customer's city
	state := "state_example" // string | Customer's state
	zipcode := "zipcode_example" // string | Customer's zipcode
	reference := "reference_example" // string | The reference of the ongoing transaction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChargeAPI.ChargeSubmitAddress(context.Background()).Address(address).City(city).State(state).Zipcode(zipcode).Reference(reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChargeAPI.ChargeSubmitAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChargeSubmitAddress`: Response
	fmt.Fprintf(os.Stdout, "Response from `ChargeAPI.ChargeSubmitAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChargeSubmitAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **address** | **string** | Customer&#39;s address | 
 **city** | **string** | Customer&#39;s city | 
 **state** | **string** | Customer&#39;s state | 
 **zipcode** | **string** | Customer&#39;s zipcode | 
 **reference** | **string** | The reference of the ongoing transaction | 

### Return type

[**Response**](Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChargeSubmitBirthday

> Response ChargeSubmitBirthday(ctx).Birthday(birthday).Reference(reference).Execute()

Submit Birthday

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	birthday := "birthday_example" // string | Customer's birthday in the format YYYY-MM-DD e.g 2016-09-21
	reference := "reference_example" // string | The reference of the ongoing transaction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChargeAPI.ChargeSubmitBirthday(context.Background()).Birthday(birthday).Reference(reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChargeAPI.ChargeSubmitBirthday``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChargeSubmitBirthday`: Response
	fmt.Fprintf(os.Stdout, "Response from `ChargeAPI.ChargeSubmitBirthday`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChargeSubmitBirthdayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **birthday** | **string** | Customer&#39;s birthday in the format YYYY-MM-DD e.g 2016-09-21 | 
 **reference** | **string** | The reference of the ongoing transaction | 

### Return type

[**Response**](Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChargeSubmitOtp

> Response ChargeSubmitOtp(ctx).Otp(otp).Reference(reference).Execute()

Submit OTP

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	otp := "otp_example" // string | Customer's OTP
	reference := "reference_example" // string | The reference of the ongoing transaction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChargeAPI.ChargeSubmitOtp(context.Background()).Otp(otp).Reference(reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChargeAPI.ChargeSubmitOtp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChargeSubmitOtp`: Response
	fmt.Fprintf(os.Stdout, "Response from `ChargeAPI.ChargeSubmitOtp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChargeSubmitOtpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **otp** | **string** | Customer&#39;s OTP | 
 **reference** | **string** | The reference of the ongoing transaction | 

### Return type

[**Response**](Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChargeSubmitPhone

> Response ChargeSubmitPhone(ctx).Phone(phone).Reference(reference).Execute()

Submit Phone

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	phone := "phone_example" // string | Customer's mobile number
	reference := "reference_example" // string | The reference of the ongoing transaction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChargeAPI.ChargeSubmitPhone(context.Background()).Phone(phone).Reference(reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChargeAPI.ChargeSubmitPhone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChargeSubmitPhone`: Response
	fmt.Fprintf(os.Stdout, "Response from `ChargeAPI.ChargeSubmitPhone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChargeSubmitPhoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phone** | **string** | Customer&#39;s mobile number | 
 **reference** | **string** | The reference of the ongoing transaction | 

### Return type

[**Response**](Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChargeSubmitPin

> Response ChargeSubmitPin(ctx).Pin(pin).Reference(reference).Execute()

Submit PIN

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	pin := "pin_example" // string | Customer's PIN
	reference := "reference_example" // string | Transaction reference that requires the PIN

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChargeAPI.ChargeSubmitPin(context.Background()).Pin(pin).Reference(reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChargeAPI.ChargeSubmitPin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChargeSubmitPin`: Response
	fmt.Fprintf(os.Stdout, "Response from `ChargeAPI.ChargeSubmitPin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChargeSubmitPinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pin** | **string** | Customer&#39;s PIN | 
 **reference** | **string** | Transaction reference that requires the PIN | 

### Return type

[**Response**](Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

