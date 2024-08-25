package main

import (
	"bytes"
	"encoding/json"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
)

var apiBaseURL = "http://localhost:8080" // Cambia esta URL según tu configuración

var statusMessages = map[string]string{
	"installing":         "Installing",
	"uninstalling":       "Uninstalling",
	"upgrading":          "Upgrading",
	"downloading":        "Downloading",
	"patience":           "This operation is taking longer than expected. Please be patient...",
	"error":              "Error",
	"success":            "Operation successful!",
	"failed":             "Operation failed.",
	"enterPackage":       "Enter the package name:",
	"checkUpdate":        "Checking for updates...",
	"updateAvailable":    "Update available! Downloading the latest version...",
	"updateNotAvailable": "No updates available. You're using the latest version.",
}

func main() {
	setupCommands()
}

func setupCommands() {
	var rootCmd = &cobra.Command{
		Use:   "latte",
		Short: "A CLI tool for managing packages",
	}

	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "install [package]",
			Short: "Install a package",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				installPackage(args[0])
			},
		},
		&cobra.Command{
			Use:   "uninstall [package]",
			Short: "Uninstall a package",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				uninstallPackage(args[0])
			},
		},
		&cobra.Command{
			Use:   "upgrade [package]",
			Short: "Upgrade a package",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				upgradePackage(args[0])
			},
		},
		&cobra.Command{
			Use:   "download [package]",
			Short: "Download a package",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				downloadPackage(args[0])
			},
		},
		&cobra.Command{
			Use:   "version",
			Short: "Show the version of LattePkg",
			Run: func(cmd *cobra.Command, args []string) {
				showVersion()
			},
		},
		&cobra.Command{
			Use:   "search [package]",
			Short: "Search for a package",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				searchPackage(args[0])
			},
		},
	)

	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		color.Cyan.Println("Usage: latte [command] [flags]")
		color.Yellow.Println("Commands:")
		color.Green.Println("  install   Install a package")
		color.Green.Println("  uninstall Uninstall a package")
		color.Green.Println("  upgrade   Upgrade a package")
		color.Green.Println("  download  Download a package")
		color.Green.Println("  version   Show the version of LattePkg")
		color.Green.Println("  search    Search for packages")
		color.Red.Println("Flags:")
		color.Magenta.Println("  --help    Show this help message")
	})

	rootCmd.Execute()
}

func executeHTTPRequest(method, endpoint string, body interface{}) (*http.Response, error) {
	var reqBody []byte
	var err error
	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, apiBaseURL+endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	return client.Do(req)
}

func installPackage(pkg string) {
	resp, err := executeHTTPRequest("POST", "/packages/install", map[string]string{"package": pkg})
	if err != nil {
		color.Red.Println(statusMessages["error"]+":", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		color.Red.Println("Error fetching package info:", resp.Status)
		return
	}

	var info map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		color.Red.Println(statusMessages["error"]+":", err)
		return
	}

	color.Green.Println(statusMessages["installing"] + " " + pkg + "...")
	color.Green.Println(statusMessages["success"])
}

func uninstallPackage(pkg string) {
	resp, err := executeHTTPRequest("POST", "/packages/uninstall", map[string]string{"package": pkg})
	if err != nil {
		color.Red.Println(statusMessages["error"]+":", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		color.Red.Println("Error uninstalling package:", resp.Status)
		return
	}

	color.Green.Println(statusMessages["uninstalling"] + " " + pkg + "...")
	color.Green.Println(statusMessages["success"])
}

func upgradePackage(pkg string) {
	resp, err := executeHTTPRequest("POST", "/packages/upgrade", map[string]string{"package": pkg})
	if err != nil {
		color.Red.Println(statusMessages["error"]+":", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		color.Red.Println("Error upgrading package:", resp.Status)
		return
	}

	color.Green.Println(statusMessages["upgrading"] + " " + pkg + "...")
	color.Green.Println(statusMessages["success"])
}

func downloadPackage(pkg string) {
	resp, err := executeHTTPRequest("GET", "/download/"+pkg, nil)
	if err != nil {
		color.Red.Println(statusMessages["error"]+":", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		color.Red.Println("Error downloading package:", resp.Status)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		color.Red.Println(statusMessages["error"]+":", err)
		return
	}

	fileName := pkg + ".tar.gz" // Cambia esto según el tipo de archivo que estés manejando
	err = ioutil.WriteFile(fileName, body, 0644)
	if err != nil {
		color.Red.Println(statusMessages["error"]+":", err)
		return
	}

	color.Green.Println(statusMessages["downloading"] + " " + pkg + "...")
	color.Green.Println(statusMessages["success"])
}

func showVersion() {
	versionInfo := "LattePkg Version 1.0.2"
	color.Greenln(versionInfo)
}

func searchPackage(query string) {
	resp, err := executeHTTPRequest("GET", "/packages/search?query="+query, nil)
	if err != nil {
		color.Red.Println(statusMessages["error"]+":", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		color.Red.Println("Error searching packages:", resp.Status)
		return
	}

	var results map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		color.Red.Println(statusMessages["error"]+":", err)
		return
	}

	color.Green.Println("Search results for", query, ":")
	if packages, ok := results["packages"].([]interface{}); ok {
		for _, pkg := range packages {
			color.Green.Println(pkg)
		}
	} else {
		color.Red.Println("Unexpected response format:", results)
	}
}
