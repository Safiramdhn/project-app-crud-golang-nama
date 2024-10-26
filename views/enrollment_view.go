package views

import (
	"encoding/json"
	"fmt"

	enrollment "github.com/Safiramdhn/project-app-crud-golang-safira/models/enrollments"
	"github.com/Safiramdhn/project-app-crud-golang-safira/services"
	"github.com/Safiramdhn/project-app-crud-golang-safira/utils"
)

func printStudentEnrollment(student_id string) {
	// for {
	utils.ClearScreen()
	studentEnroll := services.GetStudentEnrollment(student_id)
	if studentEnroll == nil {
		fmt.Println("You not enrolled in any classes")
		if !utils.PromptContinue("show class") {
			utils.ClearScreen()
			return
		}
	}

	studentClasses := services.InitiateEnrolledClass(studentEnroll)

	studentClassJson, err := json.MarshalIndent(studentClasses, "", " ")
	if err != nil {
		fmt.Println("Marshal error message: ", err)
		return
	}
	if len(studentClassJson) == 0 {
		fmt.Println("There is no classes yet")
		if !utils.PromptContinue("show class") {
			utils.ClearScreen()
			return
		}
	}
	fmt.Println(string(studentClassJson))
}

func printDeleteEnrollForm(student_id string) {
	for {
		var id string
		printStudentEnrollment(student_id)
		fmt.Println("Type enroll id to delete: ")
		fmt.Scan(&id)

		services.DeleteEnrollment(id)
		if !utils.PromptContinue("delete enrollment") {
			utils.ClearScreen()
			return
		}
	}

}

func printDeactiveEnrollment(student_id string) {
	utils.ClearScreen()
	var deactiveEnroll []enrollment.Enrollments
	enrollments := services.GetEnrollment()

	for _, enroll := range enrollments {
		if enroll.Student.Id == student_id && enroll.Status == "deleted" {
			deactiveEnroll = append(deactiveEnroll, enroll)
		}
	}

	deactiveEnrollJson, err := json.MarshalIndent(deactiveEnroll, "", " ")
	if err != nil {
		fmt.Println("Marshal error message: ", err)
		return
	}
	if len(deactiveEnrollJson) == 0 {
		fmt.Println("There is no deactive enroll yet")
	} else {
		fmt.Println(string(deactiveEnrollJson))
	}
	utils.PromptReturnToMainMenu()
}
