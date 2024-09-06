package main

// Mypy - Experimental optional static type checker for Python
// Homepage: https://www.mypy-lang.org/

import (
	"fmt"
	
	"os/exec"
)

func installMypy() {
	// Método 1: Descargar y extraer .tar.gz
	mypy_tar_url := "https://files.pythonhosted.org/packages/5c/86/5d7cbc4974fd564550b80fbb8103c05501ea11aa7835edf3351d90095896/mypy-1.11.2.tar.gz"
	mypy_cmd_tar := exec.Command("curl", "-L", mypy_tar_url, "-o", "package.tar.gz")
	err := mypy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mypy_zip_url := "https://files.pythonhosted.org/packages/5c/86/5d7cbc4974fd564550b80fbb8103c05501ea11aa7835edf3351d90095896/mypy-1.11.2.zip"
	mypy_cmd_zip := exec.Command("curl", "-L", mypy_zip_url, "-o", "package.zip")
	err = mypy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mypy_bin_url := "https://files.pythonhosted.org/packages/5c/86/5d7cbc4974fd564550b80fbb8103c05501ea11aa7835edf3351d90095896/mypy-1.11.2.bin"
	mypy_cmd_bin := exec.Command("curl", "-L", mypy_bin_url, "-o", "binary.bin")
	err = mypy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mypy_src_url := "https://files.pythonhosted.org/packages/5c/86/5d7cbc4974fd564550b80fbb8103c05501ea11aa7835edf3351d90095896/mypy-1.11.2.src.tar.gz"
	mypy_cmd_src := exec.Command("curl", "-L", mypy_src_url, "-o", "source.tar.gz")
	err = mypy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mypy_cmd_direct := exec.Command("./binary")
	err = mypy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
