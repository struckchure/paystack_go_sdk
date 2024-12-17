# \DedicatedVirtualAccountAPI

All URIs are relative to *https://api.paystack.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DedicatedAccountAddSplit**](DedicatedVirtualAccountAPI.md#DedicatedAccountAddSplit) | **Post** /dedicated_account/split | Split Dedicated Account Transaction
[**DedicatedAccountAvailableProviders**](DedicatedVirtualAccountAPI.md#DedicatedAccountAvailableProviders) | **Get** /dedicated_account/available_providers | Fetch Bank Providers
[**DedicatedAccountCreate**](DedicatedVirtualAccountAPI.md#DedicatedAccountCreate) | **Post** /dedicated_account | Create Dedicated Account
[**DedicatedAccountDeactivate**](DedicatedVirtualAccountAPI.md#DedicatedAccountDeactivate) | **Delete** /dedicated_account/{account_id} | Deactivate Dedicated Account
[**DedicatedAccountFetch**](DedicatedVirtualAccountAPI.md#DedicatedAccountFetch) | **Get** /dedicated_account/{account_id} | Fetch Dedicated Account
[**DedicatedAccountList**](DedicatedVirtualAccountAPI.md#DedicatedAccountList) | **Get** /dedicated_account | List Dedicated Accounts
[**DedicatedAccountRemoveSplit**](DedicatedVirtualAccountAPI.md#DedicatedAccountRemoveSplit) | **Delete** /dedicated_account/split | Remove Split from Dedicated Account



## DedicatedAccountAddSplit

> Response DedicatedAccountAddSplit(ctx).AccountNumber(accountNumber).Subaccount(subaccount).SplitCode(splitCode).Execute()

Split Dedicated Account Transaction

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
	accountNumber := "accountNumber_example" // string | Valid Dedicated virtual account
	subaccount := "subaccount_example" // string | Subaccount code of the account you want to split the transaction with (optional)
	splitCode := "splitCode_example" // string | Split code consisting of the lists of accounts you want to split the transaction with (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedVirtualAccountAPI.DedicatedAccountAddSplit(context.Background()).AccountNumber(accountNumber).Subaccount(subaccount).SplitCode(splitCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedVirtualAccountAPI.DedicatedAccountAddSplit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DedicatedAccountAddSplit`: Response
	fmt.Fprintf(os.Stdout, "Response from `DedicatedVirtualAccountAPI.DedicatedAccountAddSplit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDedicatedAccountAddSplitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountNumber** | **string** | Valid Dedicated virtual account | 
 **subaccount** | **string** | Subaccount code of the account you want to split the transaction with | 
 **splitCode** | **string** | Split code consisting of the lists of accounts you want to split the transaction with | 

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


## DedicatedAccountAvailableProviders

> Response DedicatedAccountAvailableProviders(ctx).Execute()

Fetch Bank Providers

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
	resp, r, err := apiClient.DedicatedVirtualAccountAPI.DedicatedAccountAvailableProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedVirtualAccountAPI.DedicatedAccountAvailableProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DedicatedAccountAvailableProviders`: Response
	fmt.Fprintf(os.Stdout, "Response from `DedicatedVirtualAccountAPI.DedicatedAccountAvailableProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDedicatedAccountAvailableProvidersRequest struct via the builder pattern


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


## DedicatedAccountCreate

> Response DedicatedAccountCreate(ctx).Customer(customer).PreferredBank(preferredBank).Subaccount(subaccount).SplitCode(splitCode).Execute()

Create Dedicated Account

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
	customer := "customer_example" // string | Customer ID or code
	preferredBank := "preferredBank_example" // string | The bank slug for preferred bank. To get a list of available banks, use the List Providers endpoint (optional)
	subaccount := "subaccount_example" // string | Subaccount code of the account you want to split the transaction with (optional)
	splitCode := "splitCode_example" // string | Split code consisting of the lists of accounts you want to split the transaction with (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedVirtualAccountAPI.DedicatedAccountCreate(context.Background()).Customer(customer).PreferredBank(preferredBank).Subaccount(subaccount).SplitCode(splitCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedVirtualAccountAPI.DedicatedAccountCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DedicatedAccountCreate`: Response
	fmt.Fprintf(os.Stdout, "Response from `DedicatedVirtualAccountAPI.DedicatedAccountCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDedicatedAccountCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customer** | **string** | Customer ID or code | 
 **preferredBank** | **string** | The bank slug for preferred bank. To get a list of available banks, use the List Providers endpoint | 
 **subaccount** | **string** | Subaccount code of the account you want to split the transaction with | 
 **splitCode** | **string** | Split code consisting of the lists of accounts you want to split the transaction with | 

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


## DedicatedAccountDeactivate

> Response DedicatedAccountDeactivate(ctx, accountId).Execute()

Deactivate Dedicated Account

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
	accountId := "accountId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedVirtualAccountAPI.DedicatedAccountDeactivate(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedVirtualAccountAPI.DedicatedAccountDeactivate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DedicatedAccountDeactivate`: Response
	fmt.Fprintf(os.Stdout, "Response from `DedicatedVirtualAccountAPI.DedicatedAccountDeactivate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDedicatedAccountDeactivateRequest struct via the builder pattern


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


## DedicatedAccountFetch

> Response DedicatedAccountFetch(ctx, accountId).Execute()

Fetch Dedicated Account

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
	accountId := "accountId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedVirtualAccountAPI.DedicatedAccountFetch(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedVirtualAccountAPI.DedicatedAccountFetch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DedicatedAccountFetch`: Response
	fmt.Fprintf(os.Stdout, "Response from `DedicatedVirtualAccountAPI.DedicatedAccountFetch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDedicatedAccountFetchRequest struct via the builder pattern


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


## DedicatedAccountList

> Response DedicatedAccountList(ctx).AccountNumber(accountNumber).Customer(customer).Active(active).Currency(currency).ProviderSlug(providerSlug).BankId(bankId).PerPage(perPage).Page(page).Execute()

List Dedicated Accounts

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
	accountNumber := "accountNumber_example" // string |  (optional)
	customer := "customer_example" // string |  (optional)
	active := true // bool |  (optional)
	currency := "currency_example" // string |  (optional)
	providerSlug := "providerSlug_example" // string |  (optional)
	bankId := "bankId_example" // string |  (optional)
	perPage := "perPage_example" // string |  (optional)
	page := "page_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedVirtualAccountAPI.DedicatedAccountList(context.Background()).AccountNumber(accountNumber).Customer(customer).Active(active).Currency(currency).ProviderSlug(providerSlug).BankId(bankId).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedVirtualAccountAPI.DedicatedAccountList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DedicatedAccountList`: Response
	fmt.Fprintf(os.Stdout, "Response from `DedicatedVirtualAccountAPI.DedicatedAccountList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDedicatedAccountListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountNumber** | **string** |  | 
 **customer** | **string** |  | 
 **active** | **bool** |  | 
 **currency** | **string** |  | 
 **providerSlug** | **string** |  | 
 **bankId** | **string** |  | 
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


## DedicatedAccountRemoveSplit

> Response DedicatedAccountRemoveSplit(ctx).AccountNumber(accountNumber).Subaccount(subaccount).SplitCode(splitCode).Execute()

Remove Split from Dedicated Account

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
	accountNumber := "accountNumber_example" // string | Valid Dedicated virtual account
	subaccount := "subaccount_example" // string | Subaccount code of the account you want to split the transaction with (optional)
	splitCode := "splitCode_example" // string | Split code consisting of the lists of accounts you want to split the transaction with (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DedicatedVirtualAccountAPI.DedicatedAccountRemoveSplit(context.Background()).AccountNumber(accountNumber).Subaccount(subaccount).SplitCode(splitCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DedicatedVirtualAccountAPI.DedicatedAccountRemoveSplit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DedicatedAccountRemoveSplit`: Response
	fmt.Fprintf(os.Stdout, "Response from `DedicatedVirtualAccountAPI.DedicatedAccountRemoveSplit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDedicatedAccountRemoveSplitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountNumber** | **string** | Valid Dedicated virtual account | 
 **subaccount** | **string** | Subaccount code of the account you want to split the transaction with | 
 **splitCode** | **string** | Split code consisting of the lists of accounts you want to split the transaction with | 

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

