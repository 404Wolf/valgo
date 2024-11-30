# AliasVal200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of this val | 
**Id** | **string** | This val’s id | 
**Version** | **int32** | The version of this val, starting at zero | 
**Code** | **NullableString** | TypeScript code associated with this val | 
**Public** | **bool** | Whether this val is available publicly on Val Town | 
**CreatedAt** | **time.Time** |  | 
**Privacy** | **string** | This val’s privacy setting. Unlisted vals do not appear on profile pages or elsewhere, but you can link to them. | 
**Type** | **string** | The type of a val. HTTP can receive web requests, Email can receive emails, Cron runs periodically, and Script can be used for libraries or one-off calculations | 
**Url** | **string** | The URL of this val on the Val Town website | 
**Links** | [**AliasVal200ResponseLinks**](AliasVal200ResponseLinks.md) |  | 
**Author** | [**NullableAliasVal200ResponseAuthor**](AliasVal200ResponseAuthor.md) |  | 
**VersionCreatedAt** | Pointer to **time.Time** |  | [optional] 
**LikeCount** | **int32** | How many likes this val has received | 
**ReferenceCount** | **int32** |  | 
**Readme** | **NullableString** | This val’s readme, as Markdown | 

## Methods

### NewAliasVal200Response

`func NewAliasVal200Response(name string, id string, version int32, code NullableString, public bool, createdAt time.Time, privacy string, type_ string, url string, links AliasVal200ResponseLinks, author NullableAliasVal200ResponseAuthor, likeCount int32, referenceCount int32, readme NullableString, ) *AliasVal200Response`

NewAliasVal200Response instantiates a new AliasVal200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAliasVal200ResponseWithDefaults

`func NewAliasVal200ResponseWithDefaults() *AliasVal200Response`

NewAliasVal200ResponseWithDefaults instantiates a new AliasVal200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AliasVal200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AliasVal200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AliasVal200Response) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *AliasVal200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AliasVal200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AliasVal200Response) SetId(v string)`

SetId sets Id field to given value.


### GetVersion

`func (o *AliasVal200Response) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AliasVal200Response) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AliasVal200Response) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetCode

`func (o *AliasVal200Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AliasVal200Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AliasVal200Response) SetCode(v string)`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *AliasVal200Response) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *AliasVal200Response) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetPublic

`func (o *AliasVal200Response) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *AliasVal200Response) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *AliasVal200Response) SetPublic(v bool)`

SetPublic sets Public field to given value.


### GetCreatedAt

`func (o *AliasVal200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AliasVal200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AliasVal200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetPrivacy

`func (o *AliasVal200Response) GetPrivacy() string`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *AliasVal200Response) GetPrivacyOk() (*string, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *AliasVal200Response) SetPrivacy(v string)`

SetPrivacy sets Privacy field to given value.


### GetType

`func (o *AliasVal200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AliasVal200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AliasVal200Response) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *AliasVal200Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AliasVal200Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AliasVal200Response) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetLinks

`func (o *AliasVal200Response) GetLinks() AliasVal200ResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AliasVal200Response) GetLinksOk() (*AliasVal200ResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AliasVal200Response) SetLinks(v AliasVal200ResponseLinks)`

SetLinks sets Links field to given value.


### GetAuthor

`func (o *AliasVal200Response) GetAuthor() AliasVal200ResponseAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *AliasVal200Response) GetAuthorOk() (*AliasVal200ResponseAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *AliasVal200Response) SetAuthor(v AliasVal200ResponseAuthor)`

SetAuthor sets Author field to given value.


### SetAuthorNil

`func (o *AliasVal200Response) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *AliasVal200Response) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetVersionCreatedAt

`func (o *AliasVal200Response) GetVersionCreatedAt() time.Time`

GetVersionCreatedAt returns the VersionCreatedAt field if non-nil, zero value otherwise.

### GetVersionCreatedAtOk

`func (o *AliasVal200Response) GetVersionCreatedAtOk() (*time.Time, bool)`

GetVersionCreatedAtOk returns a tuple with the VersionCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionCreatedAt

`func (o *AliasVal200Response) SetVersionCreatedAt(v time.Time)`

SetVersionCreatedAt sets VersionCreatedAt field to given value.

### HasVersionCreatedAt

`func (o *AliasVal200Response) HasVersionCreatedAt() bool`

HasVersionCreatedAt returns a boolean if a field has been set.

### GetLikeCount

`func (o *AliasVal200Response) GetLikeCount() int32`

GetLikeCount returns the LikeCount field if non-nil, zero value otherwise.

### GetLikeCountOk

`func (o *AliasVal200Response) GetLikeCountOk() (*int32, bool)`

GetLikeCountOk returns a tuple with the LikeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikeCount

`func (o *AliasVal200Response) SetLikeCount(v int32)`

SetLikeCount sets LikeCount field to given value.


### GetReferenceCount

`func (o *AliasVal200Response) GetReferenceCount() int32`

GetReferenceCount returns the ReferenceCount field if non-nil, zero value otherwise.

### GetReferenceCountOk

`func (o *AliasVal200Response) GetReferenceCountOk() (*int32, bool)`

GetReferenceCountOk returns a tuple with the ReferenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCount

`func (o *AliasVal200Response) SetReferenceCount(v int32)`

SetReferenceCount sets ReferenceCount field to given value.


### GetReadme

`func (o *AliasVal200Response) GetReadme() string`

GetReadme returns the Readme field if non-nil, zero value otherwise.

### GetReadmeOk

`func (o *AliasVal200Response) GetReadmeOk() (*string, bool)`

GetReadmeOk returns a tuple with the Readme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadme

`func (o *AliasVal200Response) SetReadme(v string)`

SetReadme sets Readme field to given value.


### SetReadmeNil

`func (o *AliasVal200Response) SetReadmeNil(b bool)`

 SetReadmeNil sets the value for Readme to be an explicit nil

### UnsetReadme
`func (o *AliasVal200Response) UnsetReadme()`

UnsetReadme ensures that no value is present for Readme, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


