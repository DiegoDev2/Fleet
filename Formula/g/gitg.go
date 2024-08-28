package main

import (
	"fmt"
	"log"
	"os/exec"
)

// gitgFormula represents a formula in Go.
type gitgFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg gitgFormula) Print() {
	fmt.Printf("Name: gitg\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := gitgFormula{
		Description:  "GNOME GUI client to view git repositories",
		Homepage:     "https://wiki.gnome.org/Apps/Gitg",
		URL:          "https://download.gnome.org/sources/gitg/44/gitg-44.tar.xz",
		Sha256:       "0cbc0e25ab157fe48820366df24a9ea57d105ece188f62d9bb27110198977d96",
		Dependencies: []string{"gettext", "meson", "ninja", "pkg-config", "vala", "adwaita-icon-theme", "cairo", "gdk-pixbuf", "glib", "gobject-introspection", "gpgme", "gspell", "gtk+3", "gtksourceview4", "hicolor-icon-theme", "json-glib", "libdazzle", "libgee", "libgit2-glib", "libgit2@1.7", "libhandy", "libpeas@1", "libsecret", "pango", "gettext"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installgitg(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg gitgFormula) Installgitg() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "gitg-44.tar.xz"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "gitg-44.tar"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
