
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// baobabFormula represents a formula in Go.
type baobabFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg baobabFormula) Print() {
    fmt.Printf("Name: baobab\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := baobabFormula{
        Description:  "Gnome disk usage analyzer",
        Homepage:     "https://apps.gnome.org/Baobab/",
        URL:          "https://download.gnome.org/sources/baobab/46/baobab-46.0.tar.xz",
        Sha256:       "dd385933494b6cb5e6a6b0b911c83f4c6cf9922d8a17830d61cca27d763cf659",
        Dependencies: []string{"desktop-file-utils", "gettext", "itstool", "meson", "ninja", "pkg-config", "vala", "adwaita-icon-theme", "cairo", "glib", "gtk4", "hicolor-icon-theme", "libadwaita", "pango", "gettext"},
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

    fmt.Printf("Installing formula: %s\n", "baobab")
    if err := pkg.Installbaobab(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg baobabFormula) Installbaobab() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "baobab-46.0.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "baobab-46.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
