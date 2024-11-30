# AttachmentObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** |  | 
**Filename** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] 
**Disposition** | Pointer to **string** |  | [optional] 
**ContentId** | Pointer to **string** |  | [optional] 

## Methods

### NewAttachmentObject

`func NewAttachmentObject(content string, filename string, ) *AttachmentObject`

NewAttachmentObject instantiates a new AttachmentObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentObjectWithDefaults

`func NewAttachmentObjectWithDefaults() *AttachmentObject`

NewAttachmentObjectWithDefaults instantiates a new AttachmentObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *AttachmentObject) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AttachmentObject) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AttachmentObject) SetContent(v string)`

SetContent sets Content field to given value.


### GetFilename

`func (o *AttachmentObject) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *AttachmentObject) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *AttachmentObject) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetType

`func (o *AttachmentObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttachmentObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AttachmentObject) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AttachmentObject) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDisposition

`func (o *AttachmentObject) GetDisposition() string`

GetDisposition returns the Disposition field if non-nil, zero value otherwise.

### GetDispositionOk

`func (o *AttachmentObject) GetDispositionOk() (*string, bool)`

GetDispositionOk returns a tuple with the Disposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisposition

`func (o *AttachmentObject) SetDisposition(v string)`

SetDisposition sets Disposition field to given value.

### HasDisposition

`func (o *AttachmentObject) HasDisposition() bool`

HasDisposition returns a boolean if a field has been set.

### GetContentId

`func (o *AttachmentObject) GetContentId() string`

GetContentId returns the ContentId field if non-nil, zero value otherwise.

### GetContentIdOk

`func (o *AttachmentObject) GetContentIdOk() (*string, bool)`

GetContentIdOk returns a tuple with the ContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentId

`func (o *AttachmentObject) SetContentId(v string)`

SetContentId sets ContentId field to given value.

### HasContentId

`func (o *AttachmentObject) HasContentId() bool`

HasContentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


