
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// clblasFormula represents a formula in Go.
type clblasFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg clblasFormula) Print() {
    fmt.Printf("Name: clblas\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := clblasFormula{
        Description:  "Library containing BLAS functions written in OpenCL",
        Homepage:     "https://github.com/clMathLibraries/clBLAS",
        URL:          "https://github.com/clMathLibraries/clBLAS/commit/68ce5f0b824d7cf9d71b09bb235cf219defcc7b4.patch?full_index=1",
        Sha256:       "df5dc87e9ae543a043608cf790d01b985627b5b6355356c860cfd45a47ba2c36",
        Dependencies: []string{"boost", "cmake", "opencl-headers", "opencl-icd-loader", "pocl"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installclblas(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg clblasFormula) Installclblas() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "68ce5f0b824d7cf9d71b09bb235cf219defcc7b4.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "68ce5f0b824d7cf9d71b09bb235cf219defcc7b4"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
