
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// spotify-tuiFormula represents a formula in Go.
type spotify-tuiFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg spotify-tuiFormula) Print() {
    fmt.Printf("Name: spotify-tui\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := spotify-tuiFormula{
        Description:  "Terminal-based client for Spotify",
        Homepage:     "https://github.com/Rigellute/spotify-tui",
        URL:          "https://github.com/Rigellute/spotify-tui/commit/14df9419cf72da13f3b55654686a95647ea9dfea.patch?full_index=1",
        Sha256:       "44f95b14320eb3274131f6676c1fb7bc4096735a16592a01fc1164dbe3a064e5",
        Dependencies: []string{"rust", "pkg-config", "libxcb", "openssl@1.1"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        if !isFormulaInstalled(dep) {
            fmt.Printf("Installing dependency: %s\n", dep)
            cmd := exec.Command("brew", "install", dep)
            if err := cmd.Run(); err != nil {
                log.Fatalf("Error installing dependency %s: %v", dep, err)
            }
        } else {
            fmt.Printf("Dependency %s is already installed.\n", dep)
        }
    }

    fmt.Printf("Installing formula: %s\n", "spotify-tui")
    if err := pkg.Installspotify-tui(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg spotify-tuiFormula) Installspotify-tui() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "14df9419cf72da13f3b55654686a95647ea9dfea.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "14df9419cf72da13f3b55654686a95647ea9dfea"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
