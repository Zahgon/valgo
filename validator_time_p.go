package valgo

import (
	"time"
)

// ValidatorTimeP is a type that facilitates validation for time pointer variables.
// It retains a context that records details about the validation process.
type ValidatorTimeP struct {
	context *ValidatorContext
}

// TimeP initializes a new ValidatorTimeP instance with the provided time pointer
// and optional name and title arguments for detailed error messages.
//
// Usage example:
//
//	var myTime *time.Time
//	v.TimeP(myTime, "start_time", "Start Time")
func TimeP(value *time.Time, nameAndTitle ...string) *ValidatorTimeP {
	_ = "STUB: not implemented"
	return nil
}

// Context retrieves the context associated with the validator.
func (validator *ValidatorTimeP) Context() *ValidatorContext { _ = "STUB: not implemented"; return nil }

// Not negates the result of the next validator function in the chain.
//
// Usage example:
//
//	t := time.Now()
//	Is(v.TimeP(&t).Not().Zero()).Valid()  // Will return false since t is not a zero time.
func (validator *ValidatorTimeP) Not() *ValidatorTimeP { _ = "STUB: not implemented"; return nil }

// Introduces a logical OR in the chain of validation conditions, affecting the
// evaluation order and priority of subsequent validators. A value passes the
// validation if it meets any one condition following the Or() call, adhering to
// a left-to-right evaluation. This mechanism allows for validating against
// multiple criteria where satisfying any single criterion is sufficient.
// Example:
//
//	// This validator will pass because the time is before or equal to time.Now().
//	t := time.Now()
//	isValid := v.Is(v.TimeP(&t).Nil().Or().BeforeOrEqualTo(time.Now())).Valid()
func (validator *ValidatorTimeP) Or() *ValidatorTimeP { _ = "STUB: not implemented"; return nil }

// EqualTo validates that the time pointer is equal to the specified time value.
//
// Usage example:
//
//	t1 := time.Now()
//	t2 := t1
//	Is(v.TimeP(&t1).EqualTo(t2)).Valid()  // Will return true.
func (validator *ValidatorTimeP) EqualTo(value time.Time, template ...string) *ValidatorTimeP {
	_ = "STUB: not implemented"
	return nil
}

// After validates that the time pointer is after the specified time value.
//
// Usage example:
//
//	t1 := time.Now()
//	t2 := t1.Add(-time.Hour)
//	Is(v.TimeP(&t1).After(t2)).Valid()  // Will return true.
func (validator *ValidatorTimeP) After(value time.Time, template ...string) *ValidatorTimeP {
	_ = "STUB: not implemented"
	return nil
}

// AfterOrEqualTo validates that the time pointer is after or equal to the specified time value.
//
// Usage example:
//
//	t1 := time.Now()
//	t2 := t1
//	Is(v.TimeP(&t1).AfterOrEqualTo(t2)).Valid()  // Will return true.
func (validator *ValidatorTimeP) AfterOrEqualTo(value time.Time, template ...string) *ValidatorTimeP {
	_ = "STUB: not implemented"
	return nil
}

// Before validates that the time pointer is before the specified time value.
//
// Usage example:
//
//	t1 := time.Now()
//	t2 := t1.Add(time.Hour)
//	Is(v.TimeP(&t1).Before(t2)).Valid()  // Will return true.
func (validator *ValidatorTimeP) Before(value time.Time, template ...string) *ValidatorTimeP {
	_ = "STUB: not implemented"
	return nil
}

// BeforeOrEqualTo validates that the time pointer is before or equal to the specified time value.
//
// Usage example:
//
//	t1 := time.Now()
//	t2 := t1
//	Is(v.TimeP(&t1).BeforeOrEqualTo(t2)).Valid()  // Will return true.
func (validator *ValidatorTimeP) BeforeOrEqualTo(value time.Time, template ...string) *ValidatorTimeP {
	_ = "STUB: not implemented"
	return nil
}

// Between validates that the time pointer is between the specified minimum and maximum time values (inclusive).
//
// Usage example:
//
//	t1 := time.Now()
//	min := t1.Add(-time.Hour)
//	max := t1.Add(time.Hour)
//	Is(v.TimeP(&t1).Between(min, max)).Valid()  // Will return true.
func (validator *ValidatorTimeP) Between(min time.Time, max time.Time, template ...string) *ValidatorTimeP {
	_ = "STUB: not implemented"
	return nil
}

// Zero validates that the time pointer is pointing to a zero time value.
//
// Usage example:
//
//	var t *time.Time
//	Is(v.TimeP(t).Zero()).Valid()  // Will return true as t is nil and thus pointing to a zero time.
func (validator *ValidatorTimeP) Zero(template ...string) *ValidatorTimeP {
	_ = "STUB: not implemented"
	return nil
}

// Passing allows for custom validation function to be applied on the time pointer.
//
// Usage example:
//
//	t := time.Now()
//	Is(v.TimeP(&t).Passing(func(v0 *time.Time) bool { return v0.After(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)) })).Valid()  // Custom validation.
func (validator *ValidatorTimeP) Passing(function func(v0 *time.Time) bool, template ...string) *ValidatorTimeP {
	_ = "STUB: not implemented"
	return nil
}

// InSlice validates that the time pointer is pointing to a time value present in the specified slice.
//
// Usage example:
//
//	t := time.Now()
//	validTimes := []time.Time{t, time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)}
//	Is(v.TimeP(&t).InSlice(validTimes)).Valid()  // Will return true.
func (validator *ValidatorTimeP) InSlice(slice []time.Time, template ...string) *ValidatorTimeP {
	_ = "STUB: not implemented"
	return nil
}

// Nil validates that the time pointer is nil.
//
// Usage example:
//
//	var t *time.Time
//	Is(v.TimeP(t).Nil()).Valid()  // Will return true as t is nil.
func (validator *ValidatorTimeP) Nil(template ...string) *ValidatorTimeP {
	_ = "STUB: not implemented"
	return nil
}

// NilOrZero validates that the time pointer is either nil or pointing to a zero time value.
//
// Usage example:
//
//	var t *time.Time
//	Is(v.TimeP(t).NilOrZero()).Valid()  // Will return true as t is nil.
func (validator *ValidatorTimeP) NilOrZero(template ...string) *ValidatorTimeP {
	_ = "STUB: not implemented"
	return nil
}
