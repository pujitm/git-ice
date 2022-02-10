package category

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

// State model to select the type of commit (see conventional commits)
type typeSelection struct {
	list   list.Model
	choice string
}

func (state typeSelection) Init() tea.Cmd {
	return nil
	// TODO Load from live source
	// return func() tea.Msg { return updateTypeOptions{types: getCommitTypes()} }
}

func (state typeSelection) View() string {
	if state.choice != "" {
		return state.choice
	}
	return "\n" + state.list.View()
	// return docStyle.Render(state.list.View())
}

func makeList(items []list.Item) list.Model {
	return list.New(items, selectorDelegate{}, 0, 0)
}

// Creates and returns a Program Model for choosing a git commit category
func Model() tea.Model {
	var items []list.Item
	for _, t := range getCommitTypes() {
		items = append(items, t)
	}
	list := makeList(items)
	list.Title = "Choose Commit Category"

	return typeSelection{list: list, choice: ""}
}
