
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// freedinkFormula represents a formula in Go.
type freedinkFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg freedinkFormula) Print() {
    fmt.Printf("Name: freedink\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := freedinkFormula{
        Description:  "Portable version of the Dink Smallwood game engine",
        Homepage:     "https://www.gnu.org/software/freedink/",
        URL:          "https://raw.githubusercontent.com/openbsd/ports/fc8b95c6/games/freedink/game/patches/patch-src_input_cpp",
        Sha256:       "fa06a8a87bd4f3977440cdde0fb6145b6e5b0005b266b19c059d3fd7c2ff836a",
        Dependencies: []string{"glm", "pkg-config", "check", "cxxtest", "fontconfig", "freetype", "gettext", "libzip", "sdl2", "sdl2_gfx", "sdl2_image", "sdl2_mixer", "sdl2_ttf"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installfreedink(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg freedinkFormula) Installfreedink() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "patch-src_input_cpp"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "patch-src_input_cpp"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
