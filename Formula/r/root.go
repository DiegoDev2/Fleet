
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// rootFormula represents a formula in Go.
type rootFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg rootFormula) Print() {
    fmt.Printf("Name: root\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := rootFormula{
        Description:  "Object oriented framework for large scale data analysis",
        Homepage:     "https://root.cern.ch/",
        URL:          "https://root.cern/install/all_releases/",
        Sha256:       "05df68a4414caec551cc238f46612f370ebfb05194ace9fbdbda98d87d4788a5",
        Dependencies: []string{"cmake", "ninja", "pkg-config", "cfitsio", "davix", "fftw", "freetype", "ftgl", "gcc", "gl2ps", "glew", "graphviz", "gsl", "lz4", "mysql-client", "nlohmann-json", "numpy", "openblas", "openssl@3", "pcre", "pcre2", "python@3.12", "sqlite", "tbb", "xrootd", "xxhash", "xz", "zlib", "zstd", "giflib", "jpeg-turbo", "libpng", "libtiff", "libx11", "libxext", "libxft", "libxpm", "mesa", "mesa-glu"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installroot(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg rootFormula) Installroot() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := ""
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := ""
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
