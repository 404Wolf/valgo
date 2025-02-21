# RootProjectFilesGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]FileRevisionMetadata**](FileRevisionMetadata.md) |  | 
**Links** | [**PaginationLinks**](PaginationLinks.md) |  | 

## Methods

### NewRootProjectFilesGet200Response

`func NewRootProjectFilesGet200Response(data []FileRevisionMetadata, links PaginationLinks, ) *RootProjectFilesGet200Response`

NewRootProjectFilesGet200Response instantiates a new RootProjectFilesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRootProjectFilesGet200ResponseWithDefaults

`func NewRootProjectFilesGet200ResponseWithDefaults() *RootProjectFilesGet200Response`

NewRootProjectFilesGet200ResponseWithDefaults instantiates a new RootProjectFilesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RootProjectFilesGet200Response) GetData() []FileRevisionMetadata`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RootProjectFilesGet200Response) GetDataOk() (*[]FileRevisionMetadata, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RootProjectFilesGet200Response) SetData(v []FileRevisionMetadata)`

SetData sets Data field to given value.


### GetLinks

`func (o *RootProjectFilesGet200Response) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RootProjectFilesGet200Response) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RootProjectFilesGet200Response) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


