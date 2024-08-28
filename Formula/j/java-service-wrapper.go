
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// java-service-wrapperFormula represents a formula in Go.
type java-service-wrapperFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg java-service-wrapperFormula) Print() {
    fmt.Printf("Name: java-service-wrapper\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := java-service-wrapperFormula{
        Description:  "Simplify the deployment, launch and monitoring of Java applications",
        Homepage:     "https://wrapper.tanukisoftware.com/",
        URL:          "https://downloads.sourceforge.net/project/wrapper/wrapper_src/Wrapper_3.5.59_20240723/wrapper_3.5.59_src.tar.gz",
        Sha256:       "187b9c59b2589c36223445d0b0fd9d4985dccaeca299ed046ffca2c1b2565f6d",
        Dependencies: []string{"ant", "openjdk@11", "cunit"},
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

    fmt.Printf("Installing formula: %s\n", "java-service-wrapper")
    if err := pkg.Installjava-service-wrapper(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg java-service-wrapperFormula) Installjava-service-wrapper() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "wrapper_3.5.59_src.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "wrapper_3.5.59_src.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
