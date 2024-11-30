# EmailNameAndAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Email** | **string** |  | 

## Methods

### NewEmailNameAndAddress

`func NewEmailNameAndAddress(email string, ) *EmailNameAndAddress`

NewEmailNameAndAddress instantiates a new EmailNameAndAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailNameAndAddressWithDefaults

`func NewEmailNameAndAddressWithDefaults() *EmailNameAndAddress`

NewEmailNameAndAddressWithDefaults instantiates a new EmailNameAndAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EmailNameAndAddress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailNameAndAddress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailNameAndAddress) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EmailNameAndAddress) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *EmailNameAndAddress) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmailNameAndAddress) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmailNameAndAddress) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


