
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// adios2Formula represents a formula in Go.
type adios2Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg adios2Formula) Print() {
    fmt.Printf("Name: adios2\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := adios2Formula{
        Description:  "Next generation of ADIOS developed in the Exascale Computing Program",
        Homepage:     "https://adios2.readthedocs.io",
        URL:          "https://github.com/ornladios/ADIOS2/archive/refs/tags/v2.10.1.tar.gz",
        Sha256:       "d9f8c709fc2a57f9bd8970160862a10b2070fe87cb30705e0758ca1317b17cc4",
        Dependencies: []string{"cmake", "nlohmann-json", "c-blosc", "gcc", "libfabric", "libpng", "libsodium", "mpi4py", "numpy", "open-mpi", "pugixml", "pybind11", "python@3.12", "sqlite", "yaml-cpp", "zeromq", "llvm"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installadios2(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg adios2Formula) Installadios2() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "v2.10.1.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "v2.10.1.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
