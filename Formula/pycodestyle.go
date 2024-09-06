package main

// Pycodestyle - Simple Python style checker in one Python file
// Homepage: https://pycodestyle.pycqa.org/

import (
	"fmt"
	
	"os/exec"
)

func installPycodestyle() {
	// Método 1: Descargar y extraer .tar.gz
	pycodestyle_tar_url := "https://github.com/PyCQA/pycodestyle/archive/refs/tags/2.12.1.tar.gz"
	pycodestyle_cmd_tar := exec.Command("curl", "-L", pycodestyle_tar_url, "-o", "package.tar.gz")
	err := pycodestyle_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pycodestyle_zip_url := "https://github.com/PyCQA/pycodestyle/archive/refs/tags/2.12.1.zip"
	pycodestyle_cmd_zip := exec.Command("curl", "-L", pycodestyle_zip_url, "-o", "package.zip")
	err = pycodestyle_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pycodestyle_bin_url := "https://github.com/PyCQA/pycodestyle/archive/refs/tags/2.12.1.bin"
	pycodestyle_cmd_bin := exec.Command("curl", "-L", pycodestyle_bin_url, "-o", "binary.bin")
	err = pycodestyle_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pycodestyle_src_url := "https://github.com/PyCQA/pycodestyle/archive/refs/tags/2.12.1.src.tar.gz"
	pycodestyle_cmd_src := exec.Command("curl", "-L", pycodestyle_src_url, "-o", "source.tar.gz")
	err = pycodestyle_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pycodestyle_cmd_direct := exec.Command("./binary")
	err = pycodestyle_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
