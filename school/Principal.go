package school

type Principal struct {
	StudentName string
	Grades      []int
	Age         int
}

func (s *Principal) AdmitStudent(Age int) string {

	if s.Age <= Age {
		return s.StudentName + " is admitted"
	}
	return s.StudentName + " is not admitted"
}

func (g *Principal) ExpelStudent(Scores []int) string {
	var sum int = 0
	for i := 0; i < len(Scores); i++ {
		sum += Scores[i]

		if sum/4 >= 70 {
			return g.StudentName + " is Promoted"
		}
	}
	return g.StudentName + " is Expelled"
}
