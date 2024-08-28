
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// pyenv-virtualenvwrapperFormula represents a formula in Go.
type pyenv-virtualenvwrapperFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg pyenv-virtualenvwrapperFormula) Print() {
    fmt.Printf("Name: pyenv-virtualenvwrapper\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := pyenv-virtualenvwrapperFormula{
        Description:  "Alternative to pyenv for managing virtualenvs",
        Homepage:     "https://github.com/pyenv/pyenv-virtualenvwrapper",
        URL:          "https://github.com/pyenv/pyenv-virtualenvwrapper/archive/refs/tags/v20140609.tar.gz",
        Sha256:       "0aeb3455529d63f4cd1ca55acb525e4f38e1fc7b8dca986302f475bc8596a650",
        Dependencies: []string{"pyenv"},
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

    fmt.Printf("Installing formula: %s\n", "pyenv-virtualenvwrapper")
    if err := pkg.Installpyenv-virtualenvwrapper(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg pyenv-virtualenvwrapperFormula) Installpyenv-virtualenvwrapper() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "v20140609.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "v20140609.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
