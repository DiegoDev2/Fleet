package main

import (
	"fmt"
	"log"
	"os/exec"
)

// cherrytreeFormula represents a formula in Go.
type cherrytreeFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg cherrytreeFormula) Print() {
	fmt.Printf("Name: cherrytree\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := cherrytreeFormula{
		Description:  "Hierarchical note taking application featuring rich text and syntax highlighting",
		Homepage:     "https://www.giuspen.com/cherrytree/",
		URL:          "https://github.com/giuspen/cherrytree/commit/fc1d7499067b9db9841175b5a2d6934dc65e4522.patch?full_index=1",
		Sha256:       "8ddac9e4b25533b4ec3dd3496783454a282cf8de4d07b6a1527611dccd5ecefa",
		Dependencies: []string{"cmake", "gettext", "pkg-config", "adwaita-icon-theme", "atkmm@2.28", "cairo", "cairomm@1.14", "fmt", "fribidi", "glib", "glibmm@2.66", "gspell", "gtk+3", "gtkmm3", "gtksourceview3", "gtksourceviewmm3", "libsigc++@2", "libxml++", "pango", "pangomm@2.46", "spdlog", "sqlite", "uchardet", "vte3", "at-spi2-core", "enchant", "gdk-pixbuf", "gettext", "harfbuzz"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installcherrytree(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg cherrytreeFormula) Installcherrytree() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "fc1d7499067b9db9841175b5a2d6934dc65e4522.patch?full_index=1"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "fc1d7499067b9db9841175b5a2d6934dc65e4522"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
