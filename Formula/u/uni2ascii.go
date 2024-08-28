
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// uni2asciiFormula represents a formula in Go.
type uni2asciiFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg uni2asciiFormula) Print() {
    fmt.Printf("Name: uni2ascii\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := uni2asciiFormula{
        Description:  "Bi-directional conversion between UTF-8 and various ASCII flavors",
        Homepage:     "https://billposer.org/Software/uni2ascii.html",
        URL:          "https://raw.githubusercontent.com/Homebrew/formula-patches/bb92449ad6b3878b4d6f472237152df28080df86/uni2ascii/uni2ascii-4.20.patch",
        Sha256:       "250a529eda136d0edf9e63b92a6fe95f4ef5dfad3f94e6fd8d877138ada857f8",
        Dependencies: []string{"gettext"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installuni2ascii(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg uni2asciiFormula) Installuni2ascii() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "uni2ascii-4.20.patch"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "uni2ascii-4.20"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
