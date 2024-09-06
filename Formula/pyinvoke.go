package main

// Pyinvoke - Pythonic task management & command execution
// Homepage: https://www.pyinvoke.org/

import (
	"fmt"
	
	"os/exec"
)

func installPyinvoke() {
	// Método 1: Descargar y extraer .tar.gz
	pyinvoke_tar_url := "https://files.pythonhosted.org/packages/f9/42/127e6d792884ab860defc3f4d80a8f9812e48ace584ffc5a346de58cdc6c/invoke-2.2.0.tar.gz"
	pyinvoke_cmd_tar := exec.Command("curl", "-L", pyinvoke_tar_url, "-o", "package.tar.gz")
	err := pyinvoke_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pyinvoke_zip_url := "https://files.pythonhosted.org/packages/f9/42/127e6d792884ab860defc3f4d80a8f9812e48ace584ffc5a346de58cdc6c/invoke-2.2.0.zip"
	pyinvoke_cmd_zip := exec.Command("curl", "-L", pyinvoke_zip_url, "-o", "package.zip")
	err = pyinvoke_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pyinvoke_bin_url := "https://files.pythonhosted.org/packages/f9/42/127e6d792884ab860defc3f4d80a8f9812e48ace584ffc5a346de58cdc6c/invoke-2.2.0.bin"
	pyinvoke_cmd_bin := exec.Command("curl", "-L", pyinvoke_bin_url, "-o", "binary.bin")
	err = pyinvoke_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pyinvoke_src_url := "https://files.pythonhosted.org/packages/f9/42/127e6d792884ab860defc3f4d80a8f9812e48ace584ffc5a346de58cdc6c/invoke-2.2.0.src.tar.gz"
	pyinvoke_cmd_src := exec.Command("curl", "-L", pyinvoke_src_url, "-o", "source.tar.gz")
	err = pyinvoke_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pyinvoke_cmd_direct := exec.Command("./binary")
	err = pyinvoke_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
