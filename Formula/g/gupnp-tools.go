
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// gupnp-toolsFormula represents a formula in Go.
type gupnp-toolsFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg gupnp-toolsFormula) Print() {
    fmt.Printf("Name: gupnp-tools\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := gupnp-toolsFormula{
        Description:  "Free replacements of Intel's UPnP tools",
        Homepage:     "https://wiki.gnome.org/GUPnP/",
        URL:          "https://gitlab.gnome.org/GNOME/gupnp-tools/-/commit/4e06104df81fba2cda06d4747b33e75f4cade458.diff",
        Sha256:       "a7e5c3ebf6dfd98fe17825b66b57ee40c839c19878261749f436676466faa945",
        Dependencies: []string{"meson", "ninja", "pkg-config", "gdk-pixbuf", "glib", "gssdp", "gtk+3", "gtksourceview4", "gupnp", "gupnp-av", "libsoup", "libxml2", "gettext"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installgupnp-tools(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg gupnp-toolsFormula) Installgupnp-tools() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "4e06104df81fba2cda06d4747b33e75f4cade458.diff"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "4e06104df81fba2cda06d4747b33e75f4cade458"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
