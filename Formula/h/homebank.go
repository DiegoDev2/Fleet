
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// homebankFormula represents a formula in Go.
type homebankFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg homebankFormula) Print() {
    fmt.Printf("Name: homebank\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := homebankFormula{
        Description:  "Manage your personal accounts at home",
        Homepage:     "http://homebank.free.fr",
        URL:          "https://deb.debian.org/debian/pool/main/h/homebank/",
        Sha256:       "35552c6edb31ec7f2f966cd152ca32cd8f32dac5ea6737dc8c01267542bcaba2",
        Dependencies: []string{"intltool", "pkg-config", "adwaita-icon-theme", "cairo", "fontconfig", "freetype", "gdk-pixbuf", "glib", "gtk+3", "hicolor-icon-theme", "libofx", "libsoup", "pango", "at-spi2-core", "gettext", "harfbuzz", "perl-xml-parser"},
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

    fmt.Printf("Installing formula: %s\n", "homebank")
    if err := pkg.Installhomebank(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg homebankFormula) Installhomebank() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := ""
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := ""
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
