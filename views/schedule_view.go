package views

import (
	"encoding/json"
	"fmt"

	"github.com/Safiramdhn/project-app-crud-golang-safira/services"
	"github.com/Safiramdhn/project-app-crud-golang-safira/utils"
)

func PrintScheduleForm(schdules []string) string {
	schdulesList := services.GetSchedules()

	for _, schedule := range schdulesList {
		if utils.Includes(schdules, schedule.Id) {
			schdulesList = append(schdulesList[:0], schedule)
		}
	}

	schduleJson, err := json.MarshalIndent(schdulesList, "", " ")
	if err != nil {
		fmt.Println("Marshal error message: ", err)
		return ""
	}
	if len(schduleJson) == 0 {
		fmt.Println("There is no schedule yet")
		return ""
	} else {
		var id string
		fmt.Println(string(schduleJson))
		fmt.Println("Choose schedule")
		fmt.Scan(&id)

		return id
	}
}
