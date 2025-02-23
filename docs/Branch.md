# Branch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Id** | **string** | The id of the branch | 
**Version** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**ForkedBranchId** | **NullableString** | The id of the branch this branch was forked from | 
**Links** | [**ProjectLinks**](ProjectLinks.md) |  | 

## Methods

### NewBranch

`func NewBranch(name string, id string, version int32, createdAt time.Time, updatedAt time.Time, forkedBranchId NullableString, links ProjectLinks, ) *Branch`

NewBranch instantiates a new Branch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchWithDefaults

`func NewBranchWithDefaults() *Branch`

NewBranchWithDefaults instantiates a new Branch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Branch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Branch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Branch) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *Branch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Branch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Branch) SetId(v string)`

SetId sets Id field to given value.


### GetVersion

`func (o *Branch) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Branch) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Branch) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetCreatedAt

`func (o *Branch) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Branch) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Branch) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Branch) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Branch) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Branch) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetForkedBranchId

`func (o *Branch) GetForkedBranchId() string`

GetForkedBranchId returns the ForkedBranchId field if non-nil, zero value otherwise.

### GetForkedBranchIdOk

`func (o *Branch) GetForkedBranchIdOk() (*string, bool)`

GetForkedBranchIdOk returns a tuple with the ForkedBranchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForkedBranchId

`func (o *Branch) SetForkedBranchId(v string)`

SetForkedBranchId sets ForkedBranchId field to given value.


### SetForkedBranchIdNil

`func (o *Branch) SetForkedBranchIdNil(b bool)`

 SetForkedBranchIdNil sets the value for ForkedBranchId to be an explicit nil

### UnsetForkedBranchId
`func (o *Branch) UnsetForkedBranchId()`

UnsetForkedBranchId ensures that no value is present for ForkedBranchId, not even an explicit nil
### GetLinks

`func (o *Branch) GetLinks() ProjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Branch) GetLinksOk() (*ProjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Branch) SetLinks(v ProjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


