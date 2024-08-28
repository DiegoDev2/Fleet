
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// libsoup@2Formula represents a formula in Go.
type libsoup@2Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg libsoup@2Formula) Print() {
    fmt.Printf("Name: libsoup@2\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := libsoup@2Formula{
        Description:  "HTTP client/server library for GNOME",
        Homepage:     "https://wiki.gnome.org/Projects/libsoup",
        URL:          "https://download.gnome.org/sources/libsoup/2.74/libsoup-2.74.2.tar.xz",
        Sha256:       "416f1242c8f4e786d4e3fb673a40926b9c5543bf50420bf9dbf24cae569743f2",
        Dependencies: []string{"gobject-introspection", "meson", "ninja", "pkg-config", "python@3.12", "vala", "icu4c", "glib", "glib-networking", "gnutls", "libpsl", "sqlite", "gettext", "brotli"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installlibsoup@2(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg libsoup@2Formula) Installlibsoup@2() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "libsoup-2.74.2.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "libsoup-2.74.2.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
