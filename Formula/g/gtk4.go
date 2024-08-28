package main

import (
	"fmt"
	"log"
	"os/exec"
)

// gtk4Formula represents a formula in Go.
type gtk4Formula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg gtk4Formula) Print() {
	fmt.Printf("Name: gtk4\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := gtk4Formula{
		Description:  "Toolkit for creating graphical user interfaces",
		Homepage:     "https://gtk.org/",
		URL:          "https://download.gnome.org/sources/gtk/4.14/gtk-4.14.5.tar.xz",
		Sha256:       "9e53798297c2c6e3988a73aecf5e431f5164cb3bc959a553fc84dac76001707a",
		Dependencies: []string{"docbook", "docbook-xsl", "docutils", "gettext", "gobject-introspection", "meson", "ninja", "pkg-config", "sassc", "cairo", "fontconfig", "fribidi", "gdk-pixbuf", "glib", "graphene", "harfbuzz", "hicolor-icon-theme", "jpeg-turbo", "libepoxy", "libpng", "libtiff", "pango", "gettext", "libx11", "libxcursor", "libxdamage", "libxext", "libxfixes", "libxi", "libxinerama", "libxkbcommon", "libxrandr", "wayland"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installgtk4(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg gtk4Formula) Installgtk4() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "gtk-4.14.5.tar.xz"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "gtk-4.14.5.tar"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
