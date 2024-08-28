package main

import (
	"fmt"
	"log"
	"os/exec"
)

// offlineimapFormula represents a formula in Go.
type offlineimapFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg offlineimapFormula) Print() {
	fmt.Printf("Name: offlineimap\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := offlineimapFormula{
		Description:  "Synchronizes emails between two repositories",
		Homepage:     "https://github.com/OfflineIMAP/offlineimap3",
		URL:          "https://github.com/OfflineIMAP/offlineimap3/commit/4601f50d98cffcb182fddb04f8a78c795004bc73.patch?full_index=1",
		Sha256:       "a38595f54fa70d3cdb44aec2f858c256265421171a8ec331a34cbe6041072954",
		Dependencies: []string{"certifi", "python@3.12"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installofflineimap(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg offlineimapFormula) Installofflineimap() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "4601f50d98cffcb182fddb04f8a78c795004bc73.patch?full_index=1"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "4601f50d98cffcb182fddb04f8a78c795004bc73"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
