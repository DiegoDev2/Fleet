package main

import (
	"fmt"
	"log"
	"os/exec"
)

// libtensorflowFormula represents a formula in Go.
type libtensorflowFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg libtensorflowFormula) Print() {
	fmt.Printf("Name: libtensorflow\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := libtensorflowFormula{
		Description:  "C interface for Google's OS library for Machine Intelligence",
		Homepage:     "https://www.tensorflow.org/",
		URL:          "https://github.com/tensorflow/tensorflow/commit/364b0cae088b7199a479899ef7bd99ddb9441728.patch?full_index=1",
		Sha256:       "86a225177dcee2965a1d2aa7cd5df3179365cf8ba637ba94d7fd61693f9a9fbb",
		Dependencies: []string{"bazelisk", "numpy", "python@3.12", "gnu-getopt"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installlibtensorflow(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg libtensorflowFormula) Installlibtensorflow() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "364b0cae088b7199a479899ef7bd99ddb9441728.patch?full_index=1"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "364b0cae088b7199a479899ef7bd99ddb9441728"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
