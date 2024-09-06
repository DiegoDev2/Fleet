package main

// PythonDateutil - Useful extensions to the standard Python datetime features
// Homepage: https://github.com/dateutil/dateutil

import (
	"fmt"
	
	"os/exec"
)

func installPythonDateutil() {
	// Método 1: Descargar y extraer .tar.gz
	pythondateutil_tar_url := "https://files.pythonhosted.org/packages/66/c0/0c8b6ad9f17a802ee498c46e004a0eb49bc148f2fd230864601a86dcf6db/python-dateutil-2.9.0.post0.tar.gz"
	pythondateutil_cmd_tar := exec.Command("curl", "-L", pythondateutil_tar_url, "-o", "package.tar.gz")
	err := pythondateutil_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pythondateutil_zip_url := "https://files.pythonhosted.org/packages/66/c0/0c8b6ad9f17a802ee498c46e004a0eb49bc148f2fd230864601a86dcf6db/python-dateutil-2.9.0.post0.zip"
	pythondateutil_cmd_zip := exec.Command("curl", "-L", pythondateutil_zip_url, "-o", "package.zip")
	err = pythondateutil_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pythondateutil_bin_url := "https://files.pythonhosted.org/packages/66/c0/0c8b6ad9f17a802ee498c46e004a0eb49bc148f2fd230864601a86dcf6db/python-dateutil-2.9.0.post0.bin"
	pythondateutil_cmd_bin := exec.Command("curl", "-L", pythondateutil_bin_url, "-o", "binary.bin")
	err = pythondateutil_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pythondateutil_src_url := "https://files.pythonhosted.org/packages/66/c0/0c8b6ad9f17a802ee498c46e004a0eb49bc148f2fd230864601a86dcf6db/python-dateutil-2.9.0.post0.src.tar.gz"
	pythondateutil_cmd_src := exec.Command("curl", "-L", pythondateutil_src_url, "-o", "source.tar.gz")
	err = pythondateutil_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pythondateutil_cmd_direct := exec.Command("./binary")
	err = pythondateutil_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.11")
	exec.Command("latte", "install", "python@3.11").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: six")
	exec.Command("latte", "install", "six").Run()
}
