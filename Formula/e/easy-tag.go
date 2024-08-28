
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// easy-tagFormula represents a formula in Go.
type easy-tagFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg easy-tagFormula) Print() {
    fmt.Printf("Name: easy-tag\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := easy-tagFormula{
        Description:  "Application for viewing and editing audio file tags",
        Homepage:     "https://wiki.gnome.org/Apps/EasyTAG",
        URL:          "https://download.gnome.org/sources/easytag/2.4/easytag-2.4.3.tar.xz",
        Sha256:       "8fc76e69f070f3bc6353479a850ad4cb6884c182bdd1b998624a76c880a10d34",
        Dependencies: []string{"gettext", "intltool", "itstool", "pkg-config", "adwaita-icon-theme", "at-spi2-core", "cairo", "flac", "gdk-pixbuf", "glib", "gtk+3", "harfbuzz", "hicolor-icon-theme", "id3lib", "libid3tag", "libogg", "libvorbis", "pango", "speex", "taglib", "wavpack", "gettext", "perl-xml-parser"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installeasy-tag(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg easy-tagFormula) Installeasy-tag() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "easytag-2.4.3.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "easytag-2.4.3.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
