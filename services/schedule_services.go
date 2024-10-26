package services

import (
	"fmt"

	models "github.com/Safiramdhn/project-app-crud-golang-safira/models/schedules"
	"github.com/Safiramdhn/project-app-crud-golang-safira/utils"
)

func GetSchedules() (schedules []models.Schedule) {
	var schedule models.Schedule

	file, err := utils.GetJsonFileName("schedules")
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
		schedules, err = schedule.JsonDecode(file)
		if err != nil {
			return nil
		}
	}
	return schedules
}
