
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// spidermonkeyFormula represents a formula in Go.
type spidermonkeyFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg spidermonkeyFormula) Print() {
    fmt.Printf("Name: spidermonkey\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := spidermonkeyFormula{
        Description:  "JavaScript-C Engine",
        Homepage:     "https://spidermonkey.dev",
        URL:          "https://github.com/ptomato/mozjs/commit/9f778cec201f87fd68dc98380ac1097b2ff371e4.patch?full_index=1",
        Sha256:       "a772f39e5370d263fd7e182effb1b2b990cae8c63783f5a6673f16737ff91573",
        Dependencies: []string{"pkg-config", "python@3.11", "rust", "readline", "icu4c", "nspr"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installspidermonkey(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg spidermonkeyFormula) Installspidermonkey() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "9f778cec201f87fd68dc98380ac1097b2ff371e4.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "9f778cec201f87fd68dc98380ac1097b2ff371e4"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
