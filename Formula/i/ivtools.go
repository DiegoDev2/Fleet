package main

import (
	"fmt"
	"log"
	"os/exec"
)

// ivtoolsFormula represents a formula in Go.
type ivtoolsFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg ivtoolsFormula) Print() {
	fmt.Printf("Name: ivtools\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := ivtoolsFormula{
		Description:  "X11 vector graphic servers",
		Homepage:     "https://github.com/vectaport/ivtools",
		URL:          "https://github.com/vectaport/ivtools/commit/7ce87b9159e720cf1990b6fef10ba7a8b664bcda.patch?full_index=1",
		Sha256:       "0c1e722b574df66e3c6d4114b066db99691764a04f92c77af2b7adaabde3782c",
		Dependencies: []string{"ace", "libx11", "libxext"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installivtools(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg ivtoolsFormula) Installivtools() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "7ce87b9159e720cf1990b6fef10ba7a8b664bcda.patch?full_index=1"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "7ce87b9159e720cf1990b6fef10ba7a8b664bcda"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
