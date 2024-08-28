package main

import (
	"fmt"
	"log"
	"os/exec"
)

// grpcFormula represents a formula in Go.
type grpcFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg grpcFormula) Print() {
	fmt.Printf("Name: grpc\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := grpcFormula{
		Description:  "Next generation open source RPC library and framework",
		Homepage:     "https://grpc.io/",
		URL:          "https://github.com/grpc/grpc/commit/98a96c5068da14ed29d70ca23818b5f408a2e7b4.patch?full_index=1",
		Sha256:       "a48893c79556662762c551476b92a7eca91fc350a44ad22f2fd51576e076376b",
		Dependencies: []string{"autoconf", "automake", "cmake", "libtool", "pkg-config", "abseil", "c-ares", "openssl@3", "protobuf", "re2", "llvm"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installgrpc(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg grpcFormula) Installgrpc() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "98a96c5068da14ed29d70ca23818b5f408a2e7b4.patch?full_index=1"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "98a96c5068da14ed29d70ca23818b5f408a2e7b4"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
