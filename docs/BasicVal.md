# BasicVal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of this val | 
**Id** | **string** | This val’s id | 
**Version** | **int32** | The version of this val, starting at zero | 
**Code** | **NullableString** | TypeScript code associated with this val | 
**Public** | **bool** | Whether this val is available publicly on Val Town | 
**CreatedAt** | **time.Time** |  | 
**Privacy** | **string** | This resource&#39;s privacy setting. Unlisted resources do not appear on profile pages or elsewhere, but you can link to them. | 
**Type** | **string** | The type of a val. HTTP can receive web requests, Email can receive emails, Cron runs periodically, and Script can be used for libraries or one-off calculations | 
**Url** | **string** | The URL of this resource on the Val Town website | 
**Links** | [**AliasVal200ResponseLinks**](AliasVal200ResponseLinks.md) |  | 
**Author** | [**NullableAliasVal200ResponseAuthor**](AliasVal200ResponseAuthor.md) |  | 

## Methods

### NewBasicVal

`func NewBasicVal(name string, id string, version int32, code NullableString, public bool, createdAt time.Time, privacy string, type_ string, url string, links AliasVal200ResponseLinks, author NullableAliasVal200ResponseAuthor, ) *BasicVal`

NewBasicVal instantiates a new BasicVal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicValWithDefaults

`func NewBasicValWithDefaults() *BasicVal`

NewBasicValWithDefaults instantiates a new BasicVal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BasicVal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BasicVal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BasicVal) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *BasicVal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BasicVal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BasicVal) SetId(v string)`

SetId sets Id field to given value.


### GetVersion

`func (o *BasicVal) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BasicVal) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BasicVal) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetCode

`func (o *BasicVal) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BasicVal) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BasicVal) SetCode(v string)`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *BasicVal) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *BasicVal) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetPublic

`func (o *BasicVal) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *BasicVal) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *BasicVal) SetPublic(v bool)`

SetPublic sets Public field to given value.


### GetCreatedAt

`func (o *BasicVal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BasicVal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BasicVal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetPrivacy

`func (o *BasicVal) GetPrivacy() string`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *BasicVal) GetPrivacyOk() (*string, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *BasicVal) SetPrivacy(v string)`

SetPrivacy sets Privacy field to given value.


### GetType

`func (o *BasicVal) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BasicVal) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BasicVal) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *BasicVal) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BasicVal) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BasicVal) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetLinks

`func (o *BasicVal) GetLinks() AliasVal200ResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BasicVal) GetLinksOk() (*AliasVal200ResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BasicVal) SetLinks(v AliasVal200ResponseLinks)`

SetLinks sets Links field to given value.


### GetAuthor

`func (o *BasicVal) GetAuthor() AliasVal200ResponseAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *BasicVal) GetAuthorOk() (*AliasVal200ResponseAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *BasicVal) SetAuthor(v AliasVal200ResponseAuthor)`

SetAuthor sets Author field to given value.


### SetAuthorNil

`func (o *BasicVal) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *BasicVal) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


