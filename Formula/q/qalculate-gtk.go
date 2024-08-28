
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// qalculate-gtkFormula represents a formula in Go.
type qalculate-gtkFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg qalculate-gtkFormula) Print() {
    fmt.Printf("Name: qalculate-gtk\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := qalculate-gtkFormula{
        Description:  "Multi-purpose desktop calculator",
        Homepage:     "https://qalculate.github.io/",
        URL:          "https://github.com/Qalculate/qalculate-gtk/releases/download/v5.2.0/qalculate-gtk-5.2.0.tar.gz",
        Sha256:       "d8367367da463d8abb883fe8d6feae317dc755a304f6291e88038f2cf3208bb7",
        Dependencies: []string{"intltool", "pkg-config", "adwaita-icon-theme", "cairo", "gdk-pixbuf", "glib", "gtk+3", "libqalculate", "pango", "at-spi2-core", "gettext", "harfbuzz", "perl-xml-parser"},
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

    fmt.Printf("Installing formula: %s\n", "qalculate-gtk")
    if err := pkg.Installqalculate-gtk(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg qalculate-gtkFormula) Installqalculate-gtk() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "qalculate-gtk-5.2.0.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "qalculate-gtk-5.2.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
