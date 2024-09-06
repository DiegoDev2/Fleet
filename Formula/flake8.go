package main

// Flake8 - Lint your Python code for style and logical errors
// Homepage: https://flake8.pycqa.org/

import (
	"fmt"
	
	"os/exec"
)

func installFlake8() {
	// Método 1: Descargar y extraer .tar.gz
	flake8_tar_url := "https://files.pythonhosted.org/packages/37/72/e8d66150c4fcace3c0a450466aa3480506ba2cae7b61e100a2613afc3907/flake8-7.1.1.tar.gz"
	flake8_cmd_tar := exec.Command("curl", "-L", flake8_tar_url, "-o", "package.tar.gz")
	err := flake8_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	flake8_zip_url := "https://files.pythonhosted.org/packages/37/72/e8d66150c4fcace3c0a450466aa3480506ba2cae7b61e100a2613afc3907/flake8-7.1.1.zip"
	flake8_cmd_zip := exec.Command("curl", "-L", flake8_zip_url, "-o", "package.zip")
	err = flake8_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	flake8_bin_url := "https://files.pythonhosted.org/packages/37/72/e8d66150c4fcace3c0a450466aa3480506ba2cae7b61e100a2613afc3907/flake8-7.1.1.bin"
	flake8_cmd_bin := exec.Command("curl", "-L", flake8_bin_url, "-o", "binary.bin")
	err = flake8_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	flake8_src_url := "https://files.pythonhosted.org/packages/37/72/e8d66150c4fcace3c0a450466aa3480506ba2cae7b61e100a2613afc3907/flake8-7.1.1.src.tar.gz"
	flake8_cmd_src := exec.Command("curl", "-L", flake8_src_url, "-o", "source.tar.gz")
	err = flake8_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	flake8_cmd_direct := exec.Command("./binary")
	err = flake8_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
