
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// crispy-doomFormula represents a formula in Go.
type crispy-doomFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg crispy-doomFormula) Print() {
    fmt.Printf("Name: crispy-doom\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := crispy-doomFormula{
        Description:  "Limit-removing enhanced-resolution Doom source port based on Chocolate Doom",
        Homepage:     "https://github.com/fabiangreffrath/crispy-doom",
        URL:          "https://github.com/fabiangreffrath/crispy-doom/archive/refs/tags/crispy-doom-7.0.tar.gz",
        Sha256:       "32038b4fcd8556a79441e0f482ea6597e1ae6a8ec934a415c7fadc62ed21cf22",
        Dependencies: []string{"autoconf", "automake", "pkg-config", "fluid-synth", "libpng", "libsamplerate", "sdl2", "sdl2_mixer", "sdl2_net"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installcrispy-doom(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg crispy-doomFormula) Installcrispy-doom() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "crispy-doom-7.0.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "crispy-doom-7.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
