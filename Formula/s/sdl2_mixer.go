
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// sdl2_mixerFormula represents a formula in Go.
type sdl2_mixerFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg sdl2_mixerFormula) Print() {
    fmt.Printf("Name: sdl2_mixer\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := sdl2_mixerFormula{
        Description:  "Sample multi-channel audio mixer library",
        Homepage:     "https://github.com/libsdl-org/SDL_mixer",
        URL:          "https://github.com/libsdl-org/SDL_mixer.git",
        Sha256:       "1ad00dad0c5461fdb13e972fd1e9185e7acb4754fbe122ec53324a7f7e30dcda",
        Dependencies: []string{"autoconf", "automake", "libtool", "pkg-config", "flac", "fluid-synth", "game-music-emu", "libvorbis", "libxmp", "mpg123", "opusfile", "sdl2", "wavpack"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installsdl2_mixer(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg sdl2_mixerFormula) Installsdl2_mixer() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "SDL_mixer.git"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "SDL_mixer"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
