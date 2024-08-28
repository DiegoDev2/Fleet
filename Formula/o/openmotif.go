
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// openmotifFormula represents a formula in Go.
type openmotifFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg openmotifFormula) Print() {
    fmt.Printf("Name: openmotif\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := openmotifFormula{
        Description:  "LGPL release of the Motif toolkit",
        Homepage:     "https://motif.ics.com/motif",
        URL:          "https://raw.githubusercontent.com/macports/macports-ports/8c436a9c53a7b786da8d42cda16eead0fb8733d4/x11/openmotif/files/patch-lib-xm-vendor.diff",
        Sha256:       "697ac026386dec59b82883fb4a9ba77164dd999fa3fb0569dbc8fbdca57fe200",
        Dependencies: []string{"pkg-config", "fontconfig", "freetype", "jpeg-turbo", "libice", "libpng", "libsm", "libx11", "libxext", "libxft", "libxmu", "libxp", "libxt", "xbitmaps"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installopenmotif(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg openmotifFormula) Installopenmotif() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "patch-lib-xm-vendor.diff"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "patch-lib-xm-vendor"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
