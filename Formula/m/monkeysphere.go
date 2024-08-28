
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// monkeysphereFormula represents a formula in Go.
type monkeysphereFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg monkeysphereFormula) Print() {
    fmt.Printf("Name: monkeysphere\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := monkeysphereFormula{
        Description:  "Use the OpenPGP web of trust to verify ssh connections",
        Homepage:     "https://tracker.debian.org/pkg/monkeysphere",
        URL:          "https://cpan.metacpan.org/authors/id/T/TO/TODDR/Crypt-OpenSSL-RSA-0.33.tar.gz",
        Sha256:       "bdbe630f6d6f540325746ad99977272ac8664ff81bd19f0adaba6d6f45efd864",
        Dependencies: []string{"gnu-sed", "pkg-config", "gnupg", "libassuan", "libgcrypt", "libgpg-error", "openssl@3"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installmonkeysphere(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg monkeysphereFormula) Installmonkeysphere() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "Crypt-OpenSSL-RSA-0.33.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "Crypt-OpenSSL-RSA-0.33.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
