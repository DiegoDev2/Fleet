
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// kyoto-tycoonFormula represents a formula in Go.
type kyoto-tycoonFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg kyoto-tycoonFormula) Print() {
    fmt.Printf("Name: kyoto-tycoon\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := kyoto-tycoonFormula{
        Description:  "Database server with interface to Kyoto Cabinet",
        Homepage:     "https://dbmx.net/kyototycoon/",
        URL:          "https://raw.githubusercontent.com/Homebrew/formula-patches/9925c07/kyoto-tycoon/ephemeral-ports.patch",
        Sha256:       "736603b28e9e7562837d0f376d89c549f74a76d31658bf7d84b57c5e66512672",
        Dependencies: []string{"lua", "pkg-config", "kyoto-cabinet"},
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

    fmt.Printf("Installing formula: %s\n", "kyoto-tycoon")
    if err := pkg.Installkyoto-tycoon(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg kyoto-tycoonFormula) Installkyoto-tycoon() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "ephemeral-ports.patch"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "ephemeral-ports"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
