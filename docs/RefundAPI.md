# \RefundAPI

All URIs are relative to *https://api.paystack.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RefundCreate**](RefundAPI.md#RefundCreate) | **Post** /refund | Create Refund
[**RefundFetch**](RefundAPI.md#RefundFetch) | **Get** /refund/{id} | Fetch Refund
[**RefundList**](RefundAPI.md#RefundList) | **Get** /refund | List Refunds



## RefundCreate

> Response RefundCreate(ctx).Transaction(transaction).Amount(amount).Currency(currency).CustomerNote(customerNote).MerchantNote(merchantNote).Execute()

Create Refund

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
	transaction := "transaction_example" // string | Transaction reference or id
	amount := int32(56) // int32 | Amount ( in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR ) to be refunded to the customer.  Amount cannot be more than the original transaction amount (optional)
	currency := "currency_example" // string | Three-letter ISO currency. Allowed values are NGN, GHS, ZAR or USD (optional)
	customerNote := "customerNote_example" // string | Customer reason (optional)
	merchantNote := "merchantNote_example" // string | Merchant reason (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RefundAPI.RefundCreate(context.Background()).Transaction(transaction).Amount(amount).Currency(currency).CustomerNote(customerNote).MerchantNote(merchantNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RefundAPI.RefundCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefundCreate`: Response
	fmt.Fprintf(os.Stdout, "Response from `RefundAPI.RefundCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefundCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transaction** | **string** | Transaction reference or id | 
 **amount** | **int32** | Amount ( in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR ) to be refunded to the customer.  Amount cannot be more than the original transaction amount | 
 **currency** | **string** | Three-letter ISO currency. Allowed values are NGN, GHS, ZAR or USD | 
 **customerNote** | **string** | Customer reason | 
 **merchantNote** | **string** | Merchant reason | 

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


## RefundFetch

> Response RefundFetch(ctx, id).Execute()

Fetch Refund

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
	resp, r, err := apiClient.RefundAPI.RefundFetch(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RefundAPI.RefundFetch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefundFetch`: Response
	fmt.Fprintf(os.Stdout, "Response from `RefundAPI.RefundFetch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefundFetchRequest struct via the builder pattern


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


## RefundList

> Response RefundList(ctx).PerPage(perPage).Page(page).From(from).To(to).Execute()

List Refunds

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
	from := time.Now() // time.Time | The start date (optional)
	to := time.Now() // time.Time | The end date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RefundAPI.RefundList(context.Background()).PerPage(perPage).Page(page).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RefundAPI.RefundList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefundList`: Response
	fmt.Fprintf(os.Stdout, "Response from `RefundAPI.RefundList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefundListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | Number of records to fetch per page | 
 **page** | **int32** | The section to retrieve | 
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

