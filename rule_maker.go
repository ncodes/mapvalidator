package mapvalidator

// RuleMakerCheckFunc is the function to pass to the rule maker which will be called in
// the new rule's check method.
type RuleMakerCheckFunc func(r *RuleMaker, fieldName string, fieldValue interface{}, fullMap map[string]interface{}) []error

// RuleMaker describes a rule for a RuleMaker target field. It satisfies the Rule interface
type RuleMaker struct {
	TargetField            string
	Message                string
	TargetMustExistMessage string
	CheckFunc              RuleMakerCheckFunc
}

// NewRuleMaker creates a new required rule
func NewRuleMaker(targetField, message, targetMustExistMessage string, checkFunc RuleMakerCheckFunc) *RuleMaker {
	return &RuleMaker{
		TargetField:            targetField,
		Message:                message,
		TargetMustExistMessage: targetMustExistMessage,
		CheckFunc:              checkFunc,
	}
}

// GetTargetField returns the name of the field to perform validation against
func (r *RuleMaker) GetTargetField() string {
	return r.TargetField
}

// TargetMustExist forces an error to returned if target field does not exist in map
func (r *RuleMaker) TargetMustExist() bool {
	return true
}

// GetTargetMustExistMessage returns the error message to return when target field does not exists in map
func (r *RuleMaker) GetTargetMustExistMessage() string {
	return r.TargetMustExistMessage
}

// Check performs all validations against the field value and returns every validation error failed
func (r *RuleMaker) Check(fieldName string, fieldValue interface{}, fullMap map[string]interface{}) []error {
	return r.CheckFunc(r, fieldName, fieldValue, fullMap)
}
