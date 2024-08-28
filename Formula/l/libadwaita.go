
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// libadwaitaFormula represents a formula in Go.
type libadwaitaFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg libadwaitaFormula) Print() {
    fmt.Printf("Name: libadwaita\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := libadwaitaFormula{
        Description:  "Building blocks for modern adaptive GNOME applications",
        Homepage:     "https://gnome.pages.gitlab.gnome.org/libadwaita/",
        URL:          "https://download.gnome.org/sources/libadwaita/1.5/libadwaita-1.5.3.tar.xz",
        Sha256:       "b8efe52d28183d988cc25e27fd5fc81c793f6efabc6dbb2e0e543ab511bc31cb",
        Dependencies: []string{"gettext", "gobject-introspection", "meson", "ninja", "pkg-config", "vala", "appstream", "fribidi", "glib", "graphene", "gtk4", "pango", "gettext"},
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

    fmt.Printf("Installing formula: %s\n", "libadwaita")
    if err := pkg.Installlibadwaita(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg libadwaitaFormula) Installlibadwaita() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "libadwaita-1.5.3.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "libadwaita-1.5.3.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
