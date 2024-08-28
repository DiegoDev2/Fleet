
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// gtk-gnutellaFormula represents a formula in Go.
type gtk-gnutellaFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg gtk-gnutellaFormula) Print() {
    fmt.Printf("Name: gtk-gnutella\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := gtk-gnutellaFormula{
        Description:  "Share files in a peer-to-peer (P2P) network",
        Homepage:     "https://gtk-gnutella.sourceforge.io",
        URL:          "https://downloads.sourceforge.net/project/gtk-gnutella/gtk-gnutella/1.2.3/gtk-gnutella-1.2.3.tar.xz",
        Sha256:       "70a5946bf77166b5076fe6fa1d45b69c6f512260451c96758c4cde02fbb983df",
        Dependencies: []string{"pkg-config", "at-spi2-core", "dbus", "gdk-pixbuf", "glib", "gtk+", "pango", "gettext", "harfbuzz"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installgtk-gnutella(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg gtk-gnutellaFormula) Installgtk-gnutella() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "gtk-gnutella-1.2.3.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "gtk-gnutella-1.2.3.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
