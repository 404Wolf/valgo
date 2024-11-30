# SearchVals200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]BasicVal**](BasicVal.md) |  | 
**Links** | [**PaginationLinks**](PaginationLinks.md) |  | 

## Methods

### NewSearchVals200Response

`func NewSearchVals200Response(data []BasicVal, links PaginationLinks, ) *SearchVals200Response`

NewSearchVals200Response instantiates a new SearchVals200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchVals200ResponseWithDefaults

`func NewSearchVals200ResponseWithDefaults() *SearchVals200Response`

NewSearchVals200ResponseWithDefaults instantiates a new SearchVals200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SearchVals200Response) GetData() []BasicVal`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SearchVals200Response) GetDataOk() (*[]BasicVal, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SearchVals200Response) SetData(v []BasicVal)`

SetData sets Data field to given value.


### GetLinks

`func (o *SearchVals200Response) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SearchVals200Response) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SearchVals200Response) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


