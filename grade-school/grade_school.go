// Package school ...
package school

import (
	"sort"
)

// Grade holds the grade and its students for a school
type Grade struct {
	grade    int
	students []string
}

// School holds the grades and its students
type School []Grade

// New returns a new school
func New() *School {
	var s School = []Grade{}
	return &s
}

// Add adds a student to a grade
func (s *School) Add(student string, grade int) {
	mySchool := *s
	g := Grade{grade: grade, students: []string{student}}
	for i, n := range *s {
		if n.grade == grade {
			n.students = append(n.students, student)
			mySchool[i].students = n.students
			return
		}
	}
	*s = append(*s, g)
	return
}

// Grade returns the students in a grade
func (s *School) Grade(g int) []string {
	var grade []string
	for _, n := range *s {
		if n.grade == g {
			grade = n.students
		}
	}

	return sort.StringSlice(grade)
}

// Enrollment return the roster of the school
func (s *School) Enrollment() []Grade {
	school := *s
	sort.SliceStable(school, func(i, j int) bool {
		return school[i].grade < school[j].grade
	})
	school = *s
	for i, n := range *s {
		sort.Strings(n.students)
		school[i].students = n.students
	}
	return school
}
