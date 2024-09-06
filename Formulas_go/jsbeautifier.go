package main

// Jsbeautifier - JavaScript unobfuscator and beautifier
// Homepage: https://beautifier.io

import (
	"fmt"
	
	"os/exec"
)

func installJsbeautifier() {
	// Método 1: Descargar y extraer .tar.gz
	jsbeautifier_tar_url := "https://files.pythonhosted.org/packages/69/3e/dd37e1a7223247e3ef94714abf572415b89c4e121c4af48e9e4c392e2ca0/jsbeautifier-1.15.1.tar.gz"
	jsbeautifier_cmd_tar := exec.Command("curl", "-L", jsbeautifier_tar_url, "-o", "package.tar.gz")
	err := jsbeautifier_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jsbeautifier_zip_url := "https://files.pythonhosted.org/packages/69/3e/dd37e1a7223247e3ef94714abf572415b89c4e121c4af48e9e4c392e2ca0/jsbeautifier-1.15.1.zip"
	jsbeautifier_cmd_zip := exec.Command("curl", "-L", jsbeautifier_zip_url, "-o", "package.zip")
	err = jsbeautifier_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jsbeautifier_bin_url := "https://files.pythonhosted.org/packages/69/3e/dd37e1a7223247e3ef94714abf572415b89c4e121c4af48e9e4c392e2ca0/jsbeautifier-1.15.1.bin"
	jsbeautifier_cmd_bin := exec.Command("curl", "-L", jsbeautifier_bin_url, "-o", "binary.bin")
	err = jsbeautifier_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jsbeautifier_src_url := "https://files.pythonhosted.org/packages/69/3e/dd37e1a7223247e3ef94714abf572415b89c4e121c4af48e9e4c392e2ca0/jsbeautifier-1.15.1.src.tar.gz"
	jsbeautifier_cmd_src := exec.Command("curl", "-L", jsbeautifier_src_url, "-o", "source.tar.gz")
	err = jsbeautifier_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jsbeautifier_cmd_direct := exec.Command("./binary")
	err = jsbeautifier_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
