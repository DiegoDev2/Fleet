package main

import (
	"fmt"
	"log"
	"os/exec"
)

// libdazzleFormula represents a formula in Go.
type libdazzleFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg libdazzleFormula) Print() {
	fmt.Printf("Name: libdazzle\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := libdazzleFormula{
		Description:  "GNOME companion library to GObject and Gtk+",
		Homepage:     "https://gitlab.gnome.org/GNOME/libdazzle",
		URL:          "https://download.gnome.org/sources/libdazzle/3.44/libdazzle-3.44.0.tar.xz",
		Sha256:       "7a8395fdb2257d54ffc90ebcb1a95796c87fba132b602c14ed0c42e384bb90e8",
		Dependencies: []string{"gobject-introspection", "meson", "ninja", "pkg-config", "vala", "cairo", "gdk-pixbuf", "glib", "gtk+3", "pango", "at-spi2-core", "gettext", "harfbuzz"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installlibdazzle(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg libdazzleFormula) Installlibdazzle() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "libdazzle-3.44.0.tar.xz"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "libdazzle-3.44.0.tar"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
