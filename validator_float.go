package valgo

// The [ValidatorFloat] provides functions for setting validation rules for a
// float value types, or a custom type based on a float32 or float64.
type ValidatorFloat[T ~float32 | ~float64] struct {
	context *ValidatorContext
}

// Receives a float32 value to validate.
//
// The value can also be a custom float32 type such as type Price float32.
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N" pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func Float32[T ~float32](value T, nameAndTitle ...string) *ValidatorFloat[T] {
	_ = "STUB: not implemented"
	return nil
}

// Receives a float64 value to validate.
//
// The value can also be a custom float64 type such as type Price float64.
//
// Optionally, the function can receive a name and title, in that order,
// to be displayed in the error messages. A `value_%N" pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func Float64[T ~float64](value T, nameAndTitle ...string) *ValidatorFloat[T] {
	_ = "STUB: not implemented"
	return nil
}

// Return the context of the validator. The context is useful to create a custom
// validator by extending this validator.
func (validator *ValidatorFloat[T]) Context() *ValidatorContext {
	_ = "STUB: not implemented"
	return nil

	// Invert the logical value associated with the next validator function.
	// For example:
	//
	//	// It will return false because Not() inverts the boolean value associated with the Zero() function
	//	Is(v.Float32(0).Not().Zero()).Valid()
}

func (validator *ValidatorFloat[T]) Not() *ValidatorFloat[T] { _ = "STUB: not implemented"; return nil }

// Introduces a logical OR in the chain of validation conditions, affecting the
// evaluation order and priority of subsequent validators. A value passes the
// validation if it meets any one condition following the Or() call, adhering to
// a left-to-right evaluation. This mechanism allows for validating against
// multiple criteria where satisfying any single criterion is sufficient.
// Example:
//
//	// This validator will pass because the input is Zero.
//	input := float32(0)
//	isValid := v.Is(v.Float32(input).GreaterThan(5).Or().Zero()).Valid()
func (validator *ValidatorFloat[T]) Or() *ValidatorFloat[T] { _ = "STUB: not implemented"; return nil }

// Validate if a numeric value is equal to another. This function internally uses
// the golang `==` operator.
// For example:
//
//	quantity := float32(2)
//	Is(v.Float32(quantity).EqualTo(2))
func (validator *ValidatorFloat[T]) EqualTo(value T, template ...string) *ValidatorFloat[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a numeric value is greater than another. This function internally
// uses the golang `>` operator.
// For example:
//
//	quantity := float32(3)
//	Is(v.Float32(quantity).GreaterThan(2))
func (validator *ValidatorFloat[T]) GreaterThan(value T, template ...string) *ValidatorFloat[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a numeric value is greater than or equal to another. This function
// internally uses the golang `>=` operator.
// For example:
//
//	quantity := float32(3)
//	Is(v.Float32(quantity).GreaterOrEqualTo(3))
func (validator *ValidatorFloat[T]) GreaterOrEqualTo(value T, template ...string) *ValidatorFloat[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a numeric value is less than another. This function internally
// uses the golang `<` operator.
// For example:
//
//	quantity := float32(2)
//	Is(v.Float32(quantity).LessThan(3))
func (validator *ValidatorFloat[T]) LessThan(value T, template ...string) *ValidatorFloat[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a numeric value is less than or equal to another. This function
// internally uses the golang `<=` operator.
// For example:
//
//	quantity := float32(2)
//	Is(v.Float32(quantity).LessOrEqualTo(2))
func (validator *ValidatorFloat[T]) LessOrEqualTo(value T, template ...string) *ValidatorFloat[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a number is within a range (inclusive).
// For example:
//
//	Is(v.Float32(3).Between(2,6))
func (validator *ValidatorFloat[T]) Between(min T, max T, template ...string) *ValidatorFloat[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a numeric value is zero.
//
// For example:
//
//	Is(v.Float32(0).Zero())
func (validator *ValidatorFloat[T]) Zero(template ...string) *ValidatorFloat[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a numeric value is positive (greater than zero).
//
// For example:
//
//	Is(v.Float32(5.5).Positive())
func (validator *ValidatorFloat[T]) Positive(template ...string) *ValidatorFloat[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a numeric value is negative (less than zero).
//
// For example:
//
//	Is(v.Float32(-5.5).Negative())
func (validator *ValidatorFloat[T]) Negative(template ...string) *ValidatorFloat[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a numeric value passes a custom function.
// For example:
//
//	quantity := float32(2)
//	Is(v.Float32(quantity).Passing((v float32) bool {
//		return v == getAllowedQuantity()
//	})
func (validator *ValidatorFloat[T]) Passing(function func(v T) bool, template ...string) *ValidatorFloat[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a number is present in a numeric slice.
// For example:
//
//	quantity := float32(3)
//	validQuantities := []float32{1,3,5}
//	Is(v.Float32(quantity).InSlice(validQuantities))
func (validator *ValidatorFloat[T]) InSlice(slice []T, template ...string) *ValidatorFloat[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a numeric value is NaN (Not a Number).
//
// For example:
//
//	Is(v.Float32(math.NaN()).NaN())
func (validator *ValidatorFloat[T]) NaN(template ...string) *ValidatorFloat[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a numeric value is infinite (positive or negative infinity).
//
// For example:
//
//	Is(v.Float32(math.Inf(1)).Infinite())
func (validator *ValidatorFloat[T]) Infinite(template ...string) *ValidatorFloat[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a numeric value is finite (not NaN and not infinite).
//
// For example:
//
//	Is(v.Float32(3.14).Finite())
func (validator *ValidatorFloat[T]) Finite(template ...string) *ValidatorFloat[T] {
	_ = "STUB: not implemented"
	return nil
}
