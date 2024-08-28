
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// gdlFormula represents a formula in Go.
type gdlFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg gdlFormula) Print() {
    fmt.Printf("Name: gdl\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := gdlFormula{
        Description:  "GNOME Docking Library provides docking features for GTK+ 3",
        Homepage:     "https://gitlab.gnome.org/GNOME/gdl",
        URL:          "https://gitlab.gnome.org/GNOME/gdl/-/commit/414f83eb4ad9e5576ee3d089594bf1301ff24091.diff",
        Sha256:       "715c804e6d03304bc077b99f667bbeb062c873b3bbd737182fb2cd47a295de95",
        Dependencies: []string{"gettext", "gobject-introspection", "intltool", "pkg-config", "cairo", "gdk-pixbuf", "glib", "gtk+3", "libxml2", "at-spi2-core", "gettext", "harfbuzz", "pango", "perl-xml-parser"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installgdl(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg gdlFormula) Installgdl() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "414f83eb4ad9e5576ee3d089594bf1301ff24091.diff"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "414f83eb4ad9e5576ee3d089594bf1301ff24091"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
