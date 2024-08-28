
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// viceFormula represents a formula in Go.
type viceFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg viceFormula) Print() {
    fmt.Printf("Name: vice\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := viceFormula{
        Description:  "Versatile Commodore Emulator",
        Homepage:     "https://sourceforge.net/projects/vice-emu/",
        URL:          "https://downloads.sourceforge.net/project/vice-emu/releases/vice-3.8.tar.gz",
        Sha256:       "56cbe974416e6d4e257df439151d9fa8a4594e3a305df4aff2a42c10239472ef",
        Dependencies: []string{"autoconf", "automake", "dos2unix", "pkg-config", "texinfo", "xa", "yasm", "adwaita-icon-theme", "at-spi2-core", "cairo", "flac", "gdk-pixbuf", "giflib", "glew", "glib", "gtk+3", "lame", "libogg", "libpng", "librsvg", "libvorbis", "pango", "gettext", "harfbuzz", "alsa-lib", "fontconfig", "libx11", "mesa", "pulseaudio"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installvice(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg viceFormula) Installvice() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "vice-3.8.tar.gz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "vice-3.8.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
