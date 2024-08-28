
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// mesaFormula represents a formula in Go.
type mesaFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg mesaFormula) Print() {
    fmt.Printf("Name: mesa\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := mesaFormula{
        Description:  "Graphics Library",
        Homepage:     "https://www.mesa3d.org/",
        URL:          "https://gitlab.freedesktop.org/mesa/demos/-/raw/ddc35ca0ea2f18c5011c5573b4b624c128ca7616/src/util/gl_wrap.h",
        Sha256:       "41f5a84f8f5abe8ea2a21caebf5ff31094a46953a83a738a19e21c010c433c88",
        Dependencies: []string{"bison", "meson", "ninja", "pkg-config", "python@3.12", "xorgproto", "expat", "libx11", "libxcb", "libxext", "libxfixes", "libxrandr", "elfutils", "glslang", "gzip", "libclc", "libdrm", "libva", "libvdpau", "libxml2", "libxshmfence", "libxv", "libxxf86vm", "lm-sensors", "spirv-llvm-translator", "valgrind", "wayland", "wayland-protocols", "zstd"},
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

    fmt.Printf("Installing formula: %s\n", "mesa")
    if err := pkg.Installmesa(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg mesaFormula) Installmesa() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "gl_wrap.h"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "gl_wrap"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
