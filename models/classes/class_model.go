package models

import course "github.com/Safiramdhn/project-app-crud-golang-safira/models/courses"

type Class struct {
	Id          string          `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Type        string          `json:"type"`
	Courses     []course.Course `json:"courses"`
	Instructor  string          `json:"instructor"`
	Schedules   []string        `json:"schedules"`
}
