package main

import (
	"bufio"
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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
	"packageURL":         "Enter the package URL:",
	"packageType":        "What type of package is this? (zip/git/binary/tar.gz/other):",
	"formula":            "Here is a formula template you can use:",
	"formulaExample": `package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

// Replace this URL with the actual URL of your package
const packageURL = "REPLACE_WITH_PACKAGE_URL"

// PackageType defines the type of package (zip/git/binary/tar.gz/other)
const packageType = "REPLACE_WITH_PACKAGE_TYPE"

func main() {
	// Add your installation logic here
	fmt.Println("Installing package from", packageURL)
}

func downloadPackage() error {
	// Add your download logic here
	return nil
}
`,
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
			Short: "Upload your package",
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

	cmd := exec.Command("go", "run", formulaPath)
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

	cmd := exec.Command("go", "run", formulaPath, "uninstall")
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

	cmd := exec.Command("go", "run", formulaPath, "upgrade")
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
	cmd := exec.Command("go", "run", "LattePkg/Formulas/"+string(search[0])+"/"+search+"/"+search+".go")
	color.Green.Println("Search", search, "....")
	output, err := cmd.CombinedOutput()
	if err != nil {
		color.Red.Println("Error:", err)
	} else {
		color.Green.Println(string(output))
	}
}

func downloadFormula(pkg string) (string, error) {
	baseURL := "https://raw.githubusercontent.com/CodeDiego15/LattePkg/main/Formula"
	formulaURL := fmt.Sprintf("%s/%s/%s.go", baseURL, string(pkg[0]), pkg)
	color.Green.Println("Downloading", formulaURL, "...")

	resp, err := http.Get(formulaURL)
	if err != nil {
		return "", fmt.Errorf("failed to download formula: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download formula: status code %d", resp.StatusCode)
	}

	formulaDir := filepath.Join("LattePkg", "Formulas", string(pkg[0]), pkg)
	if err := os.MkdirAll(formulaDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create directories: %v", err)
	}

	formulaPath := filepath.Join(formulaDir, fmt.Sprintf("%s.go", pkg))
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
	return filepath.Join("LattePkg", "Formulas", string(pkg[0]), pkg, pkg+".go")
}

func uploadPackage(pkg string) {
	fmt.Print(messages["packageURL"] + " ")
	url, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	url = strings.TrimSpace(url)

	fmt.Print(messages["packageType"] + " ")
	packageType, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	packageType = strings.TrimSpace(packageType)

	color.Green.Println(messages["formula"])
	color.Green.Println(messages["formulaExample"])

	color.Cyan.Println("Modify the formula by replacing the placeholders:")
	color.Cyan.Println("- `REPLACE_WITH_PACKAGE_URL` with the actual package URL")
	color.Cyan.Println("- `REPLACE_WITH_PACKAGE_TYPE` with the type of package (zip/git/binary/tar.gz/other)")

	color.Green.Println("Example formula with placeholders replaced:")
	formula := strings.Replace(messages["formulaExample"], "REPLACE_WITH_PACKAGE_URL", url, 1)
	formula = strings.Replace(formula, "REPLACE_WITH_PACKAGE_TYPE", packageType, 1)
	color.Green.Println(formula)

	// Agregar lógica para manejar el paquete proporcionado
	fmt.Printf("Uploading package: %s\n", pkg)
	// Implementa la lógica para subir el paquete aquí
}
