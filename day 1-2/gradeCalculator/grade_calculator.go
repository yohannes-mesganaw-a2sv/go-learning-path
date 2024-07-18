package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	name string
	// create a grades struct
	grades []Grade
}

type Grade struct {
	subjectName string
	gradePoint  float64
}

// calcualte the average grade
func (student Student) calculateAverageGrade() float64 {
	totalGrade := 0.0

	for _, grade := range student.grades {
		totalGrade += grade.gradePoint
	}

	return totalGrade / float64(len(student.grades))

}

func main() {
	// bufio.NewReader is used when we want to input a full line.
	reader := bufio.NewReader(os.Stdin)

	var student Student
	fmt.Println("Enter your name: ")
	for {

		fmt.Scanln(&student.name)
		line, err := reader.ReadString('\n')
		student.name = line

		if student.name == "" || err != nil {
			fmt.Println("\n Name cannot be Empty. Please enter your name.")
		} else {
			break
		}
	}

	// accept the amoung of subjects
	var subjectCount int
	var numberInput string
	fmt.Println("Enter the number of subjects: ")
	for {

		_, err := fmt.Scan(&numberInput)
		subjectCount, err = strconv.Atoi(numberInput)

		if err != nil {
			fmt.Println("\nInvalid Input. Please enter a number.")
		} else {
			break

		}
	}

	// loop through the subjects
	for i := 0; i < subjectCount; i++ {
		// create a grade object
		var grade Grade
		// accept the subject name
		fmt.Println("Enter the subject name: ")
		for {

			line, err := reader.ReadString('\n')

			if grade.subjectName == "" || err != nil {
				fmt.Println("\n Subject Name cannot be Empty. Please enter the subject name.")
			} else {
				break
			}
			grade.subjectName = line

		}

		// accept the grade
		fmt.Println("Enter the grade: ")

		for {

			_, err := fmt.Scan(&numberInput)
			grade.gradePoint, err = strconv.ParseFloat(numberInput, 64)

			// check if the grade point is a number and also with in the range of 0 - 100
			if err != nil || grade.gradePoint > 100 || grade.gradePoint < 0 {
				fmt.Println("\nInvalid Input. Please enter a number.")
			} else {
				break

			}
		}

		student.grades = append(student.grades, grade)

	}

	print("-----------------------------------")
	// print the student details
	fmt.Printf("Student Name: %s\n", student.name)
	fmt.Println("Grades: ")
	for _, grade := range student.grades {
		fmt.Printf("Subject: %s, Grade: %.2f\n", grade.subjectName, grade.gradePoint)
	}

	// calculate the average grade
	averageGrade := student.calculateAverageGrade()
	fmt.Printf("The average grade for %s is %.2f\n", student.name, averageGrade)
	print("-----------------------------------")

}
