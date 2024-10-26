package models

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Student struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *Student) JsonDecode(file *os.File) ([]Student, error) {
	var students []Student
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&students); err != nil && err != io.EOF {
		return nil, fmt.Errorf("decode error: %w", err)
	}

	return students, nil
}
