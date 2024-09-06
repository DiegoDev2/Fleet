package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var messages = map[string]string{
	"installing":         "Installing",
	"uninstalling":       "Uninstalling",
	"upgrading":          "Upgrading",
	"patience":           "This operation is taking longer than expected. Please be patient...",
	"error":              "Error",
	"success":            "Installation successful! You can now use 'latte' from anywhere.",
	"failed":             "Installation failed.",
	"enterPackage":       "Enter the package name:",
	"checkUpdate":        "Checking for updates...",
	"updateAvailable":    "Update available! Downloading the latest version...",
	"updateNotAvailable": "No updates available. You're using the latest version.",
}

func main() {
	Command()
}

func Command() {
	var rootCmd = &cobra.Command{
		Use:   "Latte",
		Short: "A CLI tool for managing packages",
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
		&cobra.Command{
			Use:   "version",
			Short: "Show the version of LattePkg",
			Run: func(cmd *cobra.Command, args []string) {
				version()
			},
		},
		&cobra.Command{
			Use:   "search [package]",
			Short: "Search the package",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				Search(args[0])
			},
		},
		&cobra.Command{
			Use:   "upload [package]",
			Short: "upload your package",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				uploadPackage(args[0])
			},
		},
	)

	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		color.Cyan.Println("Usage: latte [command] [flags]")
		color.Yellow.Println("Commands:")
		color.Green.Println("  install   Install a package")
		color.Green.Println("  uninstall Uninstall a package")
		color.Green.Println("  upgrade   Upgrade a package")
		color.Green.Println("  version   Show the version of LattePkg")
		color.Green.Println("  search    Search your packages")
		color.Green.Println("  upload    Upload your package")
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
		return "", fmt.Errorf("error executing command: %s", err)
	}
	return string(output), nil
}

func Install(pkg string) {
	formulaPath, err := getOrDownloadFormula(pkg)
	if err != nil {
		color.Red.Println(messages["error"]+":", err)
		return
	}

	cmd := exec.Command("brew", "install", formulaPath)
	color.Green.Println(messages["installing"] + " " + pkg + "...")
	output, err := executeCommand(cmd)
	if err != nil {
		color.Red.Println(messages["error"]+":", err)
	} else {
		if output != "" {
			color.Green.Println(output)
		}
	}
}

func Uninstall(pkg string) {
	formulaPath, err := getOrDownloadFormula(pkg)
	if err != nil {
		color.Red.Println(messages["error"]+":", err)
		return
	}

	cmd := exec.Command("brew", "uninstall", formulaPath)
	color.Green.Println(messages["uninstalling"] + " " + pkg + "...")
	output, err := executeCommand(cmd)
	if err != nil {
		color.Red.Println(messages["error"]+":", err)
	} else {
		if output != "" {
			color.Green.Println(output)
		}
	}
}

func Update(pkg string) {
	formulaPath, err := getOrDownloadFormula(pkg)
	if err != nil {
		color.Red.Println(messages["error"]+":", err)
		return
	}

	cmd := exec.Command("brew", "upgrade", formulaPath)
	color.Green.Println(messages["upgrading"] + " " + pkg + "...")
	output, err := executeCommand(cmd)
	if err != nil {
		color.Red.Println(messages["error"]+":", err)
	} else {
		if output != "" {
			color.Green.Println(output)
		}
	}
}

func version() {
	com := "LattePkg Version"
	ver := com + " 1.0.2"
	color.Greenln(ver)
}

func Search(search string) {
	cmd := exec.Command("brew", "search", search)
	color.Green.Println("Searching", search, "....")
	output, err := cmd.CombinedOutput()
	if err != nil {
		color.Red.Println("Error:", err)
	} else {
		color.Green.Println(string(output))
	}
}

func downloadFormula(pkg string) (string, error) {
	firstLetter := strings.ToLower(string(pkg[0]))
	baseURL := fmt.Sprintf("https://raw.githubusercontent.com/Homebrew/homebrew-core/master/Formula/%s", firstLetter)
	formulaURL := fmt.Sprintf("%s/%s.rb", baseURL, strings.ToLower(pkg))
	color.Green.Println("Downloading", formulaURL, "...")

	resp, err := http.Get(formulaURL)
	if err != nil {
		return "", fmt.Errorf("failed to download formula: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download formula: status code %d", resp.StatusCode)
	}

	formulaDir := filepath.Join("LattePkg", "Formulas", firstLetter, strings.ToLower(pkg))
	if err := os.MkdirAll(formulaDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create directories: %v", err)
	}

	formulaPath := filepath.Join(formulaDir, fmt.Sprintf("%s.rb", strings.ToLower(pkg)))
	out, err := os.Create(formulaPath)
	if err != nil {
		return "", fmt.Errorf("failed to create formula file: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to save formula: %v", err)
	}

	return formulaPath, nil
}

func getOrDownloadFormula(pkg string) (string, error) {
	localPath := getFormulaPath(pkg)
	if _, err := os.Stat(localPath); os.IsNotExist(err) {
		return downloadFormula(pkg)
	}
	return localPath, nil
}

func getFormulaPath(pkg string) string {
	firstLetter := strings.ToLower(string(pkg[0]))
	return filepath.Join("LattePkg", "Formulas", firstLetter, strings.ToLower(pkg), strings.ToLower(pkg)+".rb")
}

func uploadPackage(form string) {
	type Form struct {
		Link     string
		HomePage string
		ShaKey   string
		Dep      string
	}

	var tool Form

	for {
		fmt.Println("Ingrese el link de descarga de su herramienta:")
		_, err := fmt.Scan(&tool.Link)
		if err != nil {
			fmt.Println("Error al leer el link:", err)
			continue
		}

		fmt.Println("Ingrese la página principal de su herramienta:")
		_, err = fmt.Scan(&tool.HomePage)
		if err != nil {
			fmt.Println("Error al leer la página principal:", err)
			continue
		}

		fmt.Println("Ingrese la clave SHA:")
		_, err = fmt.Scan(&tool.ShaKey)
		if err != nil {
			fmt.Println("Error al leer la clave SHA:", err)
			continue
		}

		fmt.Println("Ingrese las dependencias (separadas por comas):")
		_, err = fmt.Scan(&tool.Dep)
		if err != nil {
			fmt.Println("Error al leer las dependencias:", err)
			continue
		}
		break
	}
}
