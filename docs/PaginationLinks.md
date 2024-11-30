# PaginationLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | **string** | URL of this page | 
**Prev** | Pointer to **string** | URL of the previous page, if any | [optional] 
**Next** | Pointer to **string** | URL of the next page, if any | [optional] 

## Methods

### NewPaginationLinks

`func NewPaginationLinks(self string, ) *PaginationLinks`

NewPaginationLinks instantiates a new PaginationLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationLinksWithDefaults

`func NewPaginationLinksWithDefaults() *PaginationLinks`

NewPaginationLinksWithDefaults instantiates a new PaginationLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *PaginationLinks) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *PaginationLinks) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *PaginationLinks) SetSelf(v string)`

SetSelf sets Self field to given value.


### GetPrev

`func (o *PaginationLinks) GetPrev() string`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *PaginationLinks) GetPrevOk() (*string, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *PaginationLinks) SetPrev(v string)`

SetPrev sets Prev field to given value.

### HasPrev

`func (o *PaginationLinks) HasPrev() bool`

HasPrev returns a boolean if a field has been set.

### GetNext

`func (o *PaginationLinks) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginationLinks) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginationLinks) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginationLinks) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


