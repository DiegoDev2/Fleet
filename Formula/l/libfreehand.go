
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// libfreehandFormula represents a formula in Go.
type libfreehandFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg libfreehandFormula) Print() {
    fmt.Printf("Name: libfreehand\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := libfreehandFormula{
        Description:  "Interpret and import Aldus/Macromedia/Adobe FreeHand documents",
        Homepage:     "https://wiki.documentfoundation.org/DLP/Libraries/libfreehand",
        URL:          "https://raw.githubusercontent.com/Homebrew/formula-patches/7bb2149f314dd174f242a76d4dde8d95d20cbae0/libfreehand/0.1.2.patch",
        Sha256:       "abfa28461b313ccf3c59ce35d0a89d0d76c60dd2a14028b8fea66e411983160e",
        Dependencies: []string{"boost", "pkg-config", "librevenge", "little-cms2"},
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

    fmt.Printf("Installing formula: %s\n", "libfreehand")
    if err := pkg.Installlibfreehand(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg libfreehandFormula) Installlibfreehand() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "0.1.2.patch"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "0.1.2"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
