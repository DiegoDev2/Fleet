package cli

import (
	"LattePkg/handlers"
	"fmt"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "Fleet",
		Short: "A package manager modern",
	}

	// Añadir comandos
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
				fmt.Println("Uninstalling " + args[0])
			},
		},
		&cobra.Command{
			Use:   "upgrade [package]",
			Short: "Upgrade a package",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("Upgrading " + args[0])
				// Aquí va la lógica para actualizar
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
			Use:   "search [package]",
			Short: "Search the package",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("Buscando el paquete: " + args[0])
				// Aquí va la lógica de búsqueda
			},
		},
	)

	return rootCmd
}
