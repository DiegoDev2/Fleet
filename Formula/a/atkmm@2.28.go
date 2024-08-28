
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// atkmm@2.28Formula represents a formula in Go.
type atkmm@2.28Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg atkmm@2.28Formula) Print() {
    fmt.Printf("Name: atkmm@2.28\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := atkmm@2.28Formula{
        Description:  "Official C++ interface for the ATK accessibility toolkit library",
        Homepage:     "https://www.gtkmm.org/",
        URL:          "https://download.gnome.org/sources/atkmm/2.28/atkmm-2.28.4.tar.xz",
        Sha256:       "fc9b8e9286aaf97bc9255158fae66cb4ea5c6a2c2c8f6cbb3681a69cd9fa39f3",
        Dependencies: []string{"meson", "ninja", "pkg-config", "at-spi2-core", "glib", "glibmm@2.66", "libsigc++@2", "gettext"},
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

    fmt.Printf("Installing formula: %s\n", "atkmm@2.28")
    if err := pkg.Installatkmm@2.28(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg atkmm@2.28Formula) Installatkmm@2.28() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "atkmm-2.28.4.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "atkmm-2.28.4.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
