
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// augeasFormula represents a formula in Go.
type augeasFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg augeasFormula) Print() {
    fmt.Printf("Name: augeas\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := augeasFormula{
        Description:  "Configuration editing tool and API",
        Homepage:     "https://augeas.net/",
        URL:          "https://github.com/hercules-team/augeas/commit/26d297825000dd2cdc45d0fa6bf68dcc14b08d7d.patch?full_index=1",
        Sha256:       "6bed3c3201eabb1849cbc729d42e33a3692069a06d298ce3f4a8bce7cdbf9f0e",
        Dependencies: []string{"autoconf", "automake", "bison", "libtool", "pkg-config", "readline"},
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

    fmt.Printf("Installing formula: %s\n", "augeas")
    if err := pkg.Installaugeas(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg augeasFormula) Installaugeas() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "26d297825000dd2cdc45d0fa6bf68dcc14b08d7d.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "26d297825000dd2cdc45d0fa6bf68dcc14b08d7d"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
