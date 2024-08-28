
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// aptosFormula represents a formula in Go.
type aptosFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg aptosFormula) Print() {
    fmt.Printf("Name: aptos\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := aptosFormula{
        Description:  "Layer 1 blockchain built to support fair access to decentralized assets for all",
        Homepage:     "https://aptosfoundation.org/",
        URL:          "https://github.com/aptos-labs/aptos-core/commit/15ec18e4a2533fc9f3e8d23f5629939a07490c23.patch?full_index=1",
        Sha256:       "92384b60959e5c9400542e3215445dc3ae634693e41de9ae8b1907ad7557c5b7",
        Dependencies: []string{"cmake", "rust", "rustfmt", "pkg-config", "zip", "openssl@3", "systemd"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installaptos(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg aptosFormula) Installaptos() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "15ec18e4a2533fc9f3e8d23f5629939a07490c23.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "15ec18e4a2533fc9f3e8d23f5629939a07490c23"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
