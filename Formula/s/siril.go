
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// sirilFormula represents a formula in Go.
type sirilFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg sirilFormula) Print() {
    fmt.Printf("Name: siril\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := sirilFormula{
        Description:  "Astronomical image processing tool",
        Homepage:     "https://www.siril.org",
        URL:          "https://free-astro.org/download/siril-1.2.3.tar.bz2",
        Sha256:       "de834dbcab8bc30d5beecf176ef0ead7535d643f8620a258d839b2753656540f",
        Dependencies: []string{"cmake", "meson", "ninja", "pkg-config", "adwaita-icon-theme", "cairo", "cfitsio", "exiv2", "ffmpeg", "ffms2", "fftw", "gdk-pixbuf", "glib", "gnuplot", "gsl", "gtk+3", "jpeg-turbo", "json-glib", "libconfig", "libheif", "libpng", "libraw", "librsvg", "libtiff", "netpbm", "opencv", "openjpeg", "pango", "wcslib", "gettext", "gtk-mac-integration", "libomp"},
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

    fmt.Printf("Installing formula: %s\n", "siril")
    if err := pkg.Installsiril(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg sirilFormula) Installsiril() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "siril-1.2.3.tar.bz2"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "siril-1.2.3.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
