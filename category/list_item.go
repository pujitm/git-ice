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
	"strings"

	"github.com/pujitm/git-ice/config"
)

type listItem struct {
	CommitType config.CommitType
}

func (t listItem) Title() string       { return strings.ToUpper(t.CommitType.Git) }
func (t listItem) Description() string { return t.CommitType.Description }
func (t listItem) FilterValue() string { return t.CommitType.Git }
