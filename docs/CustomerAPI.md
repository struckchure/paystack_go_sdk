# \CustomerAPI

All URIs are relative to *https://api.paystack.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomerCreate**](CustomerAPI.md#CustomerCreate) | **Post** /customer | Create Customer
[**CustomerDeactivateAuthorization**](CustomerAPI.md#CustomerDeactivateAuthorization) | **Post** /customer/deactivate_authorization | Deactivate Authorization
[**CustomerFetch**](CustomerAPI.md#CustomerFetch) | **Get** /customer/{code} | Fetch Customer
[**CustomerList**](CustomerAPI.md#CustomerList) | **Get** /customer | List Customers
[**CustomerRiskAction**](CustomerAPI.md#CustomerRiskAction) | **Post** /customer/set_risk_action | White/blacklist Customer
[**CustomerUpdate**](CustomerAPI.md#CustomerUpdate) | **Put** /customer/{code} | Update Customer
[**CustomerValidate**](CustomerAPI.md#CustomerValidate) | **Post** /customer/{code}/identification | Validate Customer



## CustomerCreate

> Response CustomerCreate(ctx).Email(email).FirstName(firstName).LastName(lastName).Phone(phone).Metadata(metadata).Execute()

Create Customer

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
	firstName := "firstName_example" // string | Customer's first name (optional)
	lastName := "lastName_example" // string | Customer's last name (optional)
	phone := "phone_example" // string | Customer's phone number (optional)
	metadata := "metadata_example" // string | Stringified JSON object of custom data (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.CustomerCreate(context.Background()).Email(email).FirstName(firstName).LastName(lastName).Phone(phone).Metadata(metadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.CustomerCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerCreate`: Response
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.CustomerCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Customer&#39;s email address | 
 **firstName** | **string** | Customer&#39;s first name | 
 **lastName** | **string** | Customer&#39;s last name | 
 **phone** | **string** | Customer&#39;s phone number | 
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


## CustomerDeactivateAuthorization

> Response CustomerDeactivateAuthorization(ctx).AuthorizationCode(authorizationCode).Execute()

Deactivate Authorization



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
	authorizationCode := "authorizationCode_example" // string | Authorization code to be deactivated

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.CustomerDeactivateAuthorization(context.Background()).AuthorizationCode(authorizationCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.CustomerDeactivateAuthorization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerDeactivateAuthorization`: Response
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.CustomerDeactivateAuthorization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerDeactivateAuthorizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorizationCode** | **string** | Authorization code to be deactivated | 

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


## CustomerFetch

> Response CustomerFetch(ctx, code).Execute()

Fetch Customer

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
	resp, r, err := apiClient.CustomerAPI.CustomerFetch(context.Background(), code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.CustomerFetch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerFetch`: Response
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.CustomerFetch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerFetchRequest struct via the builder pattern


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


## CustomerList

> Response CustomerList(ctx).UseCursor(useCursor).Next(next).Previous(previous).From(from).To(to).PerPage(perPage).Page(page).Execute()

List Customers



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
	useCursor := true // bool |  (optional)
	next := "next_example" // string |  (optional)
	previous := "previous_example" // string |  (optional)
	from := "from_example" // string |  (optional)
	to := "to_example" // string |  (optional)
	perPage := "perPage_example" // string |  (optional)
	page := "page_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.CustomerList(context.Background()).UseCursor(useCursor).Next(next).Previous(previous).From(from).To(to).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.CustomerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerList`: Response
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.CustomerList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **useCursor** | **bool** |  | 
 **next** | **string** |  | 
 **previous** | **string** |  | 
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


## CustomerRiskAction

> Response CustomerRiskAction(ctx).Customer(customer).RiskAction(riskAction).Execute()

White/blacklist Customer



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
	customer := "customer_example" // string | Customer's code, or email address
	riskAction := "riskAction_example" // string | One of the possible risk actions [ default, allow, deny ]. allow to whitelist.  deny to blacklist. Customers start with a default risk action.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.CustomerRiskAction(context.Background()).Customer(customer).RiskAction(riskAction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.CustomerRiskAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerRiskAction`: Response
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.CustomerRiskAction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerRiskActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customer** | **string** | Customer&#39;s code, or email address | 
 **riskAction** | **string** | One of the possible risk actions [ default, allow, deny ]. allow to whitelist.  deny to blacklist. Customers start with a default risk action.  | 

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


## CustomerUpdate

> Response CustomerUpdate(ctx, code).FirstName(firstName).LastName(lastName).Phone(phone).Metadata(metadata).Execute()

Update Customer

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
	firstName := "firstName_example" // string | Customer's first name (optional)
	lastName := "lastName_example" // string | Customer's last name (optional)
	phone := "phone_example" // string | Customer's phone number (optional)
	metadata := "metadata_example" // string | Stringified JSON object of custom data (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.CustomerUpdate(context.Background(), code).FirstName(firstName).LastName(lastName).Phone(phone).Metadata(metadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.CustomerUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerUpdate`: Response
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.CustomerUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **firstName** | **string** | Customer&#39;s first name | 
 **lastName** | **string** | Customer&#39;s last name | 
 **phone** | **string** | Customer&#39;s phone number | 
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


## CustomerValidate

> Accepted CustomerValidate(ctx, code).FirstName(firstName).LastName(lastName).Type_(type_).Country(country).Bvn(bvn).BankCode(bankCode).AccountNumber(accountNumber).Value(value).Execute()

Validate Customer



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
	firstName := "firstName_example" // string | Customer's first name
	lastName := "lastName_example" // string | Customer's last name
	type_ := "type__example" // string | Predefined types of identification.
	country := "country_example" // string | Two-letter country code of identification issuer
	bvn := "bvn_example" // string | Customer's Bank Verification Number
	bankCode := "bankCode_example" // string | You can get the list of bank codes by calling the List Banks endpoint (https://api.paystack.co/bank).
	accountNumber := "accountNumber_example" // string | Customer's bank account number.
	value := "value_example" // string | Customer's identification number. Required if type is bvn (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerAPI.CustomerValidate(context.Background(), code).FirstName(firstName).LastName(lastName).Type_(type_).Country(country).Bvn(bvn).BankCode(bankCode).AccountNumber(accountNumber).Value(value).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerAPI.CustomerValidate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerValidate`: Accepted
	fmt.Fprintf(os.Stdout, "Response from `CustomerAPI.CustomerValidate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **firstName** | **string** | Customer&#39;s first name | 
 **lastName** | **string** | Customer&#39;s last name | 
 **type_** | **string** | Predefined types of identification. | 
 **country** | **string** | Two-letter country code of identification issuer | 
 **bvn** | **string** | Customer&#39;s Bank Verification Number | 
 **bankCode** | **string** | You can get the list of bank codes by calling the List Banks endpoint (https://api.paystack.co/bank). | 
 **accountNumber** | **string** | Customer&#39;s bank account number. | 
 **value** | **string** | Customer&#39;s identification number. Required if type is bvn | 

### Return type

[**Accepted**](Accepted.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

