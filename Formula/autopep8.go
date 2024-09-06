package main

// Autopep8 - Automatically formats Python code to conform to the PEP 8 style guide
// Homepage: https://github.com/hhatto/autopep8

import (
	"fmt"
	
	"os/exec"
)

func installAutopep8() {
	// Método 1: Descargar y extraer .tar.gz
	autopep8_tar_url := "https://files.pythonhosted.org/packages/6c/52/65556a5f917a4b273fd1b705f98687a6bd721dbc45966f0f6687e90a18b0/autopep8-2.3.1.tar.gz"
	autopep8_cmd_tar := exec.Command("curl", "-L", autopep8_tar_url, "-o", "package.tar.gz")
	err := autopep8_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	autopep8_zip_url := "https://files.pythonhosted.org/packages/6c/52/65556a5f917a4b273fd1b705f98687a6bd721dbc45966f0f6687e90a18b0/autopep8-2.3.1.zip"
	autopep8_cmd_zip := exec.Command("curl", "-L", autopep8_zip_url, "-o", "package.zip")
	err = autopep8_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	autopep8_bin_url := "https://files.pythonhosted.org/packages/6c/52/65556a5f917a4b273fd1b705f98687a6bd721dbc45966f0f6687e90a18b0/autopep8-2.3.1.bin"
	autopep8_cmd_bin := exec.Command("curl", "-L", autopep8_bin_url, "-o", "binary.bin")
	err = autopep8_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	autopep8_src_url := "https://files.pythonhosted.org/packages/6c/52/65556a5f917a4b273fd1b705f98687a6bd721dbc45966f0f6687e90a18b0/autopep8-2.3.1.src.tar.gz"
	autopep8_cmd_src := exec.Command("curl", "-L", autopep8_src_url, "-o", "source.tar.gz")
	err = autopep8_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	autopep8_cmd_direct := exec.Command("./binary")
	err = autopep8_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
