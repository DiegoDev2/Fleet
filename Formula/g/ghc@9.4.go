
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// ghc@9.4Formula represents a formula in Go.
type ghc@9.4Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg ghc@9.4Formula) Print() {
    fmt.Printf("Name: ghc@9.4\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := ghc@9.4Formula{
        Description:  "Glorious Glasgow Haskell Compilation System",
        Homepage:     "https://haskell.org/ghc/",
        URL:          "https://gitlab.haskell.org/ghc/ghc/-/commit/70526f5bd8886126f49833ef20604a2c6477780a.diff",
        Sha256:       "54cdde1ca5d1b6fe3bbad8d0eac2b8c112ca1f346c4086d1e7361fa9510f1f44",
        Dependencies: []string{"autoconf", "automake", "python@3.12", "sphinx-doc", "gnu-sed", "gmp"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installghc@9.4(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg ghc@9.4Formula) Installghc@9.4() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "70526f5bd8886126f49833ef20604a2c6477780a.diff"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "70526f5bd8886126f49833ef20604a2c6477780a"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
