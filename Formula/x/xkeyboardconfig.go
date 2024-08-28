
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// xkeyboardconfigFormula represents a formula in Go.
type xkeyboardconfigFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg xkeyboardconfigFormula) Print() {
    fmt.Printf("Name: xkeyboardconfig\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := xkeyboardconfigFormula{
        Description:  "Keyboard configuration database for the X Window System",
        Homepage:     "https://www.freedesktop.org/wiki/Software/XKeyboardConfig/",
        URL:          "https://xorg.freedesktop.org/archive/individual/data/xkeyboard-config/xkeyboard-config-2.42.tar.xz",
        Sha256:       "b11a43e9842d1c854a9f01295b888c47d0c7d7e9c2b2ff1b88f6f84c97aa9f7b",
        Dependencies: []string{"gettext", "meson", "ninja", "pkg-config", "python@3.12"},
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

    fmt.Printf("Installing formula: %s\n", "xkeyboardconfig")
    if err := pkg.Installxkeyboardconfig(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg xkeyboardconfigFormula) Installxkeyboardconfig() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "xkeyboard-config-2.42.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "xkeyboard-config-2.42.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
