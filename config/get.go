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

import (
	"fmt"
	"os"
	"path"

	"github.com/BurntSushi/toml"
)

// Shortcut for GetIceConfigsFrom(FindProjectRoot(os.Getwd()))
//
// Returns an ordered slice of IceCommit configs, lowest to highest priority
func GetDefaultIceConfigs() ([]IceCommit, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return []IceCommit{}, err
	}
	projectRoot, err := FindProjectRoot(workingDir)
	if err != nil {
		return []IceCommit{}, err
	}

	configs := []IceCommit{DefaultIceCommit()}
	return append(configs, GetIceConfigsFrom(projectRoot)...), nil
}

// Returns an ordered slice of IceCommit configs.
//
// Lowest priority/specificity are first: [ DefaultConfig, ProjectConfig (if readable), PersonalConfig (if readable) ]
//
// rootDir is the directory to look for ice config files in
func GetIceConfigsFrom(rootDir string) []IceCommit {
	// least specific first
	var configs []IceCommit
	sources := [...]string{DefaultProject, DefaultPersonal}
	read := ReadConfigFrom(rootDir)
	for _, source := range sources {
		if config, err := read(source); err == nil {
			configs = append(configs, config)
		}
	}
	return configs
}

// A helper function to make reads from same root directory easier
//
// Calls ReadConfig once config path is resolved
//
// rootDir is the directory to look for ice config files in
func ReadConfigFrom(rootDir string) func(string) (IceCommit, error) {
	return func(relativeFilePath string) (IceCommit, error) {
		configPath := path.Join(rootDir, relativeFilePath)
		return ReadConfig(configPath)
	}
}

// Reads a TOML IceCommit config file from the path provided
//
// Checks if file exists and attempts to decode it if it does.
// Either step may produce an error, in which case an empty config will be returned with it.
func ReadConfig(absoluteConfigPath string) (IceCommit, error) {
	// TODO change prints to logs https://github.com/Sirupsen/logrus
	// fmt.Printf("Attempting to read config file: %s\n", absoluteConfigPath)
	var config IceCommit
	if _, err := os.Stat(absoluteConfigPath); err != nil {
		fmt.Println(fmt.Errorf("Error reading config file: %+v\n", err))
		return config, err
	}
	// fmt.Printf("Attempting to decode config file: %s\n", absoluteConfigPath)
	if _, err := toml.DecodeFile(absoluteConfigPath, &config); err != nil {
		fmt.Println(fmt.Errorf("Error decoding config file: %+v\n", err))
		return IceCommit{}, err
	}
	return config, nil
}

// Finds the directory of the project workspace by searching for a .git dir
// in the path provided and in each of its parents.
//
// Usually used with os.Getwd(), which provides the command working directory to start the search in
//
// Searches through parents recursively, so beware of stack overflows
//
// Uses FindLocalGitRepo and returns the parent of the .git directory
func FindProjectRoot(of string) (string, error) {
	git, err := FindLocalGitRepo(of)
	if err != nil {
		return "", err
	}
	return path.Join(git, ".."), nil
}

// Returns path of local git repo config
//
// Searches recursively through parent directories `of` the path passed into the parameter until it finds a .git directory
//
// Be wary of runtime stack overflows
func FindLocalGitRepo(of string) (string, error) {
	target := path.Join(of, "./.git")

	_, err := os.Stat(target)
	if os.IsNotExist(err) {
		return FindLocalGitRepo(path.Join(of, "./.."))
	}

	return target, nil
}
