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

// checks if the Branch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Branch{}

// Branch A Branch
type Branch struct {
	Name string `json:"name"`
	// The id of the branch
	Id string `json:"id"`
	Version int32 `json:"version"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	// The id of the branch this branch was forked from
	ForkedBranchId NullableString `json:"forkedBranchId"`
	Links ProjectLinks `json:"links"`
}

type _Branch Branch

// NewBranch instantiates a new Branch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBranch(name string, id string, version int32, createdAt time.Time, updatedAt time.Time, forkedBranchId NullableString, links ProjectLinks) *Branch {
	this := Branch{}
	this.Name = name
	this.Id = id
	this.Version = version
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.ForkedBranchId = forkedBranchId
	this.Links = links
	return &this
}

// NewBranchWithDefaults instantiates a new Branch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBranchWithDefaults() *Branch {
	this := Branch{}
	return &this
}

// GetName returns the Name field value
func (o *Branch) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Branch) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Branch) SetName(v string) {
	o.Name = v
}

// GetId returns the Id field value
func (o *Branch) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Branch) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Branch) SetId(v string) {
	o.Id = v
}

// GetVersion returns the Version field value
func (o *Branch) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Branch) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *Branch) SetVersion(v int32) {
	o.Version = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Branch) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Branch) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Branch) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Branch) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Branch) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Branch) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetForkedBranchId returns the ForkedBranchId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Branch) GetForkedBranchId() string {
	if o == nil || o.ForkedBranchId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ForkedBranchId.Get()
}

// GetForkedBranchIdOk returns a tuple with the ForkedBranchId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Branch) GetForkedBranchIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ForkedBranchId.Get(), o.ForkedBranchId.IsSet()
}

// SetForkedBranchId sets field value
func (o *Branch) SetForkedBranchId(v string) {
	o.ForkedBranchId.Set(&v)
}

// GetLinks returns the Links field value
func (o *Branch) GetLinks() ProjectLinks {
	if o == nil {
		var ret ProjectLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *Branch) GetLinksOk() (*ProjectLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *Branch) SetLinks(v ProjectLinks) {
	o.Links = v
}

func (o Branch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Branch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["id"] = o.Id
	toSerialize["version"] = o.Version
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["forkedBranchId"] = o.ForkedBranchId.Get()
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *Branch) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"id",
		"version",
		"createdAt",
		"updatedAt",
		"forkedBranchId",
		"links",
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

	varBranch := _Branch{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBranch)

	if err != nil {
		return err
	}

	*o = Branch(varBranch)

	return err
}

type NullableBranch struct {
	value *Branch
	isSet bool
}

func (v NullableBranch) Get() *Branch {
	return v.value
}

func (v *NullableBranch) Set(val *Branch) {
	v.value = val
	v.isSet = true
}

func (v NullableBranch) IsSet() bool {
	return v.isSet
}

func (v *NullableBranch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBranch(val *Branch) *NullableBranch {
	return &NullableBranch{value: val, isSet: true}
}

func (v NullableBranch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBranch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


