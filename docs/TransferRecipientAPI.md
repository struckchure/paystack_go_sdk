# \TransferRecipientAPI

All URIs are relative to *https://api.paystack.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransferrecipientBulk**](TransferRecipientAPI.md#TransferrecipientBulk) | **Post** /transferrecipient/bulk | Bulk Create Transfer Recipient
[**TransferrecipientCodeDelete**](TransferRecipientAPI.md#TransferrecipientCodeDelete) | **Delete** /transferrecipient/{code} | Delete Transfer Recipient
[**TransferrecipientCodePut**](TransferRecipientAPI.md#TransferrecipientCodePut) | **Put** /transferrecipient/{code} | Update Transfer recipient
[**TransferrecipientCreate**](TransferRecipientAPI.md#TransferrecipientCreate) | **Post** /transferrecipient | Create Transfer Recipient
[**TransferrecipientFetch**](TransferRecipientAPI.md#TransferrecipientFetch) | **Get** /transferrecipient/{code} | Fetch Transfer recipient
[**TransferrecipientList**](TransferRecipientAPI.md#TransferrecipientList) | **Get** /transferrecipient | List Transfer Recipients



## TransferrecipientBulk

> Response TransferrecipientBulk(ctx).Batch(batch).Execute()

Bulk Create Transfer Recipient

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
	batch := []openapiclient.TransferRecipientCreate{*openapiclient.NewTransferRecipientCreate("Type_example", "Name_example", "AccountNumber_example", "BankCode_example")} // []TransferRecipientCreate | A list of transfer recipient object. Each object should contain type, name, and bank_code.  Any Create Transfer Recipient param can also be passed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransferRecipientAPI.TransferrecipientBulk(context.Background()).Batch(batch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransferRecipientAPI.TransferrecipientBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferrecipientBulk`: Response
	fmt.Fprintf(os.Stdout, "Response from `TransferRecipientAPI.TransferrecipientBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferrecipientBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batch** | [**[]TransferRecipientCreate**](TransferRecipientCreate.md) | A list of transfer recipient object. Each object should contain type, name, and bank_code.  Any Create Transfer Recipient param can also be passed. | 

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


## TransferrecipientCodeDelete

> Response TransferrecipientCodeDelete(ctx, code).Execute()

Delete Transfer Recipient

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
	code := "code_example" // string | Transfer recipient code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransferRecipientAPI.TransferrecipientCodeDelete(context.Background(), code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransferRecipientAPI.TransferrecipientCodeDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferrecipientCodeDelete`: Response
	fmt.Fprintf(os.Stdout, "Response from `TransferRecipientAPI.TransferrecipientCodeDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Transfer recipient code | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferrecipientCodeDeleteRequest struct via the builder pattern


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


## TransferrecipientCodePut

> Response TransferrecipientCodePut(ctx, code).Name(name).Email(email).Execute()

Update Transfer recipient

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
	code := "code_example" // string | Transfer recipient code
	name := "name_example" // string | Recipient's name (optional)
	email := "email_example" // string | Recipient's email address (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransferRecipientAPI.TransferrecipientCodePut(context.Background(), code).Name(name).Email(email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransferRecipientAPI.TransferrecipientCodePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferrecipientCodePut`: Response
	fmt.Fprintf(os.Stdout, "Response from `TransferRecipientAPI.TransferrecipientCodePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Transfer recipient code | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferrecipientCodePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Recipient&#39;s name | 
 **email** | **string** | Recipient&#39;s email address | 

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


## TransferrecipientCreate

> Response TransferrecipientCreate(ctx).Type_(type_).Name(name).AccountNumber(accountNumber).BankCode(bankCode).Description(description).Currency(currency).AuthorizationCode(authorizationCode).Metadata(metadata).Execute()

Create Transfer Recipient

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
	type_ := "type__example" // string | Recipient Type (Only nuban at this time)
	name := "name_example" // string | Recipient's name
	accountNumber := "accountNumber_example" // string | Recipient's bank account number
	bankCode := "bankCode_example" // string | Recipient's bank code. You can get the list of Bank Codes by calling the List Banks endpoint
	description := "description_example" // string | A description for this recipient (optional)
	currency := "currency_example" // string | Currency for the account receiving the transfer (optional)
	authorizationCode := "authorizationCode_example" // string | An authorization code from a previous transaction (optional)
	metadata := "metadata_example" // string | Stringified JSON object of custom data (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransferRecipientAPI.TransferrecipientCreate(context.Background()).Type_(type_).Name(name).AccountNumber(accountNumber).BankCode(bankCode).Description(description).Currency(currency).AuthorizationCode(authorizationCode).Metadata(metadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransferRecipientAPI.TransferrecipientCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferrecipientCreate`: Response
	fmt.Fprintf(os.Stdout, "Response from `TransferRecipientAPI.TransferrecipientCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferrecipientCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Recipient Type (Only nuban at this time) | 
 **name** | **string** | Recipient&#39;s name | 
 **accountNumber** | **string** | Recipient&#39;s bank account number | 
 **bankCode** | **string** | Recipient&#39;s bank code. You can get the list of Bank Codes by calling the List Banks endpoint | 
 **description** | **string** | A description for this recipient | 
 **currency** | **string** | Currency for the account receiving the transfer | 
 **authorizationCode** | **string** | An authorization code from a previous transaction | 
 **metadata** | **string** | Stringified JSON object of custom data | 

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


## TransferrecipientFetch

> Response TransferrecipientFetch(ctx, code).Execute()

Fetch Transfer recipient

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
	code := "code_example" // string | Transfer recipient code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransferRecipientAPI.TransferrecipientFetch(context.Background(), code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransferRecipientAPI.TransferrecipientFetch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferrecipientFetch`: Response
	fmt.Fprintf(os.Stdout, "Response from `TransferRecipientAPI.TransferrecipientFetch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Transfer recipient code | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferrecipientFetchRequest struct via the builder pattern


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


## TransferrecipientList

> Response TransferrecipientList(ctx).PerPage(perPage).Page(page).From(from).To(to).Execute()

List Transfer Recipients

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
	resp, r, err := apiClient.TransferRecipientAPI.TransferrecipientList(context.Background()).PerPage(perPage).Page(page).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransferRecipientAPI.TransferrecipientList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransferrecipientList`: Response
	fmt.Fprintf(os.Stdout, "Response from `TransferRecipientAPI.TransferrecipientList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferrecipientListRequest struct via the builder pattern


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

