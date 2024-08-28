
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// libgedit-amtkFormula represents a formula in Go.
type libgedit-amtkFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg libgedit-amtkFormula) Print() {
    fmt.Printf("Name: libgedit-amtk\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := libgedit-amtkFormula{
        Description:  "Actions, Menus and Toolbars Kit for GTK applications",
        Homepage:     "https://gedit-technology.net",
        URL:          "https://gitlab.gnome.org/World/gedit/libgedit-amtk/-/archive/5.8.0/libgedit-amtk-5.8.0.tar.bz2",
        Sha256:       "61d99af175776561c482cbffa60cdc2438366b67519a9e69e06144deaaae556f",
        Dependencies: []string{"gobject-introspection", "meson", "ninja", "pkg-config", "glib", "gtk+3", "at-spi2-core", "cairo", "gdk-pixbuf", "gettext", "harfbuzz", "pango"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installlibgedit-amtk(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg libgedit-amtkFormula) Installlibgedit-amtk() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "libgedit-amtk-5.8.0.tar.bz2"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "libgedit-amtk-5.8.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
