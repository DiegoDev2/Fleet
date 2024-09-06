package main

// Bibtex2html - BibTeX to HTML converter
// Homepage: https://usr.lmf.cnrs.fr/~jcf/bibtex2html/index.en.html

import (
	"fmt"
	
	"os/exec"
)

func installBibtex2html() {
	// Método 1: Descargar y extraer .tar.gz
	bibtex2html_tar_url := "https://usr.lmf.cnrs.fr/~jcf/ftp/bibtex2html/bibtex2html-1.99.tar.gz"
	bibtex2html_cmd_tar := exec.Command("curl", "-L", bibtex2html_tar_url, "-o", "package.tar.gz")
	err := bibtex2html_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bibtex2html_zip_url := "https://usr.lmf.cnrs.fr/~jcf/ftp/bibtex2html/bibtex2html-1.99.zip"
	bibtex2html_cmd_zip := exec.Command("curl", "-L", bibtex2html_zip_url, "-o", "package.zip")
	err = bibtex2html_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bibtex2html_bin_url := "https://usr.lmf.cnrs.fr/~jcf/ftp/bibtex2html/bibtex2html-1.99.bin"
	bibtex2html_cmd_bin := exec.Command("curl", "-L", bibtex2html_bin_url, "-o", "binary.bin")
	err = bibtex2html_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bibtex2html_src_url := "https://usr.lmf.cnrs.fr/~jcf/ftp/bibtex2html/bibtex2html-1.99.src.tar.gz"
	bibtex2html_cmd_src := exec.Command("curl", "-L", bibtex2html_src_url, "-o", "source.tar.gz")
	err = bibtex2html_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bibtex2html_cmd_direct := exec.Command("./binary")
	err = bibtex2html_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: ocaml")
	exec.Command("latte", "install", "ocaml").Run()
}
