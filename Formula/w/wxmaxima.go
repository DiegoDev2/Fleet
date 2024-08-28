
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// wxmaximaFormula represents a formula in Go.
type wxmaximaFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg wxmaximaFormula) Print() {
    fmt.Printf("Name: wxmaxima\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := wxmaximaFormula{
        Description:  "Cross platform GUI for Maxima",
        Homepage:     "https://wxmaxima-developers.github.io/wxmaxima/",
        URL:          "https://github.com/wxMaxima-developers/wxmaxima/commit/077ec646a11bfb5aa83a478e636a715a38a9b68b.patch?full_index=1",
        Sha256:       "15fb4db52cb7e1237ee5d0934653db06809d172e2bf54709435ec24d1f7ab7a9",
        Dependencies: []string{"cmake", "gettext", "ninja", "maxima", "wxwidgets", "llvm"},
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

    fmt.Printf("Installing formula: %s\n", "wxmaxima")
    if err := pkg.Installwxmaxima(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg wxmaximaFormula) Installwxmaxima() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "077ec646a11bfb5aa83a478e636a715a38a9b68b.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "077ec646a11bfb5aa83a478e636a715a38a9b68b"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
