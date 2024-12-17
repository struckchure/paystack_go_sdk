# \DisputeAPI

All URIs are relative to *https://api.paystack.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisputeDownload**](DisputeAPI.md#DisputeDownload) | **Get** /dispute/export | Export Disputes
[**DisputeEvidence**](DisputeAPI.md#DisputeEvidence) | **Post** /dispute/{id}/evidence | Add Evidence
[**DisputeFetch**](DisputeAPI.md#DisputeFetch) | **Get** /dispute/{id} | Fetch Dispute
[**DisputeList**](DisputeAPI.md#DisputeList) | **Get** /dispute | List Disputes
[**DisputeResolve**](DisputeAPI.md#DisputeResolve) | **Put** /dispute/{id}/resolve | Resolve a Dispute
[**DisputeTransaction**](DisputeAPI.md#DisputeTransaction) | **Get** /dispute/transaction/{id} | List Transaction Disputes
[**DisputeUpdate**](DisputeAPI.md#DisputeUpdate) | **Put** /dispute/{id} | Update Dispute
[**DisputeUploadUrl**](DisputeAPI.md#DisputeUploadUrl) | **Get** /dispute/{id}/upload_url | Get Upload URL



## DisputeDownload

> Response DisputeDownload(ctx).PerPage(perPage).Page(page).Status(status).From(from).To(to).Execute()

Export Disputes

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
	resp, r, err := apiClient.DisputeAPI.DisputeDownload(context.Background()).PerPage(perPage).Page(page).Status(status).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisputeAPI.DisputeDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisputeDownload`: Response
	fmt.Fprintf(os.Stdout, "Response from `DisputeAPI.DisputeDownload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisputeDownloadRequest struct via the builder pattern


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


## DisputeEvidence

> Response DisputeEvidence(ctx, id).CustomerEmail(customerEmail).CustomerName(customerName).CustomerPhone(customerPhone).ServiceDetails(serviceDetails).DeliveryAddress(deliveryAddress).DeliveryDate(deliveryDate).Execute()

Add Evidence

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
	id := "id_example" // string | Dispute ID
	customerEmail := "customerEmail_example" // string | Customer email
	customerName := "customerName_example" // string | Customer name
	customerPhone := "customerPhone_example" // string | Customer mobile number
	serviceDetails := "serviceDetails_example" // string | Details of service offered
	deliveryAddress := "deliveryAddress_example" // string | Delivery address (optional)
	deliveryDate := time.Now() // time.Time | ISO 8601 representation of delivery date (YYYY-MM-DD) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisputeAPI.DisputeEvidence(context.Background(), id).CustomerEmail(customerEmail).CustomerName(customerName).CustomerPhone(customerPhone).ServiceDetails(serviceDetails).DeliveryAddress(deliveryAddress).DeliveryDate(deliveryDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisputeAPI.DisputeEvidence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisputeEvidence`: Response
	fmt.Fprintf(os.Stdout, "Response from `DisputeAPI.DisputeEvidence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Dispute ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisputeEvidenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customerEmail** | **string** | Customer email | 
 **customerName** | **string** | Customer name | 
 **customerPhone** | **string** | Customer mobile number | 
 **serviceDetails** | **string** | Details of service offered | 
 **deliveryAddress** | **string** | Delivery address | 
 **deliveryDate** | **time.Time** | ISO 8601 representation of delivery date (YYYY-MM-DD) | 

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


## DisputeFetch

> Response DisputeFetch(ctx, id).Execute()

Fetch Dispute

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
	id := "id_example" // string | Dispute ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisputeAPI.DisputeFetch(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisputeAPI.DisputeFetch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisputeFetch`: Response
	fmt.Fprintf(os.Stdout, "Response from `DisputeAPI.DisputeFetch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Dispute ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisputeFetchRequest struct via the builder pattern


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


## DisputeList

> Response DisputeList(ctx).PerPage(perPage).Page(page).Status(status).Transaction(transaction).From(from).To(to).Execute()

List Disputes

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
	status := "status_example" // string | Dispute Status. Acceptable values are awaiting-merchant-feedback, awaiting-bank-feedback, pending, resolved (optional)
	transaction := "transaction_example" // string | Transaction ID (optional)
	from := time.Now() // time.Time | The start date (optional)
	to := time.Now() // time.Time | The end date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisputeAPI.DisputeList(context.Background()).PerPage(perPage).Page(page).Status(status).Transaction(transaction).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisputeAPI.DisputeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisputeList`: Response
	fmt.Fprintf(os.Stdout, "Response from `DisputeAPI.DisputeList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisputeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | Number of records to fetch per page | 
 **page** | **int32** | The section to retrieve | 
 **status** | **string** | Dispute Status. Acceptable values are awaiting-merchant-feedback, awaiting-bank-feedback, pending, resolved | 
 **transaction** | **string** | Transaction ID | 
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


## DisputeResolve

> Response DisputeResolve(ctx, id).Resolution(resolution).Message(message).RefundAmount(refundAmount).UploadedFilename(uploadedFilename).Evidence(evidence).Execute()

Resolve a Dispute

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
	id := "id_example" // string | Dispute ID
	resolution := "resolution_example" // string | Dispute resolution. Accepted values, merchant-accepted, declined
	message := "message_example" // string | Reason for resolving
	refundAmount := "refundAmount_example" // string | The amount to refund, in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR
	uploadedFilename := "uploadedFilename_example" // string | Filename of attachment returned via response from the Dispute upload URL
	evidence := int32(56) // int32 | Evidence Id for fraud claims (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisputeAPI.DisputeResolve(context.Background(), id).Resolution(resolution).Message(message).RefundAmount(refundAmount).UploadedFilename(uploadedFilename).Evidence(evidence).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisputeAPI.DisputeResolve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisputeResolve`: Response
	fmt.Fprintf(os.Stdout, "Response from `DisputeAPI.DisputeResolve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Dispute ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisputeResolveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resolution** | **string** | Dispute resolution. Accepted values, merchant-accepted, declined | 
 **message** | **string** | Reason for resolving | 
 **refundAmount** | **string** | The amount to refund, in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR | 
 **uploadedFilename** | **string** | Filename of attachment returned via response from the Dispute upload URL | 
 **evidence** | **int32** | Evidence Id for fraud claims | 

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


## DisputeTransaction

> Response DisputeTransaction(ctx, id).Execute()

List Transaction Disputes

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
	id := "id_example" // string | Transaction ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisputeAPI.DisputeTransaction(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisputeAPI.DisputeTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisputeTransaction`: Response
	fmt.Fprintf(os.Stdout, "Response from `DisputeAPI.DisputeTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisputeTransactionRequest struct via the builder pattern


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


## DisputeUpdate

> Response DisputeUpdate(ctx, id).RefundAmount(refundAmount).UploadedFilename(uploadedFilename).Execute()

Update Dispute

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
	id := "id_example" // string | Dispute ID
	refundAmount := "refundAmount_example" // string | The amount to refund, in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR
	uploadedFilename := "uploadedFilename_example" // string | Filename of attachment returned via response from the Dispute upload URL (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisputeAPI.DisputeUpdate(context.Background(), id).RefundAmount(refundAmount).UploadedFilename(uploadedFilename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisputeAPI.DisputeUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisputeUpdate`: Response
	fmt.Fprintf(os.Stdout, "Response from `DisputeAPI.DisputeUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Dispute ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisputeUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refundAmount** | **string** | The amount to refund, in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR | 
 **uploadedFilename** | **string** | Filename of attachment returned via response from the Dispute upload URL | 

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


## DisputeUploadUrl

> Response DisputeUploadUrl(ctx, id).Execute()

Get Upload URL

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
	id := "id_example" // string | Dispute ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisputeAPI.DisputeUploadUrl(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisputeAPI.DisputeUploadUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisputeUploadUrl`: Response
	fmt.Fprintf(os.Stdout, "Response from `DisputeAPI.DisputeUploadUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Dispute ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisputeUploadUrlRequest struct via the builder pattern


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

