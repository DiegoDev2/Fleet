
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// ios-webkit-debug-proxyFormula represents a formula in Go.
type ios-webkit-debug-proxyFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg ios-webkit-debug-proxyFormula) Print() {
    fmt.Printf("Name: ios-webkit-debug-proxy\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := ios-webkit-debug-proxyFormula{
        Description:  "DevTools proxy for iOS devices",
        Homepage:     "https://github.com/google/ios-webkit-debug-proxy",
        URL:          "https://github.com/google/ios-webkit-debug-proxy/archive/refs/tags/v1.9.1.tar.gz",
        Sha256:       "44137a86b9210566d60a3b8b8e09f5cca01340f2db9441c6a29375df8e8e6db9",
        Dependencies: []string{"autoconf", "automake", "libtool", "pkg-config", "libimobiledevice", "libplist", "libusbmuxd", "openssl@3"},
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

    fmt.Printf("Installing formula: %s\n", "ios-webkit-debug-proxy")
    if err := pkg.Installios-webkit-debug-proxy(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg ios-webkit-debug-proxyFormula) Installios-webkit-debug-proxy() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "v1.9.1.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "v1.9.1.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
