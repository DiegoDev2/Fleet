
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// openclonkFormula represents a formula in Go.
type openclonkFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg openclonkFormula) Print() {
    fmt.Printf("Name: openclonk\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := openclonkFormula{
        Description:  "Multiplayer action game",
        Homepage:     "https://www.openclonk.org/",
        URL:          "https://github.com/openclonk/openclonk.git",
        Sha256:       "352052e8b19d4d3f9e7f311ae5f97353ad1c6ba8cb8ee5c43ead3f7bb8cf5fbb",
        Dependencies: []string{"glew", "gtk+3", "libx11", "libepoxy", "miniupnpc", "sdl2", "cmake", "freealut", "freetype", "jpeg-turbo", "libogg", "libpng", "libvorbis", "tinyxml", "pkg-config", "libxrandr", "mesa", "openal-soft"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installopenclonk(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg openclonkFormula) Installopenclonk() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "openclonk.git"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "openclonk"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
