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

package main

import (
	"fmt"
	"os"
	"pujitm/git-ice/category"
	"pujitm/git-ice/compatibility"
	"pujitm/git-ice/scope"
	"pujitm/git-ice/subject"
)

func main() {
	commitType, err := category.RunPrompt()
	handleError(err)

	compatible, err := compatibility.Prompt()
	handleError(err)

	scope, err := scope.Prompt()
	handleError(err)

	subject, err := subject.Prompt()
	handleError(err)

	header := fmt.Sprintf("%s%s: %s", commitType, formatScope(scope), subject)

	fmt.Printf("# %s\n%s", header, formatCompatibility(compatible))
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func formatScope(scope string) string {
	if scope == "" {
		return scope
	}
	return fmt.Sprintf("(%s)", scope)
}

func formatCompatibility(isCompatible bool) string {
	if isCompatible {
		return ""
	}
	return "Note: Is API-Breaking change\n"
}
