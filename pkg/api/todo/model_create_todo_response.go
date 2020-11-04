/*
 * Todo API
 *
 * Todo API
 *
 * API version: 0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package todo

import (
	"encoding/json"
	"github.com/google/uuid"
)

// CreateTodoResponse struct for CreateTodoResponse
type CreateTodoResponse struct {
	Id uuid.UUID `json:"id"`
}

// NewCreateTodoResponse instantiates a new CreateTodoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTodoResponse(id uuid.UUID, ) *CreateTodoResponse {
	this := CreateTodoResponse{}
	this.Id = id
	return &this
}

// NewCreateTodoResponseWithDefaults instantiates a new CreateTodoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTodoResponseWithDefaults() *CreateTodoResponse {
	this := CreateTodoResponse{}
	return &this
}

// GetId returns the Id field value
func (o *CreateTodoResponse) GetId() uuid.UUID {
	if o == nil  {
		var ret uuid.UUID
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateTodoResponse) GetIdOk() (*uuid.UUID, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateTodoResponse) SetId(v uuid.UUID) {
	o.Id = v
}

func (o CreateTodoResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableCreateTodoResponse struct {
	value *CreateTodoResponse
	isSet bool
}

func (v NullableCreateTodoResponse) Get() *CreateTodoResponse {
	return v.value
}

func (v *NullableCreateTodoResponse) Set(val *CreateTodoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTodoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTodoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTodoResponse(val *CreateTodoResponse) *NullableCreateTodoResponse {
	return &NullableCreateTodoResponse{value: val, isSet: true}
}

func (v NullableCreateTodoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTodoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
