package main

import (
	"fmt"
	"log"
	"os/exec"
)

// allegroFormula represents a formula in Go.
type allegroFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg allegroFormula) Print() {
	fmt.Printf("Name: allegro\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := allegroFormula{
		Description:  "C/C++ multimedia library for cross-platform game development",
		Homepage:     "https://liballeg.org/",
		URL:          "https://github.com/liballeg/allegro5/releases/download/5.2.9.1/allegro-5.2.9.1.tar.gz",
		Sha256:       "ee72b21114fea150045e575b6ab3db8ab7332deb7499d25d284e776c9a367a4b",
		Dependencies: []string{"cmake", "dumb", "flac", "freetype", "libogg", "libvorbis", "opusfile", "physfs", "theora", "webp", "opus", "jpeg-turbo", "libpng", "libx11", "libxcursor", "libxi", "libxinerama", "libxrandr", "libxscrnsaver", "mesa", "mesa-glu"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installallegro(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg allegroFormula) Installallegro() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "allegro-5.2.9.1.tar.gz"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "allegro-5.2.9.1.tar"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
