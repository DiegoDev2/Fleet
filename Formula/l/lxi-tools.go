
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// lxi-toolsFormula represents a formula in Go.
type lxi-toolsFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg lxi-toolsFormula) Print() {
    fmt.Printf("Name: lxi-tools\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := lxi-toolsFormula{
        Description:  "Open source tools for managing network attached LXI compatible instruments",
        Homepage:     "https://github.com/lxi-tools/lxi-tools",
        URL:          "https://github.com/lxi-tools/lxi-tools/archive/refs/tags/v2.7.tar.gz",
        Sha256:       "4dd930cc0b1e29dbdbe11595e0112fd420729776cf0dedcdb76bc73d092a1df6",
        Dependencies: []string{"meson", "ninja", "pkg-config", "cairo", "desktop-file-utils", "gdk-pixbuf", "glib", "gtk4", "gtksourceview5", "hicolor-icon-theme", "json-glib", "libadwaita", "liblxi", "lua", "readline", "gettext", "graphene", "harfbuzz", "pango"},
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

    fmt.Printf("Installing formula: %s\n", "lxi-tools")
    if err := pkg.Installlxi-tools(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg lxi-toolsFormula) Installlxi-tools() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "v2.7.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "v2.7.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
