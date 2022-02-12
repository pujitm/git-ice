package category

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pujitm/git-ice/config"
)

const emptyChoice = ""

// State model to select the type of commit (see conventional commits)
type selection struct {
	list   list.Model // List of git commit types to select from
	choice string     // The chosen git commit type
	hide   bool
}

func (state *selection) Value() string {
	return state.choice
}

func (state *selection) Init() tea.Cmd {
	return nil
	// TODO Load from live source
	// return func() tea.Msg { return updateTypeOptions{types: getCommitTypes()} }
}

func (state *selection) View() string {
	// Hack: Once selection is made, don't show the list view
	// Necessary bc list won't disappear once program quites via tea.Quit See Update()
	if state.hide {
		return ""
	}

	// Let the list do the heavy lifting
	return "\n" + state.list.View()
}

func makeList(items []list.Item) list.Model {
	return list.New(items, selectorDelegate{}, 0, 0)
}

// Creates and returns a Program Model for choosing a git commit category
func Model(types []config.CommitType) *selection {
	var items []list.Item
	for _, t := range types {
		items = append(items, listItem{t})
	}
	list := makeList(items)
	list.Title = "Choose Commit Category"

	return &selection{list: list, choice: emptyChoice}
}
