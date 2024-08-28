
package main

import (
    "fmt"
    "log"
    "os/exec"
)

// geeqieFormula represents a formula in Go.
type geeqieFormula struct {
    Description  string
    Homepage     string
    URL          string
    Sha256       string
    Dependencies []string
}

func (pkg geeqieFormula) Print() {
    fmt.Printf("Name: geeqie\n")
    fmt.Printf("Description: %s\n", pkg.Description)
    fmt.Printf("Homepage: %s\n", pkg.Homepage)
    fmt.Printf("URL: %s\n", pkg.URL)
    fmt.Printf("Sha256: %s\n", pkg.Sha256)
    fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
    pkg := geeqieFormula{
        Description:  "Lightweight Gtk+ based image viewer",
        Homepage:     "https://www.geeqie.org/",
        URL:          "https://github.com/BestImageViewer/geeqie/releases/download/v2.4/geeqie-2.4.tar.xz",
        Sha256:       "f0ead1a37d0842bb95b63fba36b56f3ea2668f017b70f10f5467586819ca3600",
        Dependencies: []string{"meson", "ninja", "pandoc", "pkg-config", "yelp-tools", "adwaita-icon-theme", "at-spi2-core", "cairo", "djvulibre", "evince", "exiv2", "ffmpegthumbnailer", "gdk-pixbuf", "gettext", "glib", "gspell", "gtk+3", "imagemagick", "jpeg-turbo", "jpeg-xl", "libarchive", "libheif", "libraw", "libtiff", "libx11", "little-cms2", "openjpeg", "pango", "poppler", "webp", "webp-pixbuf-loader", "enchant", "harfbuzz"},
    }

    pkg.Print()

    // Instalar dependencias
    for _, dep := range pkg.Dependencies {
        cmd := exec.Command("brew", "install", dep)
        if err := cmd.Run(); err != nil {
            log.Fatalf("Error installing dependency %s: %v", dep, err)
        }
    }

    if err := pkg.Installgeeqie(); err != nil {
        log.Fatalf("Error during installation: %v", err)
    }

    fmt.Println("Installation completed successfully.")
}

func (pkg geeqieFormula) Installgeeqie() error {
    cmd := exec.Command("curl", "-O", pkg.URL)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to download: %v", err)
    }

    tarball := "geeqie-2.4.tar.xz"
    cmd = exec.Command("tar", "-xf", tarball)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to extract tarball: %v", err)
    }

    sourceDir := "geeqie-2.4.tar"
    cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()

    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to configure and install: %v", err)
    }

    return nil
}
