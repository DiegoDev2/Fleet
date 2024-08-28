package main

import (
	"fmt"
	"log"
	"os/exec"
)

// mat2Formula represents a formula in Go.
type mat2Formula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg mat2Formula) Print() {
	fmt.Printf("Name: mat2\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := mat2Formula{
		Description:  "Metadata anonymization toolkit",
		Homepage:     "https://0xacab.org/jvoisin/mat2",
		URL:          "https://0xacab.org/jvoisin/mat2/-/commit/406924bb6164384fe0a8a8f3dc8dfe7d15577cfc.diff",
		Sha256:       "4c1c57ca8fe1eabea41d66f3ef9bd4eb2bac8ac181fceeefece4b92b5be9658d",
		Dependencies: []string{"exiftool", "ffmpeg", "gdk-pixbuf", "librsvg", "poppler", "py3cairo", "pygobject3", "python@3.12"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installmat2(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg mat2Formula) Installmat2() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "406924bb6164384fe0a8a8f3dc8dfe7d15577cfc.diff"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "406924bb6164384fe0a8a8f3dc8dfe7d15577cfc"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
