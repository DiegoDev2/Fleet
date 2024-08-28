
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// pianobarFormula represents a formula in Go.
type pianobarFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg pianobarFormula) Print() {
    fmt.Printf("Name: pianobar\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := pianobarFormula{
        Description:  "Command-line player for https://pandora.com",
        Homepage:     "https://6xq.net/pianobar/",
        URL:          "https://github.com/PromyLOPh/pianobar/commit/8bf4c1bbaa6a533f34d74f83d72eecc0beb61d4f.patch?full_index=1",
        Sha256:       "9ede8281d2c8a056b646f87582b84e20d3fc290df41c6e93336f90729794b58b",
        Dependencies: []string{"pkg-config", "ffmpeg", "json-c", "libao", "libgcrypt"},
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

    fmt.Printf("Installing formula: %s\n", "pianobar")
    if err := pkg.Installpianobar(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg pianobarFormula) Installpianobar() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "8bf4c1bbaa6a533f34d74f83d72eecc0beb61d4f.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "8bf4c1bbaa6a533f34d74f83d72eecc0beb61d4f"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
