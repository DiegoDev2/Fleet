package cmd

import (
	"fmt"

	"github.com/DiegoDev2/Fleet/internal/core/manifest"
	"github.com/DiegoDev2/Fleet/internal/handlers"
	"github.com/spf13/cobra"
)

func NewSimulateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "simulate [manifest-file]",
		Short: "Simulate installation from a manifest file",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			manifestPath := args[0]

			m, err := manifest.ParseFile(manifestPath)
			if err != nil {
				return fmt.Errorf("error parsing manifest: %w", err)
			}

			errors := manifest.Validate(m)
			if len(errors) > 0 {
				fmt.Println("Manifest validation errors:")
				for _, err := range errors {
					fmt.Printf("  - %s\n", err.Error())
				}
				return fmt.Errorf("invalid manifest")
			}

			return handlers.SimulateInstall(m)
		},
	}
}
