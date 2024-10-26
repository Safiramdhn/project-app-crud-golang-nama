package views

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Safiramdhn/project-app-crud-golang-safira/services"
	"github.com/Safiramdhn/project-app-crud-golang-safira/utils"
)

func LoginView() {
	for {
		utils.ClearScreen()
		var student_id, password string
		fmt.Println("Welcome to Choco Academy")
		fmt.Println("Please login with your student ID and password")
		fmt.Println("Student ID: ")
		fmt.Scan(&student_id)
		fmt.Println("Password: ")
		fmt.Scan(&password)

		if student_id == "" || password == "" {
			fmt.Println("Student ID and password are required")
			time.Sleep(5 * time.Second)
			continue
		}
		session, err := services.LoginService(student_id, password)
		if err != nil {
			fmt.Println("Login error message", err)
		}
		select {
		case <-session.Done():
			fmt.Println("Session expired, please login again.")
			time.Sleep(5 * time.Second)
			continue
		default:
			DashboardMenu(session)
		}
	}
}

func DashboardMenu(ctx context.Context) {
	for {
		utils.ClearScreen()
		var option int
		var logout string
		// display student name

		// dashboard menu
		fmt.Println("1. Enrollments")
		fmt.Println("2. Edit enrollments schedule")
		fmt.Println("3. Add Class")
		fmt.Println("4. Delete Enrollment")
		fmt.Println("5. Enrollment History")
		fmt.Print("99. Logout")
		fmt.Scan(&option)
		studentId := ctx.Value("student_id").(string)

		switch option {
		case 1:
			// display student enrollment
			printStudentEnrollment(studentId)
		case 2:
			// edit student class schedule
			printEditClassForm(studentId)
		case 3:
			// add class
			printAddClassForm(studentId)
		case 4:
			// delete enrollment
			printDeleteForm(studentId)
		case 5:
			printDeactiveEnrollment(studentId)
		case 99:
			fmt.Println("Are you sure want to logout? (y/n)")
			fmt.Scan(&logout)
			if strings.ToLower(logout) != "y" {
				continue
			}
			return
		default:
			fmt.Println("Invalid option")
			time.Sleep(5 * time.Second)
			continue
		}
	}
}
