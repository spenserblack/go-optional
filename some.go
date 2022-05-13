package optional

type some[T any] struct {
	value T
}

// Some is the type for an optional with an existing value.
func Some[T any](value T) Optional[T] {
	return some[T]{value}
}

func (o some[T]) Unwrap() T {
	return o.value
}

func (o some[T]) UnwrapSafe() (T, error) {
	return o.value, nil
}
