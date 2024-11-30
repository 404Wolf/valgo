/*
Val Town API

Val Town’s public API  OpenAPI JSON endpoint:  https://api.val.town/openapi.json

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the MeComments200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MeComments200ResponseDataInner{}

// MeComments200ResponseDataInner struct for MeComments200ResponseDataInner
type MeComments200ResponseDataInner struct {
	// Text of the given comment, in Markdown
	Comment string `json:"comment"`
	// The comment’s id
	Id string `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Author MeComments200ResponseDataInnerAuthor `json:"author"`
	Val MeComments200ResponseDataInnerVal `json:"val"`
}

type _MeComments200ResponseDataInner MeComments200ResponseDataInner

// NewMeComments200ResponseDataInner instantiates a new MeComments200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeComments200ResponseDataInner(comment string, id string, createdAt time.Time, author MeComments200ResponseDataInnerAuthor, val MeComments200ResponseDataInnerVal) *MeComments200ResponseDataInner {
	this := MeComments200ResponseDataInner{}
	this.Comment = comment
	this.Id = id
	this.CreatedAt = createdAt
	this.Author = author
	this.Val = val
	return &this
}

// NewMeComments200ResponseDataInnerWithDefaults instantiates a new MeComments200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeComments200ResponseDataInnerWithDefaults() *MeComments200ResponseDataInner {
	this := MeComments200ResponseDataInner{}
	return &this
}

// GetComment returns the Comment field value
func (o *MeComments200ResponseDataInner) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *MeComments200ResponseDataInner) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *MeComments200ResponseDataInner) SetComment(v string) {
	o.Comment = v
}

// GetId returns the Id field value
func (o *MeComments200ResponseDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MeComments200ResponseDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MeComments200ResponseDataInner) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *MeComments200ResponseDataInner) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *MeComments200ResponseDataInner) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *MeComments200ResponseDataInner) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetAuthor returns the Author field value
func (o *MeComments200ResponseDataInner) GetAuthor() MeComments200ResponseDataInnerAuthor {
	if o == nil {
		var ret MeComments200ResponseDataInnerAuthor
		return ret
	}

	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
func (o *MeComments200ResponseDataInner) GetAuthorOk() (*MeComments200ResponseDataInnerAuthor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Author, true
}

// SetAuthor sets field value
func (o *MeComments200ResponseDataInner) SetAuthor(v MeComments200ResponseDataInnerAuthor) {
	o.Author = v
}

// GetVal returns the Val field value
func (o *MeComments200ResponseDataInner) GetVal() MeComments200ResponseDataInnerVal {
	if o == nil {
		var ret MeComments200ResponseDataInnerVal
		return ret
	}

	return o.Val
}

// GetValOk returns a tuple with the Val field value
// and a boolean to check if the value has been set.
func (o *MeComments200ResponseDataInner) GetValOk() (*MeComments200ResponseDataInnerVal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Val, true
}

// SetVal sets field value
func (o *MeComments200ResponseDataInner) SetVal(v MeComments200ResponseDataInnerVal) {
	o.Val = v
}

func (o MeComments200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MeComments200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["comment"] = o.Comment
	toSerialize["id"] = o.Id
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["author"] = o.Author
	toSerialize["val"] = o.Val
	return toSerialize, nil
}

func (o *MeComments200ResponseDataInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"comment",
		"id",
		"createdAt",
		"author",
		"val",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMeComments200ResponseDataInner := _MeComments200ResponseDataInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMeComments200ResponseDataInner)

	if err != nil {
		return err
	}

	*o = MeComments200ResponseDataInner(varMeComments200ResponseDataInner)

	return err
}

type NullableMeComments200ResponseDataInner struct {
	value *MeComments200ResponseDataInner
	isSet bool
}

func (v NullableMeComments200ResponseDataInner) Get() *MeComments200ResponseDataInner {
	return v.value
}

func (v *NullableMeComments200ResponseDataInner) Set(val *MeComments200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMeComments200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMeComments200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeComments200ResponseDataInner(val *MeComments200ResponseDataInner) *NullableMeComments200ResponseDataInner {
	return &NullableMeComments200ResponseDataInner{value: val, isSet: true}
}

func (v NullableMeComments200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeComments200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


