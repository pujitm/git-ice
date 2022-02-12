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

import "sort"

// Resolves a unified config from the default config resolution path
//
// [ Default, Project, Personal ] (lowest to highest specificity)
//
// Using `GetDefaultIceConfigs()`
func ResolveDefaultConfig() (IceCommit, error) {
	configs, err := GetDefaultIceConfigs()
	if err != nil {
		return DefaultIceCommit(), err
	}
	return ResolveConfig(configs), nil
}

// Takes list of IceCommit configs ordered from lowest to highest priority
func ResolveConfig(configs []IceCommit) IceCommit {
	types, scopes := combineConfigs(configs)
	return IceCommit{Types: types, Scopes: scopes}
}

func combineConfigs(configs []IceCommit) ([]CommitType, []string) {
	// can make slice allocs more performant with https://stackoverflow.com/a/13427931/6656631 and some math
	var accScopes []string // scopes accumulator
	accTypes := map[string]CommitType{}
	scopeReg := map[string]bool{} // registry of unique scopes

	// appends unique scopes to accScopes
	// done this way (instead of simpler alternatives) to preserve order of scopes
	addScopes := func(scopes []string) {
		for _, scope := range scopes {
			if _, found := scopeReg[scope]; !found {
				scopeReg[scope] = true
				accScopes = append(accScopes, scope)
			}
		}
	}
	addTypes := func(types []CommitType) {
		for _, t := range types {
			accTypes[t.Git] = t
		}
	}

	for _, config := range configs {
		if config.Resolve.Scopes == ReplaceResolution {
			accScopes = []string{}
			scopeReg = map[string]bool{}
		}
		if config.Resolve.Types == ReplaceResolution {
			accTypes = map[string]CommitType{}
		}
		addScopes(config.Scopes)
		addTypes(config.Types)
	}

	return collapseTypes(accTypes), accScopes
}

// Returns list of CommitTypes in map, sorted by ordinance
func collapseTypes(typeMap map[string]CommitType) []CommitType {
	types := make([]CommitType, 0, len(typeMap))
	for _, t := range typeMap { // -> Values from typesMap
		types = append(types, t)
	}
	// sort aggregate commit types by ordinance
	sort.Slice(types, func(i, j int) bool { return types[i].Ordinal < types[j].Ordinal })
	return types
}
