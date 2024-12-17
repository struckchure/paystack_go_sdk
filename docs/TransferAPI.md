# \TransferAPI

All URIs are relative to *https://api.paystack.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransferBulk**](TransferAPI.md#TransferBulk) | **Post** /transfer/bulk | Initiate Bulk Transfer
[**TransferDisableOtp**](TransferAPI.md#TransferDisableOtp) | **Post** /transfer/disable_otp | Disable OTP requirement for Transfers
[**TransferDisableOtpFinalize**](TransferAPI.md#TransferDisableOtpFinalize) | **Post** /transfer/disable_otp_finalize | Finalize Disabling of OTP requirement for Transfers
[**TransferDownload**](TransferAPI.md#TransferDownload) | **Get** /transfer/export | Export Transfers
[**TransferEnableOtp**](TransferAPI.md#TransferEnableOtp) | **Post** /transfer/enable_otp | Enable OTP requirement for Transfers
[**TransferFetch**](TransferAPI.md#TransferFetch) | **Get** /transfer/{code} | Fetch Transfer
[**TransferFinalize**](TransferAPI.md#TransferFinalize) | **Post** /transfer/finalize_transfer | Finalize Transfer
[**TransferInitiate**](TransferAPI.md#TransferInitiate) | **Post** /transfer | Initiate Transfer
[**TransferList**](TransferAPI.md#TransferList) | **Get** /transfer | List Transfers
[**TransferResendOtp**](TransferAPI.md#TransferResendOtp) | **Post** /transfer/resend_otp | Resend OTP for Transfer
[**TransferVerify**](TransferAPI.md#TransferVerify) | **Get** /transfer/verify/{reference} | Verify Transfer



## TransferBulk

> Response TransferBulk(ctx).Source(source).Transfers(transfers).Execute()

Initiate Bulk Transfer

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
	source := "source_example" // string | Where should we transfer from? Only balance is allowed for now (optional)
	transfers := []openapiclient.TransferInitiate{*openapiclient.NewTransferInitiate("Source_example", "Amount_example", "Recipient_example")} // []TransferInitiate | A list of transfer object. Each object should contain amount, recipient, and reference (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransferAPI.TransferBulk(context.Background()).Source(source).Transfers(transfers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransferAPI.TransferBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferBulk`: Response
	fmt.Fprintf(os.Stdout, "Response from `TransferAPI.TransferBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **source** | **string** | Where should we transfer from? Only balance is allowed for now | 
 **transfers** | [**[]TransferInitiate**](TransferInitiate.md) | A list of transfer object. Each object should contain amount, recipient, and reference | 

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


## TransferDisableOtp

> Response TransferDisableOtp(ctx).Execute()

Disable OTP requirement for Transfers

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
	resp, r, err := apiClient.TransferAPI.TransferDisableOtp(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransferAPI.TransferDisableOtp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferDisableOtp`: Response
	fmt.Fprintf(os.Stdout, "Response from `TransferAPI.TransferDisableOtp`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTransferDisableOtpRequest struct via the builder pattern


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


## TransferDisableOtpFinalize

> Response TransferDisableOtpFinalize(ctx).Otp(otp).Execute()

Finalize Disabling of OTP requirement for Transfers

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
	otp := "otp_example" // string | OTP sent to business phone to verify disabling OTP requirement

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransferAPI.TransferDisableOtpFinalize(context.Background()).Otp(otp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransferAPI.TransferDisableOtpFinalize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferDisableOtpFinalize`: Response
	fmt.Fprintf(os.Stdout, "Response from `TransferAPI.TransferDisableOtpFinalize`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferDisableOtpFinalizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **otp** | **string** | OTP sent to business phone to verify disabling OTP requirement | 

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


## TransferDownload

> Response TransferDownload(ctx).PerPage(perPage).Page(page).Status(status).From(from).To(to).Execute()

Export Transfers

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
	perPage := int32(56) // int32 | Number of records to fetch per page (optional)
	page := int32(56) // int32 | The section to retrieve (optional)
	status := "status_example" // string |  (optional)
	from := time.Now() // time.Time | The start date (optional)
	to := time.Now() // time.Time | The end date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransferAPI.TransferDownload(context.Background()).PerPage(perPage).Page(page).Status(status).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransferAPI.TransferDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferDownload`: Response
	fmt.Fprintf(os.Stdout, "Response from `TransferAPI.TransferDownload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | Number of records to fetch per page | 
 **page** | **int32** | The section to retrieve | 
 **status** | **string** |  | 
 **from** | **time.Time** | The start date | 
 **to** | **time.Time** | The end date | 

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


## TransferEnableOtp

> Response TransferEnableOtp(ctx).Execute()

Enable OTP requirement for Transfers

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
	resp, r, err := apiClient.TransferAPI.TransferEnableOtp(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransferAPI.TransferEnableOtp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferEnableOtp`: Response
	fmt.Fprintf(os.Stdout, "Response from `TransferAPI.TransferEnableOtp`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTransferEnableOtpRequest struct via the builder pattern


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


## TransferFetch

> Response TransferFetch(ctx, code).Execute()

Fetch Transfer

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
	code := "code_example" // string | Transfer code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransferAPI.TransferFetch(context.Background(), code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransferAPI.TransferFetch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferFetch`: Response
	fmt.Fprintf(os.Stdout, "Response from `TransferAPI.TransferFetch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Transfer code | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferFetchRequest struct via the builder pattern


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


## TransferFinalize

> Response TransferFinalize(ctx).TransferCode(transferCode).Otp(otp).Execute()

Finalize Transfer

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
	transferCode := "transferCode_example" // string | The transfer code you want to finalize
	otp := "otp_example" // string | OTP sent to business phone to verify transfer

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransferAPI.TransferFinalize(context.Background()).TransferCode(transferCode).Otp(otp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransferAPI.TransferFinalize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferFinalize`: Response
	fmt.Fprintf(os.Stdout, "Response from `TransferAPI.TransferFinalize`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferFinalizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferCode** | **string** | The transfer code you want to finalize | 
 **otp** | **string** | OTP sent to business phone to verify transfer | 

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


## TransferInitiate

> Response TransferInitiate(ctx).Source(source).Amount(amount).Recipient(recipient).Reason(reason).Currency(currency).Reference(reference).Execute()

Initiate Transfer

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
	source := "source_example" // string | Where should we transfer from? Only balance is allowed for now
	amount := "amount_example" // string | Amount to transfer in kobo if currency is NGN and pesewas if currency is GHS.
	recipient := "recipient_example" // string | The transfer recipient's code
	reason := "reason_example" // string | The reason or narration for the transfer. (optional)
	currency := "currency_example" // string | Specify the currency of the transfer. Defaults to NGN. (optional)
	reference := "reference_example" // string | If specified, the field should be a unique identifier (in lowercase) for the object.  Only -,_ and alphanumeric characters are allowed. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransferAPI.TransferInitiate(context.Background()).Source(source).Amount(amount).Recipient(recipient).Reason(reason).Currency(currency).Reference(reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransferAPI.TransferInitiate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferInitiate`: Response
	fmt.Fprintf(os.Stdout, "Response from `TransferAPI.TransferInitiate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferInitiateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **source** | **string** | Where should we transfer from? Only balance is allowed for now | 
 **amount** | **string** | Amount to transfer in kobo if currency is NGN and pesewas if currency is GHS. | 
 **recipient** | **string** | The transfer recipient&#39;s code | 
 **reason** | **string** | The reason or narration for the transfer. | 
 **currency** | **string** | Specify the currency of the transfer. Defaults to NGN. | 
 **reference** | **string** | If specified, the field should be a unique identifier (in lowercase) for the object.  Only -,_ and alphanumeric characters are allowed. | 

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


## TransferList

> Response TransferList(ctx).PerPage(perPage).Page(page).Status(status).From(from).To(to).Execute()

List Transfers

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
	perPage := int32(56) // int32 | Number of records to fetch per page (optional)
	page := int32(56) // int32 | The section to retrieve (optional)
	status := "status_example" // string |  (optional)
	from := time.Now() // time.Time | The start date (optional)
	to := time.Now() // time.Time | The end date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransferAPI.TransferList(context.Background()).PerPage(perPage).Page(page).Status(status).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransferAPI.TransferList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferList`: Response
	fmt.Fprintf(os.Stdout, "Response from `TransferAPI.TransferList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | Number of records to fetch per page | 
 **page** | **int32** | The section to retrieve | 
 **status** | **string** |  | 
 **from** | **time.Time** | The start date | 
 **to** | **time.Time** | The end date | 

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


## TransferResendOtp

> Response TransferResendOtp(ctx).TransferCode(transferCode).Reason(reason).Execute()

Resend OTP for Transfer

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
	transferCode := "transferCode_example" // string | The transfer code that requires an OTP validation
	reason := "reason_example" // string | Either resend_otp or transfer

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransferAPI.TransferResendOtp(context.Background()).TransferCode(transferCode).Reason(reason).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransferAPI.TransferResendOtp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferResendOtp`: Response
	fmt.Fprintf(os.Stdout, "Response from `TransferAPI.TransferResendOtp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferResendOtpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferCode** | **string** | The transfer code that requires an OTP validation | 
 **reason** | **string** | Either resend_otp or transfer | 

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


## TransferVerify

> Response TransferVerify(ctx, reference).Execute()

Verify Transfer

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
	resp, r, err := apiClient.TransferAPI.TransferVerify(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransferAPI.TransferVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferVerify`: Response
	fmt.Fprintf(os.Stdout, "Response from `TransferAPI.TransferVerify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferVerifyRequest struct via the builder pattern


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

