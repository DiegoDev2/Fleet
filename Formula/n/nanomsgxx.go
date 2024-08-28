package main

import (
	"fmt"
	"log"
	"os/exec"
)

// nanomsgxxFormula represents a formula in Go.
type nanomsgxxFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg nanomsgxxFormula) Print() {
	fmt.Printf("Name: nanomsgxx\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := nanomsgxxFormula{
		Description:  "Nanomsg binding for C++11",
		Homepage:     "https://achille-roussel.github.io/nanomsgxx/doc/nanomsgxx.7.html",
		URL:          "https://github.com/achille-roussel/nanomsgxx/commit/08c6d8882e40d0279e58325d641a7abead51ca07.patch?full_index=1",
		Sha256:       "fa27cad45e6216dfcf8a26125c0ff9db65e315653c16366a82e5b39d6e4de415",
		Dependencies: []string{"pkg-config", "nanomsg"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installnanomsgxx(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg nanomsgxxFormula) Installnanomsgxx() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "08c6d8882e40d0279e58325d641a7abead51ca07.patch?full_index=1"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "08c6d8882e40d0279e58325d641a7abead51ca07"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
