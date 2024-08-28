
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// hashlinkFormula represents a formula in Go.
type hashlinkFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg hashlinkFormula) Print() {
    fmt.Printf("Name: hashlink\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := hashlinkFormula{
        Description:  "Virtual machine for Haxe",
        Homepage:     "https://hashlink.haxe.org/",
        URL:          "https://github.com/HaxeFoundation/hashlink/commit/54e97e34f29e80bcdccdb69af8ccd02bd7c0bc3a.patch?full_index=1",
        Sha256:       "2b9d5a710e10e612ba4ac82c74bec61354d9606912735e7f6a706298a68fcf75",
        Dependencies: []string{"haxe", "jpeg-turbo", "libogg", "libpng", "libuv", "libvorbis", "mbedtls", "openal-soft", "sdl2", "mesa", "mesa-glu"},
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

    fmt.Printf("Installing formula: %s\n", "hashlink")
    if err := pkg.Installhashlink(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg hashlinkFormula) Installhashlink() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "54e97e34f29e80bcdccdb69af8ccd02bd7c0bc3a.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "54e97e34f29e80bcdccdb69af8ccd02bd7c0bc3a"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
