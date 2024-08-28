package main

import (
	"fmt"
	"log"
	"os/exec"
)

// groestlcoinFormula represents a formula in Go.
type groestlcoinFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg groestlcoinFormula) Print() {
	fmt.Printf("Name: groestlcoin\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := groestlcoinFormula{
		Description:  "Decentralized, peer to peer payment network",
		Homepage:     "https://groestlcoin.org/groestlcoin-core-wallet/",
		URL:          "https://github.com/Groestlcoin/groestlcoin/releases/download/v27.0/groestlcoin-27.0.tar.gz",
		Sha256:       "3a07a117acf2e654c599b1b4d3916d42e0c3d9cdb7b13da8f184a3fa4123ea97",
		Dependencies: []string{"autoconf", "automake", "libtool", "pkg-config", "berkeley-db@5", "boost", "libevent", "miniupnpc", "zeromq", "util-linux"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installgroestlcoin(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg groestlcoinFormula) Installgroestlcoin() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "groestlcoin-27.0.tar.gz"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "groestlcoin-27.0.tar"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
