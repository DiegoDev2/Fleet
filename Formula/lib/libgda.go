
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// libgdaFormula represents a formula in Go.
type libgdaFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg libgdaFormula) Print() {
    fmt.Printf("Name: libgda\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := libgdaFormula{
        Description:  "Provides unified data access to the GNOME project",
        Homepage:     "https://www.gnome-db.org/",
        URL:          "https://gitlab.gnome.org/GNOME/libgda/-/commit/98f014c783583e3ad87ee546e8dccf34d50f1e37.diff",
        Sha256:       "2f2d257085b40ef4fccf2db68fe51407ba0f59d39672fc95fd91be3e46e91ffa",
        Dependencies: []string{"gettext", "gobject-introspection", "intltool", "meson", "ninja", "pkg-config", "vala", "glib", "iso-codes", "json-glib", "sqlite", "gettext"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installlibgda(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg libgdaFormula) Installlibgda() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "98f014c783583e3ad87ee546e8dccf34d50f1e37.diff"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "98f014c783583e3ad87ee546e8dccf34d50f1e37"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
