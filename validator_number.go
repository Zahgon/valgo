package valgo

//go:generate go run generator/main.go

// Custom generic type covering all numeric types. This type is used as the
// value type in ValidatorNumber and ValidatorNumberP.
type TypeNumber interface {
	~int |
		~int8 |
		~int16 |
		~int32 |
		~int64 |
		~uint |
		~uint8 |
		~uint16 |
		~uint32 |
		~uint64 |
		~float32 |
		~float64
}

func isNumberEqualTo[T TypeNumber](v0 T, v1 T) bool { _ = "STUB: not implemented"; return false }

func isNumberGreaterThan[T TypeNumber](v0 T, v1 T) bool { _ = "STUB: not implemented"; return false }

func isNumberGreaterOrEqualTo[T TypeNumber](v0 T, v1 T) bool {
	_ = "STUB: not implemented"
	return false
}

func isNumberLessThan[T TypeNumber](v0 T, v1 T) bool { _ = "STUB: not implemented"; return false }

func isNumberLessOrEqualTo[T TypeNumber](v0 T, v1 T) bool { _ = "STUB: not implemented"; return false }

func isNumberBetween[T TypeNumber](v T, min T, max T) bool { _ = "STUB: not implemented"; return false }

func isNumberZero[T TypeNumber](v T) bool { _ = "STUB: not implemented"; return false }

func isNumberInSlice[T TypeNumber](v T, slice []T) bool { _ = "STUB: not implemented"; return false }

// The [ValidatorNumber] provides functions for setting validation rules for a
// [TypeNumber] value type, or a custom type based on a [TypeNumber].
//
// [TypeNumber] is a generic interface defined by Valgo that generalizes any
// standard Golang type.
type ValidatorNumber[T TypeNumber] struct {
	context *ValidatorContext
}

// Receives a numeric value to validate.
//
// The value can be any golang numeric type (int64, int32, float32, uint,
// etc.) or a custom numeric type such as `type Level int32;`
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N`` pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`

func Number[T TypeNumber](value T, nameAndTitle ...string) *ValidatorNumber[T] {
	_ = "STUB: not implemented"
	return nil
}

// Return the context of the validator. The context is useful to create a custom
// validator by extending this validator.
func (validator *ValidatorNumber[T]) Context() *ValidatorContext {
	_ = "STUB: not implemented"
	return nil

	// Invert the logical value associated with the next validator function.
	// For example:
	//
	//	// It will return false because Not() inverts the boolean value associated with the Zero() function
	//	Is(v.Number(0).Not().Zero()).Valid()
}

func (validator *ValidatorNumber[T]) Not() *ValidatorNumber[T] {
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
//	// This validator will pass because the input is Zero.
//	input := 0
//	isValid := v.Is(v.Number(input).GreaterThan(5).Or().Zero()).Valid()
func (validator *ValidatorNumber[T]) Or() *ValidatorNumber[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a numeric value is equal to another. This function internally uses
// the golang `==` operator.
// For example:
//
//	quantity := 2
//	Is(v.Number(quantity).Equal(2))
func (validator *ValidatorNumber[T]) EqualTo(value T, template ...string) *ValidatorNumber[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a numeric value is greater than another. This function internally
// uses the golang `>` operator.
// For example:
//
//	quantity := 3
//	Is(v.Number(quantity).GreaterThan(2))
func (validator *ValidatorNumber[T]) GreaterThan(value T, template ...string) *ValidatorNumber[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a numeric value is greater than or equal to another. This function
// internally uses the golang `>=` operator.
// For example:
//
//	quantity := 3
//	Is(v.Number(quantity).GreaterOrEqualTo(3))
func (validator *ValidatorNumber[T]) GreaterOrEqualTo(value T, template ...string) *ValidatorNumber[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a numeric value is less than another. This function internally
// uses the golang `<` operator.
// For example:
//
//	quantity := 2
//	Is(v.Number(quantity).LessThan(3))
func (validator *ValidatorNumber[T]) LessThan(value T, template ...string) *ValidatorNumber[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a numeric value is less than or equal to another. This function
// internally uses the golang `<=` operator.
// For example:
//
//	quantity := 2
//	Is(v.Number(quantity).LessOrEqualTo(2))
func (validator *ValidatorNumber[T]) LessOrEqualTo(value T, template ...string) *ValidatorNumber[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a number is within a range (inclusive).
// For example:
//
//	Is(v.Number(3).Between(2,6))
func (validator *ValidatorNumber[T]) Between(min T, max T, template ...string) *ValidatorNumber[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a numeric value is zero.
//
// For example:
//
//	Is(v.Number(0).Zero())
func (validator *ValidatorNumber[T]) Zero(template ...string) *ValidatorNumber[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a numeric value passes a custom function.
// For example:
//
//	quantity := 2
//	Is(v.Number(quantity).Passing((v int) bool {
//		return v == getAllowedQuantity()
//	})
func (validator *ValidatorNumber[T]) Passing(function func(v T) bool, template ...string) *ValidatorNumber[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a number is present in a numeric slice.
// For example:
//
//	quantity := 3
//	validQuantities := []int{1,3,5}
//	Is(v.Number(quantity).InSlice(validQuantities))
func (validator *ValidatorNumber[T]) InSlice(slice []T, template ...string) *ValidatorNumber[T] {
	_ = "STUB: not implemented"
	return nil
}
