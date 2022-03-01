package school

type Applicant struct {
	ApplicantName string
	Grade         int
}

func (n *Applicant) SchoolApplication(grade int) string {
	if n.Grade < grade {
		return n.ApplicantName + " did not meet up the cut-off mark"
		//} else if n.Grade >= 40 && n.Age > age {
		//	return n.ApplicantName + " has been placed on probation trial"
	}
	return n.ApplicantName + " has been offered provisional admission into Ivy League School"
}
