
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// opencl-clhpp-headersFormula represents a formula in Go.
type opencl-clhpp-headersFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg opencl-clhpp-headersFormula) Print() {
    fmt.Printf("Name: opencl-clhpp-headers\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := opencl-clhpp-headersFormula{
        Description:  "C++ language header files for the OpenCL API",
        Homepage:     "https://www.khronos.org/registry/OpenCL/",
        URL:          "https://github.com/KhronosGroup/OpenCL-CLHPP/archive/refs/tags/v2024.05.08.tar.gz",
        Sha256:       "669c7dcb52412558cfe3d8ac276ff1c29af9d1121b54ef20847cfe30f3318bb7",
        Dependencies: []string{"cmake", "opencl-headers"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installopencl-clhpp-headers(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg opencl-clhpp-headersFormula) Installopencl-clhpp-headers() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "v2024.05.08.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "v2024.05.08.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
