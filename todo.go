package main

import (
	"log"
	"strings"

	todo "github.com/1set/todotxt"
)

var tasks todo.TaskList

// Initialize the task list
func initTaskList() {
	tasks = todo.NewTaskList()
	var err error
	tasks, err = todo.LoadFromPath("todo.txt")
	if err != nil {
		log.Println("Error:", err)
	}
}

func getOriginalTaskText() string {
	selected := listTasksTodoActive
	if selected == -1 {
		return ""
	}
	return tasks[selected].Original
}

func setSelectedTask(nTask string) {
	selected := listTasksTodoActive
	if selected == -1 {
		return
	}
	task, err := todo.ParseTask(nTask)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	tasks[selected] = *task
}

func updateTodoText() {
	var text string
	for id, task := range tasks {
		if id == 0 {
			text = task.Todo
			continue
		}
		text += ";"
		text += task.Todo
	}
	listTasksTodoText = text
}

func updateTaskProjects() {
	var text string
	selected := listTasksTodoActive
	if selected == -1 {
		listTaskProjectsText = ""
		return
	}
	for id, project := range tasks[selected].Projects {
		if id == 0 {
			text = project
			continue
		}
		text += ";"
		text += project
	}
	listTaskProjectsText = text
}

func updateTaskContexts() {
	var text string
	selected := listTasksTodoActive
	if selected == -1 {
		listTaskContextsText = ""
		return
	}
	for id, context := range tasks[selected].Contexts {
		if id == 0 {
			text = context
			continue
		}
		text += ";"
		text += context
	}
	listTaskContextsText = text
}

func updateTaskTags() {
	var text string
	selected := listTasksTodoActive
	if selected == -1 {
		listTaskTagsText = ""
		return
	}
	id := 0
	for key, tag := range tasks[selected].AdditionalTags {
		if id == 0 {
			text = key + ": " + tag
			id++
			continue
		}
		text += ";"
		text += key + ": " + tag
		id++
	}
	listTaskTagsText = text
}

func updateTaskPriority() {
	selected := listTasksTodoActive
	if selected == -1 {
		dropdownBoxPriorityActive = 0
		return
	}
	if !tasks[selected].HasPriority() {
		dropdownBoxPriorityActive = 0
		return
	}
	stringRef := strings.Split(dropdownBoxPriorityText, ";")
	for id, ref := range stringRef {
		if tasks[selected].Priority == ref {
			dropdownBoxPriorityActive = int32(id)
			break
		}
	}
}

func updateTaskCreatedDate() {
	selected := listTasksTodoActive
	if selected == -1 {
		labelCreatedDateText = ""
		return
	}
	if !tasks[selected].HasCreatedDate() {
		labelCreatedDateText = ""
		return
	}
	date := tasks[selected].CreatedDate
	labelCreatedDateText = date.String()
}

func updateTaskCompleteButton() {
	selected := listTasksTodoActive
	if selected == -1 {
		buttonTaskCompleteText = "#112#"
		return
	}
	if tasks[selected].IsCompleted() {
		buttonTaskCompleteText = "#113#"
	} else {
		buttonTaskCompleteText = "#112#"
	}
}

func modifyTaskState() {
	selected := listTasksTodoActive
	if selected == -1 {
		return
	}
	if tasks[selected].IsCompleted() {
		tasks[selected].Reopen()
	} else {
		tasks[selected].Complete()
	}
}

func deleteSelectedTask() {
	selected := listTasksTodoActive
	if selected == -1 {
		return
	}
	id := tasks[selected].ID
	err := tasks.RemoveTaskByID(id)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	listTasksTodoActive = -1
}

func addNewTask(nTask string) {
	task, err := todo.ParseTask(nTask)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	tasks.AddTask(task)
}

func getAllPriority() string {
	var text string
	for _, task := range tasks {
		if len(text) == 0 && task.HasPriority() {
			text = task.Priority
		}
		if task.HasPriority() && !strings.Contains(text, task.Priority) {
			text += ";"
			text += task.Priority
		}
	}
	return text
}

func getAllProjects() string {
	var text string
	var projects []string
	for _, task := range tasks {
		if task.HasProjects() {
			projects = append(projects, task.Projects...)
		}
	}
	for _, project := range projects {
		if len(text) == 0 {
			text = project
		}
		if !strings.Contains(text, project) {
			text += ";"
			text += project
		}
	}
	return text
}

func getAllContexts() string {
	var text string
	var contexts []string
	for _, task := range tasks {
		if task.HasContexts() {
			contexts = append(contexts, task.Contexts...)
		}
	}
	for _, context := range contexts {
		if len(text) == 0 {
			text = context
		}
		if !strings.Contains(text, context) {
			text += ";"
			text += context
		}
	}
	return text
}

func updateDropdownBoxFilter() {
	switch comboBoxFilterActive {
	case 0:
		dropdownBoxFilterText = "ALL;COMPLETED;NOT_COMPLETED"
	case 1:
		dropdownBoxFilterText = getAllPriority()
	case 2:
		dropdownBoxFilterText = getAllProjects()
	case 3:
		dropdownBoxFilterText = getAllContexts()
	}
}

// Save the taskList on disk
func saveTaskList() {
	err := tasks.WriteToPath("todo.txt")
	if err != nil {
		log.Fatalln("Error:", err)
	}
}
