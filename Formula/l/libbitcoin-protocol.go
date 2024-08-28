
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// libbitcoin-protocolFormula represents a formula in Go.
type libbitcoin-protocolFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg libbitcoin-protocolFormula) Print() {
    fmt.Printf("Name: libbitcoin-protocol\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := libbitcoin-protocolFormula{
        Description:  "Bitcoin Blockchain Query Protocol",
        Homepage:     "https://github.com/libbitcoin/libbitcoin-protocol",
        URL:          "https://github.com/libbitcoin/libbitcoin-protocol/archive/refs/tags/v3.8.0.tar.gz",
        Sha256:       "746249bddfec92921117988b74797a1bf4797bf954693fc36c04af6b6946bab1",
        Dependencies: []string{"autoconf", "automake", "libtool", "pkg-config", "boost@1.76", "libbitcoin-system", "libsodium", "zeromq"},
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

    fmt.Printf("Installing formula: %s\n", "libbitcoin-protocol")
    if err := pkg.Installlibbitcoin-protocol(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg libbitcoin-protocolFormula) Installlibbitcoin-protocol() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "v3.8.0.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "v3.8.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
