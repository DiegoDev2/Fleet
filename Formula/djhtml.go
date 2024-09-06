package main

// Djhtml - Django/Jinja template indenter
// Homepage: https://github.com/rtts/djhtml

import (
	"fmt"
	
	"os/exec"
)

func installDjhtml() {
	// Método 1: Descargar y extraer .tar.gz
	djhtml_tar_url := "https://files.pythonhosted.org/packages/a0/03/aac9bfb7c9b03604a2c4b0d474af22731ef41cb662fad07f956ae7bf0f6b/djhtml-3.0.6.tar.gz"
	djhtml_cmd_tar := exec.Command("curl", "-L", djhtml_tar_url, "-o", "package.tar.gz")
	err := djhtml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	djhtml_zip_url := "https://files.pythonhosted.org/packages/a0/03/aac9bfb7c9b03604a2c4b0d474af22731ef41cb662fad07f956ae7bf0f6b/djhtml-3.0.6.zip"
	djhtml_cmd_zip := exec.Command("curl", "-L", djhtml_zip_url, "-o", "package.zip")
	err = djhtml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	djhtml_bin_url := "https://files.pythonhosted.org/packages/a0/03/aac9bfb7c9b03604a2c4b0d474af22731ef41cb662fad07f956ae7bf0f6b/djhtml-3.0.6.bin"
	djhtml_cmd_bin := exec.Command("curl", "-L", djhtml_bin_url, "-o", "binary.bin")
	err = djhtml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	djhtml_src_url := "https://files.pythonhosted.org/packages/a0/03/aac9bfb7c9b03604a2c4b0d474af22731ef41cb662fad07f956ae7bf0f6b/djhtml-3.0.6.src.tar.gz"
	djhtml_cmd_src := exec.Command("curl", "-L", djhtml_src_url, "-o", "source.tar.gz")
	err = djhtml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	djhtml_cmd_direct := exec.Command("./binary")
	err = djhtml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
