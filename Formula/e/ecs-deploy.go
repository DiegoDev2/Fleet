
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// ecs-deployFormula represents a formula in Go.
type ecs-deployFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg ecs-deployFormula) Print() {
    fmt.Printf("Name: ecs-deploy\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := ecs-deployFormula{
        Description:  "CLI tool to simplify Amazon ECS deployments, rollbacks & scaling",
        Homepage:     "https://github.com/fabfuel/ecs-deploy",
        URL:          "https://files.pythonhosted.org/packages/c8/93/65e479b023bbc46dab3e092bda6b0005424ea3217d711964ccdede3f9b1b/urllib3-1.26.19.tar.gz",
        Sha256:       "3e3d753a8618b86d7de333b4223005f68720bcd6a7d2bcb9fbd2229ec7c1e429",
        Dependencies: []string{"certifi", "python@3.12"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installecs-deploy(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg ecs-deployFormula) Installecs-deploy() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "urllib3-1.26.19.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "urllib3-1.26.19.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
