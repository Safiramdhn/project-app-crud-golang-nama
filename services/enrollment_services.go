package services

import (
	"encoding/json"
	"fmt"

	enrollment "github.com/Safiramdhn/project-app-crud-golang-safira/models/enrollments"
	schedule "github.com/Safiramdhn/project-app-crud-golang-safira/models/schedules"
	"github.com/Safiramdhn/project-app-crud-golang-safira/utils"
)

func GetStudentEnrollment(student_id string) (studentEnroll []enrollment.Enrollments) {
	var enroll enrollment.Enrollments

	file, err := utils.GetJsonFileName("enrollments")
	if err != nil {
		fmt.Println("Open file error message: ", err)
		return studentEnroll
	}
	defer file.Close()

	// Check if the file is empty
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("File stat error: ", err)
		return studentEnroll
	}

	if fileInfo.Size() > 0 {
		enrollments, err := enroll.JsonDecode(file)
		if err != nil {
			return studentEnroll
		}

		for _, enrollment := range enrollments {
			if enrollment.Student.Id == student_id || enrollment.Status == "active" {
				studentEnroll = append(studentEnroll, enrollment)
			}
		}
	}
	return studentEnroll
}

func EditEnrollmentSchedule(enroll_id, class_id, schedule_id string) {
	var enroll enrollment.Enrollments
	enrollments := GetEnrollment()
	scheduleUpdated := false

	for i := range enrollments {
		if enrollments[i].Id == enroll_id && class_id == enrollments[i].Class.Id {
			schedule := GetOneSchedule(schedule_id)
			enrollments[i].Schedule = schedule
			scheduleUpdated = true
			break
		}

	}
	if !scheduleUpdated {
		fmt.Println("Notification: Schedule not updated, please check your enroll and class id again")
		return
	}

	file, err := utils.GetJsonFileName("enrollments")
	if err != nil {
		fmt.Println("Open file error message: ", err)
		return
	}
	defer file.Close()
	go enroll.JsonEncode(file, enrollments)
}

func GetEnrollment() (enrollments []enrollment.Enrollments) {
	var enrollment enrollment.Enrollments

	file, err := utils.GetJsonFileName("enrollments")
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
		enrollments, err = enrollment.JsonDecode(file)
		if err != nil {
			return nil
		}
		for _, enrollment := range enrollments {
			if enrollment.Status == "active" {
				enrollments = append(enrollments[:0], enrollment)
			}
		}
	}

	return enrollments
}

func CreateEnrollment(student_id, class_id string, student_schedule []schedule.Schedule) {
	var enroll enrollment.Enrollments
	var schdule_id string
	newClass := GetStudentClass(class_id)

	for i := 0; i < len(student_schedule); i++ {
		if utils.Includes(newClass.Schedules, student_schedule[i].Id) {
			fmt.Printf("Notification: You have classes in %+v", student_schedule[i])
			return
		} else {
			schdule_id = scheduleForm(newClass.Schedules)
		}
	}

	file, err := utils.GetJsonFileName("enrollments")
	if err != nil {
		fmt.Println("Open file error message: ", err)
		return
	}
	defer file.Close()

	// Check if the file is empty
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("File stat error: ", err)
		return
	}

	if fileInfo.Size() > 0 {
		enrollments, err := enroll.JsonDecode(file)
		if err != nil {
			return
		}

		if len(enrollments) == 0 {
			enroll = enrollment.Enrollments{
				Id:       "E001",
				Student:  GetStudentData(student_id),
				Class:    newClass,
				Schedule: GetOneSchedule(schdule_id),
				Status:   "active",
			}
		} else {
			countEnroll := len(enrollments)
			var num string
			if countEnroll > 9 {
				num = fmt.Sprintf("%d", countEnroll) // No leading zero needed
			} else {
				num = fmt.Sprintf("00%d", countEnroll) // Leading zeros for single digits
			}

			enroll = enrollment.Enrollments{
				Id:       fmt.Sprintf("E%s", num), // Use Sprintf to format the ID
				Student:  GetStudentData(student_id),
				Class:    newClass,
				Schedule: GetOneSchedule(schdule_id),
				Status:   "active",
			}
		}
		enrollments = append(enrollments, enroll)
		go enroll.JsonEncode(file, enrollments)
	}
}

func scheduleForm(schdules []string) string {
	schdulesList := GetSchedules()

	for _, schdule := range schdulesList {
		if utils.Includes(schdules, schdule.Id) {
			schdulesList = append(schdulesList[:0], schdule)
		}
	}

	schduleJson, err := json.MarshalIndent(schdulesList, "", " ")
	if err != nil {
		fmt.Println("Marshal error message: ", err)
		return ""
	}
	if len(schduleJson) == 0 {
		fmt.Println("There is no schdule yet")
		return ""
	} else {
		var id string
		fmt.Println(string(schduleJson))
		fmt.Println("Choose schedule")
		fmt.Scan(&id)

		return id
	}
}

func DeleteEnrollment(enroll_id string) {
	var enroll enrollment.Enrollments
	var newEnroll []enrollment.Enrollments
	enrollments := GetEnrollment()

	for _, enrollment := range enrollments {
		if enrollment.Id != enroll_id {
			newEnroll = append(newEnroll, enrollment)
		}
	}

	file, err := utils.GetJsonFileName("enrollments")
	if err != nil {
		fmt.Println("Open file error message: ", err)
		return
	}
	defer file.Close()
	go enroll.JsonEncode(file, enrollments)
	return
}
