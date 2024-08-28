
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// cabal-installFormula represents a formula in Go.
type cabal-installFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg cabal-installFormula) Print() {
    fmt.Printf("Name: cabal-install\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := cabal-installFormula{
        Description:  "Command-line interface for Cabal and Hackage",
        Homepage:     "https://www.haskell.org/cabal/",
        URL:          "https://downloads.haskell.org/~cabal/cabal-install-3.10.3.0/cabal-install-3.10.3.0-x86_64-linux-ubuntu20_04.tar.xz",
        Sha256:       "b7ccb975bacf8b6a7d6b5dde8a7712787473a149c3dc0ebb2de7fbd00f964844",
        Dependencies: []string{"ghc"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installcabal-install(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg cabal-installFormula) Installcabal-install() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "cabal-install-3.10.3.0-x86_64-linux-ubuntu20_04.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "cabal-install-3.10.3.0-x86_64-linux-ubuntu20_04.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
