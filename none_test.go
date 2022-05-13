package optional

import "testing"

func TestNoneUnwrapPanics(t *testing.T) {
	none := None[struct{}]()

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Did not panic")
		}
	}()

	none.Unwrap()
}

func TestNoneUnwrapSafeErr(t *testing.T) {
	none := None[struct{}]()

	if _, err := none.UnwrapSafe(); err != ErrNoneUnwrap {
		t.Fatalf(`err = %v, want ErrNoneUnwrap`, err)
	}
}
