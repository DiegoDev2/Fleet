
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// synergy-coreFormula represents a formula in Go.
type synergy-coreFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg synergy-coreFormula) Print() {
    fmt.Printf("Name: synergy-core\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := synergy-coreFormula{
        Description:  "Synergy, the keyboard and mouse sharing tool",
        Homepage:     "https://symless.com/synergy",
        URL:          "https://github.com/symless/synergy-core/archive/refs/tags/1.14.6.19-stable.tar.gz",
        Sha256:       "0b80be1af6d257aea32a7c8533fd7e7254832f6c05a2e706633e39a81e096a6b",
        Dependencies: []string{"cmake", "openssl@3", "pugixml", "qt@5", "pkg-config", "gdk-pixbuf", "glib", "libnotify", "libx11", "libxext", "libxi", "libxinerama", "libxkbfile", "libxrandr", "libxtst"},
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

    fmt.Printf("Installing formula: %s\n", "synergy-core")
    if err := pkg.Installsynergy-core(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg synergy-coreFormula) Installsynergy-core() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "1.14.6.19-stable.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "1.14.6.19-stable.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
