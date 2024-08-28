
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// emscriptenFormula represents a formula in Go.
type emscriptenFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg emscriptenFormula) Print() {
    fmt.Printf("Name: emscripten\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := emscriptenFormula{
        Description:  "LLVM bytecode to JavaScript compiler",
        Homepage:     "https://emscripten.org/",
        URL:          "https://github.com/llvm/llvm-project/archive/547917aebd1e79a8929b53f0ddf3b5185ee4df74.tar.gz",
        Sha256:       "d2f3bcf38edd5168f3e71b200bf19b6b3f0c7a459a73a8d7ad9486c7b74472bf",
        Dependencies: []string{"cmake", "node", "python@3.12", "yuicompressor", "openjdk", "openjdk"},
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

    fmt.Printf("Installing formula: %s\n", "emscripten")
    if err := pkg.Installemscripten(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg emscriptenFormula) Installemscripten() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "547917aebd1e79a8929b53f0ddf3b5185ee4df74.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "547917aebd1e79a8929b53f0ddf3b5185ee4df74.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
