package models

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

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

func (e *Enrollments) JsonDecode(file *os.File) ([]Enrollments, error) {
	var enrollments []Enrollments
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&enrollments); err != nil && err != io.EOF {
		return nil, fmt.Errorf("decode error: %w", err)
	}

	return enrollments, nil
}

func (e *Enrollments) JsonEncode(file *os.File, enrolls []Enrollments) error {
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(&enrolls); err != nil {
		return fmt.Errorf("encode error: %w", err)
	}
	return nil
}
