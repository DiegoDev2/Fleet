
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// tarsnapFormula represents a formula in Go.
type tarsnapFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg tarsnapFormula) Print() {
    fmt.Printf("Name: tarsnap\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := tarsnapFormula{
        Description:  "Online backups for the truly paranoid",
        Homepage:     "https://www.tarsnap.com/",
        URL:          "https://github.com/Tarsnap/tarsnap/commit/4af6d8350ab53d0f1f3104ce3d9072c2d5f9ef7a.patch?full_index=1",
        Sha256:       "4136b5643e25f7d5e454c014b3c13d7ad015a02e796c5c91b3e4eeca28c1556e",
        Dependencies: []string{"autoconf", "automake", "openssl@3", "e2fsprogs"},
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

    fmt.Printf("Installing formula: %s\n", "tarsnap")
    if err := pkg.Installtarsnap(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg tarsnapFormula) Installtarsnap() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "4af6d8350ab53d0f1f3104ce3d9072c2d5f9ef7a.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "4af6d8350ab53d0f1f3104ce3d9072c2d5f9ef7a"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
