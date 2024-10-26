package views

import (
	"encoding/json"
	"fmt"

	schedule "github.com/Safiramdhn/project-app-crud-golang-safira/models/schedules"
	"github.com/Safiramdhn/project-app-crud-golang-safira/services"
	"github.com/Safiramdhn/project-app-crud-golang-safira/utils"
)

func printClassList() {
	classes := services.GetClasses()

	classJson, err := json.MarshalIndent(classes, "", " ")
	if err != nil {
		fmt.Println("Marshal error message: ", err)
		return
	}
	if len(classJson) == 0 {
		fmt.Println("There is no classes yet")
	} else {
		fmt.Println(string(classJson))
	}
}

func printEditClassForm(student_id string) {
	for {
		utils.ClearScreen()
		var schedule_id, class_id, enroll_id string

		printStudentEnrollment(student_id)
		fmt.Println("Choose your enroll id: ")
		fmt.Scan(&enroll_id)
		fmt.Println("Write your class id: ")
		fmt.Scan(&class_id)

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
			go services.EditEnrollmentSchedule(enroll_id, studentClass.Id, schedule_id)
			if !utils.PromptContinue("edit class") {
				utils.ClearScreen()
				return
			}
			continue
		}
	}
}

func printAddClassForm(student_id string) {
	for {
		var class_id string
		var studentSchedule []schedule.Schedule
		utils.ClearScreen()
		printClassList()
		fmt.Println("Choose class ID")
		fmt.Scan(&class_id)

		studentEnroll := services.GetStudentEnrollment(student_id)

		for _, enroll := range studentEnroll {
			if enroll.Class.Id == class_id {
				fmt.Println("You already enrolled this class")
				break
			}
			studentSchedule = append(studentSchedule, enroll.Schedule)
		}
		go services.CreateEnrollment(student_id, class_id, studentSchedule)

		if !utils.PromptContinue("add class") {
			utils.ClearScreen()
			return
		}
	}
}
