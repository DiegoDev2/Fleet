package main

import (
	"fmt"
	"log"
	"os/exec"
)

// pymupdfFormula represents a formula in Go.
type pymupdfFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg pymupdfFormula) Print() {
	fmt.Printf("Name: pymupdf\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := pymupdfFormula{
		Description:  "Python bindings for the PDF toolkit and renderer MuPDF",
		Homepage:     "https://pymupdf.readthedocs.io/en/latest/",
		URL:          "https://files.pythonhosted.org/packages/0c/6c/1d3e88cd7b6a0f074ad6cec0dc32f9c023acd98b328eb23a183517e80e2b/PyMuPDFb-1.24.9.tar.gz",
		Sha256:       "5505f07b3dded6e791ab7d10d01f0687e913fc75edd23fdf2825a582b6651558",
		Dependencies: []string{"freetype", "python-setuptools", "swig", "mupdf", "python@3.12"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installpymupdf(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg pymupdfFormula) Installpymupdf() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "PyMuPDFb-1.24.9.tar.gz"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "PyMuPDFb-1.24.9.tar"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
