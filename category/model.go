package category

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

const emptyChoice = ""

// State model to select the type of commit (see conventional commits)
type typeSelection struct {
	list   list.Model
	choice string
}

func (state *typeSelection) Value() string {
	return state.choice
}

func (state *typeSelection) Init() tea.Cmd {
	return nil
	// TODO Load from live source
	// return func() tea.Msg { return updateTypeOptions{types: getCommitTypes()} }
}

func (state *typeSelection) View() string {
	// Hack: Once selection is made, don't show the list view
	// Necessary bc list won't disappear once program quites via tea.Quit See Update()
	if state.choice != emptyChoice {
		return ""
	}

	// Let the list do the heavy lifting
	return "\n" + state.list.View()
}

func makeList(items []list.Item) list.Model {
	return list.New(items, selectorDelegate{}, 0, 0)
}

// Creates and returns a Program Model for choosing a git commit category
func Model() *typeSelection {
	var items []list.Item
	for _, t := range getCommitTypes() {
		items = append(items, t)
	}
	list := makeList(items)
	list.Title = "Choose Commit Category"

	return &typeSelection{list: list, choice: emptyChoice}
}

func RunPrompt() (string, error) {
	commitCategory := Model()
	program := tea.NewProgram(commitCategory)
	if err := program.Start(); err != nil {
		return "", err
	}
	return commitCategory.Value(), nil
}
