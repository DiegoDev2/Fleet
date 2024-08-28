
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// libsigrokFormula represents a formula in Go.
type libsigrokFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg libsigrokFormula) Print() {
    fmt.Printf("Name: libsigrok\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := libsigrokFormula{
        Description:  "Drivers for logic analyzers and other supported devices",
        Homepage:     "https://sigrok.org/",
        URL:          "https://sigrok.org/download/binary/sigrok-firmware-fx2lafw/sigrok-firmware-fx2lafw-bin-0.1.7.tar.gz",
        Sha256:       "c876fd075549e7783a6d5bfc8d99a695cfc583ddbcea0217d8e3f9351d1723af",
        Dependencies: []string{"autoconf", "autoconf-archive", "automake", "doxygen", "graphviz", "libtool", "pkg-config", "python-setuptools", "sdcc", "swig", "glib", "glibmm@2.66", "hidapi", "libftdi", "libusb", "libzip", "nettle", "numpy", "pygobject3", "python@3.12", "gettext", "libsigc++@2"},
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

    fmt.Printf("Installing formula: %s\n", "libsigrok")
    if err := pkg.Installlibsigrok(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg libsigrokFormula) Installlibsigrok() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "sigrok-firmware-fx2lafw-bin-0.1.7.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "sigrok-firmware-fx2lafw-bin-0.1.7.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
