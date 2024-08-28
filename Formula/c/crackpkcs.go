
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// crackpkcsFormula represents a formula in Go.
type crackpkcsFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg crackpkcsFormula) Print() {
    fmt.Printf("Name: crackpkcs\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := crackpkcsFormula{
        Description:  "Multithreaded program to crack PKCS#12 files",
        Homepage:     "https://crackpkcs12.sourceforge.net/",
        URL:          "https://github.com/crackpkcs12/crackpkcs12/raw/9f7375fdc7358451add8b31aaf928ecd025d63d9/misc/utils/certs/usr0052-exportado_desde_firefox.p12",
        Sha256:       "8789861fbaf1a0fc6299756297fe646692a7b43e06c2be89a382b3dceb93f454",
        Dependencies: []string{"pkg-config", "openssl@1.1"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installcrackpkcs(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg crackpkcsFormula) Installcrackpkcs() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "usr0052-exportado_desde_firefox.p12"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "usr0052-exportado_desde_firefox"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
