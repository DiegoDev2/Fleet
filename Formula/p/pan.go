package main

import (
	"fmt"
	"log"
	"os/exec"
)

// panFormula represents a formula in Go.
type panFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg panFormula) Print() {
	fmt.Printf("Name: pan\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := panFormula{
		Description:  "Usenet newsreader that's good at both text and binaries",
		Homepage:     "https://gitlab.gnome.org/GNOME/pan",
		URL:          "https://gitlab.gnome.org/GNOME/pan/-/archive/v0.160/pan-v0.160.tar.bz2",
		Sha256:       "fa9be5e5e4c9047dbaf84acf4920464e2c04767de4826943f39d76719b8e92fb",
		Dependencies: []string{"cmake", "pkg-config", "adwaita-icon-theme", "cairo", "enchant", "gdk-pixbuf", "gettext", "glib", "gmime", "gnutls", "gtk+3", "gtkspell3", "harfbuzz", "pango", "at-spi2-core"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installpan(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg panFormula) Installpan() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "pan-v0.160.tar.bz2"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "pan-v0.160.tar"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
