
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// the_platinum_searcherFormula represents a formula in Go.
type the_platinum_searcherFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg the_platinum_searcherFormula) Print() {
    fmt.Printf("Name: the_platinum_searcher\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := the_platinum_searcherFormula{
        Description:  "Multi-platform code-search similar to ack and ag",
        Homepage:     "https://github.com/monochromegane/the_platinum_searcher",
        URL:          "https://github.com/monochromegane/the_platinum_searcher/commit/763f368fe26fa44a12e1a37598185322aa30ba8f.patch?full_index=1",
        Sha256:       "2ee0f53065663f22f3c44b30c5804e37b8cb49200a30c4513b9ef668441dd543",
        Dependencies: []string{"go"},
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

    fmt.Printf("Installing formula: %s\n", "the_platinum_searcher")
    if err := pkg.Installthe_platinum_searcher(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func isFormulaInstalled(name string) bool {
    cmd := exec.Command("brew", "list", name)
    err := cmd.Run()
    return err == nil
}

func (pkg the_platinum_searcherFormula) Installthe_platinum_searcher() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "763f368fe26fa44a12e1a37598185322aa30ba8f.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "763f368fe26fa44a12e1a37598185322aa30ba8f"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./configure && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
