package main

import (
	"fmt"
	"log"
	"os/exec"
)

// imagemagickFormula represents a formula in Go.
type imagemagickFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg imagemagickFormula) Print() {
	fmt.Printf("Name: imagemagick\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := imagemagickFormula{
		Description:  "Tools and libraries to manipulate images in many formats",
		Homepage:     "https://imagemagick.org/index.php",
		URL:          "https://imagemagick.org/archive/",
		Sha256:       "bcb224a32ed54dc6f22ca4b09444de4df5ac3c9d98c4931e8f54a21d6ce71e86",
		Dependencies: []string{"pkg-config", "fontconfig", "freetype", "ghostscript", "jpeg-turbo", "jpeg-xl", "libheif", "liblqr", "libpng", "libraw", "libtiff", "libtool", "little-cms2", "openexr", "openjpeg", "webp", "xz", "gettext", "glib", "imath", "libomp", "libx11"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installimagemagick(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg imagemagickFormula) Installimagemagick() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := ""
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := ""
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
