package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

func main() {
	var err error
	var lang string

	// Intentar cargar la preferencia de idioma
	lang, err = loadLanguagePreference()
	if err != nil || (lang != "en" && lang != "es") {
		// Preguntar el idioma si no está guardado o es inválido
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(messages["chooseLang"] + " ")
		lang, _ = reader.ReadString('\n')
		lang = lang[:len(lang)-1]

		if lang != "en" && lang != "es" {
			fmt.Println("Invalid language choice. Defaulting to English.")
			lang = "en"
		}

		// Guardar la preferencia de idioma
		err = saveLanguagePreference(lang)
		if err != nil {
			fmt.Println("Error saving language preference:", err)
			return
		}
	}

	// Cargar los mensajes del idioma seleccionado
	if lang == "es" {
		messages = loadMessagesES()
	} else {
		messages = loadMessagesEN()
	}

	Command()
}

func loadMessagesEN() map[string]string {
	return map[string]string{
		"installing":   "Installing",
		"uninstalling": "Uninstalling",
		"upgrading":    "Upgrading",
		"patience":     "This operation is taking longer than expected. Please be patient...",
		"error":        "Error",
		"success":      "Installation successful! You can now use 'turn' from anywhere.",
		"failed":       "Installation failed.",
		"enterPackage": "Enter the package name:",
		"chooseLang":   "Choose language (en/es):",
	}
}

func loadMessagesES() map[string]string {
	return map[string]string{
		"installing":   "Instalando",
		"uninstalling": "Desinstalando",
		"upgrading":    "Actualizando",
		"patience":     "Esta operación está tardando más de lo esperado. Por favor, tenga paciencia...",
		"error":        "Error",
		"success":      "¡Instalación exitosa! Ahora puede usar 'turn' desde cualquier lugar.",
		"failed":       "La instalación falló.",
		"enterPackage": "Ingrese el nombre del paquete:",
		"chooseLang":   "Elija idioma (en/es):",
	}
}

func saveLanguagePreference(lang string) error {
	file, err := os.Create("language.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(lang)
	return err
}

func loadLanguagePreference() (string, error) {
	file, err := os.Open("language.txt")
	if err != nil {
		return "", err
	}
	defer file.Close()

	var lang string
	_, err = fmt.Fscanf(file, "%s", &lang)
	return lang, err
}

func Command() {
	var rootCmd = &cobra.Command{
		Use:   "turn",
		Short: "A CLI tool for managing packages with Homebrew",
	}

	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "install [package]",
			Short: "Install a package",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				Install(args[0])
			},
		},
		&cobra.Command{
			Use:   "uninstall [package]",
			Short: "Uninstall a package",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				Uninstall(args[0])
			},
		},
		&cobra.Command{
			Use:   "upgrade [package]",
			Short: "Upgrade a package",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				Update(args[0])
			},
		},
	)

	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		color.Cyan.Println("Usage: turn [command] [flags]")
		color.Yellow.Println("Commands:")
		color.Green.Println("  install   Install a formula")
		color.Green.Println("  uninstall Uninstall a formula")
		color.Green.Println("  upgrade   Upgrade a formula")
		color.Red.Println("Flags:")
		color.Magenta.Println("  --help    Show this help message")
	})

	rootCmd.Execute()
}

func executeCommand(cmd *exec.Cmd) (string, error) {
	cmd.Stdout = nil
	cmd.Stderr = nil
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func Install(pkg string) {
	cmd := exec.Command("brew", "install", pkg)
	done := make(chan bool)

	go func() {
		color.Green.Println(messages["installing"] + " " + pkg + "...")
		output, err := executeCommand(cmd)
		if err != nil {
			color.Red.Println(messages["error"]+":", err)
		} else {
			if output != "" {
				color.Green.Println(output)
			}
		}
		done <- true
	}()

	select {
	case <-done:
		// Command finished within 30 seconds
	case <-time.After(30 * time.Second):
		color.Yellow.Println(messages["patience"])
	}
}

func Uninstall(pkg string) {
	cmd := exec.Command("brew", "uninstall", pkg)
	done := make(chan bool)

	go func() {
		color.Green.Println(messages["uninstalling"] + " " + pkg + "...")
		output, err := executeCommand(cmd)
		if err != nil {
			color.Red.Println(messages["error"]+":", err)
		} else {
			if output != "" {
				color.Green.Println(output)
			}
		}
		done <- true
	}()

	select {
	case <-done:
		// Command finished within 30 seconds
	case <-time.After(30 * time.Second):
		color.Yellow.Println(messages["patience"])
	}
}

func Update(pkg string) {
	cmd := exec.Command("brew", "upgrade", pkg)
	done := make(chan bool)

	go func() {
		color.Green.Println(messages["upgrading"] + " " + pkg + "...")
		output, err := executeCommand(cmd)
		if err != nil {
			color.Red.Println(messages["error"]+":", err)
		} else {
			if output != "" {
				color.Green.Println(output)
			}
		}
		done <- true
	}()

	select {
	case <-done:
		// Command finished within 30 seconds
	case <-time.After(30 * time.Second):
		color.Yellow.Println(messages["patience"])
	}
}
