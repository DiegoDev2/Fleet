
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// gladeFormula represents a formula in Go.
type gladeFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg gladeFormula) Print() {
    fmt.Printf("Name: glade\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := gladeFormula{
        Description:  "RAD tool for the GTK+ and GNOME environment",
        Homepage:     "https://glade.gnome.org/",
        URL:          "https://download.gnome.org/sources/glade/3.40/glade-3.40.0.tar.xz",
        Sha256:       "a68666c43b3c79b9532742a2351093d587fe0da66c369ba3e62ed35f223c4deb",
        Dependencies: []string{"docbook-xsl", "gobject-introspection", "itstool", "meson", "ninja", "pkg-config", "adwaita-icon-theme", "cairo", "gdk-pixbuf", "gettext", "glib", "gtk+3", "hicolor-icon-theme", "libxml2", "pango", "at-spi2-core", "gtk-mac-integration", "harfbuzz"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installglade(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg gladeFormula) Installglade() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "glade-3.40.0.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "glade-3.40.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
