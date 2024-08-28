package main

import (
	"fmt"
	"log"
	"os/exec"
)

// openrct2Formula represents a formula in Go.
type openrct2Formula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg openrct2Formula) Print() {
	fmt.Printf("Name: openrct2\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := openrct2Formula{
		Description:  "Open source re-implementation of RollerCoaster Tycoon 2",
		Homepage:     "https://openrct2.io/",
		URL:          "https://github.com/OpenRCT2/objects/releases/download/v1.4.7/objects.zip",
		Sha256:       "003c7d8a68a2461ac27204a361ffe6eab66e3aff262755b9830c97ce36d6fb65",
		Dependencies: []string{"cmake", "nlohmann-json", "pkg-config", "duktape", "flac", "freetype", "icu4c", "libogg", "libpng", "libvorbis", "libzip", "openssl@3", "sdl2", "speexdsp", "curl", "fontconfig", "mesa"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installopenrct2(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg openrct2Formula) Installopenrct2() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "objects.zip"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "objects"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
