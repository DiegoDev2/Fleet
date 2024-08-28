package main

import (
	"fmt"
	"log"
	"os/exec"
)

// freecivFormula represents a formula in Go.
type freecivFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg freecivFormula) Print() {
	fmt.Printf("Name: freeciv\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := freecivFormula{
		Description:  "Free and Open Source empire-building strategy game",
		Homepage:     "https://freeciv.org/",
		URL:          "https://github.com/freeciv/freeciv.git",
		Sha256:       "32b2b6cab1b82d3863f0db79617dcb41519604e6f8e0dbb4ca81fe74e224ca64",
		Dependencies: []string{"autoconf", "automake", "gettext", "libtool", "pkg-config", "adwaita-icon-theme", "at-spi2-core", "cairo", "freetype", "gdk-pixbuf", "gettext", "glib", "gtk+3", "harfbuzz", "icu4c", "pango", "readline", "sdl2", "sdl2_mixer", "sqlite", "zstd"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installfreeciv(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg freecivFormula) Installfreeciv() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "freeciv.git"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "freeciv"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
