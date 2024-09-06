package main

// Latexml - LaTeX to XML/HTML/MathML Converter
// Homepage: https://dlmf.nist.gov/LaTeXML/

import (
	"fmt"
	
	"os/exec"
)

func installLatexml() {
	// Método 1: Descargar y extraer .tar.gz
	latexml_tar_url := "https://dlmf.nist.gov/LaTeXML/releases/LaTeXML-0.8.8.tar.gz"
	latexml_cmd_tar := exec.Command("curl", "-L", latexml_tar_url, "-o", "package.tar.gz")
	err := latexml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	latexml_zip_url := "https://dlmf.nist.gov/LaTeXML/releases/LaTeXML-0.8.8.zip"
	latexml_cmd_zip := exec.Command("curl", "-L", latexml_zip_url, "-o", "package.zip")
	err = latexml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	latexml_bin_url := "https://dlmf.nist.gov/LaTeXML/releases/LaTeXML-0.8.8.bin"
	latexml_cmd_bin := exec.Command("curl", "-L", latexml_bin_url, "-o", "binary.bin")
	err = latexml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	latexml_src_url := "https://dlmf.nist.gov/LaTeXML/releases/LaTeXML-0.8.8.src.tar.gz"
	latexml_cmd_src := exec.Command("curl", "-L", latexml_src_url, "-o", "source.tar.gz")
	err = latexml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	latexml_cmd_direct := exec.Command("./binary")
	err = latexml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: perl")
	exec.Command("latte", "install", "perl").Run()
}
