
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// gdalFormula represents a formula in Go.
type gdalFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg gdalFormula) Print() {
    fmt.Printf("Name: gdal\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := gdalFormula{
        Description:  "Geospatial Data Abstraction Library",
        Homepage:     "https://www.gdal.org/",
        URL:          "https://github.com/OSGeo/gdal.git",
        Sha256:       "76dcf708838db8f957a789284b375c207300149e8f8e355f6bcaec19162793fe",
        Dependencies: []string{"doxygen", "boost", "cmake", "pkg-config", "python-setuptools", "swig", "apache-arrow", "cfitsio", "epsilon", "expat", "freexl", "geos", "giflib", "hdf5", "imath", "jpeg-turbo", "jpeg-xl", "json-c", "libaec", "libarchive", "libgeotiff", "libheif", "libkml", "liblerc", "libpng", "libpq", "libspatialite", "libtiff", "libxml2", "lz4", "netcdf", "numpy", "openexr", "openjpeg", "openssl@3", "pcre2", "poppler", "proj", "python@3.12", "qhull", "sqlite", "unixodbc", "webp", "xerces-c", "xz", "zstd", "minizip", "uriparser", "util-linux"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installgdal(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg gdalFormula) Installgdal() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "gdal.git"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "gdal"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
