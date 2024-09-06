package main

// Latex2html - LaTeX-to-HTML translator
// Homepage: https://www.latex2html.org

import (
	"fmt"
	
	"os/exec"
)

func installLatex2html() {
	// Método 1: Descargar y extraer .tar.gz
	latex2html_tar_url := "https://github.com/latex2html/latex2html/archive/refs/tags/v2024.tar.gz"
	latex2html_cmd_tar := exec.Command("curl", "-L", latex2html_tar_url, "-o", "package.tar.gz")
	err := latex2html_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	latex2html_zip_url := "https://github.com/latex2html/latex2html/archive/refs/tags/v2024.zip"
	latex2html_cmd_zip := exec.Command("curl", "-L", latex2html_zip_url, "-o", "package.zip")
	err = latex2html_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	latex2html_bin_url := "https://github.com/latex2html/latex2html/archive/refs/tags/v2024.bin"
	latex2html_cmd_bin := exec.Command("curl", "-L", latex2html_bin_url, "-o", "binary.bin")
	err = latex2html_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	latex2html_src_url := "https://github.com/latex2html/latex2html/archive/refs/tags/v2024.src.tar.gz"
	latex2html_cmd_src := exec.Command("curl", "-L", latex2html_src_url, "-o", "source.tar.gz")
	err = latex2html_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	latex2html_cmd_direct := exec.Command("./binary")
	err = latex2html_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ghostscript")
	exec.Command("latte", "install", "ghostscript").Run()
	fmt.Println("Instalando dependencia: netpbm")
	exec.Command("latte", "install", "netpbm").Run()
}
