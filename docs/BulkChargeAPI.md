# \BulkChargeAPI

All URIs are relative to *https://api.paystack.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkChargeCharges**](BulkChargeAPI.md#BulkChargeCharges) | **Get** /bulkcharge/{code}/charges | Fetch Charges in a Batch
[**BulkChargeFetch**](BulkChargeAPI.md#BulkChargeFetch) | **Get** /bulkcharge/{code} | Fetch Bulk Charge Batch
[**BulkChargeInitiate**](BulkChargeAPI.md#BulkChargeInitiate) | **Post** /bulkcharge | Initiate Bulk Charge
[**BulkChargeList**](BulkChargeAPI.md#BulkChargeList) | **Get** /bulkcharge | List Bulk Charge Batches
[**BulkChargePause**](BulkChargeAPI.md#BulkChargePause) | **Get** /bulkcharge/pause/{code} | Pause Bulk Charge Batch
[**BulkChargeResume**](BulkChargeAPI.md#BulkChargeResume) | **Get** /bulkcharge/resume/{code} | Resume Bulk Charge Batch



## BulkChargeCharges

> Response BulkChargeCharges(ctx, code).Execute()

Fetch Charges in a Batch

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
	code := "code_example" // string | Batch code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkChargeAPI.BulkChargeCharges(context.Background(), code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkChargeAPI.BulkChargeCharges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkChargeCharges`: Response
	fmt.Fprintf(os.Stdout, "Response from `BulkChargeAPI.BulkChargeCharges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Batch code | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkChargeChargesRequest struct via the builder pattern


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


## BulkChargeFetch

> Response BulkChargeFetch(ctx, code).Execute()

Fetch Bulk Charge Batch

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
	code := "code_example" // string | Batch code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkChargeAPI.BulkChargeFetch(context.Background(), code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkChargeAPI.BulkChargeFetch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkChargeFetch`: Response
	fmt.Fprintf(os.Stdout, "Response from `BulkChargeAPI.BulkChargeFetch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Batch code | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkChargeFetchRequest struct via the builder pattern


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


## BulkChargeInitiate

> Response BulkChargeInitiate(ctx).Execute()

Initiate Bulk Charge

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
	resp, r, err := apiClient.BulkChargeAPI.BulkChargeInitiate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkChargeAPI.BulkChargeInitiate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkChargeInitiate`: Response
	fmt.Fprintf(os.Stdout, "Response from `BulkChargeAPI.BulkChargeInitiate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBulkChargeInitiateRequest struct via the builder pattern


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


## BulkChargeList

> Response BulkChargeList(ctx).PerPage(perPage).Page(page).From(from).To(to).Execute()

List Bulk Charge Batches

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
	resp, r, err := apiClient.BulkChargeAPI.BulkChargeList(context.Background()).PerPage(perPage).Page(page).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkChargeAPI.BulkChargeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkChargeList`: Response
	fmt.Fprintf(os.Stdout, "Response from `BulkChargeAPI.BulkChargeList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkChargeListRequest struct via the builder pattern


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


## BulkChargePause

> Response BulkChargePause(ctx, code).Execute()

Pause Bulk Charge Batch

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
	code := "code_example" // string | Batch code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkChargeAPI.BulkChargePause(context.Background(), code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkChargeAPI.BulkChargePause``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkChargePause`: Response
	fmt.Fprintf(os.Stdout, "Response from `BulkChargeAPI.BulkChargePause`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Batch code | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkChargePauseRequest struct via the builder pattern


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


## BulkChargeResume

> Response BulkChargeResume(ctx, code).Execute()

Resume Bulk Charge Batch

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
	code := "code_example" // string | Batch code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkChargeAPI.BulkChargeResume(context.Background(), code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkChargeAPI.BulkChargeResume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkChargeResume`: Response
	fmt.Fprintf(os.Stdout, "Response from `BulkChargeAPI.BulkChargeResume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Batch code | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkChargeResumeRequest struct via the builder pattern


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

