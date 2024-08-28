
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// tarsnapperFormula represents a formula in Go.
type tarsnapperFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg tarsnapperFormula) Print() {
    fmt.Printf("Name: tarsnapper\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := tarsnapperFormula{
        Description:  "Tarsnap wrapper which expires backups using a gfs-scheme",
        Homepage:     "https://github.com/miracle2k/tarsnapper",
        URL:          "https://github.com/miracle2k/tarsnapper/commit/def72ae8499b38ab4726d826d7363490de6564fb.patch?full_index=1",
        Sha256:       "927ff17243b2e751afc7034b059365ca0373db74dcc8d917b8489058a66b2d1c",
        Dependencies: []string{"libyaml", "python@3.12", "tarsnap"},
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

    fmt.Printf("Installing formula: %s\n", "tarsnapper")
    if err := pkg.Installtarsnapper(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg tarsnapperFormula) Installtarsnapper() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "def72ae8499b38ab4726d826d7363490de6564fb.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "def72ae8499b38ab4726d826d7363490de6564fb"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
