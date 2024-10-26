package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func GetJsonFileName(file_req string) (*os.File, error) {
	var filePath string

	switch file_req {
	case "students":
		filePath = filepath.Join("models", file_req, "students.json")
	case "courses":
		filePath = filepath.Join("models", file_req, "courses.json")
	case "schedules":
		filePath = filepath.Join("models", file_req, "schedules.json")
	case "enrollments":
		filePath = filepath.Join("models", file_req, "enrollments.json")
	case "classes":
		filePath = filepath.Join("models", file_req, "classes.json")
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

func PromptContinue(action string) bool {
	fmt.Println("---------")
	fmt.Printf("%s/back to main menu? (y/n)\n", action)
	var response string
	fmt.Scan(&response)
	return response == "y"
}

func PromptReturnToMainMenu() {
	fmt.Println("---------")
	fmt.Println("Back to main menu? (y/n)")
	var response string
	fmt.Scan(&response)
	if response != "y" {
		os.Exit(0)
	}
	ClearScreen()
}

func Includes(slice []string, value string) bool {
	for _, v := range slice {
		if strings.ToLower(v) == strings.ToLower(value) {
			return true
		}
	}
	return false
}
