
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// ghc@8.6Formula represents a formula in Go.
type ghc@8.6Formula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg ghc@8.6Formula) Print() {
    fmt.Printf("Name: ghc@8.6\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := ghc@8.6Formula{
        Description:  "Glorious Glasgow Haskell Compilation System",
        Homepage:     "https://haskell.org/ghc/",
        URL:          "https://downloads.haskell.org/~ghc/8.6.5/ghc-8.6.5-x86_64-fedora27-linux.tar.xz",
        Sha256:       "cf78b53eaf336083e7a05f4a3000afbae4abe5bbc77ef80cc40e09d04ac5b4a1",
        Dependencies: []string{"python@3.10", "gmp"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installghc@8.6(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg ghc@8.6Formula) Installghc@8.6() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "ghc-8.6.5-x86_64-fedora27-linux.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "ghc-8.6.5-x86_64-fedora27-linux.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
