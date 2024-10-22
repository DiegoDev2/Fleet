// Copyright (C) 2024 Fleet Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lib

import "LattePkg/formulas"

type Tool struct {
	Name    string
	Install func()
}

var tools = []Tool{
	{
		Name:    "awscli",
		Install: formulas.InstallAws,
	},
	{
		Name:    "nmap",
		Install: formulas.InstallNmap,
	},
	{
		Name:    "nuclei",
		Install: formulas.InstallNuclei,
	},
	{
		Name:    "k9s",
		Install: formulas.InstallK9s,
	},
	{
		Name:    "git",
		Install: formulas.InstallGit,
	},
	{
		Name:    "go",
		Install: formulas.InstallGo,
	},
	{
		Name:    "aircrack-ng",
		Install: formulas.InstallAircrackNg,
	},
	{
		Name:    "node",
		Install: formulas.InstallNode,
	},
	{
		Name:    "docker",
		Install: formulas.InstallDocker,
	},
	{
		Name:    "python3",
		Install: formulas.InstallPython3,
	},
	{
		Name:    "python2",
		Install: formulas.InstallPython2,
	},
	{
		Name:    "kubectl",
		Install: formulas.InstallKubectl,
	},
	{
		Name:    "rust",
		Install: formulas.InstallCargo,
	},
	{
		Name:    "java21",
		Install: formulas.InstallJava21,
	},
	{
		Name:    "dotnet",
		Install: formulas.InstallDotnet,
	},
}
