
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// one-mlFormula represents a formula in Go.
type one-mlFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg one-mlFormula) Print() {
    fmt.Printf("Name: one-ml\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := one-mlFormula{
        Description:  "Reboot of ML, unifying its core and (now first-class) module layers",
        Homepage:     "https://people.mpi-sws.org/~rossberg/1ml/",
        URL:          "https://github.com/rossberg/1ml/commit/f99c0b3497c1f18c950dfb2ae3989573f90eaafd.patch?full_index=1",
        Sha256:       "778c9635f170a29fa6a53358e65fe85f32320eb678683ddd23e0e2c6139e7a6e",
        Dependencies: []string{"ocaml"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installone-ml(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg one-mlFormula) Installone-ml() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "f99c0b3497c1f18c950dfb2ae3989573f90eaafd.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "f99c0b3497c1f18c950dfb2ae3989573f90eaafd"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
