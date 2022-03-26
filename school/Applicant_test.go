package school

import "testing"

func TestApplicant_SchoolApplication(t *testing.T) {
	reviewApplication := []struct {
		input0     Applicant
		firstInput int
		output     string
	}{
		{Applicant{"Cynthia", 15}, 40, "Cynthia has been offered provisional admission into Ivy League School"},
		{Applicant{"Stella", 35}, 40, "Stella did not meet up the cut-off mark"},
		{Applicant{"Jonathan", 89}, 40, "Jonathan has been offered provisional admission into Ivy League School"},
		{Applicant{"Cynthia", 55}, 40, "Cynthia has been offered provisional admission into Ivy League School"},
		{Applicant{"Thompson", 33}, 40, "Thompson did not meet up the cut-off mark"},
		{Applicant{"Peterson", 30}, 40, "Peterson has been offered provisional admission into Ivy League School"},
		{Applicant{"Mary", 20}, 40, "Mary did not meet up the cut-off mark"},
	}
	for _, val := range reviewApplication {
		application := val.input0.SchoolApplication(val.firstInput)

		if application != val.output {
			t.Errorf("Expected %v, got %v", val.output, application)
		}
	}

}
