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

package message

import (
	"fmt"
	"os/exec"
)

type CommitInfo struct {
	CommitType          string
	Scope               string
	Subject             string
	BackwardsCompatible bool
	// Additional args to pass to the commit command. Slice format must comply with exec.Command
	//
	// See https://stackoverflow.com/a/24429698/6656631)
	Args []string
}

// Builds a git commit exec.Cmd with the provided info
//
// Used instead of go-git (as suggested at https://git-scm.com/book/en/v2/Appendix-B%3A-Embedding-Git-in-your-Applications-go-git)
// because (as of v5), the Commit API doesn't have the desired flexibility/options baked in
// (See https://pkg.go.dev/github.com/go-git/go-git/v5#Worktree.Commit)
func BuildCommitCommand(info CommitInfo) *exec.Cmd {
	header := BuildHeader(info.CommitType, info.Scope, info.Subject)
	footer := BuildFooter(info.BackwardsCompatible)
	args := BuildCommitArgs(header, footer, info.Args...)
	return exec.Command("git", args...)
}

func BuildCommitArgs(header, footer string, customArgs ...string) []string {
	args := []string{"commit", "-m", header}
	if footer != "" {
		args = append(args, "-m", footer)
	}
	return append(args, customArgs...)
}

func BuildHeader(cat, scope, subject string) string {
	return fmt.Sprintf("%s%s: %s", cat, formatScope(scope), subject)
}

func BuildFooter(isCompatible bool) string { return formatFooter(isCompatible) }

func formatScope(scope string) string {
	if scope == "" {
		return scope
	}
	return fmt.Sprintf("(%s)", scope)
}

func formatFooter(isCompatible bool) string {
	if isCompatible {
		return ""
	}
	return "BREAKING CHANGE"
}
