# AliasVal200ResponseLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | **string** | The URL of this val on this API | 
**Versions** | **string** | The endpoint to retrieve this val&#39;s versions | 
**Module** | **string** | The URL of this Val&#39;s source code as a module | 
**Endpoint** | Pointer to **string** | This val&#39;s web endpoint, where it serves a website or API | [optional] 

## Methods

### NewAliasVal200ResponseLinks

`func NewAliasVal200ResponseLinks(self string, versions string, module string, ) *AliasVal200ResponseLinks`

NewAliasVal200ResponseLinks instantiates a new AliasVal200ResponseLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAliasVal200ResponseLinksWithDefaults

`func NewAliasVal200ResponseLinksWithDefaults() *AliasVal200ResponseLinks`

NewAliasVal200ResponseLinksWithDefaults instantiates a new AliasVal200ResponseLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *AliasVal200ResponseLinks) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *AliasVal200ResponseLinks) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *AliasVal200ResponseLinks) SetSelf(v string)`

SetSelf sets Self field to given value.


### GetVersions

`func (o *AliasVal200ResponseLinks) GetVersions() string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *AliasVal200ResponseLinks) GetVersionsOk() (*string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *AliasVal200ResponseLinks) SetVersions(v string)`

SetVersions sets Versions field to given value.


### GetModule

`func (o *AliasVal200ResponseLinks) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *AliasVal200ResponseLinks) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *AliasVal200ResponseLinks) SetModule(v string)`

SetModule sets Module field to given value.


### GetEndpoint

`func (o *AliasVal200ResponseLinks) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *AliasVal200ResponseLinks) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *AliasVal200ResponseLinks) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *AliasVal200ResponseLinks) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


