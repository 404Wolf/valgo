# MeGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of this user | 
**Bio** | **NullableString** | The user’s biography, if they have provided one | 
**Username** | **NullableString** | The user’s handle that they chose for themselves. Does not include the @ symbol | 
**ProfileImageUrl** | **NullableString** | URL that points to the user’s profile image, if one exists | 
**Url** | **string** | URL of this user’s profile on Val Town’s website | 
**Links** | [**MeGet200ResponseLinks**](MeGet200ResponseLinks.md) |  | 
**Tier** | **NullableString** | Your account tier | 
**Email** | **string** | Your email address | 

## Methods

### NewMeGet200Response

`func NewMeGet200Response(id string, bio NullableString, username NullableString, profileImageUrl NullableString, url string, links MeGet200ResponseLinks, tier NullableString, email string, ) *MeGet200Response`

NewMeGet200Response instantiates a new MeGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeGet200ResponseWithDefaults

`func NewMeGet200ResponseWithDefaults() *MeGet200Response`

NewMeGet200ResponseWithDefaults instantiates a new MeGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MeGet200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MeGet200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MeGet200Response) SetId(v string)`

SetId sets Id field to given value.


### GetBio

`func (o *MeGet200Response) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *MeGet200Response) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *MeGet200Response) SetBio(v string)`

SetBio sets Bio field to given value.


### SetBioNil

`func (o *MeGet200Response) SetBioNil(b bool)`

 SetBioNil sets the value for Bio to be an explicit nil

### UnsetBio
`func (o *MeGet200Response) UnsetBio()`

UnsetBio ensures that no value is present for Bio, not even an explicit nil
### GetUsername

`func (o *MeGet200Response) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *MeGet200Response) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *MeGet200Response) SetUsername(v string)`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *MeGet200Response) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *MeGet200Response) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetProfileImageUrl

`func (o *MeGet200Response) GetProfileImageUrl() string`

GetProfileImageUrl returns the ProfileImageUrl field if non-nil, zero value otherwise.

### GetProfileImageUrlOk

`func (o *MeGet200Response) GetProfileImageUrlOk() (*string, bool)`

GetProfileImageUrlOk returns a tuple with the ProfileImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileImageUrl

`func (o *MeGet200Response) SetProfileImageUrl(v string)`

SetProfileImageUrl sets ProfileImageUrl field to given value.


### SetProfileImageUrlNil

`func (o *MeGet200Response) SetProfileImageUrlNil(b bool)`

 SetProfileImageUrlNil sets the value for ProfileImageUrl to be an explicit nil

### UnsetProfileImageUrl
`func (o *MeGet200Response) UnsetProfileImageUrl()`

UnsetProfileImageUrl ensures that no value is present for ProfileImageUrl, not even an explicit nil
### GetUrl

`func (o *MeGet200Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MeGet200Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MeGet200Response) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetLinks

`func (o *MeGet200Response) GetLinks() MeGet200ResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MeGet200Response) GetLinksOk() (*MeGet200ResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MeGet200Response) SetLinks(v MeGet200ResponseLinks)`

SetLinks sets Links field to given value.


### GetTier

`func (o *MeGet200Response) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *MeGet200Response) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *MeGet200Response) SetTier(v string)`

SetTier sets Tier field to given value.


### SetTierNil

`func (o *MeGet200Response) SetTierNil(b bool)`

 SetTierNil sets the value for Tier to be an explicit nil

### UnsetTier
`func (o *MeGet200Response) UnsetTier()`

UnsetTier ensures that no value is present for Tier, not even an explicit nil
### GetEmail

`func (o *MeGet200Response) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MeGet200Response) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MeGet200Response) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


