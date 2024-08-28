
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// chocolate-doomFormula represents a formula in Go.
type chocolate-doomFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg chocolate-doomFormula) Print() {
    fmt.Printf("Name: chocolate-doom\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := chocolate-doomFormula{
        Description:  "Accurate source port of Doom",
        Homepage:     "https://www.chocolate-doom.org/",
        URL:          "https://github.com/chocolate-doom/chocolate-doom/archive/refs/tags/chocolate-doom-3.1.0.tar.gz",
        Sha256:       "d314bee173f047d52b928d7865e8d8fd10ef6635b34cc89a50259ab9766e2efc",
        Dependencies: []string{"autoconf", "automake", "pkg-config", "fluid-synth", "libpng", "libsamplerate", "sdl2", "sdl2_mixer", "sdl2_net"},
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

    fmt.Printf("Installing formula: %s\n", "chocolate-doom")
    if err := pkg.Installchocolate-doom(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg chocolate-doomFormula) Installchocolate-doom() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "chocolate-doom-3.1.0.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "chocolate-doom-3.1.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
