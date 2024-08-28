package main

import (
	"fmt"
	"log"
	"os/exec"
)

// dnsvizFormula represents a formula in Go.
type dnsvizFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg dnsvizFormula) Print() {
	fmt.Printf("Name: dnsviz\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := dnsvizFormula{
		Description:  "Tools for analyzing and visualizing DNS and DNSSEC behavior",
		Homepage:     "https://github.com/dnsviz/dnsviz/",
		URL:          "https://files.pythonhosted.org/packages/8c/41/7b9a22df38bb7884012b34f2986d765691dbe41bf5e7af881dfd09f8145f/pygraphviz-1.13.tar.gz",
		Sha256:       "6ad8aa2f26768830a5a1cfc8a14f022d13df170a8f6fdfd68fd1aa1267000964",
		Dependencies: []string{"bind", "pkg-config", "swig", "json-c", "graphviz", "openssl@3", "python@3.12", "swig"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installdnsviz(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg dnsvizFormula) Installdnsviz() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "pygraphviz-1.13.tar.gz"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "pygraphviz-1.13.tar"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
