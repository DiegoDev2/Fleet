
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// shivavgFormula represents a formula in Go.
type shivavgFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg shivavgFormula) Print() {
    fmt.Printf("Name: shivavg\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := shivavgFormula{
        Description:  "OpenGL based ANSI C implementation of the OpenVG standard",
        Homepage:     "https://sourceforge.net/projects/shivavg/",
        URL:          "https://github.com/Ecognize/ShivaVG/commit/fe3bb03d7b03591b26ab7c399f51edcd130f0f4e.patch?full_index=1",
        Sha256:       "f4eb595afb053eb0a093dcf50748b54e01eff729f4589f82deb8f6f2ce8f571b",
        Dependencies: []string{"autoconf", "automake", "libtool", "freeglut", "mesa", "mesa-glu"},
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

    fmt.Printf("Installing formula: %s\n", "shivavg")
    if err := pkg.Installshivavg(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg shivavgFormula) Installshivavg() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "fe3bb03d7b03591b26ab7c399f51edcd130f0f4e.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "fe3bb03d7b03591b26ab7c399f51edcd130f0f4e"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
