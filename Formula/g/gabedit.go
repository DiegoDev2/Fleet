package main

import (
	"fmt"
	"log"
	"os/exec"
)

// gabeditFormula represents a formula in Go.
type gabeditFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg gabeditFormula) Print() {
	fmt.Printf("Name: gabedit\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := gabeditFormula{
		Description:  "GUI to computational chemistry packages like Gamess-US, Gaussian, etc.",
		Homepage:     "https://gabedit.sourceforge.net/",
		URL:          "https://sites.google.com/site/allouchear/Home/gabedit/download",
		Sha256:       "85b9f31c8d4f0b7daf320543f5da7d50a20212535bf1c753eca37aaa579c737d",
		Dependencies: []string{"pkg-config", "cairo", "gdk-pixbuf", "glib", "gtk+", "gtkglext", "pango", "at-spi2-core", "gettext", "harfbuzz", "autoconf", "automake", "libtool", "mesa", "mesa-glu"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installgabedit(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg gabeditFormula) Installgabedit() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "download"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "download"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
