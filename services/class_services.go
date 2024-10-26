package services

import (
	"fmt"

	models "github.com/Safiramdhn/project-app-crud-golang-safira/models/classes"
	"github.com/Safiramdhn/project-app-crud-golang-safira/utils"
)

func GetClasses() (classes []models.Class) {
	var class models.Class

	file, err := utils.GetJsonFileName("class")
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
		classes, err = class.JsonDecode(file)
		if err != nil {
			return nil
		}
	}
	return classes
}

func GetStudentClass(class_id string) (studenClass models.Class) {
	clasess := GetClasses()

	for _, class := range clasess {
		if class.Id == class_id {
			return class
		}
	}
	return studenClass
}
