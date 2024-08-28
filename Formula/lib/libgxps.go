package main

import (
	"fmt"
	"log"
	"os/exec"
)

// libgxpsFormula represents a formula in Go.
type libgxpsFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg libgxpsFormula) Print() {
	fmt.Printf("Name: libgxps\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := libgxpsFormula{
		Description:  "GObject based library for handling and rendering XPS documents",
		Homepage:     "https://wiki.gnome.org/Projects/libgxps",
		URL:          "https://download.gnome.org/sources/libgxps/0.3/libgxps-0.3.2.tar.xz",
		Sha256:       "0fda080a2b1da025e6d6aef7a9e4934fefaabda48cc9c080088e1146cead5558",
		Dependencies: []string{"gobject-introspection", "meson", "ninja", "pkg-config", "cairo", "freetype", "glib", "jpeg-turbo", "libarchive", "libpng", "libtiff", "little-cms2", "gettext"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installlibgxps(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg libgxpsFormula) Installlibgxps() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "libgxps-0.3.2.tar.xz"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "libgxps-0.3.2.tar"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
