# MeComments200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | **string** | Text of the given comment, in Markdown | 
**Id** | **string** | The comment’s id | 
**CreatedAt** | **time.Time** |  | 
**Author** | [**MeComments200ResponseDataInnerAuthor**](MeComments200ResponseDataInnerAuthor.md) |  | 
**Val** | [**MeComments200ResponseDataInnerVal**](MeComments200ResponseDataInnerVal.md) |  | 

## Methods

### NewMeComments200ResponseDataInner

`func NewMeComments200ResponseDataInner(comment string, id string, createdAt time.Time, author MeComments200ResponseDataInnerAuthor, val MeComments200ResponseDataInnerVal, ) *MeComments200ResponseDataInner`

NewMeComments200ResponseDataInner instantiates a new MeComments200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeComments200ResponseDataInnerWithDefaults

`func NewMeComments200ResponseDataInnerWithDefaults() *MeComments200ResponseDataInner`

NewMeComments200ResponseDataInnerWithDefaults instantiates a new MeComments200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *MeComments200ResponseDataInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *MeComments200ResponseDataInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *MeComments200ResponseDataInner) SetComment(v string)`

SetComment sets Comment field to given value.


### GetId

`func (o *MeComments200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MeComments200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MeComments200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *MeComments200ResponseDataInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MeComments200ResponseDataInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MeComments200ResponseDataInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetAuthor

`func (o *MeComments200ResponseDataInner) GetAuthor() MeComments200ResponseDataInnerAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *MeComments200ResponseDataInner) GetAuthorOk() (*MeComments200ResponseDataInnerAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *MeComments200ResponseDataInner) SetAuthor(v MeComments200ResponseDataInnerAuthor)`

SetAuthor sets Author field to given value.


### GetVal

`func (o *MeComments200ResponseDataInner) GetVal() MeComments200ResponseDataInnerVal`

GetVal returns the Val field if non-nil, zero value otherwise.

### GetValOk

`func (o *MeComments200ResponseDataInner) GetValOk() (*MeComments200ResponseDataInnerVal, bool)`

GetValOk returns a tuple with the Val field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVal

`func (o *MeComments200ResponseDataInner) SetVal(v MeComments200ResponseDataInnerVal)`

SetVal sets Val field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


