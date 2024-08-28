package main

import (
	"fmt"
	"log"
	"os/exec"
)

// cryptographyFormula represents a formula in Go.
type cryptographyFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg cryptographyFormula) Print() {
	fmt.Printf("Name: cryptography\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := cryptographyFormula{
		Description:  "Cryptographic recipes and primitives for Python",
		Homepage:     "https://cryptography.io/en/latest/",
		URL:          "https://files.pythonhosted.org/packages/9d/f1/2cb8887cad0726a5e429cc9c58e30767f58d22c34d55b075d2f845d4a2a5/setuptools-rust-1.9.0.tar.gz",
		Sha256:       "704df0948f2e4cc60c2596ad6e840ea679f4f43e58ed4ad0c1857807240eab96",
		Dependencies: []string{"pkg-config", "python@3.11", "python@3.12", "rust", "cffi", "openssl@3"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installcryptography(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg cryptographyFormula) Installcryptography() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "setuptools-rust-1.9.0.tar.gz"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "setuptools-rust-1.9.0.tar"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
