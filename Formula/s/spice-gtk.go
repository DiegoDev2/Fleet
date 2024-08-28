
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// spice-gtkFormula represents a formula in Go.
type spice-gtkFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg spice-gtkFormula) Print() {
    fmt.Printf("Name: spice-gtk\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := spice-gtkFormula{
        Description:  "GTK client/libraries for SPICE",
        Homepage:     "https://www.spice-space.org",
        URL:          "https://gitlab.freedesktop.org/spice/spice-gtk/-/commit/1511f0ad5ea67b4657540c631e3a8c959bb8d578.diff",
        Sha256:       "67c2b1d9c689dbb8eb3ed7c92996cf8c9d083d51050883593ee488957ad2a083",
        Dependencies: []string{"gobject-introspection", "intltool", "libtool", "meson", "ninja", "pkg-config", "python@3.12", "vala", "at-spi2-core", "cairo", "gdk-pixbuf", "gettext", "glib", "gstreamer", "gtk+3", "jpeg-turbo", "json-glib", "libepoxy", "libsoup", "libusb", "libx11", "lz4", "openssl@3", "opus", "pango", "phodav", "pixman", "spice-protocol", "usbredir", "gobject-introspection", "harfbuzz", "libva", "wayland"},
    }

    pkg.Print()

    // Instalar dependencias si no están instaladas
    for _, dep := range pkg.Dependencies {
        if !isDependencyInstalled(dep) {
            fmt.Printf("🛠️ Dependency %s not found. Installing...
", dep)
            cmd := exec.Command("brew", "install", dep)
            if err := cmd.Run(); err != nil {
                log.Fatalf("Error installing dependency %s: %v", dep, err)
            }
        } else {
            fmt.Printf("✅ Dependency %s is already installed.
", dep)
        }
    }

    if err := pkg.Installspice-gtk(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg spice-gtkFormula) Installspice-gtk() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "1511f0ad5ea67b4657540c631e3a8c959bb8d578.diff"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "1511f0ad5ea67b4657540c631e3a8c959bb8d578"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}

func isDependencyInstalled(dep string) bool {
    cmd := exec.Command("brew", "list", dep)
    output, err := cmd.CombinedOutput()
    return err == nil && strings.TrimSpace(string(output)) != ""
}
