# EmailsSendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to **string** | The subject line of the email | [optional] 
**From** | Pointer to [**EmailData**](EmailData.md) |  | [optional] 
**Headers** | Pointer to **map[string]string** | A set of headers to include the email that you send | [optional] 
**To** | Pointer to [**EmailDataInput**](EmailDataInput.md) |  | [optional] 
**Cc** | Pointer to [**EmailDataInput**](EmailDataInput.md) |  | [optional] 
**Bcc** | Pointer to [**EmailDataInput**](EmailDataInput.md) |  | [optional] 
**Text** | Pointer to **string** | Text content of the email, for email clients that may not support HTML | [optional] 
**Html** | Pointer to **string** | HTML content of the email. Can be specified alongside text | [optional] 
**Attachments** | Pointer to [**[]AttachmentObject**](AttachmentObject.md) | A list of attachments to add to the email | [optional] 
**ReplyToList** | Pointer to [**ReplyToList**](ReplyToList.md) |  | [optional] 

## Methods

### NewEmailsSendRequest

`func NewEmailsSendRequest() *EmailsSendRequest`

NewEmailsSendRequest instantiates a new EmailsSendRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailsSendRequestWithDefaults

`func NewEmailsSendRequestWithDefaults() *EmailsSendRequest`

NewEmailsSendRequestWithDefaults instantiates a new EmailsSendRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *EmailsSendRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailsSendRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailsSendRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *EmailsSendRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetFrom

`func (o *EmailsSendRequest) GetFrom() EmailData`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EmailsSendRequest) GetFromOk() (*EmailData, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EmailsSendRequest) SetFrom(v EmailData)`

SetFrom sets From field to given value.

### HasFrom

`func (o *EmailsSendRequest) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetHeaders

`func (o *EmailsSendRequest) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *EmailsSendRequest) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *EmailsSendRequest) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *EmailsSendRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetTo

`func (o *EmailsSendRequest) GetTo() EmailDataInput`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EmailsSendRequest) GetToOk() (*EmailDataInput, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EmailsSendRequest) SetTo(v EmailDataInput)`

SetTo sets To field to given value.

### HasTo

`func (o *EmailsSendRequest) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetCc

`func (o *EmailsSendRequest) GetCc() EmailDataInput`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *EmailsSendRequest) GetCcOk() (*EmailDataInput, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *EmailsSendRequest) SetCc(v EmailDataInput)`

SetCc sets Cc field to given value.

### HasCc

`func (o *EmailsSendRequest) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetBcc

`func (o *EmailsSendRequest) GetBcc() EmailDataInput`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *EmailsSendRequest) GetBccOk() (*EmailDataInput, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *EmailsSendRequest) SetBcc(v EmailDataInput)`

SetBcc sets Bcc field to given value.

### HasBcc

`func (o *EmailsSendRequest) HasBcc() bool`

HasBcc returns a boolean if a field has been set.

### GetText

`func (o *EmailsSendRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *EmailsSendRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *EmailsSendRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *EmailsSendRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetHtml

`func (o *EmailsSendRequest) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *EmailsSendRequest) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *EmailsSendRequest) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *EmailsSendRequest) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetAttachments

`func (o *EmailsSendRequest) GetAttachments() []AttachmentObject`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *EmailsSendRequest) GetAttachmentsOk() (*[]AttachmentObject, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *EmailsSendRequest) SetAttachments(v []AttachmentObject)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *EmailsSendRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetReplyToList

`func (o *EmailsSendRequest) GetReplyToList() ReplyToList`

GetReplyToList returns the ReplyToList field if non-nil, zero value otherwise.

### GetReplyToListOk

`func (o *EmailsSendRequest) GetReplyToListOk() (*ReplyToList, bool)`

GetReplyToListOk returns a tuple with the ReplyToList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyToList

`func (o *EmailsSendRequest) SetReplyToList(v ReplyToList)`

SetReplyToList sets ReplyToList field to given value.

### HasReplyToList

`func (o *EmailsSendRequest) HasReplyToList() bool`

HasReplyToList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


