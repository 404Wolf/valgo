/*
Val Town API

Val Town’s public API  OpenAPI JSON endpoint:  https://api.val.town/openapi.json

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package valgo

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the BasicVal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BasicVal{}

// BasicVal A Val
type BasicVal struct {
	// The name of this val
	Name string `json:"name"`
	// This val's id
	Id string `json:"id"`
	// The version of this val, starting at zero
	Version int32 `json:"version"`
	// TypeScript code associated with this val
	Code NullableString `json:"code"`
	// Whether this val is available publicly on Val Town
	Public bool `json:"public"`
	CreatedAt time.Time `json:"createdAt"`
	// This resource's privacy setting. Unlisted resources do not appear on profile pages or elsewhere, but you can link to them.
	Privacy string `json:"privacy"`
	// The type of a val. HTTP can receive web requests, Email can receive emails, Cron runs periodically, and Script can be used for libraries or one-off calculations
	Type string `json:"type"`
	// The URL of this resource on the Val Town website
	Url string `json:"url"`
	Links AliasVal200ResponseLinks `json:"links"`
	Author NullableAliasVal200ResponseAuthor `json:"author"`
}

type _BasicVal BasicVal

// NewBasicVal instantiates a new BasicVal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasicVal(name string, id string, version int32, code NullableString, public bool, createdAt time.Time, privacy string, type_ string, url string, links AliasVal200ResponseLinks, author NullableAliasVal200ResponseAuthor) *BasicVal {
	this := BasicVal{}
	this.Name = name
	this.Id = id
	this.Version = version
	this.Code = code
	this.Public = public
	this.CreatedAt = createdAt
	this.Privacy = privacy
	this.Type = type_
	this.Url = url
	this.Links = links
	this.Author = author
	return &this
}

// NewBasicValWithDefaults instantiates a new BasicVal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasicValWithDefaults() *BasicVal {
	this := BasicVal{}
	return &this
}

// GetName returns the Name field value
func (o *BasicVal) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BasicVal) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BasicVal) SetName(v string) {
	o.Name = v
}

// GetId returns the Id field value
func (o *BasicVal) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BasicVal) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BasicVal) SetId(v string) {
	o.Id = v
}

// GetVersion returns the Version field value
func (o *BasicVal) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *BasicVal) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *BasicVal) SetVersion(v int32) {
	o.Version = v
}

// GetCode returns the Code field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BasicVal) GetCode() string {
	if o == nil || o.Code.Get() == nil {
		var ret string
		return ret
	}

	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BasicVal) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// SetCode sets field value
func (o *BasicVal) SetCode(v string) {
	o.Code.Set(&v)
}

// GetPublic returns the Public field value
func (o *BasicVal) GetPublic() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Public
}

// GetPublicOk returns a tuple with the Public field value
// and a boolean to check if the value has been set.
func (o *BasicVal) GetPublicOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Public, true
}

// SetPublic sets field value
func (o *BasicVal) SetPublic(v bool) {
	o.Public = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *BasicVal) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *BasicVal) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *BasicVal) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetPrivacy returns the Privacy field value
func (o *BasicVal) GetPrivacy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value
// and a boolean to check if the value has been set.
func (o *BasicVal) GetPrivacyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Privacy, true
}

// SetPrivacy sets field value
func (o *BasicVal) SetPrivacy(v string) {
	o.Privacy = v
}

// GetType returns the Type field value
func (o *BasicVal) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BasicVal) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BasicVal) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value
func (o *BasicVal) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BasicVal) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BasicVal) SetUrl(v string) {
	o.Url = v
}

// GetLinks returns the Links field value
func (o *BasicVal) GetLinks() AliasVal200ResponseLinks {
	if o == nil {
		var ret AliasVal200ResponseLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BasicVal) GetLinksOk() (*AliasVal200ResponseLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BasicVal) SetLinks(v AliasVal200ResponseLinks) {
	o.Links = v
}

// GetAuthor returns the Author field value
// If the value is explicit nil, the zero value for AliasVal200ResponseAuthor will be returned
func (o *BasicVal) GetAuthor() AliasVal200ResponseAuthor {
	if o == nil || o.Author.Get() == nil {
		var ret AliasVal200ResponseAuthor
		return ret
	}

	return *o.Author.Get()
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BasicVal) GetAuthorOk() (*AliasVal200ResponseAuthor, bool) {
	if o == nil {
		return nil, false
	}
	return o.Author.Get(), o.Author.IsSet()
}

// SetAuthor sets field value
func (o *BasicVal) SetAuthor(v AliasVal200ResponseAuthor) {
	o.Author.Set(&v)
}

func (o BasicVal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BasicVal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["id"] = o.Id
	toSerialize["version"] = o.Version
	toSerialize["code"] = o.Code.Get()
	toSerialize["public"] = o.Public
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["privacy"] = o.Privacy
	toSerialize["type"] = o.Type
	toSerialize["url"] = o.Url
	toSerialize["links"] = o.Links
	toSerialize["author"] = o.Author.Get()
	return toSerialize, nil
}

func (o *BasicVal) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"id",
		"version",
		"code",
		"public",
		"createdAt",
		"privacy",
		"type",
		"url",
		"links",
		"author",
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

	varBasicVal := _BasicVal{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBasicVal)

	if err != nil {
		return err
	}

	*o = BasicVal(varBasicVal)

	return err
}

type NullableBasicVal struct {
	value *BasicVal
	isSet bool
}

func (v NullableBasicVal) Get() *BasicVal {
	return v.value
}

func (v *NullableBasicVal) Set(val *BasicVal) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicVal) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicVal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicVal(val *BasicVal) *NullableBasicVal {
	return &NullableBasicVal{value: val, isSet: true}
}

func (v NullableBasicVal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicVal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


