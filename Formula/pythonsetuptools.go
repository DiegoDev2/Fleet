package main

// PythonSetuptools - Easily download, build, install, upgrade, and uninstall Python packages
// Homepage: https://setuptools.pypa.io/

import (
	"fmt"
	
	"os/exec"
)

func installPythonSetuptools() {
	// Método 1: Descargar y extraer .tar.gz
	pythonsetuptools_tar_url := "https://files.pythonhosted.org/packages/3e/2c/f0a538a2f91ce633a78daaeb34cbfb93a54bd2132a6de1f6cec028eee6ef/setuptools-74.1.2.tar.gz"
	pythonsetuptools_cmd_tar := exec.Command("curl", "-L", pythonsetuptools_tar_url, "-o", "package.tar.gz")
	err := pythonsetuptools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pythonsetuptools_zip_url := "https://files.pythonhosted.org/packages/3e/2c/f0a538a2f91ce633a78daaeb34cbfb93a54bd2132a6de1f6cec028eee6ef/setuptools-74.1.2.zip"
	pythonsetuptools_cmd_zip := exec.Command("curl", "-L", pythonsetuptools_zip_url, "-o", "package.zip")
	err = pythonsetuptools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pythonsetuptools_bin_url := "https://files.pythonhosted.org/packages/3e/2c/f0a538a2f91ce633a78daaeb34cbfb93a54bd2132a6de1f6cec028eee6ef/setuptools-74.1.2.bin"
	pythonsetuptools_cmd_bin := exec.Command("curl", "-L", pythonsetuptools_bin_url, "-o", "binary.bin")
	err = pythonsetuptools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pythonsetuptools_src_url := "https://files.pythonhosted.org/packages/3e/2c/f0a538a2f91ce633a78daaeb34cbfb93a54bd2132a6de1f6cec028eee6ef/setuptools-74.1.2.src.tar.gz"
	pythonsetuptools_cmd_src := exec.Command("curl", "-L", pythonsetuptools_src_url, "-o", "source.tar.gz")
	err = pythonsetuptools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pythonsetuptools_cmd_direct := exec.Command("./binary")
	err = pythonsetuptools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
