
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// python-trove-classifiersFormula represents a formula in Go.
type python-trove-classifiersFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg python-trove-classifiersFormula) Print() {
    fmt.Printf("Name: python-trove-classifiers\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := python-trove-classifiersFormula{
        Description:  "Canonical source for classifiers on PyPI",
        Homepage:     "https://github.com/pypa/trove-classifiers",
        URL:          "https://files.pythonhosted.org/packages/c5/e9/1deb1b8113917aa735b08ef21821f3533bda2eb1fa1a16e07256dd05918f/trove-classifiers-2024.3.25.tar.gz",
        Sha256:       "60b794b825880dce65e9024a83930814c42be20b2cac8c4f162e92d83fe86eed",
        Dependencies: []string{"python@3.11", "python@3.12"},
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

    fmt.Printf("Installing formula: %s\n", "python-trove-classifiers")
    if err := pkg.Installpython-trove-classifiers(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg python-trove-classifiersFormula) Installpython-trove-classifiers() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "trove-classifiers-2024.3.25.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "trove-classifiers-2024.3.25.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
