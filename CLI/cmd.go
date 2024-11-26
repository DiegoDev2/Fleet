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

package cli

import (
	"fleet/ai"
	"fleet/handlers"
	"fleet/handlers/list"
	"fmt"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "Fleet",
		Short: "A modern package manager",
	}

	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "install [package]",
			Short: "Install a package",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				handlers.Install(args[0])
			},
		},
		&cobra.Command{
			Use:   "uninstall [package]",
			Short: "Uninstall a package",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				handlers.Uninstall(args[0])
			},
		},
		&cobra.Command{
			Use:   "list",
			Short: "List all available packages",
			Run: func(cmd *cobra.Command, args []string) {
				list.List()
			},
		},
		&cobra.Command{
			Use:   "version",
			Short: "Show the version of Fleet",
			Run: func(cmd *cobra.Command, args []string) {
				handlers.ShowVersion()
			},
		},
		&cobra.Command{
			Use:   "ai [question]",
			Short: "AI assistant",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				result := ai.ResponseAi(args[0])
				fmt.Println(result)
			},
		},
	)

	return rootCmd
}
