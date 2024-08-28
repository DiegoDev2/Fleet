
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// proxsuiteFormula represents a formula in Go.
type proxsuiteFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg proxsuiteFormula) Print() {
    fmt.Printf("Name: proxsuite\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := proxsuiteFormula{
        Description:  "Advanced Proximal Optimization Toolbox",
        Homepage:     "https://github.com/Simple-Robotics/proxsuite",
        URL:          "https://github.com/Simple-Robotics/proxsuite/releases/download/v0.6.7/proxsuite-0.6.7.tar.gz",
        Sha256:       "adf28af671df40fef0c8f9ee27812ae8778c669f6b5fc580711b18963ef00000",
        Dependencies: []string{"cmake", "doxygen", "pkg-config", "python-setuptools", "eigen", "numpy", "python@3.12", "scipy", "simde"},
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

    fmt.Printf("Installing formula: %s\n", "proxsuite")
    if err := pkg.Installproxsuite(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg proxsuiteFormula) Installproxsuite() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "proxsuite-0.6.7.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "proxsuite-0.6.7.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
