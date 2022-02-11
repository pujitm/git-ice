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

package integration_test

import (
	"log"
	"os"
	"pujitm/git-ice/config"
	"testing"
)

func getConfigs(t *testing.T) []config.IceCommit {
	workingDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Could not resolve working dir: %v\n", err)
		return []config.IceCommit{}
	}
	return config.GetIceConfigsFrom(workingDir)
}

func TestConfigDiscovery(t *testing.T) {
	configs := getConfigs(t)
	if len(configs) != 2 {
		log.Fatalf("Expected 2 config files, Found %d: %v\n", len(configs), configs)
	}
}

// TODO add test cases
func TestConfigResolution(t *testing.T) {
	configs := []config.IceCommit{config.DefaultIceCommit()}
	configs = append(configs, getConfigs(t)...)

	resolved := config.ResolveConfig(configs)
	t.Logf("Resolved config: %v\n", resolved)
}
