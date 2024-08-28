
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// at-spi2-coreFormula represents a formula in Go.
type at-spi2-coreFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg at-spi2-coreFormula) Print() {
    fmt.Printf("Name: at-spi2-core\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := at-spi2-coreFormula{
        Description:  "Protocol definitions and daemon for D-Bus at-spi",
        Homepage:     "https://www.freedesktop.org/wiki/Accessibility/AT-SPI2",
        URL:          "https://download.gnome.org/sources/at-spi2-core/2.52/at-spi2-core-2.52.0.tar.xz",
        Sha256:       "3163183968f02b39e5476ac538b74f5e0eb5ac8a613e53da0bd90380cfd2ae30",
        Dependencies: []string{"gettext", "gobject-introspection", "meson", "ninja", "pkg-config", "dbus", "glib", "libx11", "libxi", "libxtst", "xorgproto", "gettext"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installat-spi2-core(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg at-spi2-coreFormula) Installat-spi2-core() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "at-spi2-core-2.52.0.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "at-spi2-core-2.52.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
