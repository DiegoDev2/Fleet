
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// spatialite-guiFormula represents a formula in Go.
type spatialite-guiFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg spatialite-guiFormula) Print() {
    fmt.Printf("Name: spatialite-gui\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := spatialite-guiFormula{
        Description:  "GUI tool supporting SpatiaLite",
        Homepage:     "https://www.gaia-gis.it/fossil/spatialite_gui/index",
        URL:          "https://www.gaia-gis.it/gaia-sins/spatialite-gui-sources/",
        Sha256:       "4da054619064fa91a7f6c55002b6d2d08eef58ac63040a469dfed8e7358ad746",
        Dependencies: []string{"pkg-config", "freexl", "geos", "libpq", "librasterlite2", "librttopo", "libspatialite", "libtiff", "libxlsxwriter", "libxml2", "lz4", "minizip", "openjpeg", "proj", "sqlite", "virtualpg", "webp", "wxwidgets", "xz", "zstd"},
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

    fmt.Printf("Installing formula: %s\n", "spatialite-gui")
    if err := pkg.Installspatialite-gui(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg spatialite-guiFormula) Installspatialite-gui() error {
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
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
