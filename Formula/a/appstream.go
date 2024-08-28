
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// appstreamFormula represents a formula in Go.
type appstreamFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg appstreamFormula) Print() {
    fmt.Printf("Name: appstream\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := appstreamFormula{
        Description:  "Tools and libraries to work with AppStream metadata",
        Homepage:     "https://www.freedesktop.org/wiki/Distributions/AppStream/",
        URL:          "https://github.com/ximion/appstream/commit/06eeffe7eba5c4e82a1dd548e100c6fe4f71b413.patch?full_index=1",
        Sha256:       "d0ad5853d451eb073fc64bd3e9e58e81182f4142220e0f413794752cda235d28",
        Dependencies: []string{"gobject-introspection", "gtk-doc", "itstool", "meson", "ninja", "pkg-config", "vala", "glib", "libxmlb", "libyaml", "zstd", "gettext", "gettext", "gperf", "systemd"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installappstream(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg appstreamFormula) Installappstream() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "06eeffe7eba5c4e82a1dd548e100c6fe4f71b413.patch?full_index=1"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "06eeffe7eba5c4e82a1dd548e100c6fe4f71b413"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
