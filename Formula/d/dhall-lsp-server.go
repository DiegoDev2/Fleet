
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

    fmt.Printf("Installing formula: %s\n", "dhall-lsp-server")
    if err := pkg.Installdhall-lsp-server(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
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
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
