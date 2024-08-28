package main

import (
	"fmt"
	"log"
	"os/exec"
)

// gxmlFormula represents a formula in Go.
type gxmlFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg gxmlFormula) Print() {
	fmt.Printf("Name: gxml\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := gxmlFormula{
		Description:  "GObject-based XML DOM API",
		Homepage:     "https://wiki.gnome.org/GXml",
		URL:          "https://gitlab.gnome.org/GNOME/gxml/-/commit/6551103abd5143e51814ec1dce9b36bb9a46e09f.diff",
		Sha256:       "b87f585ab782b2ff4f024c45c9a90791c2023e3703756f2eb799591e7978e640",
		Dependencies: []string{"gobject-introspection", "meson", "ninja", "pkg-config", "vala", "glib", "libgee", "libxml2", "gettext"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installgxml(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg gxmlFormula) Installgxml() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "6551103abd5143e51814ec1dce9b36bb9a46e09f.diff"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "6551103abd5143e51814ec1dce9b36bb9a46e09f"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
