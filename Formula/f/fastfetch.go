package main

import (
	"fmt"
	"log"
	"os/exec"
)

// fastfetchFormula represents a formula in Go.
type fastfetchFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg fastfetchFormula) Print() {
	fmt.Printf("Name: fastfetch\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := fastfetchFormula{
		Description:  "Like neofetch, but much faster because written mostly in C",
		Homepage:     "https://github.com/fastfetch-cli/fastfetch",
		URL:          "https://github.com/fastfetch-cli/fastfetch/archive/refs/tags/2.22.0.tar.gz",
		Sha256:       "2615c949d9c4f71b74fbd075f350bfac1fdeae1ae0f2b9a1511629e3c8824115",
		Dependencies: []string{"chafa", "cmake", "glib", "imagemagick", "pkg-config", "python@3.12", "vulkan-loader", "dbus", "ddcutil", "elfutils", "libdrm", "libx11", "libxcb", "libxrandr", "linux-headers@5.15", "mesa", "opencl-icd-loader", "pulseaudio", "rpm", "wayland"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installfastfetch(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg fastfetchFormula) Installfastfetch() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "2.22.0.tar.gz"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "2.22.0.tar"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
