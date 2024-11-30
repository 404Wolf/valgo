# ValsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Val source code as TypeScript | 
**Name** | Pointer to **string** | This val’s name | [optional] 
**Readme** | Pointer to **string** | Readme contents, as Markdown | [optional] 
**Privacy** | Pointer to **string** | This val’s privacy setting. Unlisted vals do not appear on profile pages or elsewhere, but you can link to them. | [optional] 
**Type** | Pointer to **string** | The type of the val you want to create. Note that this does not include interval vals, because they cannot be created through the API yet. | [optional] [default to "script"]

## Methods

### NewValsCreateRequest

`func NewValsCreateRequest(code string, ) *ValsCreateRequest`

NewValsCreateRequest instantiates a new ValsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValsCreateRequestWithDefaults

`func NewValsCreateRequestWithDefaults() *ValsCreateRequest`

NewValsCreateRequestWithDefaults instantiates a new ValsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ValsCreateRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ValsCreateRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ValsCreateRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *ValsCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ValsCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ValsCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ValsCreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReadme

`func (o *ValsCreateRequest) GetReadme() string`

GetReadme returns the Readme field if non-nil, zero value otherwise.

### GetReadmeOk

`func (o *ValsCreateRequest) GetReadmeOk() (*string, bool)`

GetReadmeOk returns a tuple with the Readme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadme

`func (o *ValsCreateRequest) SetReadme(v string)`

SetReadme sets Readme field to given value.

### HasReadme

`func (o *ValsCreateRequest) HasReadme() bool`

HasReadme returns a boolean if a field has been set.

### GetPrivacy

`func (o *ValsCreateRequest) GetPrivacy() string`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *ValsCreateRequest) GetPrivacyOk() (*string, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *ValsCreateRequest) SetPrivacy(v string)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *ValsCreateRequest) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetType

`func (o *ValsCreateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ValsCreateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ValsCreateRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ValsCreateRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


