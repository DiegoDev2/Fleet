package main

import (
	"fmt"
	"log"
	"os/exec"
)

// libsvgFormula represents a formula in Go.
type libsvgFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg libsvgFormula) Print() {
	fmt.Printf("Name: libsvg\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := libsvgFormula{
		Description:  "Library for SVG files",
		Homepage:     "https://cairographics.org/",
		URL:          "https://raw.githubusercontent.com/buildroot/buildroot/45c3b0ec49fac67cc81651f0bed063722a48dc29/package/libsvg/0002-Fix-undefined-symbol-png_set_gray_1_2_4_to_8.patch",
		Sha256:       "a0ca1e25ea6bd5cb9aac57ac541c90ebe3b12c1340dbc5762d487d827064e0b9",
		Dependencies: []string{"autoconf", "automake", "libtool", "pkg-config", "jpeg-turbo", "libpng"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installlibsvg(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg libsvgFormula) Installlibsvg() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "0002-Fix-undefined-symbol-png_set_gray_1_2_4_to_8.patch"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "0002-Fix-undefined-symbol-png_set_gray_1_2_4_to_8"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
