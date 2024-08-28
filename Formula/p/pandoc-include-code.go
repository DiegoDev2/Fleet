
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// pandoc-include-codeFormula represents a formula in Go.
type pandoc-include-codeFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg pandoc-include-codeFormula) Print() {
    fmt.Printf("Name: pandoc-include-code\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := pandoc-include-codeFormula{
        Description:  "Pandoc filter for including code from source files",
        Homepage:     "https://github.com/owickstrom/pandoc-include-code",
        URL:          "https://hackage.haskell.org/package/pandoc-include-code-1.5.0.0/pandoc-include-code-1.5.0.0.tar.gz",
        Sha256:       "ab22b3df53ac6762dcc8b564ebbefdc96a814ba31f93fc4dd701bbb28bacb958",
        Dependencies: []string{"cabal-install", "ghc@8.10", "pandoc"},
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

    fmt.Printf("Installing formula: %s\n", "pandoc-include-code")
    if err := pkg.Installpandoc-include-code(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg pandoc-include-codeFormula) Installpandoc-include-code() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "pandoc-include-code-1.5.0.0.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "pandoc-include-code-1.5.0.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
