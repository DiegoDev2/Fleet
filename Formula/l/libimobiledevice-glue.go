
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// libimobiledevice-glueFormula represents a formula in Go.
type libimobiledevice-glueFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg libimobiledevice-glueFormula) Print() {
    fmt.Printf("Name: libimobiledevice-glue\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := libimobiledevice-glueFormula{
        Description:  "Library with common system API code for libimobiledevice projects",
        Homepage:     "https://libimobiledevice.org/",
        URL:          "https://github.com/libimobiledevice/libimobiledevice-glue/releases/download/1.2.0/libimobiledevice-glue-1.2.0.tar.bz2",
        Sha256:       "1c30c5964567df2dcb1b86e70e41622aaf72d533effecc6ee0ef7348171eefdd",
        Dependencies: []string{"autoconf", "automake", "libtool", "pkg-config", "libplist"},
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

    fmt.Printf("Installing formula: %s\n", "libimobiledevice-glue")
    if err := pkg.Installlibimobiledevice-glue(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg libimobiledevice-glueFormula) Installlibimobiledevice-glue() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "libimobiledevice-glue-1.2.0.tar.bz2"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "libimobiledevice-glue-1.2.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
