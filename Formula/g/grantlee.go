
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// grantleeFormula represents a formula in Go.
type grantleeFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg grantleeFormula) Print() {
    fmt.Printf("Name: grantlee\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := grantleeFormula{
        Description:  "Libraries for text templating with Qt",
        Homepage:     "https://github.com/steveire/grantlee",
        URL:          "https://github.com/steveire/grantlee/commit/1efeb1cb61947e69b8c99ddbfc5b75cd27013a87.patch?full_index=1",
        Sha256:       "6c5fa321c5df2b970ec2873df610ec43dd2d50977cb0a104d0d7c4ecb90621c2",
        Dependencies: []string{"cmake", "doxygen", "graphviz", "qt"},
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

    fmt.Printf("Installing formula: %s\n", "grantlee")
    if err := pkg.Installgrantlee(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg grantleeFormula) Installgrantlee() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "1efeb1cb61947e69b8c99ddbfc5b75cd27013a87.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "1efeb1cb61947e69b8c99ddbfc5b75cd27013a87"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
