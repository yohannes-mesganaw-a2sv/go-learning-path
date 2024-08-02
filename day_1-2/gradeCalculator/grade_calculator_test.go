// grade calculator test
package main

import (
	"bufio"
	"errors"
	"strings"
	"testing"
)

// Test containsSpecialCharacterOrNumber function
func TestContainsSpecialCharacterOrNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"John", false},
		{"John123", true},
		{"John!", true},
		{"John Doe", false},
		{"", false},
	}

	for _, test := range tests {
		result := containsSpecialCharacterOrNumber(test.input)
		if result != test.expected {
			t.Errorf("containsSpecialCharacterOrNumber(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

// Test calculateAverageGrade method
func TestCalculateAverageGrade(t *testing.T) {
	tests := []struct {
		grades   []Grade
		expected float64
		err      error
	}{
		{[]Grade{{"Math", 80}, {"Science", 90}}, 85.0, nil},
		{[]Grade{}, 0.0, errors.New("No grades found")},
	}

	for _, test := range tests {
		student := Student{grades: test.grades}
		result, err := student.calculateAverageGrade()
		if result != test.expected || (err != nil && err.Error() != test.err.Error()) {
			t.Errorf("calculateAverageGrade() = %v, %v; want %v, %v", result, err, test.expected, test.err)
		}
	}
}

// Helper function to create a new reader
func newReader(input string) *bufio.Reader {
	return bufio.NewReader(strings.NewReader(input))
}

// Test newStudent function
func TestNewStudent(t *testing.T) {
	input := "John Doe\n"
	reader := newReader(input)
	student := newStudent(reader)

	if student.name != "John Doe" {
		t.Errorf("newStudent() = %v; want %v", student.name, "John Doe")
	}
}

// Test populateGrade function
func TestPopulateGrade(t *testing.T) {
	input := "2\nMath\n85\nScience\n90\n"
	reader := newReader(input)
	student := &Student{}
	populateGrade(student, reader)

	expectedGrades := []Grade{
		{"Math", 85.0},
		{"Science", 90.0},
	}

	if len(student.grades) != len(expectedGrades) {
		t.Fatalf("populateGrade() length = %d; want %d", len(student.grades), len(expectedGrades))
	}

	for i, grade := range student.grades {
		if grade.subjectName != expectedGrades[i].subjectName || grade.gradePoint != expectedGrades[i].gradePoint {
			t.Errorf("populateGrade() grade = %v; want %v", grade, expectedGrades[i])
		}
	}
}
