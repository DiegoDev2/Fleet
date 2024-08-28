
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// core-lightningFormula represents a formula in Go.
type core-lightningFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg core-lightningFormula) Print() {
    fmt.Printf("Name: core-lightning\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := core-lightningFormula{
        Description:  "Lightning Network implementation focusing on spec compliance and performance",
        Homepage:     "https://github.com/ElementsProject/lightning",
        URL:          "https://github.com/ElementsProject/lightning/releases/download/v24.05/clightning-v24.05.zip",
        Sha256:       "88d9a8df043561a371db4bc4b54a398a242d4c5bbd9d808d07b549e6b3ca7d63",
        Dependencies: []string{"autoconf", "automake", "gettext", "gnu-sed", "jq", "libtool", "lowdown", "pkg-config", "poetry", "protobuf", "bitcoin", "gmp", "libsodium"},
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

    fmt.Printf("Installing formula: %s\n", "core-lightning")
    if err := pkg.Installcore-lightning(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg core-lightningFormula) Installcore-lightning() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "clightning-v24.05.zip"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "clightning-v24.05"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
