# go-optional

[![CI](https://github.com/spenserblack/go-optional/actions/workflows/ci.yml/badge.svg)](https://github.com/spenserblack/go-optional/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/spenserblack/go-optional/branch/main/graph/badge.svg?token=ce0g4TeUBX)](https://codecov.io/gh/spenserblack/go-optional)

An `Optional` type for Go.

## Description

Function return signatures like `(T, error)` do a great job of enforcing that
errors are checked. Yet, even with an `error` return value, sometimes it can
be hard to differentiate between a non-existent `T` and a zero-value `T`.

Take this example:

```go
type SurveyAnswers struct {
	// Pets is the answer to the question "how many pets do you have?"
	//
	// Optional question.
	Pets int
}
```

The zero value of `int` is `0`. So how can we tell if the person who took the
survey answered `0`, or if they simply declined to answer?

One way is to make `Pets` type `*int`, so that it can be `nil`. But this opens
us up to nil-pointer errors. That isn't so bad in this example, but it can get
tricky if the user of your library is expected to dereference the pointer, or
if you can return a nil interface.

Instead of having values that may be nil, you can have an optional containing
that value. This makes it necessary for anyone using your library to
acknowledge that the value is optional, and take some action to handle the
optional value.

### Example

```go
type SurveyAnswers struct {
	// Pets is the answer to the question "how many pets do you have?"
	//
	// Optional question.
	Pets optional.Optional[int]
}

func (s SurveyAnswers) StringifyPetsAnswer() string {
	if s.Pets.IsNone() {
		return "User declined to answer"
	}
	return fmt.Sprintf("User has %d pets", s.Pets.Unwrap())
}
```
