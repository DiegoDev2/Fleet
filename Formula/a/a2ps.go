package main

import (
	"fmt"
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
	fmt.Printf("Name: A2ps\\n")
	fmt.Printf("Description: %s\\Any-to-PostScript filter", pkg.Description)
	fmt.Printf("Homepage: %s\\https://www.gnu.org/software/a2ps/", pkg.Homepage)
	fmt.Printf("URL: %s\\https://ftp.gnu.org/gnu/a2ps/a2ps-4.15.6.tar.gz", pkg.URL)
	fmt.Printf("Sha256: %s\\87ff9d801cb11969181d5b8cf8b65e65e5b24bb0c76a1b825e8098f2906fbdf4", pkg.Sha256)
	fmt.Printf("Dependencies: %v\\n", pkg.Dependencies)
}

func main() {
	// Crear una instancia de A2psFormulaFormula
	pkg := A2psFormulaFormula{
		Description:  "Any-to-PostScript filter",
		Homepage:     "https://www.gnu.org/software/a2ps/",
		URL:          "https://ftp.gnu.org/gnu/a2ps/a2ps-4.15.6.tar.gz",
		Sha256:       "87ff9d801cb11969181d5b8cf8b65e65e5b24bb0c76a1b825e8098f2906fbdf4",
		Dependencies: []string{"gperf", "libpaper", "bdw-gc", "pkg-config"},
	}

	// Imprimir la información de la fórmula
	pkg.Print()
	InstallA2ps()
}

func InstallA2ps() {
	var cmd *exec.Cmd = exec.Command("curl", "-O", "https://ftp.gnu.org/gnu/a2ps/a2ps-4.15.6.tar.gz")
	cmd.Run()
}
