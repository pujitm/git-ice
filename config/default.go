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

package config

// Represents a config to set up an ice commit prompt
//
// Supports TOML en/decoding
type IceCommit struct {
	Resolve Resolutions // How to resolve each section
	Types   []CommitType
	Scopes  []string
}

// Represents a type of commit, according to the conventional commit spec
//
// ie. "fix" in "fix(brain): my sanity"
type CommitType struct {
	// Bump      string // major, minor, patch, or empty
	Git         string // Text representation in git commit header (i.e. "fix", "feat")
	Description string // Purpose, utility, and use-cases for the commit type
	Ordinal     uint   // Display position (lower ordinal value -> higher display position)
}

// How to resolve differences in each config section among multiple files
//
// Options:
//
// - merge (default): merges/extends configs of lower priority/specificity. Default config DefaultIceCommit() should have lowest priority.
//
// - replace: ignores (and replaces) values provided by lower configs
//
// Future possibilities: url, function name (custom logic)
type Resolutions struct {
	Types  string
	Scopes string
}

const (
	MergeResolution   string = "merge"   // A config resolution strategy. Merges a config with base config(s) preceding it
	ReplaceResolution        = "replace" // A config resolution strategy. Ignores values from any preceding base configs, including the default one (See DefaultIceCommit)
)

// Returns a default git-ice commit configuration
func DefaultIceCommit() IceCommit {
	return IceCommit{
		Resolve: defaultResolutions(),
		Types:   defaultCommitTypes(),
		Scopes:  defaultCommitScopes(),
	}
}

func defaultCommitTypes() []CommitType {
	return []CommitType{
		{Ordinal: 0, Git: "feat", Description: "Add feature or change functionality"},
		{Ordinal: 1, Git: "fix", Description: "Bug fix"},
		{Ordinal: 2, Git: "inc", Description: "Incremental work and progress"},
		{Ordinal: 3, Git: "docs", Description: "Change or update documentation"},
		{Ordinal: 4, Git: "infra", Description: "Infrastructure to deploy, run, and manage the application"},
		{Ordinal: 5, Git: "config", Description: "Configuration for project, repo, or build"},
		{Ordinal: 6, Git: "exp", Description: "Experiment with something"},
	}
}

// eg for infra, aws, gcp, lambda, gh (for github), sls (serverless), k8s
func defaultCommitScopes() []string {
	return []string{
		"aws",
	}
}

// Returns merge resolution strategy for each section
func defaultResolutions() Resolutions {
	defaultRes := MergeResolution
	return Resolutions{
		Types:  defaultRes,
		Scopes: defaultRes,
	}
}
