
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// ansible-lintFormula represents a formula in Go.
type ansible-lintFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg ansible-lintFormula) Print() {
    fmt.Printf("Name: ansible-lint\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := ansible-lintFormula{
        Description:  "Checks ansible playbooks for practices and behaviour",
        Homepage:     "https://ansible-lint.readthedocs.io/",
        URL:          "https://files.pythonhosted.org/packages/0e/af/9f2de5bd32549a1b705af7a7c054af3878816a1267cb389c03cc4f342a51/zipp-3.20.0.tar.gz",
        Sha256:       "0145e43d89664cfe1a2e533adc75adafed82fe2da404b4bbb6b026c0157bdb31",
        Dependencies: []string{"pkg-config", "rust", "ansible", "cryptography", "libyaml", "python@3.12"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installansible-lint(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg ansible-lintFormula) Installansible-lint() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "zipp-3.20.0.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "zipp-3.20.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}