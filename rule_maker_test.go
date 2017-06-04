package mapvalidator

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

func TestRuleMaker(t *testing.T) {
	rm := NewRuleMaker("target_field", false, "my message", "target must exist message", nil)
	Equal(t, rm.targetField, "target_field")
	Equal(t, rm.message, "my message")
	Equal(t, rm.targetMustExistMessage, "target must exist message")
	Nil(t, rm.checkFunc)

	NewRuleMaker("target_field", false, "my message", "target must exist message", func(r *RuleMaker, fieldName string, fieldValue interface{}, fullMap map[string]interface{}) []error {
		Equal(t, r.targetField, "target_field")
		Equal(t, r.message, "my message")
		Equal(t, r.targetMustExistMessage, "target must exist message")
		Equal(t, r.targetRequired, r.TargetMustExist())
		return nil
	})
}
