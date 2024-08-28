
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// pkcs11-helperFormula represents a formula in Go.
type pkcs11-helperFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg pkcs11-helperFormula) Print() {
    fmt.Printf("Name: pkcs11-helper\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := pkcs11-helperFormula{
        Description:  "Library to simplify the interaction with PKCS#11",
        Homepage:     "https://github.com/OpenSC/OpenSC/wiki/pkcs11-helper",
        URL:          "https://github.com/OpenSC/pkcs11-helper/releases/download/pkcs11-helper-1.30.0/pkcs11-helper-1.30.0.tar.bz2",
        Sha256:       "3ce66c1a7dcf6725cf1b4f480e734cd4b7ff2f74835b4464c2bf182ce9640d1b",
        Dependencies: []string{"autoconf", "automake", "libtool", "pkg-config", "openssl@3"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installpkcs11-helper(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg pkcs11-helperFormula) Installpkcs11-helper() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "pkcs11-helper-1.30.0.tar.bz2"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "pkcs11-helper-1.30.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
