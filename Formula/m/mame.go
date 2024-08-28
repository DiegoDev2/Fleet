package main

import (
	"fmt"
	"log"
	"os/exec"
)

// mameFormula represents a formula in Go.
type mameFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg mameFormula) Print() {
	fmt.Printf("Name: mame\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := mameFormula{
		Description:  "Multiple Arcade Machine Emulator",
		Homepage:     "https://mamedev.org/",
		URL:          "https://github.com/mamedev/mame/archive/refs/tags/mame0268.tar.gz",
		Sha256:       "53f4beccca1de36c1502d2d40c7f811acebb0e1f7450fa4638a6f6a8aa988a69",
		Dependencies: []string{"asio", "glm", "pkg-config", "rapidjson", "sphinx-doc", "flac", "jpeg-turbo", "portaudio", "portmidi", "pugixml", "sdl2", "sqlite", "utf8proc", "zstd", "pulseaudio", "qt@5", "sdl2_ttf"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installmame(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg mameFormula) Installmame() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "mame0268.tar.gz"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "mame0268.tar"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
