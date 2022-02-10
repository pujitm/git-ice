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
	"sort"
	"strings"
)

type commitType struct {
	// Bump      string // major, minor, patch, or empty
	Ordinance uint8  // used to sort/order types
	Git       string // Text representation in git header. Must be unique
	// title       string // Display Title
	description string // Display description
}

func (t commitType) Title() string       { return strings.ToUpper(t.Git) }
func (t commitType) Description() string { return t.description }
func (t commitType) FilterValue() string { return t.Git }

var commitTypes = []commitType{
	{Ordinance: 0, Git: "feat", description: "Add feature or change functionality"},
	{Ordinance: 1, Git: "fix", description: "Bug fix"},
	{Ordinance: 2, Git: "inc", description: "Incremental work and progress"},
	{Ordinance: 3, Git: "docs", description: "Change or update documentation"},
	{Ordinance: 4, Git: "infra", description: "Infrastructure to deploy, run, and manage the application"},
	{Ordinance: 5, Git: "config", description: "Configuration for project, repo, or build"},
	{Ordinance: 6, Git: "exp", description: "Experiment with something"},
}

// TODO unit test
func sortCommitTypes(types []commitType) {
	sort.Slice(types, func(i, j int) bool { return types[i].Ordinance < types[j].Ordinance })
}

// Returns sorted list of commit types
func getCommitTypes() []commitType {
	// Don't need to sort dummy options
	// return append([]commitType{}, commitTypes...)
	return commitTypes
}
