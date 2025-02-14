# MeProjects200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Project**](Project.md) |  | 
**Links** | [**PaginationLinks**](PaginationLinks.md) |  | 

## Methods

### NewMeProjects200Response

`func NewMeProjects200Response(data []Project, links PaginationLinks, ) *MeProjects200Response`

NewMeProjects200Response instantiates a new MeProjects200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeProjects200ResponseWithDefaults

`func NewMeProjects200ResponseWithDefaults() *MeProjects200Response`

NewMeProjects200ResponseWithDefaults instantiates a new MeProjects200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MeProjects200Response) GetData() []Project`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MeProjects200Response) GetDataOk() (*[]Project, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MeProjects200Response) SetData(v []Project)`

SetData sets Data field to given value.


### GetLinks

`func (o *MeProjects200Response) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MeProjects200Response) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MeProjects200Response) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


