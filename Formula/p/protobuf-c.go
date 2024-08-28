
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// protobuf-cFormula represents a formula in Go.
type protobuf-cFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg protobuf-cFormula) Print() {
    fmt.Printf("Name: protobuf-c\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := protobuf-cFormula{
        Description:  "Protocol buffers library",
        Homepage:     "https://github.com/protobuf-c/protobuf-c",
        URL:          "https://github.com/protobuf-c/protobuf-c/commit/d95aced22df60a2f0049fc03af48c8b02ce4d474.patch?full_index=1",
        Sha256:       "7aa44807367a4547bd15b3aa9a5275d5fe4739348bf2741ca773fa47015fb01a",
        Dependencies: []string{"asciidoc", "autoconf", "automake", "libtool", "pkg-config", "abseil", "protobuf"},
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

    fmt.Printf("Installing formula: %s\n", "protobuf-c")
    if err := pkg.Installprotobuf-c(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg protobuf-cFormula) Installprotobuf-c() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "d95aced22df60a2f0049fc03af48c8b02ce4d474.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "d95aced22df60a2f0049fc03af48c8b02ce4d474"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
