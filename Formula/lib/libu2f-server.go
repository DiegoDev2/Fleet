
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// libu2f-serverFormula represents a formula in Go.
type libu2f-serverFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg libu2f-serverFormula) Print() {
    fmt.Printf("Name: libu2f-server\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := libu2f-serverFormula{
        Description:  "Server-side of the Universal 2nd Factor (U2F) protocol",
        Homepage:     "https://developers.yubico.com/libu2f-server/",
        URL:          "https://github.com/Yubico/libu2f-server/commit/f7c4983b31909299c47bf9b2627c84b6bfe225de.patch?full_index=1",
        Sha256:       "012d1d759604ea80f6075b74dc9c7d8a864e4e5889fb82a222db93a6bd72cd1b",
        Dependencies: []string{"check", "gengetopt", "help2man", "pkg-config", "json-c", "openssl@3"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installlibu2f-server(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg libu2f-serverFormula) Installlibu2f-server() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "f7c4983b31909299c47bf9b2627c84b6bfe225de.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "f7c4983b31909299c47bf9b2627c84b6bfe225de"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
