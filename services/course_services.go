package services

import (
	"encoding/json"
	"fmt"
	"io"

	models "github.com/Safiramdhn/project-app-crud-golang-safira/models/courses"
	"github.com/Safiramdhn/project-app-crud-golang-safira/utils"
)

func GetCourses() []models.Course {
	// Get courses from json file
	var courses []models.Course

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
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&courses); err != nil && err != io.EOF {
			fmt.Println("decode error: %w", err)
			return nil
		}

	}
	return courses
}
