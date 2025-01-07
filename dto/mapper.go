package dto

type (
	Mapper[I, O any] func(I) O
)

func (convert Mapper[I, O]) PtrO() func(I) *O {
	return func(input I) *O {
		output := convert(input)
		return &output
	}
}

func (convert Mapper[I, O]) PtrI() func(*I) O {
	return func(input *I) O {
		if input != nil {
			return convert(*input)
		}
		var zero O
		return zero
	}
}

func (convert Mapper[I, O]) PtrIO() func(*I) *O {
	return func(input *I) *O {
		if input != nil {
			return convert.PtrO()(*input)
		}
		return nil
	}
}

// PointerFunc returns a mapper function (T->*T) which returns the allocated address of a T type value.
func PointerFunc[T any]() func(T) *T {
	return func(t T) *T { return &t }
}

// IndirectFunc returns a mapper function (*T->T) which indirect a non-nil value of type *T to its value.
// In case of nil input and safe mode returns the zero value of T, else panics.
func IndirectFunc[T any](safe bool) func(*T) T {
	return func(t *T) T {
		if safe {
			if t == nil {
				var zero T
				return zero
			}
		}
		return *t
	}
}
