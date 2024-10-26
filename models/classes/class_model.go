package models

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	course "github.com/Safiramdhn/project-app-crud-golang-safira/models/courses"
	schedule "github.com/Safiramdhn/project-app-crud-golang-safira/models/schedules"
)

type Class struct {
	Id          string              `json:"id"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	Type        string              `json:"type"`
	Courses     []course.Course     `json:"courses"`
	Instructor  string              `json:"instructor"`
	Schedules   []schedule.Schedule `json:"schedules"`
}

func (c *Class) JsonDecode(file *os.File) ([]Class, error) {
	var classes []Class
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&classes); err != nil && err != io.EOF {
		return nil, fmt.Errorf("decode error: %w", err)
	}

	return classes, nil
}
