/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstore

import (
	"bytes"
	"encoding/json"
)

// Animal struct for Animal
type Animal struct {
	ClassName string `json:"className"`
	Color *string `json:"color,omitempty"`
}

// NewAnimal instantiates a new Animal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnimal(className string, ) *Animal {
    this := Animal{}
    this.ClassName = className
    var color string = "red"
    this.Color = &color
    return &this
}

// NewAnimalWithDefaults instantiates a new Animal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnimalWithDefaults() *Animal {
    this := Animal{}
    var color string = "red"
    this.Color = &color
    return &this
}

// GetClassName returns the ClassName field value
func (o *Animal) GetClassName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassName
}

// SetClassName sets field value
func (o *Animal) SetClassName(v string) {
	o.ClassName = v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *Animal) GetColor() string {
	if o == nil || o.Color == nil {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Animal) GetColorOk() (string, bool) {
	if o == nil || o.Color == nil {
		var ret string
		return ret, false
	}
	return *o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *Animal) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *Animal) SetColor(v string) {
	o.Color = &v
}

type NullableAnimal struct {
	Value Animal
	ExplicitNull bool
}

func (v NullableAnimal) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableAnimal) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}