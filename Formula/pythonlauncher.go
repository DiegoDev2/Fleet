package main

// PythonLauncher - Launch your Python interpreter the lazy/smart way
// Homepage: https://github.com/brettcannon/python-launcher

import (
	"fmt"
	
	"os/exec"
)

func installPythonLauncher() {
	// Método 1: Descargar y extraer .tar.gz
	pythonlauncher_tar_url := "https://github.com/brettcannon/python-launcher/archive/refs/tags/v1.0.1.tar.gz"
	pythonlauncher_cmd_tar := exec.Command("curl", "-L", pythonlauncher_tar_url, "-o", "package.tar.gz")
	err := pythonlauncher_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pythonlauncher_zip_url := "https://github.com/brettcannon/python-launcher/archive/refs/tags/v1.0.1.zip"
	pythonlauncher_cmd_zip := exec.Command("curl", "-L", pythonlauncher_zip_url, "-o", "package.zip")
	err = pythonlauncher_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pythonlauncher_bin_url := "https://github.com/brettcannon/python-launcher/archive/refs/tags/v1.0.1.bin"
	pythonlauncher_cmd_bin := exec.Command("curl", "-L", pythonlauncher_bin_url, "-o", "binary.bin")
	err = pythonlauncher_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pythonlauncher_src_url := "https://github.com/brettcannon/python-launcher/archive/refs/tags/v1.0.1.src.tar.gz"
	pythonlauncher_cmd_src := exec.Command("curl", "-L", pythonlauncher_src_url, "-o", "source.tar.gz")
	err = pythonlauncher_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pythonlauncher_cmd_direct := exec.Command("./binary")
	err = pythonlauncher_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
