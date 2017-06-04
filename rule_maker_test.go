package mapvalidator

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

func TestRuleMaker(t *testing.T) {
	rm := NewRuleMaker("target_field", "my message", "target must exist message", nil)
	Equal(t, rm.TargetField, "target_field")
	Equal(t, rm.Message, "my message")
	Equal(t, rm.TargetMustExistMessage, "target must exist message")
	Nil(t, rm.CheckFunc)

	NewRuleMaker("target_field", "my message", "target must exist message", func(r *RuleMaker, fieldName string, fieldValue interface{}, fullMap map[string]interface{}) []error {
		Equal(t, r.TargetField, "target_field")
		Equal(t, r.Message, "my message")
		Equal(t, r.TargetMustExistMessage, "target must exist message")
		return nil
	})
}
