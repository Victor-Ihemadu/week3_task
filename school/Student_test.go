package school

import "testing"

func TestStudent_TakeCourse(t *testing.T) {
	takeCourseTest := []struct {
		firstInput Student
		inputFunc  string
		result     string
	}{
		{Student{"Bernard", []int{}, "Chemistry", 11}, "Chemistry", "You have been offered provisional admission for Chemistry"},
		{Student{"Bernard", []int{}, "Physics", 11}, "Chemistry", "Course is not available"},
		{Student{"Bernard", []int{}, "Agriculture", 11}, "Chemistry", "Course is not available"},
		{Student{"Bernard", []int{}, "Computer", 11}, "Computer", "You have been offered provisional admission for Computer"},
		{Student{"Bernard", []int{}, "Mathematics", 11}, "Engineering", "Course is not available"},
		{Student{"Bernard", []int{}, "Economics", 11}, "PHE", "Course is not available"},
		{Student{"Bernard", []int{}, "Business Studies", 11}, "Business Studies", "You have been offered provisional admission for Business Studies"},
	}
	for _, val := range takeCourseTest {
		offeredCourse := val.firstInput.TakeCourse(val.inputFunc)

		if offeredCourse > val.result {
			t.Errorf("Expected %v, got %v", val.result, offeredCourse)
		}
	}
}
