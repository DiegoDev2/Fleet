
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// rbenv-default-gemsFormula represents a formula in Go.
type rbenv-default-gemsFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg rbenv-default-gemsFormula) Print() {
    fmt.Printf("Name: rbenv-default-gems\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := rbenv-default-gemsFormula{
        Description:  "Auto-installs gems for Ruby installs",
        Homepage:     "https://github.com/rbenv/rbenv-default-gems",
        URL:          "https://github.com/rbenv/rbenv-default-gems/commit/ead67889c91c53ad967f85f5a89d986fdb98f6fb.patch?full_index=1",
        Sha256:       "ac6a5654c11d3ef74a97029ed86b8a7b6ae75f4ca7ff4d56df3fb35e7ae0acb8",
        Dependencies: []string{"rbenv"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installrbenv-default-gems(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg rbenv-default-gemsFormula) Installrbenv-default-gems() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "ead67889c91c53ad967f85f5a89d986fdb98f6fb.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "ead67889c91c53ad967f85f5a89d986fdb98f6fb"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
