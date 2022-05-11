package optional

// NoneUnwrap is an error that occurs when None is unwrapped.
type NoneUnwrap struct{}

// ErrNoneUnwrap is the static NoneUnwrap error.
var ErrNoneUnwrap error = NoneUnwrap{}

func (NoneUnwrap) Error() string {
	return "Unwrap on None"
}
