package valgo

// Implementation of the Go error interface in Valgo. The [Validation.Error()]
// method returns a value of this type.
//
// There is a function in this type, [Errors()], that returns a list of errors
// in a [Validation] session.
type Error struct {
	errors          map[string]*valueError
	marshalJsonFunc func(e *Error) ([]byte, error)
}

type errorTemplate struct {
	key      string
	template *string
	params   map[string]interface{}
}

// Contains information about each invalid field value returned by the
// [Validation] session.
type valueError struct {
	name           *string
	title          *string
	errorTemplates map[string]*errorTemplate
	errorMessages  []string
	messages       []string
	dirty          bool
	validator      *Validation
}

// The title of the invalid field value.
func (ve *valueError) Title() string {
	_ = "STUB: not implemented"
	// Lazy load the title
	return ""
}

// The name of the invalid field value.
func (ve *valueError) Name() string {
	_ = "STUB: not implemented"

	// Error messages related to an invalid field value.
	return ""
}

func (ve *valueError) Messages() []string { _ = "STUB: not implemented"; return nil }

func (ve *valueError) buildMessageFromTemplate(et *errorTemplate) string {
	_ = "STUB: not implemented"
	return ""
}

// Ensure interface{} values are string in order to be handle by fasttemplate

// Return the error message associated with a Valgo error.
func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

// Return a map with each Invalid value error.
func (e *Error) Errors() map[string]*valueError { _ = "STUB: not implemented"; return nil }

func (e *Error) prepareErrorsForMarshal() map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// Returns the JSON encoding of the validation error messages.
//
// A custom function can be set either by passing it as a parameter to
// [validation.Error()] or through [FactoryOptions].
func (e *Error) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Returns the JSON encoding of the validation error messages with the given prefix and indent.
//
// This function does not call a custom marshalJsonFunc function if it is set.
func (e *Error) MarshalJSONIndent(prefix, indent string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Returns the JSON encoding of the validation error messages with the pretty format.
//
// This is a shortcut for MarshalJSONIndent("", "  ").
// It does not call a custom marshalJSONFunc function if it is set.
func (e *Error) MarshalJSONPretty() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
