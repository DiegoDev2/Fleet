
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// xdotoolFormula represents a formula in Go.
type xdotoolFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg xdotoolFormula) Print() {
    fmt.Printf("Name: xdotool\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := xdotoolFormula{
        Description:  "Fake keyboard/mouse input and window management for X",
        Homepage:     "https://www.semicomplete.com/projects/xdotool/",
        URL:          "https://github.com/jordansissel/xdotool/commit/dffc9a1597bd96c522a2b71c20301f97c130b7a8.patch?full_index=1",
        Sha256:       "447fa42ec274eb7488bb4aeeccfaaba0df5ae747f1a7d818191698035169a5ef",
        Dependencies: []string{"pkg-config", "libx11", "libxinerama", "libxkbcommon", "libxtst"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installxdotool(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg xdotoolFormula) Installxdotool() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "dffc9a1597bd96c522a2b71c20301f97c130b7a8.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "dffc9a1597bd96c522a2b71c20301f97c130b7a8"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
