
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// bitcoinFormula represents a formula in Go.
type bitcoinFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg bitcoinFormula) Print() {
    fmt.Printf("Name: bitcoin\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := bitcoinFormula{
        Description:  "Decentralized, peer to peer payment network",
        Homepage:     "https://bitcoincore.org/",
        URL:          "https://github.com/fanquake/bitcoin/commit/9b03fb7603709395faaf0fac409465660bbd7d81.patch?full_index=1",
        Sha256:       "1d56308672024260e127fbb77f630b54a0509c145e397ff708956188c96bbfb3",
        Dependencies: []string{"autoconf", "automake", "boost", "libtool", "pkg-config", "libevent", "miniupnpc", "zeromq", "util-linux"},
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

    fmt.Printf("Installing formula: %s\n", "bitcoin")
    if err := pkg.Installbitcoin(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg bitcoinFormula) Installbitcoin() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "9b03fb7603709395faaf0fac409465660bbd7d81.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "9b03fb7603709395faaf0fac409465660bbd7d81"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
