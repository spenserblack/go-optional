package optional_test

import (
	"fmt"

	"github.com/spenserblack/go-optional"
)

// SurveyAnswers is a collection of answers to a survey.
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

func main() {
	answerSet1 := SurveyAnswers{optional.Some(0)}

	// Sadly, Go's type inference is not yet strong enough to infer T for None()
	answerSet2 := SurveyAnswers{optional.None[int]()}

	fmt.Println(answerSet1.StringifyPetsAnswer())
	fmt.Println(answerSet2.StringifyPetsAnswer())
	// Output:
	// User has 0 pets
	// User declined to answer
}
