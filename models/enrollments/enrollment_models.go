package models

import (
	class "github.com/Safiramdhn/project-app-crud-golang-safira/models/classes"
	schedule "github.com/Safiramdhn/project-app-crud-golang-safira/models/schedules"
	student "github.com/Safiramdhn/project-app-crud-golang-safira/models/students"
)

type Enrollments struct {
	Id       string            `json:"id"`
	Student  student.Student   `json:"student"`
	Class    class.Class       `json:"class"`
	Schedule schedule.Schedule `json:"schedule"`
	Status   string            `json:"status"`
}

type EnrolledClass struct {
	Id         string
	Class_id   string
	Title      string
	Class_type string
	Instructor string
	Schedule   schedule.Schedule
}
