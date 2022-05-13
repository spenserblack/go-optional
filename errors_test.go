package optional

import "testing"

func TestNoneUnwrapErr(t *testing.T) {
	want := "Unwrap on None"

	if actual := ErrNoneUnwrap.Error(); actual != want {
		t.Fatalf(`ErrNoneUnwrap.Error() = %q, want %q`, actual, want)
	}
}
