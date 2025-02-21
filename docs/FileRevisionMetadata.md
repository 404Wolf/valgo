# FileRevisionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Id** | **string** | The id of the resource | 
**Path** | **string** |  | 
**Version** | **int32** |  | 
**UpdatedAt** | **time.Time** |  | 
**Type** | **string** |  | 
**Links** | [**FileRevisionMetadataLinks**](FileRevisionMetadataLinks.md) |  | 

## Methods

### NewFileRevisionMetadata

`func NewFileRevisionMetadata(name string, id string, path string, version int32, updatedAt time.Time, type_ string, links FileRevisionMetadataLinks, ) *FileRevisionMetadata`

NewFileRevisionMetadata instantiates a new FileRevisionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileRevisionMetadataWithDefaults

`func NewFileRevisionMetadataWithDefaults() *FileRevisionMetadata`

NewFileRevisionMetadataWithDefaults instantiates a new FileRevisionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FileRevisionMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileRevisionMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileRevisionMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *FileRevisionMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileRevisionMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileRevisionMetadata) SetId(v string)`

SetId sets Id field to given value.


### GetPath

`func (o *FileRevisionMetadata) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FileRevisionMetadata) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FileRevisionMetadata) SetPath(v string)`

SetPath sets Path field to given value.


### GetVersion

`func (o *FileRevisionMetadata) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FileRevisionMetadata) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FileRevisionMetadata) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetUpdatedAt

`func (o *FileRevisionMetadata) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FileRevisionMetadata) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FileRevisionMetadata) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetType

`func (o *FileRevisionMetadata) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FileRevisionMetadata) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FileRevisionMetadata) SetType(v string)`

SetType sets Type field to given value.


### GetLinks

`func (o *FileRevisionMetadata) GetLinks() FileRevisionMetadataLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FileRevisionMetadata) GetLinksOk() (*FileRevisionMetadataLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FileRevisionMetadata) SetLinks(v FileRevisionMetadataLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


