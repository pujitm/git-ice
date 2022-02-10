// Copyright 2022 Pujit Mehrotra
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package category

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
// itemStyle = list.NewDefaultItemStyles()
)

type selectorDelegate struct{}

// TODO add help funcs for each item

func (d selectorDelegate) Height() int                             { return 1 }
func (d selectorDelegate) Spacing() int                            { return 0 }
func (d selectorDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d selectorDelegate) Render(writer io.Writer, m list.Model, index int, item list.Item) {
	// For example, see bubbles.list.DefaultRenderFunc
	// https://github.com/charmbracelet/bubbles/blob/1d489252fe50b60b50e7b24679e02323441aec51/list/defaultitem.go#L128
	var (
		title, desc string
		// titleLength  int
		itemSelected = index == m.Index()
		itemStyle    = getItemStyles()
	)

	if i, ok := item.(list.DefaultItem); ok {
		title = i.Title()
		// titleLength = len(title)
		desc = i.Description()
	} else {
		return
	}

	if itemSelected {
		title = itemStyle.SelectedTitle.Render(title)
		desc = itemStyle.SelectedDesc.Render(desc)
	} else {
		title = itemStyle.NormalTitle.Render(title)
		desc = itemStyle.NormalDesc.Render(desc)
	}

	// Used to justify rendered text
	const (
		// TODO get length of longest item.Title() dynamically
		widestTitleLength = 6
		// a constant gap to left justify commitType title in relation to its description
		// Most likely dependent on visual configurations via lipgloss + stuff I don't understand/care to understand
		padDistance = 25
	)
	fmt.Fprintf(writer, "%-*s%s", padDistance+widestTitleLength, title, desc)
}

func getItemStyles() list.DefaultItemStyles {
	s := list.NewDefaultItemStyles()
	s.SelectedTitle = s.NormalTitle.Copy().
		Foreground(lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"})

	return s
}
