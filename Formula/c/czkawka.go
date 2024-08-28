
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// czkawkaFormula represents a formula in Go.
type czkawkaFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg czkawkaFormula) Print() {
    fmt.Printf("Name: czkawka\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := czkawkaFormula{
        Description:  "Duplicate file utility",
        Homepage:     "https://github.com/qarmin/czkawka",
        URL:          "https://github.com/qarmin/czkawka/archive/refs/tags/7.0.0.tar.gz",
        Sha256:       "3194269936ed396ae49ec4dbccf8ac6525af961a2fb98910c6d1efa916216f84",
        Dependencies: []string{"rust", "adwaita-icon-theme", "cairo", "ffmpeg", "gdk-pixbuf", "glib", "gtk4", "libheif", "librsvg", "pango", "pkg-config", "webp-pixbuf-loader", "gettext", "graphene", "harfbuzz"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installczkawka(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg czkawkaFormula) Installczkawka() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "7.0.0.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "7.0.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
