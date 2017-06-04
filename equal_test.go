package mapvalidator

import (
	"fmt"
	"testing"

	. "github.com/stretchr/testify/assert"
)

func TestEqualStringWithMsgFail(t *testing.T) {
	var rules = []Rule{
		EqualStringWithMsg("age", "120", "age is different"),
	}
	m := map[string]interface{}{
		"age": "100",
	}
	errs := Validate(m, rules...)
	Contains(t, errs, fmt.Errorf("age is different"))

	m = map[string]interface{}{
		"age": 100,
	}
	errs = Validate(m, rules...)
	Contains(t, errs, fmt.Errorf("age is different"))
}

func TestEqualStringWithMsgSuccess(t *testing.T) {
	var rules = []Rule{
		EqualStringWithMsg("age", "120", "age is different"),
	}
	m := map[string]interface{}{
		"age": "120",
	}
	errs := Validate(m, rules...)
	Empty(t, errs)
}

func TestEqualIntWithMsgFail(t *testing.T) {
	var rules = []Rule{
		EqualIntWithMsg("age", 120, "age is different"),
	}
	m := map[string]interface{}{
		"age": 100,
	}
	errs := Validate(m, rules...)
	Contains(t, errs, fmt.Errorf("age is different"))

	m = map[string]interface{}{
		"age": "100",
	}
	errs = Validate(m, rules...)
	Contains(t, errs, fmt.Errorf("age is different"))
}

func TestEqualIntWithMsgSuccess(t *testing.T) {
	var rules = []Rule{
		EqualIntWithMsg("age", 120, "age is different"),
	}
	m := map[string]interface{}{
		"age": 120,
	}
	errs := Validate(m, rules...)
	Empty(t, errs)
}

func TestEqualInt32WithMsgFail(t *testing.T) {
	var rules = []Rule{
		EqualInt32WithMsg("age", 120, "age is different"),
	}
	m := map[string]interface{}{
		"age": int32(100),
	}
	errs := Validate(m, rules...)
	Contains(t, errs, fmt.Errorf("age is different"))

	m = map[string]interface{}{
		"age": "100",
	}
	errs = Validate(m, rules...)
	Contains(t, errs, fmt.Errorf("age is different"))
}

func TestEqualInt32WithMsgSuccess(t *testing.T) {
	var rules = []Rule{
		EqualInt32WithMsg("age", 120, "age is different"),
	}
	m := map[string]interface{}{
		"age": int32(120),
	}
	errs := Validate(m, rules...)
	Empty(t, errs)
}

func TestEqualInt64WithMsgFail(t *testing.T) {
	var rules = []Rule{
		EqualInt64WithMsg("age", 120, "age is different"),
	}
	m := map[string]interface{}{
		"age": int64(100),
	}
	errs := Validate(m, rules...)
	Contains(t, errs, fmt.Errorf("age is different"))

	m = map[string]interface{}{
		"age": "100",
	}
	errs = Validate(m, rules...)
	Contains(t, errs, fmt.Errorf("age is different"))
}

func TestEqualInt64WithMsgSuccess(t *testing.T) {
	var rules = []Rule{
		EqualInt64WithMsg("age", 120, "age is different"),
	}
	m := map[string]interface{}{
		"age": int64(120),
	}
	errs := Validate(m, rules...)
	Empty(t, errs)
}

func TestEqualFloat32WithMsgFail(t *testing.T) {
	var rules = []Rule{
		EqualFloat32WithMsg("age", 120, "age is different"),
	}
	m := map[string]interface{}{
		"age": float64(100),
	}
	errs := Validate(m, rules...)
	Contains(t, errs, fmt.Errorf("age is different"))

	m = map[string]interface{}{
		"age": "100",
	}
	errs = Validate(m, rules...)
	Contains(t, errs, fmt.Errorf("age is different"))
}

func TestEqualFloat32WithMsgSuccess(t *testing.T) {
	var rules = []Rule{
		EqualFloat32WithMsg("age", 120, "age is different"),
	}
	m := map[string]interface{}{
		"age": float32(120),
	}
	errs := Validate(m, rules...)
	Empty(t, errs)
}

func TestEqualFloat64WithMsgFail(t *testing.T) {
	var rules = []Rule{
		EqualFloat64WithMsg("age", 120, "age is different"),
	}
	m := map[string]interface{}{
		"age": float64(100),
	}
	errs := Validate(m, rules...)
	Contains(t, errs, fmt.Errorf("age is different"))

	m = map[string]interface{}{
		"age": "100",
	}
	errs = Validate(m, rules...)
	Contains(t, errs, fmt.Errorf("age is different"))
}

func TestEqualFloat64WithMsgSuccess(t *testing.T) {
	var rules = []Rule{
		EqualFloat64WithMsg("age", 120, "age is different"),
	}
	m := map[string]interface{}{
		"age": float64(120),
	}
	errs := Validate(m, rules...)
	Empty(t, errs)
}
