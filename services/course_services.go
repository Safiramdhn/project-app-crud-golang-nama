package services

import (
	"fmt"

	models "github.com/Safiramdhn/project-app-crud-golang-safira/models/courses"
	"github.com/Safiramdhn/project-app-crud-golang-safira/utils"
)

func GetCourses() (courses []models.Course) {
	// Get courses from json file
	var course models.Course

	file, err := utils.GetJsonFileName("courses")
	if err != nil {
		fmt.Println("Open file error message: ", err)
		return courses
	}
	defer file.Close()

	// Check if the file is empty
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("File stat error: ", err)
		return courses
	}

	if fileInfo.Size() > 0 {
		courses, err = course.JsonDecode(file)
		if err != nil {
			return nil
		}

	}
	return courses
}
