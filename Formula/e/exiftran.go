
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// exiftranFormula represents a formula in Go.
type exiftranFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg exiftranFormula) Print() {
    fmt.Printf("Name: exiftran\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := exiftranFormula{
        Description:  "Transform digital camera jpegs and their EXIF data",
        Homepage:     "https://www.kraxel.org/blog/linux/fbida/",
        URL:          "https://raw.githubusercontent.com/Homebrew/formula-patches/185c281/exiftran/fix-build.diff",
        Sha256:       "017268a3195fb52df08ed75827fa40e8179aff0d9e54c926b0ace5f8e60702bf",
        Dependencies: []string{"pkg-config", "jpeg-turbo", "libexif", "pixman", "cairo", "fontconfig", "freetype", "ghostscript", "giflib", "libdrm", "libepoxy", "libpng", "libtiff", "libx11", "libxext", "libxpm", "libxt", "mesa", "openmotif", "poppler", "webp"},
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

    fmt.Printf("Installing formula: %s\n", "exiftran")
    if err := pkg.Installexiftran(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg exiftranFormula) Installexiftran() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "fix-build.diff"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "fix-build"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
