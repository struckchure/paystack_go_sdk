# \TransactionAPI

All URIs are relative to *https://api.paystack.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransactionChargeAuthorization**](TransactionAPI.md#TransactionChargeAuthorization) | **Post** /transaction/charge_authorization | Charge Authorization
[**TransactionCheckAuthorization**](TransactionAPI.md#TransactionCheckAuthorization) | **Post** /transaction/check_authorization | Check Authorization
[**TransactionDownload**](TransactionAPI.md#TransactionDownload) | **Get** /transaction/export | Export Transactions
[**TransactionEvent**](TransactionAPI.md#TransactionEvent) | **Get** /transaction/{id}/event | Get Transaction Event
[**TransactionFetch**](TransactionAPI.md#TransactionFetch) | **Get** /transaction/{id} | Fetch Transaction
[**TransactionInitialize**](TransactionAPI.md#TransactionInitialize) | **Post** /transaction/initialize | Initialize Transaction
[**TransactionList**](TransactionAPI.md#TransactionList) | **Get** /transaction | List Transactions
[**TransactionPartialDebit**](TransactionAPI.md#TransactionPartialDebit) | **Post** /transaction/partial_debit | Partial Debit
[**TransactionSession**](TransactionAPI.md#TransactionSession) | **Get** /transaction/{id}/session | Get Transaction Session
[**TransactionTimeline**](TransactionAPI.md#TransactionTimeline) | **Get** /transaction/timeline/{id_or_reference} | Fetch Transaction Timeline
[**TransactionTotals**](TransactionAPI.md#TransactionTotals) | **Get** /transaction/totals | Transaction Totals
[**TransactionVerify**](TransactionAPI.md#TransactionVerify) | **Get** /transaction/verify/{reference} | Verify Transaction



## TransactionChargeAuthorization

> Response TransactionChargeAuthorization(ctx).Email(email).Amount(amount).AuthorizationCode(authorizationCode).Reference(reference).Currency(currency).Metadata(metadata).SplitCode(splitCode).Subaccount(subaccount).TransactionCharge(transactionCharge).Bearer(bearer).Queue(queue).Execute()

Charge Authorization

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
	email := "email_example" // string | Customer's email address
	amount := int32(56) // int32 | Amount should be in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR
	authorizationCode := "authorizationCode_example" // string | Valid authorization code to charge
	reference := "reference_example" // string | Unique transaction reference. Only -, ., = and alphanumeric characters allowed. (optional)
	currency := "currency_example" // string | The transaction currency (optional)
	metadata := "metadata_example" // string | Stringified JSON object of custom data (optional)
	splitCode := "splitCode_example" // string | The split code of the transaction split (optional)
	subaccount := "subaccount_example" // string | The code for the subaccount that owns the payment (optional)
	transactionCharge := "transactionCharge_example" // string | A flat fee to charge the subaccount for a transaction.  This overrides the split percentage set when the subaccount was created (optional)
	bearer := "bearer_example" // string | The beare of the transaction charge (optional)
	queue := true // bool | If you are making a scheduled charge call, it is a good idea to queue them so the processing system does not get overloaded causing transaction processing errors. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionAPI.TransactionChargeAuthorization(context.Background()).Email(email).Amount(amount).AuthorizationCode(authorizationCode).Reference(reference).Currency(currency).Metadata(metadata).SplitCode(splitCode).Subaccount(subaccount).TransactionCharge(transactionCharge).Bearer(bearer).Queue(queue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionAPI.TransactionChargeAuthorization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionChargeAuthorization`: Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionAPI.TransactionChargeAuthorization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionChargeAuthorizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Customer&#39;s email address | 
 **amount** | **int32** | Amount should be in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR | 
 **authorizationCode** | **string** | Valid authorization code to charge | 
 **reference** | **string** | Unique transaction reference. Only -, ., &#x3D; and alphanumeric characters allowed. | 
 **currency** | **string** | The transaction currency | 
 **metadata** | **string** | Stringified JSON object of custom data | 
 **splitCode** | **string** | The split code of the transaction split | 
 **subaccount** | **string** | The code for the subaccount that owns the payment | 
 **transactionCharge** | **string** | A flat fee to charge the subaccount for a transaction.  This overrides the split percentage set when the subaccount was created | 
 **bearer** | **string** | The beare of the transaction charge | 
 **queue** | **bool** | If you are making a scheduled charge call, it is a good idea to queue them so the processing system does not get overloaded causing transaction processing errors. | 

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


## TransactionCheckAuthorization

> Response TransactionCheckAuthorization(ctx).Email(email).Amount(amount).AuthorizationCode(authorizationCode).Currency(currency).Execute()

Check Authorization

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
	email := "email_example" // string | Customer's email address
	amount := int32(56) // int32 | Amount should be in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR
	authorizationCode := "authorizationCode_example" // string | Valid authorization code to charge (optional)
	currency := "currency_example" // string | The transaction currency (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionAPI.TransactionCheckAuthorization(context.Background()).Email(email).Amount(amount).AuthorizationCode(authorizationCode).Currency(currency).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionAPI.TransactionCheckAuthorization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionCheckAuthorization`: Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionAPI.TransactionCheckAuthorization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionCheckAuthorizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Customer&#39;s email address | 
 **amount** | **int32** | Amount should be in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR | 
 **authorizationCode** | **string** | Valid authorization code to charge | 
 **currency** | **string** | The transaction currency | 

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


## TransactionDownload

> Response TransactionDownload(ctx).PerPage(perPage).Page(page).From(from).To(to).Execute()

Export Transactions

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
	resp, r, err := apiClient.TransactionAPI.TransactionDownload(context.Background()).PerPage(perPage).Page(page).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionAPI.TransactionDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionDownload`: Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionAPI.TransactionDownload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionDownloadRequest struct via the builder pattern


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


## TransactionEvent

> Response TransactionEvent(ctx, id).Execute()

Get Transaction Event

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
	resp, r, err := apiClient.TransactionAPI.TransactionEvent(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionAPI.TransactionEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionEvent`: Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionAPI.TransactionEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransactionEventRequest struct via the builder pattern


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


## TransactionFetch

> Response TransactionFetch(ctx, id).Execute()

Fetch Transaction



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
	id := "id_example" // string | The ID of the transaction to fetch

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionAPI.TransactionFetch(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionAPI.TransactionFetch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionFetch`: Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionAPI.TransactionFetch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the transaction to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransactionFetchRequest struct via the builder pattern


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


## TransactionInitialize

> Response TransactionInitialize(ctx).Email(email).Amount(amount).Currency(currency).Reference(reference).CallbackUrl(callbackUrl).Plan(plan).InvoiceLimit(invoiceLimit).Metadata(metadata).Channels(channels).SplitCode(splitCode).Subaccount(subaccount).TransactionCharge(transactionCharge).Bearer(bearer).Execute()

Initialize Transaction



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
	email := "email_example" // string | Customer's email address
	amount := int32(56) // int32 | Amount should be in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR
	currency := "currency_example" // string | The transaction currency (optional)
	reference := "reference_example" // string | Unique transaction reference. Only -, ., = and alphanumeric characters allowed. (optional)
	callbackUrl := "callbackUrl_example" // string | Fully qualified url, e.g. https://example.com/ . Use this to override the callback url provided on the dashboard for this transaction (optional)
	plan := "plan_example" // string | If transaction is to create a subscription to a predefined plan, provide plan code here.  This would invalidate the value provided in amount (optional)
	invoiceLimit := int32(56) // int32 | Number of times to charge customer during subscription to plan (optional)
	metadata := "metadata_example" // string | Stringified JSON object of custom data (optional)
	channels := []string{"Inner_example"} // []string | An array of payment channels to control what channels you want to make available to the user to make a payment with (optional)
	splitCode := "splitCode_example" // string | The split code of the transaction split (optional)
	subaccount := "subaccount_example" // string | The code for the subaccount that owns the payment (optional)
	transactionCharge := "transactionCharge_example" // string | A flat fee to charge the subaccount for a transaction.  This overrides the split percentage set when the subaccount was created (optional)
	bearer := "bearer_example" // string | The beare of the transaction charge (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionAPI.TransactionInitialize(context.Background()).Email(email).Amount(amount).Currency(currency).Reference(reference).CallbackUrl(callbackUrl).Plan(plan).InvoiceLimit(invoiceLimit).Metadata(metadata).Channels(channels).SplitCode(splitCode).Subaccount(subaccount).TransactionCharge(transactionCharge).Bearer(bearer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionAPI.TransactionInitialize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionInitialize`: Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionAPI.TransactionInitialize`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionInitializeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Customer&#39;s email address | 
 **amount** | **int32** | Amount should be in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR | 
 **currency** | **string** | The transaction currency | 
 **reference** | **string** | Unique transaction reference. Only -, ., &#x3D; and alphanumeric characters allowed. | 
 **callbackUrl** | **string** | Fully qualified url, e.g. https://example.com/ . Use this to override the callback url provided on the dashboard for this transaction | 
 **plan** | **string** | If transaction is to create a subscription to a predefined plan, provide plan code here.  This would invalidate the value provided in amount | 
 **invoiceLimit** | **int32** | Number of times to charge customer during subscription to plan | 
 **metadata** | **string** | Stringified JSON object of custom data | 
 **channels** | **[]string** | An array of payment channels to control what channels you want to make available to the user to make a payment with | 
 **splitCode** | **string** | The split code of the transaction split | 
 **subaccount** | **string** | The code for the subaccount that owns the payment | 
 **transactionCharge** | **string** | A flat fee to charge the subaccount for a transaction.  This overrides the split percentage set when the subaccount was created | 
 **bearer** | **string** | The beare of the transaction charge | 

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


## TransactionList

> Response TransactionList(ctx).PerPage(perPage).Page(page).From(from).To(to).Execute()

List Transactions



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
	resp, r, err := apiClient.TransactionAPI.TransactionList(context.Background()).PerPage(perPage).Page(page).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionAPI.TransactionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionList`: Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionAPI.TransactionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionListRequest struct via the builder pattern


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


## TransactionPartialDebit

> Response TransactionPartialDebit(ctx).Email(email).Amount(amount).AuthorizationCode(authorizationCode).Currency(currency).Reference(reference).AtLeast(atLeast).Execute()

Partial Debit

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
	email := "email_example" // string | Customer's email address
	amount := int32(56) // int32 | Amount should be in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR
	authorizationCode := "authorizationCode_example" // string | Valid authorization code to charge
	currency := "currency_example" // string | The transaction currency
	reference := "reference_example" // string | Unique transaction reference. Only -, ., = and alphanumeric characters allowed. (optional)
	atLeast := "atLeast_example" // string | Minimum amount to charge (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionAPI.TransactionPartialDebit(context.Background()).Email(email).Amount(amount).AuthorizationCode(authorizationCode).Currency(currency).Reference(reference).AtLeast(atLeast).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionAPI.TransactionPartialDebit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionPartialDebit`: Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionAPI.TransactionPartialDebit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionPartialDebitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Customer&#39;s email address | 
 **amount** | **int32** | Amount should be in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR | 
 **authorizationCode** | **string** | Valid authorization code to charge | 
 **currency** | **string** | The transaction currency | 
 **reference** | **string** | Unique transaction reference. Only -, ., &#x3D; and alphanumeric characters allowed. | 
 **atLeast** | **string** | Minimum amount to charge | 

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


## TransactionSession

> Response TransactionSession(ctx, id).Execute()

Get Transaction Session

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
	resp, r, err := apiClient.TransactionAPI.TransactionSession(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionAPI.TransactionSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionSession`: Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionAPI.TransactionSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransactionSessionRequest struct via the builder pattern


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


## TransactionTimeline

> Response TransactionTimeline(ctx, idOrReference).Execute()

Fetch Transaction Timeline



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
	idOrReference := "idOrReference_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionAPI.TransactionTimeline(context.Background(), idOrReference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionAPI.TransactionTimeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionTimeline`: Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionAPI.TransactionTimeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrReference** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransactionTimelineRequest struct via the builder pattern


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


## TransactionTotals

> Response TransactionTotals(ctx).PerPage(perPage).Page(page).From(from).To(to).Execute()

Transaction Totals



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
	resp, r, err := apiClient.TransactionAPI.TransactionTotals(context.Background()).PerPage(perPage).Page(page).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionAPI.TransactionTotals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionTotals`: Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionAPI.TransactionTotals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionTotalsRequest struct via the builder pattern


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


## TransactionVerify

> Response TransactionVerify(ctx, reference).Execute()

Verify Transaction



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
	reference := "reference_example" // string | The transaction reference to verify

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionAPI.TransactionVerify(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionAPI.TransactionVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionVerify`: Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionAPI.TransactionVerify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | The transaction reference to verify | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransactionVerifyRequest struct via the builder pattern


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

