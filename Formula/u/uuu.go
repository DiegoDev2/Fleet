
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// uuuFormula represents a formula in Go.
type uuuFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg uuuFormula) Print() {
    fmt.Printf("Name: uuu\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := uuuFormula{
        Description:  "Universal Update Utility, mfgtools 3.0. NXP I.MX Chip image deploy tools",
        Homepage:     "https://github.com/nxp-imx/mfgtools",
        URL:          "https://github.com/nxp-imx/mfgtools/releases/download/uuu_1.5.182/uuu_source-uuu_1.5.182.tar.gz",
        Sha256:       "19e6bebfc3fdb36ef5b0ee3c517e294c75a2ebc9460501a104c00d7586239616",
        Dependencies: []string{"cmake", "pkg-config", "libusb", "libzip", "openssl@3", "tinyxml2", "zstd"},
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

    fmt.Printf("Installing formula: %s\n", "uuu")
    if err := pkg.Installuuu(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg uuuFormula) Installuuu() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "uuu_source-uuu_1.5.182.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "uuu_source-uuu_1.5.182.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
