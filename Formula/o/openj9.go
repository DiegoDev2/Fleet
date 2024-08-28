
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// openj9Formula represents a formula in Go.
type openj9Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg openj9Formula) Print() {
    fmt.Printf("Name: openj9\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := openj9Formula{
        Description:  "High performance, scalable, Java virtual machine",
        Homepage:     "https://www.eclipse.org/openj9/",
        URL:          "https://github.com/ibmruntimes/openj9-openjdk-jdk22.git",
        Sha256:       "6e54d984bc0c058ffb7a604810dfffba210d79e12855e5c61e9295fedeff32db",
        Dependencies: []string{"autoconf", "bash", "cmake", "ninja", "pkg-config", "fontconfig", "giflib", "harfbuzz", "jpeg-turbo", "libpng", "little-cms2", "alsa-lib", "libx11", "libxext", "libxrandr", "libxrender", "libxt", "libxtst", "numactl", "nasm"},
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

    fmt.Printf("Installing formula: %s\n", "openj9")
    if err := pkg.Installopenj9(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg openj9Formula) Installopenj9() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "openj9-openjdk-jdk22.git"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "openj9-openjdk-jdk22"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
