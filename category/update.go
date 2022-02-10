package category

import (
	tea "github.com/charmbracelet/bubbletea"
)

// type updateTypeOptions struct {
// 	// New Commit Types that will replace the existing list options
// 	types []commitType
// }

func (state *typeSelection) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// case updateTypeOptions:
	// 	var items []list.Item
	// 	for _, t := range msg.types {
	// 		items = append(items, t)
	// 	}
	// 	return typeSelection{list: makeList(items), choice: state.choice}, state.list.SetItems(items)

	case tea.KeyMsg:
		switch keystroke := msg.String(); keystroke {
		case "ctrl+c":
			return state, tea.Quit

		case "enter":
			// FIXME will panic on empty lists (item nil instead of commit type)
			state.choice = state.list.SelectedItem().(commitType).Git
			// return state, func() tea.Msg { return done{nextView: INPUTS} }
			return state, tea.Quit
		}
	case tea.WindowSizeMsg:
		top, right, bottom, left := docStyle.GetMargin()
		state.list.SetSize(msg.Width-left-right, msg.Height-top-bottom)
	}

	var cmd tea.Cmd
	state.list, cmd = state.list.Update(msg)
	return state, cmd
}
