
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// tiger-vncFormula represents a formula in Go.
type tiger-vncFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg tiger-vncFormula) Print() {
    fmt.Printf("Name: tiger-vnc\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := tiger-vncFormula{
        Description:  "High-performance, platform-neutral implementation of VNC",
        Homepage:     "https://tigervnc.org/",
        URL:          "https://github.com/TigerVNC/tigervnc/archive/refs/tags/v1.14.0.tar.gz",
        Sha256:       "91194d51191ea03eea1a927148efc052dce585bdbbb0abf5a83652626285e160",
        Dependencies: []string{"cmake", "fltk", "gettext", "gmp", "gnutls", "jpeg-turbo", "nettle", "pixman", "libx11", "libxcursor", "libxdamage", "libxext", "libxfixes", "libxft", "libxi", "libxinerama", "libxrandr", "libxrender", "libxtst", "linux-pam"},
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

    fmt.Printf("Installing formula: %s\n", "tiger-vnc")
    if err := pkg.Installtiger-vnc(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg tiger-vncFormula) Installtiger-vnc() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "v1.14.0.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "v1.14.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
