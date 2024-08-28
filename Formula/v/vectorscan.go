
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// vectorscanFormula represents a formula in Go.
type vectorscanFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg vectorscanFormula) Print() {
    fmt.Printf("Name: vectorscan\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := vectorscanFormula{
        Description:  "High-performance regular expression matching library",
        Homepage:     "https://github.com/VectorCamp/vectorscan",
        URL:          "https://github.com/VectorCamp/vectorscan/commit/d9ebb20010b3f90a7a5c7bf4a5edff2eb58f2a4f.patch?full_index=1",
        Sha256:       "e61de5f0321e9020871912883dadcdc1f49cd423dab37de67b6c1e8d07115162",
        Dependencies: []string{"boost", "cmake", "pcre", "pkg-config", "ragel"},
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

    fmt.Printf("Installing formula: %s\n", "vectorscan")
    if err := pkg.Installvectorscan(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg vectorscanFormula) Installvectorscan() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "d9ebb20010b3f90a7a5c7bf4a5edff2eb58f2a4f.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "d9ebb20010b3f90a7a5c7bf4a5edff2eb58f2a4f"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
