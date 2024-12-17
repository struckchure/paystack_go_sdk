# PageCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of page | 
**Description** | Pointer to **string** | The description of the page | [optional] 
**Amount** | Pointer to **int32** | Amount should be in kobo if currency is NGN, pesewas, if currency is GHS, and cents, if currency is ZAR | [optional] 
**Slug** | Pointer to **string** | URL slug you would like to be associated with this page. Page will be accessible at https://paystack.com/pay/[slug] | [optional] 
**Metadata** | Pointer to **string** | Stringified JSON object of custom data | [optional] 
**RedirectUrl** | Pointer to **string** | If you would like Paystack to redirect to a URL upon successful payment, specify the URL here. | [optional] 
**CustomFields** | Pointer to **[]map[string]interface{}** | If you would like to accept custom fields, specify them here. | [optional] 

## Methods

### NewPageCreate

`func NewPageCreate(name string, ) *PageCreate`

NewPageCreate instantiates a new PageCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageCreateWithDefaults

`func NewPageCreateWithDefaults() *PageCreate`

NewPageCreateWithDefaults instantiates a new PageCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PageCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PageCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PageCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PageCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PageCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PageCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PageCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAmount

`func (o *PageCreate) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PageCreate) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PageCreate) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PageCreate) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetSlug

`func (o *PageCreate) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PageCreate) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PageCreate) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PageCreate) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetMetadata

`func (o *PageCreate) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PageCreate) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PageCreate) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PageCreate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *PageCreate) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *PageCreate) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *PageCreate) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *PageCreate) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetCustomFields

`func (o *PageCreate) GetCustomFields() []map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PageCreate) GetCustomFieldsOk() (*[]map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PageCreate) SetCustomFields(v []map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PageCreate) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


