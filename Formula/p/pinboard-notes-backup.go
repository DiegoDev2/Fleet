
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// pinboard-notes-backupFormula represents a formula in Go.
type pinboard-notes-backupFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg pinboard-notes-backupFormula) Print() {
    fmt.Printf("Name: pinboard-notes-backup\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := pinboard-notes-backupFormula{
        Description:  "Efficiently back up the notes you've saved to Pinboard",
        Homepage:     "https://github.com/bdesham/pinboard-notes-backup",
        URL:          "https://github.com/bdesham/pinboard-notes-backup/commit/8be2ac9107b312657f0ae68633164ac2ea85ee9e.patch?full_index=1",
        Sha256:       "19365a6c0ca7c4d6a56c8bbfd591346ab77c23b89c7e4a2aec77ddc9eba57fe7",
        Dependencies: []string{"cabal-install", "ghc"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installpinboard-notes-backup(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg pinboard-notes-backupFormula) Installpinboard-notes-backup() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "8be2ac9107b312657f0ae68633164ac2ea85ee9e.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "8be2ac9107b312657f0ae68633164ac2ea85ee9e"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}