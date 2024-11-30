# ValsList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ValsList200ResponseDataInner**](ValsList200ResponseDataInner.md) |  | 
**Links** | [**PaginationLinks**](PaginationLinks.md) |  | 

## Methods

### NewValsList200Response

`func NewValsList200Response(data []ValsList200ResponseDataInner, links PaginationLinks, ) *ValsList200Response`

NewValsList200Response instantiates a new ValsList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValsList200ResponseWithDefaults

`func NewValsList200ResponseWithDefaults() *ValsList200Response`

NewValsList200ResponseWithDefaults instantiates a new ValsList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ValsList200Response) GetData() []ValsList200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ValsList200Response) GetDataOk() (*[]ValsList200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ValsList200Response) SetData(v []ValsList200ResponseDataInner)`

SetData sets Data field to given value.


### GetLinks

`func (o *ValsList200Response) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ValsList200Response) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ValsList200Response) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


