package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func GetJsonFileName(file_req string) (*os.File, error) {
	var filePath string

	switch file_req {
	case "students":
		filePath = filepath.Join("models", "students", "students.json")
	case "courses":
		filePath = filepath.Join("models", "courses", "courses.json")
	case "schedules":
		filePath = filepath.Join("models", "schedules", "schedules.json")
	case "enrollments":
		filePath = filepath.Join("models", "enrollments", "enrollments.json")
	default:
		fmt.Println("Invalid json file name")
	}
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)
	return file, err
}

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
