package valgo

import (
	"regexp"
)

func isStringEqualTo[T ~string](v0 T, v1 T) bool { _ = "STUB: not implemented"; return false }

func isStringGreaterThan[T ~string](v0 T, v1 T) bool { _ = "STUB: not implemented"; return false }

func isStringGreaterOrEqualTo[T ~string](v0 T, v1 T) bool { _ = "STUB: not implemented"; return false }

func isStringLessThan[T ~string](v0 T, v1 T) bool { _ = "STUB: not implemented"; return false }

func isStringLessOrEqualTo[T ~string](v0 T, v1 T) bool { _ = "STUB: not implemented"; return false }

func isStringBetween[T ~string](v T, min T, max T) bool { _ = "STUB: not implemented"; return false }

func isStringEmpty[T ~string](v T) bool { _ = "STUB: not implemented"; return false }

func isStringBlank[T ~string](v T) bool { _ = "STUB: not implemented"; return false }

func isStringInSlice[T ~string](v T, slice []T) bool { _ = "STUB: not implemented"; return false }

func isStringMatchingTo[T ~string](v T, regex *regexp.Regexp) bool {
	_ = "STUB: not implemented"
	return false
}

func isStringByteMaxLength[T ~string](v T, length int) bool {
	_ = "STUB: not implemented"
	return false
}

func isStringByteMinLength[T ~string](v T, length int) bool {
	_ = "STUB: not implemented"
	return false
}

func isStringByteLength[T ~string](v T, length int) bool { _ = "STUB: not implemented"; return false }

func isStringByteLengthBetween[T ~string](v T, min int, max int) bool {
	_ = "STUB: not implemented"
	return false
}

func isStringRuneMaxLength[T ~string](v T, length int) bool {
	_ = "STUB: not implemented"
	return false
}

func isStringRuneMinLength[T ~string](v T, length int) bool {
	_ = "STUB: not implemented"
	return false
}

func isStringRuneLength[T ~string](v T, length int) bool { _ = "STUB: not implemented"; return false }

func isStringRuneLengthBetween[T ~string](v T, min int, max int) bool {
	_ = "STUB: not implemented"
	return false
}

// The `ValidatorString` provides functions for setting validation rules for
// a string value type, or a custom type based on a string.
type ValidatorString[T ~string] struct {
	context *ValidatorContext
}

// Receive a string value to validate.
//
// The value can also be a custom string type such as type Status string;.
//
// Optionally, the function can receive a name and title, in that order, to be
// displayed in the error messages. A value_%N` pattern is used as a name in the
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name phone_number will be
// humanized as Phone Number.

func String[T ~string](value T, nameAndTitle ...string) *ValidatorString[T] {
	_ = "STUB: not implemented"
	return nil
}

// Return the context of the validator. The context is useful to create a custom
// validator by extending this validator.
func (validator *ValidatorString[T]) Context() *ValidatorContext {
	_ = "STUB: not implemented"
	return nil

	// Invert the boolean value associated with the next validator function.
	// For example:
	//
	//	// It will return false because Not() inverts the boolean value associated with the Blank() function
	//	Is(v.String("").Not().Blank()).Valid()
}

func (validator *ValidatorString[T]) Not() *ValidatorString[T] {
	_ = "STUB: not implemented"
	return nil
}

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
func (validator *ValidatorString[T]) Or() *ValidatorString[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a string value is equal to another. This function internally uses
// the golang `==` operator.
// For example:
//
//	status := "running"
//	Is(v.String(status).Equal("running"))
func (validator *ValidatorString[T]) EqualTo(value T, template ...string) *ValidatorString[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a string value is greater than another. This function internally
// uses the golang `>` operator.
// For example:
//
//	section := "bb"
//	Is(v.String(section).GreaterThan("ba"))
func (validator *ValidatorString[T]) GreaterThan(value T, template ...string) *ValidatorString[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a string value is greater than or equal to another. This function
// internally uses the golang `>=` operator.
// For example:
//
//	section := "bc"
//	Is(v.String(section).GreaterOrEqualTo("bc"))
func (validator *ValidatorString[T]) GreaterOrEqualTo(value T, template ...string) *ValidatorString[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a string value is less than another. This function internally
// uses the golang `<` operator.
// For example:
//
//	section := "bb"
//	Is(v.String(section).LessThan("bc"))
func (validator *ValidatorString[T]) LessThan(value T, template ...string) *ValidatorString[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a string value is less than or equal to another. This function
// internally uses the golang `<=` operator to compare two strings.
// For example:
//
//	section := "bc"
//	Is(v.String(section).LessOrEqualTo("bc"))
func (validator *ValidatorString[T]) LessOrEqualTo(value T, template ...string) *ValidatorString[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a string value is empty. Return false if the length of the string
// is greater than zero, even if the string has only spaces.
//
// For checking if the string has only spaces, use the function `Blank()`
// instead.
// For example:
//
//	Is(v.String("").Empty()) // Will be true
//	Is(v.String(" ").Empty()) // Will be false
func (validator *ValidatorString[T]) Empty(template ...string) *ValidatorString[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a string value is blank. Blank will be true if the length
// of the string is zero or if the string only has spaces.
// For example:
//
//	Is(v.String("").Empty()) // Will be true
//	Is(v.String(" ").Empty()) // Will be true
func (validator *ValidatorString[T]) Blank(template ...string) *ValidatorString[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a string value passes a custom function.
// For example:
//
//	status := ""
//	Is(v.String(status).Passing((v string) bool {
//		return v == getNewStatus()
//	})
func (validator *ValidatorString[T]) Passing(function func(v0 T) bool, template ...string) *ValidatorString[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a string is present in a string slice.
// For example:
//
//	status := "idle"
//	validStatus := []string{"idle", "paused", "stopped"}
//	Is(v.String(status).InSlice(validStatus))
func (validator *ValidatorString[T]) InSlice(slice []T, template ...string) *ValidatorString[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a string matches a regular expression.
// For example:
//
//	status := "pre-approved"
//	regex, _ := regexp.Compile("pre-.+")
//	Is(v.String(status).MatchingTo(regex))
func (validator *ValidatorString[T]) MatchingTo(regex *regexp.Regexp, template ...string) *ValidatorString[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate the maximum length (in bytes) of a string.
// For example:
//
//	slug := "myname"
//	Is(v.String(slug).MaxBytes(6))
//
// For character count, use `MaxLength` instead.
func (validator *ValidatorString[T]) MaxBytes(length int, template ...string) *ValidatorString[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate the minimum length (in bytes) of a string.
// For example:
//
//	slug := "myname"
//	Is(v.String(slug).MinBytes(6))
//
// For character count, use `MinLength` instead.
func (validator *ValidatorString[T]) MinBytes(length int, template ...string) *ValidatorString[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate the length (in bytes) of a string.
// For example:
//
//	slug := "myname"
//	Is(v.String(slug).OfByteLength(6))
//
// For character count, use `OfLength` instead.
func (validator *ValidatorString[T]) OfByteLength(length int, template ...string) *ValidatorString[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if the length (in bytes) of a string is within a range (inclusive).
// For example:
//
//	slug := "myname"
//	Is(v.String(slug).OfByteLengthBetween(2,6))
//
// For character count, use `OfLengthBetween` instead.
func (validator *ValidatorString[T]) OfByteLengthBetween(min int, max int, template ...string) *ValidatorString[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate the maximum length (in runes/characters) of a string.
// For example:
//
//	word := "虎視眈々" // 4 runes, len(word) = 12 bytes
//	Is(v.String(word).MaxLength(4))
func (validator *ValidatorString[T]) MaxLength(length int, template ...string) *ValidatorString[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate the minimum length (in runes/characters) of a string.
// For example:
//
//	word := "虎視眈々" // 4 runes, len(word) = 12 bytes
//	Is(v.String(word).MinLength(4))
func (validator *ValidatorString[T]) MinLength(length int, template ...string) *ValidatorString[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate the length (in runes/characters) of a string.
// For example:
//
//	word := "虎視眈々" // 4 runes, len(word) = 12 bytes
//	Is(v.String(word).OfLength(4))
func (validator *ValidatorString[T]) OfLength(length int, template ...string) *ValidatorString[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if the length (in runes/characters) of a string is within a range (inclusive).
// For example:
//
//	word := "虎視眈々" // 4 runes, len(word) = 12 bytes
//	Is(v.String(word).OfLengthBetween(2,4))
func (validator *ValidatorString[T]) OfLengthBetween(min int, max int, template ...string) *ValidatorString[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if the value of a string is within a range (inclusive).
// For example:
//
//	slug := "ab"
//	Is(v.String(slug).Between("ab","ac"))
func (validator *ValidatorString[T]) Between(min T, max T, template ...string) *ValidatorString[T] {
	_ = "STUB: not implemented"
	return nil
}
