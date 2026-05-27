package valgo

type validatorFragment struct {
	errorKey       string
	template       []string
	templateParams map[string]any
	function       func() bool
	boolOperation  bool
	orOperation    bool
	isValid        bool
}

// The context keeps the state and provides the functions to control a
// custom validator.
type ValidatorContext struct {
	fragments     []*validatorFragment
	value         any
	name          *string
	title         *string
	boolOperation bool
	orOperation   bool
}

// Create a new [ValidatorContext] to be used by a custom validator.
func NewContext(value any, nameAndTitle ...string) *ValidatorContext {
	_ = "STUB: not implemented"
	return nil
}

// Invert the boolean value associated with the next validator function in
// a custom validator.
func (ctx *ValidatorContext) Not() *ValidatorContext { _ = "STUB: not implemented"; return nil }

// Add Or operation to validation.
func (ctx *ValidatorContext) Or() *ValidatorContext { _ = "STUB: not implemented"; return nil }

// Add a function to a custom validator and pass a value used for the
// validator function to be displayed in the error message.
//
// Use [AddWithParams()] if the error message requires more input values.
func (ctx *ValidatorContext) AddWithValue(function func() bool, errorKey string, value any, template ...string) *ValidatorContext {
	_ = "STUB: not implemented"
	return nil
}

// Add a function to a custom validator.
func (ctx *ValidatorContext) Add(function func() bool, errorKey string, template ...string) *ValidatorContext {
	_ = "STUB: not implemented"
	return nil
}

// Add a function to a custom validator and pass a map with values used for the
// validator function to be displayed in the error message.
func (ctx *ValidatorContext) AddWithParams(function func() bool, errorKey string, templateParams map[string]any, template ...string) *ValidatorContext {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *ValidatorContext) validateIs(validation *Validation) *Validation {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *ValidatorContext) validateCheck(validation *Validation) *Validation {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *ValidatorContext) validate(validation *Validation, shortCircuit bool) *Validation {
	_ = "STUB: not implemented"
	// valid := true
	return nil
}

// Iterating through each fragment in the context's fragment list

// If the previous fragment is not valid, the current fragment is not in an "or" operation, and the short circuit flag is true,
// we return the current state of the validation without evaluating the current fragment

// If the current fragment is a part of an "or" operation and the previous fragment in the "or" operation
// is valid, we mark the current fragment as valid and move to the next iteration

// Evaluating the validation function of the current fragment and updating the valid flag
// The valid flag will be true only if the fragment function returns a value matching the fragment's boolean operation
// and the valid flag was true before this evaluation

// If the current fragment is valid and is part of an "or" operation, we backtrack to mark all preceding
// fragments in the "or" operation chain as valid

// Breaking the loop when we reach the start of the "or" operation chain

// Setting the validation state of the current fragment
// valid = fragment.isValid && valid

// Return the value being validated in a custom validator.
func (ctx *ValidatorContext) Value() any { _ = "STUB: not implemented"; return *new(any) }
