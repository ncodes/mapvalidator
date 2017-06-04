package mapvalidator

import (
	"testing"

	"fmt"

	. "github.com/stretchr/testify/assert"
)

func TestRequired(t *testing.T) {

	var rules = []Rule{
		Required("name"),
	}
	m := map[string]interface{}{
		"name": interface{}(0),
	}
	errs := Validate(m, rules...)
	Contains(t, errs, fmt.Errorf("field 'name' is required"))

	m = map[string]interface{}{
		"name": "",
	}
	errs = Validate(m, rules...)
	Contains(t, errs, fmt.Errorf("field 'name' is required"))

	m = map[string]interface{}{
		"name": false,
	}
	errs = Validate(m, rules...)
	Contains(t, errs, fmt.Errorf("field 'name' is required"))

	m = map[string]interface{}{
		"name": map[string]interface{}{},
	}
	errs = Validate(m, rules...)
	Contains(t, errs, fmt.Errorf("field 'name' is required"))

	m = map[string]interface{}{
		"name": map[string]string{},
	}
	errs = Validate(m, rules...)
	Contains(t, errs, fmt.Errorf("field 'name' is required"))

	m = map[string]interface{}{
		"name": []string{},
	}
	errs = Validate(m, rules...)
	Contains(t, errs, fmt.Errorf("field 'name' is required"))

	m = map[string]interface{}{
		"name": []int{},
	}
	errs = Validate(m, rules...)
	Contains(t, errs, fmt.Errorf("field 'name' is required"))

	m = map[string]interface{}{
		"name": []float32{},
	}
	errs = Validate(m, rules...)
	Contains(t, errs, fmt.Errorf("field 'name' is required"))

	m = map[string]interface{}{
		"name": []float64{},
	}
	errs = Validate(m, rules...)
	Contains(t, errs, fmt.Errorf("field 'name' is required"))

	m = map[string]interface{}{}
	errs = Validate(m, rules...)
	Contains(t, errs, fmt.Errorf("field 'name' is required"))

	m = map[string]interface{}{
		"name": []float64{1, 2, 3},
	}
	errs = Validate(m, rules...)
	Empty(t, errs)
}

func TestRequiredWithMsg(t *testing.T) {

	var rules = []Rule{
		RequiredWithMsg("name", "name is required"),
	}
	m := map[string]interface{}{
		"name": interface{}(0),
	}
	errs := Validate(m, rules...)
	Contains(t, errs, fmt.Errorf("name is required"))

	m = map[string]interface{}{
		"name": "",
	}
	errs = Validate(m, rules...)
	Contains(t, errs, fmt.Errorf("name is required"))

	m = map[string]interface{}{}
	errs = Validate(m, rules...)
	Contains(t, errs, fmt.Errorf("name is required"))
}
