package main

import (
	"fmt"
	"log"
	"os/exec"
)

// modsecurityFormula represents a formula in Go.
type modsecurityFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg modsecurityFormula) Print() {
	fmt.Printf("Name: modsecurity\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := modsecurityFormula{
		Description:  "Libmodsecurity is one component of the ModSecurity v3 project",
		Homepage:     "https://github.com/owasp-modsecurity/ModSecurity",
		URL:          "https://gitlab.archlinux.org/archlinux/packaging/packages/libmodsecurity/-/raw/5c78cfaaeb00c842731c52851341884c74bdc9b2/libxml-includes.patch",
		Sha256:       "7ee0adbe5b164ca512c49e51e30ffd41e29244156a695e619dcf1d0387e69aef",
		Dependencies: []string{"autoconf", "automake", "libtool", "pkg-config", "libmaxminddb", "lua", "pcre2", "yajl"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installmodsecurity(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg modsecurityFormula) Installmodsecurity() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "libxml-includes.patch"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "libxml-includes"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
