package main

// PythonPsutil - Cross-platform lib for process and system monitoring in Python
// Homepage: https://github.com/giampaolo/psutil

import (
	"fmt"
	
	"os/exec"
)

func installPythonPsutil() {
	// Método 1: Descargar y extraer .tar.gz
	pythonpsutil_tar_url := "https://files.pythonhosted.org/packages/90/c7/6dc0a455d111f68ee43f27793971cf03fe29b6ef972042549db29eec39a2/psutil-5.9.8.tar.gz"
	pythonpsutil_cmd_tar := exec.Command("curl", "-L", pythonpsutil_tar_url, "-o", "package.tar.gz")
	err := pythonpsutil_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pythonpsutil_zip_url := "https://files.pythonhosted.org/packages/90/c7/6dc0a455d111f68ee43f27793971cf03fe29b6ef972042549db29eec39a2/psutil-5.9.8.zip"
	pythonpsutil_cmd_zip := exec.Command("curl", "-L", pythonpsutil_zip_url, "-o", "package.zip")
	err = pythonpsutil_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pythonpsutil_bin_url := "https://files.pythonhosted.org/packages/90/c7/6dc0a455d111f68ee43f27793971cf03fe29b6ef972042549db29eec39a2/psutil-5.9.8.bin"
	pythonpsutil_cmd_bin := exec.Command("curl", "-L", pythonpsutil_bin_url, "-o", "binary.bin")
	err = pythonpsutil_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pythonpsutil_src_url := "https://files.pythonhosted.org/packages/90/c7/6dc0a455d111f68ee43f27793971cf03fe29b6ef972042549db29eec39a2/psutil-5.9.8.src.tar.gz"
	pythonpsutil_cmd_src := exec.Command("curl", "-L", pythonpsutil_src_url, "-o", "source.tar.gz")
	err = pythonpsutil_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pythonpsutil_cmd_direct := exec.Command("./binary")
	err = pythonpsutil_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.11")
exec.Command("latte", "install", "python@3.11")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
