package services

import (
	"context"
	"errors"
	"fmt"

	"time"

	models "github.com/Safiramdhn/project-app-crud-golang-safira/models/students"
	"github.com/Safiramdhn/project-app-crud-golang-safira/utils"
)

func LoginService(student_id, password string) (context.Context, error) {
	// get student by id
	student := GetStudentData(student_id)
	type key string

	// password validation
	// if not match, throw error invalid password
	if student.Password != password {
		return nil, errors.New("Invalid password, please try again")
	}

	// create context with value student id
	// add timeout 60s
	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()
	sessionCtx := context.WithValue(ctxWithTimeout, key("student_id"), student_id)

	// return context and error nil
	return sessionCtx, nil
}

func GetStudentData(student_id string) models.Student {

	// get student data from json file
	var studentData models.Student
	file, err := utils.GetJsonFileName("students")
	if err != nil {
		fmt.Println("Open file error message: ", err)
		return studentData
	}
	defer file.Close()

	// Check if the file is empty
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("File stat error: ", err)
		return studentData
	}

	if fileInfo.Size() > 0 {
		students, err := studentData.JsonDecode(file)
		if err != nil {
			return studentData
		}

		for _, student := range students {
			if student.Id == student_id {
				studentData = student
			}
		}
	}
	return studentData
}
