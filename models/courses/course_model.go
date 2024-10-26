package models

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Course struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (co *Course) JsonDecode(file *os.File) ([]Course, error) {
	var courses []Course
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&courses); err != nil && err != io.EOF {
		return nil, fmt.Errorf("decode error: %w", err)
	}

	return courses, nil
}
