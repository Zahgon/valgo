package valgo

// The Typed validator's type that keeps its validator context.
// T can be any Go type (pointer, struct, slice, map, etc.).
type ValidatorTyped[T any] struct {
	context *ValidatorContext
}

// Receive a value of type T to validate.
//
// Optionally, the function can receive a name and title, in that order, to be
// displayed in the error messages. A value_%N pattern is used as a name in the
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name phone_number will be
// humanized as Phone Number.
//
// Example:
//
//	v.Is(v.Typed(user).Not().Nil())
func Typed[T any](value T, nameAndTitle ...string) *ValidatorTyped[T] {
	_ = "STUB: not implemented"
	return nil
}

// Return the context of the validator. The context is useful to create a custom
// validator by extending this validator.
func (validator *ValidatorTyped[T]) Context() *ValidatorContext {
	_ = "STUB: not implemented"
	return nil

	// Invert the logical value associated with the next validator function.
	// For example:
	//
	//	// It will return false because `Not()` inverts the boolean value associated with `EqualTo()`
	//	v.Is(v.Typed("a").Not().EqualTo("a")).Valid()
}

func (validator *ValidatorTyped[T]) Not() *ValidatorTyped[T] { _ = "STUB: not implemented"; return nil }

// Introduces a logical OR in the chain of validation conditions, affecting the
// evaluation order and priority of subsequent validators. A value passes the
// validation if it meets any one condition following the Or() call, adhering to
// a left-to-right evaluation. This mechanism allows for validating against
// multiple criteria where satisfying any single criterion is sufficient.
// Example:
//
//	// This validator will pass because the string is equals "test".
//	input := "test"
//	isValid := v.Is(v.String(input).MinLength(5).Or().EqualTo("test")).Valid()
func (validator *ValidatorTyped[T]) Or() *ValidatorTyped[T] { _ = "STUB: not implemented"; return nil }

// Validate if a value passes a custom function.
// The function receives a typed T value, enabling compile-time type safety.
//
// For example:
//
//	type Status string
//	status := Status("running")
//	isValid := v.Is(
//	  v.Typed(status).Passing(func(s Status) bool {
//	    return s == "running" || s == "paused"
//	  }),
//	).Valid()
func (validator *ValidatorTyped[T]) Passing(function func(v T) bool, template ...string) *ValidatorTyped[T] {
	_ = "STUB: not implemented"
	return nil
}

// Value is stored as interface{} inside the context; assert back to T.

// Validate if a value is nil.
// Works for nil-able kinds: pointers, slices, maps, chans, funcs, and interfaces.
// For non-nil-able types, this will return false.
//
// For example:
//
//	var s *string
//	v.Is(v.Typed(s).Nil())
func (validator *ValidatorTyped[T]) Nil(template ...string) *ValidatorTyped[T] {
	_ = "STUB: not implemented"
	return nil
}

// In Golang nil sometimes is not equal to raw nil, such as it's explained
// here: https://dev.to/arxeiss/in-go-nil-is-not-equal-to-nil-sometimes-jn8
// So, seems using reflection is the only option here
