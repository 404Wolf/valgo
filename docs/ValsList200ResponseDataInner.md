# ValsList200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValId** | **string** | Id of a val | 
**Version** | **int32** | Version of the val | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewValsList200ResponseDataInner

`func NewValsList200ResponseDataInner(valId string, version int32, createdAt time.Time, ) *ValsList200ResponseDataInner`

NewValsList200ResponseDataInner instantiates a new ValsList200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValsList200ResponseDataInnerWithDefaults

`func NewValsList200ResponseDataInnerWithDefaults() *ValsList200ResponseDataInner`

NewValsList200ResponseDataInnerWithDefaults instantiates a new ValsList200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValId

`func (o *ValsList200ResponseDataInner) GetValId() string`

GetValId returns the ValId field if non-nil, zero value otherwise.

### GetValIdOk

`func (o *ValsList200ResponseDataInner) GetValIdOk() (*string, bool)`

GetValIdOk returns a tuple with the ValId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValId

`func (o *ValsList200ResponseDataInner) SetValId(v string)`

SetValId sets ValId field to given value.


### GetVersion

`func (o *ValsList200ResponseDataInner) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ValsList200ResponseDataInner) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ValsList200ResponseDataInner) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetCreatedAt

`func (o *ValsList200ResponseDataInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ValsList200ResponseDataInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ValsList200ResponseDataInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


