
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// manimFormula represents a formula in Go.
type manimFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg manimFormula) Print() {
    fmt.Printf("Name: manim\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := manimFormula{
        Description:  "Animation engine for explanatory math videos",
        Homepage:     "https://www.manim.community",
        URL:          "https://files.pythonhosted.org/packages/1b/f9/b01e4632aed9a6ecc2b3e501feffd3af5aa0eb4e3b0283fc9525bf503c38/watchdog-4.0.1.tar.gz",
        Sha256:       "eebaacf674fa25511e8867028d281e602ee6500045b57f43b08778082f7f8b44",
        Dependencies: []string{"cython", "ninja", "pkg-config", "cairo", "ffmpeg", "fontconfig", "freetype", "glib", "numpy", "pango", "pillow", "py3cairo", "python@3.12", "scipy", "gettext", "harfbuzz", "cmake", "patchelf"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installmanim(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg manimFormula) Installmanim() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "watchdog-4.0.1.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "watchdog-4.0.1.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
