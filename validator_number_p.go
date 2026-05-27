package valgo

//go:generate go run generator/main.go

// The Numeric pointer validator type that keeps its validator context.
type ValidatorNumberP[T TypeNumber] struct {
	context *ValidatorContext
}

// Receives a numeric pointer to validate.
//
// The value can be any golang numeric pointer type (*int64, *int32, *float32, *uint,
// etc.) or a custom numeric type such as `type Level *int32;`
//
// Optionally, the function can receive a name and title, in that order,
// to be used in the error messages. A `value_%N“ pattern is used as a name in
// error messages if a name and title are not supplied; for example: value_0.
// When the name is provided but not the title, then the name is humanized to be
// used as the title as well; for example the name `phone_number` will be
// humanized as `Phone Number`
func NumberP[T TypeNumber](value *T, nameAndTitle ...string) *ValidatorNumberP[T] {
	_ = "STUB: not implemented"
	return nil
}

// Return the context of the validator. The context is useful to create a custom
// validator by extending this validator.
func (validator *ValidatorNumberP[T]) Context() *ValidatorContext {
	_ = "STUB: not implemented"
	return nil

	// Invert the boolean value associated with the next validator function.
	// For example:
	//
	//	// It will return false because Not() inverts the boolean value associated with the Zero() function
	//	n := 0
	//	Is(v.NumberP(&n).Not().Zero()).Valid()
}

func (validator *ValidatorNumberP[T]) Not() *ValidatorNumberP[T] {
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
//	isValid := v.Is(v.NumberP(&input).GreaterThan(5).Or().Zero()).Valid()
func (validator *ValidatorNumberP[T]) Or() *ValidatorNumberP[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a numeric pointer value is equal to another value. This function internally uses
// the golang `==` operator.
// For example:
//
//	quantity := 2
//	Is(v.NumberP(quantity).Equal(2))
func (validator *ValidatorNumberP[T]) EqualTo(value T, template ...string) *ValidatorNumberP[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a numeric pointer value is greater than another value. This function internally
// uses the golang `>` operator.
// For example:
//
//	quantity := 3
//	Is(v.NumberP(&quantity).GreaterThan(2))
func (validator *ValidatorNumberP[T]) GreaterThan(value T, template ...string) *ValidatorNumberP[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a numeric pointer value is greater than or equal to another value. This function
// internally uses the golang `>=` operator.
// For example:
//
//	quantity := 3
//	Is(v.NumberP(&quantity).GreaterOrEqualTo(3))
func (validator *ValidatorNumberP[T]) GreaterOrEqualTo(value T, template ...string) *ValidatorNumberP[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a numeric pointer value is less than another value. This function internally
// uses the golang `<` operator.
// For example:
//
//	quantity := 2
//	Is(v.NumberP(&quantity).LessThan(3))
func (validator *ValidatorNumberP[T]) LessThan(value T, template ...string) *ValidatorNumberP[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a numeric pointer value is less than or equal to another value. This function
// internally uses the golang `<=` operator.
// For example:
//
//	quantity := 2
//	Is(v.NumberP(&quantity).LessOrEqualTo(2))
func (validator *ValidatorNumberP[T]) LessOrEqualTo(value T, template ...string) *ValidatorNumberP[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if the value of a numeric pointer is within a range (inclusive).
// For example:
//
//	n := 3
//	Is(v.NumberP(&n).Between(2,6))
func (validator *ValidatorNumberP[T]) Between(min T, max T, template ...string) *ValidatorNumberP[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a numeric pointer value is zero.
//
// For example:
//
//	n := 0
//	Is(v.NumberP(&n).Zero())
func (validator *ValidatorNumberP[T]) Zero(template ...string) *ValidatorNumberP[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a numeric pointer value is zero or nil.
//
// For example:
//
//	var _quantity *int
//	Is(v.NumberP(_quantity).ZeroOrNil()) // Will be true
func (validator *ValidatorNumberP[T]) ZeroOrNil(template ...string) *ValidatorNumberP[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a numeric pointer value is nil.
//
// For example:
//
//	var quantity *int
//	Is(v.NumberP(quantity).Nil()) // Will be true
func (validator *ValidatorNumberP[T]) Nil(template ...string) *ValidatorNumberP[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a numeric pointer value passes a custom function.
// For example:
//
//	quantity := 2
//	Is(v.NumberP(&quantity).Passing((v *int) bool {
//		return *v == getAllowedQuantity()
//	})
func (validator *ValidatorNumberP[T]) Passing(function func(v *T) bool, template ...string) *ValidatorNumberP[T] {
	_ = "STUB: not implemented"
	return nil
}

// Validate if a numeric pointer value is present in a numeric slice.
// For example:
//
//	quantity := 3
//	validQuantities := []int{1,3,5}
//	Is(v.NumberP(&quantity).InSlice(validQuantities))
func (validator *ValidatorNumberP[T]) InSlice(slice []T, template ...string) *ValidatorNumberP[T] {
	_ = "STUB: not implemented"
	return nil
}
