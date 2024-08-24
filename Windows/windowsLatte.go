package main

import (
	"fmt"
	"os/exec"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// Mensajes que se mostrarán al usuario
var messages = map[string]string{
	"installing":         "Installing",
	"uninstalling":       "Uninstalling",
	"upgrading":          "Upgrading",
	"searching":          "Searching",
	"patience":           "This operation is taking longer than expected. Please be patient...",
	"error":              "Error",
	"success":            "Operation successful! You can now use 'Latte' from anywhere.",
	"failed":             "Operation failed.",
	"checkUpdate":        "Checking for updates...",
	"updateAvailable":    "Update available! Downloading the latest version...",
	"updateNotAvailable": "No updates available. You're using the latest version.",
}

// Punto de entrada de la aplicación
func main() {
	Command()
}

// Configura y ejecuta los comandos de la CLI
func Command() {
	var rootCmd = &cobra.Command{
		Use:   "Latte",
		Short: "A CLI tool for managing packages with Chocolatey",
	}

	// Agregar los subcomandos
	rootCmd.AddCommand(
		createCommand("install", "Install a package", Install),
		createCommand("uninstall", "Uninstall a package", Uninstall),
		createCommand("upgrade", "Upgrade a package", Update),
		createCommand("version", "Show the version of LattePkg", ShowVersion),
		createCommand("search", "Search for a package", Search),
	)

	// Configura la función de ayuda personalizada
	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		color.Cyan.Println("Usage: Latte [command] [flags]")
		color.Yellow.Println("Commands:")
		color.Green.Println("  install   Install a package")
		color.Green.Println("  uninstall Uninstall a package")
		color.Green.Println("  upgrade   Upgrade a package")
		color.Green.Println("  version   Show the version of LattePkg")
		color.Green.Println("  search    Search for a package")
		color.Red.Println("Flags:")
		color.Magenta.Println("  --help    Show this help message")
	})

	// Ejecutar el comando principal
	if err := rootCmd.Execute(); err != nil {
		color.Red.Println(messages["error"] + ": " + err.Error())
	}
}

// Crea un nuevo comando cobra
func createCommand(use, short string, handler func(string)) *cobra.Command {
	return &cobra.Command{
		Use:   use + " [package]",
		Short: short,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			handler(args[0])
		},
	}
}

// Ejecuta un comando de sistema y devuelve el resultado o un error
func executeCommand(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error executing command: %s, output: %s", err, output)
	}
	return string(output), nil
}

// Maneja la instalación de un paquete con Chocolatey
func Install(pkg string) {
	runChocoCommand("install", pkg, messages["installing"])
}

// Maneja la desinstalación de un paquete con Chocolatey
func Uninstall(pkg string) {
	runChocoCommand("uninstall", pkg, messages["uninstalling"])
}

// Maneja la actualización de un paquete con Chocolatey
func Update(pkg string) {
	runChocoCommand("upgrade", pkg, messages["upgrading"])
}

// Maneja la búsqueda de un paquete con Chocolatey
func Search(pkg string) {
	runChocoCommand("search", pkg, messages["searching"])
}

// Muestra la versión del programa
func ShowVersion(pkg string) {
	version := "LattePkg Version 1.0.2"
	color.Green.Println(version)
}

// Ejecuta un comando de Chocolatey con los argumentos proporcionados
func runChocoCommand(action, pkg, actionMsg string) {
	color.Green.Println(actionMsg + " " + pkg + "...")
	output, err := executeCommand("choco", action, pkg)
	if err != nil {
		color.Red.Println(messages["error"] + ": " + err.Error())
	} else {
		color.Green.Println(output)
		color.Green.Println(messages["success"])
	}
}
