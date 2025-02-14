# BranchesList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Branch**](Branch.md) |  | 
**Links** | [**PaginationLinks**](PaginationLinks.md) |  | 

## Methods

### NewBranchesList200Response

`func NewBranchesList200Response(data []Branch, links PaginationLinks, ) *BranchesList200Response`

NewBranchesList200Response instantiates a new BranchesList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchesList200ResponseWithDefaults

`func NewBranchesList200ResponseWithDefaults() *BranchesList200Response`

NewBranchesList200ResponseWithDefaults instantiates a new BranchesList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BranchesList200Response) GetData() []Branch`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BranchesList200Response) GetDataOk() (*[]Branch, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BranchesList200Response) SetData(v []Branch)`

SetData sets Data field to given value.


### GetLinks

`func (o *BranchesList200Response) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BranchesList200Response) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BranchesList200Response) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


