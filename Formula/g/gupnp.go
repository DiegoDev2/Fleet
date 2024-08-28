package main

import (
	"fmt"
	"log"
	"os/exec"
)

// gupnpFormula represents a formula in Go.
type gupnpFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg gupnpFormula) Print() {
	fmt.Printf("Name: gupnp\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := gupnpFormula{
		Description:  "Framework for creating UPnP devices and control points",
		Homepage:     "https://wiki.gnome.org/Projects/GUPnP",
		URL:          "https://gitlab.gnome.org/GNOME/gupnp/-/commit/00514fb62ebd341803fa44e26a6482a8c25dbd34.diff",
		Sha256:       "2b8ead2dc0824bf30dc606421cff3cddc7d8ad785910b1228602bb861601f61c",
		Dependencies: []string{"docbook-xsl", "gobject-introspection", "meson", "ninja", "pkg-config", "vala", "gettext", "glib", "gssdp", "libsoup", "libxml2", "python@3.12"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installgupnp(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg gupnpFormula) Installgupnp() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "00514fb62ebd341803fa44e26a6482a8c25dbd34.diff"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "00514fb62ebd341803fa44e26a6482a8c25dbd34"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
