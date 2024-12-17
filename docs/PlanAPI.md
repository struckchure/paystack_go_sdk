# \PlanAPI

All URIs are relative to *https://api.paystack.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlanCreate**](PlanAPI.md#PlanCreate) | **Post** /plan | Create Plan
[**PlanFetch**](PlanAPI.md#PlanFetch) | **Get** /plan/{code} | Fetch Plan
[**PlanList**](PlanAPI.md#PlanList) | **Get** /plan | List Plans
[**PlanUpdate**](PlanAPI.md#PlanUpdate) | **Put** /plan/{code} | Update Plan



## PlanCreate

> Response PlanCreate(ctx).Name(name).Amount(amount).Interval(interval).Description(description).SendInvoices(sendInvoices).SendSms(sendSms).Currency(currency).InvoiceLimit(invoiceLimit).Execute()

Create Plan

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
	name := "name_example" // string | Name of plan
	amount := int32(56) // int32 | Amount should be in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR
	interval := "interval_example" // string | Interval in words. Valid intervals are daily, weekly, monthly,biannually, annually
	description := "description_example" // string | A description for this plan (optional)
	sendInvoices := true // bool | Set to false if you don't want invoices to be sent to your customers (optional)
	sendSms := true // bool | Set to false if you don't want text messages to be sent to your customers (optional)
	currency := "currency_example" // string | Currency in which amount is set. Allowed values are NGN, GHS, ZAR or USD (optional)
	invoiceLimit := int32(56) // int32 | Number of invoices to raise during subscription to this plan.  Can be overridden by specifying an invoice_limit while subscribing. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlanAPI.PlanCreate(context.Background()).Name(name).Amount(amount).Interval(interval).Description(description).SendInvoices(sendInvoices).SendSms(sendSms).Currency(currency).InvoiceLimit(invoiceLimit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.PlanCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlanCreate`: Response
	fmt.Fprintf(os.Stdout, "Response from `PlanAPI.PlanCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlanCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Name of plan | 
 **amount** | **int32** | Amount should be in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR | 
 **interval** | **string** | Interval in words. Valid intervals are daily, weekly, monthly,biannually, annually | 
 **description** | **string** | A description for this plan | 
 **sendInvoices** | **bool** | Set to false if you don&#39;t want invoices to be sent to your customers | 
 **sendSms** | **bool** | Set to false if you don&#39;t want text messages to be sent to your customers | 
 **currency** | **string** | Currency in which amount is set. Allowed values are NGN, GHS, ZAR or USD | 
 **invoiceLimit** | **int32** | Number of invoices to raise during subscription to this plan.  Can be overridden by specifying an invoice_limit while subscribing. | 

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


## PlanFetch

> Response PlanFetch(ctx, code).Execute()

Fetch Plan

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
	code := "code_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlanAPI.PlanFetch(context.Background(), code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.PlanFetch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlanFetch`: Response
	fmt.Fprintf(os.Stdout, "Response from `PlanAPI.PlanFetch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlanFetchRequest struct via the builder pattern


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


## PlanList

> Response PlanList(ctx).PerPage(perPage).Page(page).Interval(interval).Amount(amount).From(from).To(to).Execute()

List Plans

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
	interval := "interval_example" // string | Specify interval of the plan (optional)
	amount := int32(56) // int32 | The amount on the plans to retrieve (optional)
	from := time.Now() // time.Time | The start date (optional)
	to := time.Now() // time.Time | The end date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlanAPI.PlanList(context.Background()).PerPage(perPage).Page(page).Interval(interval).Amount(amount).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.PlanList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlanList`: Response
	fmt.Fprintf(os.Stdout, "Response from `PlanAPI.PlanList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlanListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | Number of records to fetch per page | 
 **page** | **int32** | The section to retrieve | 
 **interval** | **string** | Specify interval of the plan | 
 **amount** | **int32** | The amount on the plans to retrieve | 
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


## PlanUpdate

> Response PlanUpdate(ctx, code).Name(name).Amount(amount).Interval(interval).Description(description).SendInvoices(sendInvoices).SendSms(sendSms).Currency(currency).InvoiceLimit(invoiceLimit).Execute()

Update Plan

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
	code := "code_example" // string | 
	name := "name_example" // string | Name of plan (optional)
	amount := int32(56) // int32 | Amount should be in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR (optional)
	interval := "interval_example" // string | Interval in words. Valid intervals are daily, weekly, monthly,biannually, annually (optional)
	description := true // bool | A description for this plan (optional)
	sendInvoices := true // bool | Set to false if you don't want invoices to be sent to your customers (optional)
	sendSms := true // bool | Set to false if you don't want text messages to be sent to your customers (optional)
	currency := "currency_example" // string | Currency in which amount is set. Allowed values are NGN, GHS, ZAR or USD (optional)
	invoiceLimit := int32(56) // int32 | Number of invoices to raise during subscription to this plan.  Can be overridden by specifying an invoice_limit while subscribing. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlanAPI.PlanUpdate(context.Background(), code).Name(name).Amount(amount).Interval(interval).Description(description).SendInvoices(sendInvoices).SendSms(sendSms).Currency(currency).InvoiceLimit(invoiceLimit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.PlanUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlanUpdate`: Response
	fmt.Fprintf(os.Stdout, "Response from `PlanAPI.PlanUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlanUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Name of plan | 
 **amount** | **int32** | Amount should be in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR | 
 **interval** | **string** | Interval in words. Valid intervals are daily, weekly, monthly,biannually, annually | 
 **description** | **bool** | A description for this plan | 
 **sendInvoices** | **bool** | Set to false if you don&#39;t want invoices to be sent to your customers | 
 **sendSms** | **bool** | Set to false if you don&#39;t want text messages to be sent to your customers | 
 **currency** | **string** | Currency in which amount is set. Allowed values are NGN, GHS, ZAR or USD | 
 **invoiceLimit** | **int32** | Number of invoices to raise during subscription to this plan.  Can be overridden by specifying an invoice_limit while subscribing. | 

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

