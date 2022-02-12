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
	"bytes"
	"fmt"
	"os"

	"github.com/pujitm/git-ice/category"
	"github.com/pujitm/git-ice/compatibility"
	"github.com/pujitm/git-ice/config"
	"github.com/pujitm/git-ice/message"
	"github.com/pujitm/git-ice/scope"
	"github.com/pujitm/git-ice/subject"

	"github.com/BurntSushi/toml"
)

func main() {
	InteractiveCommit()
}

func PrintDefaultConfig() {
	var def bytes.Buffer
	e := toml.NewEncoder(&def)
	e.Encode(config.DefaultIceCommit())
	fmt.Println(def.String())
}

func InteractiveCommit() {
	handleError := func(err error) {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	}
	iceConfig, err := config.ResolveDefaultConfig()
	// TODO log config at debug level
	handleError(err)

	info := message.CommitInfo{}
	info.CommitType, err = category.RunPrompt(iceConfig.Types)
	handleError(err)

	info.BackwardsCompatible, err = compatibility.Prompt()
	handleError(err)

	info.Scope, err = scope.Prompt()
	handleError(err)

	info.Subject, err = subject.Prompt()
	handleError(err)

	editBody, err := message.PromptBodyEdit()
	handleError(err)

	args := []string{}
	if editBody {
		args = append(args, "--edit")
	}

	info.Args = args
	err = message.RunCommit(info)
	handleError(err)
}
