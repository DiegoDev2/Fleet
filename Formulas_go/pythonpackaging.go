package main

// PythonPackaging - Core utilities for Python packages
// Homepage: https://packaging.pypa.io/

import (
	"fmt"
	
	"os/exec"
)

func installPythonPackaging() {
	// Método 1: Descargar y extraer .tar.gz
	pythonpackaging_tar_url := "https://files.pythonhosted.org/packages/51/65/50db4dda066951078f0a96cf12f4b9ada6e4b811516bf0262c0f4f7064d4/packaging-24.1.tar.gz"
	pythonpackaging_cmd_tar := exec.Command("curl", "-L", pythonpackaging_tar_url, "-o", "package.tar.gz")
	err := pythonpackaging_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pythonpackaging_zip_url := "https://files.pythonhosted.org/packages/51/65/50db4dda066951078f0a96cf12f4b9ada6e4b811516bf0262c0f4f7064d4/packaging-24.1.zip"
	pythonpackaging_cmd_zip := exec.Command("curl", "-L", pythonpackaging_zip_url, "-o", "package.zip")
	err = pythonpackaging_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pythonpackaging_bin_url := "https://files.pythonhosted.org/packages/51/65/50db4dda066951078f0a96cf12f4b9ada6e4b811516bf0262c0f4f7064d4/packaging-24.1.bin"
	pythonpackaging_cmd_bin := exec.Command("curl", "-L", pythonpackaging_bin_url, "-o", "binary.bin")
	err = pythonpackaging_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pythonpackaging_src_url := "https://files.pythonhosted.org/packages/51/65/50db4dda066951078f0a96cf12f4b9ada6e4b811516bf0262c0f4f7064d4/packaging-24.1.src.tar.gz"
	pythonpackaging_cmd_src := exec.Command("curl", "-L", pythonpackaging_src_url, "-o", "source.tar.gz")
	err = pythonpackaging_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pythonpackaging_cmd_direct := exec.Command("./binary")
	err = pythonpackaging_cmd_direct.Run()
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
