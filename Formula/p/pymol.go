package main

import (
	"fmt"
	"log"
	"os/exec"
)

// pymolFormula represents a formula in Go.
type pymolFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg pymolFormula) Print() {
	fmt.Printf("Name: pymol\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := pymolFormula{
		Description:  "Molecular visualization system",
		Homepage:     "https://pymol.org/",
		URL:          "https://github.com/schrodinger/pymol-open-source/commit/4d81b4a8537421e9a1c4647934d1a16e24bc51dd.patch?full_index=1",
		Sha256:       "ee5895ecd3bf731fc1ad714cc6cea17cb5dbb81cd4dab62e77554219fe7ae1ec",
		Dependencies: []string{"cmake", "glm", "msgpack-cxx", "sip", "freetype", "glew", "libpng", "netcdf", "numpy", "pyqt@5", "python@3.12", "freeglut", "mesa"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installpymol(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg pymolFormula) Installpymol() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "4d81b4a8537421e9a1c4647934d1a16e24bc51dd.patch?full_index=1"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "4d81b4a8537421e9a1c4647934d1a16e24bc51dd"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
