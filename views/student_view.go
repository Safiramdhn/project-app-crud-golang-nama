package views

import (
	"encoding/json"
	"fmt"

	schedule "github.com/Safiramdhn/project-app-crud-golang-safira/models/schedules"
	"github.com/Safiramdhn/project-app-crud-golang-safira/services"
	"github.com/Safiramdhn/project-app-crud-golang-safira/utils"
)

func printStudentEnrollment(student_id string) {
	utils.ClearScreen()
	studentEnroll := services.GetStudentEnrollment(student_id)
	if studentEnroll == nil {
		fmt.Println("You not enrolled in any classes")
		return
	}

	var studentClasses []struct {
		id         string
		class_id   string
		title      string
		class_type string
		instructor string
		schedule   schedule.Schedule
	}
	for _, enroll := range studentEnroll {
		studentClass := struct {
			id         string
			class_id   string
			title      string
			class_type string
			instructor string
			schedule   schedule.Schedule
		}{
			id:         enroll.Id,
			class_id:   enroll.Class.Id,
			title:      enroll.Class.Title,
			class_type: enroll.Class.Type,
			instructor: enroll.Class.Instructor,
			schedule:   enroll.Schedule,
		}
		studentClasses = append(studentClasses, studentClass)
	}

	studentClassJson, err := json.MarshalIndent(studentClasses, "", " ")
	if err != nil {
		fmt.Println("Marshal error message: ", err)
		return
	}
	if len(studentClassJson) == 0 {
		fmt.Println("There is no classes yet")
	} else {
		fmt.Println(string(studentClassJson))
		utils.PromptReturnToMainMenu()
	}
}

func printEditClassForm(enroll_id, class_id string) {
	utils.ClearScreen()
	var schedule_id string

	studentClass := services.GetStudentClass(class_id)
	studentClassJson, err := json.MarshalIndent(studentClass, "", " ")
	if err != nil {
		fmt.Println("Marshal error message: ", err)
		return
	}
	if len(studentClassJson) == 0 {
		fmt.Println("There is no classes yet")
	} else {
		fmt.Println(string(studentClassJson))
		fmt.Scan(&schedule_id)

		if schedule_id == "" || utils.Includes(studentClass.Schedules, schedule_id) {
			fmt.Println("Invalid schedule id")
			return
		}
		go services.EditEnrollmentSchedule(enroll_id, schedule_id)
	}
}
