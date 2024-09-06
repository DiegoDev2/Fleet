package main

// Pylint - It's not just a linter that annoys you!
// Homepage: https://github.com/pylint-dev/pylint

import (
	"fmt"
	
	"os/exec"
)

func installPylint() {
	// Método 1: Descargar y extraer .tar.gz
	pylint_tar_url := "https://files.pythonhosted.org/packages/cf/e8/d59ce8e54884c9475ed6510685ef4311a10001674c28703b23da30f3b24d/pylint-3.2.7.tar.gz"
	pylint_cmd_tar := exec.Command("curl", "-L", pylint_tar_url, "-o", "package.tar.gz")
	err := pylint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pylint_zip_url := "https://files.pythonhosted.org/packages/cf/e8/d59ce8e54884c9475ed6510685ef4311a10001674c28703b23da30f3b24d/pylint-3.2.7.zip"
	pylint_cmd_zip := exec.Command("curl", "-L", pylint_zip_url, "-o", "package.zip")
	err = pylint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pylint_bin_url := "https://files.pythonhosted.org/packages/cf/e8/d59ce8e54884c9475ed6510685ef4311a10001674c28703b23da30f3b24d/pylint-3.2.7.bin"
	pylint_cmd_bin := exec.Command("curl", "-L", pylint_bin_url, "-o", "binary.bin")
	err = pylint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pylint_src_url := "https://files.pythonhosted.org/packages/cf/e8/d59ce8e54884c9475ed6510685ef4311a10001674c28703b23da30f3b24d/pylint-3.2.7.src.tar.gz"
	pylint_cmd_src := exec.Command("curl", "-L", pylint_src_url, "-o", "source.tar.gz")
	err = pylint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pylint_cmd_direct := exec.Command("./binary")
	err = pylint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
