
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// dhall-lsp-serverFormula represents a formula in Go.
type dhall-lsp-serverFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg dhall-lsp-serverFormula) Print() {
    fmt.Printf("Name: dhall-lsp-server\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := dhall-lsp-serverFormula{
        Description:  "Language Server Protocol (LSP) server for Dhall",
        Homepage:     "https://github.com/dhall-lang/dhall-haskell/tree/master/dhall-lsp-server",
        URL:          "https://github.com/dhall-lang/dhall-haskell/commit/5e817a9c6bccf72123a3c67961af149b32d75c10.patch?full_index=1",
        Sha256:       "994efb7ca3ec4bb7af8d557d1a6393d699c3a0fd6b8ee8902478e2b1104f6aa1",
        Dependencies: []string{"cabal-install", "ghc@9.4"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installdhall-lsp-server(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg dhall-lsp-serverFormula) Installdhall-lsp-server() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "5e817a9c6bccf72123a3c67961af149b32d75c10.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "5e817a9c6bccf72123a3c67961af149b32d75c10"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
