
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// prefixsuffixFormula represents a formula in Go.
type prefixsuffixFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg prefixsuffixFormula) Print() {
    fmt.Printf("Name: prefixsuffix\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := prefixsuffixFormula{
        Description:  "GUI batch renaming utility",
        Homepage:     "https://github.com/murraycu/prefixsuffix",
        URL:          "https://download.gnome.org/sources/prefixsuffix/0.6/prefixsuffix-0.6.9.tar.xz",
        Sha256:       "1bd8de221d43b7d0d511a7a0cb6ebd3a35f045524a02917cf839eb426ae65d41",
        Dependencies: []string{"gettext", "intltool", "pkg-config", "atkmm@2.28", "glib", "glibmm@2.66", "gtk+3", "gtkmm3", "libsigc++@2", "at-spi2-core", "cairo", "cairomm@1.14", "gdk-pixbuf", "gettext", "harfbuzz", "pango", "pangomm@2.46", "perl-xml-parser"},
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

    fmt.Printf("Installing formula: %s\n", "prefixsuffix")
    if err := pkg.Installprefixsuffix(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg prefixsuffixFormula) Installprefixsuffix() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "prefixsuffix-0.6.9.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "prefixsuffix-0.6.9.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
