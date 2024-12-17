# \PaymentRequestAPI

All URIs are relative to *https://api.paystack.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaymentRequestArchive**](PaymentRequestAPI.md#PaymentRequestArchive) | **Post** /paymentrequest/archive/{id} | Archive Payment Request
[**PaymentRequestCreate**](PaymentRequestAPI.md#PaymentRequestCreate) | **Post** /paymentrequest | Create Payment Request
[**PaymentRequestFetch**](PaymentRequestAPI.md#PaymentRequestFetch) | **Get** /paymentrequest/{id} | Fetch Payment Request
[**PaymentRequestFinalize**](PaymentRequestAPI.md#PaymentRequestFinalize) | **Post** /paymentrequest/finalize/{id} | Finalize Payment Request
[**PaymentRequestList**](PaymentRequestAPI.md#PaymentRequestList) | **Get** /paymentrequest | List Payment Request
[**PaymentRequestNotify**](PaymentRequestAPI.md#PaymentRequestNotify) | **Post** /paymentrequest/notify/{id} | Send Notification
[**PaymentRequestTotals**](PaymentRequestAPI.md#PaymentRequestTotals) | **Get** /paymentrequest/totals | Payment Request Total
[**PaymentRequestUpdate**](PaymentRequestAPI.md#PaymentRequestUpdate) | **Put** /paymentrequest/{id} | Update Payment Request
[**PaymentRequestVerify**](PaymentRequestAPI.md#PaymentRequestVerify) | **Get** /paymentrequest/verify/{id} | Verify Payment Request



## PaymentRequestArchive

> Response PaymentRequestArchive(ctx, id).Execute()

Archive Payment Request

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentRequestAPI.PaymentRequestArchive(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentRequestAPI.PaymentRequestArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentRequestArchive`: Response
	fmt.Fprintf(os.Stdout, "Response from `PaymentRequestAPI.PaymentRequestArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentRequestArchiveRequest struct via the builder pattern


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


## PaymentRequestCreate

> Response PaymentRequestCreate(ctx).Customer(customer).Amount(amount).Currency(currency).DueDate(dueDate).Description(description).LineItems(lineItems).Tax(tax).SendNotification(sendNotification).Draft(draft).HasInvoice(hasInvoice).InvoiceNumber(invoiceNumber).SplitCode(splitCode).Execute()

Create Payment Request

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
	customer := "customer_example" // string | Customer id or code
	amount := int32(56) // int32 | Payment request amount. Only useful if line items and tax values are ignored.  The endpoint will throw a friendly warning if neither is available. (optional)
	currency := "currency_example" // string | Specify the currency of the invoice. Allowed values are NGN, GHS, ZAR and USD. Defaults to NGN (optional)
	dueDate := time.Now() // time.Time | ISO 8601 representation of request due date (optional)
	description := "description_example" // string | A short description of the payment request (optional)
	lineItems := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | Array of line items (optional)
	tax := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | Array of taxes (optional)
	sendNotification := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | Indicates whether Paystack sends an email notification to customer. Defaults to true (optional)
	draft := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | Indicate if request should be saved as draft. Defaults to false and overrides send_notification (optional)
	hasInvoice := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | Set to true to create a draft invoice (adds an auto incrementing invoice number if none is provided)  even if there are no line_items or tax passed (optional)
	invoiceNumber := int32(56) // int32 | Numeric value of invoice. Invoice will start from 1 and auto increment from there. This field is to help  override whatever value Paystack decides. Auto increment for subsequent invoices continue from this point. (optional)
	splitCode := "splitCode_example" // string | The split code of the transaction split. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentRequestAPI.PaymentRequestCreate(context.Background()).Customer(customer).Amount(amount).Currency(currency).DueDate(dueDate).Description(description).LineItems(lineItems).Tax(tax).SendNotification(sendNotification).Draft(draft).HasInvoice(hasInvoice).InvoiceNumber(invoiceNumber).SplitCode(splitCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentRequestAPI.PaymentRequestCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentRequestCreate`: Response
	fmt.Fprintf(os.Stdout, "Response from `PaymentRequestAPI.PaymentRequestCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentRequestCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customer** | **string** | Customer id or code | 
 **amount** | **int32** | Payment request amount. Only useful if line items and tax values are ignored.  The endpoint will throw a friendly warning if neither is available. | 
 **currency** | **string** | Specify the currency of the invoice. Allowed values are NGN, GHS, ZAR and USD. Defaults to NGN | 
 **dueDate** | **time.Time** | ISO 8601 representation of request due date | 
 **description** | **string** | A short description of the payment request | 
 **lineItems** | **[]map[string]interface{}** | Array of line items | 
 **tax** | **[]map[string]interface{}** | Array of taxes | 
 **sendNotification** | **[]map[string]interface{}** | Indicates whether Paystack sends an email notification to customer. Defaults to true | 
 **draft** | **[]map[string]interface{}** | Indicate if request should be saved as draft. Defaults to false and overrides send_notification | 
 **hasInvoice** | **[]map[string]interface{}** | Set to true to create a draft invoice (adds an auto incrementing invoice number if none is provided)  even if there are no line_items or tax passed | 
 **invoiceNumber** | **int32** | Numeric value of invoice. Invoice will start from 1 and auto increment from there. This field is to help  override whatever value Paystack decides. Auto increment for subsequent invoices continue from this point. | 
 **splitCode** | **string** | The split code of the transaction split. | 

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


## PaymentRequestFetch

> Response PaymentRequestFetch(ctx, id).Execute()

Fetch Payment Request

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentRequestAPI.PaymentRequestFetch(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentRequestAPI.PaymentRequestFetch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentRequestFetch`: Response
	fmt.Fprintf(os.Stdout, "Response from `PaymentRequestAPI.PaymentRequestFetch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentRequestFetchRequest struct via the builder pattern


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


## PaymentRequestFinalize

> Response PaymentRequestFinalize(ctx, id).Execute()

Finalize Payment Request

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentRequestAPI.PaymentRequestFinalize(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentRequestAPI.PaymentRequestFinalize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentRequestFinalize`: Response
	fmt.Fprintf(os.Stdout, "Response from `PaymentRequestAPI.PaymentRequestFinalize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentRequestFinalizeRequest struct via the builder pattern


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


## PaymentRequestList

> Response PaymentRequestList(ctx).PerPage(perPage).Page(page).Customer(customer).Status(status).Currency(currency).From(from).To(to).Execute()

List Payment Request

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
	customer := "customer_example" // string | Customer ID (optional)
	status := "status_example" // string | Invoice status to filter (optional)
	currency := "currency_example" // string | If your integration supports more than one currency, choose the one to filter (optional)
	from := time.Now() // time.Time | The start date (optional)
	to := time.Now() // time.Time | The end date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentRequestAPI.PaymentRequestList(context.Background()).PerPage(perPage).Page(page).Customer(customer).Status(status).Currency(currency).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentRequestAPI.PaymentRequestList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentRequestList`: Response
	fmt.Fprintf(os.Stdout, "Response from `PaymentRequestAPI.PaymentRequestList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentRequestListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | Number of records to fetch per page | 
 **page** | **int32** | The section to retrieve | 
 **customer** | **string** | Customer ID | 
 **status** | **string** | Invoice status to filter | 
 **currency** | **string** | If your integration supports more than one currency, choose the one to filter | 
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


## PaymentRequestNotify

> Response PaymentRequestNotify(ctx, id).Execute()

Send Notification

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentRequestAPI.PaymentRequestNotify(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentRequestAPI.PaymentRequestNotify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentRequestNotify`: Response
	fmt.Fprintf(os.Stdout, "Response from `PaymentRequestAPI.PaymentRequestNotify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentRequestNotifyRequest struct via the builder pattern


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


## PaymentRequestTotals

> Response PaymentRequestTotals(ctx).Execute()

Payment Request Total

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
	resp, r, err := apiClient.PaymentRequestAPI.PaymentRequestTotals(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentRequestAPI.PaymentRequestTotals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentRequestTotals`: Response
	fmt.Fprintf(os.Stdout, "Response from `PaymentRequestAPI.PaymentRequestTotals`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentRequestTotalsRequest struct via the builder pattern


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


## PaymentRequestUpdate

> Response PaymentRequestUpdate(ctx, id).Customer(customer).Amount(amount).Currency(currency).DueDate(dueDate).Description(description).LineItems(lineItems).Tax(tax).SendNotification(sendNotification).Draft(draft).HasInvoice(hasInvoice).InvoiceNumber(invoiceNumber).SplitCode(splitCode).Execute()

Update Payment Request

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
	id := "id_example" // string | 
	customer := "customer_example" // string | Customer id or code (optional)
	amount := int32(56) // int32 | Payment request amount. Only useful if line items and tax values are ignored.  The endpoint will throw a friendly warning if neither is available. (optional)
	currency := "currency_example" // string | Specify the currency of the invoice. Allowed values are NGN, GHS, ZAR and USD. Defaults to NGN (optional)
	dueDate := time.Now() // time.Time | ISO 8601 representation of request due date (optional)
	description := "description_example" // string | A short description of the payment request (optional)
	lineItems := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | Array of line items (optional)
	tax := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | Array of taxes (optional)
	sendNotification := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | Indicates whether Paystack sends an email notification to customer. Defaults to true (optional)
	draft := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | Indicate if request should be saved as draft. Defaults to false and overrides send_notification (optional)
	hasInvoice := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | Set to true to create a draft invoice (adds an auto incrementing invoice number if none is provided)  even if there are no line_items or tax passed (optional)
	invoiceNumber := int32(56) // int32 | Numeric value of invoice. Invoice will start from 1 and auto increment from there. This field is to help  override whatever value Paystack decides. Auto increment for subsequent invoices continue from this point. (optional)
	splitCode := "splitCode_example" // string | The split code of the transaction split. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentRequestAPI.PaymentRequestUpdate(context.Background(), id).Customer(customer).Amount(amount).Currency(currency).DueDate(dueDate).Description(description).LineItems(lineItems).Tax(tax).SendNotification(sendNotification).Draft(draft).HasInvoice(hasInvoice).InvoiceNumber(invoiceNumber).SplitCode(splitCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentRequestAPI.PaymentRequestUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentRequestUpdate`: Response
	fmt.Fprintf(os.Stdout, "Response from `PaymentRequestAPI.PaymentRequestUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentRequestUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customer** | **string** | Customer id or code | 
 **amount** | **int32** | Payment request amount. Only useful if line items and tax values are ignored.  The endpoint will throw a friendly warning if neither is available. | 
 **currency** | **string** | Specify the currency of the invoice. Allowed values are NGN, GHS, ZAR and USD. Defaults to NGN | 
 **dueDate** | **time.Time** | ISO 8601 representation of request due date | 
 **description** | **string** | A short description of the payment request | 
 **lineItems** | **[]map[string]interface{}** | Array of line items | 
 **tax** | **[]map[string]interface{}** | Array of taxes | 
 **sendNotification** | **[]map[string]interface{}** | Indicates whether Paystack sends an email notification to customer. Defaults to true | 
 **draft** | **[]map[string]interface{}** | Indicate if request should be saved as draft. Defaults to false and overrides send_notification | 
 **hasInvoice** | **[]map[string]interface{}** | Set to true to create a draft invoice (adds an auto incrementing invoice number if none is provided)  even if there are no line_items or tax passed | 
 **invoiceNumber** | **int32** | Numeric value of invoice. Invoice will start from 1 and auto increment from there. This field is to help  override whatever value Paystack decides. Auto increment for subsequent invoices continue from this point. | 
 **splitCode** | **string** | The split code of the transaction split. | 

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


## PaymentRequestVerify

> Response PaymentRequestVerify(ctx, id).Execute()

Verify Payment Request

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentRequestAPI.PaymentRequestVerify(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentRequestAPI.PaymentRequestVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentRequestVerify`: Response
	fmt.Fprintf(os.Stdout, "Response from `PaymentRequestAPI.PaymentRequestVerify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentRequestVerifyRequest struct via the builder pattern


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

