# BlobListingItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Size** | Pointer to **int32** | Size in bytes of the object | [optional] 
**LastModified** | Pointer to **time.Time** | Creation date of the object | [optional] 

## Methods

### NewBlobListingItem

`func NewBlobListingItem(key string, ) *BlobListingItem`

NewBlobListingItem instantiates a new BlobListingItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlobListingItemWithDefaults

`func NewBlobListingItemWithDefaults() *BlobListingItem`

NewBlobListingItemWithDefaults instantiates a new BlobListingItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *BlobListingItem) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *BlobListingItem) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *BlobListingItem) SetKey(v string)`

SetKey sets Key field to given value.


### GetSize

`func (o *BlobListingItem) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BlobListingItem) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BlobListingItem) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *BlobListingItem) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetLastModified

`func (o *BlobListingItem) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *BlobListingItem) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *BlobListingItem) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *BlobListingItem) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


