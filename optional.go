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
