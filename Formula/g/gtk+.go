
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// gtk+Formula represents a formula in Go.
type gtk+Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg gtk+Formula) Print() {
    fmt.Printf("Name: gtk+\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := gtk+Formula{
        Description:  "GUI toolkit",
        Homepage:     "https://gtk.org/",
        URL:          "https://gitlab.gnome.org/GNOME/gtk/uploads/2a194d81de8e8346a81816870264b3bf/gdkimage.patch",
        Sha256:       "ce5adf1a019ac7ed2a999efb65cfadeae50f5de8663638c7f765f8764aa7d931",
        Dependencies: []string{"gobject-introspection", "pkg-config", "at-spi2-core", "cairo", "gdk-pixbuf", "glib", "hicolor-icon-theme", "pango", "gettext", "harfbuzz", "fontconfig", "libx11", "libxcomposite", "libxcursor", "libxdamage", "libxext", "libxfixes", "libxinerama", "libxrandr", "libxrender"},
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

    fmt.Printf("Installing formula: %s\n", "gtk+")
    if err := pkg.Installgtk+(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg gtk+Formula) Installgtk+() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "gdkimage.patch"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "gdkimage"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
