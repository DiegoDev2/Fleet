
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// juliaFormula represents a formula in Go.
type juliaFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg juliaFormula) Print() {
    fmt.Printf("Name: julia\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := juliaFormula{
        Description:  "Fast, Dynamic Programming Language",
        Homepage:     "https://julialang.org/",
        URL:          "https://raw.githubusercontent.com/Homebrew/formula-patches/202ccbabd44bd5ab02fbdee2f51f87bb88d74417/julia/libgcc_s-1.8.5.diff",
        Sha256:       "1eea77d8024ad8bc9c733a0e0770661bc08228d335b20c4696350ed5dfdab29a",
        Dependencies: []string{"cmake", "gcc", "suite-sparse", "ca-certificates", "curl", "gmp", "libgit2", "libnghttp2", "libssh2", "mbedtls@2", "mpfr", "openblas", "openlibm", "p7zip", "pcre2", "utf8proc", "patchelf"},
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

    fmt.Printf("Installing formula: %s\n", "julia")
    if err := pkg.Installjulia(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg juliaFormula) Installjulia() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "libgcc_s-1.8.5.diff"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "libgcc_s-1.8.5"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
