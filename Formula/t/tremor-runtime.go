
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// tremor-runtimeFormula represents a formula in Go.
type tremor-runtimeFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg tremor-runtimeFormula) Print() {
    fmt.Printf("Name: tremor-runtime\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := tremor-runtimeFormula{
        Description:  "Early-stage event processing system for unstructured data",
        Homepage:     "https://www.tremor.rs/",
        URL:          "https://github.com/tremor-rs/tremor-runtime/commit/986fae5cf1022790e60175125b848dc84f67214f.patch?full_index=1",
        Sha256:       "ff772097264185213cbea09addbcdacc017eda4f90c97d0dad36b0156e3e9dbc",
        Dependencies: []string{"cmake", "pkg-config", "rust", "librdkafka", "oniguruma", "xz", "llvm@15"},
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

    fmt.Printf("Installing formula: %s\n", "tremor-runtime")
    if err := pkg.Installtremor-runtime(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg tremor-runtimeFormula) Installtremor-runtime() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "986fae5cf1022790e60175125b848dc84f67214f.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "986fae5cf1022790e60175125b848dc84f67214f"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
