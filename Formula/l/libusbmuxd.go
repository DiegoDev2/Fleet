
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// libusbmuxdFormula represents a formula in Go.
type libusbmuxdFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg libusbmuxdFormula) Print() {
    fmt.Printf("Name: libusbmuxd\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := libusbmuxdFormula{
        Description:  "USB multiplexor library for iOS devices",
        Homepage:     "https://www.libimobiledevice.org/",
        URL:          "https://github.com/libimobiledevice/libusbmuxd/releases/download/2.1.0/libusbmuxd-2.1.0.tar.bz2",
        Sha256:       "f3e9212633194e2e9436afd98d00fd5e167ec0ff13f1abfd065112dd2dd32bd3",
        Dependencies: []string{"autoconf", "automake", "libtool", "pkg-config", "libimobiledevice-glue", "libplist"},
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

    fmt.Printf("Installing formula: %s\n", "libusbmuxd")
    if err := pkg.Installlibusbmuxd(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg libusbmuxdFormula) Installlibusbmuxd() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "libusbmuxd-2.1.0.tar.bz2"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "libusbmuxd-2.1.0.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}