# \PageAPI

All URIs are relative to *https://api.paystack.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PageAddProducts**](PageAPI.md#PageAddProducts) | **Post** /page/{id}/product | Add Products
[**PageCheckSlugAvailability**](PageAPI.md#PageCheckSlugAvailability) | **Get** /page/check_slug_availability/{slug} | Check Slug Availability
[**PageCreate**](PageAPI.md#PageCreate) | **Post** /page | Create Page
[**PageFetch**](PageAPI.md#PageFetch) | **Get** /page/{id} | Fetch Page
[**PageList**](PageAPI.md#PageList) | **Get** /page | List Pages
[**PageUpdate**](PageAPI.md#PageUpdate) | **Put** /page/{id} | Update Page



## PageAddProducts

> Response PageAddProducts(ctx, id).Product(product).Execute()

Add Products

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
	product := []string{"Inner_example"} // []string | IDs of all products to add to a page

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PageAPI.PageAddProducts(context.Background(), id).Product(product).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PageAPI.PageAddProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageAddProducts`: Response
	fmt.Fprintf(os.Stdout, "Response from `PageAPI.PageAddProducts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageAddProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **product** | **[]string** | IDs of all products to add to a page | 

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


## PageCheckSlugAvailability

> Response PageCheckSlugAvailability(ctx, slug).Execute()

Check Slug Availability

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
	slug := "slug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PageAPI.PageCheckSlugAvailability(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PageAPI.PageCheckSlugAvailability``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageCheckSlugAvailability`: Response
	fmt.Fprintf(os.Stdout, "Response from `PageAPI.PageCheckSlugAvailability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageCheckSlugAvailabilityRequest struct via the builder pattern


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


## PageCreate

> Response PageCreate(ctx).Name(name).Description(description).Amount(amount).Slug(slug).Metadata(metadata).RedirectUrl(redirectUrl).CustomFields(customFields).Execute()

Create Page

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
	name := "name_example" // string | Name of page
	description := "description_example" // string | The description of the page (optional)
	amount := int32(56) // int32 | Amount should be in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR (optional)
	slug := "slug_example" // string | URL slug you would like to be associated with this page. Page will be accessible at https://paystack.com/pay/[slug] (optional)
	metadata := "metadata_example" // string | Stringified JSON object of custom data (optional)
	redirectUrl := "redirectUrl_example" // string | If you would like Paystack to redirect to a URL upon successful payment, specify the URL here. (optional)
	customFields := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | If you would like to accept custom fields, specify them here. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PageAPI.PageCreate(context.Background()).Name(name).Description(description).Amount(amount).Slug(slug).Metadata(metadata).RedirectUrl(redirectUrl).CustomFields(customFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PageAPI.PageCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageCreate`: Response
	fmt.Fprintf(os.Stdout, "Response from `PageAPI.PageCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPageCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Name of page | 
 **description** | **string** | The description of the page | 
 **amount** | **int32** | Amount should be in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR | 
 **slug** | **string** | URL slug you would like to be associated with this page. Page will be accessible at https://paystack.com/pay/[slug] | 
 **metadata** | **string** | Stringified JSON object of custom data | 
 **redirectUrl** | **string** | If you would like Paystack to redirect to a URL upon successful payment, specify the URL here. | 
 **customFields** | **[]map[string]interface{}** | If you would like to accept custom fields, specify them here. | 

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


## PageFetch

> Response PageFetch(ctx, id).Execute()

Fetch Page

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
	resp, r, err := apiClient.PageAPI.PageFetch(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PageAPI.PageFetch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageFetch`: Response
	fmt.Fprintf(os.Stdout, "Response from `PageAPI.PageFetch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageFetchRequest struct via the builder pattern


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


## PageList

> Response PageList(ctx).PerPage(perPage).Page(page).From(from).To(to).Execute()

List Pages

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
	resp, r, err := apiClient.PageAPI.PageList(context.Background()).PerPage(perPage).Page(page).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PageAPI.PageList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageList`: Response
	fmt.Fprintf(os.Stdout, "Response from `PageAPI.PageList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPageListRequest struct via the builder pattern


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


## PageUpdate

> Response PageUpdate(ctx, id).Name(name).Description(description).Amount(amount).Active(active).Execute()

Update Page

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
	name := "name_example" // string | Name of page (optional)
	description := "description_example" // string | The description of the page (optional)
	amount := int32(56) // int32 | Amount should be in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR (optional)
	active := true // bool | Set to false to deactivate page url (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PageAPI.PageUpdate(context.Background(), id).Name(name).Description(description).Amount(amount).Active(active).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PageAPI.PageUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageUpdate`: Response
	fmt.Fprintf(os.Stdout, "Response from `PageAPI.PageUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPageUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Name of page | 
 **description** | **string** | The description of the page | 
 **amount** | **int32** | Amount should be in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR | 
 **active** | **bool** | Set to false to deactivate page url | 

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

