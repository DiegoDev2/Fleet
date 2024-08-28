
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// dspdfviewerFormula represents a formula in Go.
type dspdfviewerFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg dspdfviewerFormula) Print() {
    fmt.Printf("Name: dspdfviewer\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := dspdfviewerFormula{
        Description:  "Dual-Screen PDF Viewer for latex-beamer",
        Homepage:     "https://dspdfviewer.danny-edel.de/",
        URL:          "https://github.com/dannyedel/dspdfviewer/archive/refs/tags/v1.15.1.tar.gz",
        Sha256:       "2e95e14b53711e78ccd607afb0f44ba040f27b381d23cd4a7b6cec8ebb7d09ba",
        Dependencies: []string{"cmake", "gobject-introspection", "pkg-config", "boost", "cairo", "fontconfig", "freetype", "glib", "jpeg-turbo", "libpng", "libtiff", "openjpeg", "poppler-qt5", "qt@5", "gettext"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installdspdfviewer(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg dspdfviewerFormula) Installdspdfviewer() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "v1.15.1.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "v1.15.1.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
