# ValsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | This val’s name | [optional] 
**Readme** | Pointer to **string** | Readme contents, as Markdown | [optional] 
**Privacy** | Pointer to **string** | This resource&#39;s privacy setting. Unlisted resources do not appear on profile pages or elsewhere, but you can link to them. | [optional] 
**Type** | Pointer to **string** | The type of the val you want to update. Note that this does not include interval vals, because they cannot be created through the API yet. | [optional] 

## Methods

### NewValsUpdateRequest

`func NewValsUpdateRequest() *ValsUpdateRequest`

NewValsUpdateRequest instantiates a new ValsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValsUpdateRequestWithDefaults

`func NewValsUpdateRequestWithDefaults() *ValsUpdateRequest`

NewValsUpdateRequestWithDefaults instantiates a new ValsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ValsUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ValsUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ValsUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ValsUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReadme

`func (o *ValsUpdateRequest) GetReadme() string`

GetReadme returns the Readme field if non-nil, zero value otherwise.

### GetReadmeOk

`func (o *ValsUpdateRequest) GetReadmeOk() (*string, bool)`

GetReadmeOk returns a tuple with the Readme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadme

`func (o *ValsUpdateRequest) SetReadme(v string)`

SetReadme sets Readme field to given value.

### HasReadme

`func (o *ValsUpdateRequest) HasReadme() bool`

HasReadme returns a boolean if a field has been set.

### GetPrivacy

`func (o *ValsUpdateRequest) GetPrivacy() string`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *ValsUpdateRequest) GetPrivacyOk() (*string, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *ValsUpdateRequest) SetPrivacy(v string)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *ValsUpdateRequest) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetType

`func (o *ValsUpdateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ValsUpdateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ValsUpdateRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ValsUpdateRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


