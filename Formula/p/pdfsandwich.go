package main

import (
	"fmt"
	"log"
	"os/exec"
)

// pdfsandwichFormula represents a formula in Go.
type pdfsandwichFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg pdfsandwichFormula) Print() {
	fmt.Printf("Name: pdfsandwich\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := pdfsandwichFormula{
		Description:  "Generate sandwich OCR PDFs from scanned file",
		Homepage:     "http://www.tobias-elze.de/pdfsandwich/",
		URL:          "https://downloads.sourceforge.net/project/pdfsandwich/pdfsandwich%200.1.7/pdfsandwich-0.1.7.tar.bz2",
		Sha256:       "6f77d1aa373059a0b32d879bc45f075527a28c6a0a2068b8d38d634dfd2d7d60",
		Dependencies: []string{"gawk", "ocaml", "exact-image", "ghostscript", "imagemagick", "poppler", "tesseract", "unpaper"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installpdfsandwich(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg pdfsandwichFormula) Installpdfsandwich() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "pdfsandwich-0.1.7.tar.bz2"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "pdfsandwich-0.1.7.tar"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
