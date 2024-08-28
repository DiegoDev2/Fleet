
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// frotzFormula represents a formula in Go.
type frotzFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg frotzFormula) Print() {
    fmt.Printf("Name: frotz\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := frotzFormula{
        Description:  "Infocom-style interactive fiction player",
        Homepage:     "https://661.org/proj/if/frotz/",
        URL:          "https://gitlab.com/DavidGriffith/frotz/-/raw/2.53/src/test/etude/etude.z5",
        Sha256:       "bfa2ef69f2f5ce3796b96f9b073676902e971aedb3ba690b8835bb1fb0daface",
        Dependencies: []string{"pkg-config", "freetype", "jpeg-turbo", "libao", "libmodplug", "libpng", "libsamplerate", "libsndfile", "libvorbis", "ncurses", "sdl2", "sdl2_mixer"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installfrotz(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg frotzFormula) Installfrotz() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "etude.z5"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "etude"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
