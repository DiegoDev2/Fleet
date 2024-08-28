
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// geocode-glibFormula represents a formula in Go.
type geocode-glibFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg geocode-glibFormula) Print() {
    fmt.Printf("Name: geocode-glib\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := geocode-glibFormula{
        Description:  "GNOME library for gecoding and reverse geocoding",
        Homepage:     "https://gitlab.gnome.org/GNOME/geocode-glib",
        URL:          "https://download.gnome.org/sources/geocode-glib/3.26/geocode-glib-3.26.4.tar.xz",
        Sha256:       "705672b2c649c9dad5061d9d010d6faa106f67a278e90eab7c6b6a7a8f66e9ca",
        Dependencies: []string{"gobject-introspection", "meson", "ninja", "pkg-config", "glib", "gtk+3", "json-glib", "libsoup", "gettext"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installgeocode-glib(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg geocode-glibFormula) Installgeocode-glib() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "geocode-glib-3.26.4.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "geocode-glib-3.26.4.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
