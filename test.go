package main

import (
	"archive/zip"
	"bufio"
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
	"packageType":        "What type of package is this? (git/zip/binary/other):",
	"packageURL":         "Enter the package URL:",
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
			Use:   "upload",
			Short: "Upload a new package",
			Run:   uploadPackage,
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
		color.Green.Println("  upload    Upload a new package")
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

func uploadPackage(cmd *cobra.Command, args []string) {
	fmt.Print(messages["packageURL"] + " ")
	url, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	url = strings.TrimSpace(url)

	fmt.Print(messages["packageType"] + " ")
	packageType, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	packageType = strings.TrimSpace(packageType)

	tmpDir, err := os.MkdirTemp("", "lattepkg-upload")
	if err != nil {
		color.Red.Println(messages["error"]+":", err)
		return
	}
	defer os.RemoveAll(tmpDir)

	var filePath string
	switch packageType {
	case "git":
		color.Green.Println("Cloning repository from", url)
		cmd := exec.Command("git", "clone", url, tmpDir)
		if _, err := executeCommand(cmd); err != nil {
			color.Red.Println(messages["error"]+":", err)
			return
		}
		filePath = filepath.Join(tmpDir, filepath.Base(url))
	case "zip":
		color.Green.Println("Downloading zip from", url)
		resp, err := http.Get(url)
		if err != nil {
			color.Red.Println(messages["error"]+":", err)
			return
		}
		defer resp.Body.Close()

		zipPath := filepath.Join(tmpDir, "package.zip")
		file, err := os.Create(zipPath)
		if err != nil {
			color.Red.Println(messages["error"]+":", err)
			return
		}
		defer file.Close()

		if _, err := io.Copy(file, resp.Body); err != nil {
			color.Red.Println(messages["error"]+":", err)
			return
		}

		if err := unzip(zipPath, tmpDir); err != nil {
			color.Red.Println(messages["error"]+":", err)
			return
		}
		filePath = filepath.Join(tmpDir, filepath.Base(url))
	case "binary":
		color.Green.Println("Downloading binary from", url)
		resp, err := http.Get(url)
		if err != nil {
			color.Red.Println(messages["error"]+":", err)
			return
		}
		defer resp.Body.Close()

		filePath = filepath.Join(tmpDir, "binary")
		file, err := os.Create(filePath)
		if err != nil {
			color.Red.Println(messages["error"]+":", err)
			return
		}
		defer file.Close()

		if _, err := io.Copy(file, resp.Body); err != nil {
			color.Red.Println(messages["error"]+":", err)
			return
		}
	case "other":
		color.Red.Println("Unsupported package type:", packageType)
		return
	default:
		color.Red.Println("Invalid package type")
		return
	}

	color.Green.Println("Package uploaded successfully to", filePath)
}

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		fpath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, f.Mode())
		} else {
			if err := os.MkdirAll(filepath.Dir(fpath), 0755); err != nil {
				return err
			}
			outFile, err := os.Create(fpath)
			if err != nil {
				return err
			}
			defer outFile.Close()

			if _, err := io.Copy(outFile, rc); err != nil {
				return err
			}
		}
	}
	return nil
}
