
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// vulkan-toolsFormula represents a formula in Go.
type vulkan-toolsFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg vulkan-toolsFormula) Print() {
    fmt.Printf("Name: vulkan-tools\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := vulkan-toolsFormula{
        Description:  "Vulkan utilities and tools",
        Homepage:     "https://github.com/KhronosGroup/Vulkan-Tools",
        URL:          "https://github.com/KhronosGroup/Vulkan-Tools/archive/refs/tags/v1.3.294.tar.gz",
        Sha256:       "50ab331eb582d4d5f6bbde13c1ca49bb8f48ade9b65f5e645e00605fc4cc4e67",
        Dependencies: []string{"cmake", "python@3.12", "vulkan-volk", "glslang", "vulkan-headers", "vulkan-loader", "molten-vk", "pkg-config", "libx11", "libxcb", "libxkbfile", "libxrandr", "wayland", "wayland-protocols"},
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

    fmt.Printf("Installing formula: %s\n", "vulkan-tools")
    if err := pkg.Installvulkan-tools(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg vulkan-toolsFormula) Installvulkan-tools() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "v1.3.294.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "v1.3.294.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
