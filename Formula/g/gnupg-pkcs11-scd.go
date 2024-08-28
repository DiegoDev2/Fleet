
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// gnupg-pkcs11-scdFormula represents a formula in Go.
type gnupg-pkcs11-scdFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg gnupg-pkcs11-scdFormula) Print() {
    fmt.Printf("Name: gnupg-pkcs11-scd\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := gnupg-pkcs11-scdFormula{
        Description:  "Enable the use of PKCS#11 tokens with GnuPG",
        Homepage:     "https://gnupg-pkcs11.sourceforge.net/",
        URL:          "https://github.com/alonbl/gnupg-pkcs11-scd/releases/download/gnupg-pkcs11-scd-0.10.0/gnupg-pkcs11-scd-0.10.0.tar.bz2",
        Sha256:       "7610c2def9c9575bac8b8433ad4ec55f66d40ff82f724c805ce37e9a3ad8c9a7",
        Dependencies: []string{"autoconf", "automake", "libtool", "pkg-config", "libassuan@2", "libgcrypt", "libgpg-error", "openssl@3", "pkcs11-helper"},
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

    fmt.Printf("Installing formula: %s\n", "gnupg-pkcs11-scd")
    if err := pkg.Installgnupg-pkcs11-scd(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg gnupg-pkcs11-scdFormula) Installgnupg-pkcs11-scd() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "gnupg-pkcs11-scd-0.10.0.tar.bz2"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "gnupg-pkcs11-scd-0.10.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
