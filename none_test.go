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

func TestNoneIsNotSome(t *testing.T) {
	none := None[struct{}]()

	if none.IsSome() {
		t.Fatalf("None.IsSome() = true, want false")
	}
}

func TestNoneIsNone(t *testing.T) {
	none := None[struct{}]()

	if !none.IsNone() {
		t.Fatalf("None.IsNone() = false, want true")
	}
}
