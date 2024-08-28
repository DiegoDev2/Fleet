
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// graphicsmagickFormula represents a formula in Go.
type graphicsmagickFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg graphicsmagickFormula) Print() {
    fmt.Printf("Name: graphicsmagick\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := graphicsmagickFormula{
        Description:  "Image processing tools collection",
        Homepage:     "http://www.graphicsmagick.org/",
        URL:          "https://sourceforge.net/projects/graphicsmagick/rss?path=/graphicsmagick",
        Sha256:       "0561ac8547215fedc9408cd0c6e6a9e3927302212219b5e89b3bab0c7ae685eb",
        Dependencies: []string{"pkg-config", "freetype", "jasper", "jpeg-turbo", "jpeg-xl", "libheif", "libpng", "libtiff", "libtool", "little-cms2", "webp", "zstd"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installgraphicsmagick(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg graphicsmagickFormula) Installgraphicsmagick() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "graphicsmagick"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "graphicsmagick"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
