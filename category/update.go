package category

import (
	tea "github.com/charmbracelet/bubbletea"
)

// type exitFunc func(state *selection) (tea.Model, tea.Cmd)

func (state *selection) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	exit := func(state *selection) (tea.Model, tea.Cmd) {
		state.hide = true
		return state, tea.Quit
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keystroke := msg.String(); keystroke {
		case "ctrl+c", "q":
			return exit(state)

		case "enter":
			selected := state.list.SelectedItem()
			if selected != nil {
				state.choice = selected.(listItem).CommitType.Git
			}

			return exit(state)
		}

	case tea.WindowSizeMsg:
		top, right, bottom, left := docStyle.GetMargin()
		state.list.SetSize(msg.Width-left-right, msg.Height-top-bottom)
	}

	var cmd tea.Cmd
	state.list, cmd = state.list.Update(msg)
	return state, cmd
}
