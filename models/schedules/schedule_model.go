package models

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type Schedule struct {
	Id   string    `json:"id"`
	Day  string    `json:"day"`
	Time time.Time `json:"time"`
}

func (sch *Schedule) JsonDecode(file *os.File) ([]Schedule, error) {
	var schedules []Schedule
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&schedules); err != nil && err != io.EOF {
		return nil, fmt.Errorf("decode error: %w", err)
	}

	return schedules, nil
}
