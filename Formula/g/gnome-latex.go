
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// gnome-latexFormula represents a formula in Go.
type gnome-latexFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg gnome-latexFormula) Print() {
    fmt.Printf("Name: gnome-latex\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := gnome-latexFormula{
        Description:  "LaTeX editor for the GNOME desktop",
        Homepage:     "https://gitlab.gnome.org/swilmet/gedit-tex",
        URL:          "https://gitlab.gnome.org/swilmet/gedit-tex.git",
        Sha256:       "8281dbcc2d48a6b955f832a01947da7b5b27b536b2e650c36cddf6ec807e1116",
        Dependencies: []string{"autoconf", "automake", "gtk-doc", "libtool", "desktop-file-utils", "meson", "ninja", "vala", "gettext", "gobject-introspection", "itstool", "pkg-config", "adwaita-icon-theme", "gdk-pixbuf", "glib", "gspell", "gtk+3", "libgedit-amtk", "libgedit-gtksourceview", "libgedit-tepl", "libgee", "pango", "at-spi2-core", "cairo", "enchant", "gettext", "harfbuzz", "libgedit-gfls"},
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

    fmt.Printf("Installing formula: %s\n", "gnome-latex")
    if err := pkg.Installgnome-latex(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg gnome-latexFormula) Installgnome-latex() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "gedit-tex.git"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "gedit-tex"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
