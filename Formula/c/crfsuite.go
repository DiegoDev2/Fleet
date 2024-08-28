
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// crfsuiteFormula represents a formula in Go.
type crfsuiteFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg crfsuiteFormula) Print() {
    fmt.Printf("Name: crfsuite\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := crfsuiteFormula{
        Description:  "Fast implementation of conditional random fields",
        Homepage:     "https://www.chokkan.org/software/crfsuite/",
        URL:          "https://github.com/chokkan/crfsuite/commit/a6a4a38ccc4738deb0e90fc9ff2c11868922aa11.patch?full_index=1",
        Sha256:       "8c572cb9d737e058b0a86c6eab96d1ffa8951016b50eee505491c2dae7c7c74d",
        Dependencies: []string{"autoconf", "automake", "libtool", "liblbfgs"},
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

    fmt.Printf("Installing formula: %s\n", "crfsuite")
    if err := pkg.Installcrfsuite(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg crfsuiteFormula) Installcrfsuite() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "a6a4a38ccc4738deb0e90fc9ff2c11868922aa11.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "a6a4a38ccc4738deb0e90fc9ff2c11868922aa11"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
