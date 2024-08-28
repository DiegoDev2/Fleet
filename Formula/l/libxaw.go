
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// libxawFormula represents a formula in Go.
type libxawFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg libxawFormula) Print() {
    fmt.Printf("Name: libxaw\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := libxawFormula{
        Description:  "X.Org: X Athena Widget Set",
        Homepage:     "https://www.x.org/",
        URL:          "https://gitlab.freedesktop.org/xorg/lib/libxaw/-/commit/cce2abf00fa2c9a695f1d0e5c931c70c1ba579cf.diff",
        Sha256:       "64bfebc3fcb788582abbf2589514e64b3fa62457089c77e644177f1c0a80c10f",
        Dependencies: []string{"pkg-config", "libx11", "libxext", "libxmu", "libxpm", "libxt"},
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

    fmt.Printf("Installing formula: %s\n", "libxaw")
    if err := pkg.Installlibxaw(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg libxawFormula) Installlibxaw() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "cce2abf00fa2c9a695f1d0e5c931c70c1ba579cf.diff"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "cce2abf00fa2c9a695f1d0e5c931c70c1ba579cf"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
