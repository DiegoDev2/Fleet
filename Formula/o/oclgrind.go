
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// oclgrindFormula represents a formula in Go.
type oclgrindFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg oclgrindFormula) Print() {
    fmt.Printf("Name: oclgrind\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := oclgrindFormula{
        Description:  "OpenCL device simulator and debugger",
        Homepage:     "https://github.com/jrprice/Oclgrind",
        URL:          "https://github.com/jrprice/Oclgrind/commit/6c76e7bec0aa7fa451515a5cfcb35ab2384ba6e0.patch?full_index=1",
        Sha256:       "8c1b8ec75d8d8c8d02246124b40452ec9ef1243d3e3c497fe4ffa8571cd98ade",
        Dependencies: []string{"cmake", "llvm@14", "readline", "opencl-headers"},
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

    fmt.Printf("Installing formula: %s\n", "oclgrind")
    if err := pkg.Installoclgrind(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg oclgrindFormula) Installoclgrind() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "6c76e7bec0aa7fa451515a5cfcb35ab2384ba6e0.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "6c76e7bec0aa7fa451515a5cfcb35ab2384ba6e0"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
