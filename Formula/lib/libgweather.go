package main

import (
	"fmt"
	"log"
	"os/exec"
)

// libgweatherFormula represents a formula in Go.
type libgweatherFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg libgweatherFormula) Print() {
	fmt.Printf("Name: libgweather\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := libgweatherFormula{
		Description:  "GNOME library for weather, locations and timezones",
		Homepage:     "https://wiki.gnome.org/Projects/LibGWeather",
		URL:          "https://download.gnome.org/sources/libgweather/4.4/libgweather-4.4.2.tar.xz",
		Sha256:       "5489a2ac860e4a8bbeb4545939d78eaa5b975a6cd6bd7e619f0ddb15a357f00b",
		Dependencies: []string{"gettext", "gobject-introspection", "meson", "ninja", "pkg-config", "pygobject3", "python@3.12", "geocode-glib", "glib", "gtk+3", "json-glib", "libsoup", "gettext"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installlibgweather(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg libgweatherFormula) Installlibgweather() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "libgweather-4.4.2.tar.xz"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "libgweather-4.4.2.tar"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
