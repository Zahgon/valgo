package valgo

// The [Validation] session in Valgo is the main structure for validating one or
// more values. It is called Validation in code.
//
// A [Validation] session will contain one or more Validators, where each [Validator]
// will have the responsibility to validate a value with one or more rules.
//
// There are multiple functions to create a [Validation] session, depending on the
// requirements:
//
//   - [New]()
//   - [Is](...)
//   - [In](...)
//   - [Check](...)
//   - [InRow](...)
//   - [InCell](...)
//   - [If](...)
//   - [Do](...)
//   - [When](...)
//   - [Merge](...)
//   - [AddErrorMessage](...)
//
// the function [Is](...) is likely to be the most frequently used function in your
// validations. When [Is](...) is called, the function creates a validation and
// receives a validator at the same time.
type Validation struct {
	valid bool

	_locale         *Locale
	errors          map[string]*valueError
	invalidateMap   map[string]bool
	currentIndex    int
	marshalJsonFunc func(e *Error) ([]byte, error)
}

// Options struct is used to specify options when creating a new [Validation]
// session with the [New()] function.
//
// It contains parameters for specifying a specific locale code, modify or add a
// locale, and set a custom JSON marshaler for [Error].

type Options struct {
	localeCodeDefaultFromFactory string             // Only specified by the factory
	localesFromFactory           map[string]*Locale // Only specified by the factory

	// A string field that represents the locale code to use by the [Validation]
	// session
	LocaleCode string
	// A map field that allows to modify or add a new [Locale]
	Locale *Locale
	// A function field that allows to set a custom JSON marshaler for [Error]
	MarshalJsonFunc func(e *Error) ([]byte, error)
}

// Add one or more validators to a [Validation] session.
func (validation *Validation) Is(validators ...Validator) *Validation {
	_ = "STUB: not implemented"
	return nil
}

// [If](...) is similar to [Merge](...), but merge the [Validation] session
// only when the condition is true, and returns the same [Validation] instance.
// When the condition is false, no operation is performed and the original
// instance is returned unchanged.
//
// See [Merge](...) for more information.
//
//	v.If(isAdmin, v.Is(v.String(username, "username").Not().Blank()) )
func (validation *Validation) If(condition bool, _validation *Validation) *Validation {
	_ = "STUB: not implemented"
	return nil
}

// The [Do](...) function executes the given function with the current
// [Validation] instance and returns the same instance.
//
// This allows you to extend a validation chain with additional or
// conditional rules in a concise way:
//
//	v.Is(v.String(username, "username").Not().Blank()).Do(func(val *v.Validation) {
//		if isAdmin {
//			val.Is(v.String(role, "role").Equal("admin"))
//		}
//	})
func (validation *Validation) Do(function func(val *Validation)) *Validation {
	_ = "STUB: not implemented"
	return nil
}

// [When](...) is similar to [Do](...), but executes the given function
// only when the condition is true, and returns the same [Validation] instance.
// When the condition is false, no operation is performed and the original
// instance is returned unchanged.
//
// See [Do](...) for the unconditional variant.
//
//	v.Is(v.String(username, "username").Not().Blank()).When(isAdmin, func(val *v.Validation) {
//		val.Is(v.String(role, "role").Equal("admin"))
//	})
func (validation *Validation) When(condition bool, function func(val *Validation)) *Validation {
	_ = "STUB: not implemented"
	return nil
}

// [Check](...) adds one or more validators to a [Validation] session. But unlike [Is()],
// the validators are not short-circuited.
func (validation *Validation) Check(validators ...Validator) *Validation {
	_ = "STUB: not implemented"
	return nil
}

// A [Validation] session provides this function which returns either true if
// all their validators are valid or false if any one of them is invalid.
//
// In the following example, even though the [Validator] for age is valid, the
// [Validator] for status is invalid, making the entire Validator session
// invalid.
func (validation *Validation) Valid() bool { _ = "STUB: not implemented"; return false }

// Add a map namespace to a [Validation] session.
func (validation *Validation) In(name string, _validation *Validation) *Validation {
	_ = "STUB: not implemented"
	return nil
}

// Add an indexed namespace to a [Validation] session.
func (validation *Validation) InRow(name string, index int, _validation *Validation) *Validation {
	_ = "STUB: not implemented"
	return nil
}

// Add an indexed namespace to a [Validation] session where the target is a
// single, scalar value (e.g., entries of a primitive slice). This is useful
// for validating arrays or slices of primitives. Example:
//
//	validation := valgo.InCell("tag_priority", 0,
//		valgo.Is(valgo.String("", "tag_priority", "Tag priority").Not().Blank()),
//	)
//
// The example above validates the value at tag_priority[0].
func (validation *Validation) InCell(name string, index int, _validation *Validation) *Validation {
	_ = "STUB: not implemented"
	return nil
}

// Using [Merge](...) you can merge two [Validation] sessions. When two
// validations are merged, errors with the same value name will be merged. It is
// useful for reusing validation logic.
//
// The following example merges the [Validation] session returned by the
// validatePreStatus function. Since both [Validation] sessions validate a value
// with the name status, the error returned will return two error messages, and
// without duplicate the Not().Blank() error message rule.
func (validation *Validation) Merge(_validation *Validation) *Validation {
	_ = "STUB: not implemented"
	return nil
}

func (validation *Validation) merge(prefix string, _validation *Validation) *Validation {
	_ = "STUB: not implemented"
	return nil
}

// Add an error message to the [Validation] session without executing a field
// validator. By adding this error message, the [Validation] session will be
// marked as invalid.
func (v *Validation) AddErrorMessage(name string, message string) *Validation {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validation) mergeError(prefix string, err *Error) *Validation {
	_ = "STUB: not implemented"
	return nil
}

// MergeError allows merging Valgo errors from an already validated [Validation] session.
// The function takes an Valgo [Error] pointer as an argument and returns a [Validation] pointer.
func (v *Validation) MergeError(err *Error) *Validation { _ = "STUB: not implemented"; return nil }

// MergeErrorIn allows merging Valgo errors from already validated [Validation] sessions
// within a map namespace. The function takes a namespace name and an [Error] pointer
// as arguments and returns a [Validation] pointer.
func (v *Validation) MergeErrorIn(name string, err *Error) *Validation {
	_ = "STUB: not implemented"
	return nil
}

// MergeErrorInRow allows merging Valgo errors from already validated [Validation] sessions
// within an indexed namespace. The function takes a namespace name, an index, and an [Error] pointer
// as arguments and returns a [Validation] pointer.
//
// DEPRECATED: This method is deprecated in favor of MergeErrorInIndex().
// The MergeErrorInIndex() method is a generic name to cover errors added by
// InRow() and InCell() validations.
func (v *Validation) MergeErrorInRow(name string, index int, err *Error) *Validation {
	_ = "STUB: not implemented"
	return nil
}

// MergeErrorInRow allows merging Valgo errors from already validated [Validation] sessions
// within an indexed namespace. These are errors added by InRow() and InCell() validations.
// The function takes a namespace name, an index, and an [Error] pointer
// as arguments and returns a [Validation] pointer.
func (v *Validation) MergeErrorInIndex(name string, index int, err *Error) *Validation {
	_ = "STUB: not implemented"
	return nil
}

func (validation *Validation) invalidate(name *string, title *string, fragment *validatorFragment) {
	_ = "STUB: not implemented"
	return
}

// Return a map with the information for each invalid field validator
// in the Validation session.
func (session *Validation) Errors() map[string]*valueError { _ = "STUB: not implemented"; return nil }

// Error returns the validation errors as a standard Go error interface.
//
// DEPRECATED: This method is deprecated in favor of ToError() or ToValgoError().
// The Error() method name conflicts with Go's error interface implementation
// convention, where Error() typically implements the error interface for a type.
//
// Use ToError() for standard error handling or ToValgoError() for detailed
// validation error information.
func (validation *Validation) Error(marshalJsonFun ...func(e *Error) ([]byte, error)) error {
	_ = "STUB: not implemented"
	return nil
}

// ToError returns the validation errors as a standard Go error interface.
//
// This method is useful for idiomatic error handling and integration with
// Go's native error system. It returns the same underlying error value as
// ToValgoError() but typed as the error interface.
//
// Example:
//
//	val := Is(String("", "name").Not().Blank())
//	if err := val.ToError(); err != nil {
//	    log.Printf("Validation failed: %v", err)
//	    return err
//	}
//
// An optional JSON marshaling function can be passed to customize how the
// validation errors are serialized into JSON. If no function is provided,
// a default marshaling behavior is used.
func (validation *Validation) ToError(marshalJsonFun ...func(e *Error) ([]byte, error)) error {
	_ = "STUB: not implemented"
	// We cannot simply return validation.ToValgoError(marshalJsonFun...) because
	// when ToValgoError returns nil, it's a nil *Error (concrete type), not a nil
	// error interface. This causes issues with error checking functions like
	// assert.NoError() which expect a proper nil error interface.
	return nil
}

// ToValgoError returns the validation errors as a *valgo.Error type, providing
// access to rich, structured error details. It's essentially a shortcut to
// `ToError().(*valgo.Error)`.
//
// This method returns the underlying *valgo.Error type directly, exposing
// detailed validation information such as per-field messages, templates,
// and localized titles. It's the single source of truth for validation errors.
//
// Example:
//
//	val := Is(String("", "name").Not().Blank())
//	if errInfo := val.ToValgoError(); errInfo != nil {
//	    for field, valueError := range errInfo.Errors() {
//	        fmt.Printf("Field '%s': %v\n", field, valueError.Messages())
//	    }
//	}
//
// An optional JSON marshaling function can be passed to customize how the
// validation errors are serialized into JSON. If no function is provided,
// a default marshaling behavior is used.
func (validation *Validation) ToValgoError(marshalJsonFun ...func(e *Error) ([]byte, error)) *Error {
	_ = "STUB: not implemented"
	return nil
}

// Return true if a specific field validator is valid.
func (validation *Validation) IsValid(name string) bool { _ = "STUB: not implemented"; return false }

func (validation *Validation) getOrCreateValueError(name string, title *string) *valueError {
	_ = "STUB: not implemented"
	return nil
}

func newValidation(options ...Options) *Validation { _ = "STUB: not implemented"; return nil }

// If the factory has default locale specified, we try to use it as fallback

// Skipping default option will return nil, so we can use the factory
// locale default

// If locale entries were specified, then we merge it with the calculated
// Locale from the options localeCode

// name examples:
//
//	"object.users[1].value"
//
// namespaces generated:
//
//	"object"
//	"object.users"
//	"object.users[1]"
//	"object.users[1].value"
func (validation *Validation) addInvalidationNamespaces(name string) {
	_ = "STUB: not implemented"
	return
}

// start index of current segment (after last '.')

// First '[' in this segment: add prefix without the index.
// e.g. "object.users[1]" -> add "object.users".

// End of segment: add prefix up to this dot.
// e.g. "object.users[1].value" at '.' after "[1]" -> add "object.users[1]".

// Always add the full path
