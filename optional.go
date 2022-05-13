// Package optional provides the Optional type
package optional

// Optional wraps a type that may or may not exist.
type Optional[T any] interface {
	// Unwrap will get the optional type. Should only be used if there is 100%
	// certainty that the T exists, or when a panic is acceptable.
	Unwrap() T
	// UnwrapSafe acts like Unwrap, but will return an error instead of panicking.
	UnwrapSafe() (T, error)
}

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
