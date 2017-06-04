package mapvalidator

import (
	"fmt"
	"testing"

	. "github.com/stretchr/testify/assert"
)

func TestTypeWithMessageTypeStringFail(t *testing.T) {
	var rules = []Rule{
		TypeWithMsg("age", TypeString, "age must be a string type"),
	}
	m := map[string]interface{}{
		"age": 100,
	}
	errs := Validate(m, rules...)
	Contains(t, errs, fmt.Errorf("age must be a string type"))
}

func TestTypeWithMessageTypeStringSuccess(t *testing.T) {
	var rules = []Rule{
		TypeWithMsg("age", TypeString, "age must be a string type"),
	}
	m := map[string]interface{}{
		"age": "100",
	}
	errs := Validate(m, rules...)
	Empty(t, errs)
}

func TestTypeWithMessageTypeNumberFail(t *testing.T) {
	var rules = []Rule{
		TypeWithMsg("age", TypeNumber, "age must be a number type"),
	}
	m := map[string]interface{}{
		"age": "100",
	}
	errs := Validate(m, rules...)
	Contains(t, errs, fmt.Errorf("age must be a number type"))
}

func TestTypeWithMessageTypeNumberSuccess(t *testing.T) {
	var rules = []Rule{
		TypeWithMsg("age", TypeNumber, "age must be a string type"),
	}
	m := map[string]interface{}{
		"age": 100,
	}
	errs := Validate(m, rules...)
	Empty(t, errs)
}

func TestTypeWithMessageTypeDigitFail(t *testing.T) {
	var rules = []Rule{
		TypeWithMsg("age", TypeDigit, "age must be numeric"),
	}
	m := map[string]interface{}{
		"age": "100a",
	}
	errs := Validate(m, rules...)
	Contains(t, errs, fmt.Errorf("age must be numeric"))
}

func TestTypeWithMessageTypeDigitSuccess(t *testing.T) {
	var rules = []Rule{
		TypeWithMsg("age", TypeDigit, "age must be numeric"),
	}
	m := map[string]interface{}{
		"age": "100",
	}
	errs := Validate(m, rules...)
	Empty(t, errs)

	m = map[string]interface{}{
		"age": 100,
	}
	errs = Validate(m, rules...)
	Empty(t, errs)
}

func TestTypeWithMessageTypeMapFail(t *testing.T) {
	var rules = []Rule{
		TypeWithMsg("age", TypeMap, "data must be a map"),
	}
	m := map[string]interface{}{
		"age": "100",
	}
	errs := Validate(m, rules...)
	Contains(t, errs, fmt.Errorf("data must be a map"))
}

func TestTypeWithMessageTypeMapSuccess(t *testing.T) {
	var rules = []Rule{
		TypeWithMsg("age", TypeMap, "data must be a map"),
	}
	m := map[string]interface{}{
		"age": map[string]interface{}{},
	}
	errs := Validate(m, rules...)
	Empty(t, errs)
}
