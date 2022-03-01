package school

import "testing"

func TestTeacher_ValidateAttendance(t *testing.T) {
	checkAttendance := []struct {
		inputOne      Teacher
		inputFunction int
		output        string
	}{
		{Teacher{"Lawrence", []string{"Mathematics", "PHE", "Government"}, 69}, 65, "Lawrence fully attended classes"},
		{Teacher{"Prince", []string{"Physics", "Agriculture", "Social Studies"}, 80}, 65, "Prince fully attended classes"},
		{Teacher{"Oscar", []string{"Mathematics", "Technical Drawing", "Government"}, 50}, 65, "Oscar didn't fully attend classes"},
		{Teacher{"Johnson", []string{"English", "PHE", "Economics"}, 95}, 65, "Johnson fully attended classes"},
		{Teacher{"Charles", []string{"Music", "PHE", "Government"}, 45}, 65, "Charles didn't fully attend classes"},
		{Teacher{"Michael", []string{"Civic Education", "PHE", "Mathematics"}, 63}, 65, "Michael didn't fully attend classes"},
		{Teacher{"Harry", []string{"Government", "PHE", ""}, 71}, 65, "Harry fully attended classes"},
	}
	for _, value := range checkAttendance {
		attendance := value.inputOne.ValidateAttendance(value.inputFunction)

		if attendance != value.output {
			t.Errorf("Expected %v, got %v", value.output, attendance)
		}
	}
}

func TestTeacher_GradeStudent(t *testing.T) {
	markExam := []struct {
		firstInput     Teacher
		funcInput      string
		expectedOutput string
	}{
		{Teacher{"Harry", []string{"Agriculture", "PHE", "Music"}, 71}, "Agriculture", "Harry is graded"},
		{Teacher{"Michael", []string{"Civic Education", "PHE", "Mathematics"}, 63}, "Chemistry", "Michael has no grade assigned"},
		{Teacher{"Charles", []string{"Music", "PHE", "Government"}, 45}, "Mathematics", "Charles has no grade assigned"},
		{Teacher{"Johnson", []string{"English", "PHE", "Economics"}, 95}, "Economics", "Johnson is graded"},
		{Teacher{"Oscar", []string{"Mathematics", "Chemistry", "Government"}, 50}, "Chemistry", "Oscar is graded"},
		{Teacher{"Prince", []string{"Physics", "Agriculture", "Social Studies"}, 80}, "Computer", "Prince has no grade assigned"},
		{Teacher{"Lawrence", []string{"Mathematics", "PHE", "Government"}, 69}, "French", "Lawrence has no grade assigned"},
	}
	for _, val := range markExam {
		grade := val.firstInput.GradeStudent(val.funcInput)

		if grade != val.expectedOutput {
			t.Errorf("Expected %v, got %v", val.expectedOutput, grade)
		}
	}
}
