
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// libgedit-gflsFormula represents a formula in Go.
type libgedit-gflsFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg libgedit-gflsFormula) Print() {
    fmt.Printf("Name: libgedit-gfls\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := libgedit-gflsFormula{
        Description:  "Gedit Technology - File loading and saving",
        Homepage:     "https://gitlab.gnome.org/World/gedit/libgedit-gfls",
        URL:          "https://gitlab.gnome.org/World/gedit/libgedit-gfls/-/archive/0.1.0/libgedit-gfls-0.1.0.tar.bz2",
        Sha256:       "0e03f0d1f6337bb1ee3cb02a277f824ccf514b0ee1e23d6bc6ebca5336ec6c41",
        Dependencies: []string{"gobject-introspection", "meson", "ninja", "pkg-config", "glib", "gtk+3"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installlibgedit-gfls(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg libgedit-gflsFormula) Installlibgedit-gfls() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "libgedit-gfls-0.1.0.tar.bz2"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "libgedit-gfls-0.1.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
