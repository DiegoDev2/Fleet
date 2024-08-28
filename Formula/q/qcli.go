
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// qcliFormula represents a formula in Go.
type qcliFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg qcliFormula) Print() {
    fmt.Printf("Name: qcli\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := qcliFormula{
        Description:  "Report audiovisual metrics via libavfilter",
        Homepage:     "https://bavc.org/preserve-media/preservation-tools",
        URL:          "https://github.com/valbok/QtAVPlayer/commit/a24033d6636426d4fc1f97b6ee483bf6a7ab6072.patch?full_index=1",
        Sha256:       "14e72317eceef670be0d732541e176f143ffafc616cb698dc4052b5c1c589c92",
        Dependencies: []string{"pkg-config", "ffmpeg@6", "qt", "qwt"},
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

    fmt.Printf("Installing formula: %s\n", "qcli")
    if err := pkg.Installqcli(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg qcliFormula) Installqcli() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "a24033d6636426d4fc1f97b6ee483bf6a7ab6072.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "a24033d6636426d4fc1f97b6ee483bf6a7ab6072"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
