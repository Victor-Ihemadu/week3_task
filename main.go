package main

import (
	"fmt"
	"school_assignment/school"
)

var inst = school.Teacher{
	StudentName:     "Victor",
	TeachersCourses: []string{"Physics", "Chemistry", "Peter"},
}

func main() {
	fmt.Println(inst.GradeStudent("Physics"))
	s1 := school.Principal{"Daniel", []int{75, 69, 89, 90}, 11}
	s2 := school.Principal{"Peter", []int{60, 68, 92, 98}, 12}
	s3 := school.Principal{"John", []int{65, 86, 89, 87}, 14}
	s4 := school.Principal{"Paul", []int{90, 36, 45, 77}, 12}
	s5 := school.Principal{"Anthony", []int{64, 88, 75, 100}, 11}
	s6 := school.Principal{"Lawrence", []int{98, 43, 86, 75}, 10}
	s7 := school.Principal{"Andrew", []int{88, 75, 90, 65}, 16}
	s8 := school.Principal{"Hilary", []int{87, 90, 54, 44}, 15}
	fmt.Println(s1.AdmitStudent(12))
	fmt.Println(s2.AdmitStudent(12))
	fmt.Println(s3.AdmitStudent(12))
	fmt.Println(s4.AdmitStudent(12))
	fmt.Println(s5.AdmitStudent(12))
	fmt.Println(s6.AdmitStudent(12))
	fmt.Println(s7.AdmitStudent(12))
	fmt.Println(s8.AdmitStudent(12))
	fmt.Println(s1.ExpelStudent([]int{75, 69, 89, 90}))
	fmt.Println(s2.ExpelStudent([]int{60, 68, 92, 98}))
	fmt.Println(s3.ExpelStudent([]int{65, 86, 89, 87}))
	fmt.Println(s4.ExpelStudent([]int{90, 36, 45, 77}))
	fmt.Println(s5.ExpelStudent([]int{64, 88, 75, 100}))
	fmt.Println(s6.ExpelStudent([]int{98, 43, 86, 75}))
	fmt.Println(s7.ExpelStudent([]int{75, 69, 89, 90}))
	fmt.Println(s8.ExpelStudent([]int{87, 90, 54, 44}))
}
