package school

import "testing"

func TestPrincipal_AdmitStudent(t *testing.T) {
	admitTest := []struct {
		input1    Principal
		inputfunc int
		output    string
	}{
		{Principal{StudentName: "Victor", Grades: []int{}, Age: 17}, 12, "Victor is admitted"},
		{Principal{StudentName: "Peter", Grades: []int{}, Age: 14}, 12, "Peter is not admitted"},
		{Principal{StudentName: "James", Grades: []int{}, Age: 11}, 12, "James is admitted"},
		{Principal{StudentName: "Andrew", Grades: []int{}, Age: 13}, 12, "Andrew is not admitted"},
		{Principal{StudentName: "Paul", Grades: []int{}, Age: 10}, 12, "Paul is admitted"},
	}
	for _, value := range admitTest {
		ans := value.input1.AdmitStudent(value.inputfunc)
		if ans != value.output {
			t.Errorf("Expected %v, got %v", value.output, ans)
		}
	}
}

func TestPrincipal_ExpelStudent(t *testing.T) {
	expelTest := []struct {
		input0    Principal
		inputFunc []int
		output    string
	}{
		{Principal{"Basel", []int{}, 16}, []int{89, 67, 78, 90}, "Basel is Promoted"},
		{Principal{"Gabriel", []int{}, 11}, []int{59, 97, 78, 80}, "Gabriel is Promoted"},
		{Principal{"Tony", []int{}, 10}, []int{56, 67, 50, 60}, "Tony is Expelled"},
		{Principal{"Emmanuel", []int{}, 12}, []int{80, 97, 60, 97}, "Emmanuel is Promoted"},
		{Principal{"Pascal", []int{}, 13}, []int{58, 60, 73, 40}, "Pascal is Expelled"},
	}

	for _, value := range expelTest {
		answer := value.input0.ExpelStudent(value.inputFunc)

		if answer != value.output {
			t.Errorf("Expected %v, got %v", value.output, answer)
		}
	}
}
