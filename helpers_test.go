package fselect

import (
	"reflect"
	"testing"
)

func TestGetFieldName(t *testing.T) {
	value := reflect.Indirect(reflect.ValueOf(newPet()))
	if value.Kind() != reflect.Struct {
		t.Fatal(errInvalidValue)
	}
	fieldType := value.Type().Field(0)

	if getFieldName(&fieldType) != "first_name" {
		t.Fatal(`assert: fieldType != "first_name"`)
	}
}

func TestSliceContains(t *testing.T) {
	slice := []string{"aaa", "bbb", "ccc", "abc", "cba"}

	if sliceContains("abcd", slice) {
		t.Fatal(`"abcd" doesn't exist in slice`)
	}
	if !sliceContains("cba", slice) {
		t.Fatal(`"cba" does exist in slice`)
	}
}

func TestRepeatString(t *testing.T) {
	const expect = "test, test, test, test, test, test"

	if repeatString("test", ", ", 6) != expect {
		t.Fail()
	}
}

func newPerson() *Person {
	return &Person{"John", "Doe", 21}
}

func newPet() *Pet {
	return &Pet{"Bella", "Sky", 5}
}