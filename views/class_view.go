package views

import (
	"encoding/json"
	"fmt"

	"github.com/Safiramdhn/project-app-crud-golang-safira/services"
)

func printClassList() {
	classes := services.GetClasses()

	classJson, err := json.MarshalIndent(classes, "", " ")
	if err != nil {
		fmt.Println("Marshal error message: ", err)
		return
	}
	if len(classJson) == 0 {
		fmt.Println("There is no classes yet")
	} else {
		fmt.Println(string(classJson))
	}
}
