package services

import (
	"fmt"

	models "github.com/Safiramdhn/project-app-crud-golang-safira/models/enrollments"
	"github.com/Safiramdhn/project-app-crud-golang-safira/utils"
)

func GetStudentEnrollment(student_id string) (studentEnroll []models.Enrollments) {
	var enroll models.Enrollments

	file, err := utils.GetJsonFileName("enrollments")
	if err != nil {
		fmt.Println("Open file error message: ", err)
		return studentEnroll
	}
	defer file.Close()

	// Check if the file is empty
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("File stat error: ", err)
		return studentEnroll
	}

	if fileInfo.Size() > 0 {
		enrollments, err := enroll.JsonDecode(file)
		if err != nil {
			return studentEnroll
		}

		for _, enrollment := range enrollments {
			if enrollment.Student.Id == student_id {
				studentEnroll = append(studentEnroll, enrollment)
			}
		}
	}
	return studentEnroll
}

func EditEnrollmentSchedule(enroll_id, schedule_id string) {
	enrollments := getEnrollment()

	for i := range enrollments {
		if enrollments[i].Id == enroll_id {
			schedule := GetOneSchedule(schedule_id)
			enrollments[i].Schedule = schedule
			break
		}
	}
}

func getEnrollment() (enrollments []models.Enrollments) {
	var enrollment models.Enrollments

	file, err := utils.GetJsonFileName("enrollments")
	if err != nil {
		fmt.Println("Open file error message: ", err)
		return nil
	}
	defer file.Close()

	// Check if the file is empty
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("File stat error: ", err)
		return nil
	}

	if fileInfo.Size() > 0 {
		enrollments, err = enrollment.JsonDecode(file)
		if err != nil {
			return nil
		}
	}
	return enrollments
}
