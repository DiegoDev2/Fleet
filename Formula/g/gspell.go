
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// gspellFormula represents a formula in Go.
type gspellFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg gspellFormula) Print() {
    fmt.Printf("Name: gspell\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := gspellFormula{
        Description:  "Flexible API to implement spellchecking in GTK+ applications",
        Homepage:     "https://gitlab.gnome.org/GNOME/gspell",
        URL:          "https://download.gnome.org/sources/gspell/1.12/gspell-1.12.2.tar.xz",
        Sha256:       "1e74d8a181bc1e6fc55d1eb36a737728077969bf83e6ea38ce643abf50df2a3b",
        Dependencies: []string{"gobject-introspection", "pkg-config", "vala", "at-spi2-core", "cairo", "enchant", "gdk-pixbuf", "glib", "gtk+3", "harfbuzz", "icu4c", "pango", "autoconf", "automake", "gtk-doc", "libtool", "gettext", "gtk-mac-integration"},
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

    fmt.Printf("Installing formula: %s\n", "gspell")
    if err := pkg.Installgspell(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg gspellFormula) Installgspell() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "gspell-1.12.2.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "gspell-1.12.2.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
