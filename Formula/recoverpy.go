package main

// Recoverpy - TUI to recover overwritten or deleted data
// Homepage: https://github.com/PabloLec/recoverpy

import (
	"fmt"
	
	"os/exec"
)

func installRecoverpy() {
	// Método 1: Descargar y extraer .tar.gz
	recoverpy_tar_url := "https://files.pythonhosted.org/packages/ea/55/9b4c247c0988b791e25431269acd13ba0b784d84d3b8d46f1775fbf8693b/recoverpy-2.1.9.tar.gz"
	recoverpy_cmd_tar := exec.Command("curl", "-L", recoverpy_tar_url, "-o", "package.tar.gz")
	err := recoverpy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	recoverpy_zip_url := "https://files.pythonhosted.org/packages/ea/55/9b4c247c0988b791e25431269acd13ba0b784d84d3b8d46f1775fbf8693b/recoverpy-2.1.9.zip"
	recoverpy_cmd_zip := exec.Command("curl", "-L", recoverpy_zip_url, "-o", "package.zip")
	err = recoverpy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	recoverpy_bin_url := "https://files.pythonhosted.org/packages/ea/55/9b4c247c0988b791e25431269acd13ba0b784d84d3b8d46f1775fbf8693b/recoverpy-2.1.9.bin"
	recoverpy_cmd_bin := exec.Command("curl", "-L", recoverpy_bin_url, "-o", "binary.bin")
	err = recoverpy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	recoverpy_src_url := "https://files.pythonhosted.org/packages/ea/55/9b4c247c0988b791e25431269acd13ba0b784d84d3b8d46f1775fbf8693b/recoverpy-2.1.9.src.tar.gz"
	recoverpy_cmd_src := exec.Command("curl", "-L", recoverpy_src_url, "-o", "source.tar.gz")
	err = recoverpy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	recoverpy_cmd_direct := exec.Command("./binary")
	err = recoverpy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
