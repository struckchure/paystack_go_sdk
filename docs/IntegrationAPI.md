# \IntegrationAPI

All URIs are relative to *https://api.paystack.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IntegrationFetchPaymentSessionTimeout**](IntegrationAPI.md#IntegrationFetchPaymentSessionTimeout) | **Get** /integration/payment_session_timeout | Fetch Payment Session Timeout
[**IntegrationUpdatePaymentSessionTimeout**](IntegrationAPI.md#IntegrationUpdatePaymentSessionTimeout) | **Put** /integration/payment_session_timeout | Update Payment Session Timeout



## IntegrationFetchPaymentSessionTimeout

> Response IntegrationFetchPaymentSessionTimeout(ctx).Execute()

Fetch Payment Session Timeout

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
	resp, r, err := apiClient.IntegrationAPI.IntegrationFetchPaymentSessionTimeout(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.IntegrationFetchPaymentSessionTimeout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationFetchPaymentSessionTimeout`: Response
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.IntegrationFetchPaymentSessionTimeout`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationFetchPaymentSessionTimeoutRequest struct via the builder pattern


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


## IntegrationUpdatePaymentSessionTimeout

> Response IntegrationUpdatePaymentSessionTimeout(ctx).Body(body).Execute()

Update Payment Session Timeout

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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.IntegrationUpdatePaymentSessionTimeout(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.IntegrationUpdatePaymentSessionTimeout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationUpdatePaymentSessionTimeout`: Response
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.IntegrationUpdatePaymentSessionTimeout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationUpdatePaymentSessionTimeoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**Response**](Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

