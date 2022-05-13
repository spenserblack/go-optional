package optional

import "testing"

func TestSomeUnwrapOk(t *testing.T) {
	tests := []struct {
		o    Optional[int]
		want int
	}{
		{Some(123), 123},
		{Some(456), 456},
	}

	for _, tt := range tests {
		if v := tt.o.Unwrap(); v != tt.want {
			t.Errorf(`%v.Unwrap() = %v, want %v`, tt.o, v, tt.want)
		}
	}

}

func TestSomeUnwrapSafeOk(t *testing.T) {
	tests := []struct {
		o    Optional[int]
		want int
	}{
		{Some(123), 123},
		{Some(456), 456},
	}

	for _, tt := range tests {
		v, err := tt.o.UnwrapSafe()
		if err != nil {
			t.Errorf(`err = %v, want nil`, err)
		}
		if v != tt.want {
			t.Errorf(`%v.Unwrap() = %v, want %v`, tt.o, v, tt.want)
		}
	}
}

func TestSomeIsSome(t *testing.T) {
	some := Some(struct{}{})

	if !some.IsSome() {
		t.Fatalf("Some.IsSome() = false, want true")
	}
}

func TestSomeIsNotNone(t *testing.T) {
	some := Some(struct{}{})

	if some.IsNone() {
		t.Fatalf("Some.IsNone() = true, want false")
	}
}
