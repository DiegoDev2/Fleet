
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// rpm2cpioFormula represents a formula in Go.
type rpm2cpioFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg rpm2cpioFormula) Print() {
    fmt.Printf("Name: rpm2cpio\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := rpm2cpioFormula{
        Description:  "Tool to convert RPM package to CPIO archive",
        Homepage:     "https://svnweb.freebsd.org/ports/head/archivers/rpm2cpio/",
        URL:          "https://rpmfind.net/linux/fedora/linux/releases/39/Everything/x86_64/os/Packages/h/hello-2.12.1-2.fc39.x86_64.rpm",
        Sha256:       "10f9944f95ca54f224133cffab1cfab0c40e3adb64e4190d3d9e8f9dbed680f9",
        Dependencies: []string{"libarchive", "xz"},
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

    fmt.Printf("Installing formula: %s\n", "rpm2cpio")
    if err := pkg.Installrpm2cpio(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg rpm2cpioFormula) Installrpm2cpio() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "hello-2.12.1-2.fc39.x86_64.rpm"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "hello-2.12.1-2.fc39.x86_64"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
