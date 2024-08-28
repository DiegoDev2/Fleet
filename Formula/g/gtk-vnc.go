
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// gtk-vncFormula represents a formula in Go.
type gtk-vncFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg gtk-vncFormula) Print() {
    fmt.Printf("Name: gtk-vnc\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := gtk-vncFormula{
        Description:  "VNC viewer widget for GTK",
        Homepage:     "https://wiki.gnome.org/Projects/gtk-vnc",
        URL:          "https://gitlab.gnome.org/GNOME/gtk-vnc/-/commit/40c59c557ecf7d22d307b8cf890ce08b0376ca5a.diff",
        Sha256:       "795f35a50bb4a1976d05b961a708415246a2c3f1164c1f9d4e7931996af9f706",
        Dependencies: []string{"gettext", "gobject-introspection", "meson", "ninja", "pkg-config", "cairo", "gdk-pixbuf", "glib", "gnutls", "gtk+3", "libgcrypt", "gettext", "libx11"},
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

    fmt.Printf("Installing formula: %s\n", "gtk-vnc")
    if err := pkg.Installgtk-vnc(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg gtk-vncFormula) Installgtk-vnc() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "40c59c557ecf7d22d307b8cf890ce08b0376ca5a.diff"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "40c59c557ecf7d22d307b8cf890ce08b0376ca5a"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
