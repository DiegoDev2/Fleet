
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// uwsgiFormula represents a formula in Go.
type uwsgiFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg uwsgiFormula) Print() {
    fmt.Printf("Name: uwsgi\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := uwsgiFormula{
        Description:  "Full stack for building hosting services",
        Homepage:     "https://uwsgi-docs.readthedocs.io/en/latest/",
        URL:          "https://files.pythonhosted.org/packages/3a/7a/4c910bdc9d32640ba89f8d1dc256872c2b5e64830759f7dc346815f5b3b1/uwsgi-2.0.26.tar.gz",
        Sha256:       "052a7b88f365576d5505f51a53c7ac1085c6c671696d4107875722bbecd4e73b",
        Dependencies: []string{"pkg-config", "openssl@3", "pcre", "python@3.12", "sqlite", "yajl", "linux-pam"},
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

    fmt.Printf("Installing formula: %s\n", "uwsgi")
    if err := pkg.Installuwsgi(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg uwsgiFormula) Installuwsgi() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "uwsgi-2.0.26.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "uwsgi-2.0.26.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
