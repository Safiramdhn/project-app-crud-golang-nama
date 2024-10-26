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

		ctx := context.Background()
		timeoutCtx, cancel := services.LoginService(student_id, password, ctx)
		if timeoutCtx == nil {
			fmt.Println("Login failed. Please check your credentials.")
			time.Sleep(5 * time.Second)
			continue
		}

		defer cancel() // Cancel when the session ends

		select {
		case <-timeoutCtx.Done():
			fmt.Println("Session Timeout, please login again")
			time.Sleep(5 * time.Second)
			continue
		default:
			DashboardMenu(student_id)
		}
	}
}

func DashboardMenu(student_id string) {
	for {
		utils.ClearScreen()
		var option int
		var logout string
		// display student name

		// dashboard menu
		fmt.Println("Dashboard")
		fmt.Println("1. Enrollments")
		fmt.Println("2. Edit enrollments schedule")
		fmt.Println("3. Add Class")
		fmt.Println("4. Delete Enrollment")
		fmt.Println("5. Enrollment History")
		fmt.Println("99. Logout")
		fmt.Print("Choose Option:")
		fmt.Scan(&option)

		switch option {
		case 1:
			// display student enrollment
			printStudentEnrollment(student_id)
			utils.PromptReturnToMainMenu()
		case 2:
			// edit student class schedule
			printEditClassForm(student_id)
		case 3:
			// add class
			printAddClassForm(student_id)
		case 4:
			// delete enrollment
			printDeleteEnrollForm(student_id)
		case 5:
			printDeactiveEnrollment(student_id)
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
