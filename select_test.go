package fselect

import (
	"reflect"
	"testing"
)

const (
	errInvalidValue = "value is not a struct or pointer to struct"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Pet struct {
	FirstName string `col:"first_name"`
	LastName  string `col:"last_name"`
	Age       int    `col:"age"`
}

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

func TestInvalidVAll(t *testing.T) {
	defer func() {
		if err := recover(); err != ErrInvalidV {
			t.Fatal("Expected ErrInvalidV")
		}
	}()

	All(0)
}

func TestInvalidVAllExcept(t *testing.T) {
	defer func() {
		if err := recover(); err != ErrInvalidV {
			t.Fatal("Expected ErrInvalidV")
		}
	}()

	AllExcept(0)
}

func TestInvalidVOnly(t *testing.T) {
	defer func() {
		if err := recover(); err != ErrInvalidV {
			t.Fatal("Expected ErrInvalidV")
		}
	}()

	Only(0)
}

func TestFieldsNotFoundAllExcept(t *testing.T) {
	defer func() {
		if err := recover(); err != ErrSomeFieldsNotFound {
			t.Fatal("Expected ErrSomeFieldsNotFound")
		}
	}()

	AllExcept(newPerson(), "this is an invalid field")
}

func TestFieldsNotFoundOnly(t *testing.T) {
	defer func() {
		if err := recover(); err != ErrSomeFieldsNotFound {
			t.Fatal("Expected ErrSomeFieldsNotFound")
		}
	}()

	Only(newPerson(), "this is an invalid field")
}

func TestFields(t *testing.T) {
	p := newPerson()
	s := All(p)
	if s.Fields()[2] != "Age" {
		t.Fatal(`assert: s.Fields()[2] != "Age"`)
	}
}

func TestFieldNames(t *testing.T) {
	p := newPet()
	s := AllExcept(p, "first_name")
	if s.FieldString() != "last_name,age" {
		t.Fatal(`assert: s.FieldString() != "last_name,age"`)
	}
}

func TestBindVars(t *testing.T) {
	p := newPerson()
	s := Only(p, "Age")
	if s.BindVars() != "?" {
		t.Fatal(`assert: s.BindVars() != "?"`)
	}
}

func TestArgs(t *testing.T) {
	p := newPet()
	s := Only(p, "last_name")
	if s.Args()[0] != p.LastName {
		t.Fatal(`assert: s.Args()[0] != p.LastName`)
	}
}

func TestPreparef(t *testing.T) {
	const expect = "INSERT INTO pets (first_name,last_name,age) VALUES (?,?,?)"
	query := All(newPet()).Preparef("INSERT INTO pets (%s) VALUES (%s)")

	if query != expect {
		t.Fatal(`assert: query != expect`)
	}
}

func newPerson() *Person {
	return &Person{"John", "Doe", 21}
}

func newPet() *Pet {
	return &Pet{"Bella", "Sky", 5}
}