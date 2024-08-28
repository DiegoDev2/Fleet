
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// gtksourceviewmmFormula represents a formula in Go.
type gtksourceviewmmFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg gtksourceviewmmFormula) Print() {
    fmt.Printf("Name: gtksourceviewmm\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := gtksourceviewmmFormula{
        Description:  "C++ bindings for gtksourceview",
        Homepage:     "https://gitlab.gnome.org/GNOME/gtksourceviewmm",
        URL:          "https://download.gnome.org/sources/gtksourceviewmm/2.10/gtksourceviewmm-2.10.3.tar.xz",
        Sha256:       "cc5638fa500ed70f1de07180a729cee705fd3683305dabd4176f899bc7dfc17b",
        Dependencies: []string{"pkg-config", "gtkmm", "gtksourceview"},
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

    fmt.Printf("Installing formula: %s\n", "gtksourceviewmm")
    if err := pkg.Installgtksourceviewmm(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg gtksourceviewmmFormula) Installgtksourceviewmm() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "gtksourceviewmm-2.10.3.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "gtksourceviewmm-2.10.3.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
