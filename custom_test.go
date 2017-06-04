package mapvalidator

import (
	"fmt"
	"testing"

	. "github.com/stretchr/testify/assert"
)

func TestCustomFail(t *testing.T) {
	var rules = []Rule{
		Custom("name", "must start with '$'", func(fieldValue interface{}, fullMap map[string]interface{}) bool {
			if v, ok := fieldValue.(string); ok {
				return v[0] == '$'
			}
			return false
		}),
	}
	m := map[string]interface{}{
		"name": "ben",
	}
	errs := Validate(m, rules...)
	Contains(t, errs, fmt.Errorf("must start with '$'"))
}

func TestCustomSuccess(t *testing.T) {
	var rules = []Rule{
		Custom("name", "must start with '$'", func(fieldValue interface{}, fullMap map[string]interface{}) bool {
			if v, ok := fieldValue.(string); ok {
				return v[0] == '$'
			}
			return false
		}),
	}
	m := map[string]interface{}{
		"name": "$ben",
	}
	errs := Validate(m, rules...)
	Empty(t, errs)
}
