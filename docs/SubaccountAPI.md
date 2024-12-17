# \SubaccountAPI

All URIs are relative to *https://api.paystack.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubaccountCreate**](SubaccountAPI.md#SubaccountCreate) | **Post** /subaccount | Create Subaccount
[**SubaccountFetch**](SubaccountAPI.md#SubaccountFetch) | **Get** /subaccount/{code} | Fetch Subaccount
[**SubaccountList**](SubaccountAPI.md#SubaccountList) | **Get** /subaccount | List Subaccounts
[**SubaccountUpdate**](SubaccountAPI.md#SubaccountUpdate) | **Put** /subaccount/{code} | Update Subaccount



## SubaccountCreate

> Response SubaccountCreate(ctx).BusinessName(businessName).SettlementBank(settlementBank).AccountNumber(accountNumber).PercentageCharge(percentageCharge).Description(description).PrimaryContactEmail(primaryContactEmail).PrimaryContactName(primaryContactName).PrimaryContactPhone(primaryContactPhone).Metadata(metadata).Execute()

Create Subaccount

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
	businessName := "businessName_example" // string | Name of business for subaccount
	settlementBank := "settlementBank_example" // string | Bank code for the bank. You can get the list of Bank Codes by calling the List Banks endpoint.
	accountNumber := "accountNumber_example" // string | Bank account number
	percentageCharge := float32(3.4) // float32 | Customer's phone number
	description := "description_example" // string | A description for this subaccount (optional)
	primaryContactEmail := "primaryContactEmail_example" // string | A contact email for the subaccount (optional)
	primaryContactName := "primaryContactName_example" // string | The name of the contact person for this subaccount (optional)
	primaryContactPhone := "primaryContactPhone_example" // string | A phone number to call for this subaccount (optional)
	metadata := "metadata_example" // string | Stringified JSON object of custom data (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubaccountAPI.SubaccountCreate(context.Background()).BusinessName(businessName).SettlementBank(settlementBank).AccountNumber(accountNumber).PercentageCharge(percentageCharge).Description(description).PrimaryContactEmail(primaryContactEmail).PrimaryContactName(primaryContactName).PrimaryContactPhone(primaryContactPhone).Metadata(metadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubaccountAPI.SubaccountCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountCreate`: Response
	fmt.Fprintf(os.Stdout, "Response from `SubaccountAPI.SubaccountCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **businessName** | **string** | Name of business for subaccount | 
 **settlementBank** | **string** | Bank code for the bank. You can get the list of Bank Codes by calling the List Banks endpoint. | 
 **accountNumber** | **string** | Bank account number | 
 **percentageCharge** | **float32** | Customer&#39;s phone number | 
 **description** | **string** | A description for this subaccount | 
 **primaryContactEmail** | **string** | A contact email for the subaccount | 
 **primaryContactName** | **string** | The name of the contact person for this subaccount | 
 **primaryContactPhone** | **string** | A phone number to call for this subaccount | 
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


## SubaccountFetch

> Response SubaccountFetch(ctx, code).Execute()

Fetch Subaccount

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
	code := "code_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubaccountAPI.SubaccountFetch(context.Background(), code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubaccountAPI.SubaccountFetch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountFetch`: Response
	fmt.Fprintf(os.Stdout, "Response from `SubaccountAPI.SubaccountFetch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountFetchRequest struct via the builder pattern


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


## SubaccountList

> Response SubaccountList(ctx).PerPage(perPage).Page(page).From(from).To(to).Execute()

List Subaccounts

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
	resp, r, err := apiClient.SubaccountAPI.SubaccountList(context.Background()).PerPage(perPage).Page(page).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubaccountAPI.SubaccountList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountList`: Response
	fmt.Fprintf(os.Stdout, "Response from `SubaccountAPI.SubaccountList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountListRequest struct via the builder pattern


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


## SubaccountUpdate

> Response SubaccountUpdate(ctx, code).BusinessName(businessName).SettlementBank(settlementBank).AccountNumber(accountNumber).Active(active).PercentageCharge(percentageCharge).Description(description).PrimaryContactEmail(primaryContactEmail).PrimaryContactName(primaryContactName).PrimaryContactPhone(primaryContactPhone).Metadata(metadata).Execute()

Update Subaccount

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
	code := "code_example" // string | 
	businessName := "businessName_example" // string | Name of business for subaccount (optional)
	settlementBank := "settlementBank_example" // string | Bank code for the bank. You can get the list of Bank Codes by calling the List Banks endpoint. (optional)
	accountNumber := "accountNumber_example" // string | Bank account number (optional)
	active := true // bool | Activate or deactivate a subaccount (optional)
	percentageCharge := float32(3.4) // float32 | Customer's phone number (optional)
	description := "description_example" // string | A description for this subaccount (optional)
	primaryContactEmail := "primaryContactEmail_example" // string | A contact email for the subaccount (optional)
	primaryContactName := "primaryContactName_example" // string | The name of the contact person for this subaccount (optional)
	primaryContactPhone := "primaryContactPhone_example" // string | A phone number to call for this subaccount (optional)
	metadata := "metadata_example" // string | Stringified JSON object of custom data (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubaccountAPI.SubaccountUpdate(context.Background(), code).BusinessName(businessName).SettlementBank(settlementBank).AccountNumber(accountNumber).Active(active).PercentageCharge(percentageCharge).Description(description).PrimaryContactEmail(primaryContactEmail).PrimaryContactName(primaryContactName).PrimaryContactPhone(primaryContactPhone).Metadata(metadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubaccountAPI.SubaccountUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubaccountUpdate`: Response
	fmt.Fprintf(os.Stdout, "Response from `SubaccountAPI.SubaccountUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubaccountUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **businessName** | **string** | Name of business for subaccount | 
 **settlementBank** | **string** | Bank code for the bank. You can get the list of Bank Codes by calling the List Banks endpoint. | 
 **accountNumber** | **string** | Bank account number | 
 **active** | **bool** | Activate or deactivate a subaccount | 
 **percentageCharge** | **float32** | Customer&#39;s phone number | 
 **description** | **string** | A description for this subaccount | 
 **primaryContactEmail** | **string** | A contact email for the subaccount | 
 **primaryContactName** | **string** | The name of the contact person for this subaccount | 
 **primaryContactPhone** | **string** | A phone number to call for this subaccount | 
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

