package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var helpLabelText = map[string]string{
	"buttonNewTask":       "Create new task",
	"comboBoxFilter":      "Change the filter mode",
	"dropdownBoxFilter":   "Select the filter",
	"toggleOrder":         "Change the sorting order",
	"dropdownBoxPriority": "Modify priority's task",
	"buttonTaskEdit":      "Edit the raw todo",
	"buttonTaskComplete":  "Complete or Reopen the task",
	"buttonTaskRemove":    "Delete the task",
	"buttonAddProject":    "Add a project to the selected task",
	"buttonRemoveProject": "Delete the selected project",
	"buttonAddContext":    "Add a context to the selected task",
	"buttonRemoveContext": "Delete the selected context",
	"buttonAddTag":        "Add a tag to the selected task",
	"buttonRemoveTag":     "Delete the selected tag",
	"labelCreatedDate":    "Created date",
}

// Check if the specific element is over by mouse
func overElement(r rl.Rectangle) bool {
	var over bool = rl.CheckCollisionPointRec(rl.GetMousePosition(), r)
	return over
}

// Return the actual's key of over element
func getOverElement() (key string) {
	for key, rec := range layoutRecs {
		switch key {
		case "groupBoxTask":
			continue
		case "modalMessageBoxDelete":
			continue
		case "modalMessageBoxText":
			continue
		default:
			over := overElement(rec)
			if over {
				return key
			}
		}
	}

	return ""
}

func getHelpText() string {
	key := getOverElement()
	symbol := "#191#"
	text, ok := helpLabelText[key]
	if ok {
		return fmt.Sprintf("%s %s", symbol, text)
	}
	return ""
}
