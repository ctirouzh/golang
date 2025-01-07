package dto

// Address returns the address of an input value.
func Address[T any](input T) *T { return &input }

// Indirect makes indirect a non-nil value of type *T to its value.
// In case of nil input and safe mode returns the zero value of T, else panics.
func Indirect[T any](input *T, safe bool) T {
	if safe && input == nil {
		var zero T
		return zero
	}
	return *input
}
