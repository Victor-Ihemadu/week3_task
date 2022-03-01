package school

type Student struct {
	StudentName string
	Grades      []int
	Courses     string
	Age         int
}

func (s *Student) TakeCourse(courseOffered string) string {
	if courseOffered == s.Courses {
		return "You have been offered provisional admission for " + courseOffered
	}
	return "Course is not available"
}
