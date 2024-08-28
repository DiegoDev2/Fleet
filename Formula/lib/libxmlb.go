package main

import (
	"fmt"
	"log"
	"os/exec"
)

// libxmlbFormula represents a formula in Go.
type libxmlbFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg libxmlbFormula) Print() {
	fmt.Printf("Name: libxmlb\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := libxmlbFormula{
		Description:  "Library for querying compressed XML metadata",
		Homepage:     "https://github.com/hughsie/libxmlb",
		URL:          "https://github.com/hughsie/libxmlb/releases/download/0.3.19/libxmlb-0.3.19.tar.xz",
		Sha256:       "41146b757dcff6c8f462d6acc6f734a4e3352f70518534fbaa2f97fe28cfe204",
		Dependencies: []string{"gi-docgen", "gobject-introspection", "meson", "ninja", "pkg-config", "python@3.12", "vala", "glib", "xz", "zstd"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installlibxmlb(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg libxmlbFormula) Installlibxmlb() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "libxmlb-0.3.19.tar.xz"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "libxmlb-0.3.19.tar"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
