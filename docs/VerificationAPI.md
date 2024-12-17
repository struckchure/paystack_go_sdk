# \VerificationAPI

All URIs are relative to *https://api.paystack.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VerificationAvs**](VerificationAPI.md#VerificationAvs) | **Get** /address_verification/states | List States (AVS)
[**VerificationFetchBanks**](VerificationAPI.md#VerificationFetchBanks) | **Get** /bank | Fetch Banks
[**VerificationListCountries**](VerificationAPI.md#VerificationListCountries) | **Get** /country | List Countries
[**VerificationResolveAccountNumber**](VerificationAPI.md#VerificationResolveAccountNumber) | **Get** /bank/resolve | Resolve Account Number
[**VerificationResolveCardBin**](VerificationAPI.md#VerificationResolveCardBin) | **Get** /decision/bin/{bin} | Resolve Card BIN



## VerificationAvs

> Response VerificationAvs(ctx).Type_(type_).Country(country).Currency(currency).Execute()

List States (AVS)

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
	type_ := "type__example" // string |  (optional)
	country := "country_example" // string |  (optional)
	currency := "currency_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerificationAPI.VerificationAvs(context.Background()).Type_(type_).Country(country).Currency(currency).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationAPI.VerificationAvs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerificationAvs`: Response
	fmt.Fprintf(os.Stdout, "Response from `VerificationAPI.VerificationAvs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerificationAvsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** |  | 
 **country** | **string** |  | 
 **currency** | **string** |  | 

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


## VerificationFetchBanks

> Response VerificationFetchBanks(ctx).Country(country).PayWithBankTransfer(payWithBankTransfer).UseCursor(useCursor).PerPage(perPage).Next(next).Previous(previous).Gateway(gateway).Execute()

Fetch Banks

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
	country := "country_example" // string |  (optional)
	payWithBankTransfer := true // bool |  (optional)
	useCursor := true // bool |  (optional)
	perPage := int32(56) // int32 |  (optional)
	next := "next_example" // string |  (optional)
	previous := "previous_example" // string |  (optional)
	gateway := "gateway_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerificationAPI.VerificationFetchBanks(context.Background()).Country(country).PayWithBankTransfer(payWithBankTransfer).UseCursor(useCursor).PerPage(perPage).Next(next).Previous(previous).Gateway(gateway).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationAPI.VerificationFetchBanks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerificationFetchBanks`: Response
	fmt.Fprintf(os.Stdout, "Response from `VerificationAPI.VerificationFetchBanks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerificationFetchBanksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** |  | 
 **payWithBankTransfer** | **bool** |  | 
 **useCursor** | **bool** |  | 
 **perPage** | **int32** |  | 
 **next** | **string** |  | 
 **previous** | **string** |  | 
 **gateway** | **string** |  | 

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


## VerificationListCountries

> Response VerificationListCountries(ctx).Execute()

List Countries

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerificationAPI.VerificationListCountries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationAPI.VerificationListCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerificationListCountries`: Response
	fmt.Fprintf(os.Stdout, "Response from `VerificationAPI.VerificationListCountries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVerificationListCountriesRequest struct via the builder pattern


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


## VerificationResolveAccountNumber

> Response VerificationResolveAccountNumber(ctx).AccountNumber(accountNumber).BankCode(bankCode).Execute()

Resolve Account Number

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
	accountNumber := int32(0022728151) // int32 |  (optional)
	bankCode := int32(51) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerificationAPI.VerificationResolveAccountNumber(context.Background()).AccountNumber(accountNumber).BankCode(bankCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationAPI.VerificationResolveAccountNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerificationResolveAccountNumber`: Response
	fmt.Fprintf(os.Stdout, "Response from `VerificationAPI.VerificationResolveAccountNumber`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerificationResolveAccountNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountNumber** | **int32** |  | 
 **bankCode** | **int32** |  | 

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


## VerificationResolveCardBin

> Response VerificationResolveCardBin(ctx, bin).Execute()

Resolve Card BIN

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
	bin := "bin_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerificationAPI.VerificationResolveCardBin(context.Background(), bin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationAPI.VerificationResolveCardBin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerificationResolveCardBin`: Response
	fmt.Fprintf(os.Stdout, "Response from `VerificationAPI.VerificationResolveCardBin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bin** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerificationResolveCardBinRequest struct via the builder pattern


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

