package main

// Texi2html - Convert TeXinfo files to HTML
// Homepage: https://www.nongnu.org/texi2html/

import (
	"fmt"
	
	"os/exec"
)

func installTexi2html() {
	// Método 1: Descargar y extraer .tar.gz
	texi2html_tar_url := "https://download.savannah.gnu.org/releases/texi2html/texi2html-5.0.tar.gz"
	texi2html_cmd_tar := exec.Command("curl", "-L", texi2html_tar_url, "-o", "package.tar.gz")
	err := texi2html_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	texi2html_zip_url := "https://download.savannah.gnu.org/releases/texi2html/texi2html-5.0.zip"
	texi2html_cmd_zip := exec.Command("curl", "-L", texi2html_zip_url, "-o", "package.zip")
	err = texi2html_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	texi2html_bin_url := "https://download.savannah.gnu.org/releases/texi2html/texi2html-5.0.bin"
	texi2html_cmd_bin := exec.Command("curl", "-L", texi2html_bin_url, "-o", "binary.bin")
	err = texi2html_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	texi2html_src_url := "https://download.savannah.gnu.org/releases/texi2html/texi2html-5.0.src.tar.gz"
	texi2html_cmd_src := exec.Command("curl", "-L", texi2html_src_url, "-o", "source.tar.gz")
	err = texi2html_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	texi2html_cmd_direct := exec.Command("./binary")
	err = texi2html_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
