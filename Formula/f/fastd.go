package main

import (
	"fmt"
	"log"
	"os/exec"
)

// fastdFormula represents a formula in Go.
type fastdFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg fastdFormula) Print() {
	fmt.Printf("Name: fastd\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := fastdFormula{
		Description:  "Fast and Secure Tunnelling Daemon",
		Homepage:     "https://github.com/neocturne/fastd",
		URL:          "https://github.com/neocturne/fastd/commit/89abc48e60e182f8d57e924df16acf33c6670a9b.patch?full_index=1",
		Sha256:       "7bcac7dc288961a34830ef0552e1f9985f1b818aa37978b281f542a26fb059b9",
		Dependencies: []string{"bison", "cmake", "meson", "ninja", "pkg-config", "json-c", "libsodium", "libuecc", "openssl@3", "libcap", "libmnl"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installfastd(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg fastdFormula) Installfastd() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "89abc48e60e182f8d57e924df16acf33c6670a9b.patch?full_index=1"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "89abc48e60e182f8d57e924df16acf33c6670a9b"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
