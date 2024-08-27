package main

import (
	"fmt"
	"log"
	"os/exec"
)

// A2psFormulaFormula representa una fórmula en Go.
type A2psFormulaFormula struct {
	Description  string
	Homepage     string
	URL          string
	Sha256       string
	Dependencies []string
}

func (pkg A2psFormulaFormula) Print() {
	fmt.Printf("Name: A2ps\n")
	fmt.Printf("Description: %s\n", pkg.Description)
	fmt.Printf("Homepage: %s\n", pkg.Homepage)
	fmt.Printf("URL: %s\n", pkg.URL)
	fmt.Printf("Sha256: %s\n", pkg.Sha256)
	fmt.Printf("Dependencies: %v\n", pkg.Dependencies)
}

func main() {
	// Crear una instancia de A2psFormulaFormula
	pkg := A2psFormulaFormula{
		Description:  "Any-to-PostScript filter",
		Homepage:     "https://www.gnu.org/software/a2ps/",
		URL:          "https://ftp.gnu.org/gnu/a2ps/a2ps-4.15.6.tar.gz",
		Sha256:       "87ff9d801cb11969181d5b8cf8b65e65e5b24bb0c76a1b825e8098f2906fbdf4",
		Dependencies: []string{"pkg-config", "bdw-gc", "libpaper", "gperf"},
	}

	pkg.Print()

	// Instalar la fórmula
	if err := pkg.InstallA2ps(); err != nil {
		log.Fatalf("Error during installation: %v", err)
	}

	fmt.Println("Installation completed successfully.")
}

// InstallA2ps descarga, configura, compila e instala A2ps
func (pkg A2psFormulaFormula) InstallA2ps() error {
	// Descargar el archivo fuente
	cmd := exec.Command("curl", "-O", pkg.URL)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to download: %v", err)
	}

	// Extraer el archivo tar.gz
	tarball := "a2ps-4.15.6.tar.gz"
	cmd = exec.Command("tar", "-xf", tarball)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to extract tarball: %v", err)
	}

	// Cambiar al directorio extraído
	sourceDir := "a2ps-4.15.6"
	cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && PKG_CONFIG_PATH=/usr/local/lib/pkgconfig ./configure --sysconfdir=/etc --with-lispdir=/usr/share/emacs/site-lisp --with-packager=Homebrew --with-packager-version=4.15.6 --with-packager-bug-reports=https://github.com/Homebrew/homebrew-core/issues && make install", sourceDir))
	cmd.Stdout = log.Writer() // Redirige la salida del comando a los logs
	cmd.Stderr = log.Writer() // Redirige los errores del comando a los logs

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to configure and install: %v", err)
	}

	return nil
}
