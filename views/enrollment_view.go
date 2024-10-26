package views

import (
	"encoding/json"
	"fmt"

	enrollment "github.com/Safiramdhn/project-app-crud-golang-safira/models/enrollments"
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

	studentClasses := services.InitiateEnrolledClass(studentEnroll)

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

func printDeleteEnrollForm(student_id string) {
	var id string
	printStudentEnrollment(student_id)
	fmt.Println("Type enroll id to delete: ")
	fmt.Scan(&id)

	go services.DeleteEnrollment(id)
	utils.PromptContinue("delete enrollment")

}

func printDeactiveEnrollment(student_id string) {
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
		utils.PromptReturnToMainMenu()
	}
}
