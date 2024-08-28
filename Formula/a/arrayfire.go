package main

import (
	"fmt"
	"log"
	"os/exec"
)

// arrayfireFormula represents a formula in Go.
type arrayfireFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg arrayfireFormula) Print() {
	fmt.Printf("Name: arrayfire\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := arrayfireFormula{
		Description:  "General purpose GPU library",
		Homepage:     "https://arrayfire.com",
		URL:          "https://github.com/arrayfire/arrayfire/releases/download/v3.9.0/arrayfire-full-3.9.0.tar.bz2",
		Sha256:       "e0cdfa9839ea984d846c2fd7c9df45c337ae7159d07a590256896b1934d9670e",
		Dependencies: []string{"boost@1.85", "cmake", "doxygen", "fftw", "fmt", "freeimage", "openblas", "spdlog", "opencl-headers", "opencl-icd-loader", "pocl"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installarrayfire(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg arrayfireFormula) Installarrayfire() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "arrayfire-full-3.9.0.tar.bz2"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "arrayfire-full-3.9.0.tar"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
