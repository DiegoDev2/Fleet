
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// argyll-cmsFormula represents a formula in Go.
type argyll-cmsFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg argyll-cmsFormula) Print() {
    fmt.Printf("Name: argyll-cms\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := argyll-cmsFormula{
        Description:  "ICC compatible color management system",
        Homepage:     "https://www.argyllcms.com/",
        URL:          "https://raw.githubusercontent.com/Homebrew/formula-patches/cf70f015e7398796660da57212ff0ab90c609acf/jam/2.6.1.patch",
        Sha256:       "1850cf53c4db0e05978d52b90763b519c00fa4f2fbd6fc2753200e49943821ec",
        Dependencies: []string{"jpeg-turbo", "libpng", "libtiff", "openssl@3", "libx11", "libxext", "libxinerama", "libxrandr", "libxscrnsaver", "libxxf86vm", "xorgproto"},
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

    fmt.Printf("Installing formula: %s\n", "argyll-cms")
    if err := pkg.Installargyll-cms(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg argyll-cmsFormula) Installargyll-cms() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "2.6.1.patch"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "2.6.1"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
