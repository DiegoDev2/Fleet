
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// sh4d0wupFormula represents a formula in Go.
type sh4d0wupFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg sh4d0wupFormula) Print() {
    fmt.Printf("Name: sh4d0wup\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := sh4d0wupFormula{
        Description:  "Signing-key abuse and update exploitation framework",
        Homepage:     "https://github.com/kpcyrd/sh4d0wup",
        URL:          "https://raw.githubusercontent.com/Homebrew/formula-patches/31f5e08b1c7df4025d7042dafe756e5326151158/sh4d0wup/rust-1.80.patch",
        Sha256:       "24f3fc3919ead47c6e38c68a55d8fed0370cfddd92738519de4bd41e4da71e93",
        Dependencies: []string{"llvm", "pkg-config", "rust", "pgpdump", "gmp", "nettle", "openssl@3", "pcsc-lite", "xz", "zstd"},
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

    fmt.Printf("Installing formula: %s\n", "sh4d0wup")
    if err := pkg.Installsh4d0wup(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg sh4d0wupFormula) Installsh4d0wup() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "rust-1.80.patch"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "rust-1.80"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
