package main

import (
	"fmt"
	"log"
	"os/exec"
)

// notcursesFormula represents a formula in Go.
type notcursesFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg notcursesFormula) Print() {
	fmt.Printf("Name: notcurses\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := notcursesFormula{
		Description:  "Blingful character graphics/TUI library",
		Homepage:     "https://nick-black.com/dankwiki/index.php/Notcurses",
		URL:          "https://github.com/dankamongmen/notcurses/commit/441d66a063c7fc86436ed7ff73984050434c9142.patch?full_index=1",
		Sha256:       "aee69211bf5280bb773360a0f206e79f825ae86dbb7e05117d69acfa12917c13",
		Dependencies: []string{"cmake", "doctest", "pandoc", "pkg-config", "ffmpeg", "libdeflate", "libunistring", "ncurses"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installnotcurses(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg notcursesFormula) Installnotcurses() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "441d66a063c7fc86436ed7ff73984050434c9142.patch?full_index=1"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "441d66a063c7fc86436ed7ff73984050434c9142"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
