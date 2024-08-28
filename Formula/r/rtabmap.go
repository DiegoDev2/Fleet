
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// rtabmapFormula represents a formula in Go.
type rtabmapFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg rtabmapFormula) Print() {
    fmt.Printf("Name: rtabmap\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := rtabmapFormula{
        Description:  "Visual and LiDAR SLAM library and standalone application",
        Homepage:     "https://introlab.github.io/rtabmap",
        URL:          "https://github.com/introlab/rtabmap/archive/refs/tags/0.21.4.tar.gz",
        Sha256:       "a66f6cb1a14b01eaf94fa9fe15b28645ead7246d98e10cd2713a1a166f78cbe0",
        Dependencies: []string{"cmake", "g2o", "librealsense", "octomap", "opencv", "pcl", "pdal", "qt", "sqlite", "vtk", "boost", "flann", "glew", "libomp", "libpcap", "libpng", "lz4", "qhull"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installrtabmap(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg rtabmapFormula) Installrtabmap() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "0.21.4.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "0.21.4.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
