package services

import (
	"encoding/json"
	"fmt"
	"io"

	class "github.com/Safiramdhn/project-app-crud-golang-safira/models/classes"
	enrollment "github.com/Safiramdhn/project-app-crud-golang-safira/models/enrollments"
	"github.com/Safiramdhn/project-app-crud-golang-safira/utils"
)

func GetClasses() []class.Class {
	var classes []class.Class

	file, err := utils.GetJsonFileName("classes")
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
		decoder := json.NewDecoder(file)

		if err := decoder.Decode(&classes); err != nil && err != io.EOF {
			fmt.Println("Decode error: %w", err)
			return nil
		}
	}
	return classes
}

func GetStudentClass(class_id string) (studenClass class.Class) {
	clasess := GetClasses()

	for _, class := range clasess {
		if class.Id == class_id {
			return class
		}
	}
	return studenClass
}

func InitiateEnrolledClass(studentEnroll []enrollment.Enrollments) []enrollment.EnrolledClass {
	var studentClasses []enrollment.EnrolledClass
	for _, enroll := range studentEnroll {
		studentClass := enrollment.EnrolledClass{
			Id:         enroll.Id,
			Class_id:   enroll.Class.Id,
			Title:      enroll.Class.Title,
			Class_type: enroll.Class.Type,
			Instructor: enroll.Class.Instructor,
			Schedule:   enroll.Schedule,
		}
		studentClasses = append(studentClasses, studentClass)
	}

	return studentClasses
}
