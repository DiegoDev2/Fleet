
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// docbook-xslFormula represents a formula in Go.
type docbook-xslFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg docbook-xslFormula) Print() {
    fmt.Printf("Name: docbook-xsl\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := docbook-xslFormula{
        Description:  "XML vocabulary to create presentation-neutral documents",
        Homepage:     "https://github.com/docbook/xslt10-stylesheets",
        URL:          "https://www.linuxfromscratch.org/patches/blfs/9.1/docbook-xsl-nons-1.79.2-stack_fix-1.patch",
        Sha256:       "a92c39715c54949ba9369add1809527b8f155b7e2a2b2e30cb4b39ee715f2e30",
        Dependencies: []string{"docbook"},
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

    fmt.Printf("Installing formula: %s\n", "docbook-xsl")
    if err := pkg.Installdocbook-xsl(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg docbook-xslFormula) Installdocbook-xsl() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "docbook-xsl-nons-1.79.2-stack_fix-1.patch"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "docbook-xsl-nons-1.79.2-stack_fix-1"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
