package main

// Doc8 - Style checker for Sphinx documentation
// Homepage: https://github.com/PyCQA/doc8

import (
	"fmt"
	
	"os/exec"
)

func installDoc8() {
	// Método 1: Descargar y extraer .tar.gz
	doc8_tar_url := "https://files.pythonhosted.org/packages/11/28/b0a576233730b756ca1ebb422bc6199a761b826b86e93e5196dfa85331ea/doc8-1.1.2.tar.gz"
	doc8_cmd_tar := exec.Command("curl", "-L", doc8_tar_url, "-o", "package.tar.gz")
	err := doc8_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	doc8_zip_url := "https://files.pythonhosted.org/packages/11/28/b0a576233730b756ca1ebb422bc6199a761b826b86e93e5196dfa85331ea/doc8-1.1.2.zip"
	doc8_cmd_zip := exec.Command("curl", "-L", doc8_zip_url, "-o", "package.zip")
	err = doc8_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	doc8_bin_url := "https://files.pythonhosted.org/packages/11/28/b0a576233730b756ca1ebb422bc6199a761b826b86e93e5196dfa85331ea/doc8-1.1.2.bin"
	doc8_cmd_bin := exec.Command("curl", "-L", doc8_bin_url, "-o", "binary.bin")
	err = doc8_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	doc8_src_url := "https://files.pythonhosted.org/packages/11/28/b0a576233730b756ca1ebb422bc6199a761b826b86e93e5196dfa85331ea/doc8-1.1.2.src.tar.gz"
	doc8_cmd_src := exec.Command("curl", "-L", doc8_src_url, "-o", "source.tar.gz")
	err = doc8_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	doc8_cmd_direct := exec.Command("./binary")
	err = doc8_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
