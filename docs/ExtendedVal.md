# ExtendedVal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of this val | 
**Id** | **string** | This val&#39;s id | 
**Version** | **int32** | The version of this val, starting at zero | 
**Code** | **NullableString** | TypeScript code associated with this val | 
**Public** | **bool** | Whether this val is available publicly on Val Town | 
**CreatedAt** | **time.Time** |  | 
**Privacy** | **string** | This resource&#39;s privacy setting. Unlisted resources do not appear on profile pages or elsewhere, but you can link to them. | 
**Type** | **string** | The type of a val. HTTP can receive web requests, Email can receive emails, Cron runs periodically, and Script can be used for libraries or one-off calculations | 
**Url** | **string** | The URL of this resource on the Val Town website | 
**Links** | [**AliasVal200ResponseLinks**](AliasVal200ResponseLinks.md) |  | 
**Author** | [**NullableAliasVal200ResponseAuthor**](AliasVal200ResponseAuthor.md) |  | 
**VersionCreatedAt** | Pointer to **time.Time** |  | [optional] 
**LikeCount** | **int32** | How many likes this val has received | 
**ReferenceCount** | **int32** |  | 
**Readme** | **NullableString** | This val&#39;s readme, as Markdown | 

## Methods

### NewExtendedVal

`func NewExtendedVal(name string, id string, version int32, code NullableString, public bool, createdAt time.Time, privacy string, type_ string, url string, links AliasVal200ResponseLinks, author NullableAliasVal200ResponseAuthor, likeCount int32, referenceCount int32, readme NullableString, ) *ExtendedVal`

NewExtendedVal instantiates a new ExtendedVal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedValWithDefaults

`func NewExtendedValWithDefaults() *ExtendedVal`

NewExtendedValWithDefaults instantiates a new ExtendedVal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExtendedVal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtendedVal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtendedVal) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *ExtendedVal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExtendedVal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExtendedVal) SetId(v string)`

SetId sets Id field to given value.


### GetVersion

`func (o *ExtendedVal) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExtendedVal) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExtendedVal) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetCode

`func (o *ExtendedVal) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ExtendedVal) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ExtendedVal) SetCode(v string)`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *ExtendedVal) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ExtendedVal) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetPublic

`func (o *ExtendedVal) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *ExtendedVal) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *ExtendedVal) SetPublic(v bool)`

SetPublic sets Public field to given value.


### GetCreatedAt

`func (o *ExtendedVal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExtendedVal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExtendedVal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetPrivacy

`func (o *ExtendedVal) GetPrivacy() string`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *ExtendedVal) GetPrivacyOk() (*string, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *ExtendedVal) SetPrivacy(v string)`

SetPrivacy sets Privacy field to given value.


### GetType

`func (o *ExtendedVal) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExtendedVal) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExtendedVal) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *ExtendedVal) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ExtendedVal) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ExtendedVal) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetLinks

`func (o *ExtendedVal) GetLinks() AliasVal200ResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExtendedVal) GetLinksOk() (*AliasVal200ResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExtendedVal) SetLinks(v AliasVal200ResponseLinks)`

SetLinks sets Links field to given value.


### GetAuthor

`func (o *ExtendedVal) GetAuthor() AliasVal200ResponseAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ExtendedVal) GetAuthorOk() (*AliasVal200ResponseAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ExtendedVal) SetAuthor(v AliasVal200ResponseAuthor)`

SetAuthor sets Author field to given value.


### SetAuthorNil

`func (o *ExtendedVal) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *ExtendedVal) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetVersionCreatedAt

`func (o *ExtendedVal) GetVersionCreatedAt() time.Time`

GetVersionCreatedAt returns the VersionCreatedAt field if non-nil, zero value otherwise.

### GetVersionCreatedAtOk

`func (o *ExtendedVal) GetVersionCreatedAtOk() (*time.Time, bool)`

GetVersionCreatedAtOk returns a tuple with the VersionCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionCreatedAt

`func (o *ExtendedVal) SetVersionCreatedAt(v time.Time)`

SetVersionCreatedAt sets VersionCreatedAt field to given value.

### HasVersionCreatedAt

`func (o *ExtendedVal) HasVersionCreatedAt() bool`

HasVersionCreatedAt returns a boolean if a field has been set.

### GetLikeCount

`func (o *ExtendedVal) GetLikeCount() int32`

GetLikeCount returns the LikeCount field if non-nil, zero value otherwise.

### GetLikeCountOk

`func (o *ExtendedVal) GetLikeCountOk() (*int32, bool)`

GetLikeCountOk returns a tuple with the LikeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikeCount

`func (o *ExtendedVal) SetLikeCount(v int32)`

SetLikeCount sets LikeCount field to given value.


### GetReferenceCount

`func (o *ExtendedVal) GetReferenceCount() int32`

GetReferenceCount returns the ReferenceCount field if non-nil, zero value otherwise.

### GetReferenceCountOk

`func (o *ExtendedVal) GetReferenceCountOk() (*int32, bool)`

GetReferenceCountOk returns a tuple with the ReferenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCount

`func (o *ExtendedVal) SetReferenceCount(v int32)`

SetReferenceCount sets ReferenceCount field to given value.


### GetReadme

`func (o *ExtendedVal) GetReadme() string`

GetReadme returns the Readme field if non-nil, zero value otherwise.

### GetReadmeOk

`func (o *ExtendedVal) GetReadmeOk() (*string, bool)`

GetReadmeOk returns a tuple with the Readme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadme

`func (o *ExtendedVal) SetReadme(v string)`

SetReadme sets Readme field to given value.


### SetReadmeNil

`func (o *ExtendedVal) SetReadmeNil(b bool)`

 SetReadmeNil sets the value for Readme to be an explicit nil

### UnsetReadme
`func (o *ExtendedVal) UnsetReadme()`

UnsetReadme ensures that no value is present for Readme, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


