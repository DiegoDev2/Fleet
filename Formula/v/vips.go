
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// vipsFormula represents a formula in Go.
type vipsFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg vipsFormula) Print() {
    fmt.Printf("Name: vips\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := vipsFormula{
        Description:  "Image processing library",
        Homepage:     "https://github.com/libvips/libvips",
        URL:          "https://github.com/libvips/libvips/releases/download/v8.15.3/vips-8.15.3.tar.xz",
        Sha256:       "07568eb2e6fffd0059fbc34f678825feb5815a2b2a29fed738865e91ba36e6cb",
        Dependencies: []string{"gobject-introspection", "meson", "ninja", "pkg-config", "cairo", "cfitsio", "cgif", "fftw", "fontconfig", "gettext", "glib", "highway", "imagemagick", "jpeg-xl", "libarchive", "libexif", "libheif", "libimagequant", "libmatio", "librsvg", "libspng", "libtiff", "little-cms2", "mozjpeg", "openexr", "openjpeg", "openslide", "pango", "poppler", "webp"},
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

    fmt.Printf("Installing formula: %s\n", "vips")
    if err := pkg.Installvips(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg vipsFormula) Installvips() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "vips-8.15.3.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "vips-8.15.3.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
