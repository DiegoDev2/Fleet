package main

import (
	"fmt"
	"log"
	"os/exec"
)

// gucharmapFormula represents a formula in Go.
type gucharmapFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg gucharmapFormula) Print() {
	fmt.Printf("Name: gucharmap\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := gucharmapFormula{
		Description:  "GNOME Character Map, based on the Unicode Character Database",
		Homepage:     "https://wiki.gnome.org/Apps/Gucharmap",
		URL:          "https://www.unicode.org/Public/16.0.0/ucd/Unihan.zip",
		Sha256:       "faed0321cf85bbc18fd3edb5f03dbae9d7099ea712d75f08d1fb2f87dcc197e1",
		Dependencies: []string{"desktop-file-utils", "gobject-introspection", "gtk-doc", "itstool", "meson", "ninja", "pkg-config", "vala", "at-spi2-core", "cairo", "glib", "gtk+3", "pango", "pcre2", "gettext", "gettext"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installgucharmap(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg gucharmapFormula) Installgucharmap() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "Unihan.zip"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "Unihan"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
