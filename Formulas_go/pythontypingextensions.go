package main

// PythonTypingExtensions - Backported and experimental type hints for Python
// Homepage: https://github.com/python/typing_extensions

import (
	"fmt"
	
	"os/exec"
)

func installPythonTypingExtensions() {
	// Método 1: Descargar y extraer .tar.gz
	pythontypingextensions_tar_url := "https://files.pythonhosted.org/packages/16/3a/0d26ce356c7465a19c9ea8814b960f8a36c3b0d07c323176620b7b483e44/typing_extensions-4.10.0.tar.gz"
	pythontypingextensions_cmd_tar := exec.Command("curl", "-L", pythontypingextensions_tar_url, "-o", "package.tar.gz")
	err := pythontypingextensions_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pythontypingextensions_zip_url := "https://files.pythonhosted.org/packages/16/3a/0d26ce356c7465a19c9ea8814b960f8a36c3b0d07c323176620b7b483e44/typing_extensions-4.10.0.zip"
	pythontypingextensions_cmd_zip := exec.Command("curl", "-L", pythontypingextensions_zip_url, "-o", "package.zip")
	err = pythontypingextensions_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pythontypingextensions_bin_url := "https://files.pythonhosted.org/packages/16/3a/0d26ce356c7465a19c9ea8814b960f8a36c3b0d07c323176620b7b483e44/typing_extensions-4.10.0.bin"
	pythontypingextensions_cmd_bin := exec.Command("curl", "-L", pythontypingextensions_bin_url, "-o", "binary.bin")
	err = pythontypingextensions_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pythontypingextensions_src_url := "https://files.pythonhosted.org/packages/16/3a/0d26ce356c7465a19c9ea8814b960f8a36c3b0d07c323176620b7b483e44/typing_extensions-4.10.0.src.tar.gz"
	pythontypingextensions_cmd_src := exec.Command("curl", "-L", pythontypingextensions_src_url, "-o", "source.tar.gz")
	err = pythontypingextensions_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pythontypingextensions_cmd_direct := exec.Command("./binary")
	err = pythontypingextensions_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.11")
exec.Command("latte", "install", "python@3.11")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: mypy")
exec.Command("latte", "install", "mypy")
}
