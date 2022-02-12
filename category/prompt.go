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
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pujitm/git-ice/config"
)

func RunPrompt(types []config.CommitType) (string, error) {
	commitCategory := Model(types)
	program := tea.NewProgram(commitCategory)
	if err := program.Start(); err != nil {
		return "", err
	}
	return commitCategory.Value(), nil
}
