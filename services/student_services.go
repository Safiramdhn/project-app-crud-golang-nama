package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	models "github.com/Safiramdhn/project-app-crud-golang-safira/models/students"
	"github.com/Safiramdhn/project-app-crud-golang-safira/utils"
)

func LoginService(studentID, password string, ctx context.Context) (timeoutCtx context.Context, cancelFunc context.CancelFunc) {
	student := GetStudentData(studentID)
	if student == nil {
		// fmt.Println("Student not found")
		return nil, nil
	}

	// Validate password
	if student.Password != password {
		fmt.Println("Invalid password, please try again")
		return nil, nil
	}

	// Create a context with student ID and timeout
	timeoutCtx, cancelFunc = context.WithTimeout(ctx, 20*time.Second)

	return timeoutCtx, cancelFunc
}

func GetStudentData(studentID string) *models.Student {
	file, err := utils.GetJsonFileName("students")
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return nil
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("File stat error:", err)
		return nil
	}

	// Proceed if file has content
	if fileInfo.Size() > 0 {
		var students []models.Student
		decoder := json.NewDecoder(file)

		// Try decoding as array first
		if err := decoder.Decode(&students); err != nil {
			fmt.Println("Decode error:", err)
			return nil
		}

		// Search for the student
		for i := range students {
			if students[i].Id == studentID {
				return &students[i]
			}
		}
	}

	fmt.Println("Student not found for ID:", studentID)
	return nil
}
