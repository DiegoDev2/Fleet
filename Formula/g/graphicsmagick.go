
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
        if !isFormulaInstalled(dep) {
            fmt.Printf("Installing dependency: %s\n", dep)
            cmd := exec.Command("brew", "install", dep)
            if err := cmd.Run(); err != nil {
                log.Fatalf("Error installing dependency %s: %v", dep, err)
            }
        } else {
            fmt.Printf("Dependency %s is already installed.\n", dep)
        }
    }

    fmt.Printf("Installing formula: %s\n", "graphicsmagick")
    if err := pkg.Installgraphicsmagick(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
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
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
