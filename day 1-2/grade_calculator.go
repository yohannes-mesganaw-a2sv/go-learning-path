package main

import "fmt"

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
	var student Student
	fmt.Println("Enter your name: ")
	for {

		fmt.Scanln(&student.name)

		if student.name == "" {
			fmt.Println("\n Name cannot be Empty. Please enter your name.")
		} else {
			break
		}
	}

	// accept the amoung of subjects
	var subjectCount int
	fmt.Println("Enter the number of subjects: ")
	for {

		_, err := fmt.Scanln(&subjectCount)

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

			fmt.Scanln(&grade.subjectName)

			if grade.subjectName == "" {
				fmt.Println("\n Subject Name cannot be Empty. Please enter the subject name.")
			} else {
				break
			}
		}

		// accept the grade
		fmt.Println("Enter the grade: ")

		for {

			_, err := fmt.Scanln(&grade.gradePoint)

			if err != nil {
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
