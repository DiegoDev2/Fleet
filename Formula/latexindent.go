package main

// Latexindent - Add indentation to LaTeX files
// Homepage: https://latexindentpl.readthedocs.io

import (
	"fmt"
	
	"os/exec"
)

func installLatexindent() {
	// Método 1: Descargar y extraer .tar.gz
	latexindent_tar_url := "https://github.com/cmhughes/latexindent.pl/archive/refs/tags/V3.24.4.tar.gz"
	latexindent_cmd_tar := exec.Command("curl", "-L", latexindent_tar_url, "-o", "package.tar.gz")
	err := latexindent_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	latexindent_zip_url := "https://github.com/cmhughes/latexindent.pl/archive/refs/tags/V3.24.4.zip"
	latexindent_cmd_zip := exec.Command("curl", "-L", latexindent_zip_url, "-o", "package.zip")
	err = latexindent_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	latexindent_bin_url := "https://github.com/cmhughes/latexindent.pl/archive/refs/tags/V3.24.4.bin"
	latexindent_cmd_bin := exec.Command("curl", "-L", latexindent_bin_url, "-o", "binary.bin")
	err = latexindent_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	latexindent_src_url := "https://github.com/cmhughes/latexindent.pl/archive/refs/tags/V3.24.4.src.tar.gz"
	latexindent_cmd_src := exec.Command("curl", "-L", latexindent_src_url, "-o", "source.tar.gz")
	err = latexindent_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	latexindent_cmd_direct := exec.Command("./binary")
	err = latexindent_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: perl")
	exec.Command("latte", "install", "perl").Run()
}
