
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// msc-generatorFormula represents a formula in Go.
type msc-generatorFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg msc-generatorFormula) Print() {
    fmt.Printf("Name: msc-generator\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := msc-generatorFormula{
        Description:  "Draws signalling charts from textual description",
        Homepage:     "https://gitlab.com/msc-generator/msc-generator",
        URL:          "https://gitlab.com/api/v4/projects/31167732/packages",
        Sha256:       "b679c4ec1606b7be9c65ba4d7e320b3a48eb7f5bac32374315448ee871e0ec1a",
        Dependencies: []string{"autoconf", "automake", "bison", "help2man", "pkg-config", "cairo", "fontconfig", "gcc", "glpk", "graphviz", "libpng", "sdl2", "tinyxml2", "gnu-sed", "make", "mesa"},
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

    fmt.Printf("Installing formula: %s\n", "msc-generator")
    if err := pkg.Installmsc-generator(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg msc-generatorFormula) Installmsc-generator() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "packages"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "packages"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
