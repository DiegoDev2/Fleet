package main

import (
	"fmt"
	"log"
	"os/exec"
)

// pinocchioFormula represents a formula in Go.
type pinocchioFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg pinocchioFormula) Print() {
	fmt.Printf("Name: pinocchio\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	pkg := pinocchioFormula{
		Description:  "Efficient and fast C++ library implementing Rigid Body Dynamics algorithms",
		Homepage:     "https://stack-of-tasks.github.io/pinocchio",
		URL:          "https://github.com/stack-of-tasks/pinocchio/releases/download/v3.2.0/pinocchio-3.2.0.tar.gz",
		Sha256:       "a09640f68ac89dba6af85f0a056058747643772a95388ab7cde28d9097db5332",
		Dependencies: []string{"cmake", "doxygen", "pkg-config", "boost", "boost-python3", "console_bridge", "eigen", "eigenpy", "hpp-fcl", "python@3.12", "urdfdom", "octomap"},
	}

	pkg.Print()

	// Instalar dependencias
	for _, dep := range pkg.Dependencies {
		cmd := exec.Command("brew", "install", dep)
		if err := cmd.Run(); err != nil {
			log.Fatalf("Error installing dependency %s: %v", dep, err)
		}
	}

	if err := pkg.Installpinocchio(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

func (pkg pinocchioFormula) Installpinocchio() error {
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	tarball := "pinocchio-3.2.0.tar.gz"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	sourceDir := "pinocchio-3.2.0.tar"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
