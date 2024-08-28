
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// networkitFormula represents a formula in Go.
type networkitFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg networkitFormula) Print() {
    fmt.Printf("Name: networkit\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := networkitFormula{
        Description:  "Performance toolkit for large-scale network analysis",
        Homepage:     "https://networkit.github.io",
        URL:          "https://github.com/networkit/networkit/commit/165503580caac864c7a31558b4c5fee27bcb007e.patch?full_index=1",
        Sha256:       "67bd2d1fe3ebccb42ccdd1f7cf5aeea40967caa4e9bc96cc69737dc14ffa9654",
        Dependencies: []string{"cmake", "cython", "ninja", "python-setuptools", "tlx", "libnetworkit", "numpy", "python@3.12", "scipy", "libomp"},
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

    fmt.Printf("Installing formula: %s\n", "networkit")
    if err := pkg.Installnetworkit(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg networkitFormula) Installnetworkit() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "165503580caac864c7a31558b4c5fee27bcb007e.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "165503580caac864c7a31558b4c5fee27bcb007e"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
