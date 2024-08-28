
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// dronedbFormula represents a formula in Go.
type dronedbFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg dronedbFormula) Print() {
    fmt.Printf("Name: dronedb\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := dronedbFormula{
        Description:  "Free and open source software for aerial data storage",
        Homepage:     "https://github.com/DroneDB/DroneDB",
        URL:          "https://github.com/DroneDB/DroneDB/commit/28aa869dee5920c2d948e1b623f2f9d518bdcb1e.patch?full_index=1",
        Sha256:       "50e581aad0fd3226fe5999cc91f9a61fdcbc42c5ba2394d9def89b70183f9c96",
        Dependencies: []string{"cmake", "gdal", "libspatialite", "libzip", "pdal"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installdronedb(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg dronedbFormula) Installdronedb() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "28aa869dee5920c2d948e1b623f2f9d518bdcb1e.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "28aa869dee5920c2d948e1b623f2f9d518bdcb1e"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
