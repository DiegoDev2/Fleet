
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// gsettings-desktop-schemasFormula represents a formula in Go.
type gsettings-desktop-schemasFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg gsettings-desktop-schemasFormula) Print() {
    fmt.Printf("Name: gsettings-desktop-schemas\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := gsettings-desktop-schemasFormula{
        Description:  "GSettings schemas for desktop components",
        Homepage:     "https://download.gnome.org/sources/gsettings-desktop-schemas/",
        URL:          "https://download.gnome.org/sources/gsettings-desktop-schemas/46/gsettings-desktop-schemas-46.1.tar.xz",
        Sha256:       "872eb6f9edaf59999a3fb88a29133ca93eac0c4561a54fe7b84686bf3e1fb381",
        Dependencies: []string{"gettext", "gobject-introspection", "meson", "ninja", "pkg-config", "glib"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installgsettings-desktop-schemas(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg gsettings-desktop-schemasFormula) Installgsettings-desktop-schemas() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "gsettings-desktop-schemas-46.1.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "gsettings-desktop-schemas-46.1.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
