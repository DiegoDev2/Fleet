package main

// Pyflow - Installation and dependency system for Python
// Homepage: https://github.com/David-OConnor/pyflow

import (
	"fmt"
	
	"os/exec"
)

func installPyflow() {
	// Método 1: Descargar y extraer .tar.gz
	pyflow_tar_url := "https://github.com/David-OConnor/pyflow/archive/refs/tags/0.3.1.tar.gz"
	pyflow_cmd_tar := exec.Command("curl", "-L", pyflow_tar_url, "-o", "package.tar.gz")
	err := pyflow_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pyflow_zip_url := "https://github.com/David-OConnor/pyflow/archive/refs/tags/0.3.1.zip"
	pyflow_cmd_zip := exec.Command("curl", "-L", pyflow_zip_url, "-o", "package.zip")
	err = pyflow_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pyflow_bin_url := "https://github.com/David-OConnor/pyflow/archive/refs/tags/0.3.1.bin"
	pyflow_cmd_bin := exec.Command("curl", "-L", pyflow_bin_url, "-o", "binary.bin")
	err = pyflow_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pyflow_src_url := "https://github.com/David-OConnor/pyflow/archive/refs/tags/0.3.1.src.tar.gz"
	pyflow_cmd_src := exec.Command("curl", "-L", pyflow_src_url, "-o", "source.tar.gz")
	err = pyflow_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pyflow_cmd_direct := exec.Command("./binary")
	err = pyflow_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
