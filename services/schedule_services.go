package services

import (
	"encoding/json"
	"fmt"
	"io"

	enrollment "github.com/Safiramdhn/project-app-crud-golang-safira/models/enrollments"
	schedule "github.com/Safiramdhn/project-app-crud-golang-safira/models/schedules"
	"github.com/Safiramdhn/project-app-crud-golang-safira/utils"
)

func GetSchedules() (schedules []schedule.Schedule) {
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
		decoder := json.NewDecoder(file)

		if err := decoder.Decode(&schedules); err != nil && err != io.EOF {
			fmt.Println("decode error: %w", err)
			return nil
		}

	}
	return schedules
}

func GetOneSchedule(schedule_id string) (schedule schedule.Schedule) {
	schedules := GetSchedules()

	for _, scheduleData := range schedules {
		if scheduleData.Id == schedule_id {
			return scheduleData
		}
	}
	return schedule
}

func GetStudentSchedule(enrollments []enrollment.Enrollments) []schedule.Schedule {
	var schedules []schedule.Schedule
	for _, enrollment := range enrollments {
		schedules = append(schedules, enrollment.Schedule)
	}
	return schedules
}
