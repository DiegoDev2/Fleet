package main

import (
	"fmt"
	"log"
	"os/exec"
)

// ocpFormula represents a formula in Go.
type ocpFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg ocpFormula) Print() {
	fmt.Printf("Name: ocp\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := ocpFormula{
		Description:  "UNIX port of the Open Cubic Player",
		Homepage:     "https://stian.cubic.org/project-ocp.php",
		URL:          "https://ftp.gnu.org/gnu/unifont/unifont-15.0.06/unifont-15.0.06.tar.gz",
		Sha256:       "36668eb1326d22e1466b94b3929beeafd10b9838bf3d41f4e5e3b52406ae69f1",
		Dependencies: []string{"pkg-config", "xa", "ancient", "cjson", "flac", "freetype", "game-music-emu", "jpeg-turbo", "libdiscid", "libpng", "libvorbis", "mad", "sdl2", "libogg", "util-linux", "alsa-lib"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installocp(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg ocpFormula) Installocp() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "unifont-15.0.06.tar.gz"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "unifont-15.0.06.tar"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
