
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// eigenpyFormula represents a formula in Go.
type eigenpyFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg eigenpyFormula) Print() {
    fmt.Printf("Name: eigenpy\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := eigenpyFormula{
        Description:  "Python bindings of Eigen library with Numpy support",
        Homepage:     "https://github.com/stack-of-tasks/eigenpy",
        URL:          "https://github.com/stack-of-tasks/eigenpy/commit/b36cded3d855557bd69f63b215b9c45ecb8b0255.patch?full_index=1",
        Sha256:       "ab5f5cfe66d23a1128de4340f49de28487b1ae7082bc48e71b68521fda540e5d",
        Dependencies: []string{"boost", "cmake", "doxygen", "pkg-config", "python-setuptools", "boost-python3", "eigen", "numpy", "python@3.12", "scipy"},
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

    fmt.Printf("Installing formula: %s\n", "eigenpy")
    if err := pkg.Installeigenpy(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg eigenpyFormula) Installeigenpy() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "b36cded3d855557bd69f63b215b9c45ecb8b0255.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "b36cded3d855557bd69f63b215b9c45ecb8b0255"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
