
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// udp2raw-multiplatformFormula represents a formula in Go.
type udp2raw-multiplatformFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg udp2raw-multiplatformFormula) Print() {
    fmt.Printf("Name: udp2raw-multiplatform\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := udp2raw-multiplatformFormula{
        Description:  "Multi-platform(cross-platform) version of udp2raw-tunnel client",
        Homepage:     "https://github.com/wangyu-/udp2raw-multiplatform",
        URL:          "https://github.com/wangyu-/udp2raw-multiplatform/archive/refs/tags/20230206.0.tar.gz",
        Sha256:       "23603c4582835dcb9428c1c3e553802937568f9e7ea1bd283bf329541562ffdc",
        Dependencies: []string{"libnet"},
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

    fmt.Printf("Installing formula: %s\n", "udp2raw-multiplatform")
    if err := pkg.Installudp2raw-multiplatform(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg udp2raw-multiplatformFormula) Installudp2raw-multiplatform() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "20230206.0.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "20230206.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
