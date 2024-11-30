/*
Val Town API

Val Town’s public API  OpenAPI JSON endpoint:  https://api.val.town/openapi.json

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package valgo

import (
	"encoding/json"
)

// checks if the EmailsSendRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailsSendRequest{}

// EmailsSendRequest Fields for an email to be sent
type EmailsSendRequest struct {
	// The subject line of the email
	Subject *string `json:"subject,omitempty"`
	From *EmailData `json:"from,omitempty"`
	// A set of headers to include the email that you send
	Headers map[string]string `json:"headers,omitempty"`
	To *EmailDataInput `json:"to,omitempty"`
	Cc *EmailDataInput `json:"cc,omitempty"`
	Bcc *EmailDataInput `json:"bcc,omitempty"`
	// Text content of the email, for email clients that may not support HTML
	Text *string `json:"text,omitempty"`
	// HTML content of the email. Can be specified alongside text
	Html *string `json:"html,omitempty"`
	// A list of attachments to add to the email
	Attachments []AttachmentObject `json:"attachments,omitempty"`
	ReplyToList *ReplyToList `json:"replyToList,omitempty"`
}

// NewEmailsSendRequest instantiates a new EmailsSendRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailsSendRequest() *EmailsSendRequest {
	this := EmailsSendRequest{}
	return &this
}

// NewEmailsSendRequestWithDefaults instantiates a new EmailsSendRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailsSendRequestWithDefaults() *EmailsSendRequest {
	this := EmailsSendRequest{}
	return &this
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *EmailsSendRequest) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailsSendRequest) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *EmailsSendRequest) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *EmailsSendRequest) SetSubject(v string) {
	o.Subject = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *EmailsSendRequest) GetFrom() EmailData {
	if o == nil || IsNil(o.From) {
		var ret EmailData
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailsSendRequest) GetFromOk() (*EmailData, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *EmailsSendRequest) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given EmailData and assigns it to the From field.
func (o *EmailsSendRequest) SetFrom(v EmailData) {
	o.From = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *EmailsSendRequest) GetHeaders() map[string]string {
	if o == nil || IsNil(o.Headers) {
		var ret map[string]string
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailsSendRequest) GetHeadersOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Headers) {
		return map[string]string{}, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *EmailsSendRequest) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *EmailsSendRequest) SetHeaders(v map[string]string) {
	o.Headers = v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *EmailsSendRequest) GetTo() EmailDataInput {
	if o == nil || IsNil(o.To) {
		var ret EmailDataInput
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailsSendRequest) GetToOk() (*EmailDataInput, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *EmailsSendRequest) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given EmailDataInput and assigns it to the To field.
func (o *EmailsSendRequest) SetTo(v EmailDataInput) {
	o.To = &v
}

// GetCc returns the Cc field value if set, zero value otherwise.
func (o *EmailsSendRequest) GetCc() EmailDataInput {
	if o == nil || IsNil(o.Cc) {
		var ret EmailDataInput
		return ret
	}
	return *o.Cc
}

// GetCcOk returns a tuple with the Cc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailsSendRequest) GetCcOk() (*EmailDataInput, bool) {
	if o == nil || IsNil(o.Cc) {
		return nil, false
	}
	return o.Cc, true
}

// HasCc returns a boolean if a field has been set.
func (o *EmailsSendRequest) HasCc() bool {
	if o != nil && !IsNil(o.Cc) {
		return true
	}

	return false
}

// SetCc gets a reference to the given EmailDataInput and assigns it to the Cc field.
func (o *EmailsSendRequest) SetCc(v EmailDataInput) {
	o.Cc = &v
}

// GetBcc returns the Bcc field value if set, zero value otherwise.
func (o *EmailsSendRequest) GetBcc() EmailDataInput {
	if o == nil || IsNil(o.Bcc) {
		var ret EmailDataInput
		return ret
	}
	return *o.Bcc
}

// GetBccOk returns a tuple with the Bcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailsSendRequest) GetBccOk() (*EmailDataInput, bool) {
	if o == nil || IsNil(o.Bcc) {
		return nil, false
	}
	return o.Bcc, true
}

// HasBcc returns a boolean if a field has been set.
func (o *EmailsSendRequest) HasBcc() bool {
	if o != nil && !IsNil(o.Bcc) {
		return true
	}

	return false
}

// SetBcc gets a reference to the given EmailDataInput and assigns it to the Bcc field.
func (o *EmailsSendRequest) SetBcc(v EmailDataInput) {
	o.Bcc = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *EmailsSendRequest) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailsSendRequest) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *EmailsSendRequest) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *EmailsSendRequest) SetText(v string) {
	o.Text = &v
}

// GetHtml returns the Html field value if set, zero value otherwise.
func (o *EmailsSendRequest) GetHtml() string {
	if o == nil || IsNil(o.Html) {
		var ret string
		return ret
	}
	return *o.Html
}

// GetHtmlOk returns a tuple with the Html field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailsSendRequest) GetHtmlOk() (*string, bool) {
	if o == nil || IsNil(o.Html) {
		return nil, false
	}
	return o.Html, true
}

// HasHtml returns a boolean if a field has been set.
func (o *EmailsSendRequest) HasHtml() bool {
	if o != nil && !IsNil(o.Html) {
		return true
	}

	return false
}

// SetHtml gets a reference to the given string and assigns it to the Html field.
func (o *EmailsSendRequest) SetHtml(v string) {
	o.Html = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *EmailsSendRequest) GetAttachments() []AttachmentObject {
	if o == nil || IsNil(o.Attachments) {
		var ret []AttachmentObject
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailsSendRequest) GetAttachmentsOk() ([]AttachmentObject, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *EmailsSendRequest) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []AttachmentObject and assigns it to the Attachments field.
func (o *EmailsSendRequest) SetAttachments(v []AttachmentObject) {
	o.Attachments = v
}

// GetReplyToList returns the ReplyToList field value if set, zero value otherwise.
func (o *EmailsSendRequest) GetReplyToList() ReplyToList {
	if o == nil || IsNil(o.ReplyToList) {
		var ret ReplyToList
		return ret
	}
	return *o.ReplyToList
}

// GetReplyToListOk returns a tuple with the ReplyToList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailsSendRequest) GetReplyToListOk() (*ReplyToList, bool) {
	if o == nil || IsNil(o.ReplyToList) {
		return nil, false
	}
	return o.ReplyToList, true
}

// HasReplyToList returns a boolean if a field has been set.
func (o *EmailsSendRequest) HasReplyToList() bool {
	if o != nil && !IsNil(o.ReplyToList) {
		return true
	}

	return false
}

// SetReplyToList gets a reference to the given ReplyToList and assigns it to the ReplyToList field.
func (o *EmailsSendRequest) SetReplyToList(v ReplyToList) {
	o.ReplyToList = &v
}

func (o EmailsSendRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailsSendRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	if !IsNil(o.Cc) {
		toSerialize["cc"] = o.Cc
	}
	if !IsNil(o.Bcc) {
		toSerialize["bcc"] = o.Bcc
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Html) {
		toSerialize["html"] = o.Html
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.ReplyToList) {
		toSerialize["replyToList"] = o.ReplyToList
	}
	return toSerialize, nil
}

type NullableEmailsSendRequest struct {
	value *EmailsSendRequest
	isSet bool
}

func (v NullableEmailsSendRequest) Get() *EmailsSendRequest {
	return v.value
}

func (v *NullableEmailsSendRequest) Set(val *EmailsSendRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailsSendRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailsSendRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailsSendRequest(val *EmailsSendRequest) *NullableEmailsSendRequest {
	return &NullableEmailsSendRequest{value: val, isSet: true}
}

func (v NullableEmailsSendRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailsSendRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

