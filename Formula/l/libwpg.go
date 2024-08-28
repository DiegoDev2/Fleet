
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// libwpgFormula represents a formula in Go.
type libwpgFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg libwpgFormula) Print() {
    fmt.Printf("Name: libwpg\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := libwpgFormula{
        Description:  "Library for reading and parsing Word Perfect Graphics format",
        Homepage:     "https://libwpg.sourceforge.net/",
        URL:          "https://dev-www.libreoffice.org/src/",
        Sha256:       "c079bc2c2fc9c98e967d13ffa15b47ab25efc59ba66f731ce3758d1265017368",
        Dependencies: []string{"pkg-config", "librevenge", "libwpd"},
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

    fmt.Printf("Installing formula: %s\n", "libwpg")
    if err := pkg.Installlibwpg(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg libwpgFormula) Installlibwpg() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := ""
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := ""
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}