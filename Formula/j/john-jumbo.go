
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// john-jumboFormula represents a formula in Go.
type john-jumboFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg john-jumboFormula) Print() {
    fmt.Printf("Name: john-jumbo\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := john-jumboFormula{
        Description:  "Enhanced version of john, a UNIX password cracker",
        Homepage:     "https://www.openwall.com/john/",
        URL:          "https://github.com/openwall/john/commit/8152ac071bce1ebc98fac6bed962e90e9b92d8cf.patch?full_index=1",
        Sha256:       "efb4e3597c47930d63f51efbf18c409f436ea6bd0012a4290b05135a54d7edd4",
        Dependencies: []string{"pkg-config", "gmp", "openssl@3"},
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

    fmt.Printf("Installing formula: %s\n", "john-jumbo")
    if err := pkg.Installjohn-jumbo(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg john-jumboFormula) Installjohn-jumbo() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "8152ac071bce1ebc98fac6bed962e90e9b92d8cf.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "8152ac071bce1ebc98fac6bed962e90e9b92d8cf"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
