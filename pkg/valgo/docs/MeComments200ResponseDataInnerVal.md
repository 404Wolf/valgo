# MeComments200ResponseDataInnerVal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the val that is being commented on | 
**Id** | **string** |  | 
**Version** | **int32** |  | 
**Privacy** | **string** | This val’s privacy setting. Unlisted vals do not appear on profile pages or elsewhere, but you can link to them. | 
**Author** | [**NullableAliasVal200ResponseAuthor**](AliasVal200ResponseAuthor.md) |  | 

## Methods

### NewMeComments200ResponseDataInnerVal

`func NewMeComments200ResponseDataInnerVal(name string, id string, version int32, privacy string, author NullableAliasVal200ResponseAuthor, ) *MeComments200ResponseDataInnerVal`

NewMeComments200ResponseDataInnerVal instantiates a new MeComments200ResponseDataInnerVal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeComments200ResponseDataInnerValWithDefaults

`func NewMeComments200ResponseDataInnerValWithDefaults() *MeComments200ResponseDataInnerVal`

NewMeComments200ResponseDataInnerValWithDefaults instantiates a new MeComments200ResponseDataInnerVal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MeComments200ResponseDataInnerVal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MeComments200ResponseDataInnerVal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MeComments200ResponseDataInnerVal) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *MeComments200ResponseDataInnerVal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MeComments200ResponseDataInnerVal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MeComments200ResponseDataInnerVal) SetId(v string)`

SetId sets Id field to given value.


### GetVersion

`func (o *MeComments200ResponseDataInnerVal) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MeComments200ResponseDataInnerVal) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MeComments200ResponseDataInnerVal) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetPrivacy

`func (o *MeComments200ResponseDataInnerVal) GetPrivacy() string`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *MeComments200ResponseDataInnerVal) GetPrivacyOk() (*string, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *MeComments200ResponseDataInnerVal) SetPrivacy(v string)`

SetPrivacy sets Privacy field to given value.


### GetAuthor

`func (o *MeComments200ResponseDataInnerVal) GetAuthor() AliasVal200ResponseAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *MeComments200ResponseDataInnerVal) GetAuthorOk() (*AliasVal200ResponseAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *MeComments200ResponseDataInnerVal) SetAuthor(v AliasVal200ResponseAuthor)`

SetAuthor sets Author field to given value.


### SetAuthorNil

`func (o *MeComments200ResponseDataInnerVal) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *MeComments200ResponseDataInnerVal) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


