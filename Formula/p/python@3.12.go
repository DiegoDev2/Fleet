
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// python@3.12Formula represents a formula in Go.
type python@3.12Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg python@3.12Formula) Print() {
    fmt.Printf("Name: python@3.12\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := python@3.12Formula{
        Description:  "Interpreted, interactive, object-oriented programming language",
        Homepage:     "https://www.python.org/",
        URL:          "https://raw.githubusercontent.com/Homebrew/formula-patches/6d2fba8de3159182025237d373a6f4f78b8bd203/python/3.11-sysconfig.diff",
        Sha256:       "8bfe417c815da4ca2c0a2457ce7ef81bc9dae310e20e4fb36235901ea4be1658",
        Dependencies: []string{"pkg-config", "mpdecimal", "openssl@3", "sqlite", "xz", "berkeley-db@5", "libnsl"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installpython@3.12(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg python@3.12Formula) Installpython@3.12() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "3.11-sysconfig.diff"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "3.11-sysconfig"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
