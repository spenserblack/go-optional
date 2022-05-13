package optional

type none[T any] struct{}

// None is the none type, signifying that the T doesn't exist.
func None[T any]() Optional[T] {
	return none[T]{}
}

func (none[T]) Unwrap() T {
	panic("Unwrap on None")
}

func (none[T]) UnwrapSafe() (T, error) {
	var zero T
	return zero, ErrNoneUnwrap
}
