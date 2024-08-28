
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// geditFormula represents a formula in Go.
type geditFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg geditFormula) Print() {
    fmt.Printf("Name: gedit\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := geditFormula{
        Description:  "GNOME text editor",
        Homepage:     "https://gedit-technology.github.io/apps/gedit/",
        URL:          "https://download.gnome.org/sources/gedit/47/gedit-47.0.tar.xz",
        Sha256:       "d44f0ce398f5a58684a6797d457a6056b4f976e8bce80d93b14b4e831a94e79b",
        Dependencies: []string{"desktop-file-utils", "docbook-xsl", "gettext", "gtk-doc", "itstool", "meson", "ninja", "pkg-config", "adwaita-icon-theme", "cairo", "gdk-pixbuf", "glib", "gobject-introspection", "gsettings-desktop-schemas", "gspell", "gtk+3", "libgedit-amtk", "libgedit-gfls", "libgedit-gtksourceview", "libgedit-tepl", "libpeas@1", "libxml2", "pango", "gettext", "gtk-mac-integration"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installgedit(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg geditFormula) Installgedit() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "gedit-47.0.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "gedit-47.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
