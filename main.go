package main

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Todo.txt: controls initialization
//------------------------------------------------------------------------------------------

// Define text as const
const (
	buttonNewTaskText       = "NEW"                                                  // Button: ButtonNewTask
	comboBoxFilterText      = "ALL;PRIORITY;PROJECTS;CONTEXTS"                       // ComboBox: ComboBoxFilter
	statusBarText           = "Code by Lugh - Version @test"                         // StatusBar: StatusBar
	groupBoxTaskText        = "TASK"                                                 // GroupBox: GroupBoxTask
	buttonCreditsText       = "#214#"                                                // Button: ButtonCredits
	buttonParametersText    = "#142#"                                                // Button: ButtonParameters
	labelPriorityText       = "PRIORITY:"                                            // Label: LabelPriority
	dropdownBoxPriorityText = "NONE;A;B;C;D;E;F;H;I;J;K;L;M;N;O;P;Q;R;S;T;U;V;X;Y;Z" // DropdownBox: DropdownBoxPriority
	labelProjectsText       = "PROJECTS:"                                            // Label: LabelProjects
	buttonAddProjectText    = "#008#"                                                // Button: ButtonAddProject
	buttonRemoveProjectText = "#143#"                                                // Button: ButtonRemoveProject
	labelContextsText       = "CONTEXTS:"                                            // Label: LabelContexts
	labelTagsText           = "TAGS:"                                                // Label: LabelTags
	buttonRemoveContextText = "#143#"                                                // Button: ButtonRemoveContext
	buttonAddContextText    = "#008#"                                                // Button: ButtonAddContext
	buttonAddTagText        = "#008#"                                                // Button: ButtonAddTag
	buttonRemoveTagText     = "#143#"                                                // Button: ButtonRemoveTag
	buttonTaskRemoveText    = "#143#"                                                // Button: ButtonTaskRemove
	buttonTaskEditText      = "#023#"                                                // Button: ButtonTaskEdit
)

// Define text as var
var (
	toggleOrderText        = ""      // Toggle: ToggleOrder
	labelHelpText          = ""      // Label: LabelHelp
	listTasksTodoText      = ""      // ListView: listTasksTodo
	listTaskProjectsText   = ""      // ListView: ListTaskProjects
	listTaskContextsText   = ""      // ListView: ListTaskContexts
	listTaskTagsText       = ""      // ListView: ListTaskTags
	labelCreatedDateText   = ""      // Label: LabelCreatedDate
	buttonTaskCompleteText = "#112#" // Button: ButtonTaskComplete
	dropdownBoxFilterText  = ""      // DropdownBox: DropdownBoxFilter
)

// Define anchors
var anchorGroupTask = rl.NewVector2(384, 48)
var anchorModalBox = rl.NewVector2(224, 136)

// Define controls variables
var (
	buttonNewTaskPressed              = false // Button: ButtonNewTask
	comboBoxFilterActive        int32 = 0     // ComboBox: ComboBoxFilter
	dropdownBoxFilterEditMode         = false
	dropdownBoxFilterActive     int32 = 0 // dropdownBox: DropdownBoxFilter
	listTasksTodoScrollIndex    int32 = 0
	listTasksTodoActive         int32 = 0     // ListView: listTasksTodo
	buttonCreditsPressed              = false // Button: ButtonCredits
	buttonParametersPressed           = false // Button: ButtonParameters
	toggleOrderActive                 = false // Toggle: ToggleOrder
	dropdownBoxPriorityEditMode       = false
	dropdownBoxPriorityActive   int32 = 0 // DropdownBox: DropdownBoxPriority
	listTaskProjectsScrollIndex int32 = 0
	listTaskProjectsActive      int32 = 0     // ListView: ListTaskProjects
	buttonAddProjectPressed           = false // Button: ButtonAddProject
	buttonRemoveProjectPressed        = false // Button: ButtonRemoveProject
	listTaskContextsScrollIndex int32 = 0
	listTaskContextsActive      int32 = 0 // ListView: ListTaskContexts
	listTaskTagsScrollIndex     int32 = 0
	listTaskTagsActive          int32 = 0     // ListView: ListTaskTags
	buttonRemoveContextPressed        = false // Button: ButtonRemoveContext
	buttonAddContextPressed           = false // Button: ButtonAddContext
	buttonAddTagPressed               = false // Button: ButtonAddTag
	buttonRemoveTagPressed            = false // Button: ButtonRemoveTag
	buttonTaskRemovePressed           = false // Button: ButtonTaskRemove
	buttonTaskCompletePressed         = false // Button: ButtonTaskComplete
	buttonTaskEditPressed             = false // Button: ButtonTaskEdit
)

// Define controls rectangles
var layoutRecs = map[string]rl.Rectangle{
	"buttonNewTask":         rl.NewRectangle(8, 8, 64, 24),                                          // Button: ButtonNewTask
	"comboBoxFilter":        rl.NewRectangle(96, 8, 120, 24),                                        // ComboBox: ComboBoxFilter
	"dropdownBoxFilter":     rl.NewRectangle(224, 8, 120, 24),                                       // DropdownBox: DropdownBoxFilter
	"listTasksTodo":         rl.NewRectangle(8, 48, 368, 376),                                       // ListView: ListTasksTodo
	"statusBar":             rl.NewRectangle(0, 430, 800, 20),                                       // StatusBar: StatusBar
	"groupBoxTask":          rl.NewRectangle(anchorGroupTask.X+0, anchorGroupTask.Y+0, 408, 376),    // GroupBox: GroupBoxTask
	"buttonCredits":         rl.NewRectangle(768, 8, 24, 24),                                        // Button: ButtonCredits
	"buttonParameters":      rl.NewRectangle(736, 8, 24, 24),                                        // Button: ButtonParameters
	"toggleOrder":           rl.NewRectangle(352, 8, 24, 24),                                        // Toggle: ToggleOrder
	"labelPriority":         rl.NewRectangle(anchorGroupTask.X+24, anchorGroupTask.Y+16, 72, 24),    // Label: LabelPriority
	"dropdownBoxPriority":   rl.NewRectangle(anchorGroupTask.X+96, anchorGroupTask.Y+16, 120, 24),   // DropdownBox: DropdownBoxPriority
	"listTaskProjects":      rl.NewRectangle(anchorGroupTask.X+144, anchorGroupTask.Y+56, 216, 88),  // ListView: ListTaskProjects
	"labelProjects":         rl.NewRectangle(anchorGroupTask.X+24, anchorGroupTask.Y+56, 120, 24),   // Label: LabelProjects
	"buttonAddProject":      rl.NewRectangle(anchorGroupTask.X+96, anchorGroupTask.Y+88, 24, 24),    // Button: ButtonAddProject
	"buttonRemoveProject":   rl.NewRectangle(anchorGroupTask.X+96, anchorGroupTask.Y+120, 24, 24),   // Button: ButtonRemoveProject
	"listTaskContexts":      rl.NewRectangle(anchorGroupTask.X+144, anchorGroupTask.Y+160, 216, 88), // ListView: ListTaskContexts
	"listTaskTags":          rl.NewRectangle(anchorGroupTask.X+144, anchorGroupTask.Y+264, 216, 88), // ListView: ListTaskTags
	"labelContexts":         rl.NewRectangle(anchorGroupTask.X+24, anchorGroupTask.Y+160, 120, 24),  // Label: LabelContexts
	"labelTags":             rl.NewRectangle(anchorGroupTask.X+24, anchorGroupTask.Y+264, 120, 24),  // Label: LabelTags
	"buttonRemoveContext":   rl.NewRectangle(anchorGroupTask.X+96, anchorGroupTask.Y+224, 24, 24),   // Button: ButtonRemoveContext
	"buttonAddContext":      rl.NewRectangle(anchorGroupTask.X+96, anchorGroupTask.Y+192, 24, 24),   // Button: ButtonAddContext
	"buttonAddTag":          rl.NewRectangle(anchorGroupTask.X+96, anchorGroupTask.Y+296, 24, 24),   // Button: ButtonAddTags
	"buttonRemoveTag":       rl.NewRectangle(anchorGroupTask.X+96, anchorGroupTask.Y+328, 24, 24),   // Button: ButtonRemoveTags
	"buttonTaskRemove":      rl.NewRectangle(anchorGroupTask.X+360, anchorGroupTask.Y+16, 24, 24),   // Button: ButtonTaskRemove
	"buttonTaskComplete":    rl.NewRectangle(anchorGroupTask.X+328, anchorGroupTask.Y+16, 24, 24),   // Button: ButtonTaskComplete
	"buttonTaskEdit":        rl.NewRectangle(anchorGroupTask.X+296, anchorGroupTask.Y+16, 24, 24),   // Button: ButtonTaskEdit
	"labelHelp":             rl.NewRectangle(384, 430, 416, 20),                                     // Label: LabelHelp
	"labelCreatedDate":      rl.NewRectangle(anchorGroupTask.X+304, anchorGroupTask.Y+360, 104, 16), // Label: LabelCreatedDate
	"modalMessageBoxDelete": rl.NewRectangle(anchorModalBox.X, anchorModalBox.Y, 344, 152),
	"modalMessageBoxText":   rl.NewRectangle(anchorModalBox.X-57, anchorModalBox.Y, 457, 152),
}

//------------------------------------------------------------------------------------------

func main() {

	// Initialization
	//--------------------------------------------------------------------------------------
	const screenWidth = 800
	const screenHeight = 450

	var messageBoxDeleteTask = false
	var messageBoxEditTask = false
	var messageBoxNewTask = false
	var messageBoxTasktext string
	var messageBoxSecret = true

	rl.SetConfigFlags(rl.FlagWindowHighdpi) // Active high dpi support

	rl.InitWindow(screenWidth, screenHeight, "Todo.txt")
	defer rl.CloseWindow() // Close window and OpenGL context

	rl.SetTargetFPS(60)

	initTaskList()
	defer saveTaskList()

	// Main App loop
	for !rl.WindowShouldClose() { // Detect window close button or ESC key
		// Update
		//----------------------------------------------------------------------------------
		// update toggleOrderText
		if toggleOrderActive {
			toggleOrderText = "#121#"
		} else {
			toggleOrderText = "#120#"
		}
		// update labelHelpText
		labelHelpText = getHelpText()
		// update contents
		updateTodoText()
		updateTaskProjects()
		updateTaskContexts()
		updateTaskTags()
		updateTaskPriority()
		updateTaskCreatedDate()
		updateTaskCompleteButton()
		updateDropdownBoxFilter()
		//----------------------------------------------------------------------------------

		// Draw
		//----------------------------------------------------------------------------------
		rl.BeginDrawing()

		rl.ClearBackground(rl.GetColor(uint(gui.GetStyle(gui.DEFAULT, gui.BACKGROUND_COLOR))))

		// raygui: controls drawing
		//----------------------------------------------------------------------------------
		// Draw controls
		if dropdownBoxFilterEditMode || dropdownBoxPriorityEditMode || messageBoxDeleteTask || messageBoxNewTask || messageBoxEditTask {
			gui.Lock()
		}

		buttonNewTaskPressed = gui.Button(layoutRecs["buttonNewTask"], buttonNewTaskText)
		if buttonNewTaskPressed {
			messageBoxTasktext = ""
			messageBoxNewTask = true
		}
		comboBoxFilterActive = gui.ComboBox(layoutRecs["comboBoxFilter"], comboBoxFilterText, comboBoxFilterActive)
		listTasksTodoActive = gui.ListView(layoutRecs["listTasksTodo"], listTasksTodoText, &listTasksTodoScrollIndex, listTasksTodoActive)
		gui.StatusBar(layoutRecs["statusBar"], statusBarText)
		gui.GroupBox(layoutRecs["groupBoxTask"], groupBoxTaskText)
		buttonCreditsPressed = gui.Button(layoutRecs["buttonCredits"], buttonCreditsText)
		buttonParametersPressed = gui.Button(layoutRecs["buttonParameters"], buttonParametersText)
		toggleOrderActive = gui.Toggle(layoutRecs["toggleOrder"], toggleOrderText, toggleOrderActive)
		gui.Label(layoutRecs["labelPriority"], labelPriorityText)
		listTaskProjectsActive = gui.ListView(layoutRecs["listTaskProjects"], listTaskProjectsText, &listTaskProjectsScrollIndex, listTaskProjectsActive)
		gui.Label(layoutRecs["labelProjects"], labelProjectsText)
		buttonAddProjectPressed = gui.Button(layoutRecs["buttonAddProject"], buttonAddProjectText)
		buttonRemoveProjectPressed = gui.Button(layoutRecs["buttonRemoveProject"], buttonRemoveProjectText)
		listTaskContextsActive = gui.ListView(layoutRecs["listTaskContexts"], listTaskContextsText, &listTaskContextsScrollIndex, listTaskContextsActive)
		listTaskTagsActive = gui.ListView(layoutRecs["listTaskTags"], listTaskTagsText, &listTaskTagsScrollIndex, listTaskTagsActive)
		gui.Label(layoutRecs["labelContexts"], labelContextsText)
		gui.Label(layoutRecs["labelTags"], labelTagsText)
		buttonRemoveContextPressed = gui.Button(layoutRecs["buttonRemoveContext"], buttonRemoveContextText)
		buttonAddContextPressed = gui.Button(layoutRecs["buttonAddContext"], buttonAddContextText)
		buttonAddTagPressed = gui.Button(layoutRecs["buttonAddTag"], buttonAddTagText)
		buttonRemoveTagPressed = gui.Button(layoutRecs["buttonRemoveTag"], buttonRemoveTagText)
		buttonTaskRemovePressed = gui.Button(layoutRecs["buttonTaskRemove"], buttonTaskRemoveText)
		if buttonTaskRemovePressed {
			messageBoxDeleteTask = true
		}
		buttonTaskCompletePressed = gui.Button(layoutRecs["buttonTaskComplete"], buttonTaskCompleteText)
		if buttonTaskCompletePressed {
			modifyTaskState()
		}
		buttonTaskEditPressed = gui.Button(layoutRecs["buttonTaskEdit"], buttonTaskEditText)
		if buttonTaskEditPressed {
			messageBoxTasktext = getOriginalTaskText()
			messageBoxEditTask = true
		}
		gui.Label(layoutRecs["labelHelp"], labelHelpText)
		gui.Label(layoutRecs["labelCreatedDate"], labelCreatedDateText)
		if gui.DropdownBox(layoutRecs["dropdownBoxFilter"], dropdownBoxFilterText, &dropdownBoxFilterActive, dropdownBoxFilterEditMode) {
			dropdownBoxFilterEditMode = !dropdownBoxFilterEditMode
		}
		if gui.DropdownBox(layoutRecs["dropdownBoxPriority"], dropdownBoxPriorityText, &dropdownBoxPriorityActive, dropdownBoxPriorityEditMode) {
			dropdownBoxPriorityEditMode = !dropdownBoxPriorityEditMode
		}

		gui.Unlock()

		// messageBox logic
		if messageBoxDeleteTask {
			selectOption := gui.MessageBox(layoutRecs["modalMessageBoxDelete"], "#143#Delete Task", "Are you sure to delete this task ?", "Yes;No")
			if selectOption >= 0 {
				if selectOption == 1 {
					deleteSelectedTask()
				}
				messageBoxDeleteTask = false
			}
		}
		if messageBoxNewTask {
			selectOption := gui.TextInputBox(layoutRecs["modalMessageBoxText"], "#191#New Task", "Enter new task. It can be in raw Todo.txt format, it will be parsing automatically.", "Add;Cancel", &messageBoxTasktext, 255, &messageBoxSecret)
			if selectOption >= 0 {
				if selectOption == 1 {
					addNewTask(messageBoxTasktext)
				}
				messageBoxNewTask = false
			}
		}
		if messageBoxEditTask {
			selectOption := gui.TextInputBox(layoutRecs["modalMessageBoxText"], "#191#Edit Task", "Edit the selected task. It can be in raw Todo.txt format, it will be parsing automatically.", "Save;Cancel", &messageBoxTasktext, 255, &messageBoxSecret)
			if selectOption >= 0 {
				if selectOption == 1 {
					setSelectedTask(messageBoxTasktext)
				}
				messageBoxEditTask = false
			}
		}

		//----------------------------------------------------------------------------------

		rl.EndDrawing()
		//----------------------------------------------------------------------------------
	}

	// De-Initialization
	//--------------------------------------------------------------------------------------
	// rl.CloseWindow() // Close window and OpenGL context
	//--------------------------------------------------------------------------------------

}
