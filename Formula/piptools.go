package main

// PipTools - Locking and sync for Pip requirements files
// Homepage: https://pip-tools.readthedocs.io

import (
	"fmt"
	
	"os/exec"
)

func installPipTools() {
	// Método 1: Descargar y extraer .tar.gz
	piptools_tar_url := "https://files.pythonhosted.org/packages/1a/87/1ef453f10fb0772f43549686f924460cc0a2404b828b348f72c52cb2f5bf/pip-tools-7.4.1.tar.gz"
	piptools_cmd_tar := exec.Command("curl", "-L", piptools_tar_url, "-o", "package.tar.gz")
	err := piptools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	piptools_zip_url := "https://files.pythonhosted.org/packages/1a/87/1ef453f10fb0772f43549686f924460cc0a2404b828b348f72c52cb2f5bf/pip-tools-7.4.1.zip"
	piptools_cmd_zip := exec.Command("curl", "-L", piptools_zip_url, "-o", "package.zip")
	err = piptools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	piptools_bin_url := "https://files.pythonhosted.org/packages/1a/87/1ef453f10fb0772f43549686f924460cc0a2404b828b348f72c52cb2f5bf/pip-tools-7.4.1.bin"
	piptools_cmd_bin := exec.Command("curl", "-L", piptools_bin_url, "-o", "binary.bin")
	err = piptools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	piptools_src_url := "https://files.pythonhosted.org/packages/1a/87/1ef453f10fb0772f43549686f924460cc0a2404b828b348f72c52cb2f5bf/pip-tools-7.4.1.src.tar.gz"
	piptools_cmd_src := exec.Command("curl", "-L", piptools_src_url, "-o", "source.tar.gz")
	err = piptools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	piptools_cmd_direct := exec.Command("./binary")
	err = piptools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
