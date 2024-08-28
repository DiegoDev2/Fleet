
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// gpredictFormula represents a formula in Go.
type gpredictFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg gpredictFormula) Print() {
    fmt.Printf("Name: gpredict\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := gpredictFormula{
        Description:  "Real-time satellite tracking/prediction application",
        Homepage:     "https://gpredict.oz9aec.net/",
        URL:          "https://github.com/csete/gpredict.git",
        Sha256:       "e81cbab517c5a422abe3f0cc1a2fdac16b9c6d118949d343dfd30cd4d5026e8d",
        Dependencies: []string{"autoconf", "automake", "libtool", "autoconf", "automake", "libtool", "intltool", "pkg-config", "adwaita-icon-theme", "cairo", "gdk-pixbuf", "gettext", "glib", "goocanvas", "gtk+3", "hamlib", "pango", "at-spi2-core", "harfbuzz", "perl-xml-parser"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installgpredict(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg gpredictFormula) Installgpredict() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "gpredict.git"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "gpredict"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
