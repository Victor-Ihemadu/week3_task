package school

type Teacher struct {
	StudentName       string
	TeachersCourses   []string
	AttendanceChecker int
}

func (a *Teacher) ValidateAttendance(Attendance int) string {
	if a.AttendanceChecker < Attendance {
		return a.StudentName + " didn't fully attend classes"
	}
	return a.StudentName + " fully attended classes"
}

func (c *Teacher) GradeStudent(courseSubmitted string) string {
	for _, value := range c.TeachersCourses {
		if value == courseSubmitted {
			return c.StudentName + " is graded"
		}
	}
	return c.StudentName + " has no grade assigned"
}
