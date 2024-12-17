# \SplitAPI

All URIs are relative to *https://api.paystack.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SplitAddSubaccount**](SplitAPI.md#SplitAddSubaccount) | **Post** /split/{id}/subaccount/add | Add Subaccount to Split
[**SplitCreate**](SplitAPI.md#SplitCreate) | **Post** /split | Create Split
[**SplitFetch**](SplitAPI.md#SplitFetch) | **Get** /split/{id} | Fetch Split
[**SplitList**](SplitAPI.md#SplitList) | **Get** /split | List/Search Splits
[**SplitRemoveSubaccount**](SplitAPI.md#SplitRemoveSubaccount) | **Post** /split/{id}/subaccount/remove | Remove Subaccount from split
[**SplitUpdate**](SplitAPI.md#SplitUpdate) | **Put** /split/{id} | Update Split



## SplitAddSubaccount

> Response SplitAddSubaccount(ctx, id).Subaccount(subaccount).Share(share).Execute()

Add Subaccount to Split

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
	id := "application/json" // string | 
	subaccount := "subaccount_example" // string | Subaccount code of the customer or partner (optional)
	share := "share_example" // string | The percentage or flat quota of the customer or partner (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SplitAPI.SplitAddSubaccount(context.Background(), id).Subaccount(subaccount).Share(share).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SplitAPI.SplitAddSubaccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SplitAddSubaccount`: Response
	fmt.Fprintf(os.Stdout, "Response from `SplitAPI.SplitAddSubaccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSplitAddSubaccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subaccount** | **string** | Subaccount code of the customer or partner | 
 **share** | **string** | The percentage or flat quota of the customer or partner | 

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


## SplitCreate

> Response SplitCreate(ctx).Name(name).Type_(type_).Subaccounts(subaccounts).Currency(currency).BearerType(bearerType).BearerSubaccount(bearerSubaccount).Execute()

Create Split

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
	name := "name_example" // string | Name of the transaction split
	type_ := "type__example" // string | The type of transaction split you want to create.
	subaccounts := []openapiclient.SplitSubaccounts{*openapiclient.NewSplitSubaccounts()} // []SplitSubaccounts | A list of object containing subaccount code and number of shares
	currency := "currency_example" // string | The transaction currency
	bearerType := "bearerType_example" // string | This allows you specify how the transaction charge should be processed (optional)
	bearerSubaccount := "bearerSubaccount_example" // string | This is the subaccount code of the customer or partner that would bear the transaction charge if you specified subaccount as the bearer type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SplitAPI.SplitCreate(context.Background()).Name(name).Type_(type_).Subaccounts(subaccounts).Currency(currency).BearerType(bearerType).BearerSubaccount(bearerSubaccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SplitAPI.SplitCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SplitCreate`: Response
	fmt.Fprintf(os.Stdout, "Response from `SplitAPI.SplitCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSplitCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Name of the transaction split | 
 **type_** | **string** | The type of transaction split you want to create. | 
 **subaccounts** | [**[]SplitSubaccounts**](SplitSubaccounts.md) | A list of object containing subaccount code and number of shares | 
 **currency** | **string** | The transaction currency | 
 **bearerType** | **string** | This allows you specify how the transaction charge should be processed | 
 **bearerSubaccount** | **string** | This is the subaccount code of the customer or partner that would bear the transaction charge if you specified subaccount as the bearer type | 

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


## SplitFetch

> Response SplitFetch(ctx, id).Execute()

Fetch Split

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
	resp, r, err := apiClient.SplitAPI.SplitFetch(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SplitAPI.SplitFetch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SplitFetch`: Response
	fmt.Fprintf(os.Stdout, "Response from `SplitAPI.SplitFetch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSplitFetchRequest struct via the builder pattern


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


## SplitList

> Response SplitList(ctx).Name(name).Active(active).SortBy(sortBy).From(from).To(to).PerPage(perPage).Page(page).Execute()

List/Search Splits

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
	name := "name_example" // string |  (optional)
	active := "active_example" // string |  (optional)
	sortBy := "sortBy_example" // string |  (optional)
	from := "from_example" // string |  (optional)
	to := "to_example" // string |  (optional)
	perPage := "perPage_example" // string |  (optional)
	page := "page_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SplitAPI.SplitList(context.Background()).Name(name).Active(active).SortBy(sortBy).From(from).To(to).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SplitAPI.SplitList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SplitList`: Response
	fmt.Fprintf(os.Stdout, "Response from `SplitAPI.SplitList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSplitListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **active** | **string** |  | 
 **sortBy** | **string** |  | 
 **from** | **string** |  | 
 **to** | **string** |  | 
 **perPage** | **string** |  | 
 **page** | **string** |  | 

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


## SplitRemoveSubaccount

> Response SplitRemoveSubaccount(ctx, id).Subaccount(subaccount).Share(share).Execute()

Remove Subaccount from split

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
	subaccount := "subaccount_example" // string | Subaccount code of the customer or partner (optional)
	share := "share_example" // string | The percentage or flat quota of the customer or partner (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SplitAPI.SplitRemoveSubaccount(context.Background(), id).Subaccount(subaccount).Share(share).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SplitAPI.SplitRemoveSubaccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SplitRemoveSubaccount`: Response
	fmt.Fprintf(os.Stdout, "Response from `SplitAPI.SplitRemoveSubaccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSplitRemoveSubaccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subaccount** | **string** | Subaccount code of the customer or partner | 
 **share** | **string** | The percentage or flat quota of the customer or partner | 

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


## SplitUpdate

> Response SplitUpdate(ctx, id).Name(name).Active(active).BearerType(bearerType).BearerSubaccount(bearerSubaccount).Execute()

Update Split

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
	name := "name_example" // string | Name of the transaction split (optional)
	active := true // bool | Toggle status of split. When true, the split is active, else it's inactive (optional)
	bearerType := "bearerType_example" // string | This allows you specify how the transaction charge should be processed (optional)
	bearerSubaccount := "bearerSubaccount_example" // string | This is the subaccount code of the customer or partner that would bear the transaction charge if you specified subaccount as the bearer type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SplitAPI.SplitUpdate(context.Background(), id).Name(name).Active(active).BearerType(bearerType).BearerSubaccount(bearerSubaccount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SplitAPI.SplitUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SplitUpdate`: Response
	fmt.Fprintf(os.Stdout, "Response from `SplitAPI.SplitUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSplitUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Name of the transaction split | 
 **active** | **bool** | Toggle status of split. When true, the split is active, else it&#39;s inactive | 
 **bearerType** | **string** | This allows you specify how the transaction charge should be processed | 
 **bearerSubaccount** | **string** | This is the subaccount code of the customer or partner that would bear the transaction charge if you specified subaccount as the bearer type | 

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

