
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// cmu-sphinxbaseFormula represents a formula in Go.
type cmu-sphinxbaseFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg cmu-sphinxbaseFormula) Print() {
    fmt.Printf("Name: cmu-sphinxbase\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := cmu-sphinxbaseFormula{
        Description:  "Lightweight speech recognition engine for mobile devices",
        Homepage:     "https://cmusphinx.sourceforge.net/",
        URL:          "https://github.com/cmusphinx/sphinxbase.git",
        Sha256:       "078ae69b46fd06a8fcd68df8457e8c96e078c0276a027ae640ac18aad99d6669",
        Dependencies: []string{"autoconf", "automake", "libtool", "swig", "pkg-config", "libsamplerate", "libsndfile"},
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

    fmt.Printf("Installing formula: %s\n", "cmu-sphinxbase")
    if err := pkg.Installcmu-sphinxbase(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg cmu-sphinxbaseFormula) Installcmu-sphinxbase() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "sphinxbase.git"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "sphinxbase"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
