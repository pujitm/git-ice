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
	"os"

	"github.com/erikgeiser/promptkit/confirmation"
)

func PromptBodyEdit() (bool, error) {
	editBodyPrompt := confirmation.New("Edit message body?", confirmation.No)
	return editBodyPrompt.RunPrompt()
}

// Runs a git commit operation based on the info provided
//
// Copied from exec.Run docs:
//
// The returned error is nil if the command runs, has no problems copying stdin, stdout, and stderr, and exits with a zero exit status.
//
// If the command starts but does not complete successfully, the error is of type *exec.ExitError. Other error types may be returned for other situations.
func RunCommit(info CommitInfo) error {
	commitCmd := BuildCommitCommand(info)
	// To ensure interactivity with terminal editors like vim
	commitCmd.Stdout = os.Stdout
	commitCmd.Stdin = os.Stdin
	commitCmd.Stderr = os.Stderr
	return commitCmd.Run()
}
