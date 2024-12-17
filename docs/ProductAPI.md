# \ProductAPI

All URIs are relative to *https://api.paystack.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductCreate**](ProductAPI.md#ProductCreate) | **Post** /product | Create Product
[**ProductDelete**](ProductAPI.md#ProductDelete) | **Delete** /product/{id} | Delete Product
[**ProductFetch**](ProductAPI.md#ProductFetch) | **Get** /product/{id} | Fetch Product
[**ProductList**](ProductAPI.md#ProductList) | **Get** /product | List Products
[**ProductUpdate**](ProductAPI.md#ProductUpdate) | **Put** /product/{id} | Update product



## ProductCreate

> Response ProductCreate(ctx).Name(name).Description(description).Price(price).Currency(currency).Limited(limited).Quantity(quantity).Execute()

Create Product

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
	name := "name_example" // string | Name of product
	description := "description_example" // string | The description of the product
	price := int32(56) // int32 | Price should be in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR
	currency := "currency_example" // string | Currency in which price is set. Allowed values are: NGN, GHS, ZAR or USD
	limited := true // bool | Set to true if the product has limited stock. Leave as false if the product has unlimited stock (optional)
	quantity := int32(56) // int32 | Number of products in stock. Use if limited is true (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductCreate(context.Background()).Name(name).Description(description).Price(price).Currency(currency).Limited(limited).Quantity(quantity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductCreate`: Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Name of product | 
 **description** | **string** | The description of the product | 
 **price** | **int32** | Price should be in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR | 
 **currency** | **string** | Currency in which price is set. Allowed values are: NGN, GHS, ZAR or USD | 
 **limited** | **bool** | Set to true if the product has limited stock. Leave as false if the product has unlimited stock | 
 **quantity** | **int32** | Number of products in stock. Use if limited is true | 

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


## ProductDelete

> Response ProductDelete(ctx, id).Execute()

Delete Product

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
	resp, r, err := apiClient.ProductAPI.ProductDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductDelete`: Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductDeleteRequest struct via the builder pattern


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


## ProductFetch

> Response ProductFetch(ctx, id).Execute()

Fetch Product

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
	resp, r, err := apiClient.ProductAPI.ProductFetch(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductFetch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductFetch`: Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductFetch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductFetchRequest struct via the builder pattern


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


## ProductList

> Response ProductList(ctx).PerPage(perPage).Page(page).Active(active).From(from).To(to).Execute()

List Products

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
	perPage := int32(56) // int32 |  (optional)
	page := int32(56) // int32 |  (optional)
	active := true // bool |  (optional)
	from := time.Now() // time.Time | The start date (optional)
	to := time.Now() // time.Time | The end date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductList(context.Background()).PerPage(perPage).Page(page).Active(active).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductList`: Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** |  | 
 **page** | **int32** |  | 
 **active** | **bool** |  | 
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


## ProductUpdate

> Response ProductUpdate(ctx, id).Name(name).Description(description).Price(price).Currency(currency).Limited(limited).Quantity(quantity).Execute()

Update product

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
	name := "name_example" // string | Name of product (optional)
	description := "description_example" // string | The description of the product (optional)
	price := int32(56) // int32 | Price should be in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR (optional)
	currency := "currency_example" // string | Currency in which price is set. Allowed values are: NGN, GHS, ZAR or USD (optional)
	limited := true // bool | Set to true if the product has limited stock. Leave as false if the product has unlimited stock (optional)
	quantity := int32(56) // int32 | Number of products in stock. Use if limited is true (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAPI.ProductUpdate(context.Background(), id).Name(name).Description(description).Price(price).Currency(currency).Limited(limited).Quantity(quantity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.ProductUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductUpdate`: Response
	fmt.Fprintf(os.Stdout, "Response from `ProductAPI.ProductUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Name of product | 
 **description** | **string** | The description of the product | 
 **price** | **int32** | Price should be in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR | 
 **currency** | **string** | Currency in which price is set. Allowed values are: NGN, GHS, ZAR or USD | 
 **limited** | **bool** | Set to true if the product has limited stock. Leave as false if the product has unlimited stock | 
 **quantity** | **int32** | Number of products in stock. Use if limited is true | 

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

