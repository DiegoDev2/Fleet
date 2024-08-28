
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// openmodelicaFormula represents a formula in Go.
type openmodelicaFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg openmodelicaFormula) Print() {
    fmt.Printf("Name: openmodelica\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := openmodelicaFormula{
        Description:  "Open-source modeling and simulation tool",
        Homepage:     "https://openmodelica.org/",
        URL:          "https://github.com/OpenModelica/OpenModelica.git",
        Sha256:       "727699c2d6e510c6cb271594644ef0eec12f28e7b4acb047c778ffa1f68efbf0",
        Dependencies: []string{"autoconf", "automake", "cmake", "gcc", "gnu-sed", "libtool", "openjdk", "pkg-config", "boost", "gettext", "hdf5", "hwloc", "lp_solve", "omniorb", "openblas", "qt@5", "readline", "sundials"},
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

    fmt.Printf("Installing formula: %s\n", "openmodelica")
    if err := pkg.Installopenmodelica(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg openmodelicaFormula) Installopenmodelica() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "OpenModelica.git"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "OpenModelica"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
