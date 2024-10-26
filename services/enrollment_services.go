package services

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	class "github.com/Safiramdhn/project-app-crud-golang-safira/models/classes"
	enrollment "github.com/Safiramdhn/project-app-crud-golang-safira/models/enrollments"
	schedule "github.com/Safiramdhn/project-app-crud-golang-safira/models/schedules"
	student "github.com/Safiramdhn/project-app-crud-golang-safira/models/students"
	"github.com/Safiramdhn/project-app-crud-golang-safira/utils"
)

func GetStudentEnrollment(student_id string) (studentEnroll []enrollment.Enrollments) {
	var enrollments []enrollment.Enrollments

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
		decoder := json.NewDecoder(file)

		if err := decoder.Decode(&enrollments); err != nil && err != io.EOF {
			fmt.Println("decode error: %w", err)
			return nil
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
	go func() {
		if err := saveEnrollments(file, enrollments); err != nil {
			fmt.Println("Error saving enrollments:", err)
			time.Sleep(5 * time.Second)
		}
	}()
}

func GetEnrollment() []enrollment.Enrollments {
	file, err := utils.GetJsonFileName("enrollments")
	if err != nil {
		fmt.Println("Error opening enrollments file:", err)
		return nil
	}
	defer file.Close()

	var enrollments []enrollment.Enrollments
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("File stat error: ", err)
		return nil
	}

	if fileInfo.Size() > 0 {
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&enrollments); err != nil && err != io.EOF {
			fmt.Println("Error decoding enrollments:", err)
			return nil
		}
	}
	return enrollments
}

func CreateEnrollment(studentID, classID string, studentSchedule []schedule.Schedule) {
	newClass := GetStudentClass(classID)

	if len(studentSchedule) > 0 {
		for _, schedule := range studentSchedule {
			if utils.Includes(newClass.Schedules, schedule.Id) {
				fmt.Printf("Notification: Schedule conflict with class at %+v\n", schedule)
				return
			}
		}
	}

	file, err := utils.GetJsonFileName("enrollments")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	enrollments, err := loadEnrollments(file)
	if err != nil {
		fmt.Println("Error loading enrollments:", err)
		return
	}

	student := GetStudentData(studentID)
	schedule := GetOneSchedule(scheduleForm(newClass.Schedules))
	go func() {
		defer file.Close()

		newEnrollment := createNewEnrollment(enrollments, student, newClass, schedule)
		enrollments = append(enrollments, newEnrollment)

		// Encode updated enrollments in a goroutine
		if err := saveEnrollments(file, enrollments); err != nil {
			fmt.Println("Error saving enrollments:", err)
			time.Sleep(5 * time.Second)
		}
	}()
}
func loadEnrollments(file *os.File) ([]enrollment.Enrollments, error) {
	fileInfo, err := file.Stat()
	if err != nil || fileInfo.Size() == 0 {
		return []enrollment.Enrollments{}, err
	}

	var enrollments []enrollment.Enrollments
	if err := json.NewDecoder(file).Decode(&enrollments); err != nil && err != io.EOF {
		return nil, err
	}

	return enrollments, nil
}

// Generates a new enrollment with a unique ID
func createNewEnrollment(enrollments []enrollment.Enrollments, student *student.Student, class class.Class, schedule schedule.Schedule) enrollment.Enrollments {
	id := fmt.Sprintf("E%03d", len(enrollments)+1) // ID with leading zeros if necessary
	return enrollment.Enrollments{
		Id:       id,
		Student:  *student,
		Class:    class,
		Schedule: schedule,
		Status:   "active",
	}
}

// Encodes and saves the enrollments to the file
func saveEnrollments(file *os.File, enrollments []enrollment.Enrollments) error {
	file.Seek(0, 0)  // Go back to the start of the file
	file.Truncate(0) // Clear the file content

	encoder := json.NewEncoder(file)
	return encoder.Encode(&enrollments)
}

func scheduleForm(schdules []string) string {
	schdulesList := GetSchedules()
	var schdulesOption []schedule.Schedule

	for _, schedule := range schdulesList {
		if utils.Includes(schdules, schedule.Id) {
			schdulesOption = append(schdulesOption, schedule)
		}
	}

	schduleJson, err := json.MarshalIndent(schdulesOption, "", " ")
	if err != nil {
		fmt.Println("Marshal error message: ", err)
		return ""
	}
	if len(schduleJson) == 0 {
		fmt.Println("There is no schedule yet")
		return ""
	} else {
		var id string
		fmt.Println(string(schduleJson))
		fmt.Println("Choose schedule")
		fmt.Scan(&id)

		return id
	}
}

func DeleteEnrollment(enrollID string) {
	enrollments := GetEnrollment()
	if enrollments == nil {
		fmt.Println("No enrollments found.")
		return
	}

	// Flag to check if enrollment is found
	var enrollmentFound bool
	// Filter out the enrollment to be marked as deleted
	for i := range enrollments {
		if enrollments[i].Id == enrollID {
			enrollments[i].Status = "deleted"
			enrollmentFound = true
			break
		}
	}
	if !enrollmentFound {
		fmt.Printf("No enrollment found with ID: %s\n", enrollID)
		return
	}

	// Open the enrollments file for writing
	file, err := utils.GetJsonFileName("enrollments")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed after the function completes

	// Save the updated enrollments to the file
	if err := saveEnrollments(file, enrollments); err != nil {
		fmt.Println("Error saving enrollments:", err)
		time.Sleep(5 * time.Second) // Optional: keep this for a delay if necessary
	}
}

func IsAlreadyEnrolled(enrollments []enrollment.Enrollments, classID string) bool {
	for _, enrollment := range enrollments {
		if enrollment.Class.Id == classID {
			return true
		}
	}
	return false
}
